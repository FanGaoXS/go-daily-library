package main

import (
	"fmt"
	"github.com/uniplaces/carbon"
)

func main() {
	now := carbon.Now()

	fmt.Println("now is:", now)

	fmt.Println("one second later is:", now.AddSecond())
	fmt.Println("one minute later is:", now.AddMinute())
	fmt.Println("one hour later is:", now.AddHour())
	fmt.Println("3 minutes and 20 seconds later is:", now.AddMinutes(3).AddSeconds(20))
	fmt.Println("2 hours and 30 minutes later is:", now.AddHours(2).AddMinutes(30))
	fmt.Println("3 days and 2 hours later is:", now.AddDays(3).AddHours(2))

	// Eq/EqualTo：是否相等；
	// Ne/NotEqualTo：是否不等；
	// Gt/GreaterThan：是否在之后；
	// Lt/LessThan：是否在之前；
	// Lte/LessThanOrEqual：是否相同或在之前；
	// Between：是否在两个时间之间。

	// IsMonday/IsTuesday/.../IsSunday：判断周几
	// 是否是工作日，周末，闰年，过去时间还是未来时间：IsWeekday/IsWeekend/IsLeapYear/IsPast/IsFuture
	fmt.Println("IsLeapYear =", now.IsLeapYear())
	fmt.Println("IsWeekday =", now.IsWeekday())
	fmt.Println("IsWeekend =", now.IsWeekend())
	fmt.Println("IsPast =", now.IsPast())
	fmt.Println("IsFuture =", now.IsFuture())

	t1, _ := carbon.CreateFromDate(2012, 1, 31, "UTC")
	t2, _ := carbon.CreateFromDate(2012, 4, 30, "UTC")
	fmt.Println(t1.DiffInDays(t2, true))

	// 自定义工作日和周末
	// t.SetWeekStartsAt(time.Sunday)
	// t.SetWeekEndsAt(time.Saturday)
	// t.SetWeekendDays([]time.Weekday{time.Monday, time.Tuesday, time.Wednesday})
}
