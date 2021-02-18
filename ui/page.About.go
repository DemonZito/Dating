package ui

import (
	"dating/md"
	"dating/ui/style"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/screen"
	"qlova.org/seed/new/image"
	"qlova.org/seed/new/markdown"
	"qlova.org/seed/new/page"
	"qlova.org/seed/set"
	"qlova.org/seed/set/align"
	"qlova.org/seed/set/change"
	"qlova.org/seed/set/transition"
	"qlova.org/seed/use/css/units/percentage/of"
	"qlova.org/seed/use/css/units/rem"
	"qlova.tech/rgb"
)

type AboutPage struct{}

func (p AboutPage) Page(r page.Router) seed.Seed {
	var qlovaLink = client.NewString("https://github.com/qlova/seed")

	return page.New(
		transition.Fade(),
		set.Scrollable(),
		set.Color(rgb.White),

		align.Center(),

		change.When(screen.LargeToHuge,
			set.PaddingTop(rem.New(10)),
			set.MaxWidth(50%of.Parent),
		),

		markdown.New(style.Text,

			set.Margin(rem.One, nil),
			markdown.Set(md.Welcome),
		),

		image.New(
			set.Width(rem.New(10.0)),

			set.Margin(rem.One, rem.New(0.5)),

			image.Set("logo.svg"),
			client.OnClick(client.Open(qlovaLink)),
		),
	)
}
