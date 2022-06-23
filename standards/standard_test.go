package standards

import "testing"

func TestGetPercent(t *testing.T) {
	got := GetPercent(15, 30)
	want := 50.0

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestContains(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		got := Contains(make([]float64, 0), 3.4)
		want := false

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	})

	t.Run("element only of slice", func(t *testing.T) {
		got := Contains([]float64{3.4}, 3.4)
		want := true

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	})

	t.Run("element in beginning of slice", func(t *testing.T) {
		got := Contains([]float64{3.4, 3.6, 5.7}, 3.4)
		want := true

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	})

	t.Run("element in end of slice", func(t *testing.T) {
		got := Contains([]float64{3.4, 3.6, 5.7, 66.4}, 66.4)
		want := true

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	})

	t.Run("element not in slice", func(t *testing.T) {
		got := Contains([]float64{3.4, 3.6, 5.7}, 13.4)
		want := false

		if got != want {
			t.Errorf("got %t, wanted %t", got, want)
		}
	})
}
