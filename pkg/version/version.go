package version

import (
	"fmt"
	"runtime"
)

var (
	Version   string              // "v0.0.1"
	GoVersion = runtime.Version() // "1.19"
	Os        = runtime.GOOS      // linux
	Arch      = runtime.GOARCH    // amd64
	BuildDate string              // 2022-04-22_14:29:29
	Revision  string              // last Git commit hash
)

// Show function print version informations in standard output.
func Show() {
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("GoVersion: %s\n", GoVersion)
	fmt.Printf("Os: %s\n", Os)
	fmt.Printf("Arch: %s\n", Arch)
	fmt.Printf("BuildDate: %s\n", BuildDate)
	fmt.Printf("Revision: %s\n", Revision)
}
