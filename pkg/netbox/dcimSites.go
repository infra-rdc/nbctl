package netbox

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/olekukonko/tablewriter"
)

// SiteData holds the details of a site to be displayed
type SiteData struct {
	Name            string
	PhysicalAddress string
	TimeZone        string
}

// listSitesPayload retrieves the site data from NetBox
func listSitesPayload(netboxHost, token string, httpScheme string) ([]SiteData, error) {
	c := newNetboxClient(netboxHost, token, httpScheme)
	params := dcim.NewDcimSitesListParams()
	resp, err := c.Dcim.DcimSitesList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot get sites list: %w", err)
	}

	var sites []SiteData

	for _, site := range resp.Payload.Results {
		sites = append(sites, SiteData{
			Name:            getString(site.Name),
			PhysicalAddress: getString(site.PhysicalAddress),
			TimeZone:        getString(site.TimeZone),
		})
	}

	return sites, nil
}

// PrintSitesList prints the list of sites in the desired format
func PrintSitesList(netboxHost, token string, httpScheme string, jsonOpt bool, rawOpt bool, siteName string) error {
	sites, err := listSitesPayload(netboxHost, token, httpScheme)
	if err != nil {
		return err
	}

	// Print output in JSON format
	if jsonOpt {
		return printJSON(sites)
	}

	// Print result in raw format
	if rawOpt {
		printRaw(sites)
		return nil
	}

	// Print result in table format
	printTable(sites, siteName)
	return nil
}

// getString safely dereferences a string pointer
func getString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

// printJSON prints the sites in JSON format
func printJSON(sites []SiteData) error {
	jsonData, err := json.Marshal(sites)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	fmt.Printf("%s\n", jsonData)
	return nil
}

// printRaw prints the sites in raw format
func printRaw(sites []SiteData) {
	for _, site := range sites {
		fmt.Println(site)
	}
}

// printTable prints the sites in table format
func printTable(sites []SiteData, siteName string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Physical Address", "TimeZone"})
	table.SetBorder(true)

	for _, site := range sites {
		if siteName == "" || site.Name == siteName {
			table.Append([]string{site.Name, site.PhysicalAddress, site.TimeZone})
		}
	}
	table.Render()
}
