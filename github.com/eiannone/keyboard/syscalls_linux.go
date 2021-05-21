package keyboard

import (
    "github.com/general252/EasyDarwinLib/golang.org/x/sys/unix"
)

const (
    ioctl_GETATTR = unix.TCGETS
    ioctl_SETATTR = unix.TCSETS
)