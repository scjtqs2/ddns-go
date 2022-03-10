package app

import (
	log "github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
	"strings"
)

// Command 执行命令
func Command(cmd string) (string, error) {
	// 执行cmd
	var proc *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		proc = exec.Command("cmd", "/C", cmd)
		break
	case "drawin":
		proc = exec.Command("bash", "-c", cmd)
		break
	default:
		proc = exec.Command("sh", "-c", cmd)
		break
	}
	res, err := proc.Output()
	if err != nil {
		log.Errorf("extscript exec faild cmd=%s , err=%s", cmd, err.Error())
		return "", err
	}
	log.Infof("extscript exec success cmd=%s , rsp=%s", cmd, string(res))

	return strings.TrimSpace(string(res)), nil
}
