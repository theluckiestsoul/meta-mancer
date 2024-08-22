package main

import (
	"fmt"
	"log"
	"mime"
	"os"
	"path/filepath"

	"github.com/rwcarlsen/goexif/exif"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path")
		return
	}
	filePath := os.Args[1]
	fmt.Printf("Processing file: %s\n", filePath)

	fileType := getFileType(filePath)
	fmt.Printf("File type: %s\n", fileType)

	switch fileType {
	case "image/jpeg":
		extractJPEGMetadata(filePath)
	default:
		fmt.Println("Unsupported file type")
	}
}

func getFileType(filePath string) string {
	ext := filepath.Ext(filePath)
	mimeType := mime.TypeByExtension(ext)
	return mimeType
}

// Exif(Exchangeable image file format) is a standard that specifies the formats for images, sound, and ancillary tags used by digital cameras (including smartphones), scanners and other systems handling image and sound files recorded by digital cameras.
func extractJPEGMetadata(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	// Map of EXIF properties and their field names
	exifFields := map[string]exif.FieldName{
		"ImageWidth":                 exif.ImageWidth,
		"ImageLength":                exif.ImageLength,
		"BitsPerSample":              exif.BitsPerSample,
		"Compression":                exif.Compression,
		"PhotometricInterpretation":  exif.PhotometricInterpretation,
		"Orientation":                exif.Orientation,
		"SamplesPerPixel":            exif.SamplesPerPixel,
		"PlanarConfiguration":        exif.PlanarConfiguration,
		"YCbCrSubSampling":           exif.YCbCrSubSampling,
		"YCbCrPositioning":           exif.YCbCrPositioning,
		"XResolution":                exif.XResolution,
		"YResolution":                exif.YResolution,
		"ResolutionUnit":             exif.ResolutionUnit,
		"DateTime":                   exif.DateTime,
		"ImageDescription":           exif.ImageDescription,
		"Make":                       exif.Make,
		"Model":                      exif.Model,
		"Software":                   exif.Software,
		"Artist":                     exif.Artist,
		"Copyright":                  exif.Copyright,
		"ExifIFDPointer":             exif.ExifIFDPointer,
		"GPSInfoIFDPointer":          exif.GPSInfoIFDPointer,
		"InteroperabilityIFDPointer": exif.InteroperabilityIFDPointer,
		"ExifVersion":                exif.ExifVersion,
		"FlashpixVersion":            exif.FlashpixVersion,
		"ColorSpace":                 exif.ColorSpace,
		"ComponentsConfiguration":    exif.ComponentsConfiguration,
		"CompressedBitsPerPixel":     exif.CompressedBitsPerPixel,
		"PixelXDimension":            exif.PixelXDimension,
		"PixelYDimension":            exif.PixelYDimension,
		"MakerNote":                  exif.MakerNote,
		"UserComment":                exif.UserComment,
		"RelatedSoundFile":           exif.RelatedSoundFile,
		"DateTimeOriginal":           exif.DateTimeOriginal,
		"DateTimeDigitized":          exif.DateTimeDigitized,
		"SubSecTime":                 exif.SubSecTime,
		"SubSecTimeOriginal":         exif.SubSecTimeOriginal,
		"SubSecTimeDigitized":        exif.SubSecTimeDigitized,
		"ImageUniqueID":              exif.ImageUniqueID,
		"ExposureTime":               exif.ExposureTime,
		"FNumber":                    exif.FNumber,
		"ExposureProgram":            exif.ExposureProgram,
		"SpectralSensitivity":        exif.SpectralSensitivity,
		"ISOSpeedRatings":            exif.ISOSpeedRatings,
		"OECF":                       exif.OECF,
		"ShutterSpeedValue":          exif.ShutterSpeedValue,
		"ApertureValue":              exif.ApertureValue,
		"BrightnessValue":            exif.BrightnessValue,
		"ExposureBiasValue":          exif.ExposureBiasValue,
		"MaxApertureValue":           exif.MaxApertureValue,
		"SubjectDistance":            exif.SubjectDistance,
		"MeteringMode":               exif.MeteringMode,
		"LightSource":                exif.LightSource,
		"Flash":                      exif.Flash,
		"FocalLength":                exif.FocalLength,
		"SubjectArea":                exif.SubjectArea,
		"FlashEnergy":                exif.FlashEnergy,
		"SpatialFrequencyResponse":   exif.SpatialFrequencyResponse,
		"FocalPlaneXResolution":      exif.FocalPlaneXResolution,
		"FocalPlaneYResolution":      exif.FocalPlaneYResolution,
		"FocalPlaneResolutionUnit":   exif.FocalPlaneResolutionUnit,
		"SubjectLocation":            exif.SubjectLocation,
		"ExposureIndex":              exif.ExposureIndex,
		"SensingMethod":              exif.SensingMethod,
		"FileSource":                 exif.FileSource,
		"SceneType":                  exif.SceneType,
		"CFAPattern":                 exif.CFAPattern,
		"CustomRendered":             exif.CustomRendered,
		"ExposureMode":               exif.ExposureMode,
		"WhiteBalance":               exif.WhiteBalance,
		"DigitalZoomRatio":           exif.DigitalZoomRatio,
		"FocalLengthIn35mmFilm":      exif.FocalLengthIn35mmFilm,
		"SceneCaptureType":           exif.SceneCaptureType,
		"GainControl":                exif.GainControl,
		"Contrast":                   exif.Contrast,
		"Saturation":                 exif.Saturation,
		"Sharpness":                  exif.Sharpness,
		"DeviceSettingDescription":   exif.DeviceSettingDescription,
		"SubjectDistanceRange":       exif.SubjectDistanceRange,
		"LensMake":                   exif.LensMake,
		"LensModel":                  exif.LensModel,
	}

	// Iterate over the map and extract each property
	for name, field := range exifFields {
		value, err := x.Get(field)
		if err == nil {
			fmt.Printf("%s: %s\n", name, value.String())
		}
	}

}
