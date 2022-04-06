package repository

import "fmt"

type repositoryImpl struct {
	rep []RepTask
}

func New() *repositoryImpl {
	r := make([]RepTask, 0)

	return &repositoryImpl{
		rep: r,
	}
}

func (r *repositoryImpl) Create(task *RepTask) error {

	r.rep = append(r.rep, *task)
	fmt.Println(r)
	return nil
}

//подсчет кол-ва элементов
func (r *repositoryImpl) LenRep() int {

	count := len(r.rep)
	//fmt.Println(count)
	return count
}

//Read
func (r *repositoryImpl) Read(readTask *RepFilter) *RepTask {

	var tsd *RepTask

	for _, task := range r.rep {

		if task.Id == readTask.Id {
			tsd = &r.rep[readTask.Id-1]
		}
	}
	fmt.Println(tsd)
	return tsd
}

//Delete
func (r *repositoryImpl) Delete(delTask *RepFilter) error {

	for i, v := range r.rep {
		if v.Id == delTask.Id {
			copy(r.rep[i:], r.rep[i+1:])
			//r.rep[len(r.rep)-1] = nil // обнуляем "хвост"
			r.rep = r.rep[:len(r.rep)-1]
		}
	}
	fmt.Println(r)
	return nil
}

//Update
func (r *repositoryImpl) Update(upTask *RepTask) error {

	//var tsd *RepTask

	for index, _ := range r.rep {

		if r.rep[index].Id == upTask.Id {

			r.rep[index].Text = upTask.Text
			r.rep[index].IsDone = upTask.IsDone
			fmt.Println(r.rep[index])
		}
	}
	//fmt.Println(tsd)
	return nil
}
