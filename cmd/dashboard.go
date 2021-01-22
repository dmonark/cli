package cmd

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var dashboardListCmd = &cobra.Command{
	Use:    "dashboard",
	Short:  "Open dashboard in browser",
	PreRun: validateAuth,
	RunE: func(cmd *cobra.Command, args []string) error {
		url := "http://dashboard.razorpay.com/"
		var err error

		switch runtime.GOOS {
		case "linux":
			err = exec.Command("xdg-open", url).Start()
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		case "darwin":
			err = exec.Command("open", url).Start()
		default:
			err = fmt.Errorf("unsupported platform")
		}

		return err
	},
}
