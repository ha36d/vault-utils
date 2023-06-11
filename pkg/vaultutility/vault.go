package vaultutility

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/vault-client-go"
)

func VaultClient(address string, token string) (*vault.Client, error) {
	// prepare a client with the given base address
	client, err := vault.New(
		vault.WithAddress(address),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err == nil {
		// authenticate with a root token (insecure)
		err = client.SetToken(token)
	}

	return client, err
}

func LoopTree(source *vault.Client, ctx context.Context, engine string, path string, f func(context.Context, string, string, string, map[string]interface{})) {

	res := strings.HasSuffix(path, "/")

	if res {
		keys, err := source.List(ctx, fmt.Sprintf("%s/metadata/%s", engine, path))
		if err != nil {
			log.Fatal(err)
		}
		for _, key := range Keys(keys.Data) {
			res := strings.HasSuffix(key, "/")
			if res {
				LoopTree(source, ctx, engine, fmt.Sprintf("%s%s", path, key), f)
			} else {
				subkeys, err := source.Secrets.KvV2Read(ctx, fmt.Sprintf("%s%s", path, key), vault.WithMountPath(engine))
				if err != nil {
					log.Fatal(err)
				}
				f(ctx, engine, path, key, subkeys.Data.Data)
			}
		}
	}

}

func Keys(m map[string]interface{}) (keys []string) {
	s := make([]string, 0)
	for _, v := range m["keys"].([]interface{}) {
		s = append(s, fmt.Sprint(v))
	}
	return s
}
