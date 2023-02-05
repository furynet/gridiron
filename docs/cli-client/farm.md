# Farm

Farm module allows you to easily create farm activities on gridiron.

## Available Commands

| Name                              | Description                                           |
| --------------------------------- | ----------------------------------------------------- |
| [create](#grid-tx-farm-create)    | Create a new farm pool                                |
| [adjust](#grid-tx-farm-adjust)    | Adjust farm pool parameters                           |
| [destroy](#grid-tx-farm-destroy)  | Destroy the farm pool and get back the invested bonus |
| [stake](#grid-tx-farm-stake)      | Deposit liquidity token                               |
| [harvest](#grid-tx-farm-harvest)  | Get back the bonus for participating in the farm pool |
| [farmer](#grid-query-farm-farmer) | Query farmer information                              |
| [pool](#grid-query-farm-pool)     | Query the current status of a farm pool               |
| [pools](#grid-query-farm-pools)   | Query farm pool information by page                   |
| [params](#grid-query-farm-params) | Query the management parameters of the farm module    |

## grid tx farm create

Create a new farm pool and pay the handling fee and bonus.

```bash
grid tx farm create <Farm Pool Name> [flags]
```

**Flags:**

| Name, shorthand    | Required | Default | Description                                              |
| ------------------ | -------- | ------- | -------------------------------------------------------- |
| --lp-token-denom   | true     |         | The liquidity token accepted by farm pool                |
| --reward-per-block | true     |         | The reward per block,ex: 1grid,1atom                     |
| --total-reward     | true     |         | The Total reward for the farm pool                       |
| --description      | false    | ""      | The simple description of a farm pool                    |
| --start-height     | true     |         | The start height the farm pool                           |
| --editable         | false    | false   | Is it possible to adjust the parameters of the farm pool |

### grid tx farm adjust

Adjust the parameters of the pool before the farm pool ends, such as `reward-per-block`, `total-reward`.

```bash
grid tx farm adjust <Farm Pool Name> [flags]
```

**Flags:**

| Name, shorthand     | Required                                  | Default | Description                          |
| ------------------- | ----------------------------------------- | ------- | ------------------------------------ |
| --additional-reward | And `--reward-per-block` must choose one  | ""      | Bonuses added to the farm pool       |
| --reward-per-block  | And `--additional-reward` must choose one | ""      | The reward per block,ex: 1grid,1atom |

## grid tx farm destroy

Destroy the farm pool and get back the invested bonus.The rewards earned by the user farm ends at this moment, requiring the user to manually retrieve the income and the liquidity of the deposit.

```bash
grid tx farm destroy <Farm Pool Name> [flags]
```

### grid tx farm stake

The farmer participates in farm activities by staking the liquidity tokens specified by the pool. The rewards obtained by participating in the activities are related to the number of staking tokens and farm pool parameters.

```bash
grid tx farm stake <Farm Pool Name> <lp token> [flags]
```

### grid tx farm harvest

The farmer withdraws his rewards back.

```bash
grid tx farm harvest <Farm Pool Name>
```

### grid query farm farmer

Query farmer's information, including unclaimed rewards, mortgage liquidity, etc.

```bash
grid query farm farmer <Farmer Address> --pool-name <Farm Pool Name>
```

**Flags:**

| Name, shorthand | Required | Default | Description        |
| --------------- | -------- | ------- | ------------------ |
| --pool-name     | false    | ""      | the farm pool name |

### grid query farm pool

Query related information of a farm pool by name

```bash
grid query farm pool <Farm Pool Name>
```

### grid query farm pools

Paging query farm pool

```bash
grid query farm pools
```

### grid query farm params

Paging query farm pool

```bash
grid query farm params
```
