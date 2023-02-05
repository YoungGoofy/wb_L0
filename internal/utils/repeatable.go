package utils

import "time"

func DoWithTries(fn func() error, attemtPs int, delay time.Duration) (err error) {
	for attemtPs > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)

			continue
		}
		return nil
	}
	return
}
