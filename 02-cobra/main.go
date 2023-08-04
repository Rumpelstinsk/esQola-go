package main

import (
	cmd "github.com/Rumpelstinsk/esQola-go/02-cobra/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cmd.InitBeersCmd())
	rootCmd.AddCommand(cmd.InitStoresCmd())
	rootCmd.Execute()
}
