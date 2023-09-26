package time

import "time"

var LocalDateTime time.Time = time.Now()

func GetYear() int {
	return LocalDateTime.Year()
}

func GetMonth() int {
	return int(LocalDateTime.Month())
}

func GetDay() int {
	return LocalDateTime.Day()
}

func GetHour() int {
	return LocalDateTime.Hour()
}

func GetMinute() int {
	return LocalDateTime.Minute()
}

func GetSecond() int {
	return LocalDateTime.Second()
}
