ALTER DATABASE chatting_db SET timezone TO 'UTC';

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
REVOKE CREATE ON SCHEMA public FROM PUBLIC;

GRANT USAGE ON SCHEMA public to chatting_admin;
GRANT CREATE ON SCHEMA public to chatting_admin;

/* grant the schema access privilege to normal users. Without schema right, user will unable to see the tables. */
GRANT USAGE ON SCHEMA public to chatting_user;
GRANT USAGE ON SCHEMA public to chatting_readonly;