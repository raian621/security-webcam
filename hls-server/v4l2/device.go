package v4l2

import (
	"errors"

	"golang.org/x/sys/unix"
)

type Device struct {
	Path string
	Fd   uintptr
}

func NewDevice(path string) *Device { return &Device{Path: path} }

func (d *Device) Open(mode int) error {
	fd, err := unix.Open(d.Path, mode, 0)
	if err != nil {
		return err
	}
	if fd < 0 {
		return errors.New("unable to open video device")
	}

	d.Fd = uintptr(fd)
	return nil
}

func (d *Device) Close() error {
	err := unix.Close(int(d.Fd))
	d.Fd = 0
	return err
}
