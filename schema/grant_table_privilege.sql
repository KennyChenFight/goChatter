/*for normal tables */
GRANT SELECT, INSERT, UPDATE, DELETE, REFERENCES ON TABLE users to chatting_user;
GRANT SELECT, INSERT, UPDATE, DELETE, REFERENCES ON TABLE chatting to chatting_user;

GRANT SELECT ON TABLE users to chatting_readonly;
GRANT SELECT ON TABLE chatting to chatting_readonly;