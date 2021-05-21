// +build !windows,!linux

package keyboard

import (
    "github.com/general252/EasyDarwinLib/golang.org/x/sys/unix"
)

const (
    ioctl_GETATTR = unix.TIOCGETA
    ioctl_SETATTR = unix.TIOCSETA
)