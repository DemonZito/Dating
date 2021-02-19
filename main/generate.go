// +build generate wasm

//go:generate go run -tags generate generate.go
package main

import (
	_ "dating"
	_ "time/tzdata"

	"qlova.org/seed/use/wasm"
)

func main() { wasm.Generate() }
