package challenges

func NbYear(p0 int, percent float64, aug int, p int) (years int) {
	for p0 < p {
		p0 += int(float64(p0)*(percent/100)) + aug
		years += 1
	}

	return
}