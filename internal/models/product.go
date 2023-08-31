package models

import (
	"time"
)

type Product struct {
	Id_product     string     `db:"id_product" form:"id_product" json:"id_product" uri:"id_product"`
	Desc_product   string     `db:"desc_product" form:"desc_product" json:"desc_product"`
	Name_product   string     `db:"name_product" form:"name_product" json:"name_product"`
	Banner_product string     `db:"banner_product" form:"banner_product" json:"banner_product"`
	Price          string     `db:"price" form:"price" json:"price"`
	Isfavorite     bool       `db:"isfavorite" form:"isfavorite" json:"isfavorite"`
	CreatedAt      *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updated_at"`
	NameCtg        string     `db:"category" json:"name_category" `
	NameSize       string     `db:"size" json:"name_size"`
}

type ProductSet struct {
	Id_product     string     `db:"id_product" form:"id_product" json:"id_product" uri:"id_product"`
	Desc_product   string     `db:"desc_product" form:"desc_product" json:"desc_product"`
	Name_product   string     `db:"name_product" form:"name_product" json:"name_product"`
	Banner_product string     `db:"banner_product" json:"banner_product"`
	Price          string     `db:"price" form:"price" json:"price"`
	Isfavorite     bool       `db:"isfavorite" form:"isfavorite" json:"isfavorite"`
	CreatedAt      *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updated_at"`
	Category       []string   `db:"bridge_product_category" form:"category" json:"category"`
	Size           []string   `db:"bridge_product_size" form:"size" json:"size"`
}

type Category struct {
	Id_product  string `db:"id_product" form:"id_product" json:"id_product"`
	Id_category string `db:"id_category" form:"id_category" json:"id_category"`
}
type Size struct {
	Id_product string `db:"id_product" form:"id_product" json:"id_product"`
	Id_size    string `db:"id_size" form:"id_size" json:"id_size"`
}

type Metas struct {
	Page     int
	Limit    int
	Name     string
	Category string
}

type Products []Product
