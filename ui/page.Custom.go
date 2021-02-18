package ui

import (
	"dating"
	"dating/ui/style"
	"time"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/if/not"
	"qlova.org/seed/client/poll"
	"qlova.org/seed/client/screen"
	"qlova.org/seed/new/expander"
	"qlova.org/seed/new/feed"
	"qlova.org/seed/new/page"
	"qlova.org/seed/new/row"
	"qlova.org/seed/new/text"
	"qlova.org/seed/set"
	"qlova.org/seed/set/align"
	"qlova.org/seed/set/change"
	"qlova.org/seed/set/transition"
	"qlova.org/seed/set/visible"
	"qlova.org/seed/use/css/units/percentage/of"
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

		change.When(screen.TinyToSmall^screen.Portrait,
			set.Width(vmin.New(100)),
		),

		page.OnEnter(holidays.Refresh()),
		set.Color(rgb.White),

		poll.Every(time.Second/2, client.NewScript(
			holidays.Refresh().GetScript(),
			expiredHolidays.Refresh().GetScript(),
		)),

		visible.When(screen.TinyToSmall^screen.Portrait,
			row.New(set.Width(100%of.Parent),
				text.New(style.Text, text.Center(),
					text.SetSize(rem.New(1.3)),
					expander.Set(),
					text.Set("Popular"),

					client.OnClick(r.Goto(PopularPage{})),

					set.Color(rgb.LightGray),
				),
				text.New(style.Text, text.Center(),
					text.SetSize(rem.New(1.3)),
					expander.Set(),
					text.Set("Custom"),
				),
			),
		),

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
