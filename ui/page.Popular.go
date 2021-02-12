package ui

import (
	"dating"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/new/feed"
	"qlova.org/seed/new/page"
	"qlova.org/seed/set"
	"qlova.org/seed/set/transition"
	"qlova.org/seed/use/js/window"
	"qlova.tech/rgb"
)

type PopularPage struct{}

func (p PopularPage) Page(r page.Router) seed.Seed {
	var holidays = feed.With(dating.GetHolidays)

	return page.New(
		transition.Fade(),

		set.Scrollable(),
		page.OnEnter(holidays.Refresh()),
		set.Color(rgb.LightGray),
		NewHolidays(holidays),

		client.OnLoad(window.SetInterval(holidays.Refresh().GetScript(), client.NewFloat64(500))),
	)
}
