package main

import (
	"fmt"
	"go-android-uiauto/driver"
	"time"
)

func main() {
	d := driver.New()
	err := d.Connect("jjx8496hc6nntcqs")
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(d.Info())

	// doc := d.Document()
	// if doc != nil {
	// 	r := doc.FindElement(`//node[@resource-id="com.ss.android.ugc.aweme:id/qjs"]`)
	// 	if r != nil {
	// 		r.Swipe(driver2.SWIPE_LEFT)
	// 	}
	// }
	for {
		d.Swipe(driver.SWIPE_UP)
		time.Sleep(time.Second * 3)
	}
}
