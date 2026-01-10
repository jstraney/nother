package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/control"
	"github.com/jstraney/nother/pkg/gui"
	"github.com/jstraney/nother/pkg/node2d"
	"github.com/jstraney/nother/pkg/scene"
)

// MainMenuScene is the main menu scene for Pong.
type MainMenuScene struct {
	*scene.BaseScene
	titleLabel  *gui.Label
	startButton *gui.Button
	onStartGame func()
	dispatcher  *control.Dispatcher
}

// NewMainMenuScene creates a new main menu scene.
func NewMainMenuScene(dispatcher *control.Dispatcher, onStartGame func()) *MainMenuScene {
	return &MainMenuScene{
		BaseScene:   scene.NewBaseScene("main-menu"),
		onStartGame: onStartGame,
		dispatcher:  dispatcher,
	}
}

// Load initializes the scene resources.
func (m *MainMenuScene) Load() error {
	root := m.Root().(*node2d.Node2D)

	// Create title label
	m.titleLabel = gui.New("title", "Title")
	m.titleLabel.SetText("PONG")
	m.titleLabel.SetFontSize(80)
	m.titleLabel.SetColor(rl.White)
	m.titleLabel.SetPosition(rl.NewVector2(400-150, 100))
	root.AddChild(m.titleLabel)

	// Create start button
	m.startButton = gui.NewButton("start-btn", "StartButton")
	m.startButton.SetText("Start Game")
	m.startButton.SetSize(200, 60)
	m.startButton.SetColors(rl.Blue, rl.White)
	m.startButton.SetHoverColor(rl.DarkBlue)
	m.startButton.SetFontSize(24)
	m.startButton.SetPosition(rl.NewVector2(400-100, 300))
	m.startButton.BindInputDispatcher(m.dispatcher)
	m.startButton.SetOnClick(m.onStartGame)
	root.AddChild(m.startButton)

	return nil
}

// Render renders the scene.
func (m *MainMenuScene) Render() {
	m.titleLabel.Render()
	m.startButton.Render()
}
