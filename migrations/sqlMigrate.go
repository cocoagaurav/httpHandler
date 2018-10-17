package migrations

import "github.com/rubenv/sql-migrate"

func init() {

	Migration := &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			&migrate.Migration{
				Id: "123",
				Up: []string{`create table user(
								name 	varchar(50) 	notnull,
								UID   	int         	notnull,
								age   	int     		notnull,
								primarykey(UID),
							);
								create table post(
										id  int notnull,
										title varchar(30), 
										discription varchar(100),
										FOREIGN KEY (id) REFERENCES user(UID),
								);

										
							`},
				Down: []string{`
									drop table user;
									drop table post;
								`},
			},
		},
	}
}


