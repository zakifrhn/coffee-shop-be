CREATE TABLE coffeshop.bridge_product_size (
	id_bridge_ps serial4 NOT NULL,
	id_product uuid NOT NULL,
	id_size uuid NOT NULL
);


-- coffeshop.bridge_product_size foreign keys

ALTER TABLE coffeshop.bridge_product_size ADD CONSTRAINT bridge_product_size_id_product_fkey FOREIGN KEY (id_product) REFERENCES coffeshop.product(id_product) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE coffeshop.bridge_product_size ADD CONSTRAINT bridge_product_size_id_size_fkey FOREIGN KEY (id_size) REFERENCES coffeshop."size"(id_size);