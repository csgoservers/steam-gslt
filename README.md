# Steam `IGameServersService` client

![GitHub repo size](https://img.shields.io/github/repo-size/csgoservers/steam-gslt?logo=github&style=for-the-badge)
![GitHub issues](https://img.shields.io/github/issues/csgoservers/steam-gslt?logo=github&style=for-the-badge)

>See [license](LICENSE) if you plan to run this software. If you have any doubt [contact us](mailto:hi@csgoservers.xyz).

This is a Go implementation of the Steam `IGameServersService` interface. It is used to manage game servers that needs a *token* to work. If you want to use it, you need a [Steam key](https://steamcommunity.com/dev/apikey). You can manage your Steam game server tokens through the official [web](https://steamcommunity.com/dev/managegameservers).

### References

* [Web API overview](https://partner.steamgames.com/doc/webapi_overview)
* [Steam `IGameServersService` interface](https://partner.steamgames.com/doc/webapi/IGameServersService)

## How to use it?

If you want to use this client implementation in your own projects, you only need to execute the next command in your project root directory:

```bash
$ go get -u github.com/csgoservers/steam-gslt/pkg
```

In order to use it in your own projects you can use this code template:

```go
import "github.com/csgoservers/steam-gslt/pkg/client"

func main() {
  key := "use your own Steam API key"
  steam := client.New(key)
  accounts, err := steam.GetAccountList()
}
```

>This is an example that retrieves all accounts from your Steam API key.

#### CLI

>Command line application is a work in progress.

To use it you just need to clone this repository and execute the `make build` directive. If you want to see the full list of supported flags, then execute this command:

```bash
$ ./gslt-cli --help
```

## License

>This is not an official Steam product. csgoservers.xyz is not affiliated with Valve Corporation or Counter-Strike

See [LICENSE](LICENSE)

[![License](https://img.shields.io/badge/License-AGPLv3%202.0-brightgreen.svg?style=for-the-badge)](https://www.gnu.org/licenses/agpl-3.0.txt)
