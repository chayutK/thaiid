package thaiid

import (
	"strconv"
	"strings"
)

func IsValid(id string) bool {
	idSlice := sliceId(id)
	if len(idSlice) != 13 {
		return false
	}
	val := idCalculate(idSlice)

	modded := valOperating(val)

	lastDigit := strconv.Itoa(modded)

	return lastDigit == idSlice[len(idSlice)-1]
}

func idCalculate(idSlice []string) (result int) {
	result = 0

	for i, v := range idSlice[:12] {
		val, _ := strconv.Atoi(v)
		result += val * (13 - i)
	}

	return result
}

func valOperating(val int) int {
	return (11 - (val % 11)) % 10
}

func sliceId(id string) []string {
	id = strings.Replace(id, "-", "", -1)
	idSlice := strings.Split(id, "")

	return idSlice
}
