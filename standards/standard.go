package standards

func Contains(s []float64, e float64) bool {
	if len(s) == 0 {
		return false
	}

	if len(s) == 1 {
		return s[0] == e
	}

	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetPercent(value, totalValue int) float64 {
	percent := (float64(value) / float64(totalValue)) * 100
	return percent
}
