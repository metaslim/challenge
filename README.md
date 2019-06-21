# Simple Code Challenge

Currently this app is using Go internal library (encoding/json) which uses reflection. This can be improved by using library which does not use reflection such as easyjson (https://github.com/mailru/easyjson).

## Environment

Go 1.12

## Setting up locally

### GOPATH and installing the code
```
$ mkdir -p $HOME/Development/gocode
$ export GOPATH=$HOME/Development/gocode
$ go get github.com/metaslim/challenge
$ cd $GOPATH/src/github.com/metaslim/challenge
$ git status
On branch master
Your branch is up to date with 'origin/master'.

nothing to commit, working tree clean
```

### Install Dependencies
```sh
$ make install
```

## How to run tests

```sh
$ make test

```

## How to build binary

```sh
$ make build
```

## How to run binary

```sh
$ ./challenge
```

## How to run without compiling (slower compared to run binary)

```sh
$ make run
```

### Main Menu
![Search menu](readme-menu.png)

### Get Help
![Help](readme-help.png)

### Display Search Field
![Display search field](readme-field.png)

### Search and display in compact table
![Display search result in compact table](readme-table.png)

### Search and display in json
![Display search result in colored json](readme-json.png)

### Exit
![Quit](readme-quit.png)
