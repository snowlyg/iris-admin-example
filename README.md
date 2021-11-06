# iris-admin-example

This project is a example for [iris-admin](https://github.com/snowlyg/iris-admin)

### Initialize 
- Initialize project configs

```sh
go run cmd/main.go init
```

### Migrate 
- Exec migrate cmd
- with `-s` flag will seed datas into database after exec migrate command.  

```sh
go run cmd/main.go migrate -s
```

### Go run 
- Run project 
```sh
go run main.go 
```

### Open `http://127.0.0.1:8085` in browser
