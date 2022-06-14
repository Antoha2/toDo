package authrepository

import (
	"errors"
	"fmt"
	"log"

	etodo "github.com/antoha2/todo"
	"gorm.io/gorm"
)

//регистрация нового пользователя
func (r *AuthPostgres) CreateUser(user *etodo.User, userRoles *etodo.UsersRoles) error {

	//проверка имени на уникальность
	var count int64
	r.dbx.Where("username = ?", user.Username).Find(&user).Count(&count)
	if count > 0 {
		strErr := fmt.Sprintf("Username не уникален - %s", user.Username)
		//log.Println(strErr)
		return errors.New(strErr)
	}
	r.dbx.Create(&user)
	/* log.Println(user.UserId)
	log.Println(user)
	log.Println(userRoles) */

	//r.dbx.SetupJoinTable(&etodo.User{}, "Roles", &etodo.UserlistToRoles{})
	//var id int
	roles := new(etodo.Roles)
	//links := []etodo.UserlistToRoles{}
	links := etodo.UserlistToRoles{UserId: user.UserId, RoleId: roles.RoleId}

	for _, role := range userRoles.Roles {

		r.dbx.Table("userlist").Select("user_id").Where("user_id = ?", user.UserId).Scan(&links.UserId)
		r.dbx.Table("roles").Select("role_id").Where("rolesname = ?", role).Scan(&links.RoleId)
		r.dbx.Create(&links)

		/* r.dbx.Model(&links).Select("userlist.user_id", "roles.role_id").
		Where("userlist.user_id = ?", user.UserId).
		Where("roles.rolesname = ?", role).
		Create(&links) */

		/* 	r.dbx.Table("roles").Select("role_id", id).Where("rolesname = ?", role)
		log.Println(id) */
		/* r.dbx.Model(&links).Select("userlist.user_id", "roles.role_id").
		Joins("INNER JOIN userlist ON userlist.user_id = ?", user.UserId).
		Joins("INNER JOIN roles ON roles.role_id WHERE roles.rolesname = ?", role).
		Create(&links)
		*/

		/* 	r.dbx.Model(&links).Select("userlist.user_id", "roles.role_id").
		Joins("INNER JOIN userlist ON userlist.user_id = ?", user.UserId).
		Joins("INNER JOIN roles ON roles.rolesname = ?", role).
		Create(&links)
		*/
		/* 	r.dbx.Model(&etodo.UserlistToRoles{}).
		Select("userlist.user_id").
		Joins("INNER JOIN userlist ON userlist.user_id = ?", user.UserId).
		Select("roles.role_id").
		Joins("INNER JOIN roles ON roles.rolesname = ?", role).
		Create(&links) */

	}
	log.Println(links)

	//r.dbx.Where(etodo.UserlistToRoles{UserId: user.UserId, RoleId: roles.RoleId}).Find(&links)

	/* user:=User{
		user.Username : "61",
		user.Roles: []Roles{
			{RoleId : 1},
			{RolesName : "admin"},
		},
	} */

	//добавление нового пользователя в бд
	/* userlistToRoles := new(etodo.UserlistToRoles)

	//err := r.dbx.Model(&user).Association("roles").Error
	//log.Println(err)

	for _, role := range userRoles.Roles {

		 r.dbx.Model(&etodo.UserlistToRoles{}).Select("userlist.user_id", "roles.role_id").
			Joins("INNER JOIN userlist ON userlist.user_id = ?", user.UserId).
			Joins("INNER JOIN roles ON roles.rolesname = ?", role).
			Create(&userlistToRoles) */

	/*
		r.dbx.Table("userlist").Select("user_id").Where("user_id = ?", user.UserId).Scan(&userlistToRoles.UserId)
		r.dbx.Table("roles").Select("role_id").Where("rolesname = ?", role).Scan(&userlistToRoles.RoleId)
		r.dbx.Create(&userlistToRoles) */
	//}

	/* result := r.dbx.Model(&etodo.Roles{}).Select("roles.rolesname").
	Joins("INNER JOIN userlist_to_roles ON userlist_to_roles.user_id = ? AND userlist_to_roles.role_id = roles.role_id", userId).
	Scan(&roles) */

	/* INSERT INTO userlist_to_roles
	(SELECT userlist.user_id, roles.role_id FROM userlist, roles WHERE userlist.user_id = ? AND roles.rolesname = ?) */
	//query = fmt.Sprintf("INSERT INTO userlist_to_roles (SELECT userlist.user_id, roles.role_id FROM userlist, roles WHERE userlist.user_id = %v AND roles.rolesname = %v)", user.UserId, role)

	//query := "INSERT INTO userlist (firstname, lastname, username, password) VALUES ($1, $2, $3, $4) RETURNING user_id"
	//r.dbx.Table("userlist").Raw(query, user.FirstName, user.LastName, user.Username, user.Password).Scan(&user.UserId)
	//user1 := etodo.User{FirstName: user.FirstName, LastName: user.LastName, Username: user.Username, Password: user.Password}
	//log.Println(user1)
	//userlistToRoles := new(etodo.UserlistToRoles)
	//roles := new(etodo.Roles)

	/* subQerry1 := r.dbx.Select(user, "user_id").Where(user, "user_id = ?", user.UserId)
	subQerry2 := r.dbx.Select(roles, "role_id").Where(roles, "rolesname = ?", role)
	r.dbx.Table("userlist_to_roles").Create(userlistToRoles).Where(subQerry1).Where(subQerry2)
	*/
	/* r.dbx.Where(
	r.dbx.Table("?", r.dbx.Model(&user).Select("user_id").Where("user_id = ?", user.UserId)).
		Or(r.dbx.Table("?", r.dbx.Model(&roles).Select("role_id").Where("user_id = ?", role)))).
	Create(&userlistToRoles) */

	/* 	r.dbx.Table("userlist_to_roles").Select("userlist.user_id, roles.role_id", r.dbx.Joins("JOIN userlist ON userlist.user_id = ?", user.UserId).
	Joins("JOIN roles ON roles.rolesname = ?", role)).Create(&userlistToRoles) *
	r.dbx.Table("userlist").Select("user_id").Where("user_id = ?", user.UserId).Scan(&userid)
	r.dbx.Table("roles").Select("role_id").Where("rolesname = ?", role).Scan(&roleid)
	userlistToRoles := etodo.UserlistToRoles{UserId: userid, RoleId: roleid}
	r.dbx.Create(&userlistToRoles)
	*/

	//r.dbx.Table("userlist").Select("user_id").Where("user_id = ?", user.UserId).Scan(&userid)
	//r.dbx.Table("roles").Select("role_id").Where("rolesname = ?", role).Scan(&roleid)
	//userlistToRoles := etodo.UserlistToRoles{UserId: userid, RoleId: roleid}
	//r.dbx.Where("user_id = ?", r.dbx.Table("userlist").Select("user_id").Where("user_id = ?", user.UserId)).
	//	Where(r.dbx.Table("roles").Select("role_id").Where("rolesname = ?", role)).Create(&userlistToRoles)
	//r.dbx.

	/*
			INSERT INTO userlist_to_roles
			   	(SELECT userlist.user_id, roles.role_id FROM userlist, roles WHERE userlist.user_id = ? AND roles.rolesname = ?)


		for _, role := range userRoles.Roles {
			//r.dbx.Where(user, "user_id = ?", user.UserId).Find(userlistToRoles)
			//r.dbx.Where(roles, "rolesname = ?", role).Find(roles)

			r.dbx.Where().Create(userlistToRoles)
			log.Println(userlistToRoles)
			log.Println(roles)
		}

		/* for _, role := range userRoles.Roles {
			r.dbx.Where(user, "user_id = ?", user.UserId).Where(roles, "rolesname = ?", role).Create(&userlistToRoles)
			log.Println(role)
		} */
	/*

			  // r.dbx.Create(&user)


			   /*
		   		//добавление роли пользователя в бд
		   		for _, role := range userRoles.Roles {
		   			//r.dbx.Raw("INSERT INTO userlist_to_roles (SELECT userlist.user_id, roles.role_id FROM userlist, roles WHERE userlist.user_id = 22 AND roles.rolesname = 'admin')") //, user.UserId, role)
		   			//r.dbx.Table("userlist_to_roles", "roles").Create()      Raw("INSERT INTO userlist_to_roles (SELECT userlist.user_id, roles.role_id FROM userlist, roles WHERE userlist.user_id = ? AND roles.rolesname = '?')", user.UserId, role)

		   			userlistToRoles := UserlistToRoles{user_id: 25, role_id: 1}
		   			r.dbx.Table("userlist_to_roles").Select("user_id", "role_id").Create(&userlistToRoles)

		   			log.Println(role)
		   			/* if err != nil {
		   				return err
		   			}
		   		}*/
	fmt.Println("создан пользователь - ", user)

	return nil
}

//аутентификация пользователя
func (r *AuthPostgres) GetUser(username, password string) (*etodo.User, error) {
	user := new(etodo.User)
	result := r.dbx.Find(user, "username = ? AND password = ?", username, password)
	if result.RowsAffected == 0 || errors.Is(result.Error, gorm.ErrRecordNotFound) {
		strErr := "пользователь с таким логином и паролем не найден"
		log.Println(strErr)
		return nil, errors.New(strErr)
	}
	log.Println("аутентификация пользователя -", user)
	return user, nil
}

//удаление текущего пользователя
func (r *AuthPostgres) DeleteUser(userId int) error {

	user := new(etodo.User)
	r.dbx.Where("user_id = ?", userId).Delete(&user)

	result := r.dbx.Where("user_id = ?", userId).Find(&user)
	if result.RowsAffected < 1 {
		log.Println("пользователь с ( id -", userId, ") удален")
		return nil
	} else {
		strErr := fmt.Sprintf("ошибка удаления (id - %v, username - %v)", user.UserId, user.Username)
		log.Println(strErr)
		return errors.New(strErr)
	}

}

func (r *AuthPostgres) UpdateUser(user *etodo.User) error {

	if user.FirstName != "" {
		result := r.dbx.Table("userlist").Where(" user_id = ?", user.UserId).Update("firstname", user.FirstName)
		if result.RowsAffected < 1 {
			strErr := "ошибка изменения firstname"
			return errors.New(strErr)
		}
	}

	if user.LastName != "" {
		result := r.dbx.Table("userlist").Where(" user_id = ?", user.UserId).Update("lastname", user.LastName)
		if result.RowsAffected < 1 {
			strErr := "ошибка изменения lastname"
			return errors.New(strErr)
		}
	}
	if user.Password != "" {
		result := r.dbx.Table("userlist").Where(" user_id = ?", user.UserId).Update("password", user.Password)
		if result.RowsAffected < 1 {
			strErr := "ошибка изменения password"
			return errors.New(strErr)
		}
	}

	log.Println("изменены дaнные пользователя -", user.UserId)
	return nil
}

//получение роли из бд для middleware
func (r *AuthPostgres) GetRoles(userId int) []string {
	var roles []string

	//result := r.dbx.Table("userlist_to_roles", "roles").Raw("SELECT roles.rolesname FROM userlist_to_roles,	roles WHERE userlist_to_roles.user_id = $1 AND userlist_to_roles.role_id = roles.role_id", userId).Scan(&roles)

	result := r.dbx.Model(&etodo.Roles{}).Select("roles.rolesname").
		Joins("INNER JOIN userlist_to_roles ON userlist_to_roles.user_id = ? AND userlist_to_roles.role_id = roles.role_id", userId).
		Scan(&roles)

	//log.Println("roles !!!!!!!!!!!!!!!!! -", roles)

	//	result := r.dbx.Table("roles").Table("userlist_to_roles").Select("roles.rolesname").Where("userlist_to_roles.user_id = ?", userId).Where("userlist_to_roles.role_id = roles.role_id").Scan(&roles)

	/* sql1 := r.dbx.Table("userlist_to_roles").Select("role_id")
	sql2 := r.dbx.Table("roles").Select("role_id")
	sql3 := r.dbx.Table("userlist_to_roles").Select("user_id")

	result := r.dbx.Table("roles").Select("rolesname").
		Where("? = ?", sql1, sql2).
		Where("? = ?", sql3, userId).
		Scan(&roles) */

	/* result := r.dbx.Table("roles").Select("rolesname").
	Where("? = ?", r.dbx.Table("userlist_to_roles").Select("role_id"), r.dbx.Table("roles").Select("role_id")).
	Where("? = ?", r.dbx.Table("userlist_to_roles").Select("user_id"), userId).
	Scan(&roles) */

	//userlist_to_roles.user_id = $1 AND userlist_to_roles.role_id = roles.role_id", userId).Scan(&roles)

	if result.RowsAffected == 0 || errors.Is(result.Error, gorm.ErrRecordNotFound) {
		strErr := "пользователь с таким логином и паролем не найден"
		log.Println(strErr)
		return nil
	}
	//log.Println("!!!!!!!!!!!!!!!!!!!!!!! - ", roles)
	return roles
}
