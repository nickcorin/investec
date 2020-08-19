<p align="center">
<h1 align="center">Ziggy</h1>
<p align="center">Unofficial Go client library for the Investec OpenAPI</p>
</p>
<p align="center">
<p align="center"><a href="https://github.com/nickcorin/ziggy/actions?query=workflow%3AGo"><img src="https://github.com/nickcorin/ziggy/workflows/Go/badge.svg?branch=master" alt="Build Status"></a> <a href="https://goreportcard.com/report/github.com/nickcorin/ziggy"><img src="https://goreportcard.com/badge/github.com/nickcorin/ziggy?style=flat-square" alt="Go Report Card"></a> <a href="http://godoc.org/github.com/nickcorin/ziggy"><img src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square" alt="GoDoc"></a> <a href="LICENSE"><img src="https://img.shields.io/github/license/nickcorin/ziggy" alt="License"></a></p>
</p>
<p align="center">
<!-- <img src="/images/snorlax.jpg" /> -->
</p>

## Installation

To install `ziggy`, use `go get`:
```
go get github.com/nickcorin/ziggy
```

Import the `ziggy` package into your code:
```golang
package main

import "github.com/nickcorin/ziggy"

func main() {
	client := ziggy.New(nil)
}
```

## Usage

#### Creating a simple client.
```golang
client := ziggy.New(nil)
```

#### Configuring the client using `ClientOptions`.
```golang
client := ziggy.New(&ziggy.ClientOptions{
		ClientID: "MyClientID",
		ClientSecret: "z1ggYS3creT",
	}
)
```

#### Obtaining an access token.
```golang
token, err := z.GetAccessToken(context.Background(), ziggy.TokenScopeAccounts)
if err != nil {
	log.Fatal(err)
}
```

#### Fetching a list of your accounts.
```golang
accounts, err := z.GetAccounts(context.Background())
if err != nil {
	log.Fatal(err)
}
```

#### ...and then fetching an account's balance.
```golang
accounts, err := z.GetAccounts(context.Background())
if err != nil {
	log.Fatal(err)
}

balance, err := z.GetBalance(context.Background(), accounts[0].ID)
if err != nil {
	log.Fatal(err)
}
```

## Contributing
Please feel free to submit issues, fork the repositoy and send pull requests!

## License
This project is licensed under the terms of the MIT license.
