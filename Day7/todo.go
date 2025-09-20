package todo

type Todo struct {
	Task   string
	status bool
}

func (t *Todo) Add(task string) {
	t.Task = task
	t.status = false
}

func (t *Todo) MarkDone() {
	t.status = true
}

func StatusChecker(stat bool) string {
	if stat {
		return "Done"
	}
	return "Pending"
}
