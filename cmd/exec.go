package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap/zapcore"
	"github.com/actorbuf/skeleton/config"
	"github.com/actorbuf/iota/utils"
	logger "github.com/actorbuf/skeleton/log"
	"os"
)

var (
	rootCmd = &cobra.Command{}
	cfgFile string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(preOperation)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config_local.yaml", "config file")

	//rootCmd.AddCommand(apiServerCommand) // TODO, 可以继续追加命令
}

func preOperation() {
	// 配置文件加载
	config.Cfg = config.NewDefaultConfig()
	err := config.Cfg.SetPath(cfgFile).LoadConfigFile()
	if err != nil {
		panic(err)
	}

	if config.Cfg.ServiceName == "" {
		panic("service name cannot be empty! ")
	}

	developmentMod := true
	logLevel := zapcore.DebugLevel
	if config.Cfg.Environment == "prod" {
		developmentMod = false
		logLevel = zapcore.InfoLevel
	}

	_ = logger.NewDefaultLogger(config.Cfg.ServiceName, developmentMod, logLevel)
}