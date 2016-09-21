// Code generated by go-bindata.
// sources:
// static/css/app.css
// static/index.html
// static/js/app.js
// DO NOT EDIT!

package agent

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
	info  os.FileInfo
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

var _cssAppCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x56\xff\x8e\xa3\x36\x10\xfe\x3b\x3c\x85\xab\x55\xa5\xbb\xd5\xc2\xc2\xee\xed\xdd\x95\xa8\x52\xb7\x55\xd5\xad\xd4\xaa\x7f\xf5\x01\x0c\x1e\xc8\x28\xc6\x83\x6c\xb3\x49\x5a\xe5\xdd\x2b\xdb\x40\x80\x10\xf5\x56\x2b\x05\xe2\x6f\x3c\xbf\xbe\xf9\x26\x3b\xdb\xc8\x07\x56\x90\x38\xb1\x7f\xa3\x4d\x45\xca\xc6\x15\x6f\x50\x9e\x72\xf6\x0b\x75\x1a\x41\x3f\xb0\x86\x14\x99\x96\x97\xb0\x8d\x36\x2d\x17\x02\x55\x9d\xb3\x74\x1b\x6d\x1a\xae\x6b\x54\xe1\x79\x66\xfa\x06\xf2\x1d\x2c\x96\x7c\x1b\x6d\xe8\x1d\x74\x25\xe9\x10\x9f\x72\xb6\x43\x21\x40\x6d\xa3\x73\x14\x0d\x2e\x05\x9a\x56\xf2\x53\xce\x2a\x09\x47\x77\x27\xaa\x78\x07\x58\xef\x6c\xce\xb2\x34\x7d\xdf\x79\xb4\x01\x09\xa5\x75\xf8\xf8\x00\xc5\x1e\x6d\xcc\xdb\x16\xb8\xe6\xaa\x84\x9c\x29\x52\xe0\x50\x49\x87\x49\xd1\x59\x4b\xca\x21\x0f\x28\xec\x2e\x67\x5f\xd3\xf6\x78\x09\x3c\x96\x50\xb9\x9b\x5f\xdc\x97\xce\xa4\x24\x65\x41\xd9\xb5\x50\xdc\x47\x2c\x50\x43\x69\x91\x54\xce\x4a\x92\x5d\xa3\xfa\x83\x9c\xbd\xb8\x0b\x1e\xef\x7f\xfc\xff\xbf\xfb\xc7\xde\x91\x26\x69\x9c\x27\x97\x64\x1f\x5e\xf6\x29\xc4\x67\xe1\x68\x63\x2e\xb1\x76\x8e\x40\x59\xd0\xd3\x6a\x67\x01\x54\x90\x16\xa0\x63\xdd\x97\xa7\x3d\x32\x43\x12\x05\xbb\x7b\x7d\x7d\x5d\x94\x9a\x77\x96\x2e\x19\x7a\xc7\x49\xb9\x83\x72\xdf\x70\xbd\x4f\xb0\x0c\x35\x9a\x74\x70\x0e\x45\xd5\x76\xd6\x8c\xb8\xb1\x34\xae\xd4\x0e\x7b\x81\x9a\x96\x2b\xf7\xa6\xa0\xb4\x20\x46\x16\x19\xfc\x07\x72\x96\x65\xd3\x3a\x5f\xf0\x02\xcd\x4d\x93\xf4\xca\x24\xb1\x68\x25\x24\xa8\x2a\x4a\x2a\x04\xe9\x4d\x7c\xc1\x04\x94\xa4\x79\x68\x4f\xa7\x04\x68\x89\x0a\xb6\xd7\xc9\x5c\x19\xf6\x95\x76\x6c\x58\x8b\xcf\x34\x5c\x4a\xd0\x8b\xd0\xd2\xe4\x07\x68\x96\x97\xeb\x4e\xc5\xc6\x72\x6b\xd6\xc1\x0b\xe2\x3d\x8d\xd9\x35\x60\x0c\xaf\xc1\x9b\x0d\x18\x4b\x6d\xce\x9e\x47\x48\x87\x03\x6a\x19\xf7\xc0\x90\xb3\x07\xf9\xbe\x16\x74\x64\x92\x17\x20\x1d\xb6\xec\xb4\x21\x9d\xb3\x96\x70\x00\x7e\x33\x55\x7d\x6e\x1c\x55\x48\xbf\x25\x83\xa1\xbe\x1a\x24\xb7\xf8\x0e\xe3\x00\x64\x23\x11\x3c\x3a\x01\x81\x96\xf4\xc0\xef\x61\x88\x9f\xd2\x3e\x9f\xbb\xcb\xf9\x54\x2c\x46\x75\xf9\x0e\x9b\x96\xb4\xe5\xca\x6e\xa7\x6e\x79\x61\x48\x76\xd6\xb9\xf5\xe5\x71\x62\xd3\x0f\x40\xea\x47\xc2\x5a\x6a\xc2\x73\x28\xb1\x7b\x9a\x28\xc8\xf7\x8b\x28\x25\xd5\x13\x75\x08\x80\x99\x92\x2d\x87\x6e\x25\x99\x61\x10\x7d\x40\x37\xc7\xf0\x32\x84\x0b\xf7\x8e\x62\xb7\xa7\x6a\xc4\x25\xd4\xd9\x40\x47\xea\xec\xc3\xf2\x10\xb4\x0e\x87\xa0\xf5\xd5\x21\xaf\x9d\xa6\xf9\xe3\xf0\x38\xf5\x86\xca\x4f\xc9\x7a\x5c\xde\xa9\x63\x10\x49\x47\xa0\xbb\x34\x7d\x4a\xbf\x88\x5b\x60\x17\xc4\x04\xfc\xf5\x73\xf6\x29\x4d\x6f\x81\xc7\x48\x7a\x78\xad\xe1\xe4\xb1\x15\x49\x49\x87\x51\x6c\x56\x7b\xdf\xb7\xfc\x79\xa8\x7f\xe8\xfa\xf0\xca\xcb\x7d\xad\xa9\x53\x22\x67\xba\x2e\xf8\x87\xf4\x81\xf5\xff\x49\xf6\x71\xda\xd3\x67\x8f\x1f\xdb\x19\x5e\x7b\x2e\x3c\x65\x33\x91\xe5\x02\x3b\x93\xb3\x61\x53\x3c\xde\xff\xf9\xd7\xcf\xbf\xff\xf1\xeb\xfd\x63\xf4\x53\x03\x02\x39\xfb\xd0\xf0\xe3\x20\xe3\x5f\x9e\x3f\xb7\xc7\x8f\x2e\xfa\x61\xb1\xdd\x5e\x1f\xe7\x68\x33\xdb\x06\x9b\xc9\x3d\x3d\x1f\xd7\xd7\xc1\x42\xff\x03\x6b\xc6\x6f\x87\x9a\x5c\xf1\x71\xe6\x2f\x68\x69\xef\xf6\x42\xfa\x39\x66\xa1\xfe\x4b\xea\xc4\x85\xa4\x72\xef\x5c\xcf\xf4\xee\xc5\xeb\xdd\xb5\xf6\xac\x5f\x5e\x91\x6e\xe6\x97\xf7\xf9\xcc\xd1\x66\x47\x07\x75\xdb\x66\x88\x64\xe9\x62\x2a\xf8\xeb\xc5\x3c\xfb\xed\xfd\xf6\xf7\x6f\xb3\x8e\x4e\x17\xb3\xdb\xcc\xbe\xa5\xd3\x1f\x09\x57\x6d\xd5\x74\x98\xf8\x9f\x69\xcc\x4c\x26\x16\xdd\xea\x7f\x87\x5c\xf7\xea\x1c\xfd\x17\x00\x00\xff\xff\x86\x4a\x28\x36\x94\x09\x00\x00")

func cssAppCssBytes() ([]byte, error) {
	return bindataRead(
		_cssAppCss,
		"css/app.css",
	)
}

func cssAppCss() (*asset, error) {
	bytes, err := cssAppCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "css/app.css", size: 2452, mode: os.FileMode(420), modTime: time.Unix(1474465379, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x57\xdf\x8f\x1b\x27\x10\x7e\xf6\xfd\x15\x1c\x2f\x4e\xa4\xbb\xdd\xde\xa9\x95\x22\x6b\xbd\x51\x74\x49\xa5\x4a\x95\x1a\x25\x7d\xa9\xa2\x3c\x60\x98\xc5\xd8\x2c\xac\x80\x75\xce\x72\xfc\xbf\x57\xc0\xfe\x3e\xdb\x77\x57\xa9\x4f\xc6\xcc\x7c\x1f\x33\x1f\xc3\x2c\x64\xd7\x1f\xff\x7a\xf8\xfb\x9f\xcf\x9f\xd0\xda\x95\x32\xbf\xca\xfc\x0f\x52\xfc\x96\x54\xd5\x12\xff\x80\x55\x65\x34\xc5\xf9\xd5\x2c\x5b\x03\x61\xf9\xd5\x6c\x96\x39\xe1\x24\x78\x97\x95\x50\x6c\x89\x19\x71\x24\x79\xd0\xaa\x10\x3c\xf9\x6c\x34\x37\xa4\xfc\x60\xb8\xfd\xf6\xcb\x77\x9c\x37\xf8\x2c\x0d\x98\x80\x96\x42\x6d\x91\x01\xb9\xc4\xd6\xed\x25\xd8\x35\x80\xc3\x68\x6d\xa0\x58\xe2\x34\xa5\x4c\x6d\x6c\x42\xa5\xae\x59\x21\x89\x81\x84\xea\x32\x25\x1b\xf2\x98\x4a\xb1\xb2\xa9\x85\x92\x28\x27\xe8\x6d\x2d\xd2\xfb\xe4\x3e\xb9\xef\x66\x92\x52\xa8\x84\x5a\x8b\x9f\x5d\x84\x5a\x9b\x92\xaa\xea\x9d\x83\x4b\x8e\xbe\x29\x7e\x4b\xa5\x26\xdb\xef\xe8\x80\x98\xb0\x95\x24\xfb\x05\x52\x5a\x01\x3a\x66\x69\x74\xba\x9a\x65\x69\xa3\x43\xb6\xd2\x6c\x8f\x5a\x4c\x20\x12\x88\x4a\x62\xed\x12\x4b\x4d\xb7\xa8\xd0\x52\xea\x1f\x48\x50\xad\x70\x9e\xa5\x22\xae\x05\xd4\x09\xad\x5a\x47\xaa\x95\x33\x5a\xc6\x40\x66\xd9\xfa\xae\x35\x04\xbd\xe2\xec\x2c\xb3\x15\x51\xaf\xd4\xdb\x43\x1a\x74\x17\x95\x50\x55\xed\x2c\xf2\x6b\x82\x72\x31\xb0\x98\x80\xa0\xdb\x25\xb6\x6b\xfd\xe3\x8f\xe8\xb2\x44\xd7\xfd\xbf\x2e\x78\x9f\xfb\x5d\x1c\x30\xb1\x6b\x59\x6b\x81\x1a\xe2\x42\x9b\xb2\x21\x0c\x96\x83\xe7\x50\x0b\xd4\x53\x1d\xdb\x8c\xc6\xf8\x58\x50\x42\x15\x1a\x15\x02\x24\x6b\xbc\x66\x5f\x1d\x71\xb5\x8d\x88\x94\x89\xdd\x49\x70\x80\xd9\xe0\x39\x46\xf7\x89\xd3\x35\xd0\x6d\x49\xcc\x76\x98\x73\xab\xbf\x02\xea\x80\xa1\xf7\x68\xce\x0d\x80\x9a\xa3\x05\x9a\x1b\x60\xf3\x3e\xed\xc1\x0e\x9c\x82\x75\x7f\x02\x94\x09\xdb\x4f\xe0\x08\x9f\x75\x7b\x37\x02\x3e\x8c\x80\x1f\xcf\x00\x9b\xdd\x31\xd0\x58\xdf\xbc\xc5\xf9\xd0\x77\xb8\xd7\x51\x1b\xc5\x6f\xbd\xe4\x83\xd5\x5a\x49\x5e\xa4\x49\xa8\xaf\x2f\xb5\x52\x42\xf1\xcb\xb2\x34\xba\x34\x38\x5b\x12\x29\xc1\xe0\x49\xa5\x0e\x98\x9a\x61\xe0\xfa\xf4\x28\x42\xa2\x79\x1c\x8c\xb3\xe8\x37\x7b\xb4\xdb\xa6\x56\xb7\x7e\xa7\x2d\xbe\x98\xa3\x07\x67\xab\xbc\xd0\x26\x4b\x57\x39\x8a\x41\x12\xae\xdb\xa3\xb3\x26\x8a\x03\xfb\xe0\x7c\x2a\x61\xd5\xe1\x7a\x13\x09\x87\x29\x60\xcf\x5a\x09\x16\x58\x0f\x07\x14\x6c\x9f\x05\x43\xc7\xe3\x05\x86\xeb\x27\x14\x54\x33\x18\x71\x78\x09\x1e\x34\x83\x09\xd1\xb0\xe6\x27\xe5\xdf\xb1\x17\x42\x82\x4d\x24\x28\xee\xd6\x28\x47\x77\xf8\x84\x6c\x17\x4e\xd8\xec\x77\x8f\xbf\xac\x7a\x2d\x26\x98\xcc\x82\x04\xea\x06\xf6\x55\xed\x5c\x53\x44\xba\xf2\xfd\xcd\x2e\x71\xe1\xfb\x01\x2a\x90\x50\x28\x04\x19\xac\xa5\x66\xbe\x29\xc7\x86\x91\xf8\xf9\xb0\x09\x81\xef\xf9\xac\x9f\xcf\xe7\x4f\xcd\xb9\x50\xfc\x25\x2d\x63\xdc\x2b\xc6\x0e\xe1\x70\xac\xf4\xe3\xe0\xdc\xf8\x88\x91\xdb\x57\xd0\x1c\x1d\x6f\x45\x82\xf9\x5e\xcf\x75\xed\x4e\x64\xe7\x37\x28\xf1\xa6\x8e\x44\x92\x15\x48\x2f\x4b\x87\xca\xad\x63\xba\x76\x59\x1a\x4c\xcf\x94\xff\xeb\x03\x03\x63\xce\x05\xe6\x4d\x67\x02\x0b\x26\xeb\x18\x18\xf3\x7f\x05\x46\x38\xa8\xb3\x9a\x45\xe3\x99\xe0\x1a\x63\xf8\x39\x1b\xdd\xf9\xad\x1f\xef\x7a\xac\xdb\xd3\x85\xdc\x7e\x15\x1d\x31\xbe\xe7\x0e\x1b\x72\xfc\xba\x35\x33\x33\xa9\x09\x13\x8a\x2f\x50\x70\x4d\x84\xe2\x37\xfe\xfa\x40\x56\x12\xd8\x02\x5d\x5b\xb2\x03\x86\x7e\xfe\x6c\xcc\x71\x7c\xdd\xb5\xad\x9b\x96\xa6\x32\xa2\x24\x66\xef\x11\x8d\xe3\x0d\x0a\xad\x77\x81\x50\x87\x6d\x9c\x8f\x38\x3f\x1c\x46\x84\x6f\x9e\xf4\x5a\x08\xe6\xd0\x6b\xbf\x86\xd1\x5b\xdf\x5d\x62\x82\xcf\x34\x95\xc1\xf1\xb4\x67\x4e\xc8\xa4\x23\xbc\x44\x48\xb2\x83\x37\x6f\x47\xf7\x83\x36\xf5\x81\x84\x64\x07\x53\x05\x7b\x01\xbd\xf1\xac\x7e\x43\x01\xa3\xe3\x50\xbf\x38\xd3\xfa\x1e\x71\xfe\x95\xec\xe0\x70\x68\xe6\xdf\xcf\xd9\x7c\x31\x9f\x4f\x05\x7a\x45\x4b\x7c\x22\x00\x37\xb0\x6f\x55\x68\x97\x55\xfc\xb6\x4d\x2b\xea\xc1\x86\xa6\xee\x43\xbf\x83\x50\x71\xf9\x97\x30\x3a\x1f\xd2\x74\xff\xfc\xd9\x0a\x27\xbe\x8f\x42\x01\x27\x4e\xec\x00\x95\x60\x2d\xe1\x30\xf8\x70\x36\xd5\x63\x4c\x2f\x6d\xdf\x12\xfa\xea\x1a\xdb\xd1\xf1\x38\x59\xba\x1b\xf9\x36\x1e\xee\xb6\xe1\x4f\x49\xc4\xe8\x92\xdb\x9d\xe8\xa1\x8c\xc0\x84\xd3\x26\x5c\x48\x89\x50\x60\xf0\x34\x99\x60\xc7\xf9\x89\xd5\x66\x59\x65\xa0\xbf\x6e\xf3\x01\xc9\xb0\xc2\x74\xed\x16\x93\x76\x7c\x83\xc0\x98\xc5\xa4\x15\xde\xa0\xd0\x50\x16\x4f\xda\xd0\xd1\xaf\x5e\x99\xf8\x74\x49\x7d\x56\xf1\x1a\x4f\x8d\xa8\x1c\xb2\x86\xbe\xe0\xd1\x42\x14\xaf\x25\x31\xc9\xc6\xa6\x77\xc9\x6f\xc9\xbb\x6e\xc2\x3f\x59\x36\xe1\x7e\x1d\xf9\xfe\x03\x35\x85\xf4\x2e\xb9\x4f\x7e\xf5\x23\xcf\x85\xe8\x9a\x18\x0b\x6e\x89\x6b\x57\xbc\xbb\x40\xbd\xb1\xe9\x0e\xa4\x7e\xbc\x1c\xc0\x26\x3e\x95\x26\x3e\x59\xea\xdf\x3e\xf9\x55\x96\xc6\x37\xe3\xbf\x01\x00\x00\xff\xff\x53\xda\x3c\x3f\x44\x0e\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 3652, mode: os.FileMode(420), modTime: time.Unix(1470673026, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsAppJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x58\x6d\x6f\xdc\xb8\x11\xfe\xee\x5f\x31\x11\x16\x59\x29\xde\x50\xf6\x5d\x51\xa0\xde\xaa\x41\x90\x4b\xd1\x00\x97\xb6\x88\x83\xfb\xb2\xe7\x02\xb4\x34\x2b\x31\xa1\x48\x95\xa4\xd6\xde\xcb\xee\x7f\x2f\xf8\x22\x89\x5a\xaf\xd3\xbb\x7c\x88\x29\xce\x70\xf8\xf0\xe1\xbc\x71\xf3\x57\x50\x73\x79\x4f\x39\x2c\x56\xb4\x44\x78\x95\x5f\x5c\xec\xa8\x02\xac\x98\x91\x0a\x0a\xa0\x25\x12\xfb\x91\x26\x7e\x2a\xc9\xd6\x17\x7e\x44\x34\x9a\xdb\x46\x3e\xfc\x5b\x31\x61\x3e\x52\x55\x33\x91\x6e\x29\xd7\x38\x69\x2c\xee\xb9\x2c\xbf\xde\x96\x4a\x72\xce\x44\x0d\x05\x7c\x10\x5b\x26\x98\xd9\xc7\x46\x3e\x37\xd8\x62\x9a\xd0\x12\x73\x63\x87\x79\xcd\x4c\xd3\xdf\xcf\x77\x7a\xdb\x1b\xe9\x0d\xbd\x77\x73\x1f\x84\x91\xbf\x30\x7c\x48\x8d\xea\xed\x8e\x11\xea\x5b\x34\xbf\x50\xde\x23\x14\xb0\xed\x45\x69\x98\x14\xe9\x8e\xf2\x0c\xbe\x5d\x00\x58\x35\x06\x05\x8c\x96\xb5\x66\x52\x90\x4a\x96\xa4\x93\x9a\x59\xe5\xcf\xf2\x83\xa8\xf0\x31\x1d\x55\x38\x3a\x23\xa4\x46\xf3\xae\x57\x5a\xaa\x34\xcb\xd6\x17\x30\xd9\xf0\xfb\xd9\x4d\x56\xc0\xe0\x70\x80\xd7\xd7\xd9\xfa\xe2\xe8\x41\xd1\xae\xb3\x3c\x8a\xba\xe7\x54\x91\x56\x56\x3d\xc7\x74\xf9\x80\xf7\x9d\x92\xe5\x72\x05\x9b\xbb\x01\xbe\x66\xa2\xb4\xa8\xd3\x11\xf6\x84\x59\x97\x94\x5b\xd9\x66\x93\xb4\x3a\x59\x5d\x5f\x5d\x5d\xdd\xad\x60\x93\xe8\x64\xf5\x67\x3f\x6a\xc7\x51\x93\xac\x7e\xf8\x93\x1b\x55\xc9\xea\xc7\x6b\x2f\x35\x4d\xb2\xba\xfe\xe1\xee\xce\x02\x57\x68\x7a\x25\x26\x76\x2a\x6a\xd0\x6f\xe5\x37\xdb\x41\x01\x97\x02\x1f\xe0\x27\x6a\x30\xcd\x5e\x5b\xf9\xda\x49\xb7\x52\xa5\x03\x85\x57\x6b\x60\xf0\x57\x8f\x8c\x70\x14\xb5\x69\xd6\xc0\x2e\x2f\x07\x4b\x01\x38\x14\x5e\x65\xc3\xee\xd6\x61\x9e\x6d\xd3\x9d\x5d\xb9\xb9\xbe\xcb\x06\x30\x3b\xb8\x04\xbd\xb9\x1a\x75\x2c\x86\x8f\xd4\x34\x44\xc9\x5e\x54\xe9\x2e\x77\xda\x5e\x7a\x74\xff\x87\x85\xc9\xeb\xc4\xce\x1e\xd7\x17\x47\x77\x2f\x17\xb4\xeb\x48\xc5\x94\xbd\xb4\x9d\xf5\xab\x5a\x26\x2b\x38\xa1\x34\xac\xfd\x16\x0c\x69\xa3\x58\x69\x6e\x20\x79\x9b\xac\xdc\x14\x67\xe2\xeb\xcd\xb4\x48\xaf\x00\x57\x40\x8d\x51\x7a\x7e\xba\x6a\x05\x66\x1d\x7d\x97\x0d\x96\x5f\x63\xc7\x9b\xd4\x01\x4a\x8e\x54\x7d\x66\x2d\xca\xde\xa4\x26\x5b\x8f\x02\xb6\x4d\xab\x0c\x90\x18\x7c\x34\xa9\xf3\x82\xb4\xca\x22\xb9\xb1\x1c\xa2\x19\x96\xba\x4d\x56\x60\x7d\x60\x54\x3a\x86\xbf\x9a\x2c\x1e\xa8\x29\x9b\xd4\x81\x25\xb4\x96\xd1\xd9\x75\x8c\xa6\x82\x02\xc6\x3b\xd6\xd1\x6e\xce\x7c\x3a\x59\x8e\x48\x3f\x5e\x1c\xcf\x30\xcc\x65\xfd\xbb\x19\x5e\xbe\x5b\x9e\x67\xb8\x94\x1d\xae\x00\x39\xb6\x67\x89\xb6\xbe\xff\xc0\x44\x25\x1f\x88\x1d\x5a\xbd\xd8\x59\xa8\x02\xe1\x5c\x32\xf6\x3c\x97\x2e\xe2\x59\xb7\x09\xd9\x4a\xce\xe5\x03\x14\x60\x53\xc7\x20\xb2\x06\x89\x14\xe9\xd2\xaf\x5a\x46\xe7\xc1\x1d\x0a\x13\x33\x37\x19\xff\x89\x6d\xb7\x36\x60\x91\xf8\xef\x7f\x20\xab\x1b\x03\xaf\x01\x49\xc9\x19\x0a\xe3\x27\x22\x72\xed\xda\x0e\x55\x89\xc2\x5d\x6a\x64\xa5\x80\x2b\x78\x63\x2f\x15\x6e\x26\x83\x9f\x65\x97\x4f\x4a\xd9\xab\xeb\xab\xab\xb9\xad\xf1\x30\xa3\xd1\xa2\x80\x99\x16\xdb\xa6\x83\x52\x51\xcc\x28\x18\x62\x6f\xd2\x3d\x21\xc8\x0f\x26\x71\x25\xcb\xbe\x45\x61\xc8\x7f\x7b\x54\xfb\x5b\x97\x18\xa5\x4a\x93\xb0\x82\xb0\x52\x8a\x24\x23\xda\xec\x39\x92\x8a\xe9\x8e\xd3\xfd\x68\x06\xde\xc0\xd2\xd5\x83\x25\xdc\xc0\x52\x48\x81\xcb\x53\x17\x83\x91\xf4\xb0\xe6\x67\x59\xcf\x22\x88\x6d\xd3\xf9\x01\x22\x9e\xa0\x80\xbf\x84\x7f\xa7\x41\x31\x24\x5f\x7b\xc9\x28\x4c\xea\x1d\x29\x73\xf7\xad\x50\xb3\xdf\xd0\xde\xf7\xb0\x63\x36\x77\x97\x45\xd0\x42\xf3\xc4\x29\x56\x50\x51\x43\x63\x80\x33\x27\x04\xc0\x39\x53\x6f\x39\x4f\x13\xdd\x51\x4b\xd2\x56\xaa\xf7\xb4\x6c\xa6\x5c\x6f\xe7\x63\x53\x00\x76\x86\x28\x6c\xe5\x0e\xd3\xc8\x85\x8e\x4f\x42\x73\x06\x55\xd3\x1d\xfe\x0e\xa4\x79\x7e\x6f\x13\x2b\x94\xbd\x52\xd6\x6f\x98\xad\x79\x70\xbf\x87\x96\x09\xe0\xb2\x06\x14\x46\xed\x67\x07\x73\xe9\xb8\xa5\x8f\xa9\xf0\xe6\xc8\xcf\xb2\xfe\xd7\x76\xab\xd1\xbc\x1e\x3e\x3f\xd2\xc7\x5b\xf6\x1b\x46\x68\x1f\x1a\xc6\xd1\x57\xe9\xd9\xe1\xac\xf3\xb6\x50\x8c\x86\x36\xe2\x6e\x1d\x89\xd9\x36\x7d\xd1\x66\x70\xaf\x90\x7e\x8d\xe7\xc5\xe5\xe5\xfa\xc4\x8a\xa5\xc9\x1a\x1a\x9c\xb3\x54\x48\x0d\xbe\x0f\x77\x1d\xf8\x5e\x9f\xf2\x6a\x73\xed\x3b\x29\x8c\x0f\xc4\x96\xdc\x3f\xd1\x28\x39\xd5\xfa\x9f\xb4\x45\x27\xef\x62\x39\x12\xda\x75\x28\xaa\x77\x0d\xe3\x95\xbf\xb9\x58\xdc\x92\x85\x42\x51\xa1\xc2\xea\x24\xcd\x4c\x3e\x09\xb1\x8b\xff\x9f\x5c\xab\x7a\x31\xf9\xc9\x42\x49\x69\x6e\x7d\xba\x5c\x34\xc6\x74\x2b\x58\x18\x5f\x19\xa2\x76\x61\x4a\x97\x56\x1d\x0a\x98\x96\xad\x87\x36\x48\x74\xbd\x71\xe5\x99\x8c\x43\x7f\x43\xba\x91\x0f\x37\xf0\x4d\xf6\xe6\xc6\x82\x5f\x01\x2a\x15\x46\xb4\x46\x61\x6e\x5c\x9f\x77\xf4\x89\x7c\xcb\x38\xde\xc0\x72\x39\x7d\xe9\x1b\x10\x3d\xe7\xbe\x28\x5b\x4f\xd3\xa8\x76\xa8\xdc\x55\x87\xbd\x7b\x65\x13\x33\x97\x25\x75\x9d\x55\x47\x4d\x23\x68\x8b\x44\x61\xc7\x69\x89\x69\xbe\xf9\xcf\xaf\xf9\xdd\xe5\x22\x5f\x25\x49\x06\x97\x90\xe8\xbd\x28\x93\x01\xb9\x35\xe4\x70\x87\xc1\xb7\xe3\x20\xd9\xb9\x69\xfb\xff\x0e\xb9\x7c\x24\x5a\x63\xda\x2b\x1e\xdc\xdf\x6a\x69\xa2\xb0\x94\x42\x60\x69\xce\xd5\xe9\x1d\x51\x68\xd4\xde\x5f\xc9\x68\x96\xd9\x7b\x4c\xfc\xfe\x44\x8a\xbe\xb3\x1d\xd1\xb9\xe5\xb6\xb3\x21\xac\x82\x17\x45\x01\xac\x9a\x3c\xde\x19\xb0\x92\xf5\x54\xa7\xb1\x65\x66\xc8\x2c\xb3\xc6\x66\x94\x85\x50\x1e\xb1\x3b\x11\xed\x3a\x3e\xc3\x47\xa4\x28\x1b\x2a\xea\x19\xa0\x70\x44\x1c\x31\x68\x32\x4e\x41\x01\xe3\xf8\x59\xb3\x79\xee\xc8\xb2\x10\x62\x27\x1e\x9b\x9c\xdb\x20\x38\x47\xc1\x0b\xef\x4f\xc4\x3a\x83\x6d\x87\xe3\xef\xa8\xae\x3f\xb5\x0e\x63\x5b\x17\xf3\xe1\xf6\x74\xe5\x14\x0a\x88\x4d\xc1\xcb\x97\xb3\xef\x4d\xf4\x11\xb2\x89\x0b\x06\xef\x7e\x21\xd7\xfc\xfd\x19\xcd\x09\xcd\xb0\x97\x2d\x97\x6e\x69\x44\x4a\x78\x18\x79\xc2\xf5\xf4\x08\x90\x22\x4d\x9c\xcd\xa7\x9d\x90\x5d\x56\xd2\xb2\xc1\x21\xd7\x7a\x9a\x9e\xc3\x3d\x3d\x4e\xea\xe1\x61\x11\x2e\x7f\xe2\x3d\x3d\xef\x0e\x99\x83\xd8\x50\x51\x71\x8c\x20\x8e\x3d\x61\xe2\x08\x78\x27\xc5\x96\xd5\xe1\x4f\xaf\x5c\x04\x3a\x52\x62\xe8\xb3\xab\x1a\xe8\x2e\x7c\x78\xdb\x2b\xdd\x04\xd6\x6c\x45\xf6\xd2\xf0\x04\xf0\x0d\x88\x55\x39\x15\xfc\xcd\xce\xbf\x7c\x39\xf3\x86\x2c\x0a\x91\xc8\x67\x8a\x61\xcb\xb1\xc7\x73\x19\x71\x05\xe1\xb5\xf7\xe4\x4c\xdf\xc7\x9f\xe7\x8e\x25\x60\xc2\x30\xca\xdd\x21\xc2\x76\xc3\x19\x46\x77\x9a\xbb\xea\x53\x68\x3a\x7a\xcb\x95\xb2\xdb\x87\x8d\xce\x5f\xd0\x79\xc0\x91\xb5\x53\xc4\x71\x0c\xf9\xef\x38\x1a\xf2\x1c\x1f\x0d\x0a\xfb\x60\xd5\x90\x96\x8d\x94\x1a\x61\xab\x64\x0b\xb6\x0c\xe8\x9b\x3c\xbc\x9c\x49\x29\xdb\x9c\x7e\xa1\x8f\x52\xd5\xb9\x7b\x56\x2b\xc4\xbc\xa5\xda\xa0\xca\x39\xbb\x77\x73\xad\xac\x30\x1b\x43\xc4\x7e\x41\x01\xf9\xaf\x24\x35\xb2\xe5\x87\x2f\x5a\x8a\xc3\x17\x7d\x28\xb5\x3e\x34\xa6\xe5\x87\x5a\x1e\xf6\xf4\x4d\xcb\x0f\xba\x39\x3c\xb6\x3c\x5b\xe4\xc4\xa0\x36\x01\xf5\x1b\xf8\x84\xf5\xfb\xc7\x8e\x2c\xae\x6d\x73\xc7\x04\x5b\x0e\x90\x4b\xa9\x94\x7f\x3f\x8f\x64\xfb\xcd\x8a\x02\x96\xfb\x96\x2f\xb3\x61\xf3\xe5\x9e\xb6\x7c\xb9\x7e\xaa\x55\xcb\x48\xa9\x96\x9c\x8a\x3a\xa8\x4d\x81\x72\xeb\x1f\xf2\x69\x66\x9f\xe3\x1f\x65\x15\x7e\x50\xb0\xab\xf2\xe4\xd2\x1d\x76\x40\xc4\x25\xad\xbc\x0f\x38\xee\x5c\x68\xce\x9e\xbe\xcf\x87\xe6\xe1\x10\xaa\x40\x48\x4a\xbd\x52\xcf\x87\x2b\xdb\xa6\x4e\xc1\xd6\x82\xdd\xe0\x4a\xf3\x9f\x28\xd2\xdd\x88\xca\xbf\x1c\xd9\x16\x5c\x16\x3a\xef\x4e\x59\x28\xa7\x86\x2a\x93\xdb\x37\x15\x55\xc6\x67\x69\x3b\x3a\x5b\xce\xa8\x02\xca\x15\xd2\x6a\xff\xa9\x17\xc2\xff\x08\xe3\x22\x26\x7c\x8e\xb9\xcf\x5a\x20\x5e\x3e\x65\xe3\x61\x1e\xdd\x39\x6d\x51\xf7\xf3\xae\xef\x20\x5d\xef\xeb\x97\x55\x59\x66\xc4\x34\x28\x4e\x7f\xba\x98\x19\xb1\xb9\xf5\x04\xcc\x1b\x58\x7e\xf2\x06\xb0\x72\xef\x82\xdb\x30\x1e\xaa\xc1\xd0\xda\xcc\x0c\xcf\x4d\xba\x66\x64\x6d\x83\xed\xc7\xe9\x45\x7c\x8c\x22\x4b\xa1\xee\xce\xc0\x71\x67\xb2\x32\xd7\x46\x84\x65\x19\xd9\x32\x41\x39\xdf\x7f\xef\x24\x9e\x26\xbf\xef\xc5\xd4\xb7\x85\x12\x51\xca\xb6\x65\x26\xe4\xdf\xb1\x88\x9e\xbb\x1d\x2f\x39\xc7\xba\x9d\x7e\x86\x74\xa9\xa7\xce\x60\x96\xac\xbe\x77\x01\xce\xde\x69\xa5\x7d\x96\xdb\x51\xfb\x8f\x53\x3b\x01\xff\xa3\xcc\x4e\x4c\x9c\x23\x36\x74\x6e\x3b\x3c\xef\xe7\x27\x61\xf5\x4c\xa1\x7f\x2e\x4d\xaf\x5d\xab\xfd\xbf\x00\x00\x00\xff\xff\x7d\xb7\x54\x41\x0c\x15\x00\x00")

func jsAppJsBytes() ([]byte, error) {
	return bindataRead(
		_jsAppJs,
		"js/app.js",
	)
}

func jsAppJs() (*asset, error) {
	bytes, err := jsAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/app.js", size: 5388, mode: os.FileMode(420), modTime: time.Unix(1470673026, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"css/app.css": cssAppCss,
	"index.html": indexHtml,
	"js/app.js": jsAppJs,
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
	"css": &bintree{nil, map[string]*bintree{
		"app.css": &bintree{cssAppCss, map[string]*bintree{}},
	}},
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
	"js": &bintree{nil, map[string]*bintree{
		"app.js": &bintree{jsAppJs, map[string]*bintree{}},
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
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

