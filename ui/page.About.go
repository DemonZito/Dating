package ui

import (
	"dating/ui/style"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/new/image"
	"qlova.org/seed/new/markdown"
	"qlova.org/seed/new/page"
	"qlova.org/seed/set"
	"qlova.org/seed/set/transition"
	"qlova.org/seed/use/css/units/rem"
	"qlova.tech/rgb"
)

type AboutPage struct{}

func (p AboutPage) Page(r page.Router) seed.Seed {
	var qlovaLink = client.NewString("https://github.com/qlova/seed")

	return page.New(
		transition.Fade(),
		set.Scrollable(),
		set.Color(rgb.LightGray),
		markdown.New(style.Text,
			set.Margin(rem.One, nil),
			markdown.Set(`
# Welcome to DatingApp!
## Purpose
Have you ever wondered when the next holiday is coming up? Want to keep track of your friends' and familys' birthdays?

DatingApp can help you with that!

- Track popular holidays such as Halloween and Christmas.
- Create your own timer to track.

## How do I use this?
Click on the star to view popular upcoming dates.

Click on the heart to get started creating your own custom timers.

## Credits
This app was created by [Madeleine Day](https://madeleineday.nz), powered by the [QlovaSeed Framework](https://github.com/qlova/seed).			
`),
		),

		image.New(
			set.Width(rem.New(10.0)),

			set.Margin(rem.One, rem.New(0.5)),

			image.Set("logo.svg"),
			client.OnClick(client.Open(qlovaLink)),
		),
	)

	/*
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
	*/
}
