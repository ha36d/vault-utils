package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"context"

	"github.com/ha36d/vault-utils/pkg/targz"
	"github.com/ha36d/vault-utils/pkg/utils"
	"github.com/ha36d/vault-utils/pkg/vaultutility"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// restoreCmd represents the copy command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore tar.gz to kv data",
	Long:  `Restore tar.gz to kv data`,
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

		myFile, err := os.Open(backup)
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		defer myFile.Close()

		targz.Untar(fmt.Sprintf("%s/%s", osPath, srcengine), myFile)
		err = os.RemoveAll(fmt.Sprintf("%s/%s", osPath, srcengine))
		if err != nil {
			log.Fatal(err)
		}

		for engine, property := range resp.Data {
			engineType := property.(map[string]interface{})
			if engineType["type"] == "kv" && utils.Contains(strings.Split(srcengine, ","), strings.TrimSuffix(engine, "/")) {
				vaultutility.LoopTree(source, ctx, engine, "/", saveSecretToFile)
			}
		}

		log.Println("Job Finished!")

	},
}

func saveSecretToKv(ctx context.Context, engine string, path string, secret string, subkeys map[string]interface{}) {

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Printf("File Name: %s\n", info.Name())
		return nil
	})

	//  content, err := ioutil.ReadFile("./config.json")
	//
	// var payload Data
	// err = json.Unmarshal(content, &payload)

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
	if err := os.MkdirAll(fmt.Sprintf("%s/%s%s", osPath, engine, path), 0700); err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(fmt.Sprintf("%s/%s%s%s.json", osPath, engine, path, secret), body, 0400)
	if err != nil {
		log.Println(err)
	}
}

func init() {

	restoreCmd.Flags().StringP("restore", "f", "", "Backup file path")
	viper.BindPFlag("restore", restoreCmd.Flags().Lookup("restore"))

	rootCmd.AddCommand(backupCmd)
}
