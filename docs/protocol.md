
# Protocol

## Fetch currencies

Endpoint : `GET /get-currencies`

```go
currencies, err := client.Protocol.GetCurrencies()
```

**Response:**

```json
{
    "message": {
        "Currencies": [
            {
                "ID": 1,
                "CreatedAt": "2023-04-05T15:54:09.080227+05:30",
                "UpdatedAt": "2023-04-05T15:54:09.080227+05:30",
                "DeletedAt": null,
                "Name": "Ether",
                "Abbreviation": "ETH",
                "ContractAddress": "0x799c6832d187243f3367902079A72fb3Fd61cdF7",
                "Decimal": 18,
                "Native": false,
                "TokenOrder": 5
            },
            {
                "ID": 2,
                "CreatedAt": "2023-04-05T15:54:09.081032+05:30",
                "UpdatedAt": "2023-04-05T15:54:09.081032+05:30",
                "DeletedAt": null,
                "Name": "DAI",
                "Abbreviation": "DAI",
                "ContractAddress": "0xEe146Fac7b2fce5FdBE31C36d89cF92f6b006F80",
                "Decimal": 18,
                "Native": false,
                "TokenOrder": 1
            },
            {
                "ID": 3,
                "CreatedAt": "2023-04-05T15:54:09.081483+05:30",
                "UpdatedAt": "2023-04-05T15:54:09.081483+05:30",
                "DeletedAt": null,
                "Name": "WBTC",
                "Abbreviation": "WBTC",
                "ContractAddress": "0x0b6D9aB4c80889b65A61050470CBC5523d8Ce48D",
                "Decimal": 8,
                "Native": false,
                "TokenOrder": 4
            },
            {
                "ID": 4,
                "CreatedAt": "2023-04-05T15:54:09.081866+05:30",
                "UpdatedAt": "2023-04-05T15:54:09.081866+05:30",
                "DeletedAt": null,
                "Name": "USDC",
                "Abbreviation": "USDC",
                "ContractAddress": "0xE9573B8A0AF951431bcBD194E8cc3AeE654Cd723",
                "Decimal": 6,
                "Native": false,
                "TokenOrder": 2
            },
            {
                "ID": 5,
                "CreatedAt": "2023-04-05T15:54:09.082253+05:30",
                "UpdatedAt": "2023-04-05T15:54:09.082253+05:30",
                "DeletedAt": null,
                "Name": "Tether",
                "Abbreviation": "USDT",
                "ContractAddress": "0xCE47C48fDF8c9355FDbE4DacC1e1954914D65Be6",
                "Decimal": 6,
                "Native": false,
                "TokenOrder": 3
            },
            {
                "ID": 6,
                "CreatedAt": "2023-04-05T15:54:09.082924+05:30",
                "UpdatedAt": "2023-04-05T15:54:09.082924+05:30",
                "DeletedAt": null,
                "Name": "Matic",
                "Abbreviation": "MATIC",
                "ContractAddress": "0x",
                "Decimal": 18,
                "Native": true,
                "TokenOrder": 0
            }
        ]
    },
    "statusCode": 200
}
```

-------------------------------------------------------------------------------------------------------

### Monitor states

Endpoint : `GET /monitor-states`

```go
states, err := client.Protocol.MonitorStates()
```

**Response:**

```json
{
    "message": {
        "ContractAddress": "0xD805b041E5C49F5b671a771861651C717eB80DE5",
        "ContractSettlementRoot": "1b207a6b43887f3ca1ec84024b500d7a9ecdc8039a0b7ee1b955eb06b98f4b2b",
        "ContractLastFinalizedQueueIndex": "9",
        "ContractCurrentQueueIndex": "25",
        "ContractSettlementDataMD5": "0000000000000000000000000000000000000000000000000000000000000000",
        "ContractAPKG2": [
            "288e613acd44b4090806359bac8dd82315fad5e2abc8acd4cb480f7a8b6fac16",
            "0a841ec430685bceb2be91ad99b392ab077c6ced01ba3cdd40f4e47683026d4f",
            "22cd4d041418d1e16601fe966087ce27a1ed0f9b8f3fd236896a3316fb0df1aa",
            "1ed29a91c42c6b2e4edf23a361eeefe225c8e5392028facf24043f9cc2df5448"
        ],
        "ContractBlockNumber": "4",
        "BackendSettlementRoot": "1b207a6b43887f3ca1ec84024b500d7a9ecdc8039a0b7ee1b955eb06b98f4b2b",
        "BackendLastFinalizedQueueIndex": "9",
        "BackendCurrentQueueIndex": "25",
        "BackendBlockNumber": "4",
        "BackendAPKG2": [
            "288e613acd44b4090806359bac8dd82315fad5e2abc8acd4cb480f7a8b6fac16",
            "0a841ec430685bceb2be91ad99b392ab077c6ced01ba3cdd40f4e47683026d4f",
            "22cd4d041418d1e16601fe966087ce27a1ed0f9b8f3fd236896a3316fb0df1aa",
            "1ed29a91c42c6b2e4edf23a361eeefe225c8e5392028facf24043f9cc2df5448"
        ]
    },
    "statusCode": 200
}
```

-------------------------------------------------------------------------------------------------------
