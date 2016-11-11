// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/kildevaeld/apprun"
	_ "github.com/kildevaeld/apprun/node"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func printError(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "%s\n", e)
		os.Exit(1)
	}
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "app_run",
	Short: "",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		var config apprun.Config

		if err := viper.Unmarshal(&config); err != nil {
			printError(err)
		}

		config.Branch = viper.GetString("branch")
		config.Remote = viper.GetString("remote")
		config.Kind = viper.GetString("kind")
		config.Command = viper.GetString("command")
		config.Workspace = viper.GetString("workspace")
		config.Environ = os.Environ()
		if len(args) > 0 {
			config.Command = args[0]
		}

		if len(args) > 1 {
			config.Args = args[1:]
		}

		app, e := apprun.New(config)
		if e != nil {
			printError(e)
		}

		if err := NewProcess("Initializing application ... ", app.Init); err != nil {
			printError(err)
		}

		if viper.GetBool("update") {

			if err := NewProcess("Updating application ... ", app.Update); err != nil {
				printError(err)
			}
		}

		if err := NewProcess("Building application ... ", app.Build); err != nil {
			printError(err)
		}

		printError(app.Run())

	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.app_run.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	flags := RootCmd.Flags()

	flags.StringP("workspace", "w", "/workspace", "Path to workspace")
	flags.StringP("remote", "r", "", "Remote")
	flags.StringP("branch", "b", "", "")
	flags.StringP("kind", "k", "node", "")
	flags.BoolP("update", "u", false, "")

	viper.BindPFlags(flags)

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".app_run") // name of config file (without extension)
	viper.AddConfigPath("$HOME")    // adding home directory as first search path
	viper.SetEnvPrefix("APPRUN")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

}
