package todo

type Task struct {
	Title       string
	Description string
	Completed   bool
}

func NewTask(title string, description string) Task {
	return Task{
		Title:       title,
		Description: description,
		Completed:   false,
	}

}

func (t *Task) IsCompleted() {
	t.Completed = true
}
func (t *Task) UnComplete() {
	t.Completed = false
}	
