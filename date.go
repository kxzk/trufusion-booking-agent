package main

import (
	"strconv"
	"time"
)

func getNextWeekDate() time.Time {
	// since remote server is not PST
	tz, _ := time.LoadLocation("US/Pacific")
	oneWeek := time.Hour * 24 * 7
	return time.Now().In(tz).Add(oneWeek)
}

// convertMonth abbreviates a month's name.
// However, I am unsure as to whether this works
// for every month or just longer months.
// For instance, I don't know if June or July
// will be shortened.
// TODO: observe June, July for insight into
// abbreviation strategy
func truncMonth(m string) string {
	return m[:3]
}

// convertWeekday grab first three letters of weekday
func truncWeekday(w string) string {
	return w[:3]
}

func getFormattedDate(t time.Time, c ClassTime) string {
	// TODO: refactor to something more elegant, pretty gross
	// example -> Mon.+Apr++4%2C+2022++8%3A30+am
	// (month) `.+` (dow) `++` (day) `%2c+` (year) `++` | (hour) `%3A` (minute) `+` (meridiem)
	parts := []string{".+", "++", "%2C+", "++", "%3A", "+"}

	dow := truncWeekday(t.Weekday().String())
	mon := truncMonth(t.Month().String())
	day := strconv.Itoa(t.Day())
	year := strconv.Itoa(t.Year())
	hour := strconv.Itoa(c.hour)
	min := strconv.Itoa(c.minute)

	return dow + parts[0] + mon + parts[1] + day + parts[2] + year + parts[3] + hour + parts[4] + min + parts[5] + c.meridiem
}
