package netbox

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/olekukonko/tablewriter"
)

// ModuleData holds the details of a module to be displayed
type ModuleData struct {
	Name   string
	Serial string
}

// listModulesPayload retrieves the module data from NetBox
func listModulesPayload(netboxHost, token string, httpScheme string) ([]ModuleData, error) {
	c := newNetboxClient(netboxHost, token, httpScheme)
	params := dcim.NewDcimModulesListParams()
	resp, err := c.Dcim.DcimModulesList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot get modules list: %w", err)
	}

	var modules []ModuleData

	for _, module := range resp.Payload.Results {
		modules = append(modules, ModuleData{
			Name:   getString(module.Device.Name),
			Serial: getString(module.Serial),
		})
	}

	return modules, nil
}

// PrintModulesList prints the list of modules in the desired format
func PrintModulesList(netboxHost, token string, httpScheme string, jsonOpt bool, rawOpt bool, moduleName string) error {
	modules, err := listModulesPayload(netboxHost, token, httpScheme)
	if err != nil {
		return err
	}

	// Print output in JSON format
	if jsonOpt {
		return printJSON(modules)
	}

	// Print result in raw format
	if rawOpt {
		printRaw(modules)
		return nil
	}

	// Print result in table format
	printTable(modules, moduleName)
	return nil
}

// getString safely dereferences a string pointer
func getString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

// printJSON prints the modules in JSON format
func printJSON(modules []ModuleData) error {
	jsonData, err := json.Marshal(modules)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	fmt.Printf("%s\n", jsonData)
	return nil
}

// printRaw prints the modules in raw format
func printRaw(modules []ModuleData) {
	for _, module := range modules {
		fmt.Println(module)
	}
}

// printTable prints the modules in table format
func printTable(modules []ModuleData, moduleName string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Serial"})
	table.SetBorder(true)

	for _, module := range modules {
		if moduleName == "" || module.Name == moduleName {
			table.Append([]string{module.Name, module.Serial})
		}
	}
	table.Render()
}
