package start

import (
	"fmt"

	"github.com/hugplus/go-walker/core"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	//AppRouters = make([]func(), 0)
	configYml string
	StartCmd  = &cobra.Command{
		Use:     "start",
		Short:   "Get Application config info",
		Example: "go-walker start -c resources/config.dev.yml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "resources/config.dev.yml", "Start server with provided configuration file")
}

func run() {
	if configYml == "" {
		panic("找不到配置文件")
	}
	v := viper.New()
	v.SetConfigFile(configYml)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&core.Cfg); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&core.Cfg); err != nil {
		fmt.Println(err)
	}
	core.Init()
	core.Run(&AppRouters)
}
