package v4l2

import (
	"bytes"
	"unsafe"

	"golang.org/x/sys/unix"
)

// https://www.kernel.org/doc/html/v4.9/media/uapi/v4l/vidioc-enum-fmt.html#fmtdesc-flags

var (
	V4L2_FMT_FLAG_COMPRESSED uint32 = 0x0001
	V4L2_FMT_FLAG_EMULATED   uint32 = 0x0002
)

type Format struct {
	PixelFormat uint32
	Description string
	Flags       uint32
}

func (f Format) PixelFormatString() string {
	return string(
		[]byte{
			byte(f.PixelFormat),
			byte(f.PixelFormat >> 8),
			byte(f.PixelFormat >> 16),
			byte(f.PixelFormat >> 24),
		},
	)
}

func (d *Device) QueryFormats(streamType uint32) ([]Format, error) {
	v4l2_fmtdesc := struct {
		index       uint32
		streamType  uint32
		flags       uint32
		description [32]byte
		pixelFormat uint32
		_           [4]uint32 // reserved for "future use"
	}{
		index:      0,
		streamType: streamType,
	}

	formats := make([]Format, 0)

	for {
		errno := ioctl(d.Fd, VIDIOC_ENUM_FMT, unsafe.Pointer(&v4l2_fmtdesc))
		if errno == unix.EINVAL {
			break
		} else if errno != 0 {
			return []Format{}, errno
		}

		// find index of null terminator
		n := bytes.IndexByte(v4l2_fmtdesc.description[:], 0)
		formats = append(formats, Format{
			PixelFormat: v4l2_fmtdesc.pixelFormat,
			Description: string(v4l2_fmtdesc.description[:n]),
			Flags:       v4l2_fmtdesc.flags,
		})

		v4l2_fmtdesc.index++
	}

	return formats, nil
}
