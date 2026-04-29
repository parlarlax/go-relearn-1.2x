package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== 1. Time formatting (Go reference time) ===")
	now := time.Now()
	fmt.Println("RFC3339:", now.Format(time.RFC3339))
	fmt.Println("Kitchen:", now.Format(time.Kitchen))
	fmt.Println("Custom:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("Date only:", now.Format("02-Jan-2006"))

	fmt.Println("\n=== 2. Parsing ===")
	t, err := time.Parse("2006-01-02", "2026-04-28")
	if err != nil {
		fmt.Println("parse error:", err)
	} else {
		fmt.Println("parsed:", t.Format(time.RFC3339))
	}

	t2, _ := time.Parse(time.RFC3339, "2026-04-28T15:30:00+07:00")
	fmt.Println("parsed with tz:", t2)

	fmt.Println("\n=== 3. Duration arithmetic ===")
	d1 := 2*time.Hour + 30*time.Minute
	d2 := 45 * time.Minute
	fmt.Println("total:", d1+d2)
	fmt.Println("hours:", d1.Hours())
	fmt.Println("minutes:", d1.Minutes())

	fmt.Println("\n=== 4. Time arithmetic ===")
	future := now.Add(24 * time.Hour)
	fmt.Println("tomorrow:", future.Format("2006-01-02"))
	diff := future.Sub(now)
	fmt.Println("diff:", diff)

	fmt.Println("\n=== 5. Comparing times ===")
	a := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	b := time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("before:", a.Before(b))
	fmt.Println("after:", a.After(b))

	fmt.Println("\n=== 6. Timer ===")
	timer := time.NewTimer(50 * time.Millisecond)
	<-timer.C
	fmt.Println("timer fired")

	fmt.Println("\n=== 7. Timer — stop/reset ===")
	timer2 := time.NewTimer(5 * time.Second)
	stopped := timer2.Stop()
	fmt.Println("stopped:", stopped)
	timer2.Reset(50 * time.Millisecond)
	<-timer2.C
	fmt.Println("reset timer fired")

	fmt.Println("\n=== 8. Ticker ===")
	ticker := time.NewTicker(30 * time.Millisecond)
	done := make(chan bool)
	go func() {
		count := 0
		for {
			select {
			case <-ticker.C:
				count++
				if count >= 3 {
					done <- true
					return
				}
			}
		}
	}()
	<-done
	ticker.Stop()
	fmt.Println("ticker: 3 ticks received")

	fmt.Println("\n=== 9. AfterFunc ===")
	time.AfterFunc(50*time.Millisecond, func() {
		fmt.Println("  afterfunc: callback!")
	})
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n=== 10. Since / Until ===")
	start := time.Now()
	time.Sleep(10 * time.Millisecond)
	fmt.Println("elapsed:", time.Since(start).Round(time.Millisecond))
	deadline := time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("until 2027:", time.Until(deadline).Round(time.Hour*24))
}
