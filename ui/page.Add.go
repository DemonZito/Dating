package ui

import (
	"dating"
	"dating/ui/style"

	"qlova.org/seed"
	"qlova.org/seed/client"
	"qlova.org/seed/client/clientside"
	"qlova.org/seed/client/if/the"
	"qlova.org/seed/new/button"
	"qlova.org/seed/new/datebox"
	"qlova.org/seed/new/expander"
	"qlova.org/seed/new/hourbox"
	"qlova.org/seed/new/page"
	"qlova.org/seed/new/row"
	"qlova.org/seed/new/spacer"
	"qlova.org/seed/new/text"
	"qlova.org/seed/new/textbox"
	"qlova.org/seed/set"
	"qlova.org/seed/set/align"
	"qlova.org/seed/set/center"
	"qlova.org/seed/set/transition"
	"qlova.org/seed/use/css/units/percentage/of"
	"qlova.org/seed/use/css/units/rem"
	"qlova.org/seed/use/js"
	"qlova.tech/rgb"
)

type AddPage struct{}

func (p AddPage) Page(r page.Router) seed.Seed {
	var name = new(clientside.String)
	var date = new(clientside.Time)
	var hours = new(clientside.Duration)

	return page.New(
		transition.Fade(),

		textbox.New(style.Text, style.Border,
			set.MinWidth(rem.New(30)),
			set.Width(50%of.Parent),
			set.MaxWidth(100%of.Parent),
			set.Margin(nil, rem.New(0.5)),
			align.Center(),

			text.SetSize(rem.New(3.0)),
			text.Center(),

			textbox.SetPlaceholder("Name of Event"),
			textbox.Update(name),
		),

		spacer.New(rem.One*3),

		row.New(
			center.This(
				text.New(style.Text,
					align.Center(),
					text.Set("Date:  "),
				),
				datebox.New(style.Text,
					datebox.Update(date),

					//client.OnLoad(date.Set(time.Now())),
				),

				spacer.New(rem.One*2),

				text.New(style.Text,
					align.Center(),
					text.Set("Time:  "),
				),
				hourbox.New(style.Text,
					hourbox.Update(hours),
				),
			),
		),

		expander.New(),

		row.New(
			expander.New(),
			button.New(style.Text, style.Border,
				set.Color(rgb.Chartreuse),
				set.Margin(rem.New(0.5)),
				set.Padding(rem.New(2.0), rem.One),
				text.SetSize(rem.New(2.0)),

				text.Set("DONE"),

				client.OnClick(
					client.If(js.NewValue("(%v + %v > Date.parse(new Date()))", date, hours),
						client.Run(dating.AddCustom, name, the.Time(date, hours)),
						client.Run(dating.SaveCustom),
						r.Goto(CustomPage{})),
				),
			),
		),
	)
}
