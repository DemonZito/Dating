package dating

import "qlova.org/seed/use/wasm"

func init() {
	wasm.Export(AddCustom)
	wasm.Export(GetHolidays)
	wasm.Export(GetCustom)
	wasm.Export(SaveCustom)
	wasm.Export(LoadCustom)
	wasm.Export(DeleteCustom)
	wasm.Export(DeleteExpired)
	wasm.Export(GetExpired)
	wasm.Export(DownloadCustom)
	wasm.Export(LoadReader)
	wasm.Export(DownloadPopular)
}
