# Service

Service module allows you to define, bind, invoke services on the GRID Hub. [Read more about iService](../features/service.md).

## 可用命令

| Name                                                    | Description                                                        |
| ------------------------------------------------------- | ------------------------------------------------------------------ |
| [define](#grid-tx-service-define)                       | Define a new service                                               |
| [definition](#grid-query-service-definition)            | Query a service definition                                         |
| [bind](#grid-tx-service-bind)                           | Bind a service                                                     |
| [binding](#grid-query-service-binding)                  | Query a service binding                                            |
| [bindings](#grid-query-service-bindings)                | Query all bindings of a service definition                         |
| [set-withdraw-addr](#grid-tx-service-set-withdraw-addr) | Set a withdrawal address for a provider                            |
| [withdraw-addr](#grid-query-service-withdraw-addr)      | Query the withdrawal address of a provider                         |
| [update-binding](#grid-tx-service-update-binding)       | Update an existing service binding                                 |
| [disable](#grid-tx-service-disable)                     | Disable an available service binding                               |
| [enable](#grid-tx-service-enable)                       | Enable an unavailable service binding                              |
| [refund-deposit](#grid-tx-service-refund-deposit)       | Refund all deposit from a service binding                          |
| [call](#grid-tx-service-call)                           | Initiate a service call                                            |
| [request](#grid-query-service-request)                  | Query a request by the request ID                                  |
| [requests](#grid-query-service-requests)                | Query active requests by the service binding or request context ID |
| [respond](#grid-tx-service-respond)                     | Respond to a service request                                       |
| [response](#grid-query-service-response)                | Query a response by the request ID                                 |
| [responses](#grid-query-service-responses)              | Query active responses by the request context ID and batch counter |
| [request-context](#grid-query-service-request-context)  | Query a request context                                            |
| [update](#grid-tx-service-update)                       | Update a request context                                           |
| [pause](#grid-tx-service-pause)                         | Pause a running request context                                    |
| [start](#grid-tx-service-start)                         | Start a paused request context                                     |
| [kill](#grid-tx-service-kill)                           | Terminate a request context                                        |
| [fees](#grid-query-service-fees)                        | Query the earned fees of a provider                                |
| [withdraw-fees](#grid-tx-service-withdraw-fees)         | Withdraw the earned fees of a provider                             |
| [schema](#grid-query-service-schema)                    | Query the system schema by the schema name                         |
| [params](#grid-query-service-params)                    | Query values set as service parameters.                            |

## grid tx service define

Define a new service.

```bash
grid tx service define [flags]
```

**Flags:**

| Name, shorthand      | Default | Description                                       | Required |
| -------------------- | ------- | ------------------------------------------------- | -------- |
| --name               |         | Service name                                      | Yes      |
| --description        |         | Service description                               |          |
| --author-description |         | Service author description                        |          |
| --tags               |         | Service tags                                      |          |
| --schemas            |         | Content or file path of service interface schemas | Yes      |

### define a service

```bash
grid tx service define \
    --name=<service name> \
    --description=<service description> \
    --author-description=<author description>
    --tags=tag1,tag2 \
    --schemas=<schemas content or path/to/schemas.json> \
    --chain-id=gridiron \
    --from=<key-name> \
    --fees=0.3grid
```

### Schemas content example

```json
{
    "input": {
        "$schema": "http://json-schema.org/draft-04/schema#",
        "title": "BioIdentify service input body",
        "description": "BioIdentify service input body specification",
        "type": "object",
        "properties": {
            "id": {
                "description": "id",
                "type": "string"
            },
            "name": {
                "description": "name",
                "type": "string"
            },
            "data": {
                "description": "data",
                "type": "string"
            }
        },
        "required": [
            "id",
            "data"
        ]
    },
    "output": {
        "$schema": "http://json-schema.org/draft-04/schema#",
        "title": "BioIdentify service output body",
        "description": "BioIdentify service output body specification",
        "type": "object",
        "properties": {
            "data": {
                "description": "result data",
                "type": "string"
            }
        },
        "required": [
            "data"
        ]
    }
}
```

## grid query service definition

Query a service definition.

```bash
grid query service definition [service-name] [flags]
```

### Query a service definition

Query the detailed info of the service definition with the specified service name.

```bash
grid query service definition <service name>
```

## grid tx service bind

Bind a service.

```bash
grid tx service bind [flags]
```

**Flags:**

| Name, shorthand | Default | Description                                                                                                                   | Required |
| --------------- | ------- | ----------------------------------------------------------------------------------------------------------------------------- | -------- |
| --service-name  |         | Service name                                                                                                                  | Yes      |
| --deposit       |         | Deposit of the binding                                                                                                        | Yes      |
| --pricing       |         | Pricing content or file path, which is an instance of [Gridiron Service Pricing JSON Schema](../features/service-pricing.json) | Yes      |
| --qos           |         | Minimum response time                                                                                                         | Yes      |
| --options       |         | Non-functional requirements options                                                                                           | Yes      |
| --provider      |         | Provider address, default to the owner                                                                                        |          |

### Bind an existing service definition

The deposit needs to satisfy the minimum deposit requirement, which is the maximal one between `price` * `MinDepositMultiple` and `MinDeposit` (`MinDepositMultiple` and `MinDeposit` are the system parameters, which can be modified through the governance).

```bash
grid tx service bind \
    --service-name=<service name> \
    --deposit=10000grid \
    --pricing=<pricing content or path/to/pricing.json> \
    --qos=50 \
    --options=<non-functional requirements options content or path/to/options.json> \
    --chain-id=gridiron \
    --from=<key-name> \
    --fees=0.3grid
```

### Pricing content example

```json
{
    "price": "1grid"
}
```

## grid query service binding

Query a service binding.

```bash
grid query service binding <service name> <provider>
```

## grid query service bindings

Query all bindings of a service definition.

```bash
grid query service bindings [service-name] [flags]
```

### Query service binding list

```bash
grid query service bindings <service name> <owner address>
```

## grid tx service update-binding

Update a service binding.

```bash
grid tx service update-binding [service-name] [provider-address] [flags]
```

**Flags:**

| Name, shorthand | Default | Description                                                                                                                                       | Required |
| --------------- | ------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| --deposit       |         | Deposit added for the binding, not updated if empty                                                                                               |          |
| --pricing       |         | Pricing content or file path, which is an instance of [Gridiron Service Pricing JSON Schema](../features/service-pricing.md), not updated if empty |          |
| --qos           |         | Minimum response time, not updated if set to 0                                                                                                    |          |
| --options       |         | Non-functional requirements options                                                                                                               |          |

### Update an existing service binding

The following example updates the service binding with the additional 10 GRID deposit

```bash
grid tx service update-binding <service-name> <provider-address> \
    --deposit=10grid \
    --options=<non-functional requirements options content or path/to/options.json> \
    --pricing='{"price":"1grid"}' \
    --qos=50 \
    --chain-id=<chain-id> \
    --from=<key name> \
    --fees=0.3grid
```

## grid tx service set-withdraw-addr

Set a withdrawal address for a provider.

```bash
grid tx service set-withdraw-addr [withdrawal-address] [flags]
```

## grid query service withdraw-addr

Query the withdrawal address of a provider.

```bash
grid query service withdraw-addr [provider] [flags]
```

## grid tx service disable

Disable an available service binding.

```bash
grid tx service disable [service-name] [provider-address] [flags]
```

## grid tx service enable

Enable an unavailable service binding.

```bash
grid tx service enable [service-name] [provider-address] [flags]
```

**Flags:**

| Name, shorthand | Default | Description                            | Required |
| --------------- | ------- | -------------------------------------- | -------- |
| --deposit       |         | deposit added for enabling the binding |          |

### Enable an unavailable service binding

The following example enables an unavailable service binding with the additional 10 GRID deposit.

```bash
grid tx service enable <service name> <provider-address> --chain-id=gridiron --from=<key-name> --fees=0.3grid --deposit=10grid
```

## grid tx service refund-deposit

Refund all deposits from a service binding.

```bash
grid tx service refund-deposit [service-name] [provider-address] [flags]
```

### Refund all deposits from an unavailable service binding

Before refunding, you should [disable](#grid-tx-service-disable) the service binding first.

```bash
grid tx service refund-deposit <service name> <provider-address> --chain-id=gridiron --from=<key-name> --fees=0.3grid
```

## grid tx service call

Initiate a service call.

```bash
grid tx service call [flags]
```

**Flags:**

| Name, shorthand   | Default | Description                                                                                                            | Required |
| ----------------- | ------- | ---------------------------------------------------------------------------------------------------------------------- | -------- |
| --service-name    |         | Service name                                                                                                           | Yes      |
| --providers       |         | Provider list to request                                                                                               | Yes      |
| --service-fee-cap |         | Maximum service fee to pay for a single request                                                                        | Yes      |
| --data            |         | Content or file path of the request input, which is an Input JSON Schema instance                                      | Yes      |
| --timeout         |         | Request timeout                                                                                                        | Yes      |
| --repeated        | false   | Indicate if the reqeust is repetitive (Temporarily disabled in gridiron-v1.0.0, will be activated after a few versions) |          |
| --frequency       |         | Request frequency when repeated, default to `timeout`                                                                  |          |
| --total           |         | Request count when repeated, -1 means unlimited                                                                        |          |

### Initiate a service invocation request

```bash
grid tx service call \
    --service-name=<service name> \
    --providers=<provider list> \
    --service-fee-cap=1grid \
    --data=<request input or path/to/input.json> \
    --timeout=100 \
    --repeated \
    --frequency=150 \
    --total=100 \
    --chain-id=gridiron \
    --from=<key name> \
    --fees=0.3grid
```

### Input example

```json
{
    "header": {
        ...
    },
    "body": {
        "id": "1",
        "name": "irisnet",
        "data": "facedata"
    }
}
```

## grid query service request

Query a request by the request ID.

```bash
grid query service request [request-id] [flags]
```

### Query a service request

```bash
grid query service request <request-id>
```

:::tip
You can retrieve the `request-id` in [Query request_id through rpc interface](#Query request_id through rpc interface) or [grid query service requests](#grid query service requests).
:::

### Query request_id through rpc interface

Query `block_results` according to `block height` through `rpc interface`, find `new_batch_request_provider` in `end_block_events`, decode the result with base64 to get `request_id`.

```bash
curl -X POST -d '{"jsonrpc":"2.0","id":1,"method":"block_results","params":["10604"]}' http://localhost:26657
```

## grid query service requests

Query active requests by the service binding or request context ID.

```bash
grid query service requests [service-name] [provider] | [request-context-id] [batch-counter] [flags]
```

### Query active requests of a service binding

```bash
grid query service requests <service name> <provider>
```

### Query service requests by the request context ID and batch counter

```bash
grid query service requests <request-context-id> <batch-counter>
```

## grid tx service respond

Respond to a service request.

```bash
grid tx service respond [flags]
```

**Flags:**

| Name, shorthand | Default | Description                                                                                                                                | Required |
| --------------- | ------- | ------------------------------------------------------------------------------------------------------------------------------------------ | -------- |
| --request-id    |         | ID of the request to respond to                                                                                                            | Yes      |
| --result        |         | Content or file path of the response result, which is an instance of [Gridiron Service Result JSON Schema](../features/service-result.json) | Yes      |
| --data          |         | Content or file path of the response output, which is an Output JSON Schema instance                                                       |          |

### Respond to a service request

```bash
grid tx service respond \
    --request-id=<request-id> \
    --result=<response result or path/to/result.json> \
    --data=<response output or path/to/output.json>
    --chain-id=gridiron \
    --from=<key-name> \
    --fees=0.3grid
```

:::tip
You can retrieve the `request-id` in [Query request_id through rpc interface](#Query request_id through rpc interface) or [grid query service requests](#grid query service requests).
:::

### Result example

```json
{
    "code": 200,
    "message": ""
}
```

### Output example

```json
{
    "header": {
        ...
    },
    "body": {
        "data": "userdata"
    }
}
```

## grid query service response

Query a service response.

```bash
grid query service response [request-id] [flags]
```

:::tip
You can retrieve the `request-id` in [Query request_id through rpc interface](#Query request_id through rpc interface) or [grid query service requests](#grid query service requests).
:::

## grid query service responses

Query active responses by the request context ID and batch counter.

```bash
grid query service responses [request-context-id] [batch-counter] [flags]
```

### Query responses by the request context ID and batch counter

```bash
grid query service responses <request-context-id> <batch-counter>
```

## grid query service request-context

Query a request context.

```bash
grid query service request-context [request-context-id] [flags]
```

### Query a request context

```bash
grid query service request-context <request-context-id>
```

:::tip
You can retrieve the `request-context-id` in the result of [service call](#grid-tx-service-call)
:::

## grid tx service update

Update a request context.

```bash
grid tx service update [request-context-id] [flags]
```

**Flags:**

| Name, shorthand   | Default | Description                                                           | Required |
| ----------------- | ------- | --------------------------------------------------------------------- | -------- |
| --providers       |         | Provider list to request, not updated if empty                        |          |
| --service-fee-cap |         | Maximum service fee to pay for a single request, not updated if empty |          |
| --timeout         |         | Request timeout, not updated if set to 0                              |          |
| --frequency       |         | Request frequency, not updated if set to 0                            |          |
| --total           |         | Request count, not updated if set to 0                                |          |

### Update a request context

```bash
grid tx service update <request-context-id> \
    --providers=<provider list> \
    --service-fee-cap=1grid \
    --timeout=0 \
    --frequency=150 \
    --total=100 \
    --chain-id=gridiron \
    --from=<key name> \
    --fees=0.3grid
```

## grid tx service pause

Pause a running request context.

```bash
grid tx service pause [request-context-id] [flags]
```

### Pause a running request context

```bash
grid tx service pause <request-context-id>
```

## grid tx service start

Start a paused request context.

```bash
grid tx service start [request-context-id] [flags]
```

### Start a paused request context

```bash
grid tx service start <request-context-id>
```

## grid tx service kill

Terminate a request context.

```bash
grid tx service kill [request-context-id] [flags]
```

### Kill a request context

```bash
grid tx service kill <request-context-id>
```

## grid query service fees

Query the earned fees of a provider.

```bash
grid query service fees [provider] [flags]
```

## grid tx service withdraw-fees

Withdraw the earned fees of a provider.

```bash
grid tx service withdraw-fees [provider-address] [flags]
```

## grid query service schema

Query the system schema by the schema name, only pricing and result allowed.

```bash
grid query service schema [schema-name] [flags]
```

### Query the service pricing schema

```bash
grid query service schema pricing
```

### Query the response result schema

```bash
grid query service schema result
```

## grid query service params

Query values set as service parameters.

```bash
grid query service params [flags]
```
