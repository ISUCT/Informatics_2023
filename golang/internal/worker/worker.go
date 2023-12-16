package worker

import (
	"fmt"
)

type worker struct {
	name   string
	status string
	salary int
}

func NewWorker(name_worker string, status string, salary int) (worker, error) {
	var w worker = worker{
		name: name_worker,
	}
	w.SetStatus(status)
	var err = w.SetSalary(salary)
	return w, err
}

func (w *worker) QuitJob() {
	fmt.Printf("%s уволился", w.GetName())
}

func (w *worker) SetName(name string) {
	w.name = name
}

func (w *worker) SetStatus(status string) {
	w.status = status
}

func (w *worker) SetSalary(salary int) error {
	if salary >= 15000 {
		w.salary = salary
		return nil
	} else {
		return fmt.Errorf("Минимальная зарплата - 15000")
	}
}

func (w *worker) GetName() string {
	return w.name
}

func (w *worker) GetStatus() string {
	return w.status
}

func (w *worker) GetSalary() int {
	return w.salary
}
