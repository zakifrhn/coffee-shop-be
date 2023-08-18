package models

import "time"

type User struct {
	Id_user    string     `db:"id_user" form:"id_user" json:"id_user" uri:"id_user" valid:"-"`
	Email_user string     `db:"email" form:"email" json:"email" valid:"email, required~please input in type email"`
	Pass_user  string     `db:"pass" form:"pass" json:"pass" valid:"alphanum, stringlength(8|15)~password harus lebih sama dengan dari 8"`
	Role       string     `db:"role" json:"role" form:"role" valid:"-"`
	Phone_user string     `db:"phone_number" form:"phone_number" json:"phone_number" valid:"numeric"`
	CreatedAt  *time.Time `db:"created_at" json:"created_at" valid:"-"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}

type Meta struct {
	Page   int
	Limit  int
	Offset int
}
