package xtimer

import (
    "testing"
    "time"
    "log"
)

func TestStartTicker(t *testing.T) {
    ticker := StartTicker(time.Second, func(now time.Time) {
        log.Println("Time:", now.Format("2006-01-02 15:04:05"))
    })

    i := 0
    for {
        i++
        if i == 10 {
            // Stop ticker
            close(ticker)
            break
        }
        time.Sleep(time.Second)
    }

    log.Println("end")
}

func TestStartTimer(t *testing.T) {
    timer := StartTimer(time.Second, func(now time.Time) {
        log.Println("Time:", now.Format("2006-01-02 15:04:05"))
    })

    i := 0
    for {
        i++
        if i == 10 {
            // Stop timer
            close(timer)
            break
        }
        time.Sleep(time.Second)
    }

    log.Println("end")
}