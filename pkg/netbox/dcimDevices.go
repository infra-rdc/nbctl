package netbox

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
	"github.com/olekukonko/tablewriter"
)

// DeviceData holds the details of a device to be displayed
type DeviceData struct {
	Name     string
	Type     string
	Tenant   string
	Serial   string
	Location string
	Site     string
	Rack     string
	Status   string
}

// listDevicesPayload retrieves the device data from NetBox
func listDevicesPayload(netboxHost, token string, httpScheme string, rackName string) ([]DeviceData, error) {
	c := newNetboxClient(netboxHost, token, httpScheme)
	params := dcim.NewDcimDevicesListParams()
	resp, err := c.Dcim.DcimDevicesList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot get devices list: %w", err)
	}

	var devices []DeviceData

	for _, device := range resp.Payload.Results {
		status := *device.Status.Value
		if status == "active" {
			status = color.GreenString(status)
		} else if status == "offline" {
			status = color.RedString(status)
		}
		if rackName == "" || (device.Rack != nil && device.Rack.Name != nil && *device.Rack.Name == rackName) {
			devices = append(devices, DeviceData{
				Name:     getString(device.Name),
				Type:     getString(device.DeviceType.Model),
				Tenant:   getString(device.Tenant.Name),
				Serial:   getString(device.Serial),
				Location: getString(device.Location.Name),
				Site:     getString(device.Site.Name),
				Rack:     getString(device.Rack.Name),
				Status:   status,
			})
		}
	}

	return devices, nil
}

// PrintDevicesList prints the list of devices in the desired format
func PrintDevicesList(netboxHost, token string, httpScheme string, jsonOpt bool, rawOpt bool, deviceName string, rackName string) error {
	devices, err := listDevicesPayload(netboxHost, token, httpScheme, rackName)
	if err != nil {
		return err
	}

	// Print output in JSON format
	if jsonOpt {
		return printJSON(devices)
	}

	// Print result in raw format
	if rawOpt {
		printRaw(devices)
		return nil
	}

	// Print result in table format
	printTable(devices, deviceName)
	return nil
}

// getString safely dereferences a string pointer
func getString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

// printJSON prints the devices in JSON format
func printJSON(devices []DeviceData) error {
	jsonData, err := json.Marshal(devices)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	fmt.Printf("%s\n", jsonData)
	return nil
}

// printRaw prints the devices in raw format
func printRaw(devices []DeviceData) {
	for _, device := range devices {
		fmt.Println(device)
	}
}

// printTable prints the devices in table format
func printTable(devices []DeviceData, deviceName string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Type", "Tenant", "Serial", "Location", "Site", "Rack", "Status"})
	table.SetBorder(true)

	for _, device := range devices {
		if deviceName == "" || device.Name == deviceName {
			table.Append([]string{
				device.Name, device.Type, device.Tenant, device.Serial,
				device.Location, device.Site, device.Rack, device.Status,
			})
		}
	}
	table.Render()
}
