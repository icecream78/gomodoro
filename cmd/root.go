/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/cheggaaa/pb/v3"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	// wTime, rTime, borderTime int = 25 * 60, 5 * 60, 1 * 60
	wTime, rTime, borderTime int = 10, 5 * 60, 5
	sCounter                 *StepsCounter
	workTimer                *Timer
	rTimer                   *Timer
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gomidoro",
	Short: "CLI app for increasing your productivity with Pomodoro method",
	Long:  ``,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		r, err := cmd.Flags().GetInt("count")
		if err != nil {
			fmt.Println("Error occuer")
			return err
		}
		sCounter = NewStepsCounter(r)
		workTimer = NewTimer(wTime, borderTime)
		rTimer = NewTimer(rTime, borderTime)

		pb.RegisterElement("wtimer", workTimer, true)
		pb.RegisterElement("rtimer", rTimer, true)
		pb.RegisterElement("steps", sCounter, true)

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		widget := NewWidget(workTimer, sCounter)
		widget.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolP("daemon", "d", false, "Run timer in background as daemon")
	rootCmd.Flags().BoolP("break", "b", false, "Run break timer")
	rootCmd.Flags().IntP("count", "c", 1, "Repeat timer [COUNT] times")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".gomidoro")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
