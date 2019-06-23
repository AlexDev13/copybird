package common

import (
	"github.com/copybird/copybird/operator"
	"github.com/spf13/cobra"
	"log"
	//"log"
	//"github.com/spf13/cobra"
)

type App struct {
	cmmRoot        *cobra.Command
	cmdBackup      *cobra.Command
	cmdOperator    *cobra.Command
	vars           map[string]interface{}
}

func NewApp() *App {
	return &App{
		vars:          make(map[string]interface{}),
	}
}

func (a *App) Run() error {
	var rootCmd = &cobra.Command{Use: "copybird"}
	a.cmdBackup = &cobra.Command{
		Use:   "backup",
		Short: "Start new backup",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run:   cmdCallback(a.DoBackup),
	}
	a.cmdOperator = &cobra.Command{
		Use:   "operator",
		Short: "Start Kubernetes operator",
		Run: func(cmd *cobra.Command, args []string) {
			operator.Run()
		},
	}
	rootCmd.AddCommand(a.cmdBackup)
	a.Setup()
	return rootCmd.Execute()
}

func (a *App) DoBackup() error {
	return nil
}

func cmdCallback(f func() error) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		err := f()
		if err != nil {
			log.Printf("cmd err: %s", err)
		}
	}
}
