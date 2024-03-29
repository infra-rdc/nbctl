package netbox

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/olekukonko/tablewriter"
)

func listDevicesPayload(netboxHost, token string, httpScheme string) ([][]string, error) {
	c := newNetboxClient(netboxHost, token, httpScheme)
	params := dcim.NewDcimDevicesListParams()
	resp, err := c.Dcim.DcimDevicesList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("Cannot get devices list: %s", err)
	}

	var data [][]string

	for _, device := range resp.Payload.Results {
		data = append(
			data, []string{
				*device.Name,
				*device.DeviceType.Model,
				*device.Tenant.Name,
				*&device.Serial,
				*device.Location.Name,
				*device.Site.Name,
				*device.Rack.Name,
				*device.Status.Value})
	}

	return data, nil
}

func PrintDevicesList(netboxHost, token string, httpScheme string, jsonOpt bool, rawOpt bool, deviceName string) error {
	data, err := listDevicesPayload(netboxHost, token, httpScheme)
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
		table.SetHeader([]string{"Name", "Type", "Tenant", "Serial", "Location", "Site", "Rack", "Status"})
		table.SetBorder(true)

		if deviceName == "" {
			for _, v := range data {
				table.Append(v)
			}

		} else {
			for _, v := range data {
				for _, x := range v {
					if x == deviceName {
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
