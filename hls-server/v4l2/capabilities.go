package v4l2

import (
	"unsafe"
)

type Capability uint32

var (
	V4L2_CAP_VIDEO_CAPTURE        Capability = 0x00000001
	V4L2_CAP_VIDEO_CAPTURE_MPLANE Capability = 0x00001000
	V4L2_CAP_VIDEO_OUTPUT         Capability = 0x00000002
	V4L2_CAP_VIDEO_OUTPUT_MPLANE  Capability = 0x00002000
	V4L2_CAP_VIDEO_M2M            Capability = 0x00004000
	V4L2_CAP_VIDEO_M2M_MPLANE     Capability = 0x00008000
	V4L2_CAP_VIDEO_OVERLAY        Capability = 0x00000004
	V4L2_CAP_VBI_CAPTURE          Capability = 0x00000010
	V4L2_CAP_VBI_OUTPUT           Capability = 0x00000020
	V4L2_CAP_SLICED_VBI_CAPTURE   Capability = 0x00000040
	V4L2_CAP_SLICED_VBI_OUTPUT    Capability = 0x00000080
	V4L2_CAP_RDS_CAPTURE          Capability = 0x00000100
	V4L2_CAP_VIDEO_OUTPUT_OVERLAY Capability = 0x00000200
	V4L2_CAP_HW_FREQ_SEEK         Capability = 0x00000400
	V4L2_CAP_RDS_OUTPUT           Capability = 0x00000800
	V4L2_CAP_TUNER                Capability = 0x00010000
	V4L2_CAP_AUDIO                Capability = 0x00020000
	V4L2_CAP_RADIO                Capability = 0x00040000
	V4L2_CAP_MODULATOR            Capability = 0x00080000
	V4L2_CAP_SDR_CAPTURE          Capability = 0x00100000
	V4L2_CAP_EXT_PIX_FORMAT       Capability = 0x00200000
	V4L2_CAP_SDR_OUTPUT           Capability = 0x00400000
	V4L2_CAP_READWRITE            Capability = 0x01000000
	V4L2_CAP_ASYNCIO              Capability = 0x02000000
	V4L2_CAP_STREAMING            Capability = 0x04000000
	V4L2_CAP_TOUCH                Capability = 0x10000000
	V4L2_CAP_DEVICE_CAPS          Capability = 0x80000000
)

func (c Capability) String() string {
	switch c {
	case V4L2_CAP_VIDEO_CAPTURE:
		return "V4L2_CAP_VIDEO_CAPTURE"
	case V4L2_CAP_VIDEO_CAPTURE_MPLANE:
		return "V4L2_CAP_VIDEO_CAPTURE_MPLANE"
	case V4L2_CAP_VIDEO_OUTPUT:
		return "V4L2_CAP_VIDEO_OUTPUT"
	case V4L2_CAP_VIDEO_OUTPUT_MPLANE:
		return "V4L2_CAP_VIDEO_OUTPUT_MPLANE"
	case V4L2_CAP_VIDEO_M2M:
		return "V4L2_CAP_VIDEO_M2M"
	case V4L2_CAP_VIDEO_M2M_MPLANE:
		return "V4L2_CAP_VIDEO_M2M_MPLANE"
	case V4L2_CAP_VIDEO_OVERLAY:
		return "V4L2_CAP_VIDEO_OVERLAY"
	case V4L2_CAP_VBI_CAPTURE:
		return "V4L2_CAP_VBI_CAPTURE"
	case V4L2_CAP_VBI_OUTPUT:
		return "V4L2_CAP_VBI_OUTPUT"
	case V4L2_CAP_SLICED_VBI_CAPTURE:
		return "V4L2_CAP_SLICED_VBI_CAPTURE"
	case V4L2_CAP_SLICED_VBI_OUTPUT:
		return "V4L2_CAP_SLICED_VBI_OUTPUT"
	case V4L2_CAP_RDS_CAPTURE:
		return "V4L2_CAP_RDS_CAPTURE"
	case V4L2_CAP_VIDEO_OUTPUT_OVERLAY:
		return "V4L2_CAP_VIDEO_OUTPUT_OVERLAY"
	case V4L2_CAP_HW_FREQ_SEEK:
		return "V4L2_CAP_HW_FREQ_SEEK"
	case V4L2_CAP_RDS_OUTPUT:
		return "V4L2_CAP_RDS_OUTPUT"
	case V4L2_CAP_TUNER:
		return "V4L2_CAP_TUNER"
	case V4L2_CAP_AUDIO:
		return "V4L2_CAP_AUDIO"
	case V4L2_CAP_RADIO:
		return "V4L2_CAP_RADIO"
	case V4L2_CAP_MODULATOR:
		return "V4L2_CAP_MODULATOR"
	case V4L2_CAP_SDR_CAPTURE:
		return "V4L2_CAP_SDR_CAPTURE"
	case V4L2_CAP_EXT_PIX_FORMAT:
		return "V4L2_CAP_EXT_PIX_FORMAT"
	case V4L2_CAP_SDR_OUTPUT:
		return "V4L2_CAP_SDR_OUTPUT"
	case V4L2_CAP_READWRITE:
		return "V4L2_CAP_READWRITE"
	case V4L2_CAP_ASYNCIO:
		return "V4L2_CAP_ASYNCIO"
	case V4L2_CAP_STREAMING:
		return "V4L2_CAP_STREAMING"
	case V4L2_CAP_TOUCH:
		return "V4L2_CAP_TOUCH"
	case V4L2_CAP_DEVICE_CAPS:
		return "V4L2_CAP_DEVICE_CAPS"
	default:
		return "unknown capability"
	}
}

func (c Capability) Capabilities() []Capability {
	capabilities := []Capability{}
	for i := 0; i < 32; i++ {
		if c&1<<i > 0 {
			capabilities = append(capabilities, 1<<i)
		}
	}
	return capabilities
}

// ref https://www.kernel.org/doc/html/v4.9/media/uapi/v4l/vidioc-querycap.html
func (d *Device) QueryCapabilities() (uint32, uint32, error) {
	v4l2_capability := struct {
		driver       [16]byte
		card         [32]byte
		busInfo      [32]byte
		version      uint32
		capabilities uint32
		deviceCaps   uint32
		_            [3]uint32 // reserved
	}{}

	err := ioctl(d.Fd, VIDIOC_QUERYCAP, unsafe.Pointer(&v4l2_capability))
	if err != 0 {
		return v4l2_capability.capabilities, v4l2_capability.deviceCaps, err
	}
	return v4l2_capability.capabilities, v4l2_capability.deviceCaps, nil
}
