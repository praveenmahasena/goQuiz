package internal

import (
	"context"

	"github.com/praveenmahasena/goquiz/internal/files"
	"github.com/praveenmahasena/goquiz/internal/game"
)

func Run(ctx context.Context) error {
	file, fileErr := files.OpenFile(ctx)

	if fileErr != nil {
		return fileErr
	}

	return game.Play(ctx,file)
}
