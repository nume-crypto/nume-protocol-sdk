# Nume Protocol Go Client

Golang bindings for interacting with the Nume Protocol API
This is primarily meant for DAPPs who wish to perform interactions with the  Nume Protocol API programatically

## Usage

```go
import (
    "github.com/nume-crypto/nume-protocol-sdk"
)
client := nume.NewClient("NUME-API-KEY", nume.BASE_URL)
```

Note: All methods below return a `map[string]interface{}`.

### Documentation

- [User](docs/user.md)
- [Transaction](docs/transaction.md)
- [Protocol](docs/protocol.md)
