package main

import (
	"learning/mocking/countdown"
	"os"
	"time"
)

func main() {
	sleeper := &countdown.ConfigurableSleeper{Duration: 1 * time.Second, SleepFn: time.Sleep}
	countdown.Countdown(os.Stdout, sleeper)
}
