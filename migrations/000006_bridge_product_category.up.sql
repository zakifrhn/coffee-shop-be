CREATE TABLE coffeshop.bridge_product_category (
	id_bridge_pc serial4 NOT NULL,
	id_product uuid NOT NULL,
	id_category uuid NOT NULL
);


-- coffeshop.bridge_product_category foreign keys

ALTER TABLE coffeshop.bridge_product_category ADD CONSTRAINT bridge_product_category_id_category_fkey FOREIGN KEY (id_category) REFERENCES coffeshop.category(id_category) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE coffeshop.bridge_product_category ADD CONSTRAINT bridge_product_category_id_product_fkey FOREIGN KEY (id_product) REFERENCES coffeshop.product(id_product) ON DELETE CASCADE ON UPDATE CASCADE;