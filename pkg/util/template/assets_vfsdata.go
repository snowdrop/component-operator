// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package template

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 10, 11, 9, 22, 1, 958566131, time.UTC),
		},
		"/innerloop": &vfsgen۰DirInfo{
			name:    "innerloop",
			modTime: time.Date(2018, 10, 11, 9, 9, 14, 352956734, time.UTC),
		},
		"/innerloop/deploymentconfig": &vfsgen۰CompressedFileInfo{
			name:             "deploymentconfig",
			modTime:          time.Date(2018, 10, 11, 9, 9, 14, 351417908, time.UTC),
			uncompressedSize: 2181,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\xcf\x6f\xeb\x36\x0c\xbe\xe7\xaf\x20\x7a\xb7\xfd\x5e\x81\x77\xd1\x3b\x65\x49\x37\xb4\x40\x3b\xa3\xe9\x72\x2d\x18\x99\x71\xd4\xe9\x17\x24\xd9\x43\x60\xf4\x7f\x1f\x14\xa7\x8e\x9c\x38\x2d\x36\x60\xf3\xc9\x22\xfd\x91\x1f\xc9\xcf\x14\x5a\xb1\x26\xe7\x85\xd1\x0c\xd0\x5a\x9f\x1b\x4b\xda\xef\xc4\x36\xe4\xc2\x14\xed\xf7\xd9\x9f\x42\x57\x0c\x96\x64\xa5\xd9\x2b\xd2\x61\x61\xf4\x56\xd4\x33\x45\x01\x2b\x0c\xc8\x66\x00\x12\x37\x24\x7d\x7c\x83\x18\x83\x41\xd7\xe5\x4f\xa8\xe8\xfd\x7d\x06\xa0\x51\x51\x6a\xf1\x96\x78\xfc\xd4\x91\x95\x82\xa3\x67\xf0\x7d\x06\xe0\x49\x12\x0f\xc6\x5d\x09\x02\x50\x0d\x04\xf8\x81\xc0\xd8\xef\x83\xc3\x40\xf5\xbe\x87\x87\xbd\x25\x06\xcf\x46\x4a\xa1\xeb\x19\x40\x20\x65\x25\x06\xea\xbd\x29\xf3\xf8\xa4\xec\xaf\x24\xff\x9a\x40\x7c\xce\x2b\x8d\xb6\x8f\x6a\xe3\xc3\x8d\x0e\x28\x34\xb9\x21\x59\x06\xe8\xea\x24\x75\x06\x19\x4f\x0e\x45\x8b\xae\x90\x62\x53\xf8\xc6\x92\x6b\x85\x37\xae\x2a\x62\xf6\xc4\x90\xc7\xf3\x80\xe1\x46\x29\xd4\x15\xfb\x2a\xc8\x46\xe8\xf4\x3c\x7c\x4e\xba\x4d\xb1\x7d\x45\x0f\xf3\xf5\xfc\x75\x5e\x96\xaf\xcb\xfb\xe7\xc1\x09\xd0\xa2\x6c\x88\x41\x71\x6a\x8c\x9f\x86\x2e\xef\x7e\xf9\xe3\xb7\x4b\xe0\x4d\x70\x0d\xdd\x7c\x02\x79\x2d\x7f\x7f\x7e\x99\xc0\xfd\xf8\xf6\xed\xc7\x15\x5c\x64\xf9\x30\x9f\x60\x89\xd6\xe6\x6f\xe8\x06\x47\xd7\x81\x43\x5d\x13\xe4\x2b\x4b\x3c\xbf\xd3\xad\x87\x64\xd8\xd9\xe4\x30\x47\x11\xbb\x2e\x5f\xc7\xb7\xc4\xd7\x75\xa4\xab\xe4\x2c\x14\xd6\x91\x70\x45\x6d\xe6\x6f\x05\x8b\x1a\xf4\xe1\x66\xec\x2f\x1b\x29\x4b\x23\x05\xdf\x33\x98\xcb\xbf\x70\x7f\xea\xe2\x35\x0e\xd6\xb8\x30\x52\xcd\xa0\xac\xd2\xb8\x70\x40\x1c\xaa\x8a\xa7\x11\x75\xeb\x4c\x30\xdc\x48\x06\x2f\x8b\x72\xb0\xb7\x46\x36\x8a\x1e\x4d\xa3\xc7\x51\x55\xb4\x94\x18\x76\x6c\x52\x44\x49\xdc\x9e\xa8\xdf\xa1\xa3\x2a\x8b\xff\xd6\x95\x28\x41\xd9\x02\x5d\x10\x5b\xe4\x89\x58\x3e\xf0\xea\x36\xc5\x0a\x2d\xe2\xae\xb9\xf8\x63\x26\x25\xba\x78\x5c\xae\x2e\x47\xe4\x1a\x9d\xbd\x61\x8b\xac\x68\xbc\x2b\xa4\xe1\x28\x0b\x7f\x2b\x0a\xd7\xe8\x9f\xdc\x28\x2b\x24\x4d\xfa\xd1\x7b\x52\x1b\x49\x3f\x37\x8d\x90\x15\x4b\x15\x5e\x1c\x4c\x68\xed\xc5\x90\x3f\x9a\xbe\x3a\x75\xa8\x1f\xdb\xbf\x1c\x3b\x37\x76\x9f\x4d\xb5\xdb\x91\x37\x8d\xe3\xe4\x19\x74\xa7\xe1\x06\x72\x4a\x68\x0c\xc2\xe8\x47\xf2\x3e\x66\xe8\x7b\x5e\x51\x5b\x24\xce\x4c\x9a\xfa\x33\xd0\x91\xd2\xaf\x42\xd2\xff\x20\x90\x3e\x72\x3a\x5d\x65\xc3\x7e\x29\xdc\xa8\xb6\x6b\xe8\x6c\x52\x38\x00\x36\x5e\x68\x3e\x90\x0e\xeb\x43\xfc\x85\x44\xa1\x58\x42\x87\x47\xc3\xd3\x19\x36\x38\x51\xd7\x47\xa9\x65\xc7\x3b\xe4\x3e\x8e\x6a\xb1\x8b\x9b\x62\x36\x8c\xae\x3f\x97\xe8\x50\x0d\xcc\xb1\x09\x46\x61\x10\x9c\x41\x5c\x6b\xe7\x1b\x3f\xe6\x4a\xaa\xbc\x32\xdb\xad\x33\x09\xcd\xfe\xe2\x3d\x30\x58\x05\x47\xa8\x5e\xb0\xfe\x42\x23\x47\xa9\xfd\xf7\x05\x9c\xef\xa4\x7f\xcc\x7c\xbc\x12\xff\x0e\x00\x00\xff\xff\x10\xdf\x2e\xb7\x85\x08\x00\x00"),
		},
		"/innerloop/imagestream": &vfsgen۰CompressedFileInfo{
			name:             "imagestream",
			modTime:          time.Date(2018, 10, 8, 10, 12, 58, 129088114, time.UTC),
			uncompressedSize: 620,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x41\xeb\xdb\x30\x0c\xc5\xef\xf9\x14\xe2\xcf\xae\x4d\xf8\xef\xe8\xb2\x43\xd9\x2e\x83\x31\xc6\x0a\xbb\xab\x8e\xd2\x79\xb5\x2d\x61\x39\x85\x62\xf2\xdd\x47\x1d\x37\xed\x60\x27\xdb\x3f\xe9\x49\xcf\xaf\x94\x0f\x11\x03\x1d\x44\xc0\x7c\x82\xfe\x3b\x06\x82\x65\x41\x71\xbf\x28\xa9\xe3\x68\xe0\xfa\xde\x5d\x5c\x1c\x0d\x7c\x73\x9a\x3b\x97\x29\xa8\xe9\x4a\x81\x84\xf1\x4c\xd0\x1f\x85\x6c\xff\x35\xe0\x99\x14\x96\xa5\xdb\xc1\xab\xd6\xdd\x79\xcf\x42\x51\x7f\xbb\x29\xf7\x8e\x87\xeb\x7b\x07\xb0\x0e\xac\xaa\x63\x4e\x84\xa1\x03\x08\x94\x71\xc4\x8c\xa6\x03\x00\xf0\x78\x22\xaf\xeb\x1d\x00\x45\x0c\x3c\xad\x2e\x4b\xe5\xf7\xd7\x1d\x57\xd7\x95\xa9\x90\x6d\x7a\xe6\xcb\x2c\x3f\xd8\x3b\x7b\x7b\x4c\xf1\x6c\xd1\x1b\x98\xd0\x2b\x55\x54\x8a\x9b\xa0\xff\xc2\xf6\x42\xa9\x9a\x69\x83\x33\x9e\xdb\xea\x5d\xeb\x39\xc4\xc8\x19\xb3\xe3\xf8\x39\x8c\xda\xda\x00\x70\xc3\x9b\x55\x00\x1b\x46\x35\xf0\x96\xe6\xb8\xfb\x83\x57\x34\xc3\xac\x69\xa8\xbb\x07\xfd\xe8\x86\x34\xc7\xbd\xe5\x20\xce\xd3\x7f\xeb\xa8\x4a\xe1\xe4\x69\x7f\x9a\x9d\x1f\xcd\x30\x92\x78\xbe\x05\x8a\x59\x87\x8a\x50\xe4\xad\x2d\x2b\x85\xe2\xb8\xb9\x99\x12\x87\xa7\x8d\x35\xe3\x97\xcf\x6d\x95\x2d\xb7\x9f\x24\xbc\xa9\x5d\x10\x4e\xb9\x25\x06\xe5\x81\xd7\x66\x8f\x99\x34\x37\x94\x68\xa2\x44\xd1\xd2\xbf\xf1\x02\xe4\x9b\x90\x81\x23\xcf\xc9\x3e\x02\xae\xfe\xda\xf1\x37\x00\x00\xff\xff\xec\x98\x21\x83\x6c\x02\x00\x00"),
		},
		"/innerloop/pvc": &vfsgen۰CompressedFileInfo{
			name:             "pvc",
			modTime:          time.Date(2018, 10, 4, 9, 16, 23, 850123135, time.UTC),
			uncompressedSize: 250,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\xbd\x8a\xc3\x30\x10\x84\x7b\x3d\xc5\xbe\xc0\x19\xae\x55\xeb\xfa\x8e\x03\x83\xfb\x3d\x79\x08\x22\xfa\x8b\x76\x1d\x08\x46\xef\x1e\xe4\x98\x10\x70\xb9\x33\xdf\xec\x0c\x17\x3f\xa3\x8a\xcf\xc9\xd2\xfd\xdb\x5c\x7d\x5a\x2c\xfd\x75\x45\x14\x49\xe7\x1c\xd6\x88\x31\xb0\x8f\x26\x42\x79\x61\x65\x6b\x88\x12\x47\x58\xda\xb6\x61\x2a\x70\xc3\xa4\xb9\xf2\x05\xc3\x2f\x47\xb4\x66\x88\x02\xff\x23\x48\x07\x89\xb8\x94\x9d\x7c\x9b\x1f\xe9\x43\x93\x02\xd7\x61\x76\x0e\x22\x3f\x79\xc1\x9e\xfd\x3a\x15\x74\x6b\xff\x51\x21\x79\xad\x0e\x47\x47\xc5\x6d\x85\xe8\x71\x11\xc9\x8b\x3f\x2f\x1c\xb9\xb0\xf3\xfa\x68\xcd\x3c\x03\x00\x00\xff\xff\x76\xbd\xed\x76\xfa\x00\x00\x00"),
		},
		"/innerloop/route": &vfsgen۰CompressedFileInfo{
			name:             "route",
			modTime:          time.Date(2018, 8, 14, 17, 14, 16, 0, time.UTC),
			uncompressedSize: 172,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcd\x31\x0e\xc2\x30\x0c\x85\xe1\x3d\xa7\xf0\x09\x82\x58\x73\x08\x06\x90\xd8\x4d\xfb\x10\x16\x4d\x6c\x25\xa6\x4b\xd5\xbb\xa3\x28\x13\xa2\xf3\xff\x3e\x3d\x36\xb9\xa3\x36\xd1\x92\xa8\xea\xc7\x11\xd5\x50\xda\x4b\x9e\x1e\x45\x4f\xeb\x39\xbc\xa5\xcc\x89\xae\xbd\x85\x0c\xe7\x99\x9d\x53\x20\x2a\x9c\x91\x68\xdb\xe2\x85\x33\xf6\x3d\x10\x2d\xfc\xc0\xd2\x7a\x23\x62\xb3\xdf\xf8\x0f\x9a\x61\xea\x63\xd7\x41\xc6\xd1\x0d\x75\x95\x09\x47\xe2\x1b\x00\x00\xff\xff\xdd\xc7\x7e\xf9\xac\x00\x00\x00"),
		},
		"/innerloop/service": &vfsgen۰CompressedFileInfo{
			name:             "service",
			modTime:          time.Date(2018, 10, 4, 7, 33, 27, 665611911, time.UTC),
			uncompressedSize: 257,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\x31\x8a\xc3\x30\x10\x45\x7b\x9d\x62\x2e\xb0\x86\x6d\xd5\x6e\xbf\x18\xbc\x6c\x3f\x91\x7f\x8c\x88\xa4\x19\xa4\xc1\x10\x8c\xef\x1e\x2c\x52\x24\x24\xe9\x86\x79\xef\xc1\x67\x8d\xff\xa8\x2d\x4a\xf1\xb4\x7e\xbb\x4b\x2c\xb3\xa7\x09\x75\x8d\x01\x2e\xc3\x78\x66\x63\xef\x88\x0a\x67\x78\xda\xb6\xe1\x97\x33\xf6\xdd\x11\x25\x3e\x21\xb5\x83\x11\xb1\xea\x33\x7c\x0d\x9a\x22\x1c\xb2\x4a\xb5\x5e\x7d\xf5\xb3\x2b\x93\x22\x0c\xa3\x54\xbb\xb7\x5a\xc5\x24\x48\xf2\xf4\xf7\x33\xf6\x8f\x71\x5d\x60\xe3\xfb\xa0\x21\x21\x98\xd4\x8f\x5b\x66\x68\x92\x6b\x46\xb1\x20\xe5\x1c\x97\x07\x7e\x0b\x00\x00\xff\xff\x7c\x58\xa5\x67\x01\x01\x00\x00"),
		},
		"/servicecatalog": &vfsgen۰DirInfo{
			name:    "servicecatalog",
			modTime: time.Date(2018, 10, 11, 9, 25, 12, 784142577, time.UTC),
		},
		"/servicecatalog/servicebinding": &vfsgen۰CompressedFileInfo{
			name:             "servicebinding",
			modTime:          time.Date(2018, 10, 11, 9, 25, 12, 782687771, time.UTC),
			uncompressedSize: 273,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\x31\x6e\xc3\x30\x0c\x45\x77\x9d\x82\x43\x67\x15\xd9\x0a\x01\x1d\xda\x03\x64\xa8\x81\xee\x8c\xfc\x6b\x10\xb1\x69\xc1\x14\xb2\x08\xbc\x7b\x21\xb7\xa9\x9b\x51\xff\xe9\xf1\x93\xad\x3d\x29\x2f\x78\x2b\x85\xd2\x2b\xc5\x33\x2f\x20\xf7\xd0\x1a\x6d\xac\x13\x28\x0e\x05\x39\x0e\xd8\x6e\x92\x61\xe4\xce\x45\x3e\xb1\x99\xac\x9a\xc8\x7e\xe2\xcc\x95\xe7\x75\x8a\xd7\x17\x8b\xb2\x3e\xdf\x4e\x17\x54\x3e\x85\xab\xe8\x98\xe8\x57\x7d\x17\x1d\x45\xa7\xb0\xa0\xf2\xc8\x95\x53\x20\xea\xc5\x89\x5a\xdb\x5b\xdd\x03\xd1\xcc\x17\xcc\xd6\x19\x11\x97\xd2\xe1\x7d\xbd\x9d\x1f\xce\xbf\xd8\x0a\x72\x57\x44\xad\xb2\x66\x7c\xe0\x2b\x3d\xfc\x3d\xe6\x1b\xf2\x86\x7a\xbe\xe7\xc3\xdf\x73\x3f\x19\x3a\xba\x7f\x07\x00\x00\xff\xff\x85\x6f\x46\xd9\x11\x01\x00\x00"),
		},
		"/servicecatalog/serviceinstance": &vfsgen۰CompressedFileInfo{
			name:             "serviceinstance",
			modTime:          time.Date(2018, 10, 11, 8, 47, 22, 758236171, time.UTC),
			uncompressedSize: 333,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xbd\x6a\x04\x21\x14\x85\x7b\x9f\xe2\x16\xa9\x0d\xdb\x05\x21\x45\x08\x29\x92\x62\x09\x0c\xa4\x3f\xeb\x5c\x16\x59\xe7\x8e\x78\xcd\x12\x10\xdf\x3d\xe8\xee\xe4\x87\x2d\xfd\xce\x39\x7e\x62\xad\x77\x82\x85\x9f\x52\x22\xf7\x48\x76\x8f\x85\xa9\x35\x53\x2b\x65\xc8\x91\xc9\x4e\x89\xbd\x9d\x38\x9f\x83\x67\xa5\xd6\x90\xc2\x07\x67\x0d\xab\x38\xd2\x0b\xf6\x28\x88\xeb\xd1\x9e\x1e\xd4\x86\xf5\xfe\xbc\x3b\x70\xc1\xce\x9c\x82\xcc\x8e\xae\xd3\x57\xd1\x02\xf1\x6c\x16\x2e\x98\x51\xe0\x0c\x51\x37\x3b\xaa\x75\x68\x5b\x33\x44\x11\x07\x8e\xda\x33\x22\xa4\xd4\xc3\xed\x7d\x23\xff\xdd\xfc\xc1\x9a\xd8\xf7\x89\x8f\x9f\x5a\x38\x5f\x8d\xcf\x11\xaa\x2f\x5f\x85\xb3\x20\xee\x37\xd3\xa0\xe3\xaa\xff\xed\xf7\x08\xb9\x29\x77\x38\xba\x09\x19\x0b\x17\xce\x7a\xe1\x3f\xc7\xb7\x69\x95\xf1\x5d\x2c\x73\x6b\xdf\x01\x00\x00\xff\xff\xdd\xed\x80\x29\x4d\x01\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/innerloop"].(os.FileInfo),
		fs["/servicecatalog"].(os.FileInfo),
	}
	fs["/innerloop"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/innerloop/deploymentconfig"].(os.FileInfo),
		fs["/innerloop/imagestream"].(os.FileInfo),
		fs["/innerloop/pvc"].(os.FileInfo),
		fs["/innerloop/route"].(os.FileInfo),
		fs["/innerloop/service"].(os.FileInfo),
	}
	fs["/servicecatalog"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/servicecatalog/servicebinding"].(os.FileInfo),
		fs["/servicecatalog/serviceinstance"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr: gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}