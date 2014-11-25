package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func bash_buildpack_bash() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xac, 0x55,
		0x51, 0x6f, 0xe2, 0x46, 0x10, 0x7e, 0x36, 0xbf, 0x62, 0xe4, 0x43, 0xb9,
		0xa4, 0x92, 0x43, 0xe9, 0xf5, 0x29, 0x17, 0x4e, 0x4d, 0x02, 0xd7, 0xa2,
		0x26, 0x21, 0x25, 0xf0, 0x94, 0x46, 0x68, 0x59, 0x8f, 0x61, 0x95, 0x65,
		0xd7, 0xdd, 0x5d, 0x87, 0x43, 0x69, 0xfe, 0x7b, 0x67, 0x6d, 0x83, 0xed,
		0x40, 0x7a, 0x2f, 0x27, 0x21, 0x61, 0xef, 0xec, 0xcc, 0x7c, 0x33, 0xf3,
		0xcd, 0xe7, 0x56, 0x8c, 0x5c, 0x32, 0x83, 0x60, 0x51, 0x22, 0x77, 0x18,
		0xcf, 0x52, 0xe6, 0x96, 0xfb, 0xa7, 0x8a, 0xad, 0xb0, 0xd5, 0x9a, 0x67,
		0x42, 0xc6, 0x29, 0xe3, 0x4f, 0x51, 0xfe, 0x74, 0x7c, 0x02, 0x2f, 0xad,
		0x60, 0x7b, 0x37, 0x46, 0xcb, 0x7b, 0xe1, 0xa5, 0x37, 0x00, 0x53, 0xc0,
		0xd2, 0x54, 0x0a, 0xce, 0x9c, 0xd0, 0x0a, 0x32, 0x2b, 0xd4, 0x02, 0x84,
		0xb2, 0x8e, 0x49, 0x89, 0x31, 0xec, 0xe2, 0xd8, 0xb0, 0x15, 0xa0, 0xb2,
		0x99, 0xc1, 0xc8, 0xe7, 0xb5, 0xad, 0xa0, 0x4a, 0x61, 0xd1, 0x65, 0x29,
		0x7c, 0x81, 0x4e, 0x8c, 0xcf, 0x1d, 0x95, 0x49, 0xd9, 0x34, 0x7a, 0x60,
		0xf0, 0x2f, 0x05, 0x8d, 0x51, 0xb9, 0xba, 0x89, 0xeb, 0x55, 0x2a, 0x24,
		0xd6, 0x6c, 0xa9, 0xd1, 0x3c, 0xa1, 0xa3, 0xc8, 0x6d, 0x52, 0xb4, 0x95,
		0xe1, 0xb5, 0x5e, 0x51, 0x89, 0xee, 0x50, 0x4d, 0xc3, 0xc2, 0x54, 0xc1,
		0x86, 0xc4, 0xe8, 0x15, 0xfc, 0x2e, 0x1c, 0x4c, 0xc7, 0xd7, 0x54, 0x6c,
		0x0c, 0x3a, 0xf5, 0x85, 0x32, 0x09, 0x94, 0x7d, 0x25, 0x9c, 0x13, 0x76,
		0x19, 0x56, 0x61, 0x32, 0x23, 0x7b, 0x61, 0xbb, 0x1b, 0x96, 0x56, 0x7a,
		0xfe, 0x25, 0x04, 0xdf, 0x52, 0x7a, 0x7a, 0xf9, 0x74, 0x16, 0xb5, 0x8f,
		0xe7, 0xcc, 0xa2, 0x3f, 0x80, 0x76, 0xf7, 0xe4, 0x75, 0xaf, 0x29, 0x52,
		0x73, 0x0a, 0xed, 0x98, 0x59, 0xa0, 0xcb, 0x07, 0x44, 0x7e, 0x3b, 0x30,
		0xf9, 0x41, 0xa7, 0xed, 0xbd, 0xc9, 0x51, 0x24, 0xf0, 0xf0, 0x00, 0x91,
		0x82, 0xb0, 0x5d, 0x24, 0x0b, 0xe1, 0xf1, 0xf1, 0x33, 0xb8, 0x25, 0xaa,
		0x56, 0x10, 0x2c, 0x08, 0x32, 0x97, 0x5a, 0x21, 0x99, 0x09, 0x54, 0x48,
		0x7f, 0xb5, 0xa8, 0xe4, 0x1e, 0xf0, 0x78, 0xff, 0x2c, 0xf7, 0x5a, 0x22,
		0x7f, 0xd2, 0x99, 0x83, 0x28, 0xfa, 0x27, 0x13, 0xe8, 0xaa, 0xf8, 0x85,
		0x53, 0xd4, 0x9c, 0x14, 0x4a, 0x8b, 0x8d, 0x7c, 0x51, 0x14, 0x63, 0x4a,
		0xc0, 0xbb, 0xef, 0x65, 0x4e, 0x44, 0x2b, 0x30, 0x2b, 0x88, 0x4c, 0xd2,
		0x34, 0x75, 0x4e, 0x17, 0x3e, 0x49, 0x63, 0x56, 0x52, 0x58, 0x77, 0x68,
		0x50, 0xd7, 0x74, 0xfe, 0x1e, 0xcd, 0xa4, 0x85, 0xc8, 0x27, 0x6f, 0xf6,
		0xed, 0x4d, 0xe0, 0x9c, 0x73, 0x45, 0xe4, 0x0f, 0x70, 0xb9, 0x1b, 0x37,
		0x7e, 0x4b, 0x89, 0x6d, 0x39, 0x97, 0x69, 0x1a, 0xf4, 0xa6, 0x8d, 0x83,
		0x8b, 0xbb, 0xbb, 0x59, 0x7f, 0x38, 0xa6, 0x51, 0x10, 0xd5, 0xb7, 0x55,
		0x94, 0xb6, 0x3f, 0x46, 0x37, 0x83, 0x83, 0x86, 0xf1, 0xe0, 0xaf, 0xe9,
		0xe0, 0x7e, 0x32, 0x1b, 0xf6, 0x7b, 0x61, 0x9e, 0x36, 0x6a, 0x8f, 0x2f,
		0x6e, 0xfb, 0xa3, 0x9b, 0xea, 0xca, 0xfd, 0xe4, 0xe2, 0xea, 0xcf, 0x5e,
		0xc8, 0x31, 0x66, 0x26, 0xea, 0xfe, 0x4a, 0x06, 0x9e, 0x52, 0x5b, 0xa0,
		0x0a, 0xd7, 0x39, 0x0d, 0xb7, 0x85, 0x6c, 0xc3, 0x7b, 0xbc, 0x7d, 0xa3,
		0xd3, 0x94, 0xca, 0x4e, 0x8d, 0x78, 0x26, 0xb6, 0x2f, 0x90, 0xb0, 0x66,
		0x16, 0xcd, 0x4a, 0xd3, 0x78, 0xa2, 0xa5, 0xf6, 0xec, 0xf2, 0xc0, 0x40,
		0xe9, 0xb9, 0x8e, 0x37, 0x14, 0x77, 0xa9, 0xd7, 0x0a, 0xa2, 0x71, 0x79,
		0x70, 0xa6, 0xf4, 0xc2, 0x68, 0xda, 0xb9, 0xbf, 0x69, 0x72, 0x35, 0xf0,
		0xe5, 0x7b, 0x2d, 0x5f, 0x79, 0xc2, 0x19, 0xb1, 0xa2, 0x44, 0xe0, 0x01,
		0x4c, 0x2d, 0x26, 0x99, 0x24, 0xd1, 0x20, 0xfe, 0xab, 0x85, 0x85, 0x0e,
		0x24, 0xc8, 0x1c, 0xf1, 0xb8, 0x6a, 0xda, 0x15, 0x2d, 0xcc, 0xec, 0x6a,
		0x74, 0x7b, 0x3b, 0xb8, 0x9a, 0xcc, 0x26, 0xc3, 0x9b, 0xc1, 0x68, 0x3a,
		0xe9, 0x85, 0x9f, 0x7e, 0x2e, 0x02, 0xe4, 0x1d, 0xb7, 0x0e, 0x53, 0x98,
		0x53, 0xdb, 0xd7, 0xcc, 0xc4, 0xd6, 0xaf, 0x0c, 0x65, 0x10, 0x73, 0x21,
		0x85, 0xdb, 0xec, 0xf8, 0x9d, 0x34, 0xda, 0x81, 0xea, 0xb9, 0x41, 0x73,
		0xab, 0x33, 0xc3, 0x71, 0xef, 0x4a, 0x4e, 0xb2, 0x37, 0xf3, 0xf6, 0x32,
		0x52, 0x0c, 0xbc, 0xb6, 0x39, 0x97, 0xd3, 0xe1, 0x75, 0xff, 0x8e, 0xc6,
		0x30, 0x23, 0xb8, 0x8d, 0xc8, 0x4e, 0x38, 0x92, 0x96, 0xf0, 0x2b, 0x3a,
		0xbe, 0xf4, 0x9a, 0xc6, 0x33, 0xeb, 0x48, 0x0a, 0x76, 0x11, 0xfd, 0x36,
		0x78, 0x00, 0x75, 0x35, 0xdd, 0x5f, 0xd6, 0xc2, 0xcb, 0xdf, 0xdd, 0x31,
		0xbe, 0xe1, 0xe1, 0xbb, 0x11, 0x0c, 0xbf, 0xde, 0xf7, 0x3e, 0x7e, 0xf8,
		0x08, 0x06, 0x59, 0xec, 0x15, 0xa4, 0x14, 0x0f, 0x38, 0x3f, 0x3f, 0xdf,
		0x83, 0x48, 0xd7, 0xf7, 0x94, 0xac, 0x5a, 0xb4, 0xad, 0x10, 0x94, 0x60,
		0x8f, 0xea, 0x9b, 0x5a, 0x07, 0x5b, 0x2a, 0xd2, 0x71, 0xa6, 0x76, 0x14,
		0x8a, 0xa1, 0x89, 0xac, 0x33, 0x17, 0x8a, 0x9c, 0x9d, 0x17, 0xdf, 0x1a,
		0x23, 0x4e, 0xc2, 0xdd, 0xca, 0x17, 0x5a, 0x55, 0xad, 0x5e, 0xef, 0xf8,
		0x6d, 0xf5, 0x3f, 0x9d, 0xd0, 0xb5, 0x44, 0x9b, 0x9a, 0x9e, 0x0a, 0xdf,
		0xf5, 0x97, 0xca, 0xe9, 0xe1, 0xb7, 0xc7, 0xd7, 0xf0, 0x33, 0xc4, 0x9a,
		0x6e, 0x7e, 0x0f, 0xdf, 0xce, 0xeb, 0x5d, 0x6c, 0x39, 0x5d, 0x83, 0xe0,
		0xe8, 0x08, 0xde, 0x9d, 0x4b, 0xed, 0xce, 0x9c, 0x1a, 0xfe, 0x44, 0x2f,
		0x31, 0xc9, 0x56, 0x21, 0x4b, 0x35, 0x66, 0x34, 0xc7, 0x74, 0x80, 0x19,
		0xed, 0x06, 0x5a, 0xff, 0x15, 0x84, 0x02, 0x12, 0xc6, 0x55, 0x8f, 0xca,
		0xbb, 0x53, 0xc5, 0xe6, 0xf4, 0xef, 0x74, 0x89, 0x0b, 0x58, 0x93, 0x48,
		0xf8, 0x8d, 0xc6, 0xdd, 0xdd, 0x67, 0x6d, 0xf9, 0x85, 0xdb, 0xea, 0xd4,
		0x64, 0xd4, 0x1f, 0x9d, 0x81, 0x43, 0xaf, 0x7d, 0x09, 0xc1, 0x11, 0x16,
		0xe8, 0xa7, 0x90, 0xa3, 0xb5, 0xcc, 0x6c, 0x80, 0x3a, 0x4d, 0xe7, 0x6b,
		0x04, 0xee, 0x3f, 0xcb, 0x72, 0xcd, 0x36, 0x16, 0x52, 0x66, 0x2d, 0xb4,
		0x69, 0x29, 0xf2, 0x42, 0x60, 0x2d, 0xdc, 0xd2, 0x8b, 0xbb, 0xb0, 0x36,
		0xc3, 0x6d, 0xc1, 0xd4, 0x69, 0x2f, 0x9a, 0x17, 0xd5, 0xbd, 0x93, 0x46,
		0xc5, 0x8d, 0x31, 0x84, 0x07, 0x78, 0x52, 0xe2, 0x7c, 0x23, 0x56, 0xd0,
		0x10, 0x0e, 0x7a, 0xdb, 0x46, 0xaf, 0xfa, 0xf3, 0x43, 0x22, 0x17, 0xb3,
		0xfb, 0x6e, 0x28, 0x43, 0x27, 0xf4, 0xf1, 0xfd, 0x7f, 0x90, 0x5f, 0x1a,
		0xd6, 0xce, 0xe9, 0xd6, 0xa9, 0xf5, 0xfa, 0x5f, 0x00, 0x00, 0x00, 0xff,
		0xff, 0x40, 0xb8, 0x14, 0xca, 0x41, 0x09, 0x00, 0x00,
	},
		"bash/buildpack.bash",
	)
}

func bash_cmd_bash() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x9c, 0x54,
		0x41, 0x6f, 0x9c, 0x3c, 0x10, 0x3d, 0xc3, 0xaf, 0x98, 0xcf, 0x71, 0xa2,
		0xe4, 0x80, 0xf8, 0x58, 0xf5, 0x44, 0xb4, 0x55, 0xa2, 0xb6, 0xb7, 0xb6,
		0x97, 0x1c, 0xc3, 0x46, 0x72, 0xc0, 0x04, 0x14, 0xaf, 0x59, 0x61, 0xd8,
		0xa6, 0x22, 0xfc, 0xf7, 0xce, 0x18, 0x9b, 0x40, 0xb2, 0xbd, 0xf4, 0x04,
		0x9e, 0x19, 0xcf, 0x7b, 0xf3, 0xfc, 0xec, 0xb0, 0x90, 0xb9, 0x12, 0xad,
		0x84, 0xe8, 0x16, 0xbe, 0xfc, 0xf8, 0x7a, 0x17, 0x86, 0xf9, 0xbe, 0x88,
		0x54, 0x6d, 0xba, 0xcb, 0x2b, 0x18, 0xc2, 0xc0, 0xa7, 0x0b, 0x69, 0xf2,
		0x2d, 0xfb, 0x8e, 0x71, 0x03, 0xe2, 0x28, 0x6a, 0x25, 0x1e, 0x95, 0x84,
		0xbc, 0xd9, 0xef, 0x85, 0x2e, 0x0c, 0x7b, 0x2b, 0xd4, 0x66, 0xcb, 0x78,
		0x82, 0x01, 0xdf, 0x27, 0x7a, 0x96, 0xbf, 0x0d, 0x30, 0xae, 0x0d, 0x83,
		0x57, 0x30, 0xb2, 0x00, 0x66, 0x62, 0x5c, 0xa5, 0x71, 0xcc, 0xc2, 0xf1,
		0x0d, 0xcf, 0xd6, 0xad, 0x41, 0xe7, 0x5e, 0x65, 0xd3, 0xc2, 0x33, 0xd4,
		0x1a, 0xdb, 0x0c, 0xff, 0x11, 0xcd, 0xfb, 0x9b, 0xdd, 0xc8, 0xae, 0xa1,
		0x68, 0xc2, 0x20, 0x90, 0x79, 0xd5, 0x60, 0xe2, 0x99, 0x48, 0x34, 0x5a,
		0x22, 0xc8, 0x53, 0x2b, 0x0f, 0xc0, 0x1e, 0x08, 0xc4, 0x62, 0x36, 0x6d,
		0xb7, 0x42, 0xd2, 0x0e, 0xe7, 0x5f, 0xda, 0x46, 0x47, 0x48, 0xdf, 0xf5,
		0x94, 0x2f, 0x07, 0x5c, 0x9d, 0xd2, 0xeb, 0x9b, 0xcd, 0xa0, 0x62, 0x50,
		0xf6, 0x3a, 0xef, 0xea, 0x46, 0x83, 0xa0, 0x95, 0xd3, 0x6d, 0x21, 0x5b,
		0xa9, 0xed, 0xa8, 0x98, 0xc6, 0xef, 0xb0, 0x49, 0x23, 0x9e, 0x8c, 0x98,
		0x56, 0x4d, 0x2e, 0x94, 0xd5, 0xc1, 0xa9, 0xa0, 0x89, 0x2e, 0xbf, 0x5c,
		0x8c, 0x72, 0xb5, 0xe6, 0x5b, 0x6a, 0xb6, 0x54, 0x20, 0x62, 0x70, 0xf1,
		0x19, 0xe2, 0x42, 0x1e, 0x63, 0xdd, 0x2b, 0x05, 0x17, 0x17, 0x93, 0xaa,
		0xda, 0x8d, 0x15, 0x06, 0x76, 0x6e, 0x3a, 0x9e, 0x94, 0x0f, 0xc2, 0xc4,
		0x67, 0xf8, 0x17, 0xc5, 0x23, 0xdb, 0x6d, 0x6d, 0xaf, 0xf5, 0x8c, 0xb3,
		0x72, 0xef, 0x4e, 0xc8, 0x8d, 0xcb, 0x37, 0xd8, 0x55, 0x1e, 0x91, 0x31,
		0x06, 0xa9, 0x10, 0x90, 0xd7, 0x4a, 0x91, 0x8c, 0x71, 0xfa, 0x66, 0x58,
		0x67, 0x0d, 0xa2, 0x0d, 0xf0, 0x04, 0x32, 0x96, 0xf1, 0x9b, 0x0c, 0x85,
		0x0f, 0x83, 0xd1, 0x39, 0x67, 0xc2, 0x83, 0xe9, 0xf8, 0x1d, 0xc5, 0xc4,
		0x92, 0x4a, 0x66, 0x4e, 0x9e, 0xcc, 0x9b, 0x48, 0x98, 0xbc, 0x06, 0x53,
		0xd5, 0x65, 0x07, 0x3e, 0x8c, 0x85, 0xab, 0xf8, 0xeb, 0x2b, 0x74, 0x6d,
		0x2f, 0x7d, 0xda, 0x74, 0xa2, 0xeb, 0xcd, 0xf6, 0xff, 0x30, 0xa8, 0x4b,
		0xf0, 0xaa, 0xce, 0x6e, 0xb5, 0x32, 0x3e, 0x70, 0x8c, 0x67, 0x7c, 0xa5,
		0xe3, 0x35, 0x74, 0x95, 0xd4, 0x38, 0x04, 0x1f, 0x16, 0xfa, 0x61, 0x1d,
		0xdb, 0x8d, 0xb8, 0xfb, 0x86, 0x74, 0x50, 0x06, 0x51, 0xa8, 0xed, 0xfd,
		0x3d, 0x86, 0x28, 0x07, 0xbb, 0xdd, 0xbc, 0xd1, 0x9d, 0xd7, 0xcf, 0x06,
		0x4c, 0x9f, 0x57, 0xde, 0x11, 0x29, 0xd8, 0x42, 0xca, 0x3b, 0x66, 0x1b,
		0x3a, 0x5a, 0xe5, 0xbb, 0x10, 0xad, 0x8f, 0x4d, 0xf8, 0x65, 0xa9, 0x23,
		0xd2, 0x75, 0xaa, 0xb8, 0xa2, 0xfd, 0x65, 0xed, 0x3c, 0x31, 0x5b, 0xe3,
		0xf6, 0xc3, 0xb5, 0x4d, 0x6d, 0x21, 0xba, 0x0a, 0x31, 0xd7, 0xbe, 0x72,
		0x8d, 0x9c, 0xb5, 0x82, 0x43, 0x5b, 0xeb, 0xae, 0x04, 0x06, 0x70, 0x1e,
		0x6d, 0x3e, 0x19, 0x38, 0x37, 0x19, 0xda, 0xcc, 0x0d, 0xb5, 0x82, 0xff,
		0x28, 0xc7, 0x44, 0xc7, 0xc2, 0x98, 0xfe, 0xf1, 0x14, 0x12, 0xd5, 0xcd,
		0x50, 0x0b, 0xac, 0xf7, 0x68, 0xd3, 0xf6, 0xbf, 0x00, 0x62, 0x26, 0xf5,
		0x15, 0x33, 0xe8, 0xe4, 0x70, 0xff, 0xf1, 0x6a, 0xbc, 0xd4, 0x1d, 0xf0,
		0x49, 0xde, 0x90, 0x74, 0x72, 0x76, 0xaa, 0xa4, 0x3a, 0x9c, 0xba, 0xc4,
		0x77, 0x55, 0xf3, 0xcb, 0x00, 0x65, 0x91, 0x39, 0x8e, 0xb1, 0x17, 0xf6,
		0x22, 0xd3, 0x40, 0xa7, 0x6e, 0xb2, 0x68, 0x9f, 0xc8, 0x89, 0x64, 0x01,
		0x7f, 0x6a, 0x14, 0x5a, 0x9c, 0x1b, 0x4d, 0xe6, 0x65, 0xa7, 0xb1, 0xad,
		0x21, 0xe9, 0x07, 0x5f, 0x99, 0x33, 0x50, 0x02, 0x45, 0xc1, 0x1d, 0x53,
		0xd9, 0xc2, 0xdb, 0x03, 0xb5, 0x89, 0xcf, 0x69, 0xd2, 0x98, 0x5e, 0xa9,
		0x29, 0xa8, 0x31, 0x04, 0xeb, 0x35, 0xc4, 0x11, 0x5e, 0xa3, 0xc5, 0x76,
		0xfb, 0xb8, 0x9c, 0x38, 0x19, 0xc7, 0x44, 0x47, 0x34, 0x97, 0x7b, 0x3b,
		0x92, 0xd9, 0xb9, 0xee, 0x7e, 0xda, 0x97, 0x87, 0x44, 0x5a, 0x5c, 0x4b,
		0x2f, 0x97, 0x55, 0x25, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x23,
		0x50, 0x2b, 0x44, 0x06, 0x00, 0x00,
	},
		"bash/cmd.bash",
	)
}

func bash_fn_bash() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8c, 0x51,
		0x41, 0x6e, 0xc2, 0x30, 0x10, 0x3c, 0xc7, 0xaf, 0x18, 0xad, 0x22, 0x01,
		0xaa, 0xa2, 0x08, 0xae, 0x34, 0x3d, 0x56, 0xea, 0x1b, 0x28, 0x07, 0xcb,
		0xac, 0x89, 0xd5, 0xd4, 0x89, 0x6c, 0x03, 0xaa, 0x28, 0x7f, 0xef, 0x1a,
		0xd2, 0x92, 0x43, 0x55, 0x55, 0xb9, 0x64, 0x77, 0x66, 0x67, 0x76, 0xd6,
		0xca, 0xfa, 0x4a, 0x87, 0x7d, 0x9c, 0x2f, 0x70, 0x56, 0xc5, 0x8e, 0x4d,
		0xa7, 0x03, 0x63, 0xc7, 0xd1, 0x34, 0xf4, 0xe2, 0xe3, 0xc0, 0x26, 0x41,
		0xc3, 0x1e, 0xbc, 0x49, 0xae, 0xf7, 0xb3, 0x08, 0x21, 0x1f, 0xde, 0xd9,
		0xa7, 0x48, 0xaa, 0xe8, 0x7a, 0xa3, 0xbb, 0xdc, 0xe9, 0x9c, 0xe7, 0xa6,
		0x9c, 0xa7, 0x8f, 0x81, 0x51, 0x2e, 0xf1, 0x89, 0x7d, 0xe0, 0x01, 0xdf,
		0x6a, 0x63, 0x59, 0x1d, 0x41, 0x53, 0x03, 0x12, 0xa0, 0x65, 0xbd, 0x43,
		0xb5, 0x5c, 0xa8, 0x82, 0x4d, 0xdb, 0xa3, 0x62, 0x50, 0x79, 0x1e, 0x05,
		0xeb, 0x1a, 0x35, 0xbd, 0x7a, 0xba, 0x64, 0xa2, 0x3e, 0xbd, 0xa1, 0x7a,
		0x6e, 0x30, 0xab, 0x9b, 0xfa, 0x3c, 0x04, 0xe7, 0x13, 0xe8, 0x91, 0xca,
		0x25, 0x3d, 0xd1, 0x65, 0x26, 0x78, 0x0a, 0xc8, 0x5c, 0xc8, 0xa7, 0x2e,
		0x2a, 0xa7, 0xca, 0x16, 0xff, 0x4e, 0x95, 0xa1, 0xe0, 0x86, 0x5c, 0x51,
		0x1e, 0xc8, 0x44, 0xf9, 0xe1, 0xa3, 0xe4, 0xa3, 0x5f, 0x82, 0x45, 0x33,
		0x59, 0x9e, 0xd6, 0xb8, 0x6e, 0x5f, 0xe6, 0xfe, 0xe8, 0xee, 0xbc, 0xed,
		0xff, 0x70, 0x8f, 0x13, 0x7b, 0xba, 0x73, 0xac, 0x6f, 0x72, 0x26, 0xc4,
		0xb6, 0x3f, 0xc5, 0xfe, 0x10, 0x0c, 0x4b, 0xbd, 0xa2, 0xf1, 0x3a, 0x54,
		0x5a, 0x8f, 0x72, 0x3e, 0xbe, 0x18, 0xa4, 0x5a, 0xfc, 0x40, 0xb8, 0x01,
		0xd7, 0xcd, 0x26, 0x80, 0x2a, 0x9c, 0xc5, 0x66, 0x23, 0xa3, 0x77, 0x49,
		0xc2, 0x76, 0xbb, 0x46, 0x6a, 0xd9, 0xab, 0xa2, 0xb8, 0x25, 0x13, 0x5d,
		0x39, 0xa1, 0x76, 0x1d, 0x2a, 0x8f, 0x87, 0x95, 0xf4, 0x6f, 0xc3, 0xd6,
		0x49, 0x9c, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xaf, 0x12, 0xe0,
		0x23, 0x02, 0x00, 0x00,
	},
		"bash/fn.bash",
	)
}

func bash_herokuish_bash() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x74, 0x54,
		0xef, 0x6f, 0xdb, 0x36, 0x10, 0xfd, 0x2c, 0xfd, 0x15, 0x57, 0xce, 0x6d,
		0xd6, 0x00, 0x5a, 0xb0, 0xae, 0x9f, 0x6a, 0x68, 0xa8, 0x93, 0xb9, 0x4b,
		0xb1, 0x2e, 0x35, 0x9a, 0xb6, 0xc3, 0x90, 0x14, 0x01, 0x23, 0x9e, 0x65,
		0x22, 0x12, 0x49, 0x90, 0x54, 0x9a, 0xa0, 0xc8, 0xff, 0xde, 0x23, 0xf5,
		0xc3, 0x92, 0x9d, 0xe8, 0x93, 0x78, 0xef, 0xdd, 0xe3, 0xf1, 0xdd, 0x91,
		0xa9, 0x5c, 0xc3, 0xc5, 0x05, 0xb0, 0xd9, 0xf1, 0xe2, 0xfc, 0xf4, 0xea,
		0xeb, 0xf2, 0xd3, 0xf9, 0xfb, 0xb3, 0x77, 0x1f, 0x19, 0x64, 0x95, 0x07,
		0xf6, 0x9a, 0xc1, 0xb7, 0x6f, 0x73, 0xf0, 0x1b, 0x54, 0x69, 0x82, 0xc5,
		0x46, 0x03, 0x7b, 0xf6, 0x0c, 0xfe, 0xd7, 0x8d, 0x05, 0x77, 0xef, 0x3c,
		0xd6, 0x70, 0xcc, 0xdd, 0x06, 0xa4, 0x03, 0xdd, 0x78, 0xd0, 0x6b, 0x10,
		0xdc, 0xe3, 0x1b, 0x18, 0x69, 0x7d, 0x3c, 0x63, 0xa3, 0xcc, 0x55, 0x85,
		0xdc, 0x21, 0x34, 0xa6, 0xb4, 0x5c, 0x20, 0x78, 0xdd, 0xe6, 0xbf, 0x06,
		0x6d, 0xa1, 0xb4, 0x48, 0xc9, 0xf6, 0xb7, 0xc0, 0xbf, 0x93, 0x1e, 0x5e,
		0xa5, 0x6b, 0x99, 0xa6, 0x14, 0x14, 0x5a, 0x55, 0xf7, 0xc0, 0x8d, 0xb9,
		0x32, 0xdc, 0x6f, 0x72, 0x36, 0xfb, 0xb1, 0x58, 0xad, 0xae, 0x56, 0x8b,
		0xcf, 0xa7, 0x6f, 0xb2, 0x23, 0x0a, 0x3f, 0xb0, 0x2d, 0x0b, 0xd5, 0xed,
		0xc0, 0x5a, 0x9e, 0x7d, 0xed, 0x59, 0xbe, 0x36, 0x47, 0x04, 0x8d, 0x99,
		0xd7, 0x8d, 0xac, 0xc4, 0xc0, 0x3d, 0xfe, 0xf2, 0xfe, 0xc3, 0x5f, 0x63,
		0x76, 0x84, 0xc7, 0xfc, 0x82, 0x17, 0x1b, 0x1c, 0xf8, 0x27, 0x8b, 0x93,
		0xd3, 0xe5, 0x98, 0x1f, 0xe1, 0x3d, 0x7d, 0xc3, 0x8b, 0x9b, 0xe9, 0x1e,
		0xab, 0xc5, 0xc9, 0x3f, 0x7b, 0xfb, 0x04, 0x9a, 0xa3, 0xe4, 0x54, 0x60,
		0x51, 0x71, 0x4b, 0x06, 0x29, 0x63, 0xe5, 0xad, 0xac, 0xb0, 0x44, 0x71,
		0xd5, 0x38, 0xb4, 0x39, 0x53, 0xfa, 0x5a, 0x8b, 0x7b, 0xf6, 0x38, 0xa5,
		0xb4, 0xba, 0x31, 0x81, 0x13, 0x7f, 0x48, 0x08, 0x95, 0x6b, 0x2c, 0x66,
		0x61, 0x6b, 0xf7, 0xeb, 0x4b, 0xf8, 0x91, 0x26, 0xf5, 0x8d, 0x90, 0x16,
		0x32, 0x03, 0x97, 0x69, 0x92, 0xb0, 0x59, 0x6f, 0x27, 0xeb, 0xd6, 0xbd,
		0x71, 0xfd, 0x7a, 0x6b, 0x4f, 0x1f, 0xd9, 0x1a, 0x30, 0xe1, 0x0c, 0x47,
		0x64, 0xe9, 0x43, 0x9a, 0x8e, 0x36, 0xec, 0x0b, 0x15, 0xe8, 0x8a, 0x9c,
		0x9d, 0x6f, 0xf4, 0x77, 0x07, 0x01, 0x06, 0x87, 0xde, 0x4b, 0x55, 0x3a,
		0x6a, 0x34, 0x1d, 0x41, 0xf9, 0x35, 0xb0, 0xe7, 0xd9, 0x1f, 0xaf, 0x1c,
		0xfc, 0x02, 0xcf, 0xdd, 0xa5, 0xea, 0xd4, 0xfb, 0x1e, 0xe7, 0xa3, 0x52,
		0x43, 0xd8, 0x98, 0x4a, 0x16, 0xdc, 0x4b, 0xad, 0x5a, 0x35, 0xd1, 0x90,
		0x46, 0x09, 0xb6, 0x51, 0x5e, 0xd6, 0xd8, 0x25, 0xf7, 0xad, 0xcf, 0x47,
		0xe7, 0xa2, 0xf0, 0x2a, 0x24, 0xd0, 0xd0, 0xad, 0xc9, 0x36, 0x07, 0x6b,
		0x1a, 0x3a, 0x81, 0x6b, 0xa9, 0x42, 0xfe, 0x75, 0x98, 0x4b, 0x22, 0x4b,
		0xab, 0x55, 0x8d, 0xca, 0x77, 0x42, 0xdb, 0xb9, 0xc8, 0x27, 0x96, 0x24,
		0xec, 0x3f, 0x6d, 0x6f, 0x42, 0x22, 0xb9, 0x8a, 0x85, 0xd7, 0xf6, 0xbe,
		0xaf, 0x24, 0xd2, 0x5c, 0x97, 0xbf, 0x9d, 0x93, 0x7c, 0x62, 0x20, 0x29,
		0xf7, 0xe6, 0xb5, 0x93, 0x05, 0x95, 0x6e, 0x4f, 0x35, 0xde, 0x78, 0x18,
		0x96, 0x7c, 0xd7, 0x6b, 0x18, 0xce, 0x22, 0x95, 0xf3, 0xbc, 0xaa, 0x50,
		0x6c, 0x27, 0xce, 0xc5, 0x56, 0xdc, 0xa2, 0x75, 0xa4, 0xf7, 0x54, 0x33,
		0xa0, 0xc3, 0x81, 0x2b, 0x01, 0xae, 0x31, 0x46, 0x5b, 0x4f, 0x22, 0x7d,
		0x54, 0xaa, 0xb5, 0x1e, 0x6e, 0xae, 0xc0, 0xdb, 0x28, 0xe9, 0xa5, 0xaf,
		0xb0, 0x15, 0x8c, 0xc0, 0xec, 0xe0, 0x12, 0x2f, 0x7e, 0xff, 0x3b, 0x0b,
		0xdf, 0x9f, 0x07, 0x30, 0x3b, 0x0c, 0x24, 0xa9, 0x04, 0xf9, 0xd7, 0xb2,
		0xbe, 0x6f, 0xc8, 0x69, 0x08, 0x77, 0x02, 0x2a, 0xa9, 0x70, 0x0e, 0x42,
		0xd3, 0xe1, 0xfa, 0x17, 0x27, 0x84, 0x18, 0xe4, 0x39, 0x64, 0xd9, 0xe1,
		0xe8, 0x9d, 0x49, 0x26, 0xe2, 0x07, 0x91, 0x46, 0x51, 0xac, 0x1c, 0xee,
		0x82, 0x10, 0xbf, 0x83, 0x5e, 0x8b, 0x60, 0x7a, 0x33, 0x12, 0xba, 0x80,
		0x18, 0x2a, 0x19, 0x5f, 0x91, 0xb6, 0x1e, 0x1a, 0xbd, 0x46, 0x8a, 0x52,
		0x0a, 0xca, 0xd8, 0xbb, 0x63, 0x0c, 0x66, 0x6f, 0x43, 0x5a, 0xcd, 0xa5,
		0x1a, 0xe8, 0x90, 0xa1, 0x06, 0x23, 0x0d, 0xae, 0xb9, 0xac, 0xe6, 0x6d,
		0xdd, 0x9f, 0x3f, 0x2d, 0x4e, 0x96, 0xe1, 0x65, 0x84, 0x17, 0x2f, 0x20,
		0x72, 0xee, 0xd2, 0x34, 0x29, 0x6a, 0x91, 0xe1, 0x5d, 0xb0, 0x31, 0x0e,
		0xa6, 0x9b, 0x44, 0x3a, 0x5b, 0x27, 0xb4, 0x4c, 0xb9, 0x6d, 0xcf, 0x80,
		0x7d, 0xa1, 0x01, 0x0c, 0xbd, 0xe8, 0x1a, 0x3a, 0x69, 0xe7, 0x58, 0x6a,
		0x88, 0x67, 0xf1, 0xef, 0x09, 0xac, 0x53, 0x79, 0x02, 0xad, 0xa4, 0xf3,
		0x7b, 0xb5, 0xb8, 0xaa, 0x29, 0x81, 0xfd, 0xcb, 0x15, 0x2f, 0x31, 0x3c,
		0xb7, 0xc3, 0x3d, 0x0b, 0xc0, 0x4e, 0x11, 0x21, 0x94, 0xc9, 0x3a, 0xfc,
		0xef, 0xc7, 0x4b, 0x54, 0x68, 0xe9, 0x2d, 0xdf, 0x47, 0xda, 0xff, 0xbd,
		0x9d, 0x8d, 0xd5, 0x45, 0xb8, 0x93, 0xad, 0x09, 0xab, 0x6e, 0xe5, 0xa2,
		0x1d, 0x74, 0xb1, 0x43, 0x31, 0x50, 0xe8, 0xba, 0xa6, 0xf5, 0x4e, 0x1d,
		0x7d, 0x66, 0x46, 0xa7, 0xdd, 0x29, 0x65, 0x80, 0xf0, 0x0e, 0x8b, 0xc7,
		0x11, 0xc3, 0x6d, 0x98, 0x29, 0x02, 0xc3, 0xed, 0x67, 0xb3, 0xf3, 0xe5,
		0x87, 0x77, 0x8c, 0x1a, 0x00, 0x34, 0x48, 0x47, 0x51, 0xf1, 0x65, 0x92,
		0x4c, 0xb7, 0x20, 0xd6, 0x5b, 0x36, 0x9f, 0x07, 0x42, 0xd0, 0x1d, 0xe3,
		0x61, 0x3d, 0x82, 0xa3, 0xdb, 0x84, 0xef, 0xf4, 0x2b, 0x82, 0x87, 0x14,
		0x8f, 0x15, 0xd1, 0xe1, 0x19, 0x6b, 0x93, 0xe8, 0x4e, 0x39, 0x5e, 0xa4,
		0x0f, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xc3, 0xcc, 0x4c, 0x91,
		0x07, 0x00, 0x00,
	},
		"bash/herokuish.bash",
	)
}

func bash_procfile_bash() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x74, 0x94,
		0xdd, 0x4e, 0xe3, 0x3a, 0x10, 0xc7, 0xaf, 0x9b, 0xa7, 0x18, 0x45, 0x11,
		0xb4, 0xe7, 0xe0, 0x06, 0x6e, 0x0f, 0xea, 0xd1, 0xae, 0xc4, 0x6a, 0xb9,
		0x01, 0x56, 0xdc, 0xb2, 0xa8, 0x72, 0xe3, 0x69, 0x62, 0xd5, 0x89, 0xb3,
		0xb6, 0xc3, 0xc7, 0xb2, 0xbc, 0xfb, 0xfa, 0x23, 0x69, 0x63, 0x1a, 0xae,
		0x12, 0x8f, 0x67, 0xec, 0xdf, 0xfc, 0x67, 0x3c, 0x49, 0xab, 0x64, 0xb1,
		0xe5, 0x02, 0x49, 0x4b, 0x95, 0xc6, 0xf9, 0x02, 0xde, 0x92, 0x19, 0xc3,
		0x42, 0x50, 0x85, 0xc0, 0x50, 0x17, 0xab, 0xf4, 0x3b, 0x1a, 0x28, 0x64,
		0x5d, 0xd3, 0x86, 0x81, 0x36, 0x8a, 0x37, 0x25, 0x6c, 0xa5, 0x02, 0x0a,
		0x2e, 0x14, 0xb5, 0x06, 0xf3, 0xda, 0x22, 0x6c, 0x95, 0xac, 0xe1, 0x47,
		0x7f, 0x58, 0x7a, 0x38, 0xc3, 0x6d, 0xae, 0xd2, 0xec, 0xc2, 0x9a, 0x0a,
		0x6a, 0x20, 0xcd, 0x68, 0xdb, 0xae, 0x5b, 0x6a, 0xaa, 0x7c, 0xef, 0x0c,
		0x7f, 0xe0, 0x95, 0xd6, 0x82, 0x94, 0xe8, 0xf6, 0x5d, 0x40, 0x9a, 0xbc,
		0x27, 0x07, 0x32, 0x6d, 0xa8, 0x32, 0x53, 0x64, 0xf7, 0x5d, 0x13, 0x43,
		0x0c, 0x98, 0x11, 0x0c, 0x98, 0x4a, 0xc9, 0xae, 0xac, 0x00, 0x5f, 0xb0,
		0x98, 0x26, 0xdb, 0x5f, 0xe5, 0x5c, 0x20, 0x9b, 0xc7, 0xa2, 0x0c, 0x50,
		0x8b, 0x88, 0xca, 0xb9, 0x7e, 0x06, 0x45, 0x35, 0x28, 0xcb, 0x21, 0xeb,
		0x33, 0xe8, 0x9a, 0x56, 0xf1, 0x27, 0x1b, 0x50, 0x22, 0x83, 0x4e, 0xa3,
		0x82, 0x67, 0x6e, 0x2a, 0xb8, 0x46, 0x25, 0x77, 0x1d, 0x11, 0x7c, 0x87,
		0x80, 0xcd, 0xd3, 0x98, 0x21, 0x44, 0xf2, 0xdf, 0x48, 0x9c, 0xfb, 0x68,
		0x43, 0xa3, 0x21, 0x95, 0xac, 0x71, 0x64, 0x12, 0x92, 0x32, 0x62, 0xe3,
		0x3f, 0x9a, 0xec, 0xca, 0x2d, 0x92, 0x59, 0x74, 0x7d, 0xf6, 0x25, 0xca,
		0xc0, 0x25, 0xa5, 0x43, 0x0a, 0x86, 0x1b, 0x2b, 0x54, 0x7a, 0xc5, 0x75,
		0x21, 0x9f, 0xd0, 0xd7, 0x78, 0x2c, 0xac, 0xb6, 0x7c, 0x7c, 0x0b, 0x0f,
		0x0f, 0x40, 0xb6, 0x56, 0x8d, 0x4d, 0xc7, 0x05, 0xfb, 0x58, 0xc4, 0xc7,
		0xc7, 0x4b, 0xab, 0x34, 0x36, 0xc9, 0x6c, 0x26, 0x64, 0x41, 0x45, 0x08,
		0xb4, 0x2b, 0xff, 0xb5, 0x42, 0xcf, 0x5d, 0xfd, 0xa7, 0x62, 0x87, 0xfa,
		0xef, 0xf0, 0x55, 0xdb, 0xff, 0x17, 0xaa, 0x4a, 0x0d, 0x58, 0x54, 0x72,
		0x61, 0xaf, 0x9d, 0xb9, 0x1f, 0x48, 0xf7, 0xbe, 0xbd, 0xd8, 0x3d, 0x17,
		0x90, 0xff, 0x21, 0x7b, 0xf3, 0xbf, 0x79, 0x0e, 0xf9, 0x19, 0xbc, 0xbb,
		0x10, 0x85, 0xa6, 0x53, 0x16, 0x64, 0xcb, 0xf7, 0xd8, 0x3a, 0xc6, 0x5e,
		0x2a, 0x14, 0x48, 0xf5, 0x14, 0x36, 0xc3, 0x2d, 0xed, 0x84, 0x59, 0x0f,
		0xf8, 0xd1, 0x7a, 0x2a, 0x8d, 0xe1, 0xac, 0x28, 0x8d, 0x21, 0xaa, 0x57,
		0x31, 0x44, 0x1f, 0x27, 0x67, 0xd1, 0xd2, 0x2c, 0xba, 0xc1, 0x11, 0xc1,
		0xc9, 0x09, 0xfc, 0xb4, 0xbb, 0x7d, 0xee, 0x57, 0x61, 0x3f, 0xae, 0x88,
		0x7f, 0x83, 0x99, 0xb6, 0x57, 0x17, 0x06, 0xd9, 0xba, 0xa1, 0x35, 0x06,
		0x31, 0xa2, 0xd3, 0x3e, 0x11, 0x25, 0x9c, 0x7b, 0x2b, 0x8f, 0x8e, 0xec,
		0x1a, 0x16, 0xbf, 0xbd, 0xa1, 0xc1, 0x42, 0x9b, 0xf4, 0x62, 0x32, 0x0b,
		0x6d, 0x6d, 0x3e, 0xfd, 0x48, 0x41, 0xc7, 0x84, 0xc0, 0x1b, 0xfb, 0x84,
		0x84, 0x86, 0xbd, 0xcf, 0xe2, 0x12, 0x98, 0xf4, 0xf9, 0xbc, 0xb4, 0x52,
		0xb9, 0x47, 0x8e, 0xab, 0x5e, 0xc7, 0xc1, 0x25, 0xcf, 0xd0, 0x0b, 0xc2,
		0x64, 0x83, 0x9e, 0xf1, 0x08, 0xa2, 0x6f, 0xe9, 0x00, 0xa2, 0x2b, 0xd9,
		0x1a, 0x57, 0xd5, 0xa6, 0x13, 0xa2, 0x14, 0x72, 0x93, 0xcc, 0xea, 0x1d,
		0xe3, 0x0a, 0x48, 0x3b, 0x1e, 0x31, 0xcb, 0x3e, 0x68, 0x69, 0xb3, 0xf2,
		0x70, 0xbe, 0x85, 0x1c, 0xdf, 0x84, 0x4b, 0xfe, 0xcf, 0x52, 0x57, 0x3d,
		0xa9, 0x96, 0x9d, 0x2a, 0xdc, 0xc3, 0x1f, 0xa6, 0x99, 0xc7, 0xaa, 0xa8,
		0xae, 0x80, 0xa8, 0x78, 0x38, 0xf5, 0x8f, 0x32, 0x70, 0xf5, 0x09, 0x5e,
		0xdf, 0xdd, 0x7c, 0x5b, 0x1d, 0x38, 0xec, 0x01, 0xee, 0x29, 0xd7, 0x92,
		0x01, 0xf1, 0xce, 0x23, 0xc6, 0xd4, 0xfe, 0x8f, 0x1f, 0xe9, 0xda, 0x79,
		0xba, 0x61, 0x59, 0xc9, 0xe7, 0x06, 0xc8, 0xfd, 0xd4, 0xf6, 0x7f, 0xb1,
		0xa9, 0xb4, 0xe3, 0xad, 0x4d, 0xc7, 0x67, 0x46, 0x84, 0xf1, 0x3c, 0x09,
		0x9c, 0xa1, 0xd5, 0xdd, 0x9a, 0x33, 0xd7, 0xd3, 0xf3, 0xfb, 0xaf, 0xb7,
		0x57, 0x77, 0x37, 0xff, 0x5e, 0x9c, 0x9f, 0x9f, 0x2f, 0x5c, 0x21, 0x0e,
		0x0e, 0xae, 0xb3, 0x56, 0x69, 0x97, 0xbd, 0x05, 0x6f, 0xdb, 0x4b, 0xc9,
		0x8c, 0x32, 0xe6, 0x2f, 0xb5, 0xe9, 0xfc, 0xea, 0xb8, 0x1d, 0xdb, 0x84,
		0x94, 0xdc, 0x75, 0x45, 0xf0, 0x49, 0xfb, 0x3f, 0x17, 0x9a, 0x7a, 0x6f,
		0x3f, 0xf8, 0x5c, 0x4b, 0x13, 0xa2, 0x2b, 0x14, 0x02, 0xf2, 0x0d, 0x6f,
		0xf2, 0x8d, 0xd3, 0x33, 0x58, 0x19, 0xd7, 0x74, 0x23, 0xd0, 0x96, 0x99,
		0x6a, 0xfd, 0x2c, 0x15, 0xeb, 0xed, 0xb6, 0x68, 0x05, 0x92, 0x0d, 0x65,
		0xbe, 0xc1, 0x83, 0xad, 0x91, 0xa4, 0x50, 0x48, 0x0d, 0x06, 0x2d, 0x83,
		0xb1, 0x8b, 0xee, 0x0f, 0xb6, 0x72, 0xca, 0x86, 0x85, 0xd4, 0x70, 0x7a,
		0xda, 0x2f, 0x03, 0x7e, 0xf8, 0x3f, 0xaa, 0x8c, 0x33, 0x47, 0x99, 0xc4,
		0xf3, 0xd4, 0xd7, 0x62, 0x15, 0x39, 0x1c, 0xd7, 0x25, 0xda, 0x7f, 0xff,
		0x1b, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x05, 0xaf, 0xe4, 0x6f, 0x07, 0x00,
		0x00,
	},
		"bash/procfile.bash",
	)
}

func bash_slug_bash() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x7c, 0x53,
		0x41, 0x8b, 0xdb, 0x3c, 0x10, 0x3d, 0xc7, 0xbf, 0x62, 0x3e, 0x11, 0xd8,
		0xe4, 0xa0, 0xe4, 0x4b, 0xaf, 0xc5, 0x85, 0x92, 0x96, 0xb2, 0x50, 0xda,
		0xd2, 0x4d, 0xa0, 0xb0, 0x5d, 0x82, 0x56, 0x9e, 0xd8, 0x02, 0xd9, 0x12,
		0x92, 0xdc, 0x66, 0xdd, 0xf6, 0xbf, 0x77, 0x64, 0x3b, 0x89, 0xdd, 0xdd,
		0x24, 0xa7, 0x78, 0x34, 0xf3, 0xe6, 0xcd, 0x9b, 0x79, 0x89, 0x43, 0x91,
		0x99, 0x4a, 0x3f, 0x81, 0xd7, 0x75, 0xbe, 0xb3, 0x22, 0x14, 0x29, 0x5b,
		0x86, 0xd2, 0x2e, 0xe3, 0xf7, 0x22, 0xe4, 0x0d, 0x4b, 0x92, 0xf8, 0x97,
		0xab, 0xd2, 0x1a, 0x17, 0x66, 0x73, 0xf8, 0x95, 0x4c, 0x32, 0x94, 0x5a,
		0x38, 0x84, 0x0c, 0xbd, 0x4c, 0xd9, 0x6d, 0xfb, 0x02, 0x02, 0xf2, 0x46,
		0x59, 0x8b, 0x59, 0x0b, 0x05, 0x41, 0xb8, 0x47, 0xa1, 0x35, 0xec, 0x9d,
		0x29, 0x61, 0xfb, 0xf5, 0x23, 0x18, 0x07, 0x77, 0x9b, 0x77, 0xb7, 0x9f,
		0x80, 0x9d, 0x01, 0x6a, 0xa7, 0x53, 0x36, 0x5d, 0x51, 0x44, 0xed, 0xe1,
		0xfe, 0x1e, 0xd8, 0x74, 0xa6, 0x3d, 0xf0, 0xb7, 0x30, 0x15, 0xd6, 0xb6,
		0x6c, 0xe6, 0x0c, 0x1e, 0x1e, 0x5e, 0x43, 0x28, 0xb0, 0x4a, 0x26, 0x13,
		0x87, 0xa1, 0x76, 0x15, 0xac, 0x92, 0x09, 0xea, 0x63, 0x05, 0x61, 0x8c,
		0x72, 0x24, 0x05, 0x80, 0x13, 0x8a, 0x81, 0x65, 0x86, 0x3f, 0x96, 0x55,
		0x4d, 0x2c, 0x38, 0xa7, 0x52, 0xf7, 0x04, 0xaf, 0xa8, 0xc0, 0x3b, 0xc9,
		0xe0, 0x77, 0x24, 0x08, 0xfc, 0xd0, 0xac, 0x29, 0x72, 0x6c, 0xc6, 0x22,
		0xae, 0xc7, 0x08, 0x22, 0xc2, 0xc5, 0x94, 0xbd, 0x4a, 0xfe, 0xf4, 0xa2,
		0xe4, 0x58, 0xa1, 0x13, 0x01, 0x5f, 0x92, 0xe5, 0x43, 0xff, 0x76, 0x55,
		0x18, 0xe2, 0x0c, 0x44, 0xd8, 0x61, 0x45, 0x02, 0x5a, 0x4b, 0xe8, 0xda,
		0x48, 0xa1, 0xc1, 0xaa, 0xbc, 0xd9, 0x19, 0x1b, 0x94, 0xa9, 0x5a, 0x6d,
		0x7e, 0x16, 0x4a, 0x16, 0x6d, 0x14, 0xde, 0x9c, 0xc7, 0x3a, 0xcd, 0x3c,
		0x48, 0x4f, 0x19, 0xe7, 0xb5, 0x47, 0x2e, 0x4d, 0x69, 0x1d, 0x7a, 0xcf,
		0xad, 0x33, 0xb9, 0x13, 0x65, 0x1a, 0x73, 0x3a, 0xf6, 0x7d, 0x8f, 0xc8,
		0x46, 0xe5, 0x95, 0x71, 0x38, 0xec, 0x44, 0x9a, 0xf2, 0x3d, 0x0d, 0xfc,
		0x58, 0x2b, 0x9d, 0xb5, 0x23, 0x2f, 0x17, 0xe7, 0xcc, 0x91, 0xd2, 0xcf,
		0x00, 0xa8, 0xf7, 0x37, 0xb8, 0x54, 0xd9, 0xb5, 0x8e, 0x92, 0x4e, 0x07,
		0x74, 0x61, 0xfa, 0x0c, 0x05, 0xbe, 0x13, 0x36, 0xe7, 0x78, 0x90, 0xba,
		0xce, 0x30, 0xbd, 0x59, 0xe4, 0x2a, 0xdc, 0x74, 0xc1, 0xf5, 0x88, 0x18,
		0xeb, 0x82, 0x32, 0xd2, 0x3d, 0x9d, 0x6f, 0x0c, 0x02, 0xfd, 0x26, 0x8b,
		0xe1, 0x9c, 0x3b, 0xaf, 0x1a, 0xa4, 0x53, 0x9b, 0x65, 0x35, 0xf0, 0xbb,
		0x02, 0xce, 0xf9, 0xb4, 0x66, 0x59, 0x07, 0x9a, 0x79, 0x35, 0x27, 0x8a,
		0x41, 0x05, 0x8d, 0xc0, 0xd6, 0x24, 0x9e, 0xd2, 0xc7, 0x8d, 0xc5, 0x5a,
		0x50, 0xbe, 0x2f, 0x8a, 0x5f, 0xec, 0x74, 0x01, 0x78, 0xb8, 0x64, 0x8b,
		0xf7, 0xed, 0x0b, 0x1c, 0x4f, 0xe4, 0x9f, 0xed, 0x07, 0xd3, 0x9a, 0x62,
		0xf6, 0x65, 0xbb, 0x99, 0xf7, 0xd6, 0xf8, 0xbc, 0xdd, 0x5c, 0xb1, 0xc6,
		0x7f, 0xdd, 0x5a, 0x06, 0x73, 0xbe, 0xe8, 0x8b, 0xa8, 0xf1, 0x75, 0x67,
		0xfc, 0x7f, 0xc5, 0x1c, 0xb4, 0x3e, 0x22, 0x04, 0x7c, 0x33, 0xee, 0xd4,
		0x41, 0x8d, 0xcc, 0x31, 0x7c, 0xef, 0x2c, 0xf1, 0x37, 0x00, 0x00, 0xff,
		0xff, 0x93, 0x63, 0xfe, 0x2d, 0x4b, 0x04, 0x00, 0x00,
	},
		"bash/slug.bash",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"bash/buildpack.bash": bash_buildpack_bash,
	"bash/cmd.bash": bash_cmd_bash,
	"bash/fn.bash": bash_fn_bash,
	"bash/herokuish.bash": bash_herokuish_bash,
	"bash/procfile.bash": bash_procfile_bash,
	"bash/slug.bash": bash_slug_bash,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"bash": &_bintree_t{nil, map[string]*_bintree_t{
		"buildpack.bash": &_bintree_t{bash_buildpack_bash, map[string]*_bintree_t{
		}},
		"cmd.bash": &_bintree_t{bash_cmd_bash, map[string]*_bintree_t{
		}},
		"fn.bash": &_bintree_t{bash_fn_bash, map[string]*_bintree_t{
		}},
		"herokuish.bash": &_bintree_t{bash_herokuish_bash, map[string]*_bintree_t{
		}},
		"procfile.bash": &_bintree_t{bash_procfile_bash, map[string]*_bintree_t{
		}},
		"slug.bash": &_bintree_t{bash_slug_bash, map[string]*_bintree_t{
		}},
	}},
}}