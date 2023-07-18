package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"context"

	"github.com/ha36d/vault-utils/pkg/targz"
	"github.com/ha36d/vault-utils/pkg/utils"
	"github.com/ha36d/vault-utils/pkg/vaultutility"
	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// restoreCmd represents the copy command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore tar.gz to kv data",
	Long:  `Restore tar.gz to kv data`,
	Run: func(cmd *cobra.Command, args []string) {

		dstaddr := viper.GetString("addr")
		dsttoken := viper.GetString("token")
		dstengine := viper.GetString("engine")
		restore := viper.GetString("restore")

		verbose = viper.GetBool("verbose")

		destination, err := vaultutility.VaultClient(dstaddr, dsttoken)

		if err != nil {
			log.Fatal(err)
		}

		ctx := context.Background()

		resp, err := destination.System.MountsListSecretsEngines(ctx)
		if err != nil {
			log.Fatal(err)
		}

		osPath, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}

		myFile, err := os.Open(restore)
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		defer myFile.Close()

		if err = targz.Untar(fmt.Sprintf("%s/%s", osPath, "vault-restore"), myFile); err != nil {
			log.Fatalf("unable to untar the file: %v", err)
		}

		engineVersion := ""
		versions := []string{"1", "2"}

		for engine, property := range resp.Data {
			engineProperty := property.(map[string]interface{})
			if engineProperty["options"] != nil {
				engineOption := engineProperty["options"].(map[string]interface{})
				if engineOption["version"] != nil {
					engineVersion = engineOption["version"].(string)
				}
			}
			if utils.Contains(versions, engineVersion) && engineProperty["type"] == "kv" && (dstengine == "" || utils.Contains(strings.Split(dstengine, ","), strings.TrimSuffix(engine, "/"))) {
				saveSecretToKv(destination, ctx, engine)
			}
		}

		defer os.RemoveAll(fmt.Sprintf("%s/%s", osPath, "vault-restore"))
		log.Println("Job Finished!")

	},
}

func saveSecretToKv(destination *vault.Client, ctx context.Context, engine string) {

	verbose = viper.GetBool("verbose")

	osPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	filepath.Walk(fmt.Sprintf("%s/%s/%s", osPath, "vault-restore", engine), func(path string, info os.FileInfo, err error) error {

		var payload map[string]interface{}

		if err != nil {
			log.Fatalf(err.Error())
		}
		if !info.IsDir() {
			fileName := info.Name()
			content, err := ioutil.ReadFile(path)
			err = json.Unmarshal(content, &payload)
			if err != nil {
				log.Println(err)
			}

			_, err = destination.Secrets.KvV2Write(ctx, fmt.Sprintf("%s%s", strings.TrimPrefix(filepath.Dir(path), fmt.Sprintf("%s/%s/%s", osPath, "vault-restore", strings.TrimSuffix(engine, "/"))), strings.TrimSuffix(fileName, filepath.Ext(fileName))), schema.KvV2WriteRequest{
				Data: payload,
			}, vault.WithMountPath(engine))
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})

}

func init() {

	restoreCmd.Flags().StringP("restore", "f", "", "Backup file path")
	viper.BindPFlag("restore", restoreCmd.Flags().Lookup("restore"))

	rootCmd.AddCommand(restoreCmd)
}
