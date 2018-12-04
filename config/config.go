package config

import (
	"going_rest/checker"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func LoadConfig(cmd *cobra.Command) {
	checker.CheckErr(viper.BindPFlags(cmd.Flags()))

	// from the environment
	viper.SetEnvPrefix("GOREST")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.BindEnv("consul_url")
	viper.BindEnv("consul_path")

	viper.AutomaticEnv()

	configFile := viper.GetString("config")

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		// from a config file
		/*viper.SetConfigName("config")
		viper.AddConfigPath("./")*/

		//from consul
		consulUrl := viper.GetString("consul_url")
		consulPath := viper.GetString("consul_path")

		if consulUrl == "" {
			log.Fatal("CONSUL_URL missing")
		}
		if consulPath == "" {
			log.Fatal("CONSUL_PATH missing")
		}

		viper.AddRemoteProvider("consul", consulUrl, consulPath)
		viper.SetConfigType("yml")
		checker.CheckErr(viper.ReadRemoteConfig())

	}

	// NOTE: this will require that you have config file somewhere in the paths specified. It can be reading from JSON, TOML, YAML, HCL, and Java properties files.
	//checker.CheckErr(viper.ReadInConfig())

	LoadDB()
	LoadApp()

}
