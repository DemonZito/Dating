package ui

import (
	"dating/ui/style"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/new/column"
	"qlova.org/seed/new/image"
	"qlova.org/seed/new/page"
	"qlova.org/seed/new/row"
	"qlova.org/seed/new/spacer"
	"qlova.org/seed/new/text"
	"qlova.org/seed/new/text/rich"
	"qlova.org/seed/set"
	"qlova.org/seed/set/align"
	"qlova.org/seed/use/css/units/percentage/of"
	"qlova.org/seed/use/css/units/rem"
	"qlova.tech/rgb"
)

func NewSidebar() seed.Seed {
	var col = column.New()

	var grad = set.Gradient{From: rgb.Hex("ffafbd"), To: rgb.Hex("ffc3a0"), Direction: 1i, Side: false}

	NewOption := func(title rich.Text, icon string, nav page.Page) seed.Seed {
		return row.New(
			client.OnClick(page.RouterOf(col).Goto(nav)),
			set.Margin(rem.New(10.0), rem.One, nil, nil),

			image.New(
				set.Width(rem.New(2.0)),

				align.Center(),
				set.Margin(rem.New(0.5)),

				image.Set(icon),
			),

			text.New(style.Text,
				set.If.Medium().Portrait(
					set.Hidden(),
				),
				set.If.Small().Portrait(
					set.Hidden(),
				),
				set.If.Tiny().Portrait(
					set.Hidden(),
				),

				text.SetSize(rem.New(1.5)),
				text.SetColor(rgb.Black),
				align.Center(),
				text.Set(title),
			),

			spacer.New(rem.New(3.0)),
		)
	}

	SideBar := col.With(
		set.If.Medium().Portrait(
			row.Set(),
			set.Width(100%of.Parent),
			set.Height(rem.New(5.0)),
		),
		set.If.Small().Portrait(
			row.Set(),
			set.Width(100%of.Parent),
			set.Height(rem.New(5.0)),
		),
		set.If.Tiny().Portrait(
			row.Set(),
			set.Width(100%of.Parent),
			set.Height(rem.New(5.0)),
		),
		set.MinWidth(rem.New(37.0)),
		set.Height(100%of.Parent),

		image.New(
			image.Set("eternal-love.svg"),
			set.Width(rem.New(4.0)),
			set.Margin(rem.New(10.0), rem.One, nil, nil),
		),

		spacer.New(rem.New(2.0)),
		NewOption("Home", "home.svg", AboutPage{}),
		NewOption("Popular", "popular.svg", PopularPage{}),
		NewOption("Custom", "calendar.svg", CustomPage{}),
		NewOption("Add", "health-normal.svg", AddPage{}),

		//expander.New(),
	)

	grad.AddTo(SideBar)
	return SideBar
}
