[English](README.md) | 简体中文

# go-android-uiautomator

用go开发android自动化，所以需要学点go的语法再继续往下。

项目参考于[openatx/uiautomator2](https://github.com/openatx/uiautomator2)，项目中的dump方式实际就是借用了大佬的[u2.jar](https://public.uiauto.devsleep.com/u2jar)，输入法是从[senzhk/ADBKeyBoard](https://github.com/senzhk/ADBKeyBoard)修改的，这里将本项目使用到的库和项目列出，若有遗漏，请联系我，我会第一时间补上并致以歉意。

- [![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/beevik/etree.svg?label=etree)](https://github.com/beevik/etree/releases)
- [![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/sevlyar/go-daemon.svg?label=go-daemon)](https://github.com/sevlyar/go-daemon/releases)

向大佬致敬。

目前还处于开发中，所以使用上有什么问题或bug还请大家提交一下Issues，我只要不出意外，那肯定是会看到的( •͈ᴗ⁃͈)ᓂ- - -♡

还有就是文档所涉及演示的APP为**抖音v32.0.0**，文档有些没有讲到的请在使用时查看源码。

## 准备工作

1. 准备安卓手机

2. 开发阶段可选

   - 安装Python，再通过pip安装[uiauto.dev](https://uiauto.devsleep.com/)，用于查看页面结构

     > ```shell
     > pip install -U uiautodev
     > 
     > uiauto.dev
     > ```

   - 连接电脑开启USB调试

   - 使用无线调试

> [!NOTE]
>
> 首次的话会自动安装输入法[star-ime](https://github.com/shi-yunsheng/star-ime)，如果弹出安装提示，请选择允许。

## 快速开始
```shell
go get github.com/shi-yunsheng/driver
```


```go
import "github.com/shi-yunsheng/driver"

func main() {
    d := driver.New()
    
    d.Connect("192.168.10.128") // 通过adb devices查看设备的序列号或设备IP

	d.StartApp("com.ss.android.ugc.aweme") // 启动抖音

	searchBtn, _ := d.WaitElement(driver.By{Selector: driver.ContentDesc, Value: "搜索"})
	if searchBtn != nil {
		searchBtn.Tap()
	}

	dy := "ismeSYS"
	msg := "哈喽，老石"

	search, _ := d.WaitElement(driver.By{Selector: driver.ResourceID, Value: "com.ss.android.ugc.aweme:id/et_search_kw"})
	if search != nil {
		search.Input(dy)
		search.Search()
	}

	user, _ := d.WaitElement(driver.By{Selector: driver.EndsWithText, Value: fmt.Sprintf("抖音号：%s，按钮", dy)})
	if user != nil {
		user.Tap()
	}

	more, _ := d.WaitElement(driver.By{Selector: driver.ContentDesc, Value: "更多"})
	if more != nil {
		more.Tap()
	}

	sendBtn, _ := d.WaitElement(driver.By{Selector: driver.Text, Value: "发私信"})
	if sendBtn != nil {
		sendBtn.Tap()
	}

	message, _ := d.WaitElement(driver.By{Selector: driver.Text, Value: "发送消息"})
	if message != nil {
		message.Input(msg)
		message.Send()

		for i := 0; i < 5; i++ {
			d.Back()
		}
	}

	defer d.Cleanup()
}
```



## API Document

### New()

创建一个`driver`，用于驱动整个自动化工作。

> ```go
> func main() {
>    	d := driver.New()
> }
> ```

### Connect()

用于在PC端开发测试时连接设备使用，如果是有线方式，参数则是对应的序列号，如果通过无线方式，参数可使用IP。

> [!NOTE]
>
> 如果编译到安卓执行，需要注释或删除掉~~~

> ```go
> d.Connect("192.168.10.128") // WLAN
> // or USB d.Connect("22e3e987") 
> ```

### Info()

获取设备信息。

> ```go
> deviceInfo := d.Info()
> fmt.Println(deviceInfo)
> ```

### StartApp()

启动应用，传递应用包名。

> ```go
> d.StartApp("com.ss.android.ugc.aweme") // 启动抖音
> ```

### StopApp()

停止应用，传递应用包名。

> ```go
> d.StopApp("com.ss.android.ugc.aweme")
> ```

### RestartApp()

重启应用，传递应用包名。

> ```go
> d.RestartApp("com.ss.android.ugc.aweme")
> ```

### InstallApp()

安装应用，传递安装包路径。

> ```go
> d.InstallApp("/data/local/tmp/douyin.apk")
> ```

### UninstallApp()

卸载应用，传递应用包名。

> ```go
> d.UninstallApp("com.ss.android.ugc.aweme")
> ```

### Document()

获取当前页面结构。

> ```go
> doc := d.Document()
> fmt.Println(doc.RawXML)
> ```

### FindElement()

查找指定元素节点，如果节点不存在，返回`nil`，目前只能通过`xpath`方式查找。

> ```go
> doc := d.Document()
> search := doc.FindElement(`//node[@resource-id="com.ss.android.ugc.aweme:id/et_search_kw"]`)
> if search != nil {
>    	search.Input("老石话不多")
>    	search.Search()
> }
> ```

### FindElements()

同`FindElement()`，只是`FindElements()`返回的是一组元素节点。

> ```go
> doc := d.Document()
> navbar := doc.FindElements(`//node[@resource-id="com.ss.android.ugc.aweme:id/0zg"]`)
> for _, bar := range navbar {
>    	fmt.Println(bar.Text())
> }
> ```

### WaitElement()

等待元素节点出现，接受 `By` 类型参数，包含选择器类型、值和超时时间(毫秒)。如果超时返回 `nil` 和错误。

支持的选择器类型:
- Text: 通过文本内容匹配
- ContentDesc: 通过 content-desc 属性匹配  
- Class: 通过类名匹配
- ResourceID: 通过资源ID匹配
- StartsWithText: 通过文本前缀匹配
- EndsWithText: 通过文本后缀匹配
- StartsWithContentDesc: 通过 content-desc 前缀匹配
- EndsWithContentDesc: 通过 content-desc 后缀匹配
- StartsWithClass: 通过类名前缀匹配
- EndsWithClass: 通过类名后缀匹配
- StartsWithResourceID: 通过资源ID前缀匹配
- EndsWithResourceID: 通过资源ID后缀匹配

> ```go
> d := driver.New()
> 
> search, _ := d.WaitElement(driver.By{
>    	Selector: driver.ResourceID,
>    	Value: "com.ss.android.ugc.aweme:id/et_search_kw",
>    })
>    
> if search != nil {
>    	search.Input("老石话不多")
>    	search.Search()
> }
> ```

### Text()

获取元素节点的`text`属性值，没有返回空`""`。

> ```go
> doc := d.Document()
> likes := doc.FindElement(`//node[@resource-id="com.ss.android.ugc.aweme:id/ffl"]`)
> if likes != nil {
>    	fmt.Println(likes.Text())
> }
> ```

### ContentDesc()

获取元素节点的`content-desc`属性值，没有返回空`""`。

> ```go
> doc := d.Document()
> likes := doc.FindElement(`//node[@resource-id="com.ss.android.ugc.aweme:id/ffl"]`)
> if likes != nil {
>    	fmt.Println(likes.ContentDesc())
> }
> ```

### GetBounds()

获取元素节点的`bounds`属性值，没有返回`nil`。

> ```go
> doc := d.Document()
> likes := doc.FindElement(`//node[@resource-id="com.ss.android.ugc.aweme:id/ffl"]`)
> if likes != nil {
>    	fmt.Println(likes.ContentDesc())
> }
> ```

### GetAttribute()

如果还有其他没有给出的属性，可是该方法获取。

> ```go
> doc := d.Document()
> like := doc.FindElement(`//node[@resource-id="com.ss.android.ugc.aweme:id/ffl"]`)
> if like != nil {
>    	fmt.Println(likes.GetAttribute("resource-id"))
> }
> ```

### Tap()

+ `driver` `Tap()`

  传入要点击的坐标（x,y）进行点击。

  > ```go
  > d := driver.New()
  > d.Tap(100, 200) // 点击屏幕坐标(100, 200)处
  > ```

+ `element` `Tap()`

  直接点击该元素。

  > ```go
  > doc := d.Document()
  > like := doc.FindElement(`//node[@resource-id="com.ss.android.ugc.aweme:id/ffl"]`)
  > if likes != nil {
  >    	like.Tap() // 点击该元素
  > }
  > ```

> [!NOTE]
>
> 坐标是相对于屏幕左上角，并且坐标的取值范围为x: [0, 屏幕分辨率宽度]，y: [0, 屏幕分辨率高度]

### LongTap()

+ `driver` `LongTap()`

  传入要长按的坐标（x,y）进行长按。

  > ```go
  > d := driver.New()
  > d.LongTap(400, 600) // 长按屏幕坐标(400, 600)处
  > ```

+ `element` `LongTap()`

  直接长按该元素。

  > ```go
  > doc := d.Document()
  > chat := doc.FindElement(`//node[@resource-id="com.ss.android.ugc.aweme:id/wvq"][@class="android.widget.Button"]`)
  > if chat != nil {
  >    	chat.LongTap()
  > }
  > ```

### Swipe()

+ `driver` `Swipe()`

  根据传入的滑动方向在屏幕进行滑动。

  > ```go
  > d := driver.New()
  > d.Swipe(driver.SWIPE_DOWN) // 往下滑动
  > ```

+ `element` `Swipe()`

  根据传入的滑动方向在元素的范围内进行滑动。

  > ```go
  > doc := d.Document()
  > tabbar := doc.FindElement(`//node[@resource-id="com.ss.android.ugc.aweme:id/vf5"]`)
  > if tabbar != nil {
  > 	tabbar.Swipe(driver.SWIPE_RIGHT) // 往右滑动
  > }
  > ```

### Input()

+ `driver` `Input()`

  传入输入框的坐标输入文本。

  > ```go
  > d := driver.New()
  > d.Input(400, 600, "老石")
  > ```

+ `element` `Input()`

  如果元素节点是可输入，会直接在该元素输入文本。

  > ```go
  > d := driver.New()
  > search, _ := d.WaitElement(driver.By{
  >    	Selector: driver.ResourceID,
  >    	Value: "com.ss.android.ugc.aweme:id/et_search_kw",
  > })
  > 
  > if search != nil {
  >    	search.Input("老石话不多")
  >    	search.Search()
  > }
  > ```

### Screenshot()

+ `driver` `Screenshot()`

  对当前屏幕进行截图。

  > ```go
  > img := d.Screenshot()
  > ```
  
+ `element` `Screenshot()`

  对节点进行截图。

  > ```go
  > d := driver.New()
  > avatar, _ := d.WaitElement(driver.By{
  >    	Selector: driver.ResourceID,
  >    	Value: "com.ss.android.ugc.aweme:id/user_avatar",
  > })
  > 
  > if avatar != nil {
  >        avatar.Screenshot()
  >    }
  > ```

> [!CAUTION]
>
> 如果节点被挡住，那么截图只会有未被挡住的部分。

### Run()

执行安卓shell命令。

> ```go
> if ls, err := d.Run("ls"); err != nil {
>     fmt.Println(ls)
> }
> ```

### Cleanup()

清理操作，一般用在程序完成时。

> ```go
> defer d.Cleanup()
> ```
