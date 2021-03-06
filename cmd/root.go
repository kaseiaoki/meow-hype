/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/kaseiaoki/meow/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	snooze   string
	note     string
	after    string
	interval string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = newRootCmd()

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mw",
		Short: "meow! this is notifer",
		Long: `# meow
		meow is desktop toast notice tool.
		# usage
		### 1 default 
		"mw --note <Note to be displayed in the notification> --after <Interval between notifications (sec)> "
		
		Simple desktop notification.
		### 2 with command
		"mw <any command> --note <Note to be displayed in the notificatio> --after <Interval between notifications of running(sec)>"
		  
		Desktop notification after command execution is complete.
		## options
		### --minute bool
		Set interval in minutes
		### --hour bool
		Set interval in hour
		### --snooze string
		Set snooze(WIP)
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&note, "note", "meow!", "note")
	rootCmd.PersistentFlags().StringVar(&after, "after", "1s", "after(second)")
	rootCmd.PersistentFlags().StringVar(&interval, "interval", "10s", "interval(second)")
	rootCmd.PersistentFlags().StringVar(&snooze, "snooze", "0s", "snooze")
	rootCmd.PersistentFlags().StringVar(&config.CfgFile, "config", "", "config file (default is $HOME/.meow.toml)")
}

func initConfig() {
	if config.CfgFile != "" {
		viper.SetConfigFile(config.CfgFile)
	} else {
		conf, err := os.UserConfigDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		viper.AddConfigPath(conf)
		viper.SetConfigName(".meow")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&config.ToastConf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
