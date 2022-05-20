package models

import (
	"log"

	_ "modernc.org/sqlite"
)

type Task struct {
	Id        int64  `json:"id,omitempty"`
	Text      string `json:"text,omitempty"`
	ListId    int    `json:"listId"`
	Completed bool   `json:"completed"`
}

func GetTasksByListId(id int) ([]Task, error) {

	var tasks []Task

	rows, err := db.Query("select task.* from task where listId = ?", id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.Id, &t.Text, &t.ListId, &t.Completed); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil

}

func CreateTask(id int, task *Task) *Task {
	res, err := db.NamedExec(`INSERT INTO task(text, listId, completed) VALUES (:text, :listId, :completed)`,
		map[string]interface{}{
			"text":      task.Text,
			"listId":    id,
			"completed": task.Completed,
		})

	if err != nil {
		log.Fatal(err)
	}
	v, _ := res.LastInsertId()
	task.Id = v
	task.ListId = id

	return task
}

// func CreateTask(id int, task *Task) (*Task, error) {
// 	res, err := db.NamedExec(`INSERT INTO task(text, listId, completed) VALUES (:text, :listId, :completed)`,
// 		map[string]interface{}{
// 			"text":      task.Text,
// 			"listId":    id,
// 			"completed": task.Completed,
// 		})

// 	if err != nil {
// 		return nil, err
// 	}
// 	v, _ := res.LastInsertId()
// 	task.Id = v
// 	task.ListId = id

// 	return task, nil
// }

func UpdateTask(id int, v bool) (Task, error) {
	var t Task
	_, err := db.Exec("update task set completed = ? where id = ?", v, id)

	if err != nil {
		return t, err
	}
	row := db.QueryRow("select * from task where id = ?", id)

	if err := row.Scan(&t.Id, &t.Text, &t.ListId, &t.Completed); err != nil {

		return t, err
	}
	return t, nil
}

func DeleteTask(id int) error {

	_, err := db.Exec("delete from task where id = ?", id)

	return err

}
