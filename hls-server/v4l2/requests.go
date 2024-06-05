package v4l2

var (
	// _IOC_NONE  uint32 = 0
	_IOC_WRITE uint32 = 1
	_IOC_READ  uint32 = 2

	// [dir:2] [size:14] [type:8] [nr:8]
	_IOC_NRBITS   uint32 = 8
	_IOC_TYPEBITS uint32 = 8
	_IOC_SIZEBITS uint32 = 14

	_IOC_NRSHIFT   uint32 = 0
	_IOC_TYPESHIFT uint32 = (_IOC_NRSHIFT + _IOC_NRBITS)
	_IOC_SIZESHIFT uint32 = (_IOC_TYPESHIFT + _IOC_TYPEBITS)
	_IOC_DIRSHIFT  uint32 = (_IOC_SIZESHIFT + _IOC_SIZEBITS)
)

func ioctlEncode(dir, _type, nr, size uint32) uintptr {
	return uintptr(((dir << _IOC_DIRSHIFT) | (_type << _IOC_TYPESHIFT) |
		(nr << _IOC_NRSHIFT) | (size << _IOC_SIZESHIFT)))
}

// func ioctlEncodeNone(_type byte, nr, size uint32) uintptr {
// 	return ioctlEncode(_IOC_NONE, uint32(_type), nr, size)
// }

func ioctlEncodeR(_type byte, nr, size uint32) uintptr {
	return ioctlEncode(_IOC_READ, uint32(_type), nr, size)
}

// func ioctlEncodeW(_type byte, nr, size uint32) uintptr {
// 	return ioctlEncode(_IOC_WRITE, uint32(_type), nr, size)
// }

func ioctlEncodeWR(_type byte, nr, size uint32) uintptr {
	return ioctlEncode(_IOC_READ|_IOC_WRITE, uint32(_type), nr, size)
}

var (
	VIDIOC_QUERYCAP = ioctlEncodeR('V', 0, 104)
	VIDIOC_ENUM_FMT = ioctlEncodeWR('V', 2, 64)
	// VIDIOC_G_FMT               = ioctlEncodeWR('V', 4, v4l2_format)
	// VIDIOC_S_FMT               = ioctlEncodeWR('V', 5, v4l2_format)
	// VIDIOC_REQBUFS             = ioctlEncodeWR('V', 8, v4l2_requestbuffers)
	// VIDIOC_QUERYBUF            = ioctlEncodeWR('V', 9, v4l2_buffer)
	// VIDIOC_G_FBUF              = ioctlEncodeR('V', 10, v4l2_framebuffer)
	// VIDIOC_S_FBUF              = ioctlEncodeW('V', 11, v4l2_framebuffer)
	// VIDIOC_OVERLAY             = ioctlEncodeW('V', 14, int)
	// VIDIOC_QBUF                = ioctlEncodeWR('V', 15, v4l2_buffer)
	// VIDIOC_EXPBUF              = ioctlEncodeWR('V', 16, v4l2_exportbuffer)
	// VIDIOC_DQBUF               = ioctlEncodeWR('V', 17, v4l2_buffer)
	// VIDIOC_STREAMON            = ioctlEncodeW('V', 18, int)
	// VIDIOC_STREAMOFF           = ioctlEncodeW('V', 19, int)
	// VIDIOC_G_PARM              = ioctlEncodeWR('V', 21, v4l2_streamparm)
	// VIDIOC_S_PARM              = ioctlEncodeWR('V', 22, v4l2_streamparm)
	// VIDIOC_G_STD               = ioctlEncodeR('V', 23, v4l2_std_id)
	// VIDIOC_S_STD               = ioctlEncodeW('V', 24, v4l2_std_id)
	// VIDIOC_ENUMSTD             = ioctlEncodeWR('V', 25, v4l2_standard)
	// VIDIOC_ENUMINPUT           = ioctlEncodeWR('V', 26, v4l2_input)
	// VIDIOC_G_CTRL              = ioctlEncodeWR('V', 27, v4l2_control)
	// VIDIOC_S_CTRL              = ioctlEncodeWR('V', 28, v4l2_control)
	// VIDIOC_G_TUNER             = ioctlEncodeWR('V', 29, v4l2_tuner)
	// VIDIOC_S_TUNER             = ioctlEncodeW('V', 30, v4l2_tuner)
	// VIDIOC_G_AUDIO             = ioctlEncodeR('V', 33, v4l2_audio)
	// VIDIOC_S_AUDIO             = ioctlEncodeW('V', 34, v4l2_audio)
	// VIDIOC_QUERYCTRL           = ioctlEncodeWR('V', 36, v4l2_queryctrl)
	// VIDIOC_QUERYMENU           = ioctlEncodeWR('V', 37, v4l2_querymenu)
	// VIDIOC_G_INPUT             = ioctlEncodeR('V', 38, int)
	// VIDIOC_S_INPUT             = ioctlEncodeWR('V', 39, int)
	// VIDIOC_G_EDID              = ioctlEncodeWR('V', 40, v4l2_edid)
	// VIDIOC_S_EDID              = ioctlEncodeWR('V', 41, v4l2_edid)
	// VIDIOC_G_OUTPUT            = ioctlEncodeR('V', 46, int)
	// VIDIOC_S_OUTPUT            = ioctlEncodeWR('V', 47, int)
	// VIDIOC_ENUMOUTPUT          = ioctlEncodeWR('V', 48, v4l2_output)
	// VIDIOC_G_AUDOUT            = ioctlEncodeR('V', 49, v4l2_audioout)
	// VIDIOC_S_AUDOUT            = ioctlEncodeW('V', 50, v4l2_audioout)
	// VIDIOC_G_MODULATOR         = ioctlEncodeWR('V', 54, v4l2_modulator)
	// VIDIOC_S_MODULATOR         = ioctlEncodeW('V', 55, v4l2_modulator)
	// VIDIOC_G_FREQUENCY         = ioctlEncodeWR('V', 56, v4l2_frequency)
	// VIDIOC_S_FREQUENCY         = ioctlEncodeW('V', 57, v4l2_frequency)
	// VIDIOC_CROPCAP             = ioctlEncodeWR('V', 58, v4l2_cropcap)
	// VIDIOC_G_CROP              = ioctlEncodeWR('V', 59, v4l2_crop)
	// VIDIOC_S_CROP              = ioctlEncodeW('V', 60, v4l2_crop)
	// VIDIOC_G_JPEGCOMP          = ioctlEncodeR('V', 61, v4l2_jpegcompression)
	// VIDIOC_S_JPEGCOMP          = ioctlEncodeW('V', 62, v4l2_jpegcompression)
	// VIDIOC_QUERYSTD            = ioctlEncodeR('V', 63, v4l2_std_id)
	// VIDIOC_TRY_FMT             = ioctlEncodeWR('V', 64, v4l2_format)
	// VIDIOC_ENUMAUDIO           = ioctlEncodeWR('V', 65, v4l2_audio)
	// VIDIOC_ENUMAUDOUT          = ioctlEncodeWR('V', 66, v4l2_audioout)
	// VIDIOC_G_PRIORITY          = ioctlEncodeR('V', 67, uint32) /* enum v4l2_priority */
	// VIDIOC_S_PRIORITY          = ioctlEncodeW('V', 68, uint32) /* enum v4l2_priority */
	// VIDIOC_G_SLICED_VBI_CAP    = ioctlEncodeWR('V', 69, v4l2_sliced_vbi_cap)
	// VIDIOC_LOG_STATUS          = ioctlEncodeNone('V', 70)
	// VIDIOC_G_EXT_CTRLS         = ioctlEncodeWR('V', 71, v4l2_ext_controls)
	// VIDIOC_S_EXT_CTRLS         = ioctlEncodeWR('V', 72, v4l2_ext_controls)
	// VIDIOC_TRY_EXT_CTRLS       = ioctlEncodeWR('V', 73, v4l2_ext_controls)
	VIDIOC_ENUM_FRAMESIZES     = ioctlEncodeWR('V', 74, 44)
	VIDIOC_ENUM_FRAMEINTERVALS = ioctlEncodeWR('V', 75, 52)
	// VIDIOC_G_ENC_INDEX         = ioctlEncodeR('V', 76, v4l2_enc_idx)
	// VIDIOC_ENCODER_CMD         = ioctlEncodeWR('V', 77, v4l2_encoder_cmd)
	// VIDIOC_TRY_ENCODER_CMD     = ioctlEncodeWR('V', 78, v4l2_encoder_cmd)
)
