
# Vault-utils

Yet Another Utility for Vault. Actions supported:
    
- Copy from one Vault to another
- Backup from one Vault to tar.gz
- Restore from tar.gz to Vault


## Documentation

### Copy
This can used for copy the kv data from one Vault to another:

```
Usage:
  vault-utils copy [flags]

Flags:
  -d, --dstaddr string     Destination vault address to write to
  -f, --dstengine string   Destination vault kv engines to write to
  -k, --dsttoken string    Destination vault token to write to
  -h, --help               help for copy

Global Flags:
  -s, --addr string     Vault address
  -c, --config string   config file (default is $HOME/.vault-utils.yaml)
  -e, --engine string   Vault kv comma separated engine names
  -t, --token string    Vault token
  -v, --verbose         verbosity
```

Example:

```
./vault-utils copy --addr "http://127.0.0.1:8200" --token "someToken" --engine "kv,secret" --dstaddr "http://127.0.0.1:8202" --dsttoken "someToken" --dstengine "secret" -v
2023/05/29 20:07:51 No config file found! Counting on flags!
Using config:
--- Configuration ---
	token = someToken
	dsttoken = someToken
	engine = kv,secret
	dstengine = secret
	verbose = %!s(bool=true)
	addr = http://127.0.0.1:8200
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

### Backup
This can used for copy the kv data from one Vault to another:

```
Usage:
  vault-utils backup [flags]

Flags:
  -f, --backup string   Backup file path
  -h, --help            help for backup

Global Flags:
  -s, --addr string     Vault address
  -c, --config string   config file (default is $HOME/.vault-utils.yaml)
  -e, --engine string   Vault kv comma separated engine names
  -t, --token string    Vault token
  -v, --verbose         verbosity
```

Example:

```
./vault-utils backup --addr "http://127.0.0.1:8200" --token "someToken" -f some.tar.gz -v
2023/05/29 20:07:51 No config file found! Counting on flags!
2023/05/29 20:07:51 Job Finished!
```

Instead of flags, it is possible to use the config file:
```
./vault-utils backup -c config.yaml
```

### Restore
This can used for copy the kv data from one Vault to another:

```
Usage:
  vault-utils restore [flags]

Flags:
  -f, --backup string   Backup file path
  -h, --help            help for restore

Global Flags:
  -s, --addr string     Vault address
  -c, --config string   config file (default is $HOME/.vault-utils.yaml)
  -e, --engine string   Vault kv comma separated engine names
  -t, --token string    Vault token
  -v, --verbose         verbosity
```

Example:

```
./vault-utils restore --addr "http://127.0.0.1:8200" --token "someToken" -f some.tar.gz -v
2023/05/29 20:07:51 No config file found! Counting on flags!
2023/05/29 20:07:51 Job Finished!
```

Instead of flags, it is possible to use the config file:
```
./vault-utils restore -c config.yaml
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