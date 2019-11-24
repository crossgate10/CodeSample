package main

import (
	"GrabOrders/pool/repo"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	Execute()
}

var rootCmd = &cobra.Command{
	Use:   "pool",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var getOrdersCmd = &cobra.Command{
	Use:   "gen",
	Short: "",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Gen: " + strings.Join(args, " "))
		n, _ := strconv.Atoi(args[0])
		repo.GenOrders(int64(n))
	},
}

func Execute() {
	rootCmd.AddCommand(getOrdersCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
