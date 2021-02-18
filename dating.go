package dating

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"time"

	"qlova.org/seed/client"
	"qlova.org/seed/client/clientrpc"
	"qlova.org/seed/client/clientsafe"
	"qlova.org/seed/use/js"
)

type Holiday struct {
	ID          string
	Name        string
	Image       string
	Distance    string
	DisplayTime string
	IsCustom    string
	IsExpired   string

	Time time.Time `mirror:"ignore"`

	distance time.Duration
	nextTime func() time.Time
}

var Holidays = []Holiday{}

func readPopular(r io.Reader) {
	var rawHolidays []HolidayJSON
	var err = json.NewDecoder(r).Decode(&rawHolidays)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, hol := range rawHolidays {
		t := time.Date(time.Now().Year(), time.Month(hol.Month), hol.Day, 0, 0, 0, 0, time.Local)

		// If date has already passed, increment the year
		if time.Now().After(t) {
			t = t.AddDate(1, 0, 0)
		}
		if hol.Generate {
			continue
		}
		Holidays = append(Holidays, Holiday{
			Name:        hol.Name,
			Time:        t,
			Image:       "https://loremflickr.com/400/400/" + hol.Name + "?lock=1",
			DisplayTime: t.Format("3:04 pm - _2 Jan 2006"),
			nextTime:    func() time.Time { return t },
		})
	}
}

var Custom = []Holiday{}
var Expired = []Holiday{}

func AddCustom(name string, date time.Time) error {

	date = date.Local()

	if time.Since(date) > 0 {
		return clientsafe.Err(fmt.Errorf("Past Date Error"), "Cannot select a date in the past!")
	}

	fmt.Println(time.Since(date))

	Custom = append(Custom, Holiday{
		Name:        name,
		Image:       "https://loremflickr.com/400/400/" + name + "?lock=1",
		Time:        date,
		DisplayTime: date.Format("3:04 pm - _2 Jan 2006"),
		nextTime:    func() time.Time { return date },
		IsCustom:    "True",
	})

	return nil
}

func DeleteCustom(sid string) client.Script {
	var id, _ = strconv.Atoi(sid)
	Custom = append(Custom[:id], Custom[id+1:]...)
	return SaveCustom()
}

func DeleteExpired(sid string) client.Script {
	var id, _ = strconv.Atoi(sid)
	Expired = append(Expired[:id], Expired[id+1:]...)
	return SaveCustom()
}

func formatTime(dist time.Duration) string {
	var days = int(math.Floor(dist.Hours() / 24))
	var hours = dist.Hours()
	var minutes = dist.Minutes()
	var seconds = dist.Seconds()

	// Format in days and hours
	if days >= 1 {
		return fmt.Sprintf("%v days, %v hours", days, int(math.Ceil(hours))%24)
	} else if hours >= 1 {
		return fmt.Sprintf("%v hours, %v minutes", int(math.Floor(hours)), int(math.Ceil(minutes))%60)
	} else if minutes >= 1 {
		return fmt.Sprintf("%v minutes, %v seconds", int(math.Floor(minutes)), int(math.Ceil(seconds))%60)
	} else {
		return fmt.Sprintf("%v seconds", int(math.Ceil(seconds)))
	}
}

func update(p *[]Holiday, expired bool) {
	var overwrite = 0
	var h = *p

	for i := range h {
		if h[i].nextTime == nil {
			t := h[i].Time
			h[i].nextTime = func() time.Time { return t }
		}

		h[i].distance = h[i].nextTime().Sub(time.Now())

		if !expired && h[i].distance < 0 {
			h[i].IsExpired = "true"
			Expired = append(Expired, h[i])
			continue
		}
		h[i].Distance = formatTime(h[i].distance)

		h[overwrite] = h[i]
		overwrite++
	}

	*p = h[:overwrite]
	h = *p

	sort.Slice(h, func(i, j int) bool {
		return h[i].distance < h[j].distance
	})
}

func GetHolidays() []Holiday {
	update(&Holidays, true)

	return Holidays
}

func GetCustom() []Holiday {
	update(&Custom, false)

	return Custom
}

func GetExpired() []Holiday {
	update(&Expired, true)

	return Expired
}

func SaveCustom() client.Script {
	var SaveDates struct {
		Custom  []Holiday
		Expired []Holiday
	}
	SaveDates.Custom = Custom
	SaveDates.Expired = Expired

	b, err := json.Marshal(SaveDates)
	if err != nil {
		panic(err)
	}

	return js.Func("window.localStorage.setItem").Run(client.NewString("custom.dates"), client.NewString(string(b)))
}

func LoadCustom(custom string) {
	var SaveDates struct {
		Custom  []Holiday
		Expired []Holiday
	}
	err := json.Unmarshal([]byte(custom), &SaveDates)
	if err != nil {
		fmt.Println(err)
	}

	Custom = SaveDates.Custom
	Expired = SaveDates.Expired
}

func DownloadCustom(request clientrpc.Request) {
	request.SetHeader("Content-Disposition", "filename=CustomDates.json")

	var SaveDates struct {
		Custom  []Holiday
		Expired []Holiday
	}
	SaveDates.Custom = Custom
	SaveDates.Expired = Expired

	json.NewEncoder(request.Writer()).Encode(SaveDates)
}

func LoadReader(r io.Reader) client.Script {
	b, _ := io.ReadAll(r)
	LoadCustom(string(b))
	return SaveCustom()
}
