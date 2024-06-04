package main

// import "testing"

// func TestParseSettingsOutput(t *testing.T) {
// 	t.Parallel()

// 	// wantWebcamFormats := []WebcamFormat{
// 	// 	{
// 	// 		FormatName:     "dsafasd",
// 	// 		WebcamSettings: []WebcamSettings{},
// 	// 	},
// 	// }

// 	mockedOutput := `ioctl: VIDIOC_ENUM_FMT

// 		Type: Video Capture

// 		[0]: 'YUYV' (YUYV 4:2:2)
// 						Size: Discrete 640x480
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 160x120
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 176x144
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 320x180
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 320x240
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 352x288
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 424x240
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 480x270
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 640x360
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 800x448
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 800x600
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 848x480
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 960x540
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 1024x576
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 1280x720
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 1600x896
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 1920x1080
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 2304x1296
// 										Interval: Discrete 0.500s (2.000 fps)
// 						Size: Discrete 2304x1536
// 										Interval: Discrete 0.500s (2.000 fps)
// 		[1]: 'MJPG' (Motion-JPEG, compressed)
// 						Size: Discrete 640x480
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 160x120
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 320x240
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 352x288
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 424x240
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 480x270
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 640x360
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 800x448
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 800x600
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 848x480
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 960x540
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 1024x576
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 1280x720
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 1600x896
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 1920x1080
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 	*/               Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 176x144
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 320x180
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 320x240
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 352x288
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 424x240
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 480x270
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 640x360
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 800x448
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 800x600
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 848x480
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 960x540
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 1024x576
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 1280x720
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 1600x896
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// 						Size: Discrete 1920x1080
// 										Interval: Discrete 0.033s (30.000 fps)
// 										Interval: Discrete 0.042s (24.000 fps)
// 										Interval: Discrete 0.050s (20.000 fps)
// 										Interval: Discrete 0.067s (15.000 fps)
// 										Interval: Discrete 0.100s (10.000 fps)
// 										Interval: Discrete 0.133s (7.500 fps)
// 										Interval: Discrete 0.200s (5.000 fps)
// `
// 	parseVideoSettingsOutput([]byte(mockedOutput))
// 	t.Fail()
// }