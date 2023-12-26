dev: 
	weaver generate .
	go build -o bin/main .
	weaver multi deploy weaver.toml

db:
	cat resources/sql/create.sql | sqlite3 dev.db
