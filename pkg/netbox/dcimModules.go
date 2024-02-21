package netbox

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/olekukonko/tablewriter"
)

func listModulesPayload(netboxHost, token string, httpScheme string) ([][]string, error) {
	c := newNetboxClient(netboxHost, token, httpScheme)
	params := dcim.NewDcimModulesListParams()
	resp, err := c.Dcim.DcimModulesList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("Cannot get modules list: %s", err)
	}

	var data [][]string

	for _, module := range resp.Payload.Results {
		data = append(
			data, []string{
				*module.Device.Name, *&module.Serial})
	}

	return data, nil
}

func PrintModulesList(netboxHost, token string, httpScheme string, jsonOpt bool, rawOpt bool, moduleName string) error {
	data, err := listModulesPayload(netboxHost, token, httpScheme)
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
		table.SetHeader([]string{"Name", "Serial"})
		table.SetBorder(true)

		if moduleName == "" {
			for _, v := range data {
				table.Append(v)
			}

		} else {
			for _, v := range data {
				for _, x := range v {
					if x == moduleName {
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
