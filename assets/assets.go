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

var _pluginsCodeampGraphqlSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\xcd\x72\x23\xb9\x0d\xbe\xeb\x29\xe8\xda\x8b\xa6\xca\x4f\xa0\x5b\x66\xec\x8c\x9d\xcc\x24\x8e\xb5\x73\x48\x6d\xf9\x40\xb7\x60\x89\x71\x37\xd9\x4b\xb2\x3d\xa3\x4a\xe5\xdd\x53\xfc\x6d\x80\x64\xcb\x96\x77\x52\x95\x8b\x2d\xa2\x9b\x1f\x09\x10\xf8\x08\xa0\x4d\xc7\x7b\xae\xd9\xaf\x62\x80\x55\xfc\xfd\x97\xed\xdf\xff\xb6\x5a\x99\xee\x00\x03\x67\xff\x5e\x31\xf6\xfb\x04\xfa\xb8\x61\xff\x70\xff\x56\x8c\x0d\x93\xe5\x56\x28\xb9\x61\x5f\xe3\xaf\xd5\x7f\x56\xab\x5f\xe2\x73\x7b\x1c\x21\xfc\xf4\x73\x7f\x61\xdf\x0c\xe8\x15\x63\x93\x01\xbd\x16\xbb\x0d\xbb\xbd\xfa\xb0\x49\xc2\xf0\xd4\xc4\xc7\x66\xfd\x61\xc3\x7e\x73\x92\x87\x0b\xff\xf0\x4e\xab\x7f\x41\x67\x57\x8c\x8d\xe1\x57\x04\xb8\x64\xa6\x9f\xf6\x1b\xb6\xb5\x5a\xc8\xfd\x25\x93\x7c\x80\x79\x04\xf2\x45\x68\x25\x07\x90\xf6\xf6\x2a\x89\x3f\x6c\x10\x5a\x46\x36\x33\xb4\x59\xc7\x1f\x5b\xe0\xba\x3b\xe4\xd7\xc3\xf0\x56\x8e\x93\xbd\x64\x23\xd7\x7c\x30\x1b\x76\xc7\xf7\x42\x72\xab\xb4\x97\xcf\xd8\x5f\x84\xb1\x61\xeb\x7f\x06\x6e\x27\x0d\x6e\x81\xa7\xf8\x73\xbd\x38\x3b\xbe\x3c\xcf\xde\x82\x7e\x11\x9d\x9f\x6d\xe2\xcf\xe5\xd9\xf1\xe5\x6a\x36\x33\x23\x74\x08\x62\xeb\x86\xde\xc4\xdb\x59\x10\x2d\x7d\x0f\x3d\x70\xe3\x17\xd4\xf1\xe7\xf2\x82\xf1\xe5\x79\xc1\xeb\xd9\xe2\x0e\x01\x1d\xc0\x6c\x55\x74\x60\x6e\x0b\x68\xca\x43\x05\xc2\x5e\xb8\x16\xfc\xb1\x8f\x06\xe8\x34\xd8\x93\xfa\xbb\x17\x90\xfa\x42\xee\x7b\x88\xe2\x0c\xe0\x5d\x27\xaf\x9f\x1f\x66\x57\xb8\xfe\x61\x41\x1a\xa1\xa4\x37\x9b\xd3\x22\x09\xcc\x7a\xc9\xa3\x7e\xcb\x93\xa8\xc3\x66\x31\xf2\xaf\x59\xe6\x8f\xa0\x7c\x93\x9e\x03\x5a\x7c\x3e\x91\x02\xe1\xbe\x90\xa6\x2d\x80\x1e\x84\xc9\x8b\xcf\x23\x37\xc9\x85\xf6\x45\x88\xd6\x1c\xbb\x3e\x60\xd3\x28\xc6\xec\x27\x0d\xdc\x42\xda\xfa\x8a\xb1\xce\x0b\xe2\xa6\xd3\x99\x66\xb7\x2f\xa2\x20\x04\xf6\xb8\xa3\x10\x93\x17\x9c\x03\x11\x77\x11\xd5\xcf\xbb\x88\x8a\xaf\xa3\x3c\xfb\x63\xe1\x9e\xc1\x17\xac\x1a\x11\x80\xb1\x6a\x4c\xd3\x03\x95\x5c\x14\x13\xe2\x9a\x31\x64\xf2\x9a\x31\x62\xd6\x51\x9e\x83\xae\x88\x41\xac\xf9\x2c\x0a\x9a\x9f\x03\x71\x05\x3d\x90\x5d\xec\xbc\xe0\x1c\x88\xdb\x61\x54\xda\xb2\x81\xcb\x63\xe6\x10\xc6\x2d\x53\xd2\xbf\x20\xfc\xe3\xc4\x33\x09\xd1\x6c\xe2\xbc\xf4\x20\x21\x27\xca\xb8\x88\x5e\x46\xcd\x94\x42\x86\xd8\xca\xb1\xcb\x1a\x51\x4f\xde\x9e\x1b\x14\x3b\xde\x86\xf9\xd9\x76\x05\x2e\x31\xe0\xbb\x70\xa9\x41\x13\x2e\xb1\xea\xbb\x70\xa3\x1d\x10\x41\x64\x33\x20\x3e\xc3\x04\xb2\xc1\x44\x97\x60\xaf\xc9\xfc\x6c\x06\x0a\x1b\xac\xf0\x47\x60\xa3\x15\x28\x6c\x30\xc2\x1f\x81\xad\x8d\x90\x09\x1c\x39\x85\x27\xe1\xc0\xc5\x89\x7f\x29\x83\x2f\x68\x8e\xb1\x92\x23\xbc\x0d\xab\x56\x37\x63\x31\x74\xfa\x6f\x03\xc3\xf1\xd4\x42\x6c\x05\x57\xb8\xb8\xe2\x05\x36\x87\x96\x1f\xa2\xc8\x72\xe3\x14\x58\xd7\x3f\xfc\x32\xbc\xef\x17\x56\xf9\x2e\xec\x41\x48\xb6\x17\x2f\x20\x13\xbb\xde\x5e\x31\x2e\x77\x34\xf1\xf1\x37\x18\xde\x46\xba\x3f\xaf\x7f\x34\x76\x11\xae\x34\x12\xdb\x50\x5e\x88\xd1\xad\x93\x78\x9d\x5f\x70\x98\xf1\x67\x76\x91\x24\x20\x87\x5a\x22\x46\x8f\x7e\x07\x62\x3a\xda\x12\x31\x3a\xf3\x3b\x10\x4b\xad\xcb\x4b\x6f\xc6\x2c\x2f\xf4\x4d\x75\xed\x17\x97\xd9\x69\x63\x94\x57\xe3\x4f\x5b\x08\xd9\x28\xca\x82\x75\xfe\x47\x0a\xb9\x0c\x1e\xe7\x1b\x59\x2f\x97\xd0\xa3\xac\x64\x3d\xd1\x71\x28\x06\x90\x60\x8e\x8c\xe0\x94\x31\x32\x68\x42\x41\x92\xcc\xca\x84\x8d\x04\x14\xcb\x66\x45\x90\x30\x2f\x8b\x84\x69\xed\x8f\x4a\x3d\x0f\x5c\x3f\xa3\x74\xe6\x31\x8a\xee\x48\x65\xe2\xd2\x89\x8f\x4a\xf5\xc0\x65\x98\xf9\x19\x2c\xfb\x2c\x2c\xfb\xa4\x86\x41\xf8\x9d\xee\xc1\x7e\x16\x36\x8e\xd7\x39\x82\xfd\xec\xaa\x78\xf1\x32\x09\xdf\x33\x2a\xc6\xf7\x79\x5c\x4e\x87\x57\x42\x5a\xd0\x4f\xbc\x83\x59\xe6\xd3\xb9\x4e\x4d\x8e\xc8\x6f\xa5\x8d\x53\x50\xfe\x1e\xb2\x3f\x24\x70\xdc\xd5\x83\x37\xc8\x09\x18\x97\xe3\x5b\x2d\x5c\xbe\x90\xf2\xd0\x87\x08\x3e\x67\xe3\x01\x7b\x1e\x9f\x0f\x1d\xe6\xce\xc8\xb9\xce\x49\xd0\x59\xf0\x1e\x6c\x3f\x39\x81\xa3\x12\x2c\x80\x23\xc1\xf9\xe0\x71\x72\x02\x47\xd5\x61\x00\x47\x82\xf3\xc1\xe3\xe4\x04\xee\x4b\x69\x8f\xea\x7e\xf9\x99\xd1\x17\xdd\xac\x81\x8b\x1e\xd3\xfb\x88\x03\x8f\x44\x58\x60\xbb\xdd\xc6\xf7\x02\xa8\x55\x88\x45\x8a\x15\xca\x44\xde\xc9\x06\x30\x86\xef\x01\xaf\xeb\x82\x1e\x8f\x0f\xdc\x1c\xc8\xbe\xb8\x06\x69\x6f\x0a\xa9\x86\x27\x3c\x6c\x6d\x31\xe5\xbc\xd8\x23\xde\xb0\xc5\x4e\x0d\x03\x97\x3b\x8c\x8e\xbb\x08\x17\xb4\x6c\x26\x89\xdf\x45\x75\x36\xee\x42\x75\xe6\x74\xe5\xd5\xc3\x05\xad\x80\x49\xba\xe4\x9e\xb9\x8d\x9e\x50\xca\xb1\xf4\xd8\xab\xa3\x7b\x7d\x6b\x35\xb7\xb0\x3f\x86\xc2\x6d\xc5\x58\xef\x6e\x7d\x30\xe6\x4e\xab\x47\xc8\x52\x0d\x7c\x27\x6a\xf1\xa8\xc1\x95\x3f\x37\x4a\x3d\xa7\xf5\x82\xc9\x70\xfe\xe6\xcd\x86\x4b\x6f\x6a\xba\xd2\x26\xcf\x70\xc4\x43\x61\xae\xe0\x89\x4f\xbd\x25\xac\xd7\xa9\x5e\xe9\x93\x2a\xa6\x9e\x4b\xed\xcd\xad\x2e\x00\x26\x92\x62\x7f\xc5\x7e\x5e\x78\x3f\xd1\x33\xec\x14\xb5\x76\xcb\x17\x82\x6f\xba\xe8\x69\x9d\xcf\x0b\xe8\x1c\x2e\x89\x8f\x4e\x9f\x70\xa5\xae\x30\xdb\x98\x5e\x52\xee\xc6\x8d\x1a\xe2\xc0\xce\xcd\x5e\x39\x89\x6e\x9c\xee\xe1\xf7\x09\x8c\x2d\xa4\x5f\xc4\x20\x88\x6c\x80\x41\xe9\x63\xe3\xe5\xf0\xa0\x7a\xdf\x3a\x8a\x90\xbe\x1d\xf0\x59\xf3\x0e\xee\x40\x0b\xb5\x6b\x44\x46\x8e\x8a\x05\xa5\x6b\xdf\xc0\xd7\x0f\xb9\x7a\xde\x10\xb1\xf4\x94\xb8\xb6\xe2\x89\x7b\x17\x0a\x4d\x0d\xc6\x0e\xc0\x77\x91\xa3\x72\x4b\xcd\xeb\xc3\x45\xdf\x92\x1b\xcb\x2d\x50\xb6\x29\xfa\x2c\x4b\x5d\x16\x3f\xf3\x6b\x4d\x72\x67\x39\x85\xb1\x5c\x13\xc1\x93\x90\xc2\x1c\xa8\x09\xef\x55\xdf\x3f\xf2\xee\xb9\xba\xf5\x63\x1e\x82\x6f\x93\x57\x1c\x06\xb7\x4c\x83\xb6\xa3\x32\xc2\x2a\x7d\xa4\x47\x1b\xeb\xa0\x2c\xd9\x0b\xfb\x4d\xf7\x85\xe4\x4e\x2b\xab\x3a\x45\xc4\xda\xf0\x3b\x2d\x5e\xb8\x85\xbf\xd2\xa8\x74\x0f\xa6\xc7\x5e\x74\x85\x3c\x37\x46\xcd\x41\x7d\xbf\xf2\xac\xe7\xb4\x8f\x9a\x9e\xe8\xb6\x16\xfd\xd2\x6e\xd2\xee\xf2\xb8\x2f\xfa\x41\xef\xe9\x65\xbe\xd2\x6d\xbd\x64\xc6\xb7\x83\x91\x22\x75\x03\xf6\x74\xc3\x72\x09\x02\xf7\x30\x01\x7b\x60\xb3\x53\xb8\x17\xf6\xa3\xe6\xb2\x23\xb7\x65\xa7\xa4\x15\x72\x52\x93\x09\xc6\x24\xa4\x0c\x24\xff\xad\x93\xdc\x94\xcf\xa2\x13\x58\xba\x72\x8b\x56\x69\xb8\x42\xb2\xec\x15\xda\x52\xc3\xa8\xa4\x0f\x10\xc4\x38\xe5\x95\xc8\xbb\x03\x38\xe6\x27\x5b\x29\xb8\xfe\x64\xb0\x29\xf9\x24\xf6\x33\x35\xb4\xb4\xa8\x6a\x19\x1c\x4b\x4b\xea\xb4\x88\xa9\x55\x63\x2e\x50\x54\xb5\xaf\xc9\x58\x35\x7c\x2a\xa4\x15\x31\xfd\x04\xbe\xc1\xcc\x8b\x6a\x4f\xcc\xc1\x4b\x3a\x97\x7d\xd6\x42\xe7\xd2\x62\x38\x8c\xb6\x62\x2f\x23\xf1\x96\x14\xb3\xf0\xac\x54\xbd\x74\x8d\x25\x53\x34\x8c\x4d\x09\x96\xb1\x26\xc5\x9e\xf0\x0c\x26\x5c\xbc\xae\xfc\x5f\xd2\xa6\xce\xf6\x89\x49\xd5\x22\x23\x36\xa9\x93\xc6\x6d\x3b\xf2\xe8\xd9\xce\xdf\x1c\x4e\x05\x39\xd9\x7b\xf8\x6a\xd5\xd2\x00\x7d\xcf\xf2\x7a\xd4\xd7\xc0\x12\x19\xe0\x72\xd3\x79\x09\x01\x27\x1c\xe7\x81\x47\x7f\x42\xb7\xbe\x47\xd8\x87\x1c\xc3\x0d\x88\x23\x62\x08\xdc\xc5\xaf\x2d\x8c\x2e\xf7\xd9\x1c\x28\x22\xa9\xb0\x69\x3c\x7f\xe5\x28\xdd\xc1\x3d\x3c\x4e\xa2\xaf\x54\x4b\xd9\x18\xde\x14\x6e\xb0\xd7\x9b\x6a\xae\xfd\x86\xca\xa2\x5d\x3d\xc4\xb5\xee\x94\x0e\x6e\x76\xf1\xd0\xf0\xff\x45\xcd\x5a\x55\xc3\x55\x25\x0b\xc0\x75\x25\x71\x03\xbc\xb7\x07\x3f\xf0\xaf\x34\xaa\x8a\xc6\x2b\x8b\x15\x46\xfa\x7c\x10\x3b\xa1\xc4\xa4\x8d\x0f\x0c\xde\xb2\x89\x33\xfe\xf9\xa7\xaf\x5f\x02\xd6\x7b\x8f\xd9\x6f\xe1\x93\x92\x96\x0b\x09\x9a\x55\x3b\x28\x0d\x1d\xdc\x55\x69\x74\x24\x39\x98\x53\xd7\x24\xcc\x5c\x30\xa8\x07\x28\x4f\x6a\xe0\x3f\xb6\x93\x8e\x31\x10\x05\xdf\x24\x7f\xe1\xa2\x0f\x57\x1b\x85\x2e\xcd\x5b\x61\xfa\x84\xdd\x1e\x68\x26\x5e\x38\x1b\xd6\xc3\x57\x3f\x07\x18\x30\xc0\xc8\x2d\x66\x1e\x21\x85\x15\xbc\xbf\x82\x9e\x1f\xb7\xd0\x29\xb9\x33\x69\xea\xe8\x93\xfe\x42\x68\xc5\x00\x6a\xb2\x85\xd4\x4c\x5d\x07\xc6\xfc\x7a\xd0\x60\x0e\xca\xc5\x55\x90\x3f\x71\xd1\x4f\x1a\x2a\xf9\xc1\xda\xf1\x06\xf8\x0e\xb4\xf3\x7a\xa4\xf7\x4d\x7e\x90\xfc\xbf\x65\x9d\xe2\x2d\x6f\xa7\x32\xc0\x8a\x3a\xb0\x2a\xb6\x5a\xee\x90\x3f\xf3\xd4\x71\xfe\xff\x54\x7b\x2d\x96\x55\xb8\x72\xc6\xea\x95\x1f\x70\x5e\x57\xef\x3d\x45\xfe\x62\xf1\x5e\x98\x3a\x7f\x5d\xa9\xb7\xf1\x5a\x2d\x5f\xa5\x02\x0b\xb5\x3d\xb9\x2e\x17\xf9\x72\xa9\x26\x6f\xeb\xb0\x4c\x63\xf3\x67\x94\xc8\x62\x5e\x70\x06\x89\xb9\xba\xb6\x62\xaf\xe6\x2e\xc2\xf7\x1c\x7a\xb8\xd5\xc7\x9c\x40\x66\xe7\xb2\x25\x4d\xe9\x8b\x25\xf0\x47\x80\xb3\xbd\xa7\x99\xf0\x2f\x9e\x4a\x3b\xf3\x2f\xd3\x66\xea\x09\xcd\x6c\xbe\x95\xfa\x9c\xd4\xe4\xb2\x69\xb6\x4b\x9c\xeb\x16\xf2\xb7\x25\xf3\xa7\xac\x4e\x3f\x7c\x90\x2d\xb7\xbe\x89\xf8\x1d\x97\xee\x52\xb6\x74\xe9\xc4\x48\xa4\xe7\xae\x37\x2f\x57\x45\xe1\x5e\x73\x59\x05\x4d\xfd\x39\xa5\x69\xff\x16\x13\x2d\xba\xc2\x1b\x17\x0a\xa6\x59\x5c\x68\xb6\x5c\x2b\x26\xa8\xe9\x16\xb6\x19\xcc\xf7\xdf\x00\x00\x00\xff\xff\xba\xbc\xb3\x6c\x29\x27\x00\x00")

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

	info := bindataFileInfo{name: "plugins/codeamp/graphql/schema.graphql", size: 10025, mode: os.FileMode(420), modTime: time.Unix(1549324702, 0)}
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

	info := bindataFileInfo{name: "plugins/codeamp/graphql/static/index.html", size: 2810, mode: os.FileMode(420), modTime: time.Unix(1533663930, 0)}
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

