CREATE TABLE coffeshop.product (
	id_product uuid NOT NULL DEFAULT gen_random_uuid(),
	desc_product varchar(255) NOT NULL,
	name_product varchar(255) NOT NULL,
	banner_product varchar(255) NULL,
	price varchar(50) NULL,
	isfavorite bool NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	CONSTRAINT product_pk PRIMARY KEY (id_product)
);