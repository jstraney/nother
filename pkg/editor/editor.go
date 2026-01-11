package editor

import (
	"context"
	"fmt"

	"github.com/jstraney/nother/pkg/registry"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Editor struct
type Editor struct {
	ctx      context.Context
	Registry *registry.Store
}

// NewApp creates a new App application struct
func New() *Editor {
	return &Editor{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *Editor) Start(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *Editor) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *Editor) Weezafoo(name string) string {
	return fmt.Sprintf("Hello %s, It's Weezafoo!", name)
}

// OpenProjectDialog opens a folder/directory picker dialog
func (a *Editor) OpenProjectDialog() (string, error) {
	path, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Open Nother Project Directory",
	})
	if err != nil {
		return "", err
	}
	return path, nil
}
