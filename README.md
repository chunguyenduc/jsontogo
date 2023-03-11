# jsontogo

[![Deployment](https://github.com/chunguyenduc/jsontogo/actions/workflows/ci.yml/badge.svg?event=push)](https://github.com/chunguyenduc/jsontogo/actions/workflows/ci.yml)
![Coverage](https://github.com/chunguyenduc/jsontogo/blob/badge/badge.svg?branch=badge)

> A CLI to convert JSON to Go struct

Inspired by [mholt](https://github.com/mholt/json-to-go). Written in Go

```
go install github.com/chunguyenduc/jsontogo
```

# Usage
 ```
 jsontogo - a CLI to convert JSON to Go struct

Usage:
  jsontogo [flags]

Flags:
  -f, --file_input string    read input from JSON file
  -o, --file_output string   write output to Go file
  -h, --help                 help for jsontogo
  -n, --name string          name of struct
 ```
# Examples
```
$ jsontogo '{"login":"my_login","password":"my_password"}'

type Autogenrated struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

``` 
Custom struct name, with nested JSON

```
$ jsontogo '{"user":{"fullname":"Bob Chan","email":"bob@email.com","password":"password"}}' -n User

type User struct {
	User struct {
		Fullname string `json:"fullname"`
		Email string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

``` 

We can also convert from JSON file and output to file
```
$ jsontogo -f data.json -o data 
```

# Contact
[Duc Chu](https://www.linkedin.com/in/nguyenducchu1999/)

# License
[MIT License](LICENSE)




