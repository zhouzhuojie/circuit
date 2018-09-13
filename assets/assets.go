// Code generated by go-bindata.
// sources:
// plugins/codeamp/graphql/schema.graphql
// plugins/codeamp/graphql/static/index.html
// DO NOT EDIT!

package assets

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

var _pluginsCodeampGraphqlSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\xdd\x72\xdb\xb8\x0e\xbe\xf7\x53\xd0\xd3\x1b\x77\x26\x4f\xe0\xcb\xc6\x3d\x49\xce\x69\xcf\x7a\xe3\xf6\x2a\x93\x0b\x46\x86\x6d\x6e\x24\x52\x25\x29\xb7\x9e\x9d\x7d\xf7\x1d\xfe\x0a\xfc\x91\x13\xa7\xdd\x99\xbd\x49\xc4\x4f\xe2\x47\x00\x04\x40\x80\x56\x0d\x6d\xa9\x24\x5f\x58\x07\x33\xff\xfc\xdf\xcd\x6f\xff\x9f\xcd\x54\x73\x80\x8e\x92\x3f\x67\x84\x7c\x1b\x40\x9e\x96\xe4\x77\xf3\x6f\x46\x48\x37\x68\xaa\x99\xe0\x4b\xf2\xd9\x3f\xcd\xfe\x9a\xcd\xde\xf9\xf7\xfa\xd4\x83\x7b\xb4\x73\xdf\x91\xaf\x0a\xe4\x8c\x90\x41\x81\x5c\xb0\xed\x92\xdc\xad\xde\x2f\x03\xe8\xde\x2a\xff\x5a\x2d\xde\x2f\xc9\x83\x41\x1e\xe7\xf6\xe5\x5a\x8a\x3f\xa0\xd1\x33\x42\x7a\xf7\xe4\x09\xae\x88\x6a\x87\xfd\x92\x6c\xb4\x64\x7c\x7f\x45\x38\xed\x60\x1c\x01\x3f\x32\x29\x78\x07\x5c\xdf\xad\x02\xfc\x7e\x89\xd8\x22\xb3\x1a\xa9\xd5\xc2\x3f\x6c\x80\xca\xe6\x10\x3f\x77\xc3\x3b\xde\x0f\xfa\x8a\xf4\x54\xd2\x4e\x2d\xc9\x9a\xee\x19\xa7\x5a\x48\x8b\x8f\xdc\x9f\x98\xd2\x4e\xf4\xff\x00\xd5\x83\x04\xb3\xc0\xce\x3f\x2e\x26\x67\xfb\x8f\xc7\xd9\x1b\x90\x47\xd6\xd8\xd9\xca\x3f\x4e\xcf\xf6\x1f\x17\xb3\x89\xea\xa1\x41\x14\x1b\x33\xb4\x26\xde\x8c\x80\xb7\xf4\x3d\xb4\x40\x95\x5d\x50\xfa\xc7\xe9\x05\xfd\xc7\xe3\x82\x1f\x47\x8b\x1b\x06\xb4\x01\xa3\x55\xd1\x86\x19\x11\xd0\x94\xc7\x82\x84\x1c\xa9\x64\xf4\xa9\xf5\x06\x68\x24\xe8\xb3\xfa\x9b\x0f\x90\xfa\x8c\xef\x5b\xf0\x70\x24\xb0\xae\x13\xd7\x8f\x2f\xa3\x2b\x7c\xfc\xa1\x81\x2b\x26\xb8\x35\x9b\xd1\x22\x00\x6a\x31\xe5\x51\x0f\x71\x52\xea\xb0\x11\x46\xfe\x35\x62\x76\x0b\xf2\x2f\xd3\x7d\x40\x8b\x8f\x3b\x92\x31\xdc\x67\x68\x10\x01\x64\xc7\x54\x5c\x7c\x1c\x99\x49\x26\xb4\xe7\x2e\x5a\x63\xec\xda\x80\x0d\x23\x1f\xb3\xd7\x12\xa8\x86\x20\xfa\x8c\x90\xc6\x02\x5e\xe8\xb0\xa7\xd1\xed\xb3\x28\x70\x81\xdd\x6f\x53\x8a\xc1\x02\x97\x50\x78\x29\xbc\xfa\x51\x0a\xaf\xf8\xc2\xe3\xd1\x1f\x33\xf7\x74\xbe\xa0\x45\x8f\x08\x94\x16\x7d\x98\xee\x52\xc9\x3c\x9b\xe0\xd7\xf4\x21\x13\xd7\xf4\x11\xb3\xf0\x78\x0c\xba\x2c\x06\xb1\xe6\x23\xe4\x34\xbf\x84\x62\x05\x2d\x24\x52\x6c\x2d\x70\x09\x45\xaa\x48\x70\xea\x44\x1b\x13\xff\x0b\x94\x1c\x22\x81\x19\x64\x9c\x1b\x37\x3f\x6a\x97\xf1\x26\x2a\xbe\x89\x37\x55\x39\xf0\x26\x7a\xbf\x89\xd7\xdb\x01\x85\x70\x34\x03\xca\x38\x38\xc4\x97\x38\x15\x05\xda\x8f\xc9\xfc\x68\x86\x94\xd6\x59\xe1\x67\x68\xbd\x15\x52\x5a\x67\x84\x9f\xa1\x2d\x8d\x10\x53\x2c\x72\x0a\x9b\x26\x5d\xb6\x0c\x19\x32\xcd\xb1\x13\x9a\x63\xae\xe0\x08\xaf\xe3\x2a\xd5\x8d\x5c\x04\xed\xfe\xeb\xc8\x82\x92\x79\x22\xf7\x9b\x1d\xe0\x45\xfc\x60\x49\x22\x18\x0d\x17\x80\x44\xd5\x9c\xd1\xef\xf3\x1b\x18\x83\xc2\x39\xa3\xdf\xe2\x37\x30\xe6\x5a\xe7\xc9\x7a\xe4\xcc\x0f\xa2\x65\x71\x5c\x65\x49\xf8\xbc\x31\xf2\x94\xfe\xcb\x16\x42\x36\xf2\x98\xb3\xce\x3f\xa4\x90\xa9\x3c\xf1\x39\x19\xf5\x32\x85\x28\x3a\x4d\x17\x43\x3a\x76\x45\x2c\x02\xc2\x62\x0f\xae\x3e\x98\xfb\x03\x39\x3d\x08\x93\xe2\xa8\x30\x61\xa5\x70\xc2\xd8\xa8\x08\x02\xe3\xb2\x08\x0c\x6b\x7f\x10\xe2\xb9\xa3\xf2\x19\x1d\xc3\x4f\x1e\x5a\x27\x15\xb5\x39\x06\x3f\x08\xd1\x02\xe5\x6e\xe6\x0d\x68\x72\xc3\x34\xb9\x16\x5d\xc7\xac\xa4\x7b\xd0\x37\x4c\xfb\x71\x90\xce\x94\x43\x77\xab\x79\x51\x74\x5b\x8c\xc3\xf7\xc8\x8a\xf9\x6d\xfd\x11\xcb\xb8\x19\xe3\x1a\xe4\x8e\x36\x30\x62\xb6\x0c\x69\xc4\x60\xd2\xdb\x1d\xd7\x7e\x0a\xaa\x3b\x5d\xd5\x82\x00\xc2\xba\xbe\x05\x6b\x90\x33\x34\xa6\x36\xd5\x92\x81\x1a\xeb\xa7\x47\x4f\x3e\x56\x91\x8e\x7b\x1c\x5f\x4e\xed\xe6\x8e\xcc\xb1\x3e\x0f\xd4\x11\x78\x0b\xb7\x9d\x1c\xc8\x51\xeb\xe0\xc8\x11\x70\x39\xb9\x9f\x1c\xc8\x51\x57\xe3\xc8\x11\x70\x39\xb9\x9f\x1c\xc8\x6d\x0b\x68\x59\xcd\x93\x9d\xe9\x7d\xd1\xcc\xea\x28\x6b\x43\xa5\x3d\x4f\x0b\xd9\x2c\xc2\x5c\xb6\xdb\x2e\x6d\x0f\x9b\x5a\x25\xb1\x48\xb6\x42\x5e\x80\x1a\xac\x03\xa5\xe8\x1e\xf0\xba\x26\xe8\xf1\xf8\x40\xd5\x21\x91\x8b\x4a\xe0\xfa\x36\x43\x25\xec\xf0\xb0\x26\x62\xa8\xd5\xb0\x47\xbc\x42\xc4\x46\x74\x1d\xe5\x5b\xcc\x8e\xbb\xdf\x79\xda\xee\x25\xe5\xd0\xbc\xd8\x9b\x5e\x48\x93\x54\x1e\x4c\x5b\xf0\x38\x4f\x3b\xb7\xa4\x88\x30\xef\x8c\xa0\x67\x94\x32\x59\xba\x6f\xc5\xc9\x7c\xbe\xd1\x92\x6a\xd8\x9f\x5c\xc3\x31\x23\xa4\x65\x47\xe0\xa0\xd4\x5a\x8a\x27\x88\xa8\x04\xba\x65\x25\xdc\x4b\x30\x65\xfb\xad\x10\xcf\x61\x3d\x67\x32\x5c\xd5\x58\xb3\xe1\x96\x31\x35\x5d\x6e\x93\x67\x38\xe1\x21\x53\x2b\xd8\xd1\xa1\xd5\x49\xd6\x6b\x44\x2b\xe4\x59\x15\xc3\x5d\x41\xe9\xcd\xb5\xee\x15\x27\x92\x4c\xbe\x4c\x9e\x23\x6d\x87\x74\x0f\x1b\x91\x5a\xbb\xe6\x0b\xce\x37\x4d\xf4\xd4\xf6\xe7\x08\x32\x86\x4b\xc8\x47\xe7\x77\xb8\x50\x97\xa9\x8d\x2f\xba\xd2\xdc\x8d\x2f\x18\x12\x07\x36\x6e\xf6\xc2\x4e\x34\xfd\x70\x0f\xdf\x06\x50\x3a\x43\x3f\xb1\x8e\x25\x58\x07\x9d\x90\xa7\xca\xc7\xee\x45\xf1\xbd\x36\x29\x82\xdb\x36\xf6\x46\xd2\x06\xd6\x20\x99\xd8\xbe\x14\x87\xa1\xf9\xc3\x07\xca\x2b\xe2\x30\xb5\x3d\x95\x9a\xed\xa8\x75\x0c\xd7\x62\x13\x72\x00\xba\xf5\x99\x27\x5e\xf0\x58\x29\x29\x6b\x6b\xb8\xd2\x54\x43\x9a\x43\xb2\xae\x7f\xaa\xe7\xb7\x33\x3f\x97\xa9\xeb\xa2\xad\x56\x9a\xca\x04\xd8\x31\xce\xd4\x21\xf5\x86\x7b\xd1\xb6\x4f\xb4\x79\x2e\xce\x72\x5f\x5d\xe0\x33\xe2\x05\x37\xc0\x17\x78\x4e\xdb\x5e\x28\xa6\x85\x3c\xa5\xa9\xcc\xd7\xfc\x11\xd9\x33\xfd\x55\xb6\x19\xb2\x96\x42\x8b\x46\x24\xb0\x54\x74\x2d\xd9\x91\x6a\xf8\x5f\x1a\x6b\xe6\xc5\xf0\xd4\xb2\x26\xc3\xe3\x35\x9d\x3a\x88\xef\x2b\x9b\xcb\x8c\xf6\x5e\xd3\x33\x77\x7f\xd9\xed\x5d\x33\x48\x73\x24\xdc\x67\xb7\x13\x6f\xb9\x59\xbb\xf8\xee\xef\xc2\xbb\x32\xc0\xbe\x55\xbd\x91\xda\x33\xfd\x41\x52\xde\x24\xa7\x5b\x23\xb8\x66\x7c\x10\x83\x72\x66\x4a\x92\x28\x24\xf5\x6a\x59\x94\x86\xfa\x13\xd9\x76\x2a\x34\xb3\x2b\x39\x97\xf2\x23\xf6\x42\x9a\x11\x5d\x2f\xb8\x75\x7d\x94\x21\xb2\x14\x99\x65\xe1\xb3\x01\x23\xf8\x8e\xed\xc7\xf0\xae\xc9\x5b\x74\x19\x38\x1e\xa6\x04\xaf\x25\x97\x5a\xf7\x37\x91\x66\x0a\xb9\x06\xa5\x45\x77\x9d\xa1\x45\x72\xf9\x05\x39\x03\x67\x4f\xd4\x15\xe2\x3c\x3a\xa5\x73\x7e\x73\x97\xe9\x9c\x5b\x0c\x87\xc2\x86\xed\xb9\x4f\x9e\x79\x9a\x98\x78\x97\xab\x9e\x3b\xc1\x94\x29\x2a\xc6\x4e\x93\x24\x21\xd5\x34\x79\xc6\x33\x08\x33\x61\x38\xb3\x7f\x93\x8b\xcf\x68\x1f\x5f\xee\x4c\x66\xb5\x6a\xfa\x4b\x23\xb4\x1e\x63\xe9\xde\x8e\xb7\xd8\xe7\xc2\x39\x91\xdd\xfd\x0e\x52\xd3\x00\xfd\x42\x62\xf5\x28\x53\xf9\x54\xd8\xe3\x46\xd0\x78\x49\x42\x9e\xa4\x2e\x4b\xdc\xdb\x1d\xba\xb3\x77\x5a\xad\x3b\xfd\xcd\x20\x71\x44\x4c\x81\xef\x85\x4b\x0b\xa3\x03\x7a\x34\x07\x8a\xc8\x14\xac\x1a\xcf\x1e\x1b\x42\x36\x70\x0f\x4f\x03\x6b\x0b\xd5\x42\x9d\x84\x85\xc2\x57\xb6\xa5\x50\xd5\xb5\x2f\xab\xf9\xf3\xb9\xb5\x82\xdf\x0b\xb1\x16\xd2\xf9\xdf\xfc\xb1\x12\x18\x93\x2a\xd7\x0a\xfd\x55\x81\x39\xe2\xb2\xf8\xbf\x05\xda\xea\x83\x1d\xd8\x4f\x2a\x8d\x40\xe5\x93\xc9\xa6\xe0\x5a\x70\x4d\x19\x07\x69\x95\xab\x99\x3a\x6a\xe9\x9c\x48\x48\x64\x8f\x18\x62\xe1\x96\xc1\xcd\x9c\xd0\xc6\x12\xe4\x66\xea\xe8\x8f\xcd\x20\xbd\x67\x7a\xe0\x2b\xa7\x47\xca\x5a\xd3\x04\xe4\xd4\xb9\x6e\x05\xa7\x2d\x70\xf5\x21\xab\x5c\x53\x17\xc0\x7a\xd8\x6e\xe1\x00\x1d\x26\xe8\xa9\xc6\xf9\x80\x71\xa6\x19\x6d\x57\xd0\xd2\xd3\x06\x1a\xc1\xb7\x2a\x4c\xed\x6d\x91\x9c\x81\x9a\x75\x20\x06\x9d\xa1\x6a\x68\x1a\x50\xea\xcb\x41\x82\x3a\x08\xe3\xed\x0e\xdf\x51\xd6\x0e\x12\x0a\xfc\xa0\x75\x7f\x0b\x74\x0b\xd2\xb8\x1c\xd2\xfb\x36\xbe\x08\xce\x57\xb3\x4e\xf6\x95\xb5\x53\xee\xf6\x59\xdf\x54\x34\x27\x35\x77\x88\x3f\x16\x94\xd1\xf7\xef\xe9\x55\x8a\x9e\x12\x2b\x92\x5f\xf8\xbf\xac\xc8\x5b\xda\xdf\xc9\xb6\x36\x33\x6a\xbc\x8d\x2f\xc5\x78\xa9\xcb\x2d\x8e\xe2\x89\xae\x37\x39\xae\x26\xd3\xd2\x54\xb7\x5a\xfd\x81\x37\x35\x68\x72\x5f\x7c\xb1\x39\xab\xb5\xe6\xa4\x98\x79\xc5\x96\x1a\xa1\x2e\x71\xe5\xd4\x3d\x2b\xf3\x55\xf5\x28\xb9\xc2\x65\x56\x86\xbf\xae\x8e\x9c\x50\x2a\xde\xe7\x8d\xb7\xe1\x89\xc8\xb5\x8b\x72\x2b\xb1\xe9\xa2\xb3\xc3\x37\xb9\xe7\x4b\x27\xfa\x6c\x71\xe9\x7a\xe3\x72\x85\x03\xee\x25\xe5\x93\xfe\x32\x11\x7d\xe5\xeb\x71\x81\xc9\x4d\x7f\xe5\x42\xce\x34\x93\x0b\x8d\x96\xab\xd5\x09\xa9\xe9\x26\xc4\x74\xe6\xfb\x3b\x00\x00\xff\xff\x36\xb6\xea\x83\xf6\x23\x00\x00")

func pluginsCodeampGraphqlSchemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsCodeampGraphqlSchemaGraphql,
		"plugins/codeamp/graphql/schema.graphql",
	)
}

func pluginsCodeampGraphqlSchemaGraphql() (*asset, error) {
	bytes, err := pluginsCodeampGraphqlSchemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/codeamp/graphql/schema.graphql", size: 9206, mode: os.FileMode(420), modTime: time.Unix(1536860044, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsCodeampGraphqlStaticIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\xeb\x6f\xdb\x36\x10\xff\xde\xbf\xe2\xe6\x6c\x90\x5d\xd8\x94\xd3\xf5\x01\xa8\x76\x86\xb6\x49\xbb\x16\x69\xd3\xc6\x19\x8a\x7d\x2b\x4d\x9e\x2c\xa6\x14\xa9\x1e\x29\x3b\x6a\x91\xff\x7d\xa0\x64\xc9\x8a\x91\x0c\xd9\xb0\x41\x1f\x4c\xde\xe3\x77\x0f\xde\xc3\xb3\x9f\x8e\xcf\x5e\x5d\xfc\xf9\xf1\x04\x32\x9f\xeb\xa3\x07\xb3\xe6\x07\x60\x96\x21\x97\xe1\x00\x30\x73\xbe\xd2\xd8\x9c\x01\x96\x56\x56\xf0\x63\x7b\x01\xc8\x50\xad\x32\x9f\xc0\xe1\x74\xfa\xcb\xf3\x8e\x9a\x73\x5a\x29\x93\xc0\x74\x47\xda\x28\xe9\xb3\x7d\x39\xbb\x46\x4a\xb5\xdd\x24\x90\x29\x29\xd1\xb4\x9c\xeb\xed\xef\xc1\x8a\x78\x91\xa9\x6f\xfa\x76\x8b\xeb\x6c\x5f\x81\x5d\x6e\xfc\xc4\xdb\xaf\x68\x7a\x1a\x4b\x2e\xbe\xae\xc8\x96\x46\x26\xa0\x95\x41\x4e\x93\x15\x71\xa9\xd0\xf8\xe1\x41\xfa\x2c\x7c\x63\x38\xc0\x47\xe1\x1b\xed\x9c\x5b\x5a\x92\x48\x93\xa5\xf5\xde\xe6\x09\x1c\x16\x57\xe0\xac\x56\x12\x0e\xe4\x34\x7c\x3b\xc9\xd4\x1a\x3f\x49\x79\xae\x74\x95\x80\xab\x9c\xc7\x7c\x0c\x13\x5e\x14\x1a\x27\xed\x35\x5a\x70\x03\xaf\x89\x1b\xa1\x9c\xb0\xd1\x18\x22\xb6\x78\xfd\x61\x71\xac\x5c\xa1\x79\x35\x39\xc7\x55\xa9\x39\x05\xfa\x02\x57\x16\xe1\x8f\xb7\xd1\x18\xea\x63\x47\xfa\xfc\x31\xb0\x7f\x47\xbd\x46\xaf\x04\x87\x0f\x58\x62\x34\x86\xac\x25\x8c\x21\x3a\x2d\x85\x92\x1c\xde\x10\x37\x32\xf0\x38\x29\xae\xc7\xe0\xb8\x71\x13\x87\xa4\xd2\x9d\xd3\x05\x97\x52\x99\x55\x02\xcf\x8a\x2b\x38\x7c\x5c\x5c\xc1\xd3\xe2\x6a\x2f\x26\xa7\xbe\x63\x52\x33\x6f\x26\x7a\x16\xf7\x6a\x62\xa6\x95\xf9\x0a\x84\x7a\x3e\xa8\xa9\x2e\x43\xf4\x03\xc8\x08\xd3\xf9\x20\xf3\xbe\x70\x49\x1c\x0b\x69\x2e\x1d\x13\xda\x96\x32\xd5\x9c\x90\x09\x9b\xc7\xfc\x92\x5f\xc5\x5a\x2d\x5d\xdc\xbe\x73\x3c\x65\x87\x53\xf6\xa8\xbb\x33\xe1\xdc\x00\xe2\xdb\x0a\x31\x7e\x08\x67\x6b\x24\x52\x12\x1d\x3c\x8c\xdb\x02\x68\x35\x27\xc2\x1a\xcf\x95\x41\x02\xb6\x0e\x69\x58\x6a\x9c\xa0\x54\xde\xd2\x2d\xc5\xf4\xf4\xe9\xdf\x87\xe8\x04\xa9\xc2\x83\x23\x71\xef\x90\x52\xf4\x22\x8b\x1f\xb1\x29\xfb\xb5\x39\xb3\x5c\x19\x76\xe9\x06\x47\xb3\xb8\x81\xfb\xf7\xd8\x84\x5c\xf8\xf8\xf0\x09\x7b\xc2\x1e\x37\x97\xff\x15\x7c\x22\x6d\xfe\x1f\x1a\xb8\xf3\xb1\xf7\xe1\x67\x71\x3b\x85\x66\x61\xec\x6c\x2d\x4a\xb5\x06\xa1\xb9\x73\xf3\x41\xd7\xed\x83\xa3\x77\x9f\x2f\xe0\xa2\x6e\xfc\x99\x32\x45\xe9\x41\xc9\x3e\x1f\x0a\xcd\x05\x66\x56\x4b\xa4\xf9\x60\x27\xbc\xb2\xe8\x20\x43\xc2\x60\x59\xaa\x75\xcf\x46\x00\x68\x5d\x1b\x1c\x9d\x5a\x1e\xda\x85\x31\xd6\x97\xeb\xa7\x62\xcd\x09\x1c\x72\x12\x19\xcc\x61\xa3\x8c\xb4\x1b\xa6\xad\xe0\x5e\x59\xc3\x1a\xc6\xf3\x4e\xb0\xe0\xc4\x73\xf4\x48\x0e\xe6\xf0\xe3\xba\x61\x48\x2b\xca\x1c\x8d\x67\x2b\xf4\x27\x1a\xc3\xf1\x65\xf5\x56\x0e\xa3\x2e\x8e\x68\xc4\xd6\x5c\x97\x08\x73\x08\xd0\x7a\xe1\x2d\xf1\x15\x06\x85\xb7\x1e\xf3\x61\xd4\x3a\x9c\x5c\x6e\xfc\x45\xa3\xf1\xfc\x41\x0d\x9e\x96\x46\x04\x57\xa0\x16\xf9\x74\xfa\x3a\xd4\x24\xd2\x70\x7b\xfd\x18\x1c\x72\xa3\xae\x37\x84\x35\xce\x43\x8b\x02\xf3\x7f\xe0\x5c\xdb\x48\x1a\x3d\x84\x07\xdc\x06\xd9\x35\x5d\xf4\x42\x08\x2c\x7c\x94\x40\x14\x66\xa4\x6a\x52\x14\x5f\x3a\x6b\xa2\xf1\x4e\xea\x95\x35\x1e\x8d\x9f\x5c\x54\x05\xde\x2a\xdb\xf6\x6b\x6b\x4f\xa5\x30\x6c\x1d\x1e\xf5\xec\xdd\xc8\x94\xbb\x3b\x53\xe3\x2e\xdc\xde\x16\xb8\x2d\x80\xfb\x86\x70\x9f\x20\x6e\x48\xbf\x28\x7d\x66\x49\x7d\xaf\xf9\x51\xb2\x4b\xff\x6f\xf0\xe5\x25\x72\x42\x82\x9f\x7f\xb4\xc4\xeb\x2f\x90\x80\x29\xb5\xee\x10\xae\xf7\xd7\x21\xa1\x2f\xc9\x40\x3d\x7e\x86\x51\xfc\xad\x44\xaa\xa2\x71\x2f\x92\x1c\x7d\x66\x65\x02\x51\x61\x9d\xef\xf9\xb2\x8d\x7a\xdc\x5b\x86\xb2\x4a\xe0\xdd\xe2\xec\x03\x73\x9e\x94\x59\xa9\xb4\xda\x2b\x9d\x9d\xb0\x20\x94\x68\xbc\xe2\xda\x25\x10\x29\x23\x74\x19\xb6\x51\xeb\xdc\x88\xf9\x0c\xcd\xb0\x2b\xc8\x21\xa1\x2b\xac\x71\xd8\x7f\xb4\xad\xeb\x2d\x8b\x79\xbc\xf2\xc3\xee\x5d\xee\xc6\x78\x69\x65\xd5\xc7\xf1\x54\xdd\x78\xb9\x2d\x6e\x1d\x49\xc1\xc9\xe1\x4d\xcd\xdd\xc3\x5f\x83\xe0\x5e\x64\x30\x44\x22\x4b\xa3\xdb\x40\xfa\x9a\x3d\xc5\xce\xc7\x86\xd6\xdc\xe3\x18\xce\xd1\x48\x24\x98\xbd\xa9\x2b\xef\xd3\x29\xc4\x47\xa0\x8c\xb7\xe0\x33\xac\x13\xcc\x5a\xc9\x05\x62\x4d\x3c\x3f\x79\x71\xfc\xfe\x04\x94\xa9\x6f\xde\x16\xa0\x71\x8d\x1a\x6c\x0a\x3e\x53\x0e\x72\x2b\x4b\x1d\x18\xa0\x91\x93\x81\xdc\x12\x02\x5f\xda\xd2\xb7\x48\x99\xdd\x40\x65\x4b\x10\xdc\x80\x28\x9d\xb7\xb9\xfa\x8e\xd0\x79\xb0\xac\xa0\x20\xbb\x56\x61\xb4\x81\x54\x69\x8a\x84\xc6\x43\xdd\xc6\x0e\x2c\xb5\x30\xe1\xbf\x42\xc8\x33\xd7\x20\x32\xa5\x25\x60\x33\x01\x5c\xe3\xf2\x79\xd8\x12\xc7\x67\xef\x19\xd5\x21\x0e\xb7\x19\xa8\xc9\x4c\x10\x72\x8f\xdb\x91\x31\x6c\x4d\xf7\xab\x30\x6d\x46\x51\xb2\x37\x9a\x76\x05\xd3\x9e\xee\x9c\x41\x6d\x33\x47\xa3\x5a\x72\x9b\xf9\x9b\xbb\xa4\x59\x21\xb3\xb8\xf9\x8b\xfb\x57\x00\x00\x00\xff\xff\x45\xab\x31\x54\xfa\x0a\x00\x00")

func pluginsCodeampGraphqlStaticIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsCodeampGraphqlStaticIndexHtml,
		"plugins/codeamp/graphql/static/index.html",
	)
}

func pluginsCodeampGraphqlStaticIndexHtml() (*asset, error) {
	bytes, err := pluginsCodeampGraphqlStaticIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/codeamp/graphql/static/index.html", size: 2810, mode: os.FileMode(420), modTime: time.Unix(1535585208, 0)}
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
	"plugins/codeamp/graphql/schema.graphql": pluginsCodeampGraphqlSchemaGraphql,
	"plugins/codeamp/graphql/static/index.html": pluginsCodeampGraphqlStaticIndexHtml,
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
	"plugins": &bintree{nil, map[string]*bintree{
		"codeamp": &bintree{nil, map[string]*bintree{
			"graphql": &bintree{nil, map[string]*bintree{
				"schema.graphql": &bintree{pluginsCodeampGraphqlSchemaGraphql, map[string]*bintree{}},
				"static": &bintree{nil, map[string]*bintree{
					"index.html": &bintree{pluginsCodeampGraphqlStaticIndexHtml, map[string]*bintree{}},
				}},
			}},
		}},
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

