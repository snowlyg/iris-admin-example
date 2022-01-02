# Example for iris with iris-admin

### Initialize 
- Initialize project 
```sh
 go run iris/cmd/migrate/main.go init
```

### Refresh 
- Refresh project 
```sh
 go run iris/cmd/migrate/main.go refresh
```

### Migrate Databases
- Migrate project databases
```sh
 go run iris/cmd/migrate/main.go migrate
```

### Rollback Databases
- Rollback project databases
```sh
 go run iris/cmd/migrate/main.go rollback
```

**Notice**
- When you exec migrate or refresh command, it's will seed data to database.
- If you just exepct create tables to your database and don't want to seed datas into this tables, you need to use `--seed=false` or `-s=false`. 


### Go run 
- Run project 
```sh
 go run iris/main.go 
 # or
 go build -o example iris/main.go 
 ./example
```

### Open `http://127.0.0.1:8085` in browser
