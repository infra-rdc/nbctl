package netbox

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/olekukonko/tablewriter"
)

func listSitesPayload(netboxHost, token string, httpScheme string) ([][]string, error) {
	c := newNetboxClient(netboxHost, token, httpScheme)
	params := dcim.NewDcimSitesListParams()
	resp, err := c.Dcim.DcimSitesList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("Cannot get sites list: %s", err)
	}

	var data [][]string

	for _, site := range resp.Payload.Results {
		data = append(
			data, []string{
				*site.Name,
				*&site.PhysicalAddress,
				*site.TimeZone})
	}

	return data, nil
}

func PrintSitesList(netboxHost, token string, httpScheme string, jsonOpt bool, rawOpt bool, siteName string) error {
	data, err := listSitesPayload(netboxHost, token, httpScheme)
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
		table.SetHeader([]string{"Name", "Physical Address", "TimeZone"})
		table.SetBorder(true)

		if siteName == "" {
			for _, v := range data {
				table.Append(v)
			}

		} else {
			for _, v := range data {
				for _, x := range v {
					if x == siteName {
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
