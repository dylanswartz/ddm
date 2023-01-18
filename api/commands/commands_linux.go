package commands

import (
	"syscall"
)

func Reboot() {
    syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART)
}