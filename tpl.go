// Code generated by go-bindata. DO NOT EDIT.
// sources:
// tmpl/dashboard.tmpl (26.296kB)
// tmpl/monitor.tmpl (1.548kB)
// tmpl/screenboard.tmpl (8.376kB)
// tmpl/timeboard.tmpl (2.261kB)

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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

var _tmplDashboardTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3d\x4b\x73\xdb\x38\xd2\xf7\xfc\x0a\x94\x4e\xdf\x97\x9a\x49\xee\x5b\x93\x43\x12\xdb\x93\x54\xd9\x89\x37\xf6\x64\x32\xbb\xb3\xc5\x82\x48\x48\xc2\x9a\x24\x18\x10\x94\xac\xa8\xf4\xdf\xb7\xf0\x20\x88\x27\x05\xc9\x52\x5e\x65\x1f\x52\x61\x77\xa3\x9b\x68\x34\x1a\x8d\x46\x83\xa2\xa8\x25\x1d\xcd\x11\x98\x14\x90\xc1\x82\xcc\xb3\x02\xb6\x8b\x29\x81\xb4\x98\x80\x49\x51\x64\x9b\x0d\x78\xf6\xb6\x00\xdb\xed\x04\x6c\x9e\x00\xc0\x30\x2b\x11\x90\x7f\x2f\xc0\x84\x63\x6f\x05\x68\xbb\x9d\x3c\x01\xa0\x84\x6b\xd2\xb1\x8c\xad\x1b\xd4\xa3\x2f\x05\xe8\x96\x43\x24\xcd\x66\x03\xf0\x0c\x3c\x3b\x43\x6d\x4e\x71\xc3\x30\xa9\xc1\xaf\xdb\xed\x13\x00\x0a\x03\xf2\x02\xfc\xf6\xdb\xf9\xfb\xdb\x27\x9b\x0d\x40\x6d\x0e\x1b\xf4\x7a\x01\x29\xcc\x19\xa2\xad\xdd\x74\xbb\x7d\xc2\xe9\x38\xdb\x5f\x01\xaa\xf9\x9b\x6e\x36\xcf\x9f\x72\x11\x4f\x9f\x0b\xb6\x4a\xde\x1f\xb4\x54\x72\x3a\x5a\x0e\x6f\xcf\xc1\xfd\x7b\x8d\x33\xb8\x45\x55\x53\x42\x86\x3e\x42\x8a\xe1\xb4\x44\xad\x62\xb7\xd9\x00\x0a\xeb\x39\x0a\x51\x08\x02\xa6\xc0\xd9\x52\xc1\x85\x26\x01\xa8\x61\x85\x0c\x3d\xbe\xe3\x8f\xf2\x55\xb4\xcc\x6b\x8a\x66\xf8\x5e\x09\x02\xa0\x91\x8f\x7d\x0b\x85\x1d\xda\x44\x3a\x60\xe8\x7c\x06\xbb\x92\x69\x7e\x85\x7a\x7e\x01\x22\x7a\x96\x68\xcd\x24\xc2\x7f\xeb\x6b\x4f\x6a\x44\x2b\x70\x5c\xb3\xef\x08\xc3\xb3\xf5\x25\x6e\xfb\x37\xab\x05\x20\x2b\x39\x84\xf7\xf6\xdf\x9b\x8d\x52\xf1\x40\xba\xdd\x4e\x36\x9b\x67\xdb\xed\xe4\x97\xcd\x06\xd5\xc5\x76\xfb\x1f\x4b\x52\x7c\xdc\xae\x29\x6a\x11\xdb\x3d\x7a\x3d\x5d\x78\x0c\xb3\x46\xa0\x9d\xa1\x1c\x1d\xcc\x98\x01\x25\x98\x50\xdc\x88\xc6\xcd\x08\x80\x25\x2c\x3b\x34\x60\x3f\x8a\xc7\x1e\x1d\x1c\x57\x6b\xe8\x4e\x34\xea\xaa\xb7\x7f\xe2\x62\xae\x55\xbc\x12\x0f\xaa\x5b\x9b\x0d\x98\x96\x24\xbf\x03\x13\x49\x73\x4d\xd1\x04\x3c\x33\x15\xc6\x35\x2a\x7d\x8b\x09\x5d\x61\xb6\xf0\xe0\xd2\x2b\x71\x2b\x57\x4a\xb9\xd7\x1e\x8c\xeb\xe4\x53\xaf\x63\x00\xd6\x16\xe2\xaf\x01\xb1\xc2\x05\x5b\x68\xc4\x9f\xe2\x49\x23\x17\x08\xcf\x17\xac\x47\xbe\x91\x4f\x0a\x1b\x54\xa2\x78\xcb\xdd\x1a\xf6\x50\x52\x25\xe6\xac\x96\xfd\x3d\x43\x33\x5c\x63\xc3\x95\x0a\xdc\x33\xe5\x74\xb3\x62\x40\xf7\x1a\x18\xf4\x7b\xb3\x80\x65\x49\x56\x52\xcd\xa6\x8e\xa5\x74\x3c\x03\xe8\x33\xf8\xbf\x33\xf4\x01\xcd\x6e\x18\xc5\xf5\x5c\xf2\xfd\x7f\x30\x81\x25\xa2\x2c\x9b\x53\xd8\x2c\x26\x83\x32\x4c\xde\xb7\x6f\x6f\x2f\xcf\x6d\x9e\xc3\x6c\x10\xcb\x86\x89\x90\x6b\x4b\xc4\x17\xf5\xab\x8c\xc1\x26\xaa\x35\x57\xca\x0d\xfe\x12\x90\x94\xb5\x1c\x2c\xc7\x6c\x20\x3b\x4c\xc2\xcb\x12\xcf\xeb\x80\x08\x28\xe0\xe6\x42\x29\x29\x87\xf9\x99\x20\x26\x66\x02\x21\x85\xe3\x0a\x45\xf5\x5d\x21\x17\x2e\xad\xc7\xc3\x30\x0e\x18\x66\x8b\x39\xe1\xf0\x12\xdd\x34\xd0\xee\x2a\x00\x25\x5e\xa2\xac\xe5\xf0\x7e\xd1\xef\xe9\xcc\x8e\x3a\xee\x59\xfe\x45\x3b\x6a\xcf\x91\xa3\x29\x8a\xf7\xe2\x25\x37\xdc\xb7\x85\xd5\x09\x69\xcc\xb8\xe8\x7b\xd0\xd3\x84\x47\xca\xe5\xf8\x11\x7f\x11\x93\xcd\xe4\xb8\xc4\x5f\xac\x40\xa8\xa7\x19\xe5\x28\x00\x65\x8b\x76\xcf\x3b\xe1\xd8\xdd\x79\xd7\xaf\x12\xce\xd4\x3b\xbd\x02\xae\x29\xca\x71\x6b\x3a\x20\x20\xe3\x15\x05\x95\xd3\x6c\xa0\x0a\x8e\x9b\xcb\xf4\x8f\x1a\x33\x8b\x5f\xc7\x01\x7d\xe0\xc6\xff\x9f\xf6\x72\xb7\xe8\x9e\x05\x66\x28\xba\x67\xce\x04\xd5\x74\x0f\x1e\xa3\x7c\xc1\x57\xb8\x84\xe1\x89\x12\xa8\x99\x1c\x18\xbd\x0f\xe8\x73\x87\x5a\xd6\xba\xb3\x59\x2d\xaa\x1a\x6d\x60\xa9\x84\x85\xa6\xf4\x15\x62\x14\xe7\xff\xec\x10\x5d\x3b\xb3\xfa\x73\xaf\x17\x93\x24\x36\xa1\x83\x53\x52\x0b\x79\x2d\xd4\xe1\x4d\x11\x00\xa4\x9e\xac\x79\x62\xd0\xee\xf4\x1e\x83\x00\x52\x35\x90\xa1\x5b\xe2\xf2\xe7\x70\x8a\x32\x46\x34\x7b\x4d\x99\xce\xfd\x6d\x9d\x53\x04\x5b\xf4\x3b\x21\x85\x23\x00\x2b\x54\x36\xe7\x38\x69\xe6\x16\xb9\xc3\x2f\x2e\xe4\x3d\x2d\x10\x7d\xe5\x0e\x02\xe1\xd0\x6c\xba\xee\x5f\xbf\xa7\x4a\x7f\x79\xd1\xe2\x0c\xd3\x20\xe3\x02\x53\x8b\x33\xa7\x4b\x67\x7d\xb3\x20\x2b\x11\x26\xd7\xcc\xe1\xde\x2e\xc8\x4a\x86\xc8\xb5\x9e\xb1\x26\xf5\xe1\xeb\x82\x13\x9e\x06\x28\x6c\x2b\x4c\x9b\xad\x28\xbf\xcb\x5a\x06\x59\xd7\x9e\x66\xce\xbe\xe6\x12\x2c\x1d\x09\x99\x83\xcd\xf3\x87\x34\x6f\xf6\x3b\x25\x5d\xc3\xdf\xde\xe4\x36\xef\x81\x8a\xa1\x26\xda\x83\xa7\xcf\xd0\xe2\xb6\x0f\x2b\xc7\x8c\x05\x33\xc3\x88\x7b\x9a\x44\xf7\x0d\xe7\xb6\xaf\x63\x1c\x60\xee\x0a\x39\x45\x70\x3f\x18\x61\x9d\x62\x13\x05\x6e\x19\xc5\xd3\x8e\x87\xcd\x3f\xbe\x1f\x0f\xc6\xd4\xb6\x5f\x4f\x9d\xf3\x6c\x5d\xba\x5e\x5c\x87\x92\x21\x64\x2b\x60\xe6\x1b\x0f\x51\x03\x2c\x11\x63\x6e\x03\x00\x1a\x05\xef\xf3\x1c\xea\xd1\x76\x1a\x91\xd7\x4c\xe8\xc7\x29\x9c\x8d\x81\x7f\xfe\x54\xba\xbf\x12\xcd\x39\x21\xac\x0b\x20\xff\x2b\x77\x1c\x35\x61\xa0\xed\x9a\x86\x50\x86\x0a\x30\x5d\x03\xd8\xe0\x7d\x1d\x16\x5a\xa2\x9a\x65\x2d\xa3\x08\x56\xa7\x31\x4e\xdf\x90\x3e\x0b\x88\x1a\x12\x3f\x18\x88\xcf\xdf\x73\xfe\xb2\xde\x1e\x4c\x75\x41\xee\xc1\x04\xcf\x81\x2e\xdd\x2f\x9c\xdf\xa3\xbc\x63\x6e\xdc\xc9\x1d\x44\x86\x34\xaa\x0f\xef\x2c\xfa\x07\x87\x78\xb2\x03\x7c\xaf\x54\xe2\xfa\x44\xa1\xde\x11\x47\xe1\x1b\x6b\x6b\x46\x11\xca\x78\xc8\xed\x2a\xaa\x0f\xd0\xbd\xd8\xdc\x8c\xca\x13\xfb\xf8\x9a\x94\xc4\x8e\x73\x72\x01\xd1\xa1\x1f\x7f\x48\x63\x75\x41\x02\x36\x3b\x23\x8e\xc9\x6a\xaa\xef\x76\x23\xb2\x40\x90\x55\xb0\xf9\xf1\x57\xb0\x23\xed\x44\x7e\xcc\xe5\x2b\x92\x11\xd9\xd9\xed\xe3\x2e\x6c\x7d\x96\xe9\x2f\x78\x8f\xdb\x48\x9a\x49\xe0\xc2\x79\x26\x1f\xb5\x16\x90\x50\x9e\x09\x4e\x51\xe9\x26\x99\x04\x4c\x9f\x2a\xf1\x87\x3d\x76\x2a\x39\xf4\x47\x56\xc0\xfa\xcd\x89\x78\x48\x67\x78\x85\xdd\x24\x58\x85\xf5\xfc\xe5\xc8\x3d\x58\xc1\x7b\x97\x15\xbc\xd7\xac\xe0\xfd\x7e\x1b\xd5\xb2\x2b\xd0\xbf\x10\x75\x37\xc2\x58\x62\xb2\x2f\x1c\xa5\xb7\xa9\x9a\xf8\xe0\x08\xe9\x44\x69\x3a\x65\x69\x22\x24\x88\x99\x9a\x44\x46\x7c\x8f\x42\xba\x01\x47\xc8\xd8\x82\x1e\x47\x2d\xb3\xc1\xb8\x39\x2d\x62\xfe\xca\x2a\xfb\x0a\x51\xe7\x82\xb4\xe9\x6b\x49\xe2\x5a\xf1\xfc\x29\xf8\x20\x46\x2c\xf4\x82\xbf\x80\x69\xc7\x40\xd5\x95\x0c\x37\x25\xd2\x6b\x87\x50\x40\x0b\xc4\xb9\x85\xa4\x65\x8e\xc6\x46\x56\x99\x0b\x5c\xba\x9e\x45\x7b\xa8\x00\x6e\xc6\x41\x41\xef\x1f\x5f\xae\x0e\xd9\x72\x7d\x95\xe5\x60\x70\x87\x6e\x64\x63\x2e\x82\x3e\x4e\x58\xd0\xcf\xa2\x84\xbd\x56\xbc\x70\xc8\x30\x12\x30\xf8\xe1\xc2\x68\xb0\x90\x14\x2a\x8c\x79\x7c\xd5\xe2\xa2\xc4\x4d\x98\x75\x36\xe3\x28\x95\x7f\x37\x88\x93\x37\xfc\x7c\x52\xf8\xab\x1d\x9f\x17\x99\xb1\xe4\xf5\x54\xe9\x6f\x2e\x5a\x78\x4b\x9f\xe4\x3b\xac\x7f\x3d\xd5\x41\x09\xc3\x10\xc8\x1d\xdd\x77\xa4\xf0\x13\xd3\x35\x29\xec\xa4\xb4\xa6\x4a\x0b\xf4\xdf\x11\x69\xe3\x6f\x88\xeb\xf1\x6a\x92\x55\x02\x95\x2d\x04\x4e\x0e\x8c\x4d\x1f\x7e\x71\x4f\x84\xc8\xa2\x05\x25\xc8\x7c\x9b\x2d\xc0\xa0\x4e\xe2\x1f\x4f\x09\x0e\x79\x37\x41\xb3\x4f\xe2\xcd\x88\xc7\x88\xa3\xf2\x56\x40\x4c\xee\x82\xe6\xd8\x69\x3d\x3c\xa3\xb0\xf2\x76\xeb\x76\x81\x90\xfc\xeb\x68\xe9\x17\x09\x3d\x44\x72\x05\xfd\x13\xa1\x87\x0a\x36\x3d\xba\x9b\x13\x6e\x25\xa8\x0f\x70\xe5\x53\x1a\xb7\x2b\x48\xe7\xce\x94\xaf\x24\x48\x87\xa5\xe2\xe9\xc1\x4a\x29\xc9\xfc\xa4\x69\xac\x4b\x32\x6f\x91\x9d\x5b\x28\x25\xa8\xdf\x48\xc8\xa7\x34\xb5\x1c\x31\x1d\xf3\x9a\x94\x5d\x55\xb7\x6e\xb2\x42\xc0\x5e\x80\xc9\x30\x0b\x14\xa1\x3f\x0f\x92\xac\x62\x41\x56\x67\x90\x21\xc9\xc4\xb6\x0e\x1e\x2c\x16\x90\xa1\x4c\x8a\x55\x8e\xc2\x69\x91\xe4\x2a\x78\x9b\x2b\xd4\xb6\x70\x1e\x15\x54\x49\xb4\x2f\xcb\x6e\x97\x24\x4e\x35\x39\xc3\x6d\x53\x42\x7b\x38\x7a\x31\x85\xc2\xe9\xb4\x81\xd5\x22\x71\x46\x11\xca\x22\xab\xbe\x8b\x69\x39\x80\xcb\xfa\x7b\x22\x3b\xf8\xf7\xe4\x1f\x3a\xe5\x24\xfb\x35\xf9\x05\xfc\x3d\x11\x67\x6e\x1a\x27\x0e\xdc\x38\x6a\xd7\xeb\x1c\x30\xaf\x2a\x58\x73\x3d\xec\x73\xa4\x75\x42\x63\xbf\xe9\xaa\x0a\xd2\xb5\xb7\xcc\xb6\x12\x6e\xad\xb4\x26\xed\x81\x03\xa5\x87\x63\xa3\x90\x69\x7c\x94\x79\x5c\x10\x5a\x41\x9b\xa1\x32\xa7\x6c\x26\x51\x8a\xb5\x4d\xbf\x47\x86\xf2\x9a\xa2\x19\xa2\xa8\xce\x91\x9f\xab\xcc\x9a\x01\x69\xa6\x2d\x8d\x36\x69\x92\xde\x60\xb9\xb5\x7f\x4d\x3a\x77\x97\xbc\xc0\x2a\x11\x90\xe5\x12\xa9\x0a\xd7\xec\x16\xc9\x53\xff\x12\xb6\xec\x96\xe2\xf9\x1c\x51\x54\xf8\x53\xbf\x84\x2d\xcb\x98\xc6\x0f\x53\xdf\x6e\xf7\x50\x8b\xaf\x09\x0b\xae\xae\xaf\x49\xcd\xdc\xf3\xe9\x5c\xc1\xb4\x82\xe5\x63\x9a\x62\x5f\xc1\xfc\x8e\x87\x42\x75\xe1\xa7\x9b\xa7\x1a\x97\x59\x99\x67\xb7\xcd\xcf\x93\x83\xb6\x2c\xe1\x16\x3b\x87\xdc\xc2\x00\x18\x16\x07\xdd\xfd\xb0\x0b\xa2\x24\xe3\xe2\x94\xd7\xc4\x39\xfb\xc5\xf9\x5d\xd6\x90\x76\x28\xaa\x93\x34\x89\x9d\xc7\xf9\xdd\x79\x31\x77\xcb\xdd\xf2\xbb\x0c\x71\xa8\xc1\x53\x50\x3d\x38\xbc\x11\x1e\x33\xb9\x54\xeb\x80\xf8\xe6\x07\xcc\xc0\x5f\x92\x79\x48\xc2\x50\xbc\x1b\xc6\xf3\x48\x51\x2e\x40\x5e\x22\x42\x25\x0b\x5f\x36\xd5\x7b\xda\xb7\x76\xf4\x09\xcc\x1c\x69\x81\xdc\x5d\x27\x00\x58\x40\x55\x27\x25\x49\x3c\x5f\x1f\xe9\x20\x30\xab\x91\xba\xc0\xc1\x80\xee\xe3\x08\x81\xea\x8c\xa2\x08\x74\x23\x57\x6d\x5f\x38\x8a\xd0\xd2\x5f\xce\xe7\x14\xcd\xa1\x77\xce\x27\xff\xa0\x81\xed\x4b\xfd\x0c\x90\xdb\xe9\xc8\x1e\xde\x10\x77\x01\x73\xe4\x96\xff\xf0\xbf\x99\x80\xf7\x4e\x4a\x3c\xec\xcf\xfc\x6d\xcd\x10\x5d\x42\x37\x2b\x07\xc4\x88\x29\x54\x9f\xcd\x56\x8f\x01\x56\x41\x19\xb1\xc4\x53\xa8\xd2\x3a\x42\x15\xca\x3e\xed\x63\x27\x57\x5d\xc9\x70\xdc\x16\x9e\x3f\x05\x16\x85\x97\x22\x65\x88\x52\xc8\x83\x12\xa0\x6e\xef\x80\x86\x92\x25\x2e\x10\x7d\xc0\x4b\xdd\x20\x48\xf3\x45\xdc\x76\x23\xf8\x56\x82\xa3\x56\x19\xce\x11\xee\x9f\x69\x3f\x74\x44\x1f\x3a\x56\xa1\xba\x26\x60\x39\xdb\xa1\xaa\xc9\x22\xd0\xc5\x4f\x11\xc5\x1c\x69\xfe\x44\xdf\x7f\x70\xbc\xb8\xc2\x21\x41\xa5\x80\xcb\x59\x24\x69\x46\x14\xbe\x5b\x8c\x17\x8e\x6b\x6c\x64\xf7\x24\xff\x54\xd0\xee\x2a\x29\xd1\xab\x1d\xe0\xd7\x46\x9c\x8f\x16\x2a\xb7\x68\x21\x71\x62\x2b\x67\x95\x4e\x1e\x28\x22\x36\xfc\x89\x06\x30\x22\x62\xc7\x20\x86\x27\x44\xd2\x58\x8f\xce\xa1\xc0\xd1\x72\x22\xdb\xa0\x0f\xb6\x56\x75\xb7\xcd\x36\x1e\x7f\x3c\xe4\x78\xe4\x65\x53\x8d\x86\x28\x11\x3c\x6c\xaa\x58\x88\x32\x04\x73\xe3\x51\xca\x09\x3a\xa4\x02\x8a\x6b\x4a\x72\xd4\xb6\x5a\xac\xd7\x35\x59\xc6\x3f\x10\xc5\x3a\x3f\x42\xd3\x48\x54\x4c\x09\x43\x88\xe9\xd9\xbb\xcc\x89\xdb\x41\x66\x5a\xc1\x84\xb9\x6e\xe5\x8b\x80\x83\x96\x0b\x93\x51\x7d\x2a\x17\x30\xaf\x86\x7a\x17\xff\x0b\x5c\xb2\x40\x81\xb6\x38\xb2\x60\x56\x8d\x76\x4f\xb9\x1f\xff\xb0\x7b\xee\x9d\xf3\xc4\xf4\xce\x5f\xbd\x8e\x24\x2d\x3e\x32\xa3\xd7\xba\x10\x17\xce\x60\x29\xd3\x23\xde\x19\xbe\xb9\x3f\x77\x48\x7d\x3a\x9d\x0a\xf5\x48\x2d\xca\x7c\xc0\xab\x54\x4d\x1b\x36\x42\x71\x07\x80\x42\x46\x7c\xdf\x9e\x0f\x28\xf3\xc2\x80\x84\xec\x37\xa0\xf2\xae\xa5\x2b\x40\x5e\xc8\x0c\x5e\xc7\x4c\x63\x7b\x8a\xaa\xa3\x41\x33\x5d\xcb\x48\xf5\x6a\xee\xa7\x36\x84\x72\x04\x36\x9b\xce\xed\xec\x86\xdd\xe8\x10\x81\x17\xe3\x02\x67\x41\x81\x17\x07\x09\x7c\x5b\xc1\x39\x72\x8f\x5c\xf8\x9f\x38\x9f\xc9\x8c\xc3\x17\x4d\xf9\xd0\x09\x17\x5c\x11\x8f\x35\xe3\xcc\xd8\xc8\xd3\x20\x1c\x10\x4e\x4c\xe4\xe9\xed\xe4\x65\xd1\xf2\x4d\x3b\x46\x5a\xaf\x16\x0b\x6a\xa8\x8c\x42\x07\xaa\xa4\x3c\x91\x34\x07\xef\xb6\x98\x32\x1e\xf3\xd2\x98\x41\xf9\x0d\xef\xb5\x7d\x83\x32\x50\x99\x88\x62\x70\x5a\x3e\x26\xa2\x4e\x97\x88\xfa\x86\x51\xde\x4f\x17\xb6\x0e\xa2\xdd\xc8\x35\x4e\x19\x0e\x3b\x82\xaa\x2a\x31\x74\x03\x0d\x28\x60\xfa\xfa\x2b\x7f\x48\xaf\x66\x39\x91\x0b\x1e\x8d\x0e\x77\x6c\xdc\xc7\xd8\x85\x36\xb6\x3b\xb7\xb4\xa7\x59\x22\x0c\xe4\x0e\x27\xd6\xe6\x90\x31\x44\x9b\x92\x78\xb7\x08\x0e\x73\x62\x41\xbc\xae\x6e\xde\xc3\xcb\xc9\xf9\x15\xc4\x8e\x78\xb9\x4f\xb1\x89\xea\x22\xee\xbf\x61\xf1\xdd\x6e\x0b\x77\x6c\xdc\x33\xf1\x43\x63\xa7\x44\x53\xfe\x2b\xa6\x45\x17\x31\xba\x23\xfd\x69\xb5\x78\xf4\x32\xc9\xe4\xf2\x65\x73\x82\xa8\x7d\xe1\xa7\x91\x9b\x03\x9f\xe2\x37\x07\x7c\xd4\xfd\xe3\xcd\x81\xc7\x9b\x03\x43\x19\xbc\xb7\x23\xe0\x1b\xc3\x57\x6b\x91\x91\xf7\x6a\xac\x08\xcd\xa6\x6b\x59\x23\x69\x5f\x23\xb6\x5a\x1d\xbb\xf0\xb0\x45\x74\x89\x73\xe4\xd7\xcf\x3f\x7f\x1a\x2e\xce\x5f\x51\xd8\x34\x88\x7a\x1f\x9d\x49\xab\xde\x91\x39\x28\xbb\xef\x33\x05\xb3\xd2\x54\xa9\x47\xe7\x37\xf2\xfd\xed\xe3\x7d\x44\xf1\x72\x28\x54\xe9\x49\x1e\xbc\x5b\x51\xba\xca\x4a\xb4\x44\x65\x46\xa6\xff\x45\x39\xc3\xcb\xc4\x9d\x8b\xfb\xe2\x1f\x31\x5a\x05\x3e\xce\x82\x56\xce\xd7\x59\x14\xd5\x5e\xda\xb8\xe4\x2f\xf8\xbe\x7f\xbf\xb7\x67\xb6\x76\x4a\x62\x7c\x4f\x25\xd6\x22\xbd\xba\xe2\x9c\x52\x42\x5f\x75\xe2\x2b\x59\x5e\x91\x05\xe2\xc8\x6c\x2a\xb1\x43\xb1\x85\xd9\x26\x69\x6f\xca\xf5\x70\x45\x8a\x80\xb6\x2a\x0e\x35\xb4\x25\xa8\x52\xcb\x2e\x2a\xf4\x27\xae\x0b\xb2\x72\x8b\x39\x2a\x94\xad\x14\xc2\xba\xd0\x6f\x34\x38\xf6\x44\xe4\x42\xb9\xe1\xa2\x23\x7d\xea\x21\x88\x1f\x2e\x2e\x1d\x33\xcc\x7c\xdc\x4c\x3f\x6e\xa6\x0f\xbf\xf3\x83\xf2\x8e\x62\xb6\x1e\xed\xd5\x18\x51\x8d\xd8\x8a\xd0\xbb\xef\xb4\x7b\x1f\xba\xf1\xf1\x8a\xe0\x69\xf7\x3d\x8e\x57\x3c\xf9\x11\x3c\x43\x51\x95\xa8\x81\xef\x2b\xf5\x95\xab\xe6\x52\x67\x50\xef\x11\x06\x7f\xe7\xf7\xa6\x8d\x44\x49\x8d\xe4\xf7\x19\xfd\xa3\xb4\x9a\xaf\x36\x1c\xa5\xcf\xd3\x7a\xda\xfd\x65\x04\x94\xad\x44\x58\x5f\xbf\xed\x29\xbf\xd1\xa1\x5d\xbf\x00\xc0\x02\x32\x18\x3d\x58\x8b\x10\x54\x3d\x38\x38\x8a\xe7\xf7\x0d\x45\xad\x97\x1b\xe7\x7f\x68\x40\xf5\xdf\xf6\x18\x20\xfb\xa9\x5a\xe4\xe1\xc4\x37\x4d\xbd\xed\x32\xc7\x64\xe2\xeb\xa7\x66\xce\xce\xf9\xfe\xe9\x61\xda\x3e\xe4\xc4\xe6\x68\xa9\x30\x2b\xe8\xbb\x94\xd7\x79\xfd\xaa\x6a\x09\x37\xaa\xa9\x25\x60\x6f\x21\xb2\x9d\x57\x66\x6c\xde\x1d\x56\x19\xc6\x81\x70\x6f\x21\x57\x90\xde\xb9\xbb\x11\xc3\xfc\x14\xd6\x40\x56\x02\x14\x8a\x66\x0e\xf6\x75\x7b\x6c\x86\x43\x27\xb7\x3b\xcf\x6d\x47\xb3\xb7\x47\xca\x4a\x1c\xcd\xc6\x92\x02\x65\xd2\x94\xb8\x65\x93\x63\x04\xc9\x47\x88\x70\x1f\xe3\xd9\xc7\x78\xd6\x17\xbd\x33\x3e\xda\x75\x3a\xf4\x6b\x7c\xca\x9c\x64\x52\x51\x98\xa3\x4c\x25\x38\x4e\x73\x14\x7b\x5e\x2f\xad\x5e\xa1\x7a\xa9\x97\xe1\x7a\xb9\xdf\xb7\x87\xc7\x52\x3f\x02\x96\x94\xfa\x19\x65\xde\xc0\xda\x5b\xdf\xdb\x06\xd6\xd6\xd2\xae\xa9\x0e\x10\xb0\x20\xab\x37\xd8\x71\x35\x62\x0d\x5d\x60\x7d\xf7\x49\x13\x1d\xb4\x48\x8b\x2c\x4b\x40\x00\x92\x70\x27\x1d\x73\xa0\x90\x4b\xc8\x50\x9d\xaf\x43\x17\xac\x24\xc2\xbc\x59\x25\x21\x07\xc9\x79\x45\x11\xbc\x2b\xc8\x2a\x70\x8b\x73\xaa\x51\x83\xac\x81\xfc\x20\x69\x67\xc6\x67\x16\x7d\x81\xe6\x47\x18\xcd\x0b\xaa\x26\xf8\x20\xb1\x1f\xd4\xef\x5e\x18\x3f\x3a\x60\x88\xed\x7f\x15\x43\xfe\x02\xc1\x20\xd7\x6a\xb5\xbf\x5c\xfc\x05\x05\x2e\x17\xf2\x10\xcb\xb9\x59\x68\x50\xee\x6f\xef\xa7\xb8\xc6\x18\x73\x76\x16\x52\x7d\xc0\xbf\x5d\x61\x96\x2f\x46\x48\xf9\x7f\x65\xf2\x5e\xb5\x28\x90\xc3\x79\xcc\x81\x8a\xbc\xbd\xe5\x38\x77\xfc\x26\x89\xfc\x4b\xf9\x65\x12\x2d\x3e\xf4\xf3\x04\xf2\xcf\xfa\x91\x02\x4d\x3f\xb8\x65\xeb\xd7\x0a\xc2\x0b\x6d\xe0\x63\xfd\x9a\x62\xec\x93\xfd\x01\x69\xde\xb7\xfb\xf7\x5a\x77\x0d\x19\x35\x6a\x19\x2a\x0e\xdc\xce\x38\x78\xf9\xe1\x08\xd3\x58\xe4\xbf\x3b\x5e\x61\xd7\xaf\x4a\xa8\xc1\x10\x64\xdb\x27\xff\x0b\x00\x00\xff\xff\x39\x04\x13\xd1\xb8\x66\x00\x00")

func tmplDashboardTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplDashboardTmpl,
		"tmpl/dashboard.tmpl",
	)
}

func tmplDashboardTmpl() (*asset, error) {
	bytes, err := tmplDashboardTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/dashboard.tmpl", size: 26296, mode: os.FileMode(0644), modTime: time.Unix(1597675860, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6d, 0x8, 0x48, 0x6e, 0xda, 0x28, 0xc2, 0xbc, 0xc1, 0xe0, 0x60, 0xe7, 0x9d, 0xa7, 0x99, 0x80, 0xe5, 0xd7, 0xfa, 0xc2, 0xa1, 0x4a, 0xb3, 0x29, 0xdc, 0x7f, 0xfd, 0x5b, 0x94, 0xd3, 0xdc, 0x58}}
	return a, nil
}

var _tmplMonitorTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x94\xc1\x6e\xe3\x2c\x10\xc7\xef\x7e\x8a\x91\xcf\xdf\x97\x27\x68\x0f\xab\x36\xab\xe6\xb0\x89\xb6\x8a\xd4\xc3\x6a\x85\x90\x99\xc4\xa8\x04\x52\xc0\xb1\x22\x96\x77\x5f\xd9\x40\x6c\x1c\xda\xcd\xc9\xf3\x9f\xff\xfc\x98\x0c\x63\x6b\x34\xaa\xd3\x0d\x42\xcd\xa8\xa5\x4c\x1d\xc9\x49\x49\x6e\x95\xae\xa1\x66\x8c\x38\x07\xab\x0d\x03\xef\x6b\x70\x15\x80\xa4\x27\x84\xfc\xf7\x08\xce\x01\x9a\x86\x9e\xf1\xa9\xa5\x9a\x36\x16\xb5\x81\xd5\x76\x70\x7a\x5f\x01\xd8\xeb\xf9\xbe\xa6\x1e\xc0\xfb\x21\xe3\x7d\x5d\x01\x38\xf7\x3f\xf0\x03\xac\xf6\xf4\x68\x62\xd9\xf0\xb4\x2c\xfb\xe5\x9c\xa6\xf2\x88\xc1\xe8\x7d\xf1\x68\xef\xff\x73\x0e\x25\xf3\xfe\x77\x24\xa3\x64\x01\x7a\x42\x63\xe8\x11\x73\xe8\xc3\xc3\x7a\xb7\xaf\x86\x86\x7e\xc4\xb4\xf7\xd5\x20\xc1\x08\x17\xd4\x72\x25\x49\x2a\x9d\xfb\x77\xe7\x21\x65\x56\xeb\x9b\x6d\x41\xa8\x00\x3e\x3a\xd4\xd7\x30\xa5\xd5\xcf\xf1\xf9\xcf\x7d\xcf\xde\x57\xb1\xd5\x9e\xdb\xf6\x06\x0e\x4d\xa7\xd9\x6c\x95\xe5\x87\xeb\x56\x3d\x53\x4b\x43\x46\x8e\x0a\x91\x8a\x0c\x97\x77\xbb\x8d\x82\x33\x9b\x42\x02\xbe\x62\x00\x6c\xa4\x45\x7d\xa1\x22\x64\x75\x54\x09\x4f\x72\x80\x96\xdd\x45\x70\x38\xff\x5b\xc7\xb8\xcd\x1a\xa5\xa3\x32\xef\x71\xe6\x29\x92\xf6\xfc\x84\xaa\xb3\x2f\x71\x27\x42\x44\xda\xc8\xc8\xb3\x45\xc0\x46\x36\xa2\x63\x38\xed\x15\x0f\x02\x19\xf7\x2b\x60\xe6\x9e\x2f\xc7\xf5\xd1\x71\x8d\xdf\x3b\x21\xde\xb8\x64\xaa\x4f\xf3\x1a\x65\x72\xe8\x84\x20\x7d\x48\xa4\x89\x15\x0b\xca\x23\xc3\xfe\x45\x19\xfb\x8c\x82\x5e\xe3\xcc\xb0\x27\xad\x32\x96\xb0\x51\x8b\x53\xbb\xb3\x15\x69\xeb\x0b\x15\xdd\xb8\x90\x33\x27\xde\xc4\x0c\x59\xf4\xce\xa8\xf3\xdb\x68\x35\x9a\x56\x09\x36\xdb\xcc\xb0\xb0\xcb\x8c\x9d\xe2\xc7\xf1\xcb\x31\x41\x76\xef\xc1\x02\xa0\xde\x63\x0b\x93\x94\xfd\x9b\xa9\xe6\x8d\x6a\xc9\xe5\x31\xc9\x7d\x0c\x43\xf5\x22\xf9\x35\xe2\x15\x1b\x75\x19\xde\xc2\x1c\x45\x74\xd2\x33\xe6\xd2\xfd\x09\xfb\x49\x73\xcb\x9b\xf4\x4a\x00\x34\x29\x0e\xb0\x65\xfa\x1f\x94\xe5\xa1\x89\xb6\xec\xf1\x33\x7f\x86\x2f\xae\x48\x29\xa8\x7c\xf5\x37\x00\x00\xff\xff\x04\x8a\x39\x87\x0c\x06\x00\x00")

func tmplMonitorTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplMonitorTmpl,
		"tmpl/monitor.tmpl",
	)
}

func tmplMonitorTmpl() (*asset, error) {
	bytes, err := tmplMonitorTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/monitor.tmpl", size: 1548, mode: os.FileMode(0644), modTime: time.Unix(1597657642, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x64, 0x19, 0x3a, 0x9f, 0x4a, 0x2c, 0x34, 0x55, 0xfd, 0x94, 0x8f, 0x0, 0x93, 0xfd, 0x10, 0xd2, 0x18, 0xce, 0xa5, 0x51, 0x87, 0xae, 0x8f, 0xcf, 0x9d, 0xd6, 0x12, 0xde, 0x4c, 0x8f, 0x42, 0x65}}
	return a, nil
}

var _tmplScreenboardTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x59\xdb\x72\xdb\x38\x12\x7d\xf7\x57\xa0\xf4\xbc\x71\xbe\x20\x0f\x8e\x9d\x8b\xab\xec\x8d\xd7\x72\x9c\xdd\x9d\x9a\x62\xc1\x64\x8b\x42\x99\x24\x14\x00\x94\x2d\x6b\xf8\xef\x53\xb8\x37\x28\xde\x32\x35\xe3\x17\x0b\xe7\x34\x4e\xe3\xd2\x6c\xdc\x04\x48\xde\x8a\x1c\xc8\xaa\xa0\x8a\x16\xbc\xcc\x64\x2e\x00\x9a\x27\x4e\x45\xb1\x22\xab\xe3\x91\x9c\x5f\x17\xa4\xeb\x56\xe4\x78\x46\x88\x62\xaa\x02\xf2\xc1\xe2\x0f\xa6\xd0\x75\xab\x33\x42\x8e\xc7\x77\x84\x6d\xc8\xf9\x3d\xd0\xe2\x5b\x53\x1d\x48\xd7\x9d\x11\x22\x80\x16\x19\xd7\xc5\x0f\x44\xd7\x48\x59\x5d\x07\x9a\x22\x16\xb4\xc0\x7a\x4b\x05\x38\x4c\xda\xdf\xb6\x2e\x26\x06\x6b\x3e\x40\xbd\xab\xa8\x82\x47\x2a\x18\x7d\xaa\x40\x46\xfa\x85\xa9\xed\xa4\x81\xa0\x4d\x09\xe4\xdc\x02\xca\xd9\x65\x7b\x67\x68\xba\x4e\x48\x43\x6b\xd0\xff\x5d\xf7\xff\xad\x8b\xb6\xf7\x84\xec\x04\x6c\xd8\x6b\xe0\xee\x6c\xd1\xb3\x05\x6c\x68\x5b\x29\xcf\x5e\xb9\xa2\x1e\xd5\x33\x42\x06\x7b\x34\x54\x48\x7a\xf3\x83\x15\x25\xa8\xb1\x3e\xbc\x18\xd6\x35\x5c\x1d\x76\x71\xd2\xf4\x6f\xdf\xae\x57\x8f\xfe\x37\x40\x07\x0f\xfd\x2f\x40\x61\x80\xdd\x7c\x5b\xcd\x93\x48\x78\x80\x57\x95\xd4\x09\x5d\xe8\x49\x5c\x54\xac\x6c\x12\x9d\x8c\x1a\x08\xab\x79\xa3\x79\xb9\x35\x7b\x4b\x5b\x95\x49\x8d\x60\x31\x67\x32\xa9\xf5\x15\x58\xb9\x55\x1e\xdd\xda\x92\x13\x09\xdc\xa4\xc2\x0f\x56\xa8\xad\x07\x5f\x4c\xc1\xd5\xf7\xcc\x74\x67\xec\xf8\xd9\x7e\xe8\xdf\xbe\x07\x0b\xc6\xf5\x92\x57\x5c\x78\x30\x37\x05\x57\xdb\x33\x33\xe3\x58\x03\xc6\xdc\xf7\x82\x50\xa5\x7f\xdb\x68\x22\xa4\x62\x7b\xc8\xe4\x8e\x86\x19\xbb\x61\x7b\x58\xeb\xb2\xf7\xd3\x8d\x79\x1b\x74\x5e\xc1\x15\x6c\x86\xfc\x27\x84\x62\x15\x64\x05\x6c\x42\x33\xbc\xc0\x23\x7b\xf3\x36\x84\xec\xd9\x9b\x6f\x95\xc5\x57\xc8\x1a\x79\x47\x63\xd7\x4a\xc5\xeb\xef\x0d\x53\x91\xcb\x0d\x96\xb5\x1a\xf4\x03\x89\xcd\x66\x45\x2f\x5a\xc5\x65\x4e\xe3\xf7\x42\x08\x0d\x90\x53\xc4\x36\xb3\x82\x3a\x0a\x92\x0f\xc7\x06\x49\xef\xcb\x41\x46\xb3\x8a\xf7\xf0\xb3\x05\xe9\x33\x48\x3a\xf4\xc3\x5c\x92\x61\xf4\x9f\xb0\x66\x61\x46\xa2\xf8\x7f\x5a\x10\x87\x68\x48\xc8\x4f\xdf\x46\xcf\xac\x92\x3a\x49\x0b\x51\xaf\x6d\xb6\x0a\xf8\x68\x26\x9b\xd1\x31\x4e\xfb\x62\x3f\x35\x98\x61\x49\x6c\xb6\x48\xf7\x16\x94\x60\x39\x66\x6a\x8b\x38\xc1\xc0\x2f\xeb\x2d\xbc\xaa\xcf\xac\x52\x20\x92\x3e\xeb\x69\xde\x58\x18\xcd\x73\x30\x5c\x24\x7d\xc3\x6a\x1c\xdf\xfa\x1b\xae\x63\x6c\x7b\x76\x91\xd4\x45\x59\x0a\x28\xa9\xe2\x49\x2b\x69\x44\x7d\x78\x63\xbb\x45\xca\x97\xbc\xde\x51\x01\x0f\x1c\x93\xb9\x05\x33\xc5\x63\x4a\x8b\x66\xcb\x74\xb7\x3a\x70\xfb\xd3\x9f\x1b\x34\x99\xff\xc4\x70\x91\xf4\x37\x51\x80\xf8\x98\x44\x3a\xd7\x50\xf6\x14\x16\xd2\x68\xb2\x5c\xf1\x8a\x89\x53\xc9\x82\x89\x44\xd3\x1a\x2d\x12\xfd\xf4\xaa\x04\xbd\xe4\x15\xe6\x40\x63\x59\xce\x2b\x2f\x8a\x8c\x16\x89\x5e\x37\xb9\x00\x2a\xe1\x0b\xe7\x09\xcf\x1c\x9e\x95\x9a\x70\xe2\x3d\xe3\x85\xf1\xd0\x14\x4c\x31\xde\xd0\xea\x33\x17\x35\xc5\xf9\x08\x67\xab\x01\xbb\x77\x3d\xc3\x93\xd4\xa5\x03\x2b\xd4\xca\x36\xa6\x1a\xca\x63\x23\x4b\xab\xaf\x39\xba\xc0\x8e\xf6\x29\x0a\xde\xd1\x0a\x94\x82\x94\xdc\x39\xd0\xef\x1f\x83\xcd\x62\x59\xfb\x55\xf4\x3f\x4b\xff\xfd\xe0\x0f\x33\xb1\x5c\xac\x7f\xdd\xec\x41\xa8\x94\x63\x16\x0b\x53\xec\x2c\x16\x6b\x3e\xd2\xaa\xed\x0d\xc4\xde\x40\x7e\x05\x77\xfc\xf2\x46\xd6\xb4\x84\xef\xf7\x37\xbd\x66\x6a\x34\x6b\x45\x88\x74\x64\x36\x23\xdd\x0d\xc4\x29\x99\x8b\xdd\x91\x70\x5e\xab\x43\x05\xc3\x11\x7c\x42\x49\x03\x0c\xc5\xe3\x3f\x14\x3e\xfd\xec\x38\xb3\xd6\xce\xa8\x25\x5b\x61\xfb\x37\xb1\x21\x1e\x1f\x7d\xf2\xfe\xbd\x19\x89\xbf\x38\xe2\x0f\xb4\xb4\xcb\x64\x92\x38\x14\x2d\xdd\x7a\x2a\xc9\x07\xf2\xdb\xf1\xe8\x92\x43\xb4\xee\xba\xd5\xf1\x78\xde\x75\xab\x7f\x1d\x8f\xd0\x14\x5d\xf7\xfb\xb8\x37\xdd\x46\xb7\x0f\x9a\xda\x71\xcd\x03\x26\x4f\xef\xa1\x49\xb2\x57\x0c\x91\x41\xea\x24\xad\x81\xb6\x42\x71\x33\xba\xed\xd2\xcd\x36\xc6\xbd\x26\xbd\x3b\x69\xe4\x30\xd2\x6b\xf8\x2d\x15\xcf\xc9\x30\xc7\x86\x0f\x52\x27\x0d\xaf\x8d\xd5\xc0\x4e\xf2\xef\xda\x02\xde\xd0\x27\x48\xd6\xbf\xca\x00\x7e\x07\xe4\xd8\x45\x52\x27\x69\x6b\x26\x69\x0d\xc6\xcc\xd0\xe0\x2f\x8e\x18\x2d\xe0\xcf\x43\x67\x03\xfc\xc4\x11\x33\x39\x2e\xeb\xad\x65\x72\x5a\x8e\x06\x93\x07\x46\x7c\x5a\xc2\x87\xa4\xe4\x78\x34\x52\xf7\x4e\x40\xce\x24\xe3\xe1\x24\xb3\x0b\x40\xbc\x3d\x09\x16\xb3\x47\xe6\xf4\x36\x61\xc1\x91\x68\x44\xeb\x33\x6f\x92\xb1\xd9\xf0\x26\x1d\x1b\x64\x30\x29\x74\x51\x81\x50\xd7\x57\x1e\xa6\xba\x98\xb1\xb0\x11\x8a\xf4\xb4\x4a\xab\xf8\x3d\x6c\x04\xc8\x90\x47\xf5\x01\x32\x13\x0e\x43\x67\xc8\x68\x36\xa9\x78\x03\x25\x42\x2b\x5b\xf2\xe1\xef\xb9\x05\x0a\x78\x90\xac\x4a\x32\x4c\x89\xd1\xa4\x5c\x72\x46\x34\x87\xb1\xe1\x64\x35\x16\x83\x71\x99\x47\x4b\x3b\x5e\xd5\xc7\x2e\x4b\xb6\x90\x3f\x87\xcb\x12\x53\x08\xfb\x7f\xcb\x4c\x56\xff\x22\x78\xbb\x63\x4d\xe9\xf1\xd2\x97\x9d\x08\xe2\xe7\x75\x12\x91\x44\x61\xfe\xce\x26\x7f\xbe\xe3\x32\xde\x8e\xe4\xcf\xd9\x8e\xcb\x78\xef\xe5\xe9\x59\x95\x4f\x45\x09\x89\x0c\x68\x00\xe9\x38\x83\xe9\xeb\xb3\x87\xdb\x30\x1d\x5b\x55\x87\xf9\x70\xf8\x6c\x23\x70\x03\xb0\xef\xb9\xba\x1f\xcb\x1c\x6f\xd0\x9f\xca\x64\x73\x1e\xd9\x49\x11\xb3\xb6\xe2\xb8\x36\xc9\x39\x09\x6b\x6c\x32\xa9\xb5\x66\x6f\x28\x38\xa4\x2d\x39\x91\xc0\x4d\x2a\xdc\x52\x51\xb2\x90\xd2\x6a\x5b\xf2\x17\x0a\x9e\x9b\xee\x4f\xb3\x0f\x3d\x69\xf6\xa1\x0b\x06\x9d\x6e\x3c\x88\x3d\xcb\xc1\xfd\x0b\x9d\xb0\xc5\xcc\xfd\x0f\xbd\xe9\x1b\x2f\x91\x76\xb7\xe6\x89\xae\xb9\x58\x4f\x45\x93\xcb\xf5\xf1\x91\x86\x47\x10\x78\x29\xd1\x33\x96\xed\x1d\x16\x07\x1d\x99\x4d\x27\x38\x7a\xe0\xad\xea\x69\x56\x06\xec\xab\xf6\x4d\xa7\x67\xb4\x95\x6a\xbd\xe5\x2f\x5f\x59\x3c\xc2\xd6\xad\x54\x99\xdc\xf2\x97\x6c\xab\x51\x3f\xbf\xa9\xe5\x22\xd5\x4f\x42\x70\x31\xa0\x0b\x16\xef\x29\x07\xeb\x45\xda\x37\x54\x41\x93\x1f\x4e\xc5\x2b\x47\xf4\xd4\xa3\xfd\x22\xf9\x8f\x02\xe8\x73\xc1\x5f\x9a\x53\x07\x4f\x81\xea\xb9\xc0\x75\x16\x39\xb9\x62\x52\x09\xf6\xd4\x2a\x34\xa9\xd1\x4f\x81\xd9\x9e\xab\x5e\xcd\x45\xde\xee\xdd\x8b\xda\x0d\x93\xea\xd4\x9b\x7f\x6f\xcb\x2a\x4d\xf7\xdc\xf5\xaa\x4e\xba\xbb\x62\x72\x57\xd1\x83\xbd\xef\xf0\x64\x61\x41\x7f\x9d\xe1\x9f\x9b\x7a\xa6\xf3\x2f\x09\x77\x02\x36\x20\xa0\x89\x29\xc0\xa4\xd1\x6c\x17\x71\x7c\xfb\x91\x98\x4f\xaf\x12\xac\x80\xff\x83\xe0\x97\xbc\x6d\xe2\xa7\xb0\x65\x05\x64\x6f\x20\x78\x96\x5b\xdc\xaf\x1d\x7d\xeb\x99\xb4\xd9\xd0\x12\xd6\x8a\xaa\x56\xea\xd1\x4c\x9e\xab\x6a\x43\x66\xd2\xb0\x76\x2a\x92\x17\xac\xb1\xca\x8b\x3d\xe2\x27\xb0\x01\x8f\xf6\x61\x0a\x3f\xeb\x8c\x55\xfe\x35\x8f\x78\xd1\x1a\xf2\x88\x97\xb0\xb1\xca\xbf\xe6\x31\xd9\x6f\x0f\xb9\x4c\xb6\xdf\xa3\xd5\xa7\x8f\x08\x54\xd0\x5a\x62\xd4\x1e\x25\x53\x7c\x67\x4b\xfd\xe7\xa0\x35\xc7\xd7\x54\x92\xc7\x0b\x2a\xc7\x2c\x7a\x6b\x89\xe8\xe8\x5b\xdc\x84\x82\x89\xd8\x08\x9b\xb0\x8e\x9f\x8c\xe5\x66\x45\xd6\x8a\x26\x3d\x31\x45\xdf\x15\xc7\x8d\x88\x8c\x9e\x02\x87\xbf\xf8\xb6\x6e\x24\xfa\xd2\x4d\xd1\xbe\xba\x7b\xf2\x0f\x02\x32\xa7\x3b\xb8\xdc\x52\x41\x73\x7c\xa3\x32\xb6\x94\xf2\x52\x42\x68\x7d\x65\x4b\x7e\xed\xf4\xdc\x40\x10\xe8\x93\xad\x7d\xbc\x9e\x78\x08\xef\xce\xfe\x0c\x00\x00\xff\xff\xd2\xc2\xd5\xc0\xb8\x20\x00\x00")

func tmplScreenboardTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplScreenboardTmpl,
		"tmpl/screenboard.tmpl",
	)
}

func tmplScreenboardTmpl() (*asset, error) {
	bytes, err := tmplScreenboardTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/screenboard.tmpl", size: 8376, mode: os.FileMode(0644), modTime: time.Unix(1597229994, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x18, 0xae, 0x4b, 0x1b, 0xd5, 0x4a, 0x7b, 0xf9, 0x3, 0x4b, 0x12, 0x2, 0xd6, 0xe0, 0x6, 0xfe, 0x8c, 0x15, 0x4c, 0xb9, 0x48, 0xaf, 0x52, 0x1e, 0x71, 0x5a, 0xfe, 0xf3, 0x82, 0x86, 0x77, 0xd3}}
	return a, nil
}

var _tmplTimeboardTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xcd\x6e\xdb\x38\x10\xbe\xfb\x29\x06\x42\x0e\xbb\x40\xec\x07\x58\x20\x87\x6c\x02\x07\x0b\x6c\xdb\x34\x09\xd2\x43\x51\x08\xb4\x34\x52\x88\x52\x3f\xa1\xa8\x24\x0e\xc1\x77\x2f\x38\xfc\x93\x6c\xc5\x3d\xd4\x27\xce\x37\xdf\xfc\x72\x86\x96\xc4\xa1\x1b\x65\x81\x90\x95\x4c\xb1\xb2\xab\x73\xc5\x1b\xdc\x75\x4c\x96\x19\x64\x5a\xc3\xe6\xbf\x12\x8c\xc9\x40\xaf\x00\x14\x57\x02\xc1\xfd\x2e\x9c\xf6\x81\x20\x63\xb2\x15\x40\x89\x43\x21\x79\xaf\x78\xd7\x06\xf5\xf5\x04\x72\x24\x89\xac\xcc\xbb\x56\xec\xc9\x87\xe5\xdc\x21\x2b\xbf\x58\x60\x6d\xcc\x0a\x40\xeb\x57\xae\x9e\x60\x73\x23\x59\xff\x34\x44\x50\xb2\xb6\x46\xd8\x00\x89\xb5\xd5\x51\x4a\x21\xa9\x85\x74\x92\xab\x6b\xac\x78\xcb\x29\x09\xe7\x0e\xe0\x85\xbf\xa7\x22\x1e\xf9\xbb\x55\x04\xa3\x35\xf0\x0a\x36\x97\xa3\xea\x86\x82\x09\xb4\x2a\x16\x05\x6f\x92\xb4\xc6\x64\xd6\x04\xdb\xd2\xbb\x0e\x0e\x6e\x25\x16\x7c\xf0\x41\xfb\x28\x78\x07\x49\xfb\xa1\x83\x1b\xd9\x8d\x3d\x75\xa0\xb6\x27\xb8\x80\xef\x5a\x9f\xd5\x0e\xfd\xe7\x22\x10\x8c\x09\xdd\x39\xe3\x6d\x89\x6f\xe7\x70\x86\x02\x9b\x03\x06\xaf\xbc\xda\x98\x73\xad\x29\x58\xa6\x35\x31\xe9\x44\xc8\x8f\xe5\x44\xee\x8b\xae\x47\x4a\x64\xb0\x27\x9f\xc8\xe0\x50\x1b\xc6\x11\x4e\x25\x92\x18\x7f\x94\x88\xda\xbb\xfb\x20\x74\x20\xc9\x4d\x01\x40\xcf\x04\x2a\x85\xb3\xe9\x24\xfe\xe6\xd6\x6b\xc2\x0d\x47\x6e\x5e\x09\xde\x2f\x72\xb7\x56\x11\xf8\x66\x39\x99\x4f\x4c\xfe\x44\x49\x6d\xb1\x90\x1b\xb5\x19\xe8\x67\xd6\xdb\x35\xa4\x8a\xf9\xaa\x7d\x8f\x69\x8d\xac\x90\xf2\x7b\x61\x62\x8c\xa3\xf6\x48\x42\x52\x6a\x4d\xe1\xff\x67\x3b\x14\x36\x8e\xa0\x83\x27\x3b\xf4\x68\xa4\x52\x09\xc7\x87\x58\x94\xab\xe0\x0e\x9f\x47\x1c\xd4\x62\x09\xd2\xe9\x62\x0d\xcf\x93\x5e\x7f\x1d\x51\xee\xd3\x1a\xc5\x3c\xa9\xb4\xb5\x31\x54\xef\x41\xb9\x5a\xdb\x14\xc0\x7b\x8f\x26\x97\x75\x2d\xb1\x66\xaa\x93\x2e\x09\x0b\xb6\x08\x59\x06\x7f\x5d\xe3\x1d\x56\xf7\x4a\xf2\xb6\x9e\xf2\xfe\xa6\x25\x4d\x66\x61\x4b\x13\x12\xa3\x91\x43\x1b\x75\x7d\x18\x36\x4e\x97\xd6\xbe\x17\x0e\x89\xbc\xf9\xc0\x59\x3b\xda\x73\x3f\x5e\x76\xcb\xfd\x31\xec\x78\x1a\xbc\x79\xb7\x93\xf1\x37\x5e\xaa\x27\x6b\xfa\x4a\x07\x6f\xe8\xd0\x13\x66\xa7\x9b\x3a\xb7\x99\x54\xe9\xcb\x0e\x8d\x88\x0a\x37\xd0\x57\x5d\x5b\xd2\x2b\xc9\xc4\xb6\x93\x0d\x53\x03\x4c\x47\xfb\x43\x75\x78\x9a\x53\x43\x8b\x44\xcd\x2b\xe2\xce\xba\x06\xd3\xb6\x9d\xee\xda\x7c\x3e\xac\xe7\xa6\x67\x72\x7a\xc7\x57\x09\x49\xb3\x97\x8a\x0a\xdb\xb3\xbc\x52\xc7\x01\x62\x33\xc6\x41\x75\xcd\xbf\xf5\x55\x27\xc8\x73\x41\x72\xbe\xab\xf3\x82\x90\x10\xfd\x80\xf6\x5b\x8f\xdb\x43\x8f\xd5\xa2\xc7\xed\xc7\x1e\xa7\xd7\xe6\x14\x4b\xa7\x55\xa2\x1e\x6f\xfa\xb2\x64\xfc\xe3\x3b\xc5\xd2\xf5\x3f\x60\xd3\x0b\xa6\xf0\x91\x49\xce\x76\x02\xe3\xc3\x37\xf9\x67\xb6\xdf\x08\x9e\x96\xbf\x78\x9e\xbf\xf9\x96\x35\x38\x79\x2d\x3e\x5b\x31\xdc\x57\x2f\xb1\xe2\x6f\x30\xf9\x6b\xb4\x62\xd0\x96\x58\xb1\x51\xa8\xf4\x49\xe1\x44\xfb\x51\xb2\x98\xb4\x59\xfd\x0a\x00\x00\xff\xff\xe4\xf8\xa0\x0f\xd5\x08\x00\x00")

func tmplTimeboardTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplTimeboardTmpl,
		"tmpl/timeboard.tmpl",
	)
}

func tmplTimeboardTmpl() (*asset, error) {
	bytes, err := tmplTimeboardTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/timeboard.tmpl", size: 2261, mode: os.FileMode(0644), modTime: time.Unix(1597229994, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4c, 0xea, 0xbe, 0xce, 0x31, 0x3c, 0x96, 0x4, 0x13, 0x38, 0x42, 0xae, 0x47, 0x3d, 0x5, 0xeb, 0x1d, 0xfd, 0x8b, 0xae, 0xfa, 0x12, 0xdb, 0x5f, 0xb4, 0xfa, 0x96, 0xf, 0x62, 0xec, 0xe8, 0x19}}
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
	"tmpl/dashboard.tmpl":   tmplDashboardTmpl,
	"tmpl/monitor.tmpl":     tmplMonitorTmpl,
	"tmpl/screenboard.tmpl": tmplScreenboardTmpl,
	"tmpl/timeboard.tmpl":   tmplTimeboardTmpl,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"dashboard.tmpl":   &bintree{tmplDashboardTmpl, map[string]*bintree{}},
		"monitor.tmpl":     &bintree{tmplMonitorTmpl, map[string]*bintree{}},
		"screenboard.tmpl": &bintree{tmplScreenboardTmpl, map[string]*bintree{}},
		"timeboard.tmpl":   &bintree{tmplTimeboardTmpl, map[string]*bintree{}},
	}},
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
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
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
