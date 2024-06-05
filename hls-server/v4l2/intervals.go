package v4l2

import (
	"encoding/binary"
	"errors"
	"unsafe"

	"golang.org/x/sys/unix"
)

var (
	V4L2_FRMIVAL_TYPE_DISCRETE   uint32 = 1 // discrete frame interval
	V4L2_FRMIVAL_TYPE_CONTINUOUS uint32 = 1 // discrete frame interval
	V4L2_FRMIVAL_TYPE_STEPWISE   uint32 = 1 // discrete frame interval
)

type V4L2_frmival_stepwise struct {
	Min      V4L2_fract
	Max      V4L2_fract
	Stepwise V4L2_fract
}

type V4L2_frmivalenum struct {
	// IN: Index of the given frame interval in the enumeration.
	Index uint32
	// IN: Pixel format for which the frame intervals are enumerated.
	PixelFormat uint32
	// IN: Frame width for which the frame intervals are enumerated.
	Width uint32
	// IN: Frame height for which the frame intervals are enumerated.
	Height uint32
	// OUT: Frame interval type the device supports.
	Type uint32
	// OUT: Frame interval with the given index.
	FrameIntervalBytes [24]byte
	// 8 bytes are reserved for "future use"
	_ [2]uint32
}

type V4L2_FrmiValEnumDiscrete struct {
	Index       uint32
	PixelFormat uint32
	Width       uint32
	Height      uint32
	Interval    V4L2_fract
}

type V4L2_FrmiValEnumContinuous V4L2_FrmiValEnumDiscrete

type V4L2_FrmiValEnumStepwise struct {
	Index            uint32
	PixelFormat      uint32
	Width            uint32
	Height           uint32
	StepwiseInterval V4L2_frmival_stepwise
}

func VidiocEnumFrameIntervals(fd uintptr, argp *V4L2_frmivalenum) error {
	arg := unsafe.Pointer(argp)
	errno := ioctl(fd, VIDIOC_ENUM_FRAMEINTERVALS, arg)

	if errno != 0 {
		return errors.New("error")
	}

	return nil
}

type Interval struct {
	Numerator   uint32
	Denominator uint32
}

func MarshalInterval(intervalBytes [24]byte) Interval {
	return Interval{
		Numerator:   binary.LittleEndian.Uint32(intervalBytes[:4]),
		Denominator: binary.LittleEndian.Uint32(intervalBytes[4:8]),
	}
}

// https://www.kernel.org/doc/html/v4.9/media/uapi/v4l/vidioc-enum-frameintervals.html
func (d *Device) QueryIntervals(pixfmt, width, height uint32) ([]Interval, error) {
	intervals := []Interval{}

	v4l2_frmivalenum := struct {
		index         uint32
		pixelFormat   uint32
		width         uint32
		height        uint32
		intervalType  uint32
		intervalBytes [24]byte
		_             [2]uint32
	}{
		index:       0,
		pixelFormat: pixfmt,
		width:       width,
		height:      height,
	}

	for {
		errno := ioctl(d.Fd, VIDIOC_ENUM_FRAMEINTERVALS, unsafe.Pointer(&v4l2_frmivalenum))

		if errno == unix.EINVAL {
			break
		} else if errno != 0 {
			return []Interval{}, errno
		}

		intervals = append(intervals, MarshalInterval(v4l2_frmivalenum.intervalBytes))
		v4l2_frmivalenum.index++
	}

	return intervals, nil
}

/*
Frame intervals and frame rates: The V4L2 API uses frame intervals instead of
frame rates. Given the frame interval the frame rate can be computed as follows:

frame_rate = 1 / frame_interval
*/
// func frameRate(interval V4L2_fract) float32 {
// 	return float32(interval.Demoninator) / float32(interval.Numerator)
// }
