Step 1: install PostgreSQL

    sudo apt-get install postgresql

Step 2: Creating the database and corresponding user account for the project

	Login the postgresql by either one of following command:
		1.	sudo -u postgres psql -f create_db.sql
			(It will perform OS authentication. It works only if the postgresql is installed in your local machine.)

		2.	psql -h <machine_name> -U <username> postgres -f create_db.sql
			(The user account should be superuser of postgresql server. Otherwise you are unlikely to have enough privilege.)


Step 2: Setup the privilege in the database.

	Run the following command:
		psql -h <machine_name> -U chatting_admin -f setup_db.sql chatting_db


Step 3: Create the table and foreign key in the database.

	Run the following command:
		psql -h <machine_name> -U chatting_admin -f create_table.sql chatting_db
		psql -h <machine_name> -U chatting_admin -f create_fk.sql chatting_db

Step 4: Grant table privilege to the users.

	Run the following command:
		psql -h <machine_name> -U chatting_admin -f grant_table_privilege.sql chatting_db