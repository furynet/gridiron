# Slashing

Slashing module can unjail validator previously jailed for downtime

## Available Commands

| Name                                                | Description                                     |
| --------------------------------------------------- | ----------------------------------------------- |
| [unjail](#grid-tx-slashing-unjail)                  | Unjail validator previously jailed for downtime |
| [params](#grid-query-slashing-params)               | Query the current slashing parameters           |
| [signing-info](#grid-query-slashing-signing-info)   | Query a validator's signing information         |
| [signing-infos](#grid-query-slashing-signing-infos) | Query signing information of all validators     |

## grid tx slashing unjail

Unjail validator previously jailed for downtime.

```bash
grid tx slashing unjail [flags]
```

## grid query slashing params

Query the current slashing parameters.

```bash
grid query slashing params  [flags]
```

## grid query slashing signing-info

Query a validator's signing information.

```bash
grid query slashing signing-info [validator-conspub] [flags]
```

## grid query slashing signing-infos

Query signing information of all validators.

```bash
grid query slashing signing-infos [flags]
```
