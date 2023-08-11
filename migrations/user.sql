CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE coffeshop.user (
    id_user uuid NULL DEFAULT uuid_generate_v4(),
    email varchar(255) NOT NULL,
    pass varchar(255) NOT NULL,
    phone_number varchar(20),
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    updated_at timestamp without time zone NULL,
    CONSTRAINT user_pk PRIMARY KEY (id_user)
);