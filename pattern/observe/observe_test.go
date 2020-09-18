package observe

import (
	"testing"
	"time"
)

func TestObserve(t *testing.T) {
	monkey1 := &programMonkey{
		hair:     0,
		workTime: 12 * 6 * time.Hour,
	}
	personnel1 := &personnel{
		purse:    10000,
		workTime: 6 * 8 * time.Hour,
	}
	designer1 := &designer{
		fashion:  true,
		workTime: 1 * 6 * time.Hour,
	}
	hw := &company{}
	hw.hire(monkey1)
	hw.hire(personnel1)
	hw.hire(designer1)

	hw.payroll()
}
