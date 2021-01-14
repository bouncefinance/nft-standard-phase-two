package gpool

type Task struct {
	f func()error
}

func NewTask(task func() error) *Task {
	t := Task{
		f: task,
	}

	return &t
}
func (t *Task) Execute() {
	t.f()
}
