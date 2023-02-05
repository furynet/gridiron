# NFT

`NFT` provides the ability to digitize assets. Through this module, each off-chain asset will be modeled as a unique on-chain nft.

## Available Commands

| Name                                          | Description                                                                                         |
| --------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| [issue](#grid-tx-nft-issue)                   | Specify the nft Denom (nft classification) and metadata JSON Schema to issue nft.                   |
| [transfer-denom](#grid-tx-nft-transfer-denom) | The owner of the NFT classification can transfer the ownership of the NFT classification to others. |
| [mint](#grid-tx-nft-mint)                     | Additional issuance (create) of specific nft of this type can be made.                              |
| [edit](#grid-tx-nft-edit)                     | The metadata of the specified nft can be updated.                                                   |
| [transfer](#grid-tx-nft-transfer)             | Transfer designated nft.                                                                            |
| [burn](#grid-tx-nft-burn)                     | Destroy the created nft.                                                                            |
| [supply](#grid-query-nft-supply)              | Query the total amount of nft according to Denom; accept the optional owner parameter.              |
| [owner](#grid-query-nft-owner)                | Query all nft owned by an account; you can specify the Denom parameter.                             |
| [collection](#grid-query-nft-collection)      | Query all nft according to Denom.                                                                   |
| [denom](#grid-query-nft-denom)                | Query nft denom information based on Denom.                                                         |
| [denoms](#grid-query-nft-denoms)              | Query the total amount of nft according to Denom; accept the optional owner parameter.              |
| [token](#grid-query-nft-token)                | Query specific nft based on Denom and ID.                                                           |

## grid tx nft issue

Specify the nft Denom (nft classification) and metadata JSON Schema to issue nft.

```bash
grid tx nft issue [denom-id] [flags]
```

**Flags:**

| Name, shorthand     | Required | Default                                                                                                                                                                                                                     | Description |
| ------------------- | -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------- |
| --name              |          | The name of the denom                                                                                                                                                                                                       |             |
| --uri               |          | The uri of the denom                                                                                                                                                                                                        |             |
| --data              |          | Off-chain metadata for supplementation (JSON object)                                                                                                                                                                        |             |
| --schema            |          | Denom data structure definition                                                                                                                                                                                             |             |
| --symbol            |          | The symbol of the denom                                                                                                                                                                                                     |             |
| --mint-restricted   |          | This field indicates whether there are restrictions on the issuance of NFTs under this classification, true means that only Denom owners can issue NFTs under this classification, false means anyone can                   |             |
| --update-restricted |          | This field indicates whether there are restrictions on updating NFTs under this classification, true means that no one under this classification can update the NFT, false means that only the owner of this NFT can update |             |

## grid tx nft transfer-denom

The owner of the NFT classification can transfer the ownership of the NFT classification to others.

```bash
grid tx nft transfer-denom [recipient] [denom-id]
```

## grid tx nft mint

Additional issuance (create) of specific nft of this type can be made.  

```bash
grid tx nft mint [denomID] [tokenID] [flags]
```

**Flags:**

| Name, shorthand | Required | Default                     | Description |
| --------------- | -------- | --------------------------- | ----------- |
| --uri           |          | URI of off-chain token data |             |
| --recipient     |          | Receiver of the nft         |             |
| --name          |          | The name of nft             |             |

## grid tx nft edit

The metadata of the specified nft can be updated.

```bash
grid tx nft edit [denomID] [tokenID] [flags]
```

**Flags:**

| Name, shorthand | Required | Default                     | Description |
| --------------- | -------- | --------------------------- | ----------- |
| --uri           |          | URI of off-chain token data |             |
| --name          |          | The name of nft             |             |

## grid tx nft transfer

Transfer designated nft.

```bash
grid tx nft transfer [recipient] [denomID] [tokenID] [flags]
```

**Flags:**

| Name, shorthand | Required | Default                     | Description |
| --------------- | -------- | --------------------------- | ----------- |
| --uri           |          | URI of off-chain token data |             |
| --name          |          | The name of nft             |             |

## grid tx nft burn

Destroy the created nft.

```bash
grid tx nft burn [denomID] [tokenID] [flags]
```

## grid query nft

Query nft

### grid query nft supply

```bash
grid query nft supply [denomID]
grid query nft supply [denomID] --owner=<owner address>
```

### grid query nft owner

```bash
grid query nft owner [owner address]
grid query nft owner [owner address] --denom=<denomID>
```

### grid query nft collection

```bash
grid query nft collection [denomID]
```

### grid query nft denom

```bash
grid query nft denom [denomID]
```

### grid query nft denoms

```bash
grid query nft denoms
```

### grid query nft token

```bash
grid query nft token [denomID] [tokenID]
```
