--the script to remove all tables in the database
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS chatting CASCADE;


create table users
(
    id uuid,
    email character varying(200) not null,
    password_digest character varying(1000) not null,
    name character varying(255) not null,
    self_introduction character varying(100),
    picture text,
    CONSTRAINT "users_pk" PRIMARY KEY (id)
);

create table chatting
(
    id uuid,
    sender_id uuid not null,
    receiver_id uuid not null,
    content text not null,
    CONSTRAINT "chatting_pk" PRIMARY KEY (id)
);

ALTER TABLE users ADD CONSTRAINT users_u1 UNIQUE (email);
