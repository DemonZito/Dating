package ui

import (
	"dating"
	"dating/ui/style"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/new/circle"
	"qlova.org/seed/new/column"
	"qlova.org/seed/new/filepicker"
	"qlova.org/seed/new/image"
	"qlova.org/seed/new/page"
	"qlova.org/seed/new/spacer"
	"qlova.org/seed/new/text"
	"qlova.org/seed/set"
	"qlova.org/seed/set/align"
	"qlova.org/seed/set/center"
	"qlova.org/seed/set/change"
	"qlova.org/seed/set/transition"
	"qlova.org/seed/set/visible"
	"qlova.org/seed/use/css/units/px"
	"qlova.org/seed/use/css/units/rem"
	"qlova.tech/rgb"
)

func NewToolbar() seed.Seed {
	var MenuOpen = new(clientside.Bool)

	return column.New(
		set.Overlay(set.Bottom, set.Right),
		set.OffsetRight(rem.One),

		change.When(page.Is(AddPage{}), set.Hidden()),

		visible.When(MenuOpen,
			column.New(set.Color(rgb.LightGray),
				set.Padding(nil, rem.One),
				transition.Fade(),

				set.Border(set.Solid),
				set.BorderColor(rgb.Black),
				set.BorderWidth(px.One),

				text.New(style.Text, align.Center(), text.Set("Export"), client.OnClick(client.Download(dating.DownloadCustom))),
				spacer.New(rem.One),
				text.New(style.Text, align.Center(), text.Set("Import"), client.OnClick(filepicker.SelectFile(func(f filepicker.File) client.Script {
					return client.Run(dating.LoadReader, f)
				}))),
			),
		),

		circle.New(circle.Set(rem.New(3)),
			set.Margin(rem.One),
			set.Color(style.PrimaryColor),

			set.Border(set.Solid),
			set.BorderColor(rgb.Black),
			set.BorderWidth(px.One),

			center.This(
				image.New(align.Center(), image.Set("tinker.svg"), set.Size(rem.New(2), rem.New(2)), client.OnClick(MenuOpen.Toggle())),
			),
		),
	)
}
