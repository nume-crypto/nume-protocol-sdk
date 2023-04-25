
# Transaction

## Create transaction

Endpoint : `POST /store-eth-transaction`

```go
data := map[string]interface{}{
    "data": "0x02f8aa8202c60680808303345094e9573b8a0af951431bcbd194e8cc3aee654cd72380b844a9059cbb000000000000000000000000d28c7e81b6e8ba287307c034133eb7d72302ebf30000000000000000000000000000000000000000000000000000000001dbebf0c001a05710d8b6a53cbcfc3f47db9a7c6d97d292b6c2ceac8c0dca552b89e37483f894a06d098518f913cbfb5fd00f7219b965c29233a927ac49f069186df1c4bcae64b2"
}
tx, err := client.Transaction.Create(data)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| data          | string      | Raw transaction data (if from and to address is same then its considered as a withdrawal request) |

**Response:**

```json
{
   "message":{
      "From":"0x5D9aE648d74dF50746951Cbb2FB296c836A7A0e9",
      "To":"0xD28c7E81b6E8ba287307c034133EB7d72302eBf3",
      "TokenAddress":"0xE9573B8A0AF951431bcBD194E8cc3AeE654Cd723",
      "Amount":"31190000",
      "Nonce":6,
      "TxHash":"0xfda688fc5607a9388de0f8efa6d90c331ab0d6cf32dec3222efccc282bdf85bb",
      "Signature":"5710d8b6a53cbcfc3f47db9a7c6d97d292b6c2ceac8c0dca552b89e37483f8946d098518f913cbfb5fd00f7219b965c29233a927ac49f069186df1c4bcae64b201",
      "PublicKey":"04d6ad9daabdc226a09d6d8b2e60eba2844c0ea310eda5601888a27402c93ecd46747a502aa07bbf3386fe5cec20325eea5bfb295618f32a4a5731525a6958a4eb"
   },
   "statusCode":200
}
```

-------------------------------------------------------------------------------------------------------

### Fetch transaction data

Endpoint : `GET /transaction/:txHash`

```go
tx_hash := "0xd4b1e4d1557f25129b0024e5055d6b8629698c8dab5633a9e6f22e159c1e889e"
tx, err := client.Transaction.Get(tx_hash)
```

**Parameters:**

| Name          | Type        | Description                                 |
|---------------|-------------|---------------------------------------------|
| txHash          | string      | Transaction hash |

**Response:**

```json
{
    "message": {
        "Hash": "0xfda688fc5607a9388de0f8efa6d90c331ab0d6cf32dec3222efccc282bdf85bb",
        "From": "0x5d9ae648d74df50746951cbb2fb296c836a7a0e9",
        "To": "0xd28c7e81b6e8ba287307c034133eb7d72302ebf3",
        "Currency": "0xE9573B8A0AF951431bcBD194E8cc3AeE654Cd723",
        "CurrencyId": 4,
        "Amount": "31190000",
        "Status": "SUCCESS",
        "Nonce": 6,
        "Type": "transfer",
        "CreatedAt": "2023-04-25T13:58:16.161575+05:30"
    },
    "statusCode": 200
}
```

-------------------------------------------------------------------------------------------------------
