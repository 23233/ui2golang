English | [简体中文](README_CN.md)

# go-android-uiautomator

Android automation development using Go, so you need to learn some Go syntax before continuing.

This project references [openatx/uiautomator2](https://github.com/openatx/uiautomator2). The dump method actually uses [u2.jar](https://public.uiauto.devsleep.com/u2jar) from the original project, and the input method is modified from [senzhk/ADBKeyBoard](https://github.com/senzhk/ADBKeyBoard). Here I list the libraries and projects used in this project. If there are any omissions, please contact me, and I will add them immediately with apologies.

- [![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/antchfx/xmlquery.svg?label=xmlquery)](https://github.com/antchfx/xmlquery/releases)
- [![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/sevlyar/go-daemon.svg?label=go-daemon)](https://github.com/sevlyar/go-daemon/releases)

Respect to the developers.

The project is still under development, so if you encounter any problems or bugs, please submit Issues. I will definitely see them as long as nothing unexpected happens ( •͈ᴗ⁃͈)ᓂ- - -♡

Also, the documentation demonstrates using **Douyin v32.0.0** app. For features not covered in the documentation, please refer to the source code.

## Preparation

1. Prepare an Android phone

2. Optional during development

   - Install Python, then install [uiauto.dev](https://uiauto.devsleep.com/) via pip to view page structure

     > ```shell
     > pip install -U uiautodev
     > 
     > uiauto.dev
     > ```

   - Connect to computer and enable USB debugging

   - Use wireless debugging

> [!NOTE]
>
> For first-time use, star-ime input method will be automatically installed. If an installation prompt appears, please allow it.

## Quick Start

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

Create a `driver` to drive the entire automation work.

> ```go
> func main() {
>    	d := driver.New()
> }
> ```

### Connect()

Used to connect to devices when developing and testing on PC. If using wired connection, the parameter is the corresponding serial number; if using wireless connection, the parameter can be the IP address.

> [!NOTE]
>
> If compiling to run on Android, this needs to be commented out or deleted~~~

> ```go
> d.Connect("192.168.10.128") // WLAN
> // or USB d.Connect("22e3e987") 
> ```

### Info()

Get device information.

> ```go
> deviceInfo := d.Info()
> fmt.Println(deviceInfo)
> ```

### StartApp()

Launch an application by passing the package name.

> ```go
> d.StartApp("com.ss.android.ugc.aweme") // Launch Douyin
> ```

### StopApp()

Stop an application by passing the package name.

> ```go
> d.StopApp("com.ss.android.ugc.aweme")
> ```

### RestartApp()

Restart an application by passing the package name.

> ```go
> d.RestartApp("com.ss.android.ugc.aweme")
> ```

### InstallApp()

Install an application by passing the installation package path.

> ```go
> d.InstallApp("/data/local/tmp/douyin.apk")
> ```

### UninstallApp()

Uninstall an application by passing the package name.

> ```go
> d.UninstallApp("com.ss.android.ugc.aweme")
> ```

### Document()

Get the current page structure.

> ```go
> doc := d.Document()
> fmt.Println(doc.RawXML)
> ```

### FindElement()

Find a specific element node. If the node doesn't exist, returns `nil`. Currently only supports finding through `xpath`.

> ```go
> doc := d.Document()
> search := doc.FindElement(`//node[@resource-id="com.ss.android.ugc.aweme:id/et_search_kw"]`)
> if search != nil {
>    	search.Input("老石话不多")
>    	search.Search()
> }
> ```

### FindElements()

Similar to `FindElement()`, but `FindElements()` returns a group of element nodes.

> ```go
> doc := d.Document()
> navbar := doc.FindElements(`//node[@resource-id="com.ss.android.ugc.aweme:id/0zg"]`)
> for _, bar := range navbar {
>    	fmt.Println(bar.Text())
> }
> ```

### WaitElement()

Wait for an element node to appear. Accepts a `By` type parameter containing selector type, value, and timeout duration (milliseconds). Returns `nil` and error if timeout occurs.

Supported selector types:
- Text: Match by text content
- ContentDesc: Match by content-desc attribute
- Class: Match by class name
- ResourceID: Match by resource ID
- StartsWithText: Match by text prefix
- EndsWithText: Match by text suffix
- StartsWithContentDesc: Match by content-desc prefix
- EndsWithContentDesc: Match by content-desc suffix
- StartsWithClass: Match by class name prefix
- EndsWithClass: Match by class name suffix
- StartsWithResourceID: Match by resource ID prefix
- EndsWithResourceID: Match by resource ID suffix

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

Get the `text` attribute value of an element node. Returns empty string `""` if not found.

> ```go
> doc := d.Document()
> likes := doc.FindElement(`//node[@resource-id="com.ss.android.ugc.aweme:id/ffl"]`)
> if likes != nil {
>    	fmt.Println(likes.Text())
> }
> ```

### ContentDesc()

Get the `content-desc` attribute value of an element node. Returns empty string `""` if not found.

> ```go
> doc := d.Document()
> likes := doc.FindElement(`//node[@resource-id="com.ss.android.ugc.aweme:id/ffl"]`)
> if likes != nil {
>    	fmt.Println(likes.ContentDesc())
> }
> ```

### GetBounds()

Get the `bounds` attribute value of an element node. Returns `nil` if not found.

> ```go
> doc := d.Document()
> likes := doc.FindElement(`//node[@resource-id="com.ss.android.ugc.aweme:id/ffl"]`)
> if likes != nil {
>    	fmt.Println(likes.ContentDesc())
> }
> ```

### GetAttribute()

If there are other attributes not provided, you can use this method to get them.

> ```go
> doc := d.Document()
> like := doc.FindElement(`//node[@resource-id="com.ss.android.ugc.aweme:id/ffl"]`)
> if like != nil {
>    	fmt.Println(likes.GetAttribute("resource-id"))
> }
> ```

### Tap()

+ `driver` `Tap()`

  Pass in coordinates (x,y) to tap at that location.

  > ```go
  > d := driver.New()
  > d.Tap(100, 200) // Tap at screen coordinates (100, 200)
  > ```

+ `element` `Tap()`

  Directly tap the element.

  > ```go
  > doc := d.Document()
  > like := doc.FindElement(`//node[@resource-id="com.ss.android.ugc.aweme:id/ffl"]`)
  > if likes != nil {
  >    	like.Tap() // Tap this element
  > }
  > ```

> [!NOTE]
>
> Coordinates are relative to the screen's top-left corner, and the coordinate range is x: [0, screen width], y: [0, screen height]

### LongTap()

+ `driver` `LongTap()`

  Pass in coordinates (x,y) to long press at that location.

  > ```go
  > d := driver.New()
  > d.LongTap(400, 600) // Long press at screen coordinates (400, 600)
  > ```

+ `element` `LongTap()`

  Directly long press the element.

  > ```go
  > doc := d.Document()
  > chat := doc.FindElement(`//node[@resource-id="com.ss.android.ugc.aweme:id/wvq"][@class="android.widget.Button"]`)
  > if chat != nil {
  >    	chat.LongTap()
  > }
  > ```

### Swipe()

+ `driver` `Swipe()`

  Swipe on the screen according to the input direction.

  > ```go
  > d := driver.New()
  > d.Swipe(driver.SWIPE_DOWN) // Swipe down
  > ```

+ `element` `Swipe()`

  Swipe within the element's bounds according to the input direction.

  > ```go
  > doc := d.Document()
  > tabbar := doc.FindElement(`//node[@resource-id="com.ss.android.ugc.aweme:id/vf5"]`)
  > if tabbar != nil {
  > 	tabbar.Swipe(driver.SWIPE_RIGHT) // Swipe right
  > }
  > ```

### Input()

+ `driver` `Input()`

  Input text at the specified coordinates.

  > ```go
  > d := driver.New()
  > d.Input(400, 600, "老石")
  > ```

+ `element` `Input()`

  If the element node is input-capable, directly input text into that element.

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

  Take a screenshot of the current screen.

  > ```go
  > img := d.Screenshot()
  > ```
  
+ `element` `Screenshot()`

  Take a screenshot of the node.

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
> If the node is obscured, the screenshot will only show the unobscured parts.

### Run()

Execute Android shell commands.

> ```go
> if ls, err := d.Run("ls"); err != nil {
>     fmt.Println(ls)
> }
> ```

### Cleanup()

Cleanup operations, generally used when the program completes.

> ```go
> defer d.Cleanup()
> ```
