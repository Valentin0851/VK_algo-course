package main

import (
	"testing"
)

func TestMultiple(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		var x, y, result = 2, 3, 6
		realResult := Mult(x, y)
		if realResult != result {
			t.Errorf("%d != %d", realResult, result)
		}
	})

	t.Run("medium", func(t *testing.T) {
		var x, y, result = 10, 15, 150
		realResult := Mult(x, y)
		if realResult != result {
			t.Errorf("%d != %d", realResult, result)
		}
	})

	t.Run("hard", func(t *testing.T) {
		var x, y, result = 100, 1500, 150000
		realResult := Mult(x, y)
		if realResult != result {
			t.Errorf("%d != %d", realResult, result)
		}
	})
}
