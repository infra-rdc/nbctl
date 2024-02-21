package netbox

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/olekukonko/tablewriter"
)

func listPlatformsPayload(netboxHost, token string, httpScheme string) ([][]string, error) {
	c := newNetboxClient(netboxHost, token, httpScheme)
	params := dcim.NewDcimPlatformsListParams()
	resp, err := c.Dcim.DcimPlatformsList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("Cannot get platforms list: %s", err)
	}

	var data [][]string

	for _, platform := range resp.Payload.Results {
		data = append(
			data, []string{
				*platform.Name,
				*&platform.Description})
	}

	return data, nil
}

func PrintPlatformsList(netboxHost, token string, httpScheme string, jsonOpt bool, rawOpt bool, platformName string) error {
	data, err := listPlatformsPayload(netboxHost, token, httpScheme)
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

		if platformName == "" {
			for _, v := range data {
				table.Append(v)
			}

		} else {
			for _, v := range data {
				for _, x := range v {
					if x == platformName {
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
