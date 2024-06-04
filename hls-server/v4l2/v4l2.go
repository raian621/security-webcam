package v4l2

// interfaces with the v4l2 API

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

func ioctl(fd, req uintptr, arg unsafe.Pointer) unix.Errno {
	for {
		_, _, errno := unix.Syscall(unix.SYS_IOCTL, fd, req, uintptr(arg))

		// retry if the system call was interrupted
		if errno == unix.EINTR {
			continue
		}

		// otherwise, return errno
		return errno
	}
}
