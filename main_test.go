package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestBoxing(t *testing.T) {
	t.Run("20 cakes and 25 apples", func(t *testing.T) {
		countBoxes, countEachBox := boxing(20, 25)
		assert.Equal(t, countBoxes, 20)
		assert.Equal(t, countEachBox, 1)
	})

	t.Run("30 cakes and 10 apples", func(t *testing.T) {
		countBoxes, countEachBox := boxing(30, 10)
		assert.Equal(t, countBoxes, 10)
		assert.Equal(t, countEachBox, 1)
	})

	t.Run("20 cakes and 20 apples", func(t *testing.T) {
		countBoxes, countEachBox := boxing(20, 20)
		assert.Equal(t, countBoxes, 10)
		assert.Equal(t, countEachBox, 2)
	})

	t.Run("27 cakes and 27 apples", func(t *testing.T) {
		countBoxes, countEachBox := boxing(27, 27)
		assert.Equal(t, countBoxes, 9)
		assert.Equal(t, countEachBox, 3)
	})
}
