
# Vault-utils

Yet Another Utility for Vault. Actions supported:
    
- Copy from one Vault to another


## Documentation

This can used for copy the kv data from one Vault to another:

```
Usage:
  vault-utils copy [flags]

Flags:
  -d, --dstaddr string     Destination vault address to write to
  -f, --dstengine string   Destination vault kv engines to write to
  -k, --dsttoken string    Destination vault token to write to
  -h, --help               help for copy
  -s, --srcaddr string     Source vault address to read from
  -e, --srcengine string   Source vault kv engines to read from
  -t, --srctoken string    Source vault token to read from

Global Flags:
  -c, --config string   config file (default is $HOME/.vault-utils.yaml)
  -v, --verbose         verbosity
```

Example:

```
./vault-utils copy --srcengine "kv,secret" --dstengine "secret" --srcaddr "http://127.0.0.1:8200" --srctoken "someToken" --dstaddr "http://127.0.0.1:8202" --dsttoken "someToken" -v
2023/05/29 20:07:51 No config file found! Counting on flags!
Using config:
--- Configuration ---
	srctoken = someToken
	dsttoken = someToken
	srcengine = kv,secret
	dstengine = secret
	verbose = %!s(bool=true)
	srcaddr = http://127.0.0.1:8200
	dstaddr = http://127.0.0.1:8202
---
2023/05/29 20:07:51 writing secret: secret//ss/ss
2023/05/29 20:07:51 secret key was written successfully
2023/05/29 20:07:51 writing secret: secret//tt/tt
2023/05/29 20:07:51 secret key was written successfully
2023/05/29 20:07:51 Job Finished!
```

Instead of flags, it is possible to use the config file:
```
./vault-utils copy -c config.yaml
```
## Installation

From the [release page](https://github.com/ha36d/vault-utils/releases/), find the version suitable for your env, and run it shell.

```bash
  ./vault-utils
```
    
## Contributing

Contributions are always welcome!

See `contributing.md` for ways to get started.


## Acknowledgements

This couldn't work without

 - [Hashicorp Vault](https://github.com/hashicorp/vault)
 - [Hashicorp Vault Go Client Library](https://github.com/hashicorp/vault-client-go/)
 - [Cobra](https://github.com/spf13/cobra/)
 - [Viper](https://github.com/spf13/viper)
 - [Goreleaser](https://github.com/Goreleaser)