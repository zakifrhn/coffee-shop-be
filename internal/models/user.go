package models

import "time"

type User struct {
	Id_user    string     `db:"id_user" form:"id_user" json:"id_user" uri:"id_user"`
	Email_user string     `db:"email" form:"email" json:"email"`
	Pass_user  string     `db:"pass" form:"pass" json:"pass"`
	Role       string     `db:"role" json:"role" form:"role"`
	Phone_user string     `db:"phone_number" form:"phone_number" json:"phone_number"`
	CreatedAt  *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at"`
}

type Meta struct {
	Page   int
	Limit  int
	Offset int
}
