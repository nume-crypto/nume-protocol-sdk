
# User

## Create user

Endpoint : `POST /initiate-registration`

```go
data := map[string]interface{}{
    "blsG2PublicKey": ["2a7a2b702cd1282c6d615f70b083a76cc290fcaacc9f0c4534a775b23323bbe3","12eae9c0a36ed48fbfce93a81237aa8a8f980514008726e1ba633101862cb22a","15b747f77a80cd966258a2c3769d1ea37f9782263d8e8a74c1c6860a002bcd23","03577a899a7b8e33af9b8d9af42d50bf85f5334c7e78499da6f47bcef78f4d80"],
    "blsG1PublicKey": ["1a4488372cdfad3051d1f8169ad5a242343258e37b78dcd2b0b3bfb973a5e955","1d702935f31018f7ff285b175c334baf0adb3119bdcb43902ea4a2b84b674b36"],
    "cmkId": "ABCDEFGH-XXXX-YYYY-ZZZZ-123456789011",
    "encryptedPrivateKey": "AQICAHhim9vQpNTqtsff2a9psId3gQYMUurhcnwGij0iO6SzMgG1qOPpoK202H0/I7apzbzhAAAAojCBnwYJKoZIhvcNAQcGoIGRMIGOAgEAMIGIBgkqhkiG9w0BBwEwHgYJYIZIAWUDBAEuMBEEDF+s/FUDX+6WvGgKLgIBEIBbQTfJ6W3fczjyH5aAdtMXhgSC5S8Zx/tY4lugsgCHBEOefCisfZhiDkrRsw2f408NeMra7BQBbKfBvTLs7aq46JBSnNBjQCSWdwdAYH1yWPB4z+9IY5h0vGc3bA=="
}
user, err := client.User.Create(data)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| blsG1PublicKey          | [2]string      | User's G1 public key components                        |
| blsG2PublicKey        | [4]string      | User's G2 public key components                   |
| cmkId | string | CMK ID used to encrypt user's private key |
| encryptedPrivateKey      | string      | User's encrypted private key              |

**Response:**

```json
{
    "message": {
        "Id" : "1",
        "HashedPublicKey": "190ce0ac817bf41f26c665f414e0fc1c955864aff93b28132b2f73ed65522a29",
        "DepositSignature" : "8a22e585ffa1a11d00c7a3e31cef97e510a0bde2bb7c46b87e0dab636cead4fc4959d75eefcb008964d4be3d111feddb74648b0158e5e1303cde4ad85f080c1d1c",
        "RegistrationSignature" : "bc14939486e34436c42ab7b7e2c54de1d17f580ed73352b1d8257804d3a144aa417e778d2881922f35e367a0b9359ea51f9a3b7727386cfed43bade741db17811b"
    },
    "statusCode": 200
}
```

-------------------------------------------------------------------------------------------------------

### Fetch user data

Endpoint : `GET /fetch-user-data/:hashedPublicKey`

```go
user_hashed_publickey := "190ce0ac817bf41f26c665f414e0fc1c955864aff93b28132b2f73ed65522a29"
user, err := client.User.Get(user_hashed_publickey)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| user_hashed_publickey         | string      | Hashed publickey of the user to be fetched  |

**Response:**

```json
{
    "message": {
        "EncryptedPrivateKey": "AQICAHhim9vQpNTqtsff2a9psId3gQYMUurhcnwGij0iO6SzMgG1qOPpoK202H0/I7apzbzhAAAAojCBnwYJKoZIhvcNAQcGoIGRMIGOAgEAMIGIBgkqhkiG9w0BBwEwHgYJYIZIAWUDBAEuMBEEDF+s/FUDX+6WvGgKLgIBEIBbQTfJ6W3fczjyH5aAdtMXhgSC5S8Zx/tY4lugsgCHBEOefCisfZhiDkrRsw2f408NeMra7BQBbKfBvTLs7aq46JBSnNBjQCSWdwdAYH1yWPB4z+9IY5h0vGc3bA==",
        "CmkId": "4bf5bfc2-75cd-46b2-b657-44aed46bf502",
        "BlsG1PublicKey": [
            "1a4488372cdfad3051d1f8169ad5a242343258e37b78dcd2b0b3bfb973a5e955",
            "1d702935f31018f7ff285b175c334baf0adb3119bdcb43902ea4a2b84b674b36"
        ],
        "BlsG2PublicKey": [
            "2a7a2b702cd1282c6d615f70b083a76cc290fcaacc9f0c4534a775b23323bbe3",
            "12eae9c0a36ed48fbfce93a81237aa8a8f980514008726e1ba633101862cb22a",
            "15b747f77a80cd966258a2c3769d1ea37f9782263d8e8a74c1c6860a002bcd23",
            "03577a899a7b8e33af9b8d9af42d50bf85f5334c7e78499da6f47bcef78f4d80"
        ],
        "Nonce": 12,
        "RegistrationSignature": "bc14939486e34436c42ab7b7e2c54de1d17f580ed73352b1d8257804d3a144aa417e778d2881922f35e367a0b9359ea51f9a3b7727386cfed43bade741db17811b",
        "DepositSignature": "8a22e585ffa1a11d00c7a3e31cef97e510a0bde2bb7c46b87e0dab636cead4fc4959d75eefcb008964d4be3d111feddb74648b0158e5e1303cde4ad85f080c1d1c"
    },
    "statusCode": 200
}
```

-------------------------------------------------------------------------------------------------------

### Fetch user balance

Endpoint : `GET /fetch-user-balance/:hashedPublicKey`

```go
user_hashed_publickey := "190ce0ac817bf41f26c665f414e0fc1c955864aff93b28132b2f73ed65522a29"
user, err := client.User.GetUserBalance(user_hashed_publickey)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| user_hashed_publickey          | string      | Hashed publickey of the user to be fetched |

**Response:**

```json
{
    "message": {
        "UserBalances": [
            {
                "CurrencyId": 1,
                "Amount": "0",
                "ContractAddress": "0x799c6832d187243f3367902079A72fb3Fd61cdF7",
                "Abbreviation": "ETH",
                "TokenOrder": 5
            },
            {
                "CurrencyId": 2,
                "Amount": "10739999999999998976",
                "ContractAddress": "0xEe146Fac7b2fce5FdBE31C36d89cF92f6b006F80",
                "Abbreviation": "DAI",
                "TokenOrder": 1
            },
            {
                "CurrencyId": 4,
                "Amount": "8550000",
                "ContractAddress": "0xE9573B8A0AF951431bcBD194E8cc3AeE654Cd723",
                "Abbreviation": "USDC",
                "TokenOrder": 2
            },
            {
                "CurrencyId": 3,
                "Amount": "16328",
                "ContractAddress": "0x0b6D9aB4c80889b65A61050470CBC5523d8Ce48D",
                "Abbreviation": "WBTC",
                "TokenOrder": 4
            }
        ]
    },
    "statusCode": 200
}
```

-------------------------------------------------------------------------------------------------------
