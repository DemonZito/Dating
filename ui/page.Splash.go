package ui

import (
	"qlova.org/seed"
	"qlova.org/seed/new/page"
	"qlova.org/seed/set"
	"qlova.org/seed/set/transition"
	"qlova.tech/rgb"
)

type SplashPage struct{}

func (p SplashPage) Page(r page.Router) seed.Seed {
	return page.New(
		transition.Fade(),
		set.Color(rgb.Black),
	)
}
