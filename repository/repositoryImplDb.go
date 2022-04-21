package repository

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

const countTask = 3 //максимальное кол-во записей в базе

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

	count := 0
	stmtCount, err := r.rep.DB.Query("select count(id) as count from todolist")
	if err != nil {
		panic(err)
	}
	for stmtCount.Next() {

		err := stmtCount.Scan(&count)
		if err != nil {
			panic(err)
		}
	}
	if count < countTask {

		_, err := r.rep.DB.Exec("INSERT INTO todolist (id, text, isdone) VALUES ($1, $2, $3)", task.Id, task.Text, task.IsDone)
		if err != nil {
			panic(err)
		}
		fmt.Println("создана запись - ", task)
		return nil
	}

	errStr := fmt.Sprintf("не больше %d записей", countTask)
	return errors.New(errStr)

}

//Read
func (r *repositoryImplDB) Read(readFilter *RepFilter) []RepTask {

	sliceTask := make([]RepTask, 0)

	if readFilter.Ids == nil || len(readFilter.Ids) == 0 {
		sliceTask = r.queryRead("select * from todolist", sliceTask)
		return sliceTask
	}

	ids := ""
	for i, id := range readFilter.Ids {

		if i == 0 {
			ids = strconv.Itoa(id)
		} else {
			ids = ids + ", " + strconv.Itoa(id)
		}
	}

	var task RepTask

	//stmtGet, err := r.rep.DB.Query("select * from todolist where id in ($1)", ids)
	strRead := fmt.Sprintf("select * from todolist where id in (%s)", ids)
	stmtGet, err := r.rep.DB.Query(strRead)
	if err != nil {
		panic(err)
	}
	for stmtGet.Next() {

		err := stmtGet.Scan(&task.Id, &task.Text, &task.IsDone)
		if err != nil {
			panic(err)
		}
		sliceTask = append(sliceTask, task)
		fmt.Println("считана запись -", task)
	}
	return sliceTask
}

//Delete
func (r *repositoryImplDB) Delete(delTask *RepTask) error {

	strDelete := fmt.Sprintf("delete from todolist where id=%v", delTask.Id)
	err := r.queryUpdate(strDelete, delTask)
	if err != nil {
		return err
	}
	return nil
}

//Update
func (r *repositoryImplDB) Update(upTask *RepTask) error {

	strUpdate := fmt.Sprintf("update todolist set text='%v', isdone=%v where id=%v", upTask.Text, upTask.IsDone, upTask.Id)
	err := r.queryUpdate(strUpdate, upTask)
	if err != nil {
		return err
	}
	return nil
}
