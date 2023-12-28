// Package daterange implements date ranges.
package daterange

// ----------------------------------------------------------------------------
// Imports
// ----------------------------------------------------------------------------

import (
	"errors"
	d "waydate/pkg/date"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type DateRange struct {
	first d.Date
	last  d.Date
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

var errorRange = DateRange{}

// ----------------------------------------------------------------------------
// Functions
// ----------------------------------------------------------------------------

// New create a new date range from two dates.
func New(first d.Date, last d.Date) (DateRange, error) {
	var err error = nil

	// Precondition: first <= last
	if first.After(last) {
		var message = "daterange.New: first date " + first.String() +
			" must not be after last date " +
			last.String()
		err = errors.New(message)
		return errorRange, err
	}

	dateRange := DateRange{first, last}
	// Postcondition:
	//   err <> nil or
	//   dateRange.first = first and dateRange.last = last and not dateRange.first.After(dateRange.last)
	return dateRange, err
}

// Overlaps returns true if the two date ranges overlap.
func Overlaps(dateRange1 DateRange, dateRange2 DateRange) bool {
	var result bool
	switch {
	case dateRange1.last.Before(dateRange2.first):
		result = false
	case dateRange1.first.After(dateRange2.last):
		result = false
	default:
		result = true
	}
	return result
}

// ----------------------------------------------------------------------------
// Methods
// ----------------------------------------------------------------------------

// Size returns the number of days in the date range
func (dateRange DateRange) Size() int {
	var size = d.Difference(dateRange.last, dateRange.first)
	return size
}

// InRange returns true if a date in within a date range.
func (dateRange DateRange) InRange(date d.Date) bool {
	var result bool

	switch {
	case date.Before(dateRange.first):
		result = false
	case date.After(dateRange.last):
		result = false
	default:
		result = true
	}
	return result
}

func (dateRange DateRange) String() string {
	var dateRangeString = "(" +
		dateRange.first.String() + "," +
		dateRange.last.String() + ")"
	return dateRangeString
}
