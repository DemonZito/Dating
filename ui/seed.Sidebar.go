package ui

import (
	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/new/column"
	"qlova.org/seed/new/image"
	"qlova.org/seed/new/page"
	"qlova.org/seed/new/row"
	"qlova.org/seed/set"
	"qlova.org/seed/set/align"
	"qlova.org/seed/set/center"
	"qlova.org/seed/use/css/units/percentage/of"
	"qlova.org/seed/use/css/units/rem"
	"qlova.tech/rgb"
)

func NewSidebar() seed.Seed {
	var col = column.New()

	return col.With(
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
		set.Color(rgb.Black),
		set.MinWidth(rem.New(5.0)),
		set.Height(100%of.Parent),

		/*
			text.New(style.Text,
				text.Set("DatingApp"),
				text.SetColor(rgb.White),
				text.SetSize(rem.New(2.0)),
				text.Center(),
				align.Center(),

				client.OnClick(page.RouterOf(col).Goto(AboutPage{})),
			),
		*/
		center.This(
			image.New(
				set.Width(rem.New(3.0)),

				align.Center(),
				set.Margin(rem.New(0.5)),

				image.Set("house.svg"),
				client.OnClick(page.RouterOf(col).Goto(AboutPage{})),
			),

			image.New(
				set.Width(rem.New(3.0)),

				align.Center(),
				set.Margin(rem.New(0.5)),

				image.Set("round-star.svg"),
				client.OnClick(page.RouterOf(col).Goto(PopularPage{})),
			),

			image.New(
				set.Width(rem.New(3.0)),

				align.Center(),
				set.Margin(rem.New(0.5)),

				image.Set("pencil.svg"),
				client.OnClick(page.RouterOf(col).Goto(CustomPage{})),
			),

			image.New(
				set.Width(rem.New(3.0)),

				align.Center(),
				set.Margin(rem.New(0.5)),

				image.Set("heart.svg"),
				client.OnClick(page.RouterOf(col).Goto(AddPage{})),
			),
		),
	)
}
