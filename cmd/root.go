/*
Copyright © 2024 Julien Briault <julien.briault@restosducoeur.org>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nbctl",
	Short: "Uncomplicated CLI interaction with Netbox.",
	Long:  `Uncomplicated CLI interaction with Netbox.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	jsonOpt     bool
	rawOpt      bool
	netboxToken string = os.Getenv("NETBOX_TOKEN")
	netboxHost  string = os.Getenv("NETBOX_HOST")
	httpScheme  string = os.Getenv("NETBOX_HTTP_SCHEME")
)

func init() {
	if netboxToken == "" {
		log.Fatalf("Please provide netbox API token via env var NETBOX_TOKEN")
	}

	if netboxHost == "" {
		log.Fatalf("Please provide netbox host via env var NETBOX_HOST")
	}

	if httpScheme == "" {
		log.Fatalf("Please provide netbox HTTP scheme (http or https) via env var NETBOX_HTTP_SCHEME")
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
