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
	"path"

	"github.com/icecream78/gomodoro/pomodoro"
	"github.com/icecream78/gomodoro/widget"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	template                            string = `{{ red "Work time:" }} {{bar . "[" "=" "=>" "_" "]"}} {{ string . "timer" }} {{ string . "steps" }}`
	wTime, rTime, lrTime, notify, steps int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gomodoro",
	Short: "CLI app for increasing your productivity with Pomodoro method",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bar := widget.NewBar(template)
		c := pomodoro.Config{
			WorkTime:     wTime,
			RestTime:     rTime,
			Notify:       notify,
			Steps:        steps,
			LongRestTime: lrTime,
		}
		p := pomodoro.NewPomodoroTimer(&c)
		p.Subscribe(bar)

		// run app
		p.Run()
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

	// TODO: write custom unmarshaller for values with support m,s shorthands
	rootCmd.Flags().IntVarP(&steps, "count", "c", 1, "Repeat timer [COUNT] times")
	rootCmd.Flags().IntVarP(&wTime, "step", "w", 25, "Time duration for work step in minutes")
	rootCmd.Flags().IntVarP(&rTime, "break", "b", 5, "Time duration for break step in minutes")
	rootCmd.Flags().IntVarP(&lrTime, "rest", "r", 20, "Time duration for rest step in minutes")
	rootCmd.Flags().IntVarP(&notify, "notify", "n", 20, "Remaining time in percent when need make notification")
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// TODO: generate base config on command enter
	viper.AddConfigPath(path.Join(home, ".config", "gomodoro"))
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
