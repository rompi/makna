package makna_test

import (
	"testing"

	"github.com/rompi/makna"
)

func TestCastString(t *testing.T) {
	t.Run("Success convert integer to string", func(t *testing.T) {
		var x = 5
		newX := makna.ToString(x)
		if newX != "5" {
			t.Errorf("casting int to string")
		}
	})

	t.Run("Success convert boolean to string", func(t *testing.T) {
		var x = true
		newX := makna.ToString(x)
		if newX != "true" {
			t.Errorf("casting boolean to string")
		}
	})
}

func TestCastInt32(t *testing.T) {
	t.Run("Success convert string to int32 with error", func(t *testing.T) {
		var x = "5"
		newX, err := makna.ToInt32E(x)
		if newX != int32(5) || err != nil {
			t.Errorf("casting string to int32")
		}
	})

	t.Run("Error convert string to int32 with error", func(t *testing.T) {
		var x = "a5"
		newX, err := makna.ToInt32E(x)
		if newX != int32(0) || err == nil {
			t.Errorf("casting string to int32")
		}
	})

	t.Run("Success convert string to int32", func(t *testing.T) {
		var x = "5"
		newX := makna.ToInt32(x)
		if newX != int32(5) {
			t.Errorf("casting string to int32")
		}
	})
}
