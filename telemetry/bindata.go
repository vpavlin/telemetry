// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 000001_message_type.up.sql (66B)
// 000002_bandwidth_protocol.up.sql (719B)
// 000003_index_truncate.up.sql (598B)
// 000004_envelope.table.up.sql (531B)
// 000005_pushed_envelope.up.sql (574B)
// 000006_status_version.up.sql (198B)
// 000007_waku_push_filter.up.sql (523B)
// doc.go (73B)

package telemetry

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __000001_message_typeUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x4a\x4d\x4e\xcd\x2c\x4b\x4d\xf1\x4d\x2d\x2e\x4e\x4c\x4f\x2d\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\xc8\x85\x88\x85\x54\x16\xa4\x2a\x84\x39\x06\x39\x7b\x38\x06\x69\x18\x99\x9a\x6a\x5a\x73\x01\x02\x00\x00\xff\xff\xf4\x14\x08\x7c\x42\x00\x00\x00")

func _000001_message_typeUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_message_typeUpSql,
		"000001_message_type.up.sql",
	)
}

func _000001_message_typeUpSql() (*asset, error) {
	bytes, err := _000001_message_typeUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_message_type.up.sql", size: 66, mode: os.FileMode(0644), modTime: time.Unix(1716427081, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe2, 0x43, 0xcc, 0xef, 0xad, 0x5f, 0x44, 0x58, 0x8d, 0x47, 0x70, 0x5d, 0x23, 0x30, 0xe2, 0x1f, 0xdb, 0x4d, 0xad, 0x6e, 0xd9, 0xe7, 0x50, 0x19, 0x43, 0x1c, 0x37, 0x57, 0xea, 0xc6, 0x57, 0xab}}
	return a, nil
}

var __000002_bandwidth_protocolUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x90\x3f\x4f\xf3\x30\x10\xc6\xf7\x7c\x8a\x1b\x5b\x29\xd3\x2b\x75\xea\xe4\x37\x39\xc0\xc2\x71\x2a\xdb\x41\x74\x42\x56\x72\x82\x48\x6d\x0c\xf1\x85\x81\x4f\x8f\x92\x50\x10\x7f\x5a\x10\x03\x93\x25\x3f\x8f\xee\xee\xf7\xcb\x0c\x0a\x87\xe0\xc4\x7f\x85\x20\xcf\x40\x97\x0e\xf0\x5a\x5a\x67\xe1\xbe\x0f\x1c\xea\xb0\xb3\xec\x39\x1a\xcf\x04\x8b\x04\x00\xa0\x6d\xc0\xa2\x91\x42\xc1\xc6\xc8\x42\x98\x2d\x5c\xe2\x36\x9d\xa2\xbb\x10\x59\x36\x70\x25\x4c\x76\x21\xcc\xe2\xdf\x6a\xb5\x9c\x26\xea\x4a\xa9\xb9\x71\x18\xaa\xfd\x9e\x4e\xf5\x7a\xcf\x24\x3b\xc8\xcb\x6a\x3c\x6c\x63\x30\x93\x56\x96\xfa\x8b\x56\x39\xf0\x77\xb5\xba\x27\xcf\xd4\x08\x06\xa9\x1d\x9e\xa3\xf9\x98\x87\x2e\x72\xef\xdb\x8e\x3f\x43\xdf\x0c\x5d\xfb\x30\x10\xcc\xcf\x62\x26\x4c\xdf\x71\xa4\x6f\x0b\x96\xc9\x72\x9d\x24\x3f\x95\xea\x02\xfb\x5d\xfc\x4b\xad\x3c\x6e\x94\xdd\x11\x0f\x53\x3a\xfa\x3c\xa2\xe9\x55\x63\x2e\x1c\x9e\x56\x37\xa3\xfd\x46\x9e\x50\x0e\xcd\x8b\xbb\x9e\x6a\x6a\x1f\xa9\x29\x28\x46\x7f\x4b\x11\x44\x9e\x43\x56\xaa\xaa\xd0\xb0\x9f\xff\x6c\xfb\x44\x87\x7b\xd7\xc9\x73\x00\x00\x00\xff\xff\x88\xcb\xe5\x75\xcf\x02\x00\x00")

func _000002_bandwidth_protocolUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_bandwidth_protocolUpSql,
		"000002_bandwidth_protocol.up.sql",
	)
}

func _000002_bandwidth_protocolUpSql() (*asset, error) {
	bytes, err := _000002_bandwidth_protocolUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_bandwidth_protocol.up.sql", size: 719, mode: os.FileMode(0644), modTime: time.Unix(1716427081, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xfe, 0x83, 0x69, 0xab, 0x3e, 0xf5, 0x8d, 0x44, 0xb2, 0x6e, 0x52, 0x8d, 0x27, 0xe8, 0x95, 0x28, 0x3c, 0xea, 0x29, 0x93, 0x6d, 0xa3, 0x10, 0xde, 0x9b, 0xc8, 0xa6, 0xb9, 0x80, 0xa1, 0x3, 0x6f}}
	return a, nil
}

var __000003_index_truncateUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\xb1\x6a\x80\x30\x10\x86\xf7\x3c\xc5\x8d\x0a\x5d\x3a\x67\x4a\x35\x83\xa0\x11\xd2\x08\xdd\x24\xe8\x91\x0a\xb6\x91\xe4\x5a\xfa\xf8\xc5\x4a\x2b\x18\xb5\x6b\xfe\xef\xbe\xfb\xb9\x18\xdd\xa9\x42\x18\x09\x46\x3c\xd5\x12\x02\x0e\x38\x7d\xe2\xd8\x60\x8c\xd6\x61\xe4\xec\x00\x2c\xc1\x93\x1f\xfc\xfc\x4c\x96\xa2\xb6\x84\xf7\x84\xf1\x64\xe7\xd4\x72\x58\x23\x9c\x0b\xe8\x2c\xe1\xc8\x19\x2b\xb4\x5c\xc1\x4a\x95\xf2\x25\xe9\xd3\x0f\x01\x57\x4e\x10\xb4\x2a\x49\xb3\xbf\x34\xe7\xb7\x9e\x7d\x61\x1f\x3e\xde\x4f\x65\x3b\x92\xfd\x20\xf9\xb1\x5a\x72\x89\x7e\x1a\xbf\x1e\x57\x53\x92\x64\xbf\x2f\xca\xbe\xe1\x03\x5c\xb6\x3c\x39\xdd\xb9\x74\xcb\x6e\xb4\x4c\xd4\x46\xea\xab\x3f\x03\x2d\x95\x68\x24\x14\x6d\xdd\x35\x0a\x5e\x7d\xa4\xaa\x04\xd3\xc2\x82\x18\xaa\x92\x5f\x4f\x6f\x8b\xff\x9d\xff\x0e\x00\x00\xff\xff\x11\x42\xcb\x4c\x56\x02\x00\x00")

func _000003_index_truncateUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000003_index_truncateUpSql,
		"000003_index_truncate.up.sql",
	)
}

func _000003_index_truncateUpSql() (*asset, error) {
	bytes, err := _000003_index_truncateUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000003_index_truncate.up.sql", size: 598, mode: os.FileMode(0644), modTime: time.Unix(1716427081, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcf, 0x8, 0x4, 0x47, 0xc8, 0x65, 0x38, 0x79, 0x3e, 0x37, 0xec, 0x4e, 0x1a, 0x24, 0x50, 0x3c, 0x1c, 0x75, 0xe8, 0x3b, 0x2, 0x62, 0x2, 0x52, 0x50, 0xff, 0x4a, 0x8f, 0x9d, 0x71, 0x79, 0xf6}}
	return a, nil
}

var __000004_envelopeTableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\xb1\x6e\x83\x30\x10\x86\x77\x9e\xe2\x46\x90\x98\x2a\x65\xca\x74\x85\x6b\x63\xc5\x98\xca\x98\xaa\x99\x2a\x02\xa7\x14\xa9\xb1\xa9\x0d\x91\xfa\xf6\x95\x42\x5a\x55\xa9\x42\x27\x0f\xff\xe7\xdf\xe7\xef\x32\x4d\x68\x08\x0c\xde\x4b\x02\xf1\x00\xaa\x34\x40\x2f\xa2\x32\x15\x78\x6e\xb9\x3f\x71\x47\xf6\xc4\xef\x6e\xe0\x00\x71\x04\x00\xd0\x77\x50\x91\x16\x28\xe1\x49\x8b\x02\xf5\x0e\xb6\xb4\x4b\xcf\xd1\x91\x43\x68\x0e\xbc\x69\xc2\x1b\x3c\xa3\xce\x36\xa8\xe3\xbb\xd5\x2a\x39\xd7\xaa\x5a\xca\x19\x0b\x6c\x47\x1c\x41\x28\x43\x8f\xa4\xaf\xc2\xd6\x73\x33\x72\x77\x33\x1f\xdd\xd0\xb7\x4b\xed\xc3\xb4\xaf\xa6\xbd\xf9\x0f\xbb\x7c\xcf\x6f\xf9\xb3\x16\xf9\x12\x69\x5d\xc7\xaa\x39\xf2\xe2\xa3\xde\xb5\x1c\x42\x6f\x0f\xe4\xbd\xf3\x4b\x68\x56\xaa\xca\x68\x14\xca\xfc\x55\xfc\x3a\xd9\xfe\x63\x62\x98\x8f\x78\x16\x95\xfe\xf6\x9a\x5e\x0d\x9e\xfe\x8c\x97\x44\xc9\x3a\x8a\x50\x1a\xd2\x97\x7d\x7e\xd7\x17\xf3\xf5\x00\x98\xe7\x90\x95\xb2\x2e\xd4\x4d\x4b\xeb\xe8\x2b\x00\x00\xff\xff\x9d\x3f\xc2\xc6\x13\x02\x00\x00")

func _000004_envelopeTableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000004_envelopeTableUpSql,
		"000004_envelope.table.up.sql",
	)
}

func _000004_envelopeTableUpSql() (*asset, error) {
	bytes, err := _000004_envelopeTableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000004_envelope.table.up.sql", size: 531, mode: os.FileMode(0644), modTime: time.Unix(1716524216, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x32, 0xee, 0x49, 0xa0, 0x48, 0x2b, 0x8b, 0xe8, 0xd3, 0x6a, 0xae, 0x7f, 0x62, 0x65, 0x8a, 0x45, 0xbb, 0x8a, 0xee, 0xcd, 0x13, 0xde, 0xd6, 0x33, 0xe2, 0x3f, 0x32, 0xff, 0xfe, 0xf4, 0xda, 0xe7}}
	return a, nil
}

var __000005_pushed_envelopeUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\xcf\x6a\xf2\x40\x14\xc5\xf7\x79\x8a\xbb\x08\x24\x81\xac\x3e\x70\x95\xd5\x18\xef\xa7\x83\x71\x12\x26\x63\xab\x2b\xc9\x9f\x8b\x49\x89\x49\xea\xcc\x14\x7c\xfb\xa2\xa1\xa5\x16\xaa\xab\x81\x39\x87\xdf\x3d\x9c\xb3\x48\xc1\x75\x61\x8e\x4b\x2e\x1c\x00\x80\x58\x22\x53\x08\x6a\x9f\x21\x8c\xb6\xec\x5a\xdd\x1c\x4e\x64\x9a\xa1\x06\x96\x03\x8a\xed\x06\x7c\x2f\x69\x8f\x8d\xc9\xac\x6e\xbc\x10\x3c\x49\x5d\x71\xf1\x82\xc8\xc1\x5d\x8c\x99\xe2\xe9\x04\x7a\x5d\xa1\x80\xda\x8e\x5d\x5b\x15\x86\x0e\x43\xf9\x46\x95\x01\x75\xfd\xed\x6d\xd7\x45\x0e\x8a\x05\xb8\x6e\xe4\x38\x5f\x27\xd9\x3c\x41\xe0\xff\x41\xa4\x0a\x70\xc7\x73\x95\x83\xa6\xde\x60\xff\x41\xdd\x30\x92\x06\xff\x06\x6e\x6b\xc8\x51\x72\x96\x40\x26\xf9\x86\xc9\x3d\xac\x71\x1f\xde\xa4\x13\x69\x5d\x1c\x69\x55\xe8\x06\x5e\x98\x8c\x57\x4c\xfa\xff\x66\xb3\xe0\x86\x14\xdb\x24\x99\x6c\x57\x2a\x33\xc0\x85\xc2\x25\xca\x5f\x62\x75\xa6\xc2\x50\xfd\xa7\x6e\x86\xb1\xad\x1e\xd1\x47\x5b\xe6\xb6\x54\xcf\x6c\x9a\xfa\x9a\xce\x6b\xba\x6c\xf9\xe2\x91\xaf\x1f\x6a\x12\xc5\x89\x9e\x9c\xbc\x0e\xb5\x99\x76\xba\x9f\x6d\x72\xc4\xa9\xc8\x95\x64\x5c\xa8\xfb\x4e\x0f\xb6\x6f\xdf\x2d\xc1\xf4\xf8\x53\x33\xe1\xcf\x22\xc3\xbb\xa4\xe1\x77\x9e\xc0\x09\x22\xe7\x33\x00\x00\xff\xff\x3d\x18\x50\x60\x3e\x02\x00\x00")

func _000005_pushed_envelopeUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000005_pushed_envelopeUpSql,
		"000005_pushed_envelope.up.sql",
	)
}

func _000005_pushed_envelopeUpSql() (*asset, error) {
	bytes, err := _000005_pushed_envelopeUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000005_pushed_envelope.up.sql", size: 574, mode: os.FileMode(0644), modTime: time.Unix(1719028717, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7d, 0xaf, 0x8a, 0xcb, 0x97, 0x1e, 0xc6, 0xf6, 0x86, 0xe4, 0x1b, 0x67, 0x10, 0x87, 0x8e, 0x80, 0x1d, 0x5a, 0x7d, 0x64, 0xd0, 0x89, 0x3f, 0x1e, 0x6f, 0x93, 0x87, 0x4a, 0xd7, 0x87, 0xb8, 0x5e}}
	return a, nil
}

var __000006_status_versionUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x4a\x4d\x4e\xcd\x2c\x4b\x4d\xf1\x4d\x2d\x2e\x4e\x4c\x4f\x2d\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\x2e\x49\x2c\x29\x2d\x0e\x4b\x2d\x2a\xce\xcc\xcf\x53\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x36\xd4\xb4\xe6\xc2\x66\x84\x6b\x5e\x59\x6a\x4e\x7e\x01\x59\x66\x14\xa7\xe6\x95\x90\xa8\x1f\x10\x00\x00\xff\xff\xeb\x4e\x39\x66\xc6\x00\x00\x00")

func _000006_status_versionUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000006_status_versionUpSql,
		"000006_status_version.up.sql",
	)
}

func _000006_status_versionUpSql() (*asset, error) {
	bytes, err := _000006_status_versionUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000006_status_version.up.sql", size: 198, mode: os.FileMode(0644), modTime: time.Unix(1719028717, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2b, 0x11, 0xee, 0x9f, 0x4f, 0xf5, 0x0, 0x9a, 0x98, 0xe9, 0x44, 0x21, 0x2e, 0x57, 0xf7, 0xae, 0xf3, 0xb2, 0x3d, 0x94, 0x40, 0x69, 0xa7, 0x1d, 0x62, 0x57, 0x31, 0x9f, 0x60, 0x6, 0xed, 0x80}}
	return a, nil
}

var __000007_waku_push_filterUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\xc1\x4a\x03\x31\x10\x86\xcf\xdd\xa7\x98\xe3\x2e\xec\x49\xe8\xc9\x53\x5c\x52\x1b\x5c\xd3\x92\x44\xb1\x27\x49\x37\x03\x0d\x6e\xb3\xdb\x64\x42\x7d\x7c\xa1\x82\x76\xc5\xb6\xa7\x61\xf8\x3f\xf8\x87\x6f\x1a\xc5\x99\xe1\x60\xd8\x43\xcb\x41\x2c\x40\xae\x0c\xf0\x37\xa1\x8d\x86\xa3\xfd\xc8\xeb\x9c\x76\x0b\xdf\x13\x46\x28\x8b\x99\x77\xa0\xb9\x12\xac\x85\xb5\x12\xcf\x4c\x6d\xe0\x89\x6f\xea\x62\x76\xb4\x7d\x8f\xc4\x9c\x8b\x98\x12\xbc\x32\xd5\x2c\x99\x2a\xef\xe6\xf3\xaa\x2e\x66\x23\x62\x14\x4e\x63\x70\x18\x27\xd9\xa9\x4b\xbe\xb4\xed\x0f\xa4\x70\x1c\x22\x5d\xc1\x12\x1e\x32\x86\x0e\x97\x36\xed\x6e\x42\x66\x20\xdb\xdf\xa4\x44\x70\xf8\x79\x91\xea\x86\x40\x18\xc8\x0c\xa3\xef\x2e\x1f\x9f\xb7\x29\x6f\xaf\x33\xe4\xf7\x98\xc8\xee\x47\x10\xd2\xf0\x47\xae\x26\x2d\x11\x2d\xa1\x63\xf4\x4f\x58\x00\x00\x34\x2b\xa9\x8d\x62\x42\x9a\x3f\x5f\x79\xcf\xc1\x1f\x32\xc2\xf7\x28\xcf\x5d\xd7\x30\x95\x5a\xc3\xb9\xbd\xdf\xed\x24\xa0\x2a\xaa\xfb\xaf\x00\x00\x00\xff\xff\x48\xf0\x3d\x30\x0b\x02\x00\x00")

func _000007_waku_push_filterUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000007_waku_push_filterUpSql,
		"000007_waku_push_filter.up.sql",
	)
}

func _000007_waku_push_filterUpSql() (*asset, error) {
	bytes, err := _000007_waku_push_filterUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000007_waku_push_filter.up.sql", size: 523, mode: os.FileMode(0644), modTime: time.Unix(1719271502, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5d, 0xa, 0x2c, 0x93, 0xa, 0x1f, 0xeb, 0x49, 0x60, 0xe2, 0x8, 0x46, 0xb5, 0x16, 0xa4, 0xa9, 0x7f, 0xec, 0xfb, 0xe1, 0xdc, 0x12, 0x15, 0x17, 0x1, 0x28, 0xa3, 0xca, 0xeb, 0x45, 0x81, 0x31}}
	return a, nil
}

var _docGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xc9\x31\x12\x84\x20\x0c\x05\xd0\x9e\x53\xfc\x0b\x90\xf4\x7b\x9b\xac\xfe\xc9\x38\x20\x41\x4c\xe3\xed\x6d\xac\xdf\xb4\xad\x99\x13\xf7\xd5\x4b\x51\xf5\xf8\x39\x07\x97\x25\xe1\x51\xff\xc7\xd8\x2d\x0d\x75\x36\x47\xb2\xf3\x64\xae\x07\x35\x20\xa2\x1f\x8a\x07\x44\xcb\x1b\x00\x00\xff\xff\xb6\x03\x50\xe0\x49\x00\x00\x00")

func docGoBytes() ([]byte, error) {
	return bindataRead(
		_docGo,
		"doc.go",
	)
}

func docGo() (*asset, error) {
	bytes, err := docGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "doc.go", size: 73, mode: os.FileMode(0644), modTime: time.Unix(1716427081, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xae, 0x4f, 0xb8, 0x11, 0x84, 0x79, 0xbb, 0x6c, 0xf, 0xed, 0xc, 0xfc, 0x18, 0x32, 0x9d, 0xf1, 0x7, 0x2c, 0x20, 0xde, 0xe9, 0x97, 0x0, 0x62, 0x9f, 0x5e, 0x24, 0xfc, 0x8e, 0xc2, 0xd9, 0x2d}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
var _bindata = map[string]func() (*asset, error){
	"000001_message_type.up.sql":       _000001_message_typeUpSql,
	"000002_bandwidth_protocol.up.sql": _000002_bandwidth_protocolUpSql,
	"000003_index_truncate.up.sql":     _000003_index_truncateUpSql,
	"000004_envelope.table.up.sql":     _000004_envelopeTableUpSql,
	"000005_pushed_envelope.up.sql":    _000005_pushed_envelopeUpSql,
	"000006_status_version.up.sql":     _000006_status_versionUpSql,
	"000007_waku_push_filter.up.sql":   _000007_waku_push_filterUpSql,
	"doc.go":                           docGo,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"000001_message_type.up.sql":       {_000001_message_typeUpSql, map[string]*bintree{}},
	"000002_bandwidth_protocol.up.sql": {_000002_bandwidth_protocolUpSql, map[string]*bintree{}},
	"000003_index_truncate.up.sql":     {_000003_index_truncateUpSql, map[string]*bintree{}},
	"000004_envelope.table.up.sql":     {_000004_envelopeTableUpSql, map[string]*bintree{}},
	"000005_pushed_envelope.up.sql":    {_000005_pushed_envelopeUpSql, map[string]*bintree{}},
	"000006_status_version.up.sql":     {_000006_status_versionUpSql, map[string]*bintree{}},
	"000007_waku_push_filter.up.sql":   {_000007_waku_push_filterUpSql, map[string]*bintree{}},
	"doc.go":                           {docGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
