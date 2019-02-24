package makna_test

import (
	"fmt"
	"testing"

	"github.com/rompi/makna"
)

func TestCastString(t *testing.T) {
	t.Run("Success convert int32 to string", func(t *testing.T) {
		var x = int32(5)
		newX := makna.ToString(x)
		if newX != "5" {
			t.Errorf("casting int to string")
		}
	})

	t.Run("Success convert int64 to string", func(t *testing.T) {
		var x = int64(777777777777)
		newX := makna.ToString(x)
		if newX != "777777777777" {
			t.Errorf("casting int to string")
		}
	})

	t.Run("Success convert float32 to string", func(t *testing.T) {
		var x = float32(5.4)
		newX := makna.ToString(x)
		if newX != "5.4" {
			t.Errorf("casting float32 to string")
		}
	})

	t.Run("Success convert float64 to string", func(t *testing.T) {
		var x = float64(5555555555.4)
		newX := makna.ToString(x)
		fmt.Println(newX)
		if newX != "5555555555.40" {
			t.Errorf("casting float64 to string")
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

func TestCastInt32E(t *testing.T) {
	t.Run("Success convert string to int32 with error", func(t *testing.T) {
		var x = "5"
		newX, err := makna.ToInt32E(x)
		if newX != int32(5) || err != nil {
			t.Errorf("casting string to int32 with error")
		}
	})

	t.Run("Success convert int64 to int32 with error", func(t *testing.T) {
		var x = int64(5)
		newX, err := makna.ToInt32E(x)
		if newX != int32(5) || err != nil {
			t.Errorf("casting int64 to int32 with error")
		}
	})

	t.Run("Success convert float32 to int32", func(t *testing.T) {
		var x = float32(5.0)
		newX, err := makna.ToInt32E(x)
		if newX != 5 || err != nil {
			t.Errorf("casting float32 to int32 with error")
		}
	})

	t.Run("Success convert float64 to int32", func(t *testing.T) {
		var x = float64(5.0)
		newX, err := makna.ToInt32E(x)
		if newX != 5 || err != nil {
			t.Errorf("casting float64 to int32 with error")
		}
	})

	t.Run("Error convert boolean with error", func(t *testing.T) {
		var x = true
		_, err := makna.ToInt32E(x)
		if err == nil {
			t.Errorf("casting boolean to int32 with error")
		}
	})

	t.Run("Error convert string to int32 with error", func(t *testing.T) {
		var x = "828975897248578924"
		_, err := makna.ToInt32E(x)
		if err == nil {
			t.Errorf("casting random string to int32 with error")
		}
	})

	t.Run("Error convert string to int32 with error", func(t *testing.T) {
		var x = "a3"
		_, err := makna.ToInt32E(x)
		if err == nil {
			t.Errorf("casting random string to int32 with error")
		}
	})
}

func TestCastInt32(t *testing.T) {
	t.Run("Success convert string to int32", func(t *testing.T) {
		var x = "5"
		newX := makna.ToInt32(x)
		if newX != int32(5) {
			t.Errorf("casting string to int32")
		}
	})

	t.Run("Success convert int64 to int32", func(t *testing.T) {
		var x = int64(5)
		newX := makna.ToInt32(x)
		if newX != int32(5) {
			t.Errorf("casting int64 to int32")
		}
	})

	t.Run("Success convert float32 to int32", func(t *testing.T) {
		var x = float32(5.0)
		newX := makna.ToInt32(x)
		if newX != 5 {
			t.Errorf("casting float32 to int32")
		}
	})

	t.Run("Success convert float64 to int32", func(t *testing.T) {
		var x = float64(5.0)
		newX := makna.ToInt32(x)
		if newX != 5 {
			t.Errorf("casting float64 to int32")
		}
	})

	t.Run("Error convert boolean", func(t *testing.T) {
		var x = true
		newX := makna.ToInt32(x)
		if newX != 0 {
			t.Errorf("casting boolean to int32")
		}
	})

	t.Run("Error convert overflow string to int32", func(t *testing.T) {
		var x = "828975897248578924"
		newX := makna.ToInt32(x)
		if newX != 2147483647 {
			t.Errorf("casting overflow string to int32")
		}
	})

	t.Run("Error convert random string to int32", func(t *testing.T) {
		var x = "a3"
		newX := makna.ToInt32(x)
		if newX != 0 {
			t.Errorf("casting random string to int32")
		}
	})
}
