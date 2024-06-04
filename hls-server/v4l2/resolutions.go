package v4l2

import (
	"encoding/binary"
	"unsafe"

	"golang.org/x/sys/unix"
)

// https://www.kernel.org/doc/html/v4.9/media/uapi/v4l/vidioc-enum-framesizes.html#vidioc-enum-framesizes

var (
	V4L2_FRMSIZE_TYPE_DISCRETE   uint32 = 1
	V4L2_FRMSIZE_TYPE_CONTINUOUS uint32 = 2
	V4L2_FRMSIZE_TYPE_STEPWISE   uint32 = 3
)

type V4L2_frmsize_discrete struct {
	Width  uint32
	Height uint32
}

type V4L2_frmsize_stepwise struct {
	MinWidth   uint32
	MaxWidth   uint32
	StepWidth  uint32
	MinHeight  uint32
	MaxHeight  uint32
	StepHeight uint32
}

type V4L2_frmsizeenum struct {
	Index          uint32
	PixelFormat    uint32
	Type           uint32
	FrameSizeBytes [24]byte
	_              [2]uint32
}

type Resolution struct {
	Width  uint32
	Height uint32
}

func MarshalResolution(resBytes [24]byte) Resolution {
	width := binary.LittleEndian.Uint32(resBytes[:4])
	height := binary.LittleEndian.Uint32(resBytes[4:8])

	return Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}
}

func (d *Device) QueryResolutions(pixfmt uint32) ([]Resolution, error) {
	resolutions := make([]Resolution, 0)
	v4l2_frmsizeenum := struct {
		index          uint32
		pixelFormat    uint32
		resolutionType uint32
		frameSizeBytes [24]byte
		_              [2]uint32
	}{
		index:       0,
		pixelFormat: pixfmt,
	}

	for {
		errno := ioctl(d.Fd, VIDIOC_ENUM_FRAMESIZES, unsafe.Pointer(&v4l2_frmsizeenum))

		if errno == unix.EINVAL {
			break
		} else if errno != 0 {
			return []Resolution{}, errno
		}

		resolutions = append(resolutions, MarshalResolution(v4l2_frmsizeenum.frameSizeBytes))
		v4l2_frmsizeenum.index++
	}

	return resolutions, nil
}
