ALTER TABLE chatting ADD CONSTRAINT chatting_fk1 FOREIGN KEY (sender_id) REFERENCES users (id) MATCH FULL ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE chatting ADD CONSTRAINT chatting_fk2 FOREIGN KEY (receiver_id) REFERENCES users (id) MATCH FULL ON DELETE CASCADE ON UPDATE CASCADE;