CREATE TABLE coffeshop.category (
	id_category uuid NOT NULL DEFAULT coffeshop.uuid_generate_v4(),
	name_category varchar(50) NOT NULL,
	CONSTRAINT category_pk PRIMARY KEY (id_category)
);