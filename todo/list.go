package todo

import "sync"

type List struct {
	tasks map[string]Task
	mtx   sync.RWMutex
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}

func (l *List) AddTask(task Task) error {

	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.tasks[task.Title]; ok {
		return ErrTaskAlredyExist
	}

	l.tasks[task.Title] = task
	return nil
}

func (l *List) GetTask(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[title]

	if !ok {
		return Task{}, ErrTaskNotFound
	}
	return task, nil
}
func (l *List) ListTasks() map[string]Task {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	tmp := make(map[string]Task, len(l.tasks))

	for k, v := range l.tasks {
		tmp[k] = v
	}
	return tmp
}
func (l *List) ListUnCompletedTasks() map[string]Task {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	uncompletedtasks := make(map[string]Task)

	for title, task := range l.tasks {
		if !task.Completed {
			uncompletedtasks[title] = task
		}
	}
	return uncompletedtasks
}
func (l *List) CompletedTask(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	task.IsCompleted()
																					
	l.tasks[title] = task

	return task, nil
}
func (l *List) UnCompleteTask(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrTaskNotFound
	}
	task.UnComplete()
	l.tasks[title] = task
	return task, nil
}
func (l *List) DeleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()	

	_, ok := l.tasks[title]
	if !ok {
		return ErrTaskNotFound
	}

	delete(l.tasks, title)
	return nil
}
