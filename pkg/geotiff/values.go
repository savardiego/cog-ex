package geotiff

import (
	"fmt"
	"reflect"
)

// FieldTagNames contains value and name of all tags
var FieldTagNames map[uint16]string

// FieldTags is the dual of FieldTagNames
var FieldTags map[string]uint16

// ArrayFields is an array with some tags (?)
var ArrayFields []uint16

// FieldTypeNames define types inside tiff
var FieldTypeNames map[uint16]string

// FieldTypes is the dual of FieldTypeNames
var FieldTypes map[string]uint16

// PhotometricInterpretations values
var PhotometricInterpretations map[string]uint

// GeoKeyNames contains names of geographic keys
var GeoKeyNames map[uint]string

// GeoKeys is the dual of GeoKeyNames
var GeoKeys map[string]uint

func init() {
	FieldTagNames = map[uint16]string{
		// TIFF Baseline
		0x013B: "Artist",
		0x0102: "BitsPerSample",
		0x0109: "CellLength",
		0x0108: "CellWidth",
		0x0140: "ColorMap",
		0x0103: "Compression",
		0x8298: "Copyright",
		0x0132: "DateTime",
		0x0152: "ExtraSamples",
		0x010A: "FillOrder",
		0x0121: "FreeByteCounts",
		0x0120: "FreeOffsets",
		0x0123: "GrayResponseCurve",
		0x0122: "GrayResponseUnit",
		0x013C: "HostComputer",
		0x010E: "ImageDescription",
		0x0101: "ImageLength",
		0x0100: "ImageWidth",
		0x010F: "Make",
		0x0119: "MaxSampleValue",
		0x0118: "MinSampleValue",
		0x0110: "Model",
		0x00FE: "NewSubfileType",
		0x0112: "Orientation",
		0x0106: "PhotometricInterpretation",
		0x011C: "PlanarConfiguration",
		0x0128: "ResolutionUnit",
		0x0116: "RowsPerStrip",
		0x0115: "SamplesPerPixel",
		0x0131: "Software",
		0x0117: "StripByteCounts",
		0x0111: "StripOffsets",
		0x00FF: "SubfileType",
		0x0107: "Threshholding",
		0x011A: "XResolution",
		0x011B: "YResolution",

		// TIFF Extended
		0x0146: "BadFaxLines",
		0x0147: "CleanFaxData",
		0x0157: "ClipPath",
		0x0148: "ConsecutiveBadFaxLines",
		0x01B1: "Decode",
		0x01B2: "DefaultImageColor",
		0x010D: "DocumentName",
		0x0150: "DotRange",
		0x0141: "HalftoneHints",
		0x015A: "Indexed",
		0x015B: "JPEGTables",
		0x011D: "PageName",
		0x0129: "PageNumber",
		0x013D: "Predictor",
		0x013F: "PrimaryChromaticities",
		0x0214: "ReferenceBlackWhite",
		0x0153: "SampleFormat",
		0x0154: "SMinSampleValue",
		0x0155: "SMaxSampleValue",
		0x022F: "StripRowCounts",
		0x014A: "SubIFDs",
		0x0124: "T4Options",
		0x0125: "T6Options",
		0x0145: "TileByteCounts",
		0x0143: "TileLength",
		0x0144: "TileOffsets",
		0x0142: "TileWidth",
		0x012D: "TransferFunction",
		0x013E: "WhitePoint",
		0x0158: "XClipPathUnits",
		0x011E: "XPosition",
		0x0211: "YCbCrCoefficients",
		0x0213: "YCbCrPositioning",
		0x0212: "YCbCrSubSampling",
		0x0159: "YClipPathUnits",
		0x011F: "YPosition",

		// EXIF
		0x9202: "ApertureValue",
		0xA001: "ColorSpace",
		0x9004: "DateTimeDigitized",
		0x9003: "DateTimeOriginal",
		0x8769: "Exif IFD",
		0x9000: "ExifVersion",
		0x829A: "ExposureTime",
		0xA300: "FileSource",
		0x9209: "Flash",
		0xA000: "FlashpixVersion",
		0x829D: "FNumber",
		0xA420: "ImageUniqueID",
		0x9208: "LightSource",
		0x927C: "MakerNote",
		0x9201: "ShutterSpeedValue",
		0x9286: "UserComment",

		// IPTC
		0x83BB: "IPTC",

		// ICC
		0x8773: "ICC Profile",

		// XMP
		0x02BC: "XMP",

		// GDAL
		0xA480: "GDAL_METADATA",
		0xA481: "GDAL_NODATA",

		// Photoshop
		0x8649: "Photoshop",

		// GeoTiff
		0x830E: "ModelPixelScale",
		0x8482: "ModelTiepoint",
		0x85D8: "ModelTransformation",
		0x87AF: "GeoKeyDirectory",
		0x87B0: "GeoDoubleParams",
		0x87B1: "GeoAsciiParams",
	}
	FieldTags = make(map[string]uint16)
	for k, v := range FieldTagNames {
		FieldTags[v] = k
	}

	ArrayFields = []uint16{
		FieldTags["BitsPerSample"],
		FieldTags["ExtraSamples"],
		FieldTags["SampleFormat"],
		FieldTags["StripByteCounts"],
		FieldTags["StripOffsets"],
		FieldTags["StripRowCounts"],
		FieldTags["TileByteCounts"],
		FieldTags["TileOffsets"],
	}

	FieldTypeNames = map[uint16]string{
		0x0001: "BYTE",
		0x0002: "ASCII",
		0x0003: "SHORT",
		0x0004: "LONG",
		0x0005: "RATIONAL",
		0x0006: "SBYTE",
		0x0007: "UNDEFINED",
		0x0008: "SSHORT",
		0x0009: "SLONG",
		0x000A: "SRATIONAL",
		0x000B: "FLOAT",
		0x000C: "DOUBLE",
		// introduced by BigTIFF
		0x0010: "LONG8",
		0x0011: "SLONG8",
		0x0012: "IFD8",
	}
	FieldTypes = make(map[string]uint16)
	for k, v := range FieldTypeNames {
		FieldTypes[v] = k
	}

	PhotometricInterpretations = map[string]uint{
		"WhiteIsZero":      0,
		"BlackIsZero":      1,
		"RGB":              2,
		"Palette":          3,
		"TransparencyMask": 4,
		"CMYK":             5,
		"YCbCr":            6,
		"CIELab":           8,
		"ICCLab":           9,
	}

	GeoKeyNames = map[uint]string{
		1024: "GTModelTypeGeoKey",
		1025: "GTRasterTypeGeoKey",
		1026: "GTCitationGeoKey",
		2048: "GeographicTypeGeoKey",
		2049: "GeogCitationGeoKey",
		2050: "GeogGeodeticDatumGeoKey",
		2051: "GeogPrimeMeridianGeoKey",
		2052: "GeogLinearUnitsGeoKey",
		2053: "GeogLinearUnitSizeGeoKey",
		2054: "GeogAngularUnitsGeoKey",
		2055: "GeogAngularUnitSizeGeoKey",
		2056: "GeogEllipsoidGeoKey",
		2057: "GeogSemiMajorAxisGeoKey",
		2058: "GeogSemiMinorAxisGeoKey",
		2059: "GeogInvFlatteningGeoKey",
		2060: "GeogAzimuthUnitsGeoKey",
		2061: "GeogPrimeMeridianLongGeoKey",
		2062: "GeogTOWGS84GeoKey",
		3072: "ProjectedCSTypeGeoKey",
		3073: "PCSCitationGeoKey",
		3074: "ProjectionGeoKey",
		3075: "ProjCoordTransGeoKey",
		3076: "ProjLinearUnitsGeoKey",
		3077: "ProjLinearUnitSizeGeoKey",
		3078: "ProjStdParallel1GeoKey",
		3079: "ProjStdParallel2GeoKey",
		3080: "ProjNatOriginLongGeoKey",
		3081: "ProjNatOriginLatGeoKey",
		3082: "ProjFalseEastingGeoKey",
		3083: "ProjFalseNorthingGeoKey",
		3084: "ProjFalseOriginLongGeoKey",
		3085: "ProjFalseOriginLatGeoKey",
		3086: "ProjFalseOriginEastingGeoKey",
		3087: "ProjFalseOriginNorthingGeoKey",
		3088: "ProjCenterLongGeoKey",
		3089: "ProjCenterLatGeoKey",
		3090: "ProjCenterEastingGeoKey",
		3091: "ProjCenterNorthingGeoKey",
		3092: "ProjScaleAtNatOriginGeoKey",
		3093: "ProjScaleAtCenterGeoKey",
		3094: "ProjAzimuthAngleGeoKey",
		3095: "ProjStraightVertPoleLongGeoKey",
		3096: "ProjRectifiedGridAngleGeoKey",
		4096: "VerticalCSTypeGeoKey",
		4097: "VerticalCitationGeoKey",
		4098: "VerticalDatumGeoKey",
		4099: "VerticalUnitsGeoKey",
	}
	GeoKeys = make(map[string]uint)
	for k, v := range GeoKeyNames {
		GeoKeys[v] = k
	}
}

// GetValues read values from slice
func GetValues(slice *DataSlice, fieldType uint16, count uint, offset uint) ([]interface{}, error) {
	var values []interface{}
	var readMethodName string
	fieldTypeLength, err := getFieldTypeLength(fieldType)
	if err != nil {
		return values, fmt.Errorf("cannot get field type length due to %v", err)
	}
	switch fieldType {
	case FieldTypes["BYTE"], FieldTypes["ASCII"], FieldTypes["UNDEFINED"]:
		values = make([]interface{}, count)
		readMethodName = "ReadUint8"
		break
	case FieldTypes["SBYTE"]:
		values = make([]interface{}, count)
		readMethodName = "ReadInt8"
		break
	case FieldTypes["SHORT"]:
		values = make([]interface{}, count)
		readMethodName = "ReadUint16"
		break
	case FieldTypes["SSHORT"]:
		values = make([]interface{}, count)
		readMethodName = "ReadInt16"
		break
	case FieldTypes["LONG"]:
		values = make([]interface{}, count)
		readMethodName = "ReadUint32"
		break
	case FieldTypes["SLONG"]:
		values = make([]interface{}, count)
		readMethodName = "ReadInt32"
		break
	case FieldTypes["LONG8"], FieldTypes["IFD8"]:
		values = make([]interface{}, count)
		readMethodName = "ReadUint64"
		break
	case FieldTypes["SLONG8"]:
		values = make([]interface{}, count)
		readMethodName = "ReadInt64"
		break
	case FieldTypes["RATIONAL"]:
		values = make([]interface{}, count)
		readMethodName = "ReadUint32"
		break
	default:
		values = make([]interface{}, count)
		readMethodName = "unknown"
	}

	// Normal Fields
	if fieldType != FieldTypes["RATIONAL"] && fieldType != FieldTypes["SRATIONAL"] {
		os := offset
		for i := uint(0); i < count; i++ {
			os += i * uint(fieldTypeLength)
			rv := reflect.ValueOf(slice).MethodByName(readMethodName).Call([]reflect.Value{reflect.ValueOf(os)})
			values[i] = rv[0].Interface()

		}
		// RAIONAL or SRATIONAL fields
	} else {
		os := offset
		for i := uint(0); i < count; i += 2 {
			os += i * uint(fieldTypeLength)
			rv := reflect.ValueOf(slice).MethodByName(readMethodName).Call([]reflect.Value{reflect.ValueOf(os)})
			values[i] = rv[0].Interface()
		}
	}
	return values, nil
}

func getFieldTypeLength(fieldType uint16) (uint8, error) {
	switch fieldType {
	case FieldTypes["BYTE"], FieldTypes["ASCII"], FieldTypes["SBYTE"], FieldTypes["UNDEFINED"]:
		return 1, nil
	case FieldTypes["SHORT"], FieldTypes["SSHORT"]:
		return 2, nil
	case FieldTypes["LONG"], FieldTypes["SLONG"], FieldTypes["FLOAT"]:
		return 4, nil
	case FieldTypes["LONG8"], FieldTypes["SLONG8"], FieldTypes["IFD8"], FieldTypes["RATIONAL"], FieldTypes["SRATIONAL"], FieldTypes["DOUBLE"]:
		return 8, nil
	default:
		return 0, fmt.Errorf("invalid field type: %d", fieldType)
	}
}
