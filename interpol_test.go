package interpol

import "testing"

func TestInterpolateInt(t *testing.T) {
	startMin := 0
	startMax := 1000
	endMin := 0
	endMax := 65535
	result := InterpolateInt(39, startMin, startMax, endMin, endMax)
	if result != 2555 {
		t.Errorf("integer interpolation of 0 1000 -> 0 65535 failed")
	}

	startMin = -100
	startMax = 100
	endMin = 100
	endMax = 200
	result = InterpolateInt(-150, startMin, startMax, endMin, endMax)
	if result != 75 {
		t.Errorf("integer interpolation of -100 100 -> 100 200 failed")
	}
}

func TestInterpolateFloat(t *testing.T) {
	startMin := 3201.0
	startMax := 100200.0
	endMin := 10.0
	endMax := 11.0
	result := InterpolateFloat(50200.0, startMin, startMax, endMin, endMax)
	if result != 10.484530768358436 {
		t.Errorf("float interpolation of 3201.0 100200.0 -> 10.0 11.0 failed")
	}
	startMin = -100.0
	startMax = 100.0
	endMin = 100.0
	endMax = 200.0
	result = InterpolateFloat(-150.0, startMin, startMax, endMin, endMax)
	if result != 75.0 {
		t.Errorf("float interpolation of -100.0 100.0 -> 100.0 200.0 failed")
	}
}
