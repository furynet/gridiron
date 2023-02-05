# Oracle

Oracle module allows you to manage the feed on GRID Hub

## Available Commands

| Name                              | Description                                                                          |
| --------------------------------- | ------------------------------------------------------------------------------------ |
| [create](#grid-tx-oracle-create)  | Create a new feed, the feed will be in "paused" state                                |
| [start](#grid-tx-oracle-start)    | Start a feed in "paused" state                                                       |
| [pause](#grid-tx-oracle-pause)    | Pause a feed in "running" state                                                      |
| [edit](#grid-tx-oracle-edit)      | Modify the feed information and update service invocation parameters by feed creator |
| [feed](#grid-query-oracle-feed)   | Query the feed definition                                                            |
| [feeds](#grid-query-oracle-feeds) | Query a group of feed definition                                                     |
| [value](#grid-query-oracle-value) | Query the feed result                                                                |

## grid tx oracle create

This command is used to create a new feed, the feed will be in "paused" state.

```bash
grid tx oracle create [flags]
```

**Flags:**

| Name, shorthand   | Type     | Required | Default | Description                                                                                                                    |
| ----------------- | -------- | -------- | ------- | ------------------------------------------------------------------------------------------------------------------------------ |
| --feed-name       | string   | Yes      |         | The unique identifier of the feed.                                                                                             |
| --description     | string   |          |         | The description of the feed.                                                                                                   |
| --latest-history  | uint64   | Yes      |         | The maximum Number of the latest history values to be saved for the Feed, range [1, 100].                                      |
| --service-name    | string   | Yes      |         | The name of the service to be invoked by the feed.                                                                             |
| --input           | string   | Yes      |         | The input argument (JSON) used to invoke the service.                                                                          |
| --providers       | []string | Yes      |         | The list of service provider addresses.                                                                                        |
| --service-fee-cap | string   | Yes      |         | Only providers charging a fee lower than the cap will be invoked.                                                              |
| --timeout         | int64    |          |         | The maximum number of blocks to wait for a response since a request is sent, beyond which the request will be ignored.         |
| --frequency       | uint64   |          |         | The invocation frequency of sending repeated requests.                                                                         |
| --threshold       | uint16   |          | 1       | The minimum number of responses needed for aggregation, range [1, Length(providers)].                                          |
| --aggregate-func  | string   | Yes      |         | The name of predefined function for processing the service responses, e.g. avg, max, min etc.                                  |
| --value-json-path | string   | Yes      |         | The field name or path of Service response result used to retrieve the value property of aggregate-func from response results. |

### Create a new feed

```bash
grid tx oracle create \
    --feed-name="test-feed" \
    --latest-history=10 \
    --service-name="test-service" \
    --input=<request-data> \
    --providers=<provide1_address>,<provider2_address> \
    --service-fee-cap=1grid \
    --timeout=2 \
    --frequency=10 \
    --total=10 \
    --threshold=1 \
    --aggregate-func="avg" \
    --value-json-path="high" \
    --chain-id=gridiron \
    --from=node0 \
    --fees=0.3grid
```

## grid tx oracle start

This command is used to start a feed in "paused" state

```bash
grid tx oracle start [feed-name] [flags]
```

### Start a "paused" feed

```bash
grid tx oracle start test-feed --chain-id=gridiron --from=node0 --fees=0.3grid
```

## grid tx oracle pause

This command is used to pause a feed in "running" state

```bash
grid tx oracle pause [feed-name] [flags]
```

### Pause a "running" feed

```bash
grid tx oracle pause test-feed --chain-id=gridiron --from=node0 --fees=0.3grid
```

## grid tx oracle edit

This command is used to edit an existing feed on GRID Hub.

```bash
grid tx oracle edit [feed-name] [flags]
```

**Flags:**

| Name, shorthand   | Type     | Required | Default | Description                                                                                                            |
| ----------------- | -------- | -------- | ------- | ---------------------------------------------------------------------------------------------------------------------- |
| --feed-name       | string   | Yes      |         | The unique identifier of the feed.                                                                                     |
| --description     | string   |          |         | The description of the feed.                                                                                           |
| --latest-history  | uint64   | Yes      |         | The maximum Number of the latest history values to be saved for the Feed, range [1, 100].                              |
| --providers       | []string | Yes      |         | The list of service provider addresses.                                                                                |
| --service-fee-cap | string   | Yes      |         | Only providers charging a fee lower than the cap will be invoked.                                                      |
| --timeout         | int64    |          |         | The maximum number of blocks to wait for a response since a request is sent, beyond which the request will be ignored. |
| --frequency       | uint64   |          |         | The invocation frequency of sending repeated requests.                                                                 |
| --threshold       | uint16   |          | 1       | The minimum number of responses needed for aggregation, range [1, Length(providers)].                                  |

### Edit an existed feed

```bash
grid tx oracle edit test-feed --chain-id=gridiron --from=node0 --fees=0.3grid --latest-history=5
```

## grid query oracle feed

This command is used to query a feed

```bash
grid query oracle feed [feed-name] [flags]
```

### Query an existed feed

```bash
grid query oracle feed test-feed
```

## grid query oracle feeds

This command is used to query a group of feed

```bash
grid query oracle feeds [flags]
```

**Flags:**

| Name, shorthand | Type   | Required | Default | Description                                  |
| --------------- | ------ | -------- | ------- | -------------------------------------------- |
| --state         | string |          |         | the state of the feed, e.g. paused, running. |

### Query a group of feed

```bash
grid query oracle feeds --state=running
```

## grid query oracle value

This command is used to query the result of a specified feed

```bash
grid query oracle value test-feed
```
