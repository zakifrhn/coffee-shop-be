CREATE TABLE coffeshop.category (
	id_category uuid NOT NULL DEFAULT gen_random_uuid(),
	name_category varchar(50) NOT NULL,
	CONSTRAINT category_pk PRIMARY KEY (id_category)
);