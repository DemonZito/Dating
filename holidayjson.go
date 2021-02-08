package dating

type HolidayJSON struct {
	Name     string `json:"name"`
	Month    int    `json:"month"`
	Day      int    `json:"day"`
	Generate bool   `json:"generate,omitempty"`
}
