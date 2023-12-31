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

func (r *RepoProduct) CreateProduct(data *models.ProductSet) (string, error) {

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
			VALUES($1,$2)`, &idProduct, &data.Category[i])
	}

	for i := range data.Size {
		r.MustExec(`INSERT INTO coffeshop."bridge_product_size"(
			id_product,
			id_size)
			VALUES($1,$2)`, &idProduct, &data.Size[i])
	}

	fmt.Printf("Inserted product with ID: %s\n", idProduct)
	return "Succees 1 Data product Added", errProd
}

func (r *RepoProduct) UpdateProduct(data *models.ProductSet) (string, error) {

	query := `UPDATE coffeshop."product" 
	SET
	desc_product = $2,
	name_product = $3,
	banner_product = $4,
	price = $5,
	isfavorite = $6,
	updated_at = now()
	where id_product = $1`

	_, errProd := r.Exec(query, data.Id_product, data.Desc_product, data.Name_product, data.Banner_product, data.Price, data.Isfavorite)
	if errProd != nil {
		fmt.Println(errProd)
	}

	_, errCtg := r.Exec(`DELETE FROM coffeshop."bridge_product_category" WHERE id_product = $1;`, data.Id_product)
	fmt.Println(errCtg)
	if errCtg != nil {
		fmt.Println(errCtg)
	}

	_, errSize := r.Exec(`DELETE FROM coffeshop."bridge_product_size" WHERE id_product = $1;`, data.Id_product)
	fmt.Println(errSize)
	if errSize != nil {
		fmt.Println(errSize)
	}

	for i := range data.Category {
		r.MustExec(`INSERT INTO coffeshop."bridge_product_category"(
			id_product,
			id_category)
			VALUES($1,$2)`, data.Id_product, &data.Category[i])
	}

	for i := range data.Size {
		r.MustExec(`INSERT INTO coffeshop."bridge_product_size"(
			id_product,
			id_size)
			VALUES($1,$2)`, data.Id_product, &data.Size[i])
	}

	fmt.Printf("Updated product with ID: %s\n", data.Id_product)
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
						string_agg(c.name_category, ', ') as category,
						string_agg(s.uk_product, ', ') as size
						FROM coffeshop.product p

						JOIN coffeshop.bridge_product_category bgpc ON bgpc.id_product = p.id_product
						JOIN coffeshop.category c ON bgpc.id_category = c.id_category
						JOIN coffeshop.bridge_product_size bgps ON bgps.id_product = p.id_product
						JOIN coffeshop.size s ON bgps.id_size = s.id_size
						 WHERE TRUE %s %s
		GROUP BY p.id_product LIMIT $1 OFFSET $2`, search, category)

	var products []models.Product
	err := r.Select(&products, query, limit, offset)
	if err != nil {
		return nil, err
	}
	fmt.Println(products)
	return products, nil
}
