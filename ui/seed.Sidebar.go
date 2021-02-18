package ui

import (
	"dating"
	"dating/ui/style"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/screen"
	"qlova.org/seed/new/column"
	"qlova.org/seed/new/expander"
	"qlova.org/seed/new/filepicker"
	"qlova.org/seed/new/image"
	"qlova.org/seed/new/page"
	"qlova.org/seed/new/row"
	"qlova.org/seed/new/spacer"
	"qlova.org/seed/new/text"
	"qlova.org/seed/new/text/rich"
	"qlova.org/seed/set"
	"qlova.org/seed/set/align"
	"qlova.org/seed/set/change"
	"qlova.org/seed/set/visible"
	"qlova.org/seed/use/css/units/percentage/of"
	"qlova.org/seed/use/css/units/rem"
	"qlova.tech/rgb"
)

func NewSidebar() seed.Seed {
	var sidebar = row.New()

	NewOption := func(title rich.Text, icon string, onclick client.Script) seed.Seed {
		return row.New(
			client.OnClick(onclick),

			set.Margin(nil, rem.One),
			change.When(screen.TinyToSmall^screen.Landscape,
				set.Margin(nil, nil),
			),
			image.New(
				set.Width(rem.New(2.0)),

				align.Center(),

				change.When(screen.Not(screen.TinyToSmall^screen.Landscape),
					set.Margin(rem.New(0.5), nil),
				),

				image.Set(icon),
			),

			text.New(style.Text,
				change.When(screen.TinyToSmall^screen.Portrait,
					set.Hidden(),
				),
				text.SetSize(rem.New(1.5)),
				text.SetColor(rgb.Black),
				align.Center(),
				text.Set(title),
			),
		)
	}

	return sidebar.With(
		change.When(screen.TinyToSmall^screen.Portrait,
			row.Set(),
			set.Width(100%of.Parent),
			set.Height(rem.New(5.0)),
		),

		change.When(screen.TinyToSmall^screen.Landscape,
			set.Height(100%of.Parent),
		),

		change.When(screen.MediumToHuge,
			set.Width(rem.New(37.0)),
			set.MaxWidth(25%of.Parent),
			set.Height(100%of.Parent),
		),

		set.Gradient{From: rgb.Hex("ffafbd"), To: rgb.Hex("ffc3a0"), Direction: 1i, Side: false},

		visible.When(screen.MediumToHuge, expander.New()),

		column.New(
			change.When(screen.TinyToSmall^screen.Portrait,
				align.Center(),
				row.Set(),
				set.Width(100%of.Parent),
			),
			set.Width(rem.New(15.0)),
			change.When(screen.TinyToSmall^screen.Landscape,
				set.Width(rem.New(10.0)),
			),

			// Logo
			image.New(
				set.Margin(rem.One),
				set.Width(rem.New(8.0)),
				image.Set("eternal-love.svg?v1"),

				change.When(screen.TinyToSmall,
					set.Width(rem.New(4.0)),
				),
				change.When(screen.Medium,
					set.Width(rem.New(6.0)),
				),
			),

			visible.When(screen.MediumToHuge,
				spacer.New(rem.New(2.0)),
			),

			// Icons
			expander.New(change.When(screen.MediumToHuge, set.Hidden())),
			NewOption("Home", "home.svg", page.RouterOf(sidebar).Goto(AboutPage{})),
			NewOption("Popular", "popular.svg", page.RouterOf(sidebar).Goto(PopularPage{})),
			NewOption("Custom", "calendar.svg", page.RouterOf(sidebar).Goto(CustomPage{})),
			NewOption("Add", "health-normal.svg", page.RouterOf(sidebar).Goto(AddPage{})),
			expander.New(change.When(screen.TinyToSmall^screen.Portrait, set.Hidden())),
			NewOption("Export", "save.svg", client.Download(dating.DownloadCustom)),
			NewOption("Import", "open-folder.svg", filepicker.SelectFile(func(f filepicker.File) client.Script {
				return client.Run(dating.LoadReader, f)
			})),
			expander.New(change.When(screen.MediumToHuge, set.Hidden())),
		),
		//expander.New(),
	)
}
