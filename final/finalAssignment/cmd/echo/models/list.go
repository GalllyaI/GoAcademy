package models

import (
	"fmt"

	_ "modernc.org/sqlite"
)

type List struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Tasks []Task `json:"tasks,omitempty"`
}

func GetAllLists() ([]List, error) {

	var lists []List
	rows, err := db.Queryx("select * from list")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var l List
		err = rows.StructScan(&l)
		//err ????
		lists = append(lists, List{Id: l.Id, Name: l.Name})
	}

	return lists, nil
}

func getListById(id int) (List, error) {
	row := db.QueryRowx("select * from list where id = ?", id)

	var l List
	err := row.StructScan(&l)
	if err != nil {
		return l, err
	}

	return l, nil
}

func CreateList(list *List) (*List, error) {

	res, err := db.NamedExec(`INSERT INTO list(name) VALUES (:name)`,
		map[string]interface{}{
			"name": list.Name,
		})

	if err != nil {
		return nil, err
	}
	v, _ := res.LastInsertId()
	list.Id = v

	return list, nil
}

func DeleteList(id int) error {

	_, err := getListById(id)

	if err != nil {
		return fmt.Errorf("no list with that id")
	}

	_, err = db.Exec("delete from list where id = ?", id)
	// r, _ := row.RowsAffected()

	return err

}
