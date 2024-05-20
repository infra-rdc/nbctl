package netbox

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/olekukonko/tablewriter"
)

// PlatformData holds the details of a platform to be displayed
type PlatformData struct {
	Name        string
	Description string
}

// listPlatformsPayload retrieves the platform data from NetBox
func listPlatformsPayload(netboxHost, token string, httpScheme string) ([]PlatformData, error) {
	c := newNetboxClient(netboxHost, token, httpScheme)
	params := dcim.NewDcimPlatformsListParams()
	resp, err := c.Dcim.DcimPlatformsList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot get platforms list: %w", err)
	}

	var platforms []PlatformData

	for _, platform := range resp.Payload.Results {
		platforms = append(platforms, PlatformData{
			Name:        getString(platform.Name),
			Description: getString(platform.Description),
		})
	}

	return platforms, nil
}

// PrintPlatformsList prints the list of platforms in the desired format
func PrintPlatformsList(netboxHost, token string, httpScheme string, jsonOpt bool, rawOpt bool, platformName string) error {
	platforms, err := listPlatformsPayload(netboxHost, token, httpScheme)
	if err != nil {
		return err
	}

	// Print output in JSON format
	if jsonOpt {
		return printJSON(platforms)
	}

	// Print result in raw format
	if rawOpt {
		printRaw(platforms)
		return nil
	}

	// Print result in table format
	printTable(platforms, platformName)
	return nil
}

// getString safely dereferences a string pointer
func getString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

// printJSON prints the platforms in JSON format
func printJSON(platforms []PlatformData) error {
	jsonData, err := json.Marshal(platforms)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	fmt.Printf("%s\n", jsonData)
	return nil
}

// printRaw prints the platforms in raw format
func printRaw(platforms []PlatformData) {
	for _, platform := range platforms {
		fmt.Println(platform)
	}
}

// printTable prints the platforms in table format
func printTable(platforms []PlatformData, platformName string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Description"})
	table.SetBorder(true)

	for _, platform := range platforms {
		if platformName == "" || platform.Name == platformName {
			table.Append([]string{platform.Name, platform.Description})
		}
	}
	table.Render()
}
