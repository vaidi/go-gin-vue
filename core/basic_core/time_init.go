package basic_core

import (
	"fmt"
	"time"
)

func TimeInit() {
	testTime1()
	testTime2()
	testTime3()
}

func testTime1() {
	now := time.Now()
	fmt.Printf("current time is %v\n", now)
	yearDay := now.YearDay()
	year := now.Year()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("year day %v, hour %v, minute %v year%v day %v second %v", yearDay, hour, minute, year, day, second)
}

func testTime2() {
	now := time.Now()
	unix := now.Unix()
	nano := now.UnixNano() //纳秒时间戳
	fmt.Printf("current time is %v unix:%v nano:%v\n", now, unix, nano)

}
func testTime3() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
	//定时器
	tick := time.Tick(time.Hour)
	//匿名函数
	//go func() {
	//	for i := range tick {
	//		fmt.Println("执行定时器", i)
	//	}
	//}()
	go func(a <-chan time.Time) {
		for i := range a {
			fmt.Println("1个小时执行一次 执行定时器", i)
		}
	}(tick)

}

func testTime4() {
	now := time.Now()
	fmt.Println(now)
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	inLocation, err := time.ParseInLocation("2006-01-02 15:04:05.000 Mon Jan", now.Format(time.RFC3339), location)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(inLocation)
	fmt.Println(inLocation.Sub(now))

}
