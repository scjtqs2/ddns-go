package app

import (
	"fmt"
	"os/exec"
	"runtime"
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
		fmt.Printf("extscript exec faild cmd=%s , err=%s", cmd, err.Error())
		return "", err
	}
	fmt.Printf("extscript exec success cmd=%d , rsp=%s", cmd, string(res))
	return string(res), nil
}
