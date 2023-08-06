package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/hugplus/go-walker/cmd/module"
	"github.com/hugplus/go-walker/cmd/start"
	"github.com/hugplus/go-walker/cmd/version"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:          "go-walker",
		Short:        "go-walker",
		Long:         `go-walker`,
		SilenceUsage: true,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				tip()
				return errors.New("requires at least one arg")
			}
			return nil
		},
		PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
		Run: func(cmd *cobra.Command, args []string) {
			tip()
		},
	}
)

func tip() {
	usageStr := `欢迎使用 go-walker 查看命令：go-walker --help`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(start.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
	rootCmd.AddCommand(module.StartCmd)
	//rootCmd.AddCommand(migrate.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
