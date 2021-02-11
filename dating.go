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
			Image:       "https://loremflickr.com/500/500/" + hol.Name + "?lock=1",
			DisplayTime: t.String(),
			nextTime:    func() time.Time { return t },
		})
	}
}

var Custom = []Holiday{}

func AddCustom(name string, date time.Time, hours string) {
	Custom = append(Custom, Holiday{
		Name:        name,
		Image:       "https://picsum.photos/100?" + fmt.Sprint(time.Now().UnixNano()),
		Time:        date,
		DisplayTime: date.String(),
		nextTime:    func() time.Time { return date },
		IsCustom:    "True",
	})
}

func DeleteCustom(sid string) client.Script {
	var id, _ = strconv.Atoi(sid)
	Custom = append(Custom[:id], Custom[id+1:]...)
	return SaveCustom()
}

func update(h []Holiday) {
	for i := range h {
		if h[i].nextTime == nil {
			t := h[i].Time
			h[i].nextTime = func() time.Time { return t }
		}
		h[i].distance = h[i].nextTime().Sub(time.Now())
		h[i].Distance = fmt.Sprintf("%v days", int(math.Ceil(h[i].distance.Hours()/24)))
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
