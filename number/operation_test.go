package number

import "testing"

func TestAddOperationAdd1To2ShouldResult3(t *testing.T) {
	op := newOperation()
	result := op.Add(1, 2)
	if result != 3 {
		t.Errorf("result should be %v but %v", 3, result)
	}
}

func TestAddOperationAddMinus1To2ShouldResult1(t *testing.T) {
	op := newOperation()
	result := op.Add(-1, 2)
	if result != 1 {
		t.Errorf("result should be %v but %v", 1, result)
	}
}

func TestAddOperationAdd1ToMinus2ShouldResultMinus1(t *testing.T) {
	op := newOperation()
	result := op.Add(1, -2)
	if result != -1 {
		t.Errorf("result should be %v but %v", -1, result)
	}
}
