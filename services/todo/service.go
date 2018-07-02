package todoservice

import (
	"todo/model/errors/input"
	"todo/model/errors/notfound"
	"todo/model/task"
)

// GetList a user list
func GetList(user string) ([]todotask.TodoTask, error) {
	if user == "" {
		fields := []string{"user"}
		return nil, inputerror.New(fields)
	}
	return todotask.List(user)

}

// AddTask Add a new task to the usser list
func AddTask(desc string, user string) (todotask.TodoTask, error) {
	fields := []string{}

	if desc == "" {
		fields = append(fields, "desc")
	}
	if user == "" {
		fields = append(fields, "user")
	}
	if len(fields) > 0 {
		return todotask.TodoTask{}, inputerror.New(fields)
	}

	task := todotask.New(0, desc, false, user)
	return task, task.AddTask()

}

// UpdateTask Update a given task, setting it to done or undone
func UpdateTask(id int64, user string, done bool) (bool, error) {
	task, err := todotask.GetTask(id, user)
	if err != nil {
		return false, err
	}
	if task.ID == 0 {
		err = notfounderror.New("TodoTask")
		return false, err
	}
	err = task.UpdateTask(done)
	return err == nil, err

}

// DeleteTask Delete a given task
func DeleteTask(id int64, user string) (bool, error) {
	task, err := todotask.GetTask(id, user)
	if err != nil {
		return false, err
	}
	if task.ID == 0 {
		err = notfounderror.New("TodoTask")
		// NOT found
		return false, err
	}
	err = task.DeleteTask()
	return err == nil, err

}

// GetTask Get a task from the database
func GetTask(id int64, user string) (todotask.TodoTask, error) {
	fields := []string{}
	if id == 0 {
		fields = append(fields, "id")
	}
	if user == "" {
		fields = append(fields, "user")
	}

	if len(fields) > 0 {
		return todotask.TodoTask{}, inputerror.New(fields)
	}

	return todotask.GetTask(id, user)
}
