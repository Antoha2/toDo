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

func (r *repositoryImpl) LenRep() int {

	count := len(r.rep)
	//fmt.Println(count)
	return count
}

func (r *repositoryImpl) Read(readTask *RepFilter) *RepTask {

	/* if len(r.rep) < readTask.Id {
		return nil
	} */

	var tsd *RepTask

	for _, task := range r.rep {

		if task.Id == readTask.Id {
			tsd = &r.rep[readTask.Id-1]
		}
	}
	fmt.Println(tsd)
	return tsd

}

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
