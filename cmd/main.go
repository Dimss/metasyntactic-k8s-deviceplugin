package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

var (
	Version = "v0.0.1"
	Build   string
	rootCmd = &cobra.Command{
		Use:   "foobar",
		Short: "Foobar - device plugin for baz devices",
	}
	versionCmd = &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "show foobar version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("🐾 version: %s build: %s \n", Version, Build)
		},
	}
	startCmd = &cobra.Command{
		Use:     "start",
		Aliases: []string{"s"},
		Short:   "start foobar device plugin",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("starting foobar device plugin")
		},
	}
)

func initConfig() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("FOOBAR_DEVICE_PLUGIN")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	setupLogging()

}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(versionCmd)
}

func setupLogging() {

	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		DisableColors: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := strings.TrimSuffix(filepath.Base(frame.File), filepath.Ext(frame.File))
			line := strconv.Itoa(frame.Line)
			return "", fmt.Sprintf("%s:%s", fileName, line)
		},
	})
	// Logs are always goes to STDOUT
	log.SetOutput(os.Stdout)
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
