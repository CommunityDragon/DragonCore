CREATE TABLE public."user" (
    id uuid DEFAULT uuid_generate_v4(),
    email varchar(255) NOT NULL,
    password bytea NOT NULL);

ALTER TABLE public."user" ADD CONSTRAINT user_pk PRIMARY KEY (id);

