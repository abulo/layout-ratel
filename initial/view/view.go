package view

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

var _view_backstage_common_allow_html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\xbf\x6e\x13\x41\x10\xc6\xeb\xf8\x29\x36\x1b\x41\x95\xcd\xd9\x89\x82\x62\xfb\xce\x12\x24\x29\xa8\xa0\x00\x09\xca\xbd\xdd\xb1\x6f\xe4\xbd\xdd\x65\x77\xfc\x0f\xcb\x15\x0d\x08\x84\x80\x06\x21\x81\x04\x2d\x45\x0a\x2a\x24\x40\xbc\x4c\x62\x91\x8a\x57\x40\x67\xc7\xe4\x0f\x0d\xd5\xdd\xcd\xed\xf7\xfb\xbe\x19\xcd\x4e\xa7\x4c\x43\x17\x2d\x30\x9e\x4b\xd5\x8f\x24\x7b\x90\x28\x57\x96\xce\x26\xd2\x18\x37\xda\x2a\xa8\x34\x9c\xcd\x66\xb5\x74\xfd\xe0\xce\xfe\xbd\x87\x77\x0f\x59\x55\xea\xd4\xd2\xe5\xa3\x96\x16\x20\x75\xa7\xb6\x96\x96\x40\x92\xa9\x42\x86\x08\x94\xf1\x01\x75\xc5\x1e\xaf\xea\x84\x64\xa0\x33\xff\xf2\x69\xfe\xfe\xd9\xaf\xa3\x9f\xa7\x6f\x8f\xe6\x1f\x9e\x9c\xbe\x7b\x95\x26\xcb\x3f\x2b\xa9\x95\x25\x64\x3c\x80\xd5\x10\x20\x70\xa6\x9c\x25\xb0\x94\xf1\x11\xe4\x7d\x24\xfe\xf7\x60\x41\xe4\x05\x3c\x1a\xe0\x30\xe3\x0f\xc4\xfd\x9b\x62\xdf\x95\x5e\x12\xe6\x06\x2e\xa8\x6e\x1f\x66\xa0\x7b\xb0\xa9\x8a\xe0\x4a\xc8\x1a\xfc\x8a\xd1\x10\x61\xe4\x5d\xa0\x8b\x46\xa8\xa9\xc8\x34\x0c\x51\x81\x58\x7c\x6c\x32\xb4\x48\x28\x8d\x88\x4a\x1a\xc8\x1a\x9b\xac\x94\x63\x2c\x07\xe5\xaa\x70\x15\x2b\xbd\x37\x20\x4a\x97\xa3\x01\x31\x82\x5c\x48\xef\x45\x24\x49\x83\x28\x72\x19\x44\xa4\xc9\xa5\x98\xb9\x91\xaa\xff\x5f\x10\x25\xbd\xbc\xdc\xe2\x04\xe2\x55\x65\xd7\x85\x52\x92\xd0\x40\xa0\x08\x9d\xbd\x70\x9a\xc0\x80\x2f\x9c\x85\xcc\xba\x85\xcc\xa0\xed\xb3\x00\x26\xe3\x8b\x50\xb1\x00\x20\xce\x8a\x00\xdd\x8c\x27\x01\xa2\x1b\x04\x05\x49\x95\x1d\x55\xe2\xcd\xa0\x87\x36\x31\x72\x32\xc0\x44\xc5\xb8\x7c\xdb\x52\xb1\x8a\x90\x26\xcb\x25\xa8\xa5\xb9\xd3\x13\xa6\x8c\x8c\x31\xe3\xaa\x40\xa3\x03\xd8\x5b\x4e\x4f\x16\x86\x1a\x87\x6c\x61\x55\x85\x19\x93\x90\x06\x7b\xb6\xc5\x14\x58\x82\xd0\x66\x5e\x6a\x8d\xb6\xd7\x6a\x34\xae\xb1\x7a\x9b\x77\x58\x8a\x2b\xd4\xc2\x4c\xa0\xaa\x1a\x3a\x23\x18\xb4\x20\x0a\xc0\x5e\x41\xad\xed\xdd\x00\x65\x9b\x75\x9d\x25\x11\xf1\x31\xb4\x1a\x7b\x8b\x82\x72\xc6\x85\x16\xdb\xd8\x69\xee\x1c\xec\x56\xc8\xeb\x1b\x63\xb8\xd1\x54\xed\x34\xc1\x4e\x6d\x6d\x2d\xf5\x2b\xda\xb9\x94\x6d\xd7\xfd\xf8\x8c\x35\x5a\xe2\xd9\x4e\xbd\x7e\x0e\x6b\x36\x9b\x6d\xde\x99\x3f\x7d\x7d\xf2\xfc\xcd\xf1\xd7\xcf\x27\x2f\xbf\xfd\xfe\xfe\xe2\xf8\xc7\xc7\x7f\x37\x7c\x3d\x4d\x7c\xd5\x77\xa2\x71\x58\x0d\xa9\x1a\x4e\x35\xa4\x64\x79\x75\xa6\x53\x06\x56\xb3\xd9\xec\x4f\x00\x00\x00\xff\xff\xa3\x7f\x0a\x85\x86\x03\x00\x00")

func view_backstage_common_allow_html() ([]byte, error) {
	return bindata_read(
		_view_backstage_common_allow_html,
		"view/backstage/common/allow.html",
	)
}

var _view_backstage_common_redirect_html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\x41\x4e\xc4\x30\x0c\x45\xf7\x39\x85\x95\x15\xb3\xa0\xdd\x22\xcd\x94\xab\x20\x93\x3a\x6d\x20\x4d\x2a\xc7\x65\x40\x96\xef\x8e\x34\x03\x88\x22\x31\xdb\xff\xfd\xbe\xfc\x54\x61\xa4\x98\x0a\x81\x7f\xc6\xf0\xda\x04\x27\xea\x43\x5d\x96\x5a\x7a\xa6\x31\x31\x05\xe9\x66\x59\xb2\x37\x73\xa7\x16\x38\xad\x02\xf2\xb1\xd2\xe0\x85\xde\xa5\x7f\xc1\x37\xbc\xa6\x1e\x32\x96\x69\xc3\x89\x06\xff\x3b\x0d\x33\x72\x23\x19\xfc\x26\xf1\xfe\xc1\x3f\x3a\x55\x48\x11\xba\x9f\xf5\x52\x25\x05\x02\x33\x87\x99\x58\xee\x54\xff\x76\x66\x87\xa3\x53\xa5\x32\x9a\xb9\x14\x77\x17\x2b\x32\x15\x79\xda\x38\x9b\x1d\xd4\x01\x00\x9c\x53\x19\xeb\xf9\xab\xe9\x72\x0d\x28\xa9\x96\x6e\x66\x8a\x30\xc0\x7f\xf0\xd1\x19\xe5\x46\xbb\x89\x5b\xec\x37\xe4\x4e\xfd\x55\xf5\x62\x76\xf9\xf1\x33\x00\x00\xff\xff\xfe\x58\x8b\x5e\x55\x01\x00\x00")

func view_backstage_common_redirect_html() ([]byte, error) {
	return bindata_read(
		_view_backstage_common_redirect_html,
		"view/backstage/common/redirect.html",
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
	"view/backstage/common/allow.html": view_backstage_common_allow_html,
	"view/backstage/common/redirect.html": view_backstage_common_redirect_html,
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
	"view": &_bintree_t{nil, map[string]*_bintree_t{
		"backstage": &_bintree_t{nil, map[string]*_bintree_t{
			"common": &_bintree_t{nil, map[string]*_bintree_t{
				"allow.html": &_bintree_t{view_backstage_common_allow_html, map[string]*_bintree_t{
				}},
				"redirect.html": &_bintree_t{view_backstage_common_redirect_html, map[string]*_bintree_t{
				}},
			}},
		}},
	}},
}}