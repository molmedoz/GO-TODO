package todotask

import (
	"fmt"
	"todo/model"
	"todo/model/errors/notfound"
)

// TodoTask {
// ID int64 id of the task
// Desc string description of the task
// Done bool   flag that indicates if the task is done
// User string owner of the task
// }
type TodoTask struct {
	ID   int64  `json:"id"`
	Desc string `json:"desc"`
	Done bool   `json:"done"`
	User string `json:"user"`
}

// New Generate a new TodoTask
func New(id int64, desc string, done bool, user string) TodoTask {
	return TodoTask{id, desc, done, user}
}

func (t TodoTask) create() error {

	return nil
}

// List return the tasklist of a user
func List(user string) ([]TodoTask, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, description, done, user FROM TASKS WHERE user=?;", user)
	if err != nil {
		return nil, err
	}
	list := []TodoTask{}

	defer rows.Close()
	for rows.Next() {
		var task TodoTask

		rows.Scan(&task.ID, &task.Desc, &task.Done, &task.User)
		list = append(list, task)
	}
	// return list
	return list, nil
}

// Load load a task from the database
func (t *TodoTask) Load() error {
	task := TodoTask{}

	db, err := database.GetConnection()

	if err != nil {
		return err
	}
	rows, err := db.Query("SELECT id, description, done, user FROM TASKS WHERE id=? AND user=?;", t.ID, t.User)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&task.ID, &task.Desc, &task.Done, &task.User)
		break
	}
	if task.Desc == "" {
		return notfounderror.New("TodoTask")
	}
	t.Desc = task.Desc
	t.Done = task.Done
	return nil
}

// GetTask get a task from the database
func GetTask(id int64, user string) (TodoTask, error) {
	task := TodoTask{}
	task.ID = id
	task.User = user
	return task, task.Load()
}

// AddTask Add a task to the database
func (t *TodoTask) AddTask() error {

	db, _ := database.GetConnection()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stm, err := tx.Prepare("INSERT INTO TASKS (description, user) values (?, ?);")
	if err != nil {
		return err
	}

	result, _ := stm.Exec(t.Desc, t.User)

	id, err := result.LastInsertId()
	//task := TodoTask{id, desc, false, user}

	if err != nil {
		e := tx.Rollback()
		if e != nil {
			err = e
		}
		return err
	}
	t.ID = id
	return tx.Commit()
}

// UpdateTask Update a given task, setting it to done or undone
func (t *TodoTask) UpdateTask(done bool) error {

	db, err := database.GetConnection()

	if err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stm, err := tx.Prepare("UPDATE TASKS SET done=? WHERE id = ? AND user=?")

	if err != nil {
		return err
	}

	result, err := stm.Exec(done, t.ID, t.User)
	fmt.Println(result)
	//n, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		// if e != nil {
		// 	err = e
		// }
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	t.Done = done
	return nil
}

// DeleteTask Delete a given task
func (t *TodoTask) DeleteTask() error {

	db, _ := database.GetConnection()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stm, err := tx.Prepare("DELETE FROM TASKS WHERE id = ? AND user=?")
	if err != nil {
		return err
	}

	result, err := stm.Exec(t.ID, t.User)
	result.RowsAffected()
	if err != nil {
		tx.Rollback()

		return err
	}
	return tx.Commit()
}
