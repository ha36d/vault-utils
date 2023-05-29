package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"

	"context"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy vault kv data",
	Long:  `Copy vault kv data`,
	Run: func(cmd *cobra.Command, args []string) {

		srcaddr := viper.GetString("srcaddr")
		dstaddr := viper.GetString("dstaddr")
		srctoken := viper.GetString("srctoken")
		dsttoken := viper.GetString("dsttoken")
		srcengine := viper.GetString("srcengine")
		dstengine := viper.GetString("dstengine")
		verbose = viper.GetBool("verbose")

		source := VaultClient(srcaddr, srctoken)
		destination := VaultClient(dstaddr, dsttoken)

		ctx := context.Background()

		resp, err := source.System.MountsListSecretsEngines(ctx)
		if err != nil {
			log.Fatal(err)
		}

		for engine, property := range resp.Data {
			engineType := property.(map[string]interface{})
			if engineType["type"] == "kv" && contains(strings.Split(srcengine, ","), strings.TrimSuffix(engine, "/")) {
				loopTree(source, destination, ctx, engine, dstengine, "/")
			}
		}

		log.Println("Job Finished!")

	},
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func loopTree(source *vault.Client, destination *vault.Client, ctx context.Context, engine string, dstengine string, path string) {

	res := strings.HasSuffix(path, "/")

	if res {
		keys, err := source.List(ctx, fmt.Sprintf("%s/metadata/%s", engine, path))
		if err != nil {
			log.Fatal(err)
		}
		for _, key := range Keys(keys.Data) {
			res := strings.HasSuffix(key, "/")
			if res {
				loopTree(source, destination, ctx, engine, dstengine, fmt.Sprintf("%s%s", path, key))
			} else {
				subkeys, err := source.Secrets.KvV2Read(ctx, fmt.Sprintf("%s%s", path, key), vault.WithMountPath(engine))
				if err != nil {
					log.Fatal(err)
				}
				for key, value := range subkeys.Data.Data {
					if verbose {
						log.Println("writing secret:", fmt.Sprintf("%s%s%s", engine, path, key))
					}

					_, err = destination.Secrets.KvV2Write(ctx, fmt.Sprintf("%s%s", path, key), schema.KvV2WriteRequest{
						Data: map[string]any{
							key: value,
						},
					}, vault.WithMountPath(dstengine))
					if err != nil {
						log.Fatal(err)
					}
					if verbose {
						log.Println("secret key was written successfully")
					}
				}
			}
		}
	}

}

func VaultClient(address string, token string) *vault.Client {
	// prepare a client with the given base address
	client, err := vault.New(
		vault.WithAddress(address),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	// authenticate with a root token (insecure)
	if err := client.SetToken(token); err != nil {
		log.Fatal(err)
	}

	return client
}

func Keys(m map[string]interface{}) (keys []string) {
	s := make([]string, 0)
	for _, v := range m["keys"].([]interface{}) {
		s = append(s, fmt.Sprint(v))
	}
	return s
}
func init() {

	copyCmd.Flags().StringP("sourceaddress", "s", "", "Source vault address to read from")
	viper.BindPFlag("sourceaddress", copyCmd.Flags().Lookup("sourceaddress"))
	copyCmd.Flags().StringP("dstaddr", "d", "", "Destination vault address to write to")
	viper.BindPFlag("dstaddr", copyCmd.Flags().Lookup("dstaddr"))
	copyCmd.Flags().StringP("stctoken", "t", "", "Source vault token to write to")
	viper.BindPFlag("stctoken", copyCmd.Flags().Lookup("stctoken"))
	copyCmd.Flags().StringP("dsttoken", "k", "", "Destination vault token to write to")
	viper.BindPFlag("dsttoken", copyCmd.Flags().Lookup("dsttoken"))
	copyCmd.Flags().StringP("srcengine", "e", "", "Source vault engine to write to")
	viper.BindPFlag("srcengine", copyCmd.Flags().Lookup("srcengine"))
	copyCmd.Flags().StringP("dstengine", "f", "", "Destination vault engine to write to")
	viper.BindPFlag("dstengine", copyCmd.Flags().Lookup("dstengine"))

	rootCmd.AddCommand(copyCmd)
}
