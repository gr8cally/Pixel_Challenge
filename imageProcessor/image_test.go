package imageProcessor

import (
	"io/fs"
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

	t.Run("Silver dataset", func(t *testing.T) {

		// Arrange
		image := "0c9ec855-f5ad-4586-bf3b-da489f447219.raw"
		category := "../Silver"
		files, err := ioutil.ReadDir(category)
		if err != nil {
			t.Fatal(err)
		}
		got := Compare(image, files, category)
		want := map[string]float64{
			"07820a55-9d59-453d-bb71-826418bc2874.raw": 6.2963,
			"85fb2c82-ffb8-44d2-80ca-73de8f94d8e6.raw": 6.2940,
			"bcd0ffed-a032-4ab6-aaab-6addec4ebe6b.raw": 6.2997,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	//t.Run("Gold dataset", func(t *testing.T) {
	//
	//	// Arrange
	//	image := "0a0f8f44-3b78-4bff-adee-14bc708e4ba7.raw"
	//	category := "../Gold"
	//	files, err := ioutil.ReadDir(category)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	got := Compare(image, files, category)
	//	want := map[string]float64{
	//		"70e70084-9e53-4af8-a195-8fcb395aed78.raw": 6.3288,
	//		"7155ec2a-2e9c-4b78-845a-fdd955d3aa9b.raw": 6.3228,
	//		"c7f36f26-0b1f-4d1c-ba7c-cc9be0b3eaf0.raw": 6.3216,
	//	}
	//
	//	if !reflect.DeepEqual(got, want) {
	//		t.Errorf("got %v, wanted %v", got, want)
	//	}
	//})
}

func benchmarkCompare(image, category string, files []fs.FileInfo, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Compare(image, files, category)
	}
}

func BenchmarkCompare(b *testing.B) {
	image := "1d25ea94-4562-4e19-848e-b60f1b58deee.raw"
	category := "../Bronze"
	files, err := ioutil.ReadDir(category)
	if err != nil {
		b.Fatal(err)
	}
	benchmarkCompare(image, category, files, b)
}

func BenchmarkCompare2(b *testing.B) {
	image := "0c9ec855-f5ad-4586-bf3b-da489f447219.raw"
	category := "../Silver"
	files, err := ioutil.ReadDir(category)
	if err != nil {
		b.Fatal(err)
	}
	benchmarkCompare(image, category, files, b)
}

func BenchmarkCompare3(b *testing.B) {
	image := "0a0f8f44-3b78-4bff-adee-14bc708e4ba7.raw"
	category := "../gold"
	files, err := ioutil.ReadDir(category)
	if err != nil {
		b.Fatal(err)
	}
	benchmarkCompare(image, category, files, b)
}
