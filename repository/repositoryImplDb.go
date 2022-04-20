package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var errNotFinedIdDB = errors.New("id not fined")

type repositoryImplDB struct {
	rep *sqlx.DB
}

func NewDB(dbx *sqlx.DB) *repositoryImplDB {

	return &repositoryImplDB{
		rep: dbx,
	}
}

//запрос на считку записей из DB
func (r *repositoryImplDB) queryRead(sqlQuery string, sliceTask []RepTask) []RepTask {

	var task RepTask
	stmtGet, err := r.rep.DB.Query(sqlQuery)
	if err != nil {
		panic(err)
	}
	i := 1
	for stmtGet.Next() {

		err := stmtGet.Scan(&task.Id, &task.Text, &task.IsDone)
		if err != nil {
			panic(err)
		}
		sliceTask = append(sliceTask, task)
		fmt.Println("считана запись -", task)
		i++
	}
	return sliceTask
}

//запрос на изменение/удаление записи в DB
func (r *repositoryImplDB) queryUpdate(sqlQuery string, upTask *RepTask) error {

	stmtUp, err := r.rep.DB.Exec(sqlQuery)
	if err == nil {

		count, err := stmtUp.RowsAffected()
		if count == 0 || err != nil {
			return errNotFinedIdDB
		}

	}
	fmt.Println("изменена/удалена запись c id -", upTask.Id)
	return nil
}

//подсчет кол-ва элементов
func (r *repositoryImplDB) LenRep() int {

	//count := len(r.rep)
	return 0
}

//Create
func (r *repositoryImplDB) Create(task *RepTask) error {

	strCreate := fmt.Sprintf("INSERT INTO todolist (id, text, isDone) VALUES (%d, '%s', %v)", task.Id, task.Text, task.IsDone)
	stmtIns, err := r.rep.DB.Query(strCreate)
	if err != nil {
		panic(err)
	}
	fmt.Println("создана запись - ", task)
	defer stmtIns.Close()
	return nil
}

//подсчет кол-ва элементов
/* func (r *repositoryImplDB) LenRep() int {

	count := len(r.rep)
	return count
} */

//Read

func (r *repositoryImplDB) Read(readFilter *RepFilter) []RepTask {

	sliceTask := make([]RepTask, 0)

	if readFilter.Ids == nil || len(readFilter.Ids) == 0 {
		strRead := fmt.Sprintf("select * from todolist")
		sliceTask = r.queryRead(strRead, sliceTask)
		return sliceTask
	}

	for _, id := range readFilter.Ids {
		strRead := fmt.Sprintf("select * from todolist where id=%d", id)
		sliceTask = r.queryRead(strRead, sliceTask)

	}
	return sliceTask
}

//Delete
func (r *repositoryImplDB) Delete(delTask *RepTask) error {

	strDelete := fmt.Sprintf("delete from todolist where id=%d", delTask.Id)
	err := r.queryUpdate(strDelete, delTask)
	if err != nil {
		return err
	}
	return nil
}

//Update
func (r *repositoryImplDB) Update(upTask *RepTask) error {

	strUpdate := fmt.Sprintf("update todolist set text='%s' , isdone=%v where id=%d", upTask.Text, upTask.IsDone, upTask.Id)
	err := r.queryUpdate(strUpdate, upTask)
	if err != nil {
		return err
	}
	return nil
}
