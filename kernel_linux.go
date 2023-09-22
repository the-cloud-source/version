package version

import "golang.org/x/sys/unix"

func init() {
	utsname := unix.Utsname{}
	unix.Uname(&utsname)
	kernel = str(utsname.Release)
}

func str(a [65]byte) string {

	s := ""
	for _, c := range a {
		if c != 0 {
			s += string(c)
		}
	}
	return string(s)
}
