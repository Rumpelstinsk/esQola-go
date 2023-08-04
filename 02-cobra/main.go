package main

import (
	beer_cmd "github.com/Rumpelstinsk/esQola-go/02-cobra/cmd"
	stores_cmd "github.com/Rumpelstinsk/esQola-go/02-cobra/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(beer_cmd.InitBeersCmd())
	rootCmd.AddCommand(stores_cmd.InitStoresCmd())
	rootCmd.Execute()
}
