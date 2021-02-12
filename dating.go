package dating

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"qlova.org/seed/client"
	"qlova.org/seed/use/js"
)

type Holiday struct {
	ID          string
	Name        string
	Image       string
	Distance    string
	DisplayTime string
	IsCustom    string

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
			DisplayTime: t.String(),
			nextTime:    func() time.Time { return t },
		})
	}
}

var Custom = []Holiday{}

func constructTime(date time.Time, hours string) time.Time {
	var timecomps = strings.Split(hours, ":")
	hour, _ := strconv.Atoi(timecomps[0])
	minute, _ := strconv.Atoi(timecomps[1])
	return time.Date(date.Year(), date.Month(), date.Day(), hour, minute, 0, 0, time.Local)
}

func AddCustom(name string, date time.Time, hours string) {
	fmt.Println(hours)
	var datetime = constructTime(date, hours)

	Custom = append(Custom, Holiday{
		Name:        name,
		Image:       "https://loremflickr.com/400/400" + "?lock=1",
		Time:        datetime,
		DisplayTime: datetime.String(),
		nextTime:    func() time.Time { return datetime },
		IsCustom:    "True",
	})
}

func DeleteCustom(sid string) client.Script {
	var id, _ = strconv.Atoi(sid)
	Custom = append(Custom[:id], Custom[id+1:]...)
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

func update(h []Holiday) {
	for i := range h {
		if h[i].nextTime == nil {
			t := h[i].Time
			h[i].nextTime = func() time.Time { return t }
		}
		h[i].distance = h[i].nextTime().Sub(time.Now())
		h[i].Distance = formatTime(h[i].distance)
	}

	sort.Slice(h, func(i, j int) bool {
		return h[i].distance < h[j].distance
	})
}

func GetHolidays() []Holiday {
	update(Holidays)

	return Holidays
}

func GetCustom() []Holiday {
	update(Custom)

	return Custom
}

func SaveCustom() client.Script {
	b, err := json.Marshal(Custom)
	if err != nil {
		panic(err)
	}

	return js.Func("window.localStorage.setItem").Run(client.NewString("custom.dates"), client.NewString(string(b)))
}

func LoadCustom(custom string) {
	err := json.Unmarshal([]byte(custom), &Custom)
	if err != nil {
		fmt.Println(err)
	}
}
