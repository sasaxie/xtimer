/**
 * Timer for scheduled tasks.
 */

package xtimer

import (
	"time"
)

/*
 * Execute your logic until program was done or stop this function.
 */
func StartTicker(interval time.Duration, f func(now time.Time)) chan bool {
	done := make(chan bool, 1)

	go func() {
		t := time.NewTicker(interval)
		defer t.Stop()
		for {
			select {
			case now := <-t.C:
				f(now)
			case <-done:
				return
			}
		}
	}()

	return done
}

/*
 * Regularly executed once.
 */
func StartTimer(interval time.Duration, f func(now time.Time)) chan bool {
	done := make(chan bool, 1)

	go func() {
		t := time.NewTimer(interval)
		defer t.Stop()
		for {
			select {
			case now := <-t.C:
				f(now)
			case <-done:
				return
			}
		}
	}()

	return done
}
