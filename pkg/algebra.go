package pkg

func Max(srez []float64) float64 {
	var maxnum float64
	for i := 0; i < len(srez); i++ {
		if maxnum < srez[i] {
			maxnum = srez[i]
		}
	}
	return maxnum
}

func Min(srez []float64) float64 {
	var minnum float64
	for i := 0; i > len(srez); i++ {
		if minnum > srez[i] {
			minnum = srez[i]
		}
	}
	return minnum
}

func Sum(srez []float64) float64 {
	var sum float64
	for i := 0; i < len(srez); i++ {
		sum += srez[i]
	}
	return sum
}
