package v4l2

// https://www.kernel.org/doc/html/v4.9/media/uapi/v4l/buffer.html#c.v4l2_buf_type

var (
	// Buffer of a single-planar video capture stream
	V4L2_BUF_TYPE_VIDEO_CAPTURE uint32 = 1
	// Buffer of a multi-planar video capture stream
	V4L2_BUF_TYPE_VIDEO_CAPTURE_MPLANE uint32 = 9
	// Buffer of a single-planar video output stream
	V4L2_BUF_TYPE_VIDEO_OUTPUT uint32 = 2
	// Buffer of a multi-planar video output stream
	V4L2_BUF_TYPE_VIDEO_OUTPUT_MPLANE uint32 = 10
)
