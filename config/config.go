package config

import (
	"going_rest/checker"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func LoadConfig(cmd *cobra.Command) {
	checker.CheckErr(viper.BindPFlags(cmd.Flags()))

	// from the environment
	viper.SetEnvPrefix("GOREST")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	configFile := viper.GetString("config")

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		// from a config file
		viper.SetConfigName("config")
		viper.AddConfigPath("./")
		//viper.AddConfigPath("$HOME/desktop/golang/src/petting-the-cobra")
	}

	// NOTE: this will require that you have config file somewhere in the paths specified. It can be reading from JSON, TOML, YAML, HCL, and Java properties files.
	checker.CheckErr(viper.ReadInConfig())

	LoadDB()

}
