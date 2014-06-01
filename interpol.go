package interpol

func minf(a, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}
func maxf(a, b float64) float64 {
	if a < b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func max(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//interpolate between two ranges
func InterpolateFloat(x, oMin, oMax, nMin, nMax float64) float64 {

	//check reversed input range
	reverseInput := false
	oldMin := minf(oMin, oMax)
	oldMax := maxf(oMin, oMax)
	if oldMin != oMin {
		reverseInput = true
	}

	//check reversed output range
	reverseOutput := false
	newMin := minf(nMin, nMax)
	newMax := maxf(nMin, nMax)
	if newMin != nMin {
		reverseOutput = true
	}

	portion := (x - oldMin) * (newMax - newMin) / (oldMax - oldMin)
	if reverseInput {
		portion = (oldMax - x) * (newMax - newMin) / (oldMax - oldMin)
	}

	result := portion + newMin
	if reverseOutput == true {
		result = newMax - portion
	}

	return result

}
func InterpolateInt(x, oMin, oMax, nMin, nMax int) int {

	//check reversed input range
	reverseInput := false
	oldMin := min(oMin, oMax)
	oldMax := max(oMin, oMax)
	if oldMin != oMin {
		reverseInput = true
	}

	//check reversed output range
	reverseOutput := false
	newMin := min(nMin, nMax)
	newMax := max(nMin, nMax)
	if newMin != nMin {
		reverseOutput = true
	}

	portion := (x - oldMin) * (newMax - newMin) / (oldMax - oldMin)
	if reverseInput {
		portion = (oldMax - x) * (newMax - newMin) / (oldMax - oldMin)
	}

	result := portion + newMin
	if reverseOutput == true {
		result = newMax - portion
	}

	return result

}
