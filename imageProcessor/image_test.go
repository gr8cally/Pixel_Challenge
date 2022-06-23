package imageProcessor

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestCompare(t *testing.T) {
	t.Run("Bronze dataset", func(t *testing.T) {

		// Arrange
		image := "1d25ea94-4562-4e19-848e-b60f1b58deee.raw"
		category := "../Bronze"
		files, err := ioutil.ReadDir(category)
		if err != nil {
			t.Fatal(err)
		}
		got := Compare(image, files, category)
		want := map[string]float64{
			"4424fa5a-c00d-4cd5-8525-fcf921b09ca8.raw": 6.2790,
			"6c9952ef-e5bf-4de2-817b-fd0073be8449.raw": 6.2518,
			"e3c342e2-2429-4f47-8828-f7ee0703ad38.raw": 6.2544,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
