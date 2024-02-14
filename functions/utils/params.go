package utils

import (
	"strconv"

	"github.com/fatchild/api-rolling-stones-top-500-albums/functions/consts"
)

func SanitiseYearParam(yearString string, yearGiven bool) string {
	yearSanitised := consts.VALID_YEAR_2023

	if yearGiven && StringIn(yearString, consts.VALID_YEARS) {
		yearSanitised = yearString
	}

	return yearSanitised
}

func SanitisePositionParam(posString string, posGiven bool) string {
	posSanitized := "1"

	posInt, err := strconv.Atoi(posString)
	if err == nil && posGiven && (posInt >= 1 && posInt <= 500) {
		posSanitized = posString
	}

	return posSanitized
}
