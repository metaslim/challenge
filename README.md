# Simple Code Challenge

Currently this app is using Go internal library (encoding/json) which use reflection. In the future json file parsing can be improved by using https://github.com/mailru/easyjson (without reflection).

## Setting up locally

#### GOPATH and installing the code
```
$ export GOPATH=/Users/someone/Development/gocode
$ export PATH=$GOPATH/bin:$PATH
$ echo $GOPATH
/Users/someone/Development/gocode
$ go get go get github.com/metaslim/challenge
$ cd $GOPATH/src/github.com/metaslim/challenge
$ pwd
/Users/someone/Development/gocode/src/github.com/metaslim/challenge
$ git status
On branch master
Your branch is up to date with 'origin/master'.

nothing to commit, working tree clean
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

#### How it looks like

##### Main Menu
![Search menu](search-menu.png)

##### Get Help
![Help](search-help.png)

##### Display Search Field
![Display search field](search-field.png)

##### Search and display in compact table
![Display search result in compact table](search-table.png)

##### Search and display in json
![Display search result in colored json](search-json.png)

##### Exit
![Quit](search-quit.png)

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
