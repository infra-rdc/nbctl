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
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Main dcim options vars
	siteNameOpt string
	// dcimCmd represents the dcim command
	dcimCmd = &cobra.Command{
		Use:   "dcim",
		Short: "Interact with dcim.",
		Long:  `Interact with dcim.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("dcim called")
		},
	}
)

func init() {
	rootCmd.AddCommand(dcimCmd)
}
