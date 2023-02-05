# Record

Record module allows you to manage record on GRID Hub

## Available Commands

| Name                                | Description        |
| ----------------------------------- | ------------------ |
| [create](#grid-tx-record-create)    | Create a record    |
| [record](#grid-query-record-record) | Query record by id |

## grid tx record create

Create a record

```bash
grid tx record create [digest] [digest-algo] [flags]
```

**Flags:**

| Name, shorthand | Type   | Required | Default | Description                                |
| --------------- | ------ | -------- | ------- | ------------------------------------------ |
| --uri           | string |          |         | Source uri of record, such as an ipfs link |
| --meta          | string |          |         | meta data of record                        |

## grid query record record

Query record by id

```bash
grid query record record [record-id]
```
