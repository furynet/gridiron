# Distribution

The distribution module allows you to manage your [Staking Rewards](../concepts/general-concepts.md#staking-rewards).

## Available Subcommands

| Name                                                                                    | Description                                                                                                                                           |
| --------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| [commission](#grid-query-distribution-commission)                                       | Query distribution validator commission                                                                                                               |
| [community-pool](#grid-query-distribution-community-pool)                               | Query the amount of coins in the community pool                                                                                                       |
| [params](#grid-query-distribution-params)                                               | Query distribution params                                                                                                                             |
| [rewards](#grid-query-distribution-rewards)                                             | Query all distribution delegator rewards or rewards from a particular validator                                                                       |
| [slashes](#grid-query-distribution-slashes)                                             | Query distribution validator slashes.                                                                                                                 |
| [validator-outstanding-rewards](#grid-query-distribution-validator-outstanding-rewards) | Query distribution outstanding (un-withdrawn) rewards for a validator and all their delegations                                                       |
| [fund-community-pool](#grid-tx-distribution-fund-community-pool)                        | Funds the community pool with the specified amount                                                                                                    |
| [set-withdraw-addr](#grid-tx-distribution-set-withdraw-addr)                            | Set the withdraw address for rewards associated with a delegator address                                                                              |
| [withdraw-all-rewards](#grid-tx-distribution-withdraw-all-rewards)                      | Withdraw all rewards for a single delegator                                                                                                           |
| [withdraw-rewards](#grid-tx-distribution-withdraw-rewards)                              | Withdraw rewards from a given delegation address,and optionally withdraw validator commission if the delegation address given is a validator operator |

## grid query distribution commission

Query validator commission rewards from delegators to that validator.

```bash
grid query distribution commission [validator] [flags]
```

## grid query distribution community-pool

Query all coins in the community pool which is under Governance control.

```bash
grid query distribution community-pool [flags]
```

## grid query distribution params

Query distribution params.

```bash
 grid query distribution params [flags]
```

## grid query distribution rewards

Query all rewards earned by a delegator, optionally restrict to rewards from a single validator.

```bash
grid query distribution rewards [delegator-addr] [validator-addr] [flags]
```

## grid query distribution slashes

Query all slashes of a validator for a given block range.

```bash
grid query distribution slashes [validator] [start-height] [end-height] [flags]
```

## grid query distribution validator-outstanding-rewards

Query distribution outstanding (un-withdrawn) rewards for a validator and all their delegations.

```bash
grid query distribution validator-outstanding-rewards [validator] [flags]
```

## grid tx distribution fund-community-pool

Funds the community pool with the specified amount.

```bash
grid tx distribution fund-community-pool [amount] [flags]
```

## grid tx distribution set-withdraw-addr

Set the withdraw address for rewards associated with a delegator address.

```bash
grid tx distribution set-withdraw-addr [withdraw-addr] [flags]
```

## grid tx distribution withdraw-all-rewards

Withdraw all rewards for a single delegator.

```bash
grid tx distribution withdraw-all-rewards [flags]
```

## grid tx distribution withdraw-rewards

Withdraw rewards from a given delegation address, and optionally withdraw validator commission if the delegation address given is a validator operator.

```bash
grid tx distribution withdraw-rewards [validator-addr] [flags]
```
