package ui

import (
	"dating"
	"dating/ui/style"
	"time"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/if/not"
	"qlova.org/seed/client/poll"
	"qlova.org/seed/new/feed"
	"qlova.org/seed/new/page"
	"qlova.org/seed/new/text"
	"qlova.org/seed/set"
	"qlova.org/seed/set/align"
	"qlova.org/seed/set/transition"
	"qlova.org/seed/set/visible"
	"qlova.org/seed/use/css/units/rem"
	"qlova.org/seed/use/css/units/vmin"
	"qlova.tech/rgb"
)

type CustomPage struct{}

func (p CustomPage) Page(r page.Router) seed.Seed {
	var holidays = feed.With(dating.GetCustom)
	var expiredHolidays = feed.With(dating.GetExpired)

	return page.New(
		transition.Fade(),

		set.Scrollable(),

		set.If.Medium().Portrait(
			set.Width(vmin.New(100)),
		),

		set.If.Small().Portrait(
			set.Width(vmin.New(100)),
		),

		page.OnEnter(holidays.Refresh()),
		set.Color(rgb.LightGray),

		poll.Every(time.Second/2, client.NewScript(
			holidays.Refresh().GetScript(),
			expiredHolidays.Refresh().GetScript(),
		)),

		NewHolidays(holidays),

		visible.When(not.True(expiredHolidays.Empty),
			text.New(style.Text,
				text.Set("Expired"),
				align.Center(),
				text.Center(),
				text.SetSize(rem.New(2.0)),
			),
		),

		NewHolidays(expiredHolidays),
	)
}
