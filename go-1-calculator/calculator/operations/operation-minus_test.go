package operations_test

import (
	"gb/go-1-11/calculator/operations"
	"testing"
)

func TestOperationMinus(t *testing.T) {
	minus := operations.NewOperationMinus()

	for _, ex := range []struct {
		X      float64
		Y      float64
		Result float64
		Err    error
	}{
		{100, 50, 50, nil},
		{100, 100, 0, nil},
		{200, 0, 200, nil},
	} {
		err := minus.Init(ex.X, ex.Y)
		if err != nil {
			t.Fatalf("could not init %f - %f", ex.X, ex.Y)
		}

		result, err := minus.Calc()
		if err != ex.Err {
			t.Fatalf("could not calc %f - %f", ex.X, ex.Y)
		}

		if result != ex.Result {
			t.Fatalf("could not calc %f- %f. Expected %f got %f", ex.X, ex.Y, ex.Result, result)
		}
	}

}
