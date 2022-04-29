# Experimental Cosomos baesd multi-chain structure
Built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

Regitration chain

## Get started

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

For my example, you will need two run 4 commands:

`starport chain serve -c bob.yml -r -f -v` - This will be the "identity" chain that holds the registration/delegation storage

`starport chain serve -c ashley.yml -r -f -v` - This is the other chain that sends identity information

```
starport relayer configure -a \
--source-rpc "http://0.0.0.0:26657" \
--source-faucet "http://0.0.0.0:4500" \
--source-port "identity" \
--source-version "identity-1" \
--source-gasprice "0.0000025stake" \
--source-prefix "cosmos" \
--target-rpc "http://0.0.0.0:26659" \
--target-faucet "http://0.0.0.0:4501" \
--target-port "identity" \
--target-version "identity-1" \
--target-gasprice "0.0000025stake" \
--target-prefix "cosmos"
```

`starport relayer connect`

If you need to restart a chain, you will need to remove the relayer configuration file located at `~/.starport/relayer/config.yml` and then run the relayer configure command again.

### Configure

Your blockchain in development can be configured with `config.yml`, I've created config yamls for two chains -> `ashley.yml` and `bob.yml`. These two chains are conifgured with ports that allow them to communicate with each other. To learn more, see the [Starport docs](https://docs.starport.network).

### Identity Functions

#### Create Registry (handle)

```
registryd tx identity send-ibc-registration identity channel-0 "firstHandle" --from taco --chain-id ashley --home ~/.ashley
```

In the `send-ibc-registration` command, the args are [port] [channel] [handle] --from <account on ashley chain> --chain-id <ashley> --home <location of .ashley folder>

This command will create a new dsnpId for the handle and setup an initial delegation entry with the address from the "from" user

#### Update Delegation

```
registryd tx identity send-ibc-permission-update identity channel-0 "1" “cosmos1qyw382nm5659frpg72vnmxydc4jjp9ne8y8njt” 1 --from taco --chain-id ashley --home ~/.ashley
```

In the `send-ibc-permission-update` command, the args are [port] [channel] [dsnpId] [address to update/add] [role] --from <account on ashley chain> --chain-id <ashley> --home <location of .ashley folder>

This command will retrieve the existing delegations for the given dsnpId, then if the from user is authorized to change delegations, the address passed in will either be added to the delgations for the dsnpId or the role of it's existing delegation will be updated.
### Launch

To launch your blockchain live on multiple nodes, use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

### Web Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Starport front-end development](https://github.com/tendermint/vue).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.starport.network/amparks100/registry@latest! | sudo bash
```
`amparks100/registry` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/cosmosnetwork)
