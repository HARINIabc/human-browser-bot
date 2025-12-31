package scheduler

import "time"

type Scheduler struct {
	startHour int
	endHour   int
}

func New(startHour, endHour int) *Scheduler {
	return &Scheduler{
		startHour: startHour,
		endHour:   endHour,
	}
}

func (s *Scheduler) IsAllowed() bool {
	now := time.Now().Hour()
	return now >= s.startHour && now < s.endHour
}
