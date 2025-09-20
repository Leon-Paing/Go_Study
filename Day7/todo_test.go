package todo

import "testing"

func TestAdd(t *testing.T) {
	var t1 Todo
	t1.Add("Learn Go")

	if t1.Task != "Learn Go" {
		t.Errorf("Expected `Learn Go` got %s", t1.Task)
	}
}

func TestMarkDone(t *testing.T) {
	var t1 Todo
	t1.Add("Walk the dog")
	t1.MarkDone()

	if !t1.status {
		t.Errorf("Expected true got false")
	}
}

func TestStatusChecker(t *testing.T) {
	var t1 Todo
	t1.Add("Read a book")
	t1.MarkDone()
	status := StatusChecker(t1.status)

	if status != "Done" {
		t.Errorf("Expected `Done` got %s", status)
	}
}
