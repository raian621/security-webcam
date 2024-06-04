package v4l2

import (
	"testing"

	"golang.org/x/sys/unix"
)

func TestQueryCapabilities(t *testing.T) {
	device := NewDevice("/dev/video0")
	if err := device.Open(unix.O_RDWR); err != nil {
		t.Fatal(err)
	}
	defer device.Close()

	capabilities, deviceCaps, err := device.QueryCapabilities()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("capabilities: %v", Capability(capabilities).Capabilities())
	t.Logf("device capabilities: %s", Capability(deviceCaps).Capabilities())
	// t.Fail()
}

func TestQueryFormats(t *testing.T) {
	device := NewDevice("/dev/video0")
	if err := device.Open(unix.O_RDWR); err != nil {
		t.Fatal(err)
	}
	defer device.Close()

	formats, err := device.QueryFormats(V4L2_BUF_TYPE_VIDEO_CAPTURE)
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range formats {
		t.Log(f)
		t.Log(f.PixelFormatString())
	}

	// t.Fail()
}

func TestQueryResolutions(t *testing.T) {
	device := NewDevice("/dev/video0")
	if err := device.Open(unix.O_RDWR); err != nil {
		t.Fatal(err)
	}
	defer device.Close()

	formats, err := device.QueryFormats(V4L2_BUF_TYPE_VIDEO_CAPTURE)
	if err != nil {
		t.Fatal(err)
	}

	formatResolutions := []struct {
		format      *Format
		resolutions []Resolution
	}{}

	for _, f := range formats {
		f := f
		resolutions, err := device.QueryResolutions(f.PixelFormat)
		t.Log(resolutions)
		if err != nil {
			t.Fatal(err)
		}

		formatResolutions = append(formatResolutions, struct {
			format      *Format
			resolutions []Resolution
		}{
			format:      &f,
			resolutions: resolutions,
		})
	}

	for _, f := range formatResolutions {
		t.Log("format:     ", f.format.PixelFormatString())
		t.Log("description:", f.format.Description)
		for _, r := range f.resolutions {
			t.Logf("\t%dx%d\n", r.Width, r.Height)
		}
	}

	// t.Fail()
}

func TestQueryIntervals(t *testing.T) {
	device := NewDevice("/dev/video0")
	if err := device.Open(unix.O_RDWR); err != nil {
		t.Fatal(err)
	}
	defer device.Close()

	formats, err := device.QueryFormats(V4L2_BUF_TYPE_VIDEO_CAPTURE)
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range formats {
		t.Log(f.Description)
		t.Log(f.PixelFormatString())

		resolutions, err := device.QueryResolutions(f.PixelFormat)
		if err != nil {
			t.Fatal(err)
		}

		for _, r := range resolutions {
			intervals, err := device.QueryIntervals(f.PixelFormat, r.Width, r.Height)
			if err != nil {
				t.Fatal(err)
			}

			t.Logf("\t%dx%d\n", r.Width, r.Height)

			for _, i := range intervals {
				t.Logf(
					"\t\tinterval: %d / %d - %0.3f fps",
					i.Numerator,
					i.Denominator,
					float32(i.Denominator)/float32(i.Numerator))
			}
		}
	}

	// t.Fail()
}
