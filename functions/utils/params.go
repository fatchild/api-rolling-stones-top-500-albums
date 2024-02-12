package utils

import (
	"strconv"

	"github.com/fatchild/api-rolling-stones-top-500-albums/functions/consts"
)

func SanitiseYear(yearString string, yearGiven bool) string {
	yearSanitised := consts.VALID_YEAR_2023

	if yearGiven && (yearString == "2023" || yearString == "2020" || yearString == "2012" || yearString == "2003") {
		yearSanitised = yearString
	}

	return yearSanitised
}

func SanitisePosition(posString string, posGiven bool) string {
	posSanitised := 1

	posString, err := strconv.Atoi(posInt)
	if err == nil && posGiven && posInt > 1 && posInt < 500 {
		posSanitised = posString
	}

	return posSanitised
}
