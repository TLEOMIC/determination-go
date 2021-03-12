package tool

import "time"

func SleepTillMidnight(){
	year, month, day := time.Now().Add(24*time.Hour).Date()
	time.Sleep(time.Duration(time.Date(year, month, day, 0, 0, 0, 0, time.Local).Unix()-time.Now().Unix())*time.Second)
}