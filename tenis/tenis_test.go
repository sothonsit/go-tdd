package tenis

import "testing"

func TestPointA0andB0WhenAGetPointThenPointAShouldBe15AndPointBShouldBe0(t *testing.T) {
	tenis := Tenis()
	resultA, resultB := tenis.AGetPoint()
	if resultA != 15 {
		t.Errorf("resultA should be %v but %v", 15, resultA)
	}
	if resultB != 0 {
		t.Errorf("resultB should be %v but %v", 0, resultB)
	}
}
