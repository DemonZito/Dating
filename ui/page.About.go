package ui

import (
	"dating"
	"dating/md"
	"dating/ui/style"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/new/filepicker"
	"qlova.org/seed/new/image"
	"qlova.org/seed/new/markdown"
	"qlova.org/seed/new/page"
	"qlova.org/seed/new/row"
	"qlova.org/seed/set"
	"qlova.org/seed/set/align"
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
		set.Color(rgb.White),
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

		row.New(
			align.Center(),
			image.New(
				set.Width(rem.New(5.0)),

				set.Margin(rem.One, rem.New(0.5)),

				image.Set("save.svg"),
				client.OnClick(client.Download(dating.DownloadCustom)),
			),
			image.New(
				set.Width(rem.New(5.0)),

				set.Margin(rem.One, rem.New(0.5)),

				image.Set("open-folder.svg"),
				client.OnClick(filepicker.SelectFile(func(f filepicker.File) client.Script {
					return client.Run(dating.LoadReader, f)
				})),
			),
		),
	)
}
