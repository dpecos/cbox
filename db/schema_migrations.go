package db

import migrate "github.com/rubenv/sql-migrate"

func migrations() *migrate.MemoryMigrationSource {
	return &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{

			&migrate.Migration{
				Id: "20180406",
				Up: []string{
					`create table commands (
						id integer not null primary key,
						cmd text not null,
						title text not null,
						description text,
						url text,
						created_at timestamp default current_timestamp,
						updated_at timestamp default current_timestamp
					)`,
				},
				Down: []string{
					"drop table commands",
				},
			},

			&migrate.Migration{
				Id: "20180409",
				Up: []string{
					`create table tags (
						name text primary key
					)`,
					`create table command_tags (
						command integer, 
						tag text,
						primary key (command, tag),
						foreign key(command) references commands(id),
						foreign key(tag) references tags(name)
					)`,
				},
				Down: []string{
					"drop table tags",
					"drop table command_tags",
				},
			},
		},
	}
}
