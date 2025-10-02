package utils

import (
	"embed"
	"fmt"
	"io/fs"
)

func RemoveFsPrefix(embedded embed.FS, prefix string) fs.FS {
	output, err := fs.Sub(embedded, prefix)
	if err != nil {
		panic(fmt.Errorf("failed getting the sub tree for the site files: %w", err))
	}
	return output
}
