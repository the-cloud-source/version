//go:build windows

package version

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func init() {
	v := windows.RtlGetVersion()

	kernel = fmt.Sprintf("%d.%d.%d-%#x-%#x-%#x",
		v.MajorVersion, v.MinorVersion, v.BuildNumber,
		v.ProductType, v.SuiteMask, v.PlatformId)
}
