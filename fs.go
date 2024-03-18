package coalescefs

import (
	"errors"
	"io/fs"
)

type internalFS struct {
	filesystems []fs.FS
}

// FS returns a fs.FS that is the combined coalesced union of all the file-systems in 'filesystems'.
func FS(filesystems ...fs.FS) fs.FS {
	return internalFS{
		filesystems:filesystems,
	}
}

func (receiver internalFS) Open(name string) (file fs.File, err error) {
	for _, filesystem := range receiver.filesystems {
		file, err = filesystem.Open(name)
		if errors.Is(err, fs.ErrNotExist) {
			continue
		}
		return file, err
	}

	return nil, fs.ErrNotExist
}
