package sTime

import (
	"fmt"
	"strconv"
)

type Time uint

func (t *Time) Set(time []byte) {
	hours, _ := strconv.Atoi(string(time[:2]))
	minutes, _ := strconv.Atoi(string(time[3:5]))
	seconds, _ := strconv.Atoi(string(time[6:]))

	*t = Time(seconds + 60*minutes + 3600*hours)
}

func (t Time) Hours() int {
	return int(t / 3600)
}

func (t Time) Minutes() int {
	return int((t % 3600) / 60)
}

func (t Time) Seconds() int {
	return int(t % 60)
}

func Diff(A, B Time) Time {
	if A > B {
		B += 3600 * 24
	}
	return B - A
}

func (t Time) ToString() string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hours(), t.Minutes(), t.Seconds())
}
