package main

import "strconv"

type ClassTime struct {
	hour     int
	minute   int
	meridiem string
}

type Class struct {
	name string
	time ClassTime
	id   int
}

var classURLFormat = map[string]string{
	"bootcamp":   "60+Min.+Tru+Barefoot+Bootcamp+%28All+Levels%29",
	"pilates":    "60+Min.+Tru+Hot+Pilates+%28All+Levels%29",
	"kettlebell": "60+Min.+Tru+Kettlebell+%28All+Levels%29",
}

// TODO: this is susceptible to class changes since each ID corresponds
// to a particuliar date, time, class, instructor
var classSchedule = map[string][]Class{
	"Monday":    {Class{"pilates", ClassTime{8, 30, "am"}, 113_433}, Class{"bootcamp", ClassTime{10, 15, "am"}, 113_726}},
	"Tuesday":   {Class{"pilates", ClassTime{9, 45, "am"}, 117_478}, Class{"bootcamp", ClassTime{11, 30, "am"}, 118_217}},
	"Wednesday": {Class{"bootcamp", ClassTime{10, 15, "am"}, 108_698}},
	"Friday":    {Class{"kettlebell", ClassTime{12, 00, "pm"}, 108_043}},
}

// getClassIDOffset determines by how much to increment the
// Class.id value
// TODO: this will only work for current year, will need to figure something
// else out later
func getClassIDOffset(currISO int) int {
	origISO := 14 // the ISO value when I originally captured each Class.id
	return currISO - origISO
}

func getClassLink(date string, id int, class string) string {
	// TODO: less gross way of building string
	parts := []string{"https://cart.mindbodyonline.com/sites/14486/cart/add_booking?item%5Binfo%5D=", "&item%5Bmbo_id%5D=", "&item%5Bmbo_location_id%5D=1&item%5Bname%5D=", "&item%5Btype%5D=Class"}
	return parts[0] + date + parts[1] + strconv.Itoa(id) + parts[2] + class + parts[3]
}
