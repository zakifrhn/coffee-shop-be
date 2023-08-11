CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

--SELECT * FROM pg_extension WHERE extname = 'uuid-ossp';

CREATE TABLE coffeshop.product (
    id_product uuid NOT NULL DEFAULT uuid_generate_v4(),
    desc_product varchar(255) NOT null,
    name_product varchar(255) not null,
    banner_product varchar(255),
    price varchar(50),
    isFavorite bool,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NULL,
); 

SELECT * FROM product

CREATE TABLE coffeshop.user (
    id_user uuid NULL DEFAULT uuid_generate_v4(),
    email varchar(255) NOT NULL,
    pass varchar(255) NOT NULL,
    phone_number varchar(20),
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NULL,
    CONSTRAINT user_pk PRIMARY KEY (id_user)
);

CREATE TABLE coffeshop.size (
    id_size uuid NOT NULL DEFAULT uuid_generate_v4(),
    uk_product varchar(5) NOT null,
    CONSTRAINT size_pk PRIMARY KEY (id_size)
);

create table coffeshop.bridge_product_size(
	id_bridge_ps serial4 not null,
	id_product uuid NOT NULL,
	id_size uuid NOT NULL
)

create table coffeshop.category(
    id_category uuid NOT NULL DEFAULT uuid_generate_v4(),
    name_category varchar(50) NOT null,
    CONSTRAINT category_pk PRIMARY KEY (id_category)
)

create table coffeshop.bridge_product_category(
	id_bridge_pc serial4 not null,
	id_product uuid NOT NULL,
	id_category uuid NOT NULL
)

alter table coffeshop.bridge_product_category add foreign key (id_category) references "category" (id_category)
alter table coffeshop.bridge_product_category add foreign key (id_product) references "product" (id_product)