
# User

## Fetch user data

Endpoint : `GET /fetch-user-data/:address`

```go
address := "0x5D9aE648d74dF50746951Cbb2FB296c836A7A0e9"
user, err := client.User.Get(address)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| address         | string      | Address of the user to be fetched  |

**Response:**

```json
{
    "message": {
        "Nonce": 6
    },
    "statusCode": 200
}
```

-------------------------------------------------------------------------------------------------------

### Fetch user balance

Endpoint : `GET /fetch-user-balance/:address`

```go
address := "0x5D9aE648d74dF50746951Cbb2FB296c836A7A0e9"
user, err := client.User.GetUserBalance(address)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| address          | string      | Address of the user to be fetched |

**Response:**

```json
{
    "message": {
        "UserBalances": [
            {
                "CurrencyId": 4,
                "Amount": "734630000",
                "ContractAddress": "0xE9573B8A0AF951431bcBD194E8cc3AeE654Cd723",
                "Abbreviation": "USDC"
            }
        ]
    },
    "statusCode": 200
}
```

-------------------------------------------------------------------------------------------------------

### Fetch user currency balance

Endpoint : `GET /fetch-user-balance/:address/:tokenAddress`

```go
address := "0x5D9aE648d74dF50746951Cbb2FB296c836A7A0e9"
token_address := "0xEe146Fac7b2fce5FdBE31C36d89cF92f6b006F80"
user, err := client.User.GetUserCurrencyBalance(address)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| address          | string      | Address of the user to be fetched |
| tokenAddress          | string      | Address of the token to be fetched |

**Response:**

```json
{
    "message": {
        "UserBalances": {
            "CurrencyId": 4,
            "Amount": "734630000",
            "ContractAddress": "0xE9573B8A0AF951431bcBD194E8cc3AeE654Cd723",
            "Abbreviation": "USDC"
        }
    },
    "statusCode": 200
}
```

-------------------------------------------------------------------------------------------------------

### Fetch user proofs

Endpoint : `GET /fetch-user-balance/:address/:tokenAddress`

```go
address := "0x5D9aE648d74dF50746951Cbb2FB296c836A7A0e9"
token_address := "0xEe146Fac7b2fce5FdBE31C36d89cF92f6b006F80"
user, err := client.User.GetUserProof(address)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| address          | string      | Address of the user to be fetched |
| tokenAddress          | string      | Address of the token to be fetched |

**Response:**

```json
{
    "message": {
        "AccountIndex": 0,
        "AccountProof": [
            "0x5306fccd46eeea13a7b7ba988a94514277db096872995cf086250194e5d8367a",
            "0xe41866f18bc2b9502fccc217a5572755dab1f9ca439846b75c3f91f15dd9c28b",
            "0x1989cb201b9b6ba3b3a98aeaac29638c95b11fd3e2e77aa7712a8d2095bef2a4",
            "0xe97e5ec1c1266505544143852b8fd3045e1add92ddaadfb34ee38be20f2c2e7e"
        ],
        "AccountProofHelper": [
            1,
            1,
            1
        ],
        "BalancesRoot": "0x89d7e77cdef0473233ac43d43df3c7b93c6bfbbf40bdd9c007ed747383946d99",
        "BalancesProof": [
            "0xa86d54e9aab41ae5e520ff0062ff1b4cbd0b2192bb01080a058bb170d84e6457",
            "0x4b4efd86a2cec7174648fca755d3b9caf672051f139e1b37846d357f29e0d889",
            "0x2a3c055e5aad1f95e094e401d23a52dd4975291cc3ecbaef3a11c98dfdef94b8"
        ],
        "Balance": "753820000",
        "BalancesIndex": 0
    },
    "statusCode": 200
}
```

-------------------------------------------------------------------------------------------------------
