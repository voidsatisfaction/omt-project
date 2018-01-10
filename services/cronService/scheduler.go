package cronService

// This is my custom scheduler package
// Scheduler only takes no parameter, no return function

import "time"

type schedule struct {
	work     func()
	ticker   *time.Ticker
	duration time.Duration
}

func newSchedule(work func(), duration time.Duration) *schedule {
	return &schedule{work, time.NewTicker(duration), duration}
}

// Scheduler is
type Scheduler []*schedule

func New() *Scheduler {
	return &Scheduler{}
}

// Add is
func (s *Scheduler) Add(work func(), duration time.Duration) {
	schedule := newSchedule(work, duration)
	*s = append(*s, schedule)
}

func (s *Scheduler) Len() int {
	return len(*s)
}

// Run is to execute all works
func (s *Scheduler) Run() {
	for i := 0; i < s.Len(); i++ {
		schedule := (*s)[i]
		ticker := schedule.ticker

		go func() {
			for {
				select {
				case <-ticker.C:
					schedule.work()
				}
			}
		}()
	}
}
