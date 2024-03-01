package version

import (
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/klauspost/cpuid/v2"
)

// These vars get bound at build time using the --ldflags mechanism.
var (
	Version   = "0.0.0"
	Build     = "0000"
	GitCommit = "unknown"
	GitBranch = "unknown"
	BuildTime = "unset"
	Name      = "binary-name-missing"

	runArch string
	bldArch string

	userAgent = ""
	vString   = ""
	webHash   = ""
)

func init() {

	switch runtime.GOARCH {
	case "amd64":
		runArch = runtime.GOARCH + "-v" + strconv.Itoa(cpuid.CPU.X64Level())
		bldArch = runtime.GOARCH + "-v0"
		bi, ok := debug.ReadBuildInfo()
		if ok {
			for _, a := range bi.Settings {
				if a.Key == "GOAMD64" {
					bldArch = runtime.GOARCH + "-" + a.Value
				}
			}
		}
	default:
		runArch = runtime.GOARCH
		bldArch = runtime.GOARCH
	}

	userAgent = ua()
	vString = v()
	webHash = buildWebHash()
}

func UserAgent() string {
	return userAgent
}

func String() string {
	return vString
}

func ua() string {

	var sb strings.Builder

	sb.WriteString(Name)
	if Version != "0.0.0" {
		sb.WriteString(" version=")
		sb.WriteString(Version)
	}

	if Build != "0000" || Version == "0.0.0" {
		sb.WriteString(" build=")
		sb.WriteString(Build)
	}

	sb.WriteString(" os=")
	sb.WriteString(runtime.GOOS)
	sb.WriteString("/")
	sb.WriteString(bldArch)
	if bldArch != runArch {
		sb.WriteString("/")
		sb.WriteString(runArch)
	}

	return sb.String()
}

func v() string {
	built := BuildTime
	i, err := strconv.ParseInt(BuildTime, 10, 64)
	if err == nil {
		t := time.Unix(i, 0).UTC()
		built = t.Format(time.RFC3339)
	}

	var sb strings.Builder

	sb.WriteString(Name)
	if Version != "0.0.0" {
		sb.WriteString(" version=")
		sb.WriteString(Version)
	}
	if Build != "0000" || Version == "0.0.0" {
		sb.WriteString(" build=")
		sb.WriteString(Build)
	}

	sb.WriteString(" git=")
	sb.WriteString(GitCommit)
	if GitBranch != "unknown" {
		sb.WriteString(" branch=")
		sb.WriteString(GitBranch)
	}

	sb.WriteString(" ")
	sb.WriteString(built)

	sb.WriteString(" ")
	sb.WriteString(runtime.Version())
	sb.WriteString("/")
	sb.WriteString(bldArch)

	if bldArch != runArch {
		sb.WriteString(" on ")
		sb.WriteString(runArch)
	}

	return sb.String()
}

func WebHash() string {
	return webHash
}

func buildWebHash() string {
	return strings.Replace(GitCommit+"-"+BuildTime, "+", "-", -1)
}

var kernel = "unknown"

func Kernel() string {
	return kernel
}
