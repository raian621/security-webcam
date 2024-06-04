package main

import (
	"hls-server/v4l2"
	"log"

	"golang.org/x/sys/unix"
)

// type SupportedFormat struct {
// 	PixelFormat string                `json="pixelfmt`
// 	Resolutions ResolutionWithBitrate `json="resolutions"`
// }

type ResolutionWithBitrate struct {
	Width   uint32 `json:"width"`
	Height  uint32 `json:"height"`
	MaxKbps uint32
}

// var supportedFormats = make([]SupportedFormat, 0)
var resolutions = make([]ResolutionWithBitrate, 0)

var supportedResolutions = []ResolutionWithBitrate{
	{Width: 480, Height: 360, MaxKbps: 600},
	{Width: 640, Height: 480, MaxKbps: 1500},
	{Width: 1280, Height: 720, MaxKbps: 3000},
	{Width: 1920, Height: 1080, MaxKbps: 6000},
	{Width: 3840, Height: 2160, MaxKbps: 12000},
}

func getResolutions(videoDevice string) error {
	resolutions = []ResolutionWithBitrate{}

	device := v4l2.NewDevice(videoDevice)
	device.Open(unix.O_RDONLY)
	defer func() {
		if err := device.Close(); err != nil {
			log.Println("unexpected error while closing video device")
		}
	}()
	devicePixelFormats, err := device.QueryFormats(v4l2.V4L2_BUF_TYPE_VIDEO_CAPTURE)
	if err != nil {
		return err
	}
	deviceResolutions, err := device.QueryResolutions(devicePixelFormats[0].PixelFormat)
	if err != nil {
		return err
	}

	for _, dr := range deviceResolutions {
		// binary search to find if the current resolution is supported
		l, r := 0, len(supportedResolutions)-1

		for l <= r {
			m := (l + r) / 2

			if supportedResolutions[m].Width < dr.Width {
				l = m + 1
			} else if supportedResolutions[m].Height > dr.Height {
				r = m - 1
			} else {
				if supportedResolutions[m].Height < dr.Height {
					l = m + 1
				} else if supportedResolutions[m].Height > dr.Height {
					r = m - 1
				} else {
					resolutions = append(resolutions, supportedResolutions[m])
					break
				}
			}
		}
	}

	return nil
}
