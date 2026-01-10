package game

import (
	"fmt"
	"io/fs"
	"sync"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/asset"
	"github.com/jstraney/nother/pkg/scene"
)

// Config holds global game configuration.
type Config struct {
	// Window configuration
	WindowWidth  int32
	WindowHeight int32
	WindowTitle  string

	// Performance configuration
	TargetFPS int32

	// Asset filesystem
	AssetFS fs.FS
}

// Game is the main game engine wrapper.
// It orchestrates the AssetManager, SceneManager, and main game loop.
type Game struct {
	config        Config
	assetManager  *asset.AssetManager
	sceneManager  *scene.SceneManager
	running       bool
	lastFrameTime time.Time
	mu            sync.RWMutex
}

// New creates a new Game instance with the given configuration.
func New(config Config) (*Game, error) {
	// Validate configuration
	if config.WindowWidth <= 0 {
		return nil, fmt.Errorf("window width must be positive")
	}
	if config.WindowHeight <= 0 {
		return nil, fmt.Errorf("window height must be positive")
	}
	if config.TargetFPS <= 0 {
		return nil, fmt.Errorf("target FPS must be positive")
	}
	if config.WindowTitle == "" {
		return nil, fmt.Errorf("window title cannot be empty")
	}
	if config.AssetFS == nil {
		return nil, fmt.Errorf("asset filesystem cannot be nil")
	}

	return &Game{
		config:       config,
		assetManager: asset.New(config.AssetFS),
		sceneManager: scene.New(),
		running:      false,
		lastFrameTime: time.Now(),
	}, nil
}

// Init initializes the raylib window and prepares the game for running.
// This must be called before Run().
func (g *Game) Init() error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.running {
		return fmt.Errorf("game is already running")
	}

	// Initialize raylib window
	rl.InitWindow(g.config.WindowWidth, g.config.WindowHeight, g.config.WindowTitle)

	// Set target FPS
	rl.SetTargetFPS(g.config.TargetFPS)

	g.lastFrameTime = time.Now()
	return nil
}

// Run starts the main game event loop.
// This blocks until the window is closed or Stop() is called.
func (g *Game) Run() error {
	g.mu.Lock()
	if g.running {
		g.mu.Unlock()
		return fmt.Errorf("game is already running")
	}
	g.running = true
	g.mu.Unlock()

	defer func() {
		g.mu.Lock()
		g.running = false
		g.mu.Unlock()
	}()

	// Main game loop
	for !rl.WindowShouldClose() {
		// Calculate delta time
		now := time.Now()
		deltaTime := now.Sub(g.lastFrameTime).Seconds()
		g.lastFrameTime = now

		// Update current scene
		g.sceneManager.Update(deltaTime)

		// Begin frame
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Render current scene (if it has a Render method)
		currentSceneName := g.sceneManager.Current()
		if currentScene := g.sceneManager.GetScene(currentSceneName); currentScene != nil {
			if renderable, ok := currentScene.(interface{ Render() }); ok {
				renderable.Render()
			}
		}

		// End frame
		rl.EndDrawing()
	}

	return nil
}

// Stop stops the main game loop.
func (g *Game) Stop() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.running = false
}

// IsRunning returns whether the game is currently running.
func (g *Game) IsRunning() bool {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.running
}

// Shutdown shuts down the game, releasing all resources.
// This must be called before the program exits.
func (g *Game) Shutdown() error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.running {
		return fmt.Errorf("cannot shutdown while game is running")
	}

	// Unload all scenes
	if err := g.sceneManager.Shutdown(); err != nil {
		return fmt.Errorf("failed to shutdown scene manager: %w", err)
	}

	// Unload all assets
	if err := g.assetManager.UnloadAll(); err != nil {
		return fmt.Errorf("failed to unload assets: %w", err)
	}

	// Close raylib window
	rl.CloseWindow()

	return nil
}

// AssetManager returns the asset manager for loading assets.
func (g *Game) AssetManager() *asset.AssetManager {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.assetManager
}

// SceneManager returns the scene manager for managing scenes.
func (g *Game) SceneManager() *scene.SceneManager {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.sceneManager
}

// Config returns the game configuration.
func (g *Game) Config() Config {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.config
}

// SetTargetFPS updates the target FPS.
func (g *Game) SetTargetFPS(fps int32) error {
	if fps <= 0 {
		return fmt.Errorf("target FPS must be positive")
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	g.config.TargetFPS = fps
	if g.running {
		rl.SetTargetFPS(fps)
	}
	return nil
}

// DeltaTime returns the time elapsed since the last frame.
func (g *Game) DeltaTime() float64 {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return time.Since(g.lastFrameTime).Seconds()
}

// WindowWidth returns the window width in pixels.
func (g *Game) WindowWidth() int32 {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.config.WindowWidth
}

// WindowHeight returns the window height in pixels.
func (g *Game) WindowHeight() int32 {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.config.WindowHeight
}
