package repositories

import (
	"fmt"
	"inter/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoProduct struct {
	*sqlx.DB
}

func NewProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{db}
}

// type category struct {
// 	NameCategory string `db:"name_category" json:"name_category" `
// }

func (r *RepoProduct) CreateProduct(data *models.Product) (string, error) {

	query := `INSERT INTO coffeshop."product" (
		desc_product,
		name_product,
		banner_product,
		price,
		isfavorite
	) VALUES (
		$1,$2, $3,$4, $5
	) RETURNING id_product`
	var idProduct string

	errProd := r.Get(&idProduct, query, data.Desc_product, data.Name_product, data.Banner_product, data.Price, data.Isfavorite)
	if errProd != nil {
		fmt.Println(errProd)
	}

	for i := range data.Category {
		r.MustExec(`INSERT INTO coffeshop."bridge_product_category"(
			id_product,
			id_category)
			VALUES($1,$2)`, &idProduct, &data.Category[i].Id_category)
	}

	for i := range data.Size {
		r.MustExec(`INSERT INTO coffeshop."bridge_product_size"(
			id_product,
			id_size)
			VALUES($1,$2)`, &idProduct, &data.Size[i].Id_size)
	}

	fmt.Printf("Inserted product with ID: %s\n", idProduct)
	return "Succees 1 Data product Added", errProd
}

func (r *RepoProduct) UpdateProduct(data *models.Product) (string, error) {

	// query := `UPDATE coffeshop."product"
	// 			SET
	// 		desc_product=:desc_product,
	// 		name_product=:name_product,
	// 		banner_product=:banner_product,
	// 		price=:price,
	// 		isfavorite=:isfavorite,
	// 		updated_at = now()
	// 		 WHERE
	// 		 id_product = :id_product;`
	// _, er := r.NamedExec(query, data)
	// if er != nil {
	// 	fmt.Print("ini errornya", er)
	// 	return "", er
	// }

	// return "1 data has been updated", nil

	query := `UPDATE coffeshop."product" 
		SET
		desc_product = $2,
		name_product = $3,
		banner_product = $4,
		price = $5,
		isfavorite = $6,
		updated_at = now()
		where id_product = $1`
	//var idProduct string

	_, errProd := r.Exec(query, data.Id_product, data.Desc_product, data.Name_product, data.Banner_product, data.Price, data.Isfavorite)
	if errProd != nil {
		fmt.Println(errProd)
	}

	for i := range data.Category {
		r.MustExec(`UPDATE coffeshop."bridge_product_category" SET
			id_category = $2
			where id_product = $1`, data.Id_product, &data.Category[i].Id_category)
	}

	for i := range data.Size {
		r.MustExec(`UPDATE coffeshop."bridge_product_size" SET
			id_size = $2
			where id_product= $1`, data.Id_product, &data.Size[i].Id_size)
	}

	fmt.Printf("Inserted product with ID: %s\n", data.Id_product)
	return "Succees 1 Data product Updated", errProd
}

func (r *RepoProduct) DeleteProduct(data *models.Product) (string, error) {
	query := `DELETE FROM coffeshop."product" WHERE id_product = :id_product;`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}

	return "1 data product has been Deleted", nil
}

func (r *RepoProduct) GetAllProduct(data *models.Product, page, limit int) ([]models.Product, error) {

	offset := (page - 1) * limit

	query := fmt.Sprintf(`SELECT * FROM coffeshop.product WHERE isfavorite = true ORDER BY "id_product" LIMIT $1 OFFSET $2`)

	var products []models.Product
	err := r.Select(&products, query, limit, offset)
	if err != nil {
		return nil, err
	}
	fmt.Println(err)
	return products, nil
}

func (r *RepoProduct) GetCategory(data *models.Product, page int, limit int, category string, search string) ([]models.Product, error) {

	offset := (page - 1) * limit
	if search == "" {
		search = ""
	} else {
		search = fmt.Sprintf(`AND LOWER(p.name_product) like LOWER('%s')`, "%"+search+"%")
	}

	if category == "" {
		category = ""
	} else {
		category = fmt.Sprintf(` AND LOWER(c.name_category) like LOWER('%s')`, "%"+category+"%")
	}

	query := fmt.Sprintf(`SELECT
						p.id_product,
						p.banner_product,
						p.name_product,
						p.price,
						string_agg(c.name_category, ',') as category
						FROM coffeshop.product p

						JOIN coffeshop.bridge_product_category bgpc ON bgpc.id_product = p.id_product
						JOIN coffeshop.category c ON bgpc.id_category = c.id_category
						 WHERE TRUE %s %s
		GROUP BY p.id_product LIMIT $1 OFFSET $2`, search, category)

	var products []models.Product
	err := r.Select(&products, query, limit, offset)
	if err != nil {
		//fmt.Println(err)
		return nil, err
	}
	//fmt.Println(err)
	return products, nil
}

func (r *RepoProduct) GetNameProduct(data *models.Product, page, limit int, name string) ([]models.Product, error) {

	offset := (page - 1) * limit

	query := (`SELECT
						p.id_product,
						p.banner_product,
						p.name_product,
						p.price
						FROM coffeshop.product p

						JOIN coffeshop.bridge_product_category bgpc ON bgpc.id_product = p.id_product
						JOIN coffeshop.category c ON bgpc.id_category = c.id_category
						WHERE lower(p.name_product) LIKE $3
						 GROUP BY p.id_product LIMIT $1 OFFSET $2`)

	// if name != "" {
	// 	query += " AND p.name_product ILIKE '%'|| $3 || '%'"
	// }
	var products []models.Product
	err := r.Select(&products, query, limit, offset, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	fmt.Println(err)
	return products, nil
}
