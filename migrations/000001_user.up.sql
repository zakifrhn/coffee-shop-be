CREATE TABLE coffeshop.user (
	id_user uuid NOT NULL DEFAULT gen_random_uuid(),
	email varchar(255) NOT NULL,
	pass varchar(255) NOT NULL,
	phone_number varchar(20) NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NULL,
	"role" varchar(50) NOT NULL,
	CONSTRAINT user_pk PRIMARY KEY (id_user)
);