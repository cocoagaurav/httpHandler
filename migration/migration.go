package migration

import "github.com/rubenv/sql-migrate"

func Getmigration() migrate.MemoryMigrationSource {
	var Migration = &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			&migrate.Migration{
				Id: "123",
				Up: []string{`  
								create table user(
								name 	varchar(50) NOT NULL,
								UID   	int         NOT NULL,
								age   	int     	NOT NULL,
								PRIMARY KEY(UID)
								)`,
					`create table post(
								id			int 			NOT NULL,
								title		varchar(50) 	NOT NULL,
								discription	varchar(50) 	NOT NULL,
								FOREIGN KEY (id) REFERENCES user(UID)
								)`},
				Down: []string{`
									drop table user;
									drop table post;
								`},
			},
		},
	}

	return *Migration
}
