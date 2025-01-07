// some shitty comment
package files

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
)

func OpenFile(ctx context.Context) (*os.File, error) {
	go func() {
		<-ctx.Done()
		log.Println("program is being cancelled during open file")
		os.Exit(-1)
	}()

	path, pathErr := filePath()

	if pathErr != nil {
		return nil, pathErr
	}

	return os.OpenFile(path, os.O_RDONLY, 0666)
}

func filePath() (string, error) {
	var path string

	if _, err := fmt.Fscan(os.Stdin, &path); err != nil {
		return "", err
	}

	if path == "" {
		return path, errors.New("did not provide a path")
	}

	wd, wdErr := os.Getwd()

	if wdErr != nil {
		return wd, wdErr
	}
	return wd + path, nil
}
