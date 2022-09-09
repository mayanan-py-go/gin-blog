package cron

import (
	"time"
)

// SetTime 获取到自定时间的Duration 误差在1s内
// 计算设置时间和当前时间的差值，大于当前时间则返回，否则为第二天的时间
func SetTime(hour, minute, second int) (d time.Duration) {
	now := time.Now()
	setTime := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, second, 0, now.Location())
	d = setTime.Sub(now)
	if d > 0 {
		return
	}
	return d + time.Hour*24
}
func ScheduleTask(f func()) {
	timer := time.NewTimer(SetTime(0, 0, 0))
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			timer.Reset(time.Hour * 24)
			f()
		}
	}
}
