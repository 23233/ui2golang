package main

import (
	"go-android-uiauto/driver"
)

func main() {
	d := driver.New()
	err := d.Connect("jjx8496hc6nntcqs")
	if err != nil {
		return
	}

	// fmt.Println(d.Info())

	doc := d.Document()
	if doc != nil {
		search := doc.FindElement(`//node[@resource-id="com.ss.android.ugc.aweme:id/et_search_kw"]`)
		if search != nil {
			search.Input("抖音")
		}
	}
	// for {
	// 	d.Swipe(driver.SWIPE_UP)
	// 	time.Sleep(time.Second * 3)
	// }

	d.Screenshot()
}
