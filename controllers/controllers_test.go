package controllers

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzCustomClassicOk(t *testing.T) {
	result, err := Fizzbuzzcustom(2, 5, 13, "cat", "dog")
	assert.Nil(t, err)
	if !reflect.DeepEqual(result, []string{"1", "cat", "3", "cat", "dog", "cat", "7", "cat", "9", "catdog", "11", "cat", "13"}) {
		t.Error("Expected [1, cat, 3, cat, dog, cat, 7, cat, 9, catdog, 11, cat, 13], got \n", result)
	}
}

func TestFizzBuzzCustomMutipleOfOneAnotherOk(t *testing.T) {
	result, err := Fizzbuzzcustom(2, 4, 15, "cat", "dog")
	assert.Nil(t, err)
	if !reflect.DeepEqual(result, []string{"1", "cat", "3", "catdog", "5", "cat", "7", "catdog", "9", "cat", "11", "catdog", "13", "cat", "15"}) {
		t.Error("Expected [1, cat, 3, catdog, 5, 6, 7, catdog, 9, cat, 11, catdog, 13, cat, 15], got \n", result)
	}
}

func TestFizzBuzzCustomEmptyStringOk(t *testing.T) {
	result, err := Fizzbuzzcustom(2, 5, 18, "", "dog")
	assert.Nil(t, err)
	if !reflect.DeepEqual(result, []string{"1", "", "3", "", "dog", "", "7", "", "9", "dog", "11", "", "13", "", "dog", "", "17", ""}) {
		t.Error("Expected [1, 3, dog, 7, 9, dog, 11, 13, dog, 17], got \n", result)
	}
}

func TestFizzBuzzCustomFirstZero(t *testing.T) {
	result, err := Fizzbuzzcustom(0, 3, 22, "cat", "dog")
	assert.NotNil(t, err)
	assert.Nil(t, result)
}
func TestFizzBuzzCustomSecondZero(t *testing.T) {
	result, err := Fizzbuzzcustom(2, 0, 22, "cat", "dog")
	assert.NotNil(t, err)
	assert.Nil(t, result)
}
func TestFizzBuzzCustomAllZero(t *testing.T) {
	result, err := Fizzbuzzcustom(0, 0, 22, "cat", "dog")
	assert.NotNil(t, err)
	assert.Nil(t, result)
}
