// Package file implements methods for get logs from files.
//
// Log sorting.
package file

import (
	"sort"
)

// Aggregate extend Sort interface.
type Aggregate []Log

func (a Aggregate) Len() int {
	return len(a)
}

func (a Aggregate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Aggregate) Less(i, j int) bool {
	return a[i].Timestamp.Unix() < a[j].Timestamp.Unix()

	//firstDate := a[i].Timestamp
	//secondDate := a[j].Timestamp
	//
	//firstDate.Unix()
	//
	//fmt.Println(firstDate)
	//fmt.Println(secondDate)
	//
	//return (firstDate.Year() < secondDate.Year()) &&
	//	(firstDate.Month() < secondDate.Month()) &&
	//	(firstDate.Day() < secondDate.Day()) &&
	//	(firstDate.Hour() < secondDate.Hour()) &&
	//	(firstDate.Minute() < secondDate.Minute()) &&
	//	(firstDate.Second() < secondDate.Second())
}

// SortByTimestamp sort log by timestamp.
func SortByTimestamp(logs []Log) []Log {
	sort.Sort(Aggregate(logs))
	return logs
}
