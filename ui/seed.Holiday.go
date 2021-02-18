package ui

import (
	"dating"
	"dating/ui/style"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientfmt"
	"qlova.org/seed/client/screen"
	"qlova.org/seed/new/column"
	"qlova.org/seed/new/expander"
	"qlova.org/seed/new/feed"
	"qlova.org/seed/new/image"
	"qlova.org/seed/new/row"
	"qlova.org/seed/new/text"
	"qlova.org/seed/set"
	"qlova.org/seed/set/change"
	"qlova.org/seed/set/visible"
	"qlova.org/seed/use/css/units/rem"
	"qlova.org/seed/use/js"
	"qlova.tech/rgb"
)

func NewHolidays(f *feed.Feed) seed.Seed {

	var holiday dating.Holiday
	f.Into(&holiday)

	return f.New(
		row.New(style.Border,
			set.Height(rem.New(10.0)),
			set.Margin(rem.One, rem.One/2),
			set.Color(rgb.White),
			set.Clipped(),

			set.OnHover(set.Color(rgb.Hex("f7f9fa"))),

			image.New(
				set.Width(rem.New(10.0)),
				change.When(screen.TinyToMedium^screen.Portrait,
					set.Width(rem.New(5.0)),
				),

				image.Crop(),

				image.SetTo(f.String(holiday.Image)),
			),
			column.New(
				text.New(style.Text,
					text.SetSize(rem.New(1.5)),
					change.When(screen.Small,
						text.SetSize(rem.New(1.25)),
					),

					set.Padding(rem.New(2.0), nil),
					change.When(screen.TinyToMedium^screen.Portrait,
						set.Padding(rem.New(1.0), nil),
					),

					text.SetStringTo(clientfmt.Sprintf("%v until %v",
						f.String(holiday.Distance),
						f.String(holiday.Name))),

					change.When(f.String(holiday.IsExpired),
						text.SetStringTo(clientfmt.Sprintf("%v",
							f.String(holiday.Name))),
					),
				),

				text.New(style.Text,
					text.SetSize(rem.New(1.0)),
					change.When(screen.Small,
						text.SetSize(rem.New(1.0)),
					),
					set.Padding(rem.New(2.0), rem.New(1.0)),
					change.When(screen.TinyToMedium^screen.Portrait,
						set.Padding(rem.New(1.0), rem.New(1.0)),
					),

					text.SetStringTo(f.String(holiday.DisplayTime)),
				),
			),

			expander.New(),
			visible.When(f.String(holiday.IsCustom),
				image.New(
					set.Width(rem.New(5.0)),
					change.When(screen.TinyToMedium^screen.Portrait,
						set.Width(rem.New(2.5)),
					),
					set.Margin(nil, rem.One, rem.One, rem.One),
					image.Set("cancel.svg"),

					change.When(f.String(holiday.IsExpired),
						image.Set("confirmed.svg"),
					),

					client.OnClick(
						client.If(f.String(holiday.IsExpired),
							client.Run(dating.DeleteExpired, js.String{f.Data.Index.GetValue().Call("toString")}),
						).Else(
							client.Run(dating.DeleteCustom, js.String{f.Data.Index.GetValue().Call("toString")}),
						),
						f.Refresh(),
					),
				),
			),
		),
	)
}
