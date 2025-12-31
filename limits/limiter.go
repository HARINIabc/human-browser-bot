package limits

import "time"

type Limiter struct {
	maxActions int
	actions    int
	lastReset  time.Time
}
func New(max int) *Limiter {
	return &Limiter{
		maxActions: max,
		actions:    0,
		lastReset:  time.Now(),
	}
}
func (l *Limiter) Allow() bool {
	
	if time.Since(l.lastReset) > 24*time.Hour {
		l.actions = 0
		l.lastReset = time.Now()
	}

	if l.actions >= l.maxActions {
		return false
	}

	l.actions++
	return true
}
