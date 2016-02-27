package math

import (
	"testing"
	m "introduction-to-go/11-packages/01-package/math"
)

func TestAverage(t *testing.T){
	var v float64
	v = m.Average([]float64{1,2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}


