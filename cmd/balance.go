package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var balanceCmd = &cobra.Command{
	Use:    "balance",
	Short:  "Display merchant balance",
	PreRun: validateAuth,
	RunE: func(cmd *cobra.Command, args []string) error {
		response, error := ExecuteRequest("https://api.razorpay.com/v1/balance", http.MethodGet, nil)
		if error != nil {
			fmt.Println(error.Error())
			os.Exit(1)
		}

		var data map[string]interface{}
		json.Unmarshal(response, &data)

		amount := data["balance"].(float64)
		color.Yellow("Your balance is " + fmt.Sprintf("%f ", amount) + data["currency"].(string))
		return nil

	},
}
