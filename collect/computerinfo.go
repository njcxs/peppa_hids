package collect

import (
	"io/ioutil"
	"os"
	"strings"
)

// GetComInfo 获取计算机信息
func GetComInfo() (info ComputerInfo) {
	info.IP = LocalIP
	info.Hostname, _ = os.Hostname()
	out := Cmdexec("uname -r")
	dat, err := ioutil.ReadFile("/etc/redhat-release")
	if err != nil {
		dat, _ = ioutil.ReadFile("/etc/issue")
		issue := strings.SplitN(string(dat), "\n", 2)[0]
		out2 := Cmdexec("uname -m")
		info.System = issue + " " + out + out2
	} else {
		info.System = string(dat) + " " + out
	}
	discern(&info)
	return info
}
