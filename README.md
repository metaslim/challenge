# Simple Code Challenge

## Setting up locally

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

Currently this app is using Go internal library (encoding/json) which use reflection. In the future parsing json file should able to be improved with https://github.com/mailru/easyjson without reflection.