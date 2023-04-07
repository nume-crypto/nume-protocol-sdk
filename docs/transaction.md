
# Transaction

## Create transaction

Endpoint : `POST /store-transaction`

```go
data := map[string]interface{}{
    "from": "18e6b07f0d3731e5d4cfa3acba2cac922ab085ac33966f09045525ec56f00447",
    "to": "447bF33F7c7C925eb7674bCF590AeD4Aa57e656b",
    "amount": "10",
    "nonce": 4,
    "tokenAddress": "0xCE47C48fDF8c9355FDbE4DacC1e1954914D65Be6",
    "type": "withdrawal",
    "signature": "020f5b0ba41993f001ae5e058a421267fff39ccbbedd2167f45f3bc984b284285e"
}
tx, err := client.Transaction.Create(data)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| from          | string      | User's hashed public key                       |
| to        | string      | To address  ( transfer ->  user2's hashed public key / withdrawal -> wallet address)                |
| amount | string | Token value  to be processed |
| nonce | int | Nonce provided by the user |
| tokenAddress      | string      | Token address chosen by the user for the transaction              |
| type      | string      | Transaction type ( transfer or withdrawal)             |
| signature      | string      | User's signature of the payload              |

**Response:**

```json
{
    "message": "success",
    "statusCode": 200
}
```

-------------------------------------------------------------------------------------------------------
