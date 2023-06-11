package cmd

import (
	"fmt"
	"log"
	"strings"

	"context"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ha36d/vault-utils/pkg/utils"
	"github.com/ha36d/vault-utils/pkg/vaultutility"
)

type codeReturn func()

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy kv data from one vault to another",
	Long:  `Copy kv data from one vault to another`,
	Run: func(cmd *cobra.Command, args []string) {

		srcaddr := viper.GetString("addr")
		srctoken := viper.GetString("token")
		srcengine := viper.GetString("engine")
		verbose = viper.GetBool("verbose")

		source, err := vaultutility.VaultClient(srcaddr, srctoken)

		if err != nil {
			log.Fatal(err)
		}

		ctx := context.Background()

		resp, err := source.System.MountsListSecretsEngines(ctx)
		if err != nil {
			log.Fatal(err)
		}

		for engine, property := range resp.Data {
			engineType := property.(map[string]interface{})
			if engineType["type"] == "kv" && utils.Contains(strings.Split(srcengine, ","), strings.TrimSuffix(engine, "/")) {
				vaultutility.LoopTree(source, ctx, engine, "/", copySecret)
			}
		}

		log.Println("Job Finished!")

	},
}

func copySecret(ctx context.Context, engine string, path string, secret string, subkeys map[string]interface{}) {

	dstaddr := viper.GetString("dstaddr")
	dsttoken := viper.GetString("dsttoken")
	dstengine := viper.GetString("dstengine")
	verbose = viper.GetBool("verbose")
	destination, err := vaultutility.VaultClient(dstaddr, dsttoken)

	if err != nil {
		log.Fatal(err)
	}

	for key, value := range subkeys {

		_, err := destination.Secrets.KvV2Write(ctx, fmt.Sprintf("%s%s", path, secret), schema.KvV2WriteRequest{
			Data: map[string]any{
				key: value,
			},
		}, vault.WithMountPath(dstengine))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func init() {

	copyCmd.Flags().StringP("dstaddr", "d", "", "Destination vault address to write to")
	viper.BindPFlag("dstaddr", copyCmd.Flags().Lookup("dstaddr"))
	copyCmd.Flags().StringP("dsttoken", "k", "", "Destination vault token to write to")
	viper.BindPFlag("dsttoken", copyCmd.Flags().Lookup("dsttoken"))
	copyCmd.Flags().StringP("dstengine", "f", "", "Destination vault kv engines to write to")
	viper.BindPFlag("dstengine", copyCmd.Flags().Lookup("dstengine"))

	rootCmd.AddCommand(copyCmd)
}
