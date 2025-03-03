package initial

import (
	"golang.org/x/time/rate"
)

func (initial *Initial) InitRate() *Initial {
	initial.Limiter = rate.NewLimiter(
		rate.Limit(initial.Config.Float("Rate.Limit", 10)),
		initial.Config.Int("Rate.Burst", 10),
	)
	return initial
}
