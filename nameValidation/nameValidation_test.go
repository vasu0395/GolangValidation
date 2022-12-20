package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameValidation(t *testing.T){
	t.Run("should return true when valid name passed",func(t *testing.T) {
		assert.Equal(t,true,isValidName("vasu chauhan"))
	})

	t.Run("should return false when invalid name passed",func(t *testing.T) {
		assert.Equal(t,false,isValidName("   in"))
	})
}