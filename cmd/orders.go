package cmd

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
)

var orderCmd = &cobra.Command{
	Use:   "orders",
	Short: "list all the orders",
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, element := range os.Environ() {
			variable := strings.Split(element, "=")
			fmt.Println(variable[0], "=>", variable[1])
		}
		fmt.Println("**")
		v, b := syscall.Getenv("rzp_key")
		fmt.Println(b)
		fmt.Println(v)
		fmt.Println("Printing all the orders")
		return nil
	},
}

var orderCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create order",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Creating orders")
		return nil
	},
}

func init() {
	orderCmd.AddCommand(orderCreateCmd)
}
