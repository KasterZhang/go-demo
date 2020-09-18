package observe

import (
	"fmt"
	"time"
)

type company struct {
	employees []employee
}

type employee interface {
	notify(string)
}

type programMonkey struct {
	hair     int8
	workTime time.Duration
}

type personnel struct {
	purse    int64
	workTime time.Duration
}

type designer struct {
	fashion  bool
	workTime time.Duration
}

func (c *company) payroll() {
	message := "payroll send"
	for _, employee := range c.employees {
		employee.notify(message)
	}
}

func (c *company) hire(employee employee) {
	c.employees = append(c.employees, employee)
}

func (p *programMonkey) notify(message string) {
	fmt.Println("After work ", p.workTime, " a week ", "I got ", message)
}

func (p *personnel) notify(message string) {
	fmt.Println("After work ", p.workTime, " a week ", "A message:", message)
}

func (p *designer) notify(message string) {
	fmt.Println("After work ", p.workTime, " a week ", "Cool, ", message)
}
