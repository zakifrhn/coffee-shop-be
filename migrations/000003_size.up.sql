CREATE TABLE coffeshop.size (
	id_size uuid NOT NULL DEFAULT coffeshop.uuid_generate_v4(),
	uk_product varchar(5) NOT NULL,
	CONSTRAINT size_pk PRIMARY KEY (id_size)
);