package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/control"
	"github.com/jstraney/nother/pkg/game"
)

func main() {
	// Get the asset directory (in this example, we'll use an empty fs for now)
	exeDir, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable directory: %v", err)
	}
	assetDir := filepath.Dir(exeDir)

	// Create game configuration
	config := game.Config{
		WindowWidth:  800,
		WindowHeight: 600,
		WindowTitle:  "Pong",
		TargetFPS:    60,
		AssetFS:      os.DirFS(assetDir),
	}

	// Create the game
	g, err := game.New(config)
	if err != nil {
		log.Fatalf("Failed to create game: %v", err)
	}

	// Initialize the game
	if err := g.Init(); err != nil {
		log.Fatalf("Failed to initialize game: %v", err)
	}
	defer g.Shutdown()

	// Create the input dispatcher
	dispatcher := control.New()
	dispatcher.WatchMouseEvents()

	// Create scenes with callbacks for switching
	var gameScene *GameScene
	menuScene := NewMainMenuScene(dispatcher, func() {
		// Transition to game scene
		if err := g.SceneManager().ChangeScene("game"); err != nil {
			fmt.Printf("Failed to change to game scene: %v\n", err)
		}
	})

	gameScene = NewGameScene(dispatcher, func() {
		// Transition back to menu
		if err := g.SceneManager().ChangeScene("main-menu"); err != nil {
			fmt.Printf("Failed to change to menu scene: %v\n", err)
		}
	})

	// Add scenes to the manager
	if err := g.SceneManager().AddScene("main-menu", menuScene); err != nil {
		log.Fatalf("Failed to add main menu scene: %v", err)
	}
	if err := g.SceneManager().AddScene("game", gameScene); err != nil {
		log.Fatalf("Failed to add game scene: %v", err)
	}

	// Start with the main menu
	if err := g.SceneManager().ChangeScene("main-menu"); err != nil {
		log.Fatalf("Failed to change to main menu: %v", err)
	}

	// Main game loop
	for !rl.WindowShouldClose() {
		// Poll input events
		dispatcher.Poll()

		// Update and render
		g.SceneManager().Update(1.0 / 60.0)

		// Draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// Render current scene
		if currentScene := g.SceneManager().GetScene(g.SceneManager().Current()); currentScene != nil {
			if renderable, ok := currentScene.(interface{ Render() }); ok {
				renderable.Render()
			}
		}

		rl.EndDrawing()

		// Sleep to maintain frame rate
		time.Sleep(time.Duration(1000/60) * time.Millisecond)
	}
}
