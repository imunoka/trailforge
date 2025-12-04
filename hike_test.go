package trailforge

import "testing"

func TestHike(t *testing.T) {
	h := Hike{
		Name: "Timberline to Paradise Park",
		Location: "Mt Hood, Oregon",
		DistanceMi: 9.8,
		ElevGainFt: 2800,
		Difficulty: Beast,
		Solo: true,
	}

	got := h.String()
	want := "Timberline to Paradise Park (9.8 mi, 2800 ft gain, Beast)"

	if got != want {
		t.Errorf("\ngot %q\nwant %q", got, want)
	}
}
