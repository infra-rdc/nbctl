package netbox

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/olekukonko/tablewriter"
)

func listDeviceRolesPayload(netboxHost, token string, httpScheme string) ([][]string, error) {
	c := newNetboxClient(netboxHost, token, httpScheme)
	params := dcim.NewDcimDeviceRolesListParams()
	resp, err := c.Dcim.DcimDeviceRolesList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("Cannot get device roles list: %s", err)
	}

	var data [][]string

	for _, roles := range resp.Payload.Results {
		data = append(
			data, []string{
				*roles.Name,
				*&roles.Description})
	}

	return data, nil
}

func PrintDeviceRolesList(netboxHost, token string, httpScheme string, jsonOpt bool, rawOpt bool, deviceRoleName string) error {
	data, err := listDeviceRolesPayload(netboxHost, token, httpScheme)
	if err != nil {
		return err
	}

	// Print output un json format
	if jsonOpt {
		jsonData, _ := json.Marshal(data)
		fmt.Printf(string(jsonData))
	} else if rawOpt {
		// Print result in raw format
		for _, value := range data {
			fmt.Println(value)
		}
	} else {
		// Init new table
		table := tablewriter.NewWriter(os.Stdout)

		// Set table headers
		table.SetHeader([]string{"Name", "Description"})
		table.SetBorder(true)

		if deviceRoleName == "" {
			for _, v := range data {
				table.Append(v)
			}

		} else {
			for _, v := range data {
				for _, x := range v {
					if x == deviceRoleName {
						table.Append(v)
					}
				}
			}
		}
		// Print table in std output
		table.Render()
	}
	return nil
}
