package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"context"

	"github.com/ha36d/vault-utils/pkg/targz"
	"github.com/ha36d/vault-utils/pkg/utils"
	"github.com/ha36d/vault-utils/pkg/vaultutility"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// backupCmd represents the copy command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup kv data to tar.gz",
	Long:  `Backup kv data to tar.gz`,
	Run: func(cmd *cobra.Command, args []string) {

		srcaddr := viper.GetString("addr")
		srctoken := viper.GetString("token")
		srcengine := viper.GetString("engine")
		backup := viper.GetString("backup")

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

		osPath, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}

		myFile, err := os.Create(backup)
		if err != nil {
			panic(err)
		}

		if err := os.MkdirAll(fmt.Sprintf("%s/%s", osPath, "vault-backup"), 0700); err != nil {
			log.Fatal(err)
		}
		defer os.RemoveAll(fmt.Sprintf("%s/%s", osPath, "vault-backup"))

		for engine, property := range resp.Data {
			engineType := property.(map[string]interface{})

			if engineType["type"] == "kv" && (srcengine == "" || utils.Contains(strings.Split(srcengine, ","), strings.TrimSuffix(engine, "/"))) {
				vaultutility.LoopTree(source, ctx, engine, "/", saveSecretToFile)
			}
		}

		targz.Tar(fmt.Sprintf("%s/%s", osPath, "vault-backup"), myFile)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Job Finished!")

	},
}

func saveSecretToFile(ctx context.Context, engine string, path string, secret string, subkeys map[string]interface{}) {

	verbose = viper.GetBool("verbose")

	content := make(map[string]any)

	for key, value := range subkeys {
		content[key] = value
	}
	osPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	body, err := json.Marshal(content)
	if err != nil {
		log.Println(err)
	}
	if err := os.MkdirAll(fmt.Sprintf("%s/%s/%s%s", osPath, "vault-backup", engine, path), 0700); err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(fmt.Sprintf("%s/%s/%s%s%s.json", osPath, "vault-backup", engine, path, secret), body, 0400)
	if err != nil {
		log.Println(err)
	}
}

func init() {

	backupCmd.Flags().StringP("backup", "f", "", "Backup file path")
	viper.BindPFlag("backup", backupCmd.Flags().Lookup("backup"))

	rootCmd.AddCommand(backupCmd)
}
