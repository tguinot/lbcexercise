package controllers

import (
	"errors"
	"strconv"
)

func Fizzbuzzcustom(int1 int, int2 int, limit int, str1 string, str2 string) ([]string, error) {
	var result []string
	if int1 <= 0 || int2 <= 0 {
		return nil, errors.New("input integers must be strictly positive")
	}
	for i := 1; i <= limit; i++ {
		if i%int1 == 0 && i%int2 == 0 {
			result = append(result, str1+str2)
		} else if i%int1 == 0 {
			result = append(result, str1)
		} else if i%int2 == 0 {
			result = append(result, str2)
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result, nil
}
