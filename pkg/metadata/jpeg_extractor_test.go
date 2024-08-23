package metadata

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// Mock function to capture output
func captureOutput(f func()) string {
	var buf bytes.Buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old
	buf.ReadFrom(r)
	return buf.String()
}

// Helper function to parse EXIF output into a map
func parseExifOutput(output string) map[string]string {
	exifMap := make(map[string]string)
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) == 2 {
			exifMap[parts[0]] = parts[1]
		}
	}
	return exifMap
}

func TestExtractJPEGMetadata(t *testing.T) {
	// Path to a sample JPEG file with known EXIF data
	sampleFilePath := "../../IMG_4229.JPG"

	// Expected EXIF properties (replace with actual expected values)
	expectedOutput := `XResolution: "72/1"
InteroperabilityIFDPointer: 10400
WhiteBalance: 0
ResolutionUnit: 2
ColorSpace: 1
SubSecTime: "00"
ExifIFDPointer: 360
GPSInfoIFDPointer: 10628
PixelXDimension: 6000
CustomRendered: 0
UserComment: ""
Copyright: ""
MakerNote: ""
DateTimeDigitized: "2024:04:10 01:22:57"
ShutterSpeedValue: "393216/65536"
Orientation: 1
YResolution: "72/1"
Model: "Canon EOS 80D"
FocalPlaneYResolution: "4000000/594"
FocalLength: "50/1"
SceneCaptureType: 0
ExposureProgram: 2
ApertureValue: "303104/65536"
ExifVersion: "0230"
FlashpixVersion: "0100"
DateTimeOriginal: "2024:04:10 01:22:57"
FocalPlaneResolutionUnit: 2
YCbCrPositioning: 2
Make: "Canon"
ISOSpeedRatings: 1600
LensModel: "EF-S18-135mm f/3.5-5.6 IS USM"
PixelYDimension: 4000
ExposureTime: "1/60"
ExposureMode: 0
SubSecTimeDigitized: "00"
MeteringMode: 5
DateTime: "2024:04:10 01:22:57"
Artist: ""
ComponentsConfiguration: ""
Flash: 9
FNumber: "5/1"
ExposureBiasValue: "0/1"
FocalPlaneXResolution: "6000000/921"
SubSecTimeOriginal: "00"
`

	// Capture the output of the extractJPEGMetadata function
	output := captureOutput(func() {
		ExtractJPEGMetadata(sampleFilePath)
	})

	// Parse the outputs into maps
	expectedMap := parseExifOutput(expectedOutput)
	outputMap := parseExifOutput(output)

	// Compare the maps
	if len(expectedMap) != len(outputMap) {
		t.Errorf("Expected %d EXIF fields, but got %d", len(expectedMap), len(outputMap))
	}

	for key, expectedValue := range expectedMap {
		if outputValue, exists := outputMap[key]; !exists {
			t.Errorf("Missing EXIF field: %s", key)
		} else if outputValue != expectedValue {
			t.Errorf("For EXIF field %s, expected value %s, but got %s", key, expectedValue, outputValue)
		}
	}
}
