/*
	chatting_admin is the database owner.
	It owns all objects(table, trigger, etc) in the database.
	Usually, it should be used ONLY when DBA wants to change the database schema.
	chatting_user is the account used by the golang executable.
	Thus, it is NOT allowed to create / change / any object in database.
	For normal table, only CURD privilege is granted, truncate table should NOT be granted.
	chatting_readonly is used during debugging.
	Trusted software developer will use this account to view the data in production database directly.
	Thus it should have select privilege.
*/


CREATE ROLE chatting_admin LOGIN PASSWORD 'admin_password' NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION;
CREATE ROLE chatting_user LOGIN PASSWORD 'user_password' NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION;
CREATE ROLE chatting_readonly LOGIN PASSWORD 'readonly_password' NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION;

CREATE DATABASE chatting_db with ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8' CONNECTION LIMIT = -1 template=template0;

ALTER DATABASE chatting_db OWNER TO chatting_admin;