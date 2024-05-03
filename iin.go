package main

import (
	"fmt"
	"strconv"
	"time"
)

func VerifyIIN(iin string) bool {
	if len(iin) != 12 { // 12 digits in IIN number it's 1 byte per digit
		return false
	}

	if !isValidChecksum(iin) {
		return false
	}

	return true
}

func GetDateOfBirth(iin string) (time.Time, error) {
	digit := string(iin[6])

	year := 0
	yearDigits, err := strconv.Atoi(iin[0:2])
	if err != nil {
		return time.Time{}, err
	}

	switch digit { // 0 - for foreigers
	case "1", "2":
		year = 1800 + yearDigits
	case "3", "4":
		year = 1900 + yearDigits
	case "5", "6":
		year = 2000 + yearDigits
	}

	month, err := strconv.Atoi(iin[2:4])
	if err != nil {
		return time.Time{}, err
	}

	day, err := strconv.Atoi(iin[4:6])
	if err != nil {
		return time.Time{}, err
	}

	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	return dob, nil
}

func GetGender(iin string) (string, error) {
	digit, err := strconv.Atoi(string(iin[6]))
	if err != nil {
		return "", err
	}

	if digit < 1 && digit > 6 {
		return "", fmt.Errorf("invalid IIN 7's digit")
	}

	if digit%2 == 0 {
		return "female", nil
	} else if digit%2 == 1 {
		return "male", nil
	}

	return "", fmt.Errorf("invalid IIN 7's digit")
}

// isValidChecksum checks if the checksum of the IIN is valid
// The checksum is calculated by multiplying each digit of the IIN by a weight
// and taking the remainder of the sum of the products divided by 11
// If the remainder is 10, the process is repeated with different weights
// If second time the remainder is 10, it is replaced with 0
// The last digit of the IIN is the checksum
func isValidChecksum(iin string) bool {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	sum := 0

	for i, v := range iin[:len(iin)-1] {
		digit, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}

		sum += digit * weights[i]
	}

	remainder := sum % 11
	checksum, err := strconv.Atoi(string(iin[len(iin)-1]))
	if err != nil {
		return false
	}

	weights = []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2}

	if remainder == 10 {
		sum = 0
		for i, v := range iin[:len(iin)-1] {
			digit, err := strconv.Atoi(string(v))
			if err != nil {
				return false
			}

			sum += digit * weights[i]
		}

		remainder = sum % 11
		if remainder == 10 {
			remainder = 0
		}
	}

	return remainder == checksum
}
