// ----------------------------------------------------------------------------
//
// Year Month
//
// Author: William Shaffer
// Version: 25-Apr-2024
//
// Copyright (c) 2024 William Shaffer
//
// ----------------------------------------------------------------------------

package date

// This file manages year month structures

// ----------------------------------------------------------------------------
// Imports
// ----------------------------------------------------------------------------

import (
	"strconv"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type YearMonth struct {
	Year  int
	Month int
}

type YearMonthString struct {
	Year  int
	Month string
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const (
	startYear  = 2022
	startMonth = 9
	endYear    = 2025
	endMonth   = 12
)

var namesMonth = []string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

// ----------------------------------------------------------------------------
// Factory Functions
// ----------------------------------------------------------------------------

func NewYearMonth(year int, month int) (YearMonth, error) {
	var err error = nil
	//
	// Preconditions
	//
	err = isYear(Year(year))
	if err != nil {
		return YearMonth{}, err
	}
	err = isMonth(Month(month))
	if err != nil {
		return YearMonth{}, err
	}
	//
	// Set the values
	//
	yearMonth := YearMonth{
		Year:  year,
		Month: month,
	}
	return yearMonth, err
}

func NewYearMonthFromDate(date Date) (YearMonth, error) {
	var year = int(date.Year())
	var month = int(date.Month())
	return NewYearMonth(year, month)
}

// Keys returns a slice with the year month structures for 2023 and 2024.
func Keys() ([]YearMonth, error) {
	var keys []YearMonth
	var err error
	var yearMonth YearMonth

	for year := startYear; year <= endYear; year++ {
		for month := 1; month <= 12; month++ {
			switch {
			case year == startYear && month < startMonth:
				// do nothing
			case year == endYear && month > endMonth:
				// do nothing
			default:
				yearMonth, err = NewYearMonth(year, month)
				if err != nil {
					return keys, err
				}
				keys = append(keys, yearMonth)
			}
		}
	}
	return keys, nil
}

// newYearMonthString creates a new YearMonthString with the month translated
// from the month number (1..12) to the month abbreviation.
func newYearMonthString(yearMonth YearMonth) YearMonthString {
	yms := YearMonthString{
		Year:  yearMonth.Year,
		Month: namesMonth[yearMonth.Month-1],
	}
	return yms
}

// ----------------------------------------------------------------------------
// Functions
// ----------------------------------------------------------------------------

// MonthAbbr returns the abbreviation of a

// ----------------------------------------------------------------------------
// Methods - YearMonth
// ----------------------------------------------------------------------------

// String converts a year month to a string in the format MM/YYYY.
func (yearMonth YearMonth) String() string {
	var value = strconv.Itoa(yearMonth.Month) + "/" + strconv.Itoa(yearMonth.Year)
	return value
}

// MonthString converts the current year month to a year month string.
func (yearMonth YearMonth) MonthString() YearMonthString {
	yms := newYearMonthString(yearMonth)
	return yms
}
