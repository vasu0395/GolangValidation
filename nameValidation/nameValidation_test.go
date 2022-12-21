package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameValidationBasedOnLength(t *testing.T){
	t.Run("should return true when valid name passed",func(t *testing.T) {
		assert.Equal(t,true,isValidNameBasedOnLength("vasu chauhan"))
	})

	t.Run("should return false when invalid name passed",func(t *testing.T) {
		assert.Equal(t,false,isValidNameBasedOnLength("   in"))
	})
}

func TestNameValidationBasedOnEmoji(t *testing.T){
	t.Run("Should return false when name doesn't contain emoji",func(t *testing.T) {
		assert.Equal(t,false,isEmojiPresent("vasu chauhan"))
	})

	t.Run("Should return true when name contain emoji",func(t *testing.T) {
		assert.Equal(t,true,isEmojiPresent("vasu 😊 chauhan"))
	})
}