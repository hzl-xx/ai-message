package model

type User struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
	Phone string `json:"phone"`
	Name string `json:"name"`
	Age uint8 `json:"age"`
}
// 设置表明
//func (u User) TableName() string {
//	return "user"
//
//}

func (u *User)CheckAuth(username, password string) bool {
	db.Select("id").Where(User{Username : username, Password : password}).First(u)
	if u.ID > 0 {
		return true
	}

	return false
}

func (u *User)InsertUser() error {
	return db.Create(u).Error
}

func (u *User)GetUser(pagenum int, pagesize int, params interface{}) (user []User) {
	db.Where(params).Offset(pagenum).Limit(pagesize).Find(&user)
	return
}

func (u *User)GetUserTotal(params interface{}) (count int) {
	db.Model(&User{}).Where(params).Count(&count)
	return
}

func (u *User)DeleteUser(id int) bool {
	db.Where("id=?", id).Delete(&User{})
	return true
}

func (u *User)EditUser(id int, data interface{}) bool {
	db.Model(&User{}).Where("id=?", id).Update(data)
	return true
}
