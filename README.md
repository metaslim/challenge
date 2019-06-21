# Simple Code Challenge

## How it looks like

#### Main Menu
![](menu.png)

#### Help
![](help.png)

#### Display Search Field
![](field.png)

#### Search and display in compact table
![](table.png)

#### Search and display in json
![](json.png)

#### Quit
![](quit.png)

## Setting up locally

#### GOPATH and location of the code
```
$ echo $GOPATH
/Users/someone/Development/gocode
$ pwd
/Users/someone/Development/gocode/src/github.com/metaslim/challenge
$ ls
Gopkg.lock Gopkg.toml Makefile   README.md  config     data       field.png  help.png   json.png   lib        main.go    menu.png   quit.png   table.png  vendor
```

#### Install Dependencies
```sh
make install
```

#### Run Tests
```sh
make test

```

#### Run App
```sh
make run
```

#### Command
```
"quit" to exit
"help" to get help

"describe-organizations" will return search fields for organizations
"describe-users" will return search fields for users
"describe-tickets" will return search fields for tickets

"table-organizations:tags=West" will return any organizations who has West in their Tags in compact table
"table-users:alias=Miss Coffey" will return any users whose alias is Miss Coffey in compact table
"table-tickets:status=pending" will return any tickets with pending status  in compact table

"search-organizations:tags=West" will return any organizations who has West in their Tags
"search-users:alias=Miss Coffey" will return any users whose alias is Miss Coffey
"search-tickets:status=pending" will return any tickets with pending status
```

Currently this app is using Go internal library (encoding/json) which use reflection. In the future json file parsing can be improved by using https://github.com/mailru/easyjson (without reflection).