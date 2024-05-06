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
	"github.com/infra-rdc/nbctl/pkg/netbox"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	platformName string
	// dcimPlatformsListCmd represents the list command
	dcimPlatformsListCmd = &cobra.Command{
		Use:   "list",
		Short: "List all platforms.",
		Long:  `List all platforms.`,
		Run: func(cmd *cobra.Command, args []string) {
			err := netbox.PrintPlatformsList(netboxHost, netboxToken, httpScheme, jsonOpt, rawOpt, platformName)
			if err != nil {
				log.Error(err)
			}
		},
	}
)

func init() {
	dcimPlatformsCmd.AddCommand(dcimPlatformsListCmd)
	dcimPlatformsListCmd.Flags().BoolVarP(&jsonOpt, "json", "j", false, "Enable json output")
	dcimPlatformsListCmd.Flags().BoolVarP(&rawOpt, "raw", "r", false, "Enable raw output")
	dcimPlatformsListCmd.Flags().StringVarP(&siteNameOpt, "name", "n", "", "Select specifc platform name")
}
