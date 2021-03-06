// +build windows

// Copyright (c) 2018, The TurtleCoin Developers
// Copyright (c) 2019, 2ACoin Developers
//
// Please see the included LICENSE file for more information.
//

package walletdmanager

import (
	"os/exec"
	"syscall"
)

func hideCmdWindowIfNeeded(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}
