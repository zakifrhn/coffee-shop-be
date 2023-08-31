CREATE TABLE coffeshop.size (
	id_size uuid NOT NULL DEFAULT gen_random_uuid(),
	uk_product varchar(5) NOT NULL,
	CONSTRAINT size_pk PRIMARY KEY (id_size)
);