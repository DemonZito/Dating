package ui

import (
	"dating"
	"dating/ui/style"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/screen"
	"qlova.org/seed/new/expander"
	"qlova.org/seed/new/feed"
	"qlova.org/seed/new/page"
	"qlova.org/seed/new/row"
	"qlova.org/seed/new/text"
	"qlova.org/seed/set"
	"qlova.org/seed/set/transition"
	"qlova.org/seed/set/visible"
	"qlova.org/seed/use/css/units/percentage/of"
	"qlova.org/seed/use/css/units/rem"
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
		set.Color(rgb.White),

		visible.When(screen.TinyToSmall^screen.Portrait,
			row.New(set.Width(100%of.Parent),
				text.New(style.Text, text.Center(),
					text.SetSize(rem.New(1.3)),
					expander.Set(),
					text.Set("Popular"),
				),
				text.New(style.Text, text.Center(),
					text.SetSize(rem.New(1.3)),
					expander.Set(),
					text.Set("Custom"),

					client.OnClick(r.Goto(CustomPage{})),

					set.Color(rgb.LightGray),
				),
			),
		),

		NewHolidays(holidays),

		client.OnLoad(window.SetInterval(holidays.Refresh().GetScript(), client.NewFloat64(500))),
	)
}
