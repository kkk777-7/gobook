package timeutil

import "time"

func LessMonth(t, t2 time.Time) bool {
	onemonth := t2.AddDate(0, -1, 0)
	return t.After(onemonth)
}

func LessYear(t, t2 time.Time) bool {
	oneyear := t2.AddDate(-1, 0, 0)
	return t.After(oneyear)
}
