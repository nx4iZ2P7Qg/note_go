package main

import (
	"fmt"
	"log"
	"time"
)

// 基本使用
func demo0101() {

	// We'll start by getting the current time.
	now := time.Now()
	fmt.Println(now)

	// You can build a `time` struct by providing the
	// year, month, day, etc. Times are always associated
	// with a `Location`, i.e. time zone.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	fmt.Println(then)

	// You can extract the various components of the time
	// value as expected.
	fmt.Println(then.Year())
	fmt.Println(then.Month())
	fmt.Println(then.Day())
	fmt.Println(then.Hour())
	fmt.Println(then.Minute())
	fmt.Println(then.Second())
	fmt.Println(then.Nanosecond())
	fmt.Println(then.Location())

	// The Monday-Sunday `Weekday` is also available.
	fmt.Println(then.Weekday())

	// These methods compare two times, testing if the
	// first occurs before, after, or at the same time
	// as the second, respectively.
	fmt.Println(then.Before(now))
	fmt.Println(then.After(now))
	fmt.Println(then.Equal(now))

	// The `Sub` methods returns a `Duration` representing
	// the interval between two times.
	diff := now.Sub(then)
	fmt.Println(diff)

	// We can compute the length of the duration in
	// various units.
	fmt.Println(diff.Hours())
	fmt.Println(diff.Minutes())
	fmt.Println(diff.Seconds())
	fmt.Println(diff.Nanoseconds())

	// You can use `Add` to advance a time by a given
	// duration, or with a `-` to move backwards by a
	// duration.
	fmt.Println(then.Add(diff))
	fmt.Println(then.Add(-diff))

	// Use constants like: Nanosecond, Microsecond, Millisecond,
	// Second, Minute, Hour to get new time, like this example to
	// print date after/before 30 days. Note: Substract one month
	// is not equal to 30 days.
	fmt.Println(now.Add(30 * 24 * time.Hour).Format("2006-01-02"))
	fmt.Println(now.Add(-30 * 24 * time.Hour).Format("2006-01-02"))
	fmt.Println(now.AddDate(0, -1, 0).Format("2006-01-02"))
	fmt.Println(now.AddDate(0, 0, -30).Format("2006-01-02"))
}

// 字符串转日期
func demo0102() {
	start, err := time.ParseInLocation("2006-01-02", "2020-02-27", time.Local)
	if err != nil {
		log.Fatal("error", err)
	}
	duration := time.Now().Sub(start)

	log.Println(duration.Hours() / 24)
}
