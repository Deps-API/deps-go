# deps-go

> Idiomatic Go client for the Deps API.

The `deps-go` package is a lightweight and powerful client for interacting with the Deps API in any Go environment.

## Features

- **Idiomatic Go**: Clean, generated Go code that feels natural to use.
- **Typed**: Fully typed with Go structs for a better developer experience.
- **Lightweight**: Minimal dependencies.
- **Flexible**: Configure the HTTP client, timeout, and base URL with functional options.

## Installation

```sh
go get github.com/Deps-API/deps-go@v1.0.0
```

## `Client` API

The `Client` is the main entry point for interacting with the Deps API.

**Example**: Get the API status.

```go
package main

import (
	"context"
	"fmt"
	"log"

	depscian "deps-go"
)

func main() {
	client, err := depscian.NewClient("YOUR_API_KEY")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	status, err := client.Status.Get(context.Background())
	if err != nil {
		log.Fatalf("Failed to get status: %v", err)
	}

	fmt.Printf("API Status: %+v\n", status)
}
```

### Constructor options

The `NewClient` function accepts your `apiKey` and a variable number of functional options.

| Name             | Type                | Description                                                                  |
| ---------------- | ------------------- | ---------------------------------------------------------------------------- |
| `apiKey`         | `string`            | Your Deps API key. This is a required parameter.                             |
| `WithBaseURL`    | `func(string) Option` | Overrides the base URL of the Deps API. Defaults to `https://api.depscian.tech/v2`. |
| `WithTimeout`    | `func(time.Duration) Option` | Sets a custom timeout for the HTTP client. Defaults to `30s`.                |
| `WithHTTPClient` | `func(*http.Client) Option` | Allows you to use a completely custom `http.Client`.                         |

### API Modules

The client provides access to the API through a set of services:

- `client.Admins`
- `client.Families`
- `client.Fractions`
- `client.Ghetto`
- `client.Leadership`
- `client.Map`
- `client.Online`
- `client.Player`
- `client.Sobes`
- `client.Status`

**Example**: Get the online players on a server.

```go
package main

import (
	"context"
	"fmt"
	"log"

	depscian "deps-go"
)

func main() {
	client, err := depscian.NewClient("YOUR_API_KEY")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	var serverId int = 1
	online, err := client.Online.Get(context.Background(), serverId)
	if err != nil {
		log.Fatalf("Failed to get online list: %v", err)
	}

	fmt.Printf("Online Players: %+v\n", online)
}
```

### Error Handling

All API methods return a standard Go `error` as the second return value. You can check for `nil` to determine if the request was successful.

```go
package main

import (
	"context"
	"fmt"
	"log"

	depscian "deps-go"
)

func main() {
	client, err := depscian.NewClient("YOUR_API_KEY")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	var serverId int = 1
	var nickname string = "NonExistentPlayer"
	player, err := client.Player.Find(context.Background(), serverId, nickname)
	if err != nil {
		// The generated client returns a generic error.
		// You can inspect the response body for more details.
		log.Printf("API Error: %v", err)
		if player != nil && player.Body != nil {
			log.Printf("Error Body: %s", string(player.Body))
		}
		return
	}

	fmt.Printf("Player: %+v\n", player)
}
```

## LICENSE

[MIT](LICENSE)