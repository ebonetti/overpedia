// Code generated by go-bindata. DO NOT EDIT.
// sources:
// db/base.sql
// db/indices.sql
// db/query-pages.sql
// db/query-toptenbyyear.sql
// db/test.sql
// db/types.sql
// templates/charts.html
// templates/page.html
// templates/pagelist.html
// templates/topten.html

package exporter


import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}


type asset struct {
	bytes []byte
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataDbBasesql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xdf\x6f\xab\x36\x18\x7d\xe7\xaf\xf8\xf6\x14\xa8\x58\xda\x44\x9a" +
	"\x34\x8d\xe5\xc1\x97\x38\x0d\x1b\x81\x08\x4c\x6f\x7b\xa7\x29\xa2\xc1\x69\xdd\xd2\x80\xb0\xd3\xae\x9a\xf6\xbf\x4f" +
	"\x18\x30\xce\xaf\xdb\x6a\x7b\x89\x82\x39\x9f\xfd\xf9\x9c\xe3\x63\xa6\x51\xb8\x84\xd8\x9d\xe3\x05\x02\x6f\x06\xf8" +
	"\xd6\x8b\x49\x0c\x6f\xe3\x02\x5c\x14\xbb\x68\x8a\x1d\xc3\x8d\x30\x22\xb8\x03\xbd\x8d\x0b\xc7\x30\x2e\x2f\x96\xe9" +
	"\x03\xe5\x50\xd1\xb2\xa2\x9c\x6e\x05\x87\x37\xf6\xcc\x4a\x9a\xb1\x14\xd2\x4a\xb0\x75\x4e\x39\xa4\xdb\x0c\x8a\x57" +
	"\x5a\x35\xc3\xa2\x28\xd9\x9a\x5f\x5c\x1a\xdd\x8c\x6e\xe8\xfb\x88\x78\x61\x50\x4f\x3a\x7c\x79\x5f\x17\x79\x9e\x0a" +
	"\x0a\xa6\x1f\xba\xc8\xc7\x30\x81\x01\xdd\xae\x92\x78\x98\x90\xd9\x8f\x3f\x0f\x2c\x47\x55\x12\xf4\xc5\xc7\xb2\xaa" +
	"\x94\x6d\x98\x06\x00\x40\xfd\x7f\xc5\x32\xf0\x02\x82\xaf\x71\x04\x41\x48\x20\x48\x7c\xdf\xee\xdf\x0a\x26\x72\x0a" +
	"\x37\x28\x72\xe7\x28\x32\x7f\x1a\x8d\x2d\x85\x6a\xdb\xc1\xfb\xcd\x68\xb5\xe9\x3d\x17\x55\xba\x16\x40\xf0\x2d\xf9" +
	"\x10\xbd\xab\xf2\xff\xb2\x4e\x45\xb7\xe2\x83\x3d\xf0\x62\xcd\xd2\xfc\x69\xf7\x52\xf2\x0e\xf6\xc7\x9f\x86\x25\x55" +
	"\x89\xe8\x2b\xe3\xac\xd8\xee\x29\x53\xa9\xc1\x62\x73\x42\x26\x4d\x91\x9e\xd7\xbe\xa6\xe1\xb6\xa2\xaf\xe7\xdb\xda" +
	"\x71\x5a\x69\x6f\xf5\x41\x7e\x5f\x08\xf8\x12\x86\x3e\x46\xc1\xa9\xcd\x9c\x9d\x53\x2e\xc8\x2b\xfa\x4a\x2b\xf1\x1d" +
	"\xc8\xfa\x31\xad\xde\x28\x7b\x78\x14\x30\xf3\x43\x44\xce\x40\x32\xb6\xd9\x9c\x05\x08\xf6\x42\xb9\x48\x5f\x4a\x20" +
	"\xde\x02\xc7\x04\x2d\x96\x07\xa8\x77\x9a\x56\x5d\x13\xc6\x49\x1f\xea\xa2\x7c\xda\x8d\xa7\x94\x1c\x5d\x75\x5a\xfa" +
	"\x45\x9a\xb1\xed\x03\x64\xa9\x48\x87\xc3\xa1\x94\x29\x5c\xde\x69\xb6\x9f\x45\xe1\x02\x7e\x19\xc8\x87\x0d\xcb\x69" +
	"\x99\x8a\xc7\x01\x7c\xf5\xc8\x1c\xdc\xf8\x06\xe6\x18\x4d\x71\xe4\xf4\x55\x4a\x54\xb3\x91\xd3\x6e\x85\xb3\x7b\xad" +
	"\xec\xb6\x6d\x5b\xe7\xdf\xde\x67\xda\xd6\x59\xb5\xf7\x18\xb4\xba\x9e\xd4\x52\xdf\xe9\x2b\x59\x4e\xbb\x73\xd0\xdb" +
	"\x2d\xc6\xa4\xa1\x7b\x52\x27\x10\x01\x13\xdf\x92\x08\xb9\xc4\xbc\xc3\x28\x6a\x66\xcf\x52\x41\x57\xa2\xda\x6d\xd7" +
	"\xe6\xa0\x86\x0e\xec\x7d\x19\x2d\x0b\x50\xdc\x11\x5a\x73\x89\x7c\x82\xa3\xc3\xd0\x90\x32\xa0\xe9\x14\x96\x91\xb7" +
	"\x40\xd1\x1d\xfc\x8e\xef\xc0\x6c\xb7\x6f\xd9\xea\xf5\x2c\x8c\xb0\x77\x1d\x74\xaf\xdb\x03\x6a\x41\x84\x67\x38\xc2" +
	"\x81\x8b\x63\x3d\x88\x8e\xeb\xdd\xd0\x4f\x16\x41\x23\x78\x46\x4b\xf1\xa8\x1c\x31\xc5\x33\x94\xf8\x04\xc6\x2d\x58" +
	"\x76\x79\x0c\xaf\x19\xd9\x37\xcf\x31\xb2\xc9\x35\x19\xe6\x1f\x41\xeb\x60\xda\x03\x3a\x9a\xaf\x74\x47\xb6\x4a\x6a" +
	"\x43\x9f\xd3\xb2\x61\xa2\xee\xfa\xc8\xe4\x13\x58\x0d\x0f\x07\x0d\x68\x56\x6a\x4e\x4d\x8c\x7d\xec\x92\xee\xf0\xd8" +
	"\xc0\x9f\x4e\x15\xb4\x25\x87\x1d\xf3\x27\x88\xbc\xeb\x39\x81\xdf\x42\x2f\xd0\x7a\x49\x62\x2f\xb8\xee\xb5\x31\x00" +
	"\x2c\x58\xc1\xd7\x39\x8e\x70\xd7\x10\xcb\x60\x22\x57\xe5\xdd\xb3\x63\x48\x96\x4e\x1e\x71\xc7\x30\xbc\x20\xc6\x11" +
	"\xa9\xb5\x0c\xb5\x95\x6e\x90\x9f\xe0\xf8\xf2\x62\x93\x3e\x53\x39\x5d\x1d\x83\x19\x6c\x8a\x0a\x1e\xf2\xe2\x3e\xcd" +
	"\x81\x8b\x54\x30\x2e\xd8\xfa\xe2\x52\x6e\xc4\xbc\xb2\x1b\xb9\xf6\x7e\xaf\x6c\x18\xfc\xfd\xcf\xc0\x86\xab\xda\xbf" +
	"\x27\xc9\x35\x7b\x8b\xd8\xa0\xf9\x72\x02\xe6\xa8\xae\x6b\xf7\xd7\xee\xe6\x87\xc9\x15\xa0\x60\xda\x3d\x4e\x54\x81" +
	"\xba\xdf\xbd\x60\x8a\x6f\x21\x0c\x8e\xdc\xac\x96\xe8\x7c\x26\x03\xd0\x4f\xe2\xda\x59\x87\x24\xcb\xff\xab\xf2\x99" +
	"\xbe\x3b\x06\x0a\x90\x7f\xf7\x4d\xeb\xdb\x31\x54\x72\x2e\x10\xc1\x91\x87\x7c\xef\x1b\x9e\xc2\x8d\x87\xbf\x2a\x94" +
	"\xa8\x28\x05\x14\x1b\x87\x4e\x50\x1d\x1b\x4a\xfb\xe6\x1c\x27\x81\x17\x06\x0a\x3e\x1e\xf6\x15\xa3\xe1\xb9\x22\x28" +
	"\x47\x87\x26\x29\xc7\xf5\xe6\x65\x4d\x4b\xd1\x78\xa8\xb1\xb4\x4f\x53\xd7\xe8\x6a\x9d\xef\xb8\xa8\xc3\x73\x9b\xd1" +
	"\xbf\x74\xf6\xe4\x36\xcc\xe3\xee\x2d\xe7\x88\x3b\x09\xed\xe9\x3b\x9e\xf7\x98\xca\x1a\xd3\xf7\xa4\xf8\x93\x9f\x58" +
	"\xcd\x76\x34\x06\xc7\xc3\x0b\x5b\x11\x23\xd5\xfc\x75\x5c\x27\x24\xe3\x12\x6f\x83\x1b\x22\x1f\xc7\x2e\x36\xbb\xdd" +
	"\x4b\x99\xfb\x9a\x46\xf5\xba\x44\x16\xac\x36\xbb\x3c\x6f\x20\x46\x5e\xbc\xd1\xca\xe4\x65\xce\xc4\xaa\x4c\x2b\x61" +
	"\x7e\x6e\x2e\x7b\x00\x03\x7b\x64\x69\x73\xca\xf1\xff\xaf\xd2\x61\xd4\xab\x8b\xe5\x74\xdc\x37\xf7\xe0\xf9\xb4\x6f" +
	"\x02\xe3\x73\x59\xaf\x27\xad\xbc\xc1\xf4\xe0\x3e\x34\x50\xeb\x14\xed\x33\x4b\x59\xa5\xae\xb5\xbb\x6f\x2a\xeb\xf4" +
	"\xf9\xd4\xea\xf4\x8b\xfa\x73\xe8\xee\xea\xfe\x18\xdd\x7f\x17\xc8\x6b\x54\x33\xa1\x02\x39\xc6\xbf\x01\x00\x00\xff" +
	"\xff\xf0\x82\x30\x34\x43\x0c\x00\x00")

func bindataDbBasesqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataDbBasesql,
		"db/base.sql",
	)
}



func bindataDbBasesql() (*asset, error) {
	bytes, err := bindataDbBasesqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "db/base.sql",
		size: 3139,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1525703780, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataDbIndicessql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\x5d\x6f\xa3\x38\x17\xbe\xe7\x57\x9c\xbb\x02\x42\x49\x1a\xe9\xbd" +
	"\x69\xdf\x59\x89\x4d\x3c\x1d\x56\x09\xe9\x00\x99\x99\xce\x4d\x44\xc1\x6d\x2d\x11\x60\xb1\xd3\x36\xff\x7e\x65\x9b" +
	"\x6f\x0c\xe9\x68\x3f\xee\x5a\xe7\xf8\x7c\x3c\xcf\xe3\xe3\x63\xe6\xa6\x87\xf3\x02\x53\x9c\x32\x0a\x24\x8d\xf1\x3b" +
	"\xb0\x73\x8e\x29\xb0\x97\x90\x01\x65\x21\x23\x94\x91\x88\x42\x96\x42\x1e\x3e\x63\x0a\x7a\x58\x30\x12\x25\x98\x42" +
	"\x98\xc6\xc0\xb2\x9c\x44\xd4\x30\xe7\x9a\xb6\xf2\x90\x1d\x20\x08\x1e\xee\x11\xbc\x2d\xb3\xd9\xf1\x2c\xfd\xd9\x3e" +
	"\x20\x77\xbf\x05\xfd\x2a\xca\xd2\xa7\x84\x44\xec\xca\x82\xab\x3c\x4b\xf0\x91\x44\x57\xc6\x6d\xbd\x73\x6b\x07\xc8" +
	"\x73\xec\x8d\xf3\x13\xad\xe1\x9b\x83\xbe\x0b\x37\x8c\x1c\xf1\x63\x76\x4a\x63\x0a\xb6\xaf\xf9\x68\x83\x56\x01\x6c" +
	"\x1d\x57\x3f\xe3\xb0\x30\xb8\xf7\x23\x49\xf9\xdf\x16\x6c\xed\x1f\xad\xd5\xf0\x5d\xac\x6a\xdc\xb6\xc0\xaf\x07\xee" +
	"\x88\xb2\xf0\x98\x57\x9b\xea\x05\xb9\x53\x61\x13\xbe\xd7\x0b\xda\x67\x6f\xb7\x15\x09\x15\xf8\x95\x50\x92\xa5\xf4" +
	"\x56\xd3\xe6\xa6\x23\x6a\x3c\x9e\x28\x83\x18\x3f\x91\x14\xc7\x40\x52\x08\xe1\x2d\x3c\x4b\x0c\x43\x38\x12\x4a\x49" +
	"\xfa\x0c\x38\x65\xc5\x19\xa2\xac\x28\x70\xc4\x92\x33\xb7\x0f\x4f\x09\x03\x96\xc1\x62\xb6\x30\xe7\xd3\x38\x90\x34" +
	"\x26\x11\xa6\x8f\x67\x5e\x15\x87\xe2\xbb\x13\x7c\x81\x92\x8c\x13\xc5\x05\xcd\x22\x12\x26\xa5\x19\x4f\x5f\xd7\x00" +
	"\x00\x4a\xc4\xd6\x8e\x1f\x38\xee\x2a\x00\x77\xbf\xd9\xdc\xdc\xb4\x09\x9a\x9b\xf8\x1d\xfc\x19\xdc\x67\xf9\x29\x09" +
	"\x0b\xc2\xce\xe6\x9c\xef\xe7\x42\xb0\x04\xeb\x07\x12\x5b\x20\x31\xe6\x91\x0e\x24\x16\xae\x87\x90\x88\xe5\xef\x5f" +
	"\x90\x87\x2a\x43\x70\x7c\x70\x77\x32\xac\xf8\x75\xef\x3a\x3b\x17\xec\xf2\xbf\x7e\x76\x8d\x46\xba\x39\xfe\x1b\xe9" +
	"\x80\xed\xae\x81\xb3\x4e\x68\x81\x5f\x71\xc1\xe0\x37\x58\x68\x86\x8c\xf1\x0b\x88\x56\x89\x15\x38\x65\x3c\x84\xed" +
	"\x5f\xcc\x72\x94\xb6\x3f\x76\x8e\x2b\x6a\xe0\x1e\x58\x81\x31\xec\x7d\xc7\xbd\x03\xbd\xf4\x68\x8c\x63\x68\x5e\x76" +
	"\xaf\x19\x96\x56\x1d\xdf\xa8\xc0\x21\x23\x59\x5a\xca\xa9\x5b\x5b\x9d\x7f\xe7\xa0\xf1\x3f\xa6\x90\xbe\xf3\x76\xfb" +
	"\x7b\xf8\xfd\xa1\xda\x5e\x61\xd9\x0e\x35\xc4\xb1\x8e\x55\x9e\xe3\x2a\x52\xcd\xb5\x04\xb4\x44\xb7\x5c\x8d\x71\xce" +
	"\x5e\xba\xb9\x34\x8d\xc2\xaa\x11\x6c\x2b\xa0\xd9\x07\xff\x87\x25\x8c\x03\xa9\x42\x68\x66\x5a\x97\x72\x50\x02\xdb" +
	"\xe1\x93\x0e\xc8\xac\x00\xca\x4e\x29\x9b\x40\x47\x44\xb2\xe0\x30\x93\xb8\xac\x76\x7b\x37\xd0\x4d\xe3\xe6\xe6\x29" +
	"\xc9\x42\x26\x0e\x47\xc6\xc2\xa4\x76\x35\x0e\xcc\x80\x0e\x0b\x9e\x71\x8a\x8b\x90\xe1\x03\xc5\x05\xc1\x54\xd0\x6d" +
	"\x95\xdd\xd3\x80\x83\xa4\x5f\xc1\x6f\x27\x29\xae\x2c\xe5\xb1\x11\xf9\x0c\xaa\x52\x9e\x65\x45\x59\x6f\x98\x3c\xbf" +
	"\xb4\xca\x51\x86\xe8\xe6\xa6\x70\x2d\x61\x26\x05\x8e\x7f\x21\xc5\x5e\x72\xf9\xf5\x4c\x26\x23\xd4\x58\x77\x4a\x0b" +
	"\xf2\x65\xeb\x87\xaa\x7b\x5d\xc8\x58\x46\xcc\xaf\xa5\x3c\xa6\x4c\x96\x3d\xc9\xc8\x74\x8c\xb6\xac\xaf\x67\xbc\x64" +
	"\xd1\xd8\xaa\xa6\x96\x2f\xe5\xda\xa7\xb1\x7e\xca\x11\xf1\xf3\xb0\xa0\x18\x7d\x6d\xda\x3e\xfa\xba\x2a\xad\x27\x65" +
	"\x58\x02\xd2\xc2\xa0\x0a\xd2\x70\x08\x21\x05\x85\x14\xe5\x29\xc8\xab\xba\x2f\x50\x92\x5f\x2b\xbb\x9f\x52\x86\xe3" +
	"\x49\xf1\x5a\x9b\x1a\xc7\xdb\xf8\x94\xb7\xa6\x8a\x29\xd4\x78\xa4\x51\x04\x27\xe2\x0c\x55\x73\x29\xca\x3f\xce\x9a" +
	"\xbd\x41\xfe\x0a\xe9\x02\x77\x6b\x61\x48\x2d\x4f\x76\x92\x7e\xdb\xa8\xa6\x30\x65\xe7\x58\x79\x3b\xdf\x97\xac\xf7" +
	"\xf7\x2d\xac\xa5\x01\x87\x83\xde\xe4\x2a\xb7\x08\xe3\x16\x6f\x6d\x25\xb4\x4a\x6a\x19\xd7\x30\x4c\x9a\x6e\xd0\xe7" +
	"\x40\xda\x4f\x1e\x80\x11\x1f\x6a\x04\x8d\x3e\x27\x77\xe8\x6f\x73\x22\xf7\xec\xb7\x92\x13\x03\x76\xdf\x90\x07\xfa" +
	"\xbd\xed\x05\x4e\xc0\x2f\x2d\xae\xff\x76\x1b\xea\xb9\x85\x9d\xb7\x46\x1e\xb7\xaa\x5c\xc2\x1a\xf9\x2b\xd5\xc1\x1c" +
	"\x97\xd9\x06\xfd\xd7\x25\x35\xc6\xe3\x05\x35\xce\x2f\x55\xd3\xe4\xcc\xab\xa9\x2e\xe8\xf2\xe9\x31\x28\xa3\x7e\x93" +
	"\x7c\x6c\xf4\xbc\x5e\x2c\x4c\xbd\xca\x77\xde\x4a\xca\x4c\xb2\x67\xbd\x7b\x21\xcf\xcb\x8a\x95\x57\xda\x74\x17\x6c" +
	"\xf4\x3d\xc2\x46\xa9\xd4\x29\x75\xd6\x2e\x7a\xc3\xc6\xe4\x41\x19\x4c\x4d\x9f\x60\xc9\x61\x1c\x9b\x87\xcd\x8f\x5c" +
	"\x7a\x2d\xcf\xf5\xad\x75\xf1\x75\x60\x76\x07\xac\x92\xa6\x71\xfb\xcb\x44\x8e\x8f\xea\x5c\x9e\x92\x22\x25\x5b\x7d" +
	"\x0d\x7d\x6c\x5e\x6f\xdd\x58\xf5\x08\x59\x4d\x4d\xf2\xb9\x7d\xf1\x71\x31\x9c\x37\x95\x4f\x06\xe5\x88\xc9\x7d\xf0" +
	"\xff\x95\xa3\x7e\x4f\xdb\xca\x21\xb7\x9e\xf2\x3a\x0c\x77\xc7\x7a\x91\x86\x2c\x66\x20\x2b\xc3\x12\x3b\x2f\x5d\x21" +
	"\xea\xfb\xc3\xe0\x4a\x3e\xe6\x09\x66\x78\x4c\x7a\x63\x35\xc8\xc9\x38\x27\xd1\xa0\x9e\x6a\xe2\x2c\xef\x3d\x49\x73" +
	"\x79\xf1\xf5\x39\xaf\xa2\x7a\xce\xdd\x97\xa0\xa9\xb3\x46\xb4\x2c\x57\xd1\x24\x0c\xcd\xd0\xd4\x29\xaa\xb3\x5a\xb4" +
	"\xde\x3e\xb8\x88\x70\xca\x48\x82\x0f\x31\xa1\x91\xbe\x98\xfd\xcf\x00\xfe\xe4\x77\xdc\x52\x4e\x7a\xdd\x10\x87\x7a" +
	"\x15\x79\xf7\x50\xd3\xc6\x46\x64\x55\x2e\x5a\x73\xb0\xba\x87\xb0\xe7\xb4\xf9\x7a\xe3\xb8\x6b\xf4\x03\x76\xae\xe2" +
	"\x63\x45\x2d\x46\xf1\xc5\x64\x4f\x71\x0c\x8f\x67\xd8\xd8\x01\xf2\xec\x8d\x04\x94\xa4\xf0\xe7\x49\xa8\xa0\xf9\x0e" +
	"\x32\xe5\xb1\x1c\xb9\xf9\x95\x56\x91\xd9\x94\xd1\x3f\x33\xc6\xad\x66\xbb\xf6\xe6\xe1\x27\x1a\xba\xba\xd5\x34\xed" +
	"\xaf\x00\x00\x00\xff\xff\xac\x50\x93\x65\x05\x13\x00\x00")

func bindataDbIndicessqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataDbIndicessql,
		"db/indices.sql",
	)
}



func bindataDbIndicessql() (*asset, error) {
	bytes, err := bindataDbIndicessqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "db/indices.sql",
		size: 4869,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1525703498, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataDbQuerypagessql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\x5d\x73\x9b\x38\x14\x7d\xf7\xaf\xb8\x8f\x90\x65\x93\x8d\x1f\xb3" +
	"\xbb\x0f\xd4\x26\x09\x1d\x02\x19\xc0\xcd\x64\x32\x19\x46\x06\x85\x28\x05\xc1\x48\x72\x1d\xff\xfb\x8e\x24\xbe\x8c" +
	"\x21\x4d\xfb\x66\xe9\x7e\x9e\x7b\xce\x15\xbe\x38\x23\xf4\xa5\x62\x25\x12\xa4\xa2\x1c\x5e\x2a\x06\x88\x09\x92\x16" +
	"\x98\x5b\x20\xaa\x9a\xa4\x1c\x10\xcd\x20\x2f\xaa\x2d\x2a\xce\x2e\x16\x0f\x6e\x7c\xdb\x1a\xec\x08\x8c\x05\x00\x40" +
	"\xe4\x78\xce\x2a\x86\x1a\xe5\x38\x21\x99\xbc\x57\x1e\x09\xc9\x94\xf9\x3a\x0c\xee\x60\xbf\xac\xce\xa5\x03\x57\x57" +
	"\x0f\xb7\x4e\xe8\xe8\x80\x0c\xd7\xe2\x15\xfe\x87\xcb\x85\x69\x41\x8d\x59\x8a\xa9\x20\x05\xce\x08\xcd\x48\x8a\x4f" +
	"\xab\x88\x43\x8d\xad\xb6\x96\x05\x07\x8c\x98\x05\x7b\x4c\xf2\x57\x61\x29\xbf\x26\x47\xc2\x10\xfd\x6e\x98\x10\x7c" +
	"\x73\x42\xd8\xcb\x34\x7d\x72\xed\x68\x64\x98\x72\x3c\xf2\xfb\x1b\x2e\xcf\xff\x31\x2f\x6e\x42\xc7\x8e\x9d\x28\x36" +
	"\xa6\x9c\x32\xf8\x0b\xa6\x63\x97\xa6\x75\x69\xca\x5a\xda\x3c\xae\x38\xca\x62\x47\xea\x66\xb6\x1b\xf1\xa9\x76\xc4" +
	"\x4c\x3f\xe2\xa8\x21\xcd\xc8\x74\x5b\x53\x13\x13\xfb\x3e\xea\x23\x18\x62\x3f\x20\x5c\x1a\x8e\x29\x6f\x58\xdc\x1e" +
	"\x24\x4f\x9a\x7a\xd7\x5f\x07\x0f\x9a\x11\xe3\xde\x0e\x63\x37\x76\x03\x1f\xbe\x3c\x36\xcc\x6a\x42\x07\xd2\x08\xc2" +
	"\xb5\x13\x4a\xbb\x26\xd9\xd4\x3d\xe8\xb2\x7f\x92\x00\xd6\x4e\xb4\x6a\xb2\x88\x4f\xb7\x61\x75\x9a\x9e\x69\x48\x7c" +
	"\xbe\xa3\xf9\x5c\xba\xb7\xe9\x4d\x40\x79\x3e\xb7\x72\x96\x2e\x84\x18\x43\x87\x04\xe5\xb9\xb1\xb2\xa5\x56\x9a\xad" +
	"\x18\x0a\xff\x54\x98\x5a\x82\xa7\x4c\xcf\x29\x66\x40\xb5\x46\xa6\xe4\x25\xb9\x96\x87\x12\x23\xbe\x63\xb8\xc4\x54" +
	"\x98\x3d\x38\x69\x01\x3b\x5a\x29\xd7\x81\x0b\xef\xc5\x72\x8a\x57\xd9\x6e\xc2\x60\x73\x2f\x53\x1c\x21\x9d\x1d\xd0" +
	"\x87\x33\x1a\x8f\x47\x93\x33\x6c\xa7\x83\x42\x68\x86\xdf\xa5\x7d\x79\x6c\xee\x10\x49\x5b\x87\x88\x0b\xf4\x31\x14" +
	"\x94\xe7\x93\x68\x14\x0e\x94\xe3\x94\x61\xf5\x02\xcb\x39\xcd\xbe\xab\x16\x94\x84\x76\x6b\xd4\x6d\x98\x20\x25\xde" +
	"\x56\x3b\x9a\x71\xeb\x57\x8f\xec\x7f\xb0\x04\x65\xd9\xf8\x52\x9f\xb6\xe7\x4d\x96\xb9\x73\x7d\xa3\x23\x76\xb2\x26" +
	"\xc3\x3f\x08\x97\x5f\x8c\x19\x54\x6d\x4a\x56\xed\x13\x51\x25\x6f\xbc\xa2\xcd\xcc\x55\x80\xfe\xa9\xbc\x05\x11\x45" +
	"\xfb\x9c\xa3\x2d\x17\x0c\xa5\xa2\x95\x58\x63\x23\x5c\x1d\xdb\xdb\x97\x5d\x51\x0c\xa3\x76\xac\xe8\x26\x63\xf5\x60" +
	"\x3b\x2a\xcb\x03\x7e\x17\x98\x66\x38\x93\xc6\x66\x59\x57\x81\xed\x39\xd1\xca\x31\x14\x75\x7a\x71\x9e\x9e\xaf\xae" +
	"\xe6\xb9\x7f\x7a\x3e\x09\xad\x52\x82\x8a\xb7\x5d\x59\x8f\x12\x94\x07\x59\xe9\xe9\xd9\x5c\x74\x4d\xc8\x0b\xf9\xa5" +
	"\x35\xcd\x45\x4f\x9c\x84\xa3\xd8\x02\x51\x83\xe7\x5c\xc7\xf0\x35\x70\x7d\xf0\xec\xd8\x09\x6d\xef\x58\x03\x63\xed" +
	"\xfe\xe6\xf0\x06\xd3\x50\x53\xe8\x85\x4c\x35\xcd\x03\x30\x3d\xd5\x3b\x4a\x31\x17\x86\xa8\x55\xff\xc9\xc0\xc7\x04" +
	"\xf5\x47\x20\x08\xd7\xae\x6f\x7b\x6e\xfc\x08\x89\xd1\xe9\x47\xa6\x54\x48\x46\x28\x37\x91\xeb\xdf\x40\xeb\x27\xa7" +
	"\x93\x40\xe0\x43\x1c\x6e\x9c\x45\x0f\x7f\x76\xaf\xc7\xe1\xda\xfb\x64\x7b\xc6\x6e\xfd\xca\x36\x30\xf4\x13\xdc\x9e" +
	"\xd4\x9c\xfe\x5d\xfc\x0c\x00\x00\xff\xff\x70\x5e\xd3\x18\x09\x09\x00\x00")

func bindataDbQuerypagessqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataDbQuerypagessql,
		"db/query-pages.sql",
	)
}



func bindataDbQuerypagessql() (*asset, error) {
	bytes, err := bindataDbQuerypagessqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "db/query-pages.sql",
		size: 2313,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1525703856, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataDbQuerytoptenbyyearsql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x53\x41\x6f\xa3\x3c\x10\xbd\xe7\x57\xcc\x11\x2a\x94\x26\xbd\x7e\xca" +
	"\x81\x2f\x61\x5b\x56\x14\x56\x40\x54\xed\xc9\x72\xc2\x88\x7a\x37\x31\x2c\x76\x94\xf2\xef\x57\x33\x80\xd3\x34\xd1" +
	"\x9e\x90\xdf\x3c\xcf\x3c\xbf\x37\x3c\x3e\xc8\xce\xaa\xfd\x01\xc1\x36\x2d\x58\xd4\xb0\xeb\xa1\x47\xd9\xc1\x9f\x13" +
	"\x76\xfd\xc3\xe3\xec\x2d\x2e\x5f\x18\x31\x10\x16\xe0\xcd\x00\x00\x8a\x28\x89\xd6\x25\xa3\x7c\xfe\x96\x67\xaf\x70" +
	"\x7e\x6a\xe6\x56\x1d\x71\xd7\x9c\x74\x65\x02\xa8\x51\x63\x27\x2d\x0a\x83\x9d\x42\xe3\x1d\x95\xa6\x0b\xc1\x51\x7e" +
	"\xd0\xd7\x07\xe1\xf1\x97\x3b\x6c\xd3\x38\x4b\x21\x4c\x92\xcf\xfd\x17\x34\x91\x87\xf8\x01\xe9\x53\xfb\x5b\x0d\xad" +
	"\xac\x51\xa8\x8a\x70\x66\x08\x55\x5d\x4b\x22\x82\x61\xe8\xed\x25\xca\xa3\xe1\x42\x85\xad\x7d\x87\x15\x2c\xb9\x73" +
	"\xdf\xe2\x6d\xe3\x4d\x5c\x94\x71\xba\x2e\xb9\x7c\xdd\x52\xe9\x4a\xed\xd1\xec\x7a\x67\xc0\x9d\xd6\x4f\xa3\xe8\xe5" +
	"\xe2\xa6\xb5\x98\xb3\x11\x20\xe6\xd4\x3b\x00\xd9\x75\xb2\x17\xb2\xae\xbd\x75\x58\x94\x9e\x67\x5b\x56\x2d\xac\xb2" +
	"\x07\x0c\x60\x3a\xca\x9d\xb1\x9d\xdc\x5b\x46\x86\xc7\x5e\x18\xca\x30\xe2\xd3\x30\xd2\x78\xec\xe9\x8e\x0f\x59\xbe" +
	"\x89\x72\xf8\xff\x27\x9c\x51\xd5\xef\x16\x36\x51\xb1\x66\x12\xfb\x02\x97\x87\x71\xc4\x93\xcd\xa3\x29\x01\x97\x93" +
	"\xb0\x8c\xf2\x30\x19\xdf\xf0\x25\xff\x81\x19\x4c\x39\x04\xe3\x1c\x47\xfd\x87\x69\x17\xe3\x78\xe1\x56\x83\x04\xf6" +
	"\x06\xc2\x74\xe3\xf2\x84\xd5\xa8\x6a\xee\x10\x2e\xf7\x2d\x52\x89\x84\xb2\x91\x8c\x7e\xc9\x60\x9a\x73\xcf\x06\x57" +
	"\x4c\xe2\xd7\xb8\x84\xe5\x82\x01\x1f\x04\x7c\xcf\xe2\x74\xd8\x67\x1a\x38\x38\x65\x5b\xd8\x16\x71\xfa\x0c\xde\xf8" +
	"\xd4\x61\x71\x9f\xf3\x6c\xfb\x83\x1a\x5f\x87\x4a\xe1\xd3\xf9\x97\x69\xf4\xdd\xff\x26\x80\xae\x39\x0b\xdb\x08\x62" +
	"\x8c\xb9\x0f\xf8\xcd\x36\x38\x7f\x8d\x4b\x57\xe9\x0a\x3f\x3a\xa9\x7f\x2b\x5d\x7f\xca\x98\xa8\xbe\x23\x49\xad\x4f" +
	"\xf2\xc0\x54\x34\x13\x99\xab\x34\xf3\x92\x3c\x2f\xe9\xf5\x63\x86\x9f\x6e\xd2\xcb\x74\xb7\x24\x7c\x72\x23\x09\xf9" +
	"\x6f\xf6\x37\x00\x00\xff\xff\xba\x35\x83\xe7\x47\x04\x00\x00")

func bindataDbQuerytoptenbyyearsqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataDbQuerytoptenbyyearsql,
		"db/query-toptenbyyear.sql",
	)
}



func bindataDbQuerytoptenbyyearsql() (*asset, error) {
	bytes, err := bindataDbQuerytoptenbyyearsqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "db/query-toptenbyyear.sql",
		size: 1095,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1525948399, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataDbTestsql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\xc1\x6a\xc3\x30\x10\x44\xef\xfa\x8a\xbd\x09\x02\x8e\xa0\xc7\x8a" +
	"\x9e\xfa\x09\x85\x9e\x74\x59\x2b\xeb\x76\x8b\x6a\xc9\xda\x8d\x41\x7f\x5f\x9c\x34\x46\x04\x7c\xdb\xd9\x99\x37\x30" +
	"\xee\x64\x8a\x2c\x09\x86\x15\x2e\xa8\x38\x71\xa2\x82\xfa\xfd\x66\x71\x14\xb7\x5d\xee\xc5\xbd\x7f\x7c\x5a\x18\x22" +
	"\xd8\xc0\xa0\x24\x7a\x96\x25\x59\x6f\x4e\xce\x98\x50\x84\x14\x0a\x7e\x51\x85\x3c\x4d\xde\x98\xf0\x78\xc8\xa3\x0b" +
	"\x5e\xfb\x66\xb0\x37\xf3\x1c\x65\xb5\xfe\x9e\xae\xb4\xb2\x70\x9e\x0f\x89\x3d\xd0\x53\x92\x23\x63\xfa\xb9\xfe\x96" +
	"\x43\xae\x8b\xfc\x93\x26\x30\x8c\x28\xb4\x4d\xf0\x9b\xe0\xf9\xc2\x91\x64\xd7\xda\x4a\xa7\x96\x2b\xd5\x36\x68\x2e" +
	"\x4a\xf3\xd8\x1a\x61\x7d\xb2\xee\x5b\x6e\xbf\xbf\x00\x00\x00\xff\xff\xf3\xcc\x98\x6f\x4a\x01\x00\x00")

func bindataDbTestsqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataDbTestsql,
		"db/test.sql",
	)
}



func bindataDbTestsql() (*asset, error) {
	bytes, err := bindataDbTestsqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "db/test.sql",
		size: 330,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1525705070, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataDbTypessql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x53\x3d\x6f\xf2\x30\x10\x9e\xc9\xaf\xf0\x08\x12\x7a\xa5\x17\xa9\x53" +
	"\x27\x17\x4c\x8b\x44\x01\x05\x97\x16\x21\x86\x2b\xb9\x52\x2b\x89\x13\xd9\x8e\x4a\xfe\x7d\x15\x3b\x11\x26\x84\x76" +
	"\x6b\x46\x3f\x5f\x7e\xee\x9c\x71\xc8\x28\x67\x84\x6f\x57\x8c\x7c\x8d\xb2\x7f\x25\x82\x4a\x11\x74\xa1\x30\x45\x69" +
	"\x08\x5d\x93\x7e\xd0\xdb\x40\x52\x20\xf1\xbe\xe9\x7c\x49\xf9\x30\xe8\xad\x50\x1d\x50\x1a\x91\x60\x0b\x98\xa0\xd4" +
	"\xe8\xa3\x0d\x10\x82\x8c\x7d\x27\x32\x5b\x54\xc7\x3c\xcb\xc5\xa1\x8b\x6f\x81\x9b\x6e\x16\xf5\x2d\x9d\xdb\x16\x41" +
	"\xb5\x43\x82\xc1\x7d\x10\xb4\xeb\x0a\x19\xe1\xc9\x94\x39\x8e\xbc\xd2\xda\xb5\xb6\xba\x06\xaf\x7d\x2a\x4d\x5a\x5a" +
	"\xd5\xd0\x12\x9e\x7d\x59\x4d\x68\xcd\x70\xb7\xef\x8c\x4e\xcb\x1c\x8e\x78\x8e\xe2\xc2\xb8\x29\x6e\x68\x38\x7e\xa2" +
	"\x61\xff\xee\xff\x68\xe0\x42\xe8\xbb\x36\x0a\x0e\x86\x10\xce\xde\xb8\x3b\xb3\xcd\xbb\xe9\x33\xed\xc0\x87\xe5\x72" +
	"\xce\xe8\xe2\x46\x3a\x9e\x0c\xca\x08\xa3\x3f\xba\x85\x3b\x9c\x16\x49\xf2\x83\xe6\x55\xc4\x22\xc7\x48\xc0\x8b\x4a" +
	"\x3a\xe0\xb1\x42\x30\x22\x93\x6e\xbd\xb3\x05\x67\x8f\x2c\x74\xd0\x0a\x8e\x38\xc1\xdc\x7c\x36\xeb\xae\x20\xdb\xfc" +
	"\xaa\x7a\x55\x58\xc8\x8f\x8c\x9c\x5b\x57\x6a\x72\xb1\x62\x7f\x3a\x2e\x61\x6d\xc0\xad\xb8\xe6\x74\x3f\x9d\xdd\xde" +
	"\xb1\xe7\x42\xc6\xfa\xc2\xb1\x72\xba\xf1\x14\xac\x95\x02\x19\x0b\x79\x3c\x5f\xca\x9e\x76\x3e\xba\x86\xfa\xbb\x33" +
	"\x48\x59\x40\x62\xa5\xa8\xaf\x12\xca\xfa\x37\xb9\x98\x64\x8b\xdc\xbe\x9f\x0b\xfa\x0e\x00\x00\xff\xff\xe1\xb2\x0d" +
	"\x3e\x37\x04\x00\x00")

func bindataDbTypessqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataDbTypessql,
		"db/types.sql",
	)
}



func bindataDbTypessql() (*asset, error) {
	bytes, err := bindataDbTypessqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "db/types.sql",
		size: 1079,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1525703624, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataTemplatesChartshtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xd1\xc1\x6a\xe3\x30\x10\x06\xe0\xbb\x9f\x62\xf0\x69\x77\x31\x72\xf6" +
	"\xea\xc4\x81\xec\xa6\x0d\x3d\xb4\x94\x34\x2d\x94\x92\xc3\x54\x9e\x26\x4a\x6d\xd9\x48\x93\xd8\x45\xe8\xdd\x8b\x9d" +
	"\xd4\x49\x20\x14\x7a\x1b\xcf\xff\xf9\x47\x48\x23\x2b\x8d\xaa\x18\xf8\xa3\xa2\x34\x64\x6a\x38\xde\xe0\x0e\xf7\xdb" +
	"\x10\xac\x91\x69\xb8\x66\xae\x6c\x12\xc7\x75\x5d\x8b\x95\x65\x64\x25\x85\x2c\x8b\x58\xae\xd1\xb0\x8d\xf3\x12\x33" +
	"\x32\x62\x63\xc3\xf1\x28\xde\xff\x38\x0e\x82\xef\x8b\xc7\xc1\x0e\x0d\xdc\x5d\xcd\x26\xd3\xc9\x62\x02\x29\x38\x70" +
	"\xce\xa0\x5e\x11\x88\xff\x5d\xad\xf7\x01\x00\x80\x73\xe2\xa1\x42\xed\x7d\x72\x4a\x6e\x74\x46\x0d\x7d\x99\x4e\x75" +
	"\xab\x96\xbd\xf4\x6a\x5e\xd6\x47\x02\xe0\x64\x72\xcc\xbc\x77\xbb\xc4\x39\xf1\x84\xf9\x96\xbc\x77\xae\x56\xbc\x06" +
	"\x71\x5d\x9a\x02\x99\x29\x3b\xec\xa3\xb7\x16\xb5\x39\xe9\xec\xc8\xee\xc9\x48\xd2\xac\xf2\x96\x54\x17\xc9\x1c\xf5" +
	"\xbb\xf7\x91\x39\x0b\x7d\x74\x18\x96\xfd\x74\x38\xde\xf2\xec\xfb\x34\xf5\xc3\x20\xe8\x6e\xab\x50\x7a\x96\x97\xaf" +
	"\x98\x2f\x54\x41\x96\xb1\xa8\x20\x05\x4d\x35\x4c\x91\xe9\x97\x73\xe2\x5f\xb9\xd5\x99\xd2\xab\x67\x42\x63\xc5\xad" +
	"\xd2\x3d\x14\x8f\x5a\x35\xde\xff\xf9\x3b\x18\x0c\x7e\x47\x50\x60\xf3\xa3\x26\x6c\x2e\x37\x0d\x83\xfe\xc1\x3f\x03" +
	"\x00\x00\xff\xff\xfd\xd3\x9d\x85\x48\x02\x00\x00")

func bindataTemplatesChartshtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplatesChartshtml,
		"templates/charts.html",
	)
}



func bindataTemplatesChartshtml() (*asset, error) {
	bytes, err := bindataTemplatesChartshtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "templates/charts.html",
		size: 584,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1521647185, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataTemplatesPagehtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x54\x5d\x6f\xe3\x28\x14\x7d\xf7\xaf\xb8\xcb\x73\x82\xb7\x6f\xab\xae" +
	"\x1d\x29\x6a\xbb\xab\x6a\xaa\x4e\xd5\x8f\x87\xd1\xa8\x1a\x51\xb8\xb1\x49\x6c\x40\x70\x9b\x34\x42\xfc\xf7\x11\x76" +
	"\x3e\xac\xce\x68\x9e\x80\xe3\xcb\xb9\x1f\xe7\xe0\xea\xaf\xeb\xaf\x57\xcf\xdf\x1e\x6e\xa0\xa5\xbe\x5b\x14\x55\x5e" +
	"\xa0\x13\xa6\xa9\x59\x8c\xfc\x4e\x98\x26\x25\x96\x71\x14\x0a\xb4\x1a\xd0\xe7\xbd\xc3\x94\x18\x28\x41\x62\x4e\xd6" +
	"\x69\x39\xc2\x79\x37\x46\xf7\x48\x02\x64\x2b\x7c\x40\xaa\xd9\xcb\xf3\x7f\xf3\x7f\x58\x79\xc4\x8d\xe8\xb1\x66\x5b" +
	"\x8d\x3b\x67\x3d\x31\x90\xd6\x10\x1a\xaa\xd9\x4e\x2b\x6a\x6b\x85\x5b\x2d\x71\x3e\x1c\x66\xa0\x8d\x26\x2d\xba\x79" +
	"\x90\xa2\xc3\xfa\x82\xff\x3d\x83\x5e\x1b\xdd\xbf\xf7\x67\xe8\x4c\xed\xbc\x75\xe8\x69\x5f\x33\xdb\x5c\xea\x5e\x34" +
	"\x38\xa1\x6f\x89\xdc\x65\x59\x9e\xda\xe2\x06\x1b\xe1\x50\x69\xc1\xad\x6f\xca\x21\x3c\x94\x9d\x6d\x2c\x77\xa6\xc9" +
	"\xa4\x45\x45\x9a\x3a\x5c\xc4\xa8\x57\xc0\x6f\xc3\xa1\xc3\x2b\x41\xd8\x58\xbf\xbf\x84\x18\xd1\xa8\x94\x72\xf3\x39" +
	"\x30\x25\x98\xc3\xfd\x91\xb5\x2a\xc7\xdb\x45\x51\x75\xda\x6c\xc0\x63\x57\xb3\x40\xfb\x0e\x43\x8b\x48\xe0\x3c\xae" +
	"\x90\x64\xcb\x80\xf6\x0e\x6b\x46\xf8\x41\xa5\x0c\x81\x41\xeb\x71\x55\x33\xce\xf3\xa9\x94\xd6\x23\xcf\xf0\x62\xca" +
	"\x23\x85\xb1\x46\x4b\xd1\x1d\xa3\x63\xe4\x57\x47\xec\xe5\xf1\x6e\xd0\xa1\x88\x91\xb0\x77\x9d\x20\x04\x96\xe5\xa0" +
	"\xc0\xb3\xc0\x0c\x78\x4a\x45\x51\x05\xe9\xb5\xa3\x69\xfa\xb5\xd8\x8a\x11\x65\x8b\x62\x2b\x3c\xdc\xdf\xfc\xbf\x7c" +
	"\x5c\xde\x7f\x79\x82\x1a\xbe\xc7\xe8\x85\x69\x10\xf8\xa3\x30\x1b\x6d\x9a\x90\x52\x01\x00\xf9\xc3\x00\xa5\x34\x8b" +
	"\x91\x3f\xa0\x97\x68\x48\xe7\x79\xe4\xf3\x35\x9a\x80\x9f\xc1\x5b\xa3\xf0\x63\xdc\x2e\x7b\x9b\xe5\xc8\xdb\x27\x27" +
	"\x4c\x4a\xaf\xb3\xc3\x5c\x0b\x80\xd7\x7f\x8b\xaa\x1c\x2b\xca\x83\x7c\xb3\x6a\x9f\x57\xa5\xb7\x83\x1b\x7b\x34\xef" +
	"\x6c\x51\x95\x4a\x6f\xa7\xf0\x41\xf3\x3c\x82\xaa\xbd\x80\x41\x3f\x63\x69\xa2\xe1\x68\xde\xd1\x8a\x31\xf2\x17\xa3" +
	"\xd0\x87\x3c\x6a\x75\x50\x92\x1d\x6a\x58\x9c\xc5\xad\xca\xf6\x62\x9a\x45\x1b\xf2\xf6\x9c\xfd\x37\x59\xce\xb1\x3b" +
	"\xbd\xd1\x83\x2b\xe6\x2b\x6f\x7b\x96\x59\x97\x6f\x81\xbc\x90\x94\x89\x3f\xd5\x9f\x9d\xb9\x13\x5e\x85\x4f\xec\xfc" +
	"\x4e\x9b\x4d\x98\xf2\x06\x2b\xf3\xf3\x58\xbf\xf7\x2e\x3b\x64\x2a\xb8\x13\x0d\x76\x3a\xd0\x51\xf2\xd3\xdd\x81\xf1" +
	"\x64\xdd\x71\xd2\xc5\x2f\x35\x78\x0c\xb6\xdb\xa2\xcf\xbe\xfb\xa3\x51\x20\x78\x39\xb8\x75\x1d\xca\xe3\xa5\x1f\xe7" +
	"\x57\xb6\x1e\x9a\x38\x4a\x78\x4a\x53\x0e\x52\x56\xe5\xf8\xdf\xf9\x19\x00\x00\xff\xff\xbd\xc4\x38\x8c\x88\x04\x00" +
	"\x00")

func bindataTemplatesPagehtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplatesPagehtml,
		"templates/page.html",
	)
}



func bindataTemplatesPagehtml() (*asset, error) {
	bytes, err := bindataTemplatesPagehtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "templates/page.html",
		size: 1160,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1526224456, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataTemplatesPagelisthtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcc\xc1\xad\x83\x30\x0c\x80\xe1\x3b\x53\x58\xb9\x03\x0b\x84\x48\x6f" +
	"\x07\xde\x00\x6e\x6c\xb5\x56\xd3\x80\x62\x4b\x3d\x44\xde\xbd\x82\xd2\xdb\x7f\xf8\xf5\x45\x2a\x90\x0b\xaa\x2e\xa1" +
	"\x88\xda\xc8\xd5\x9a\xb0\x86\xd4\x7b\xc3\x7a\x67\x98\xdc\x07\x00\x80\x48\x06\x84\x86\xa3\x6d\xbb\xe4\x25\xf4\x3e" +
	"\xad\x47\xb9\x1f\xef\xb4\x8a\x15\x76\x8f\x33\x59\xba\x7e\xfa\xc1\x6f\x79\xca\xce\x24\x38\xea\x63\x6b\x16\xbe\x4e" +
	"\xc5\x17\x9f\xcc\x7f\x25\x6e\x9a\xb7\xc6\x74\x29\xa7\xf8\x77\x53\x6b\x98\xed\x44\x29\xf5\xce\x95\xdc\x87\x38\x53" +
	"\x49\xc3\x27\x00\x00\xff\xff\x4f\xdb\xa9\x56\xb7\x00\x00\x00")

func bindataTemplatesPagelisthtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplatesPagelisthtml,
		"templates/pagelist.html",
	)
}



func bindataTemplatesPagelisthtml() (*asset, error) {
	bytes, err := bindataTemplatesPagelisthtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "templates/pagelist.html",
		size: 183,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1489957171, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataTemplatesToptenhtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x93\x41\x6f\xdc\x20\x10\x85\xef\xfe\x15\x94\x73\x16\x9a\x5b\x15\xe1" +
	"\xbd\xa4\xad\x54\x29\x6a\xab\x34\x39\xf4\x54\x4d\x61\x62\xcf\x06\x03\x82\x89\x93\x95\xe5\xff\x5e\x61\x67\x77\xad" +
	"\x4a\x2b\x31\xfb\x66\xde\x87\x79\xd8\xe6\xc3\xe7\x1f\xb7\x0f\xbf\x7f\x7e\x11\x3d\x0f\x7e\xdf\x98\xba\x08\x0f\xa1" +
	"\x6b\xe5\x34\xa9\x3b\x08\xdd\x3c\xcb\xaa\x23\x38\x41\xae\x95\x1c\x13\x63\x90\xc2\x01\xc3\x8e\x8f\x09\x97\xc1\x6f" +
	"\xc1\xe1\xdb\x3c\xbf\xcb\x09\x33\x45\xb7\x34\x7e\x25\x08\x67\x9d\x63\x22\xbb\xc8\x0f\xb5\x5a\xc9\x03\x32\x08\xdb" +
	"\x43\x2e\xc8\xad\x7c\x7c\xf8\xba\xfb\x24\xf5\x49\x0f\x30\x60\x2b\x47\xc2\xd7\x14\x33\x4b\x61\x63\x60\x0c\xdc\xca" +
	"\x57\x72\xdc\xb7\x0e\x47\xb2\xb8\x5b\xfe\x5c\x09\x0a\xc4\x04\x7e\x57\x2c\x78\x6c\xaf\xd5\xc7\x2b\x31\x50\xa0\xe1" +
	"\x65\xb8\x48\x17\x74\xca\x31\x61\xe6\x63\x2b\x63\x77\x43\x03\x74\xb8\xc1\xf7\xcc\xe9\x46\xeb\x73\x04\x2a\x60\x07" +
	"\x09\x1d\x81\x8a\xb9\xd3\xcb\x78\xd1\x3e\x76\x51\xa5\xd0\x55\x68\x63\x98\xd8\xe3\xbe\x1e\xae\x16\xf3\x2c\x76\xe2" +
	"\xfb\xc9\x65\xf4\xda\x6d\x1a\xe3\x29\x3c\x8b\x8c\xbe\x95\x85\x8f\x1e\x4b\x8f\xc8\x22\x65\x7c\x42\xb6\xbd\x14\x6b" +
	"\xa4\x8c\x6f\xac\x6d\x29\x52\xf4\x19\x9f\x5a\xa9\x94\x5e\x7f\xb6\x14\x6d\x63\x46\x55\x9b\xfb\x2d\xcd\x42\x88\x81" +
	"\x2c\xf8\x93\x67\x9a\xd4\xed\x49\x7b\xbc\xbf\x5b\xd2\x6e\xcc\xdf\xe8\x8e\x75\x75\x34\x2e\xf7\x39\x60\x78\x91\x7b" +
	"\xa3\x1d\x8d\x5b\xf9\x3d\x89\xc5\xd2\x5f\x6f\x4e\x65\x74\x7f\xbd\x1d\xa4\xc0\x39\x9e\x01\xd3\xc4\x38\x24\x0f\x8c" +
	"\x42\x26\xe8\xd0\x53\x61\x55\xdf\x29\x29\xd4\x3d\x84\x67\xaa\x61\x36\x8d\x79\x8a\x91\x31\x2f\x80\xb5\xac\x84\xb5" +
	"\xaa\xf0\xff\x9f\x26\x63\x89\x7e\xac\x53\x8d\x29\x36\x53\xe2\x6d\x4e\x07\x18\x61\x55\xa5\x28\xd9\x6e\xd2\x3a\x14" +
	"\x7d\xb2\xfe\xb9\xdc\xe5\xa1\xd4\xdd\x56\xc7\xfe\xb2\x99\x5e\xa2\x31\x7a\xfd\x12\xfe\x05\x00\x00\xff\xff\xca\xbe" +
	"\x65\xb5\x1a\x03\x00\x00")

func bindataTemplatesToptenhtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplatesToptenhtml,
		"templates/topten.html",
	)
}



func bindataTemplatesToptenhtml() (*asset, error) {
	bytes, err := bindataTemplatesToptenhtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "templates/topten.html",
		size: 794,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1526224461, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"db/base.sql":               bindataDbBasesql,
	"db/indices.sql":            bindataDbIndicessql,
	"db/query-pages.sql":        bindataDbQuerypagessql,
	"db/query-toptenbyyear.sql": bindataDbQuerytoptenbyyearsql,
	"db/test.sql":               bindataDbTestsql,
	"db/types.sql":              bindataDbTypessql,
	"templates/charts.html":     bindataTemplatesChartshtml,
	"templates/page.html":       bindataTemplatesPagehtml,
	"templates/pagelist.html":   bindataTemplatesPagelisthtml,
	"templates/topten.html":     bindataTemplatesToptenhtml,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"db": {Func: nil, Children: map[string]*bintree{
		"base.sql": {Func: bindataDbBasesql, Children: map[string]*bintree{}},
		"indices.sql": {Func: bindataDbIndicessql, Children: map[string]*bintree{}},
		"query-pages.sql": {Func: bindataDbQuerypagessql, Children: map[string]*bintree{}},
		"query-toptenbyyear.sql": {Func: bindataDbQuerytoptenbyyearsql, Children: map[string]*bintree{}},
		"test.sql": {Func: bindataDbTestsql, Children: map[string]*bintree{}},
		"types.sql": {Func: bindataDbTypessql, Children: map[string]*bintree{}},
	}},
	"templates": {Func: nil, Children: map[string]*bintree{
		"charts.html": {Func: bindataTemplatesChartshtml, Children: map[string]*bintree{}},
		"page.html": {Func: bindataTemplatesPagehtml, Children: map[string]*bintree{}},
		"pagelist.html": {Func: bindataTemplatesPagelisthtml, Children: map[string]*bintree{}},
		"topten.html": {Func: bindataTemplatesToptenhtml, Children: map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
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
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
