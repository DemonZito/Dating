package ui

import (
	"dating/ui/style"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/new/html/a"
	"qlova.org/seed/new/image"
	"qlova.org/seed/new/page"
	"qlova.org/seed/new/spacer"
	"qlova.org/seed/new/text"
	"qlova.org/seed/set"
	"qlova.org/seed/set/align"
	"qlova.org/seed/set/center"
	"qlova.org/seed/set/transition"
	"qlova.org/seed/use/css/units/rem"
	"qlova.org/seed/use/css/units/vmin"
	"qlova.org/seed/use/html/attr"
	"qlova.tech/rgb"
)

type AboutPage struct{}

func (p AboutPage) Page(r page.Router) seed.Seed {
	var qlovaLink = client.NewString("https://github.com/qlova/seed")

	return page.New(
		transition.Fade(),
		set.Scrollable(),
		set.Color(rgb.LightGray),

		set.If.Medium().Portrait(
			set.Width(vmin.New(100)),
		),

		set.If.Small().Portrait(
			set.Width(vmin.New(100)),
		),

		center.Seeds{
			text.New(style.Text,
				text.SetSize(rem.New(1.5)),
				align.Center(),
				text.Set("This web-app was created by:"),
			),

			a.New(style.Text,
				align.Center(),
				text.SetSize(rem.New(1.5)),
				attr.Set("href", "https://www.madeleineday.nz/"),
				attr.Set("target ", "_blank"),
				text.Set("Madeleine Day"),
			),

			spacer.New(rem.One),

			text.New(style.Text,
				text.SetSize(rem.New(1.5)),
				align.Center(),
				text.Set("using:"),
			),

			image.New(
				set.Width(rem.New(10.0)),

				align.Center(),
				set.Margin(nil, rem.New(0.5)),

				image.Set("logo.svg"),
				client.OnClick(client.Open(qlovaLink)),
			),
		},
	)
}
