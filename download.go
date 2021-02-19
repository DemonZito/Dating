// +build wasm

package dating

import (
	"fmt"
	"net/http"
	"syscall/js"
	"time"
)

var Host = js.Global().Get("location").Get("host").String()
var Protocol = js.Global().Get("location").Get("protocol").String()

var Location *time.Location

func SetTimeLocation() {
	location, err := time.LoadLocation(js.Global().Get("Intl").Call("DateTimeFormat").Call("resolvedOptions").Get("timeZone").String())
	if err == nil {
		Location = location
	} else {
		fmt.Println("Invalid Location", err)
	}
}

func DownloadPopular() {
	res, err := http.Get(Protocol + "//" + Host + "/assets/holidays.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	readPopular(res.Body)
}
