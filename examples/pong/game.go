package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jstraney/nother/pkg/area2d"
	"github.com/jstraney/nother/pkg/control"
	"github.com/jstraney/nother/pkg/gui"
	"github.com/jstraney/nother/pkg/node2d"
	"github.com/jstraney/nother/pkg/scene"
)

// GameScene is the active gameplay scene for Pong.
type GameScene struct {
	*scene.BaseScene
	ballCircle  *area2d.Circle
	paddle1Rect *Paddle
	paddle2Rect *Paddle
	scoreLabel  *gui.Label
	ballVelX    float32
	ballVelY    float32
	ballSpeed   float32
	ballRadius  float32
	score1      int
	score2      int
	dispatcher  *control.Dispatcher
	onGameOver  func()
}

// Paddle represents a player's paddle.
type Paddle struct {
	*node2d.Node2D
	width  float32
	height float32
	speed  float32
}

// NewPaddle creates a new paddle.
func NewPaddle(id string, name string) *Paddle {
	return &Paddle{
		Node2D: node2d.New(id, name),
		width:  15,
		height: 100,
		speed:  400,
	}
}

// Update updates the paddle position.
func (p *Paddle) Update(deltaTime float64) {
	// Position update happens in the scene
	p.Node2D.Update(deltaTime)
}

// Render renders the paddle as a rectangle.
func (p *Paddle) Render() {
	pos := p.GlobalPosition()
	rl.DrawRectangle(int32(pos.X), int32(pos.Y), int32(p.width), int32(p.height), rl.White)
}

// NewGameScene creates a new game scene.
func NewGameScene(dispatcher *control.Dispatcher, onGameOver func()) *GameScene {
	return &GameScene{
		BaseScene:  scene.NewBaseScene("game"),
		ballSpeed:  300,
		score1:     0,
		score2:     0,
		dispatcher: dispatcher,
		onGameOver: onGameOver,
	}
}

// Load initializes the scene resources.
func (g *GameScene) Load() error {
	root := g.Root()

	// Initialize ball
	g.ballRadius = 5
	g.ballCircle = area2d.NewCircle("ball", "Ball")
	g.ballCircle.SetRadius(g.ballRadius)
	g.ballCircle.SetPosition(rl.NewVector2(400, 300))
	g.ballCircle.SetVisible(true)
	g.ballCircle.SetFill(rl.White)
	g.ballVelX = g.ballSpeed
	g.ballVelY = g.ballSpeed * 0.5
	root.AddChild(g.ballCircle)

	// Initialize paddles
	g.paddle1Rect = NewPaddle("paddle1", "Paddle1")
	g.paddle1Rect.SetPosition(rl.NewVector2(30, 250))
	root.AddChild(g.paddle1Rect)

	g.paddle2Rect = NewPaddle("paddle2", "Paddle2")
	g.paddle2Rect.SetPosition(rl.NewVector2(755, 250))
	root.AddChild(g.paddle2Rect)

	// Score label
	g.scoreLabel = gui.New("score", "Score")
	g.scoreLabel.SetText("0 : 0")
	g.scoreLabel.SetFontSize(40)
	g.scoreLabel.SetColor(rl.White)
	g.scoreLabel.SetPosition(rl.NewVector2(350, 20))
	root.AddChild(g.scoreLabel)

	// Watch keys for paddle control
	g.dispatcher.WatchKeys(rl.KeyW, rl.KeyS, rl.KeyUp, rl.KeyDown)

	return nil
}

// Unload cleans up scene resources.
func (g *GameScene) Unload() error {
	return nil
}

// Update updates the scene.
func (g *GameScene) Update(deltaTime float64) {
	if g.Paused() {
		return
	}

	// Update ball position
	ballPos := g.ballCircle.Position()
	ballPos.X += g.ballVelX * float32(deltaTime)
	ballPos.Y += g.ballVelY * float32(deltaTime)

	// Ball bounces off top and bottom
	if ballPos.Y-g.ballRadius <= 0 || ballPos.Y+g.ballRadius >= 600 {
		g.ballVelY = -g.ballVelY
		if ballPos.Y-g.ballRadius <= 0 {
			ballPos.Y = g.ballRadius
		} else {
			ballPos.Y = 600 - g.ballRadius
		}
	}

	// Ball out of bounds (left and right)
	if ballPos.X < 0 {
		g.score2++
		g.resetBall()
		g.updateScore()
	} else if ballPos.X > 800 {
		g.score1++
		g.resetBall()
		g.updateScore()
	}

	g.ballCircle.SetPosition(ballPos)

	// Update paddle 1 (W/S keys)
	p1Pos := g.paddle1Rect.Position()
	if rl.IsKeyDown(rl.KeyW) && p1Pos.Y > 0 {
		p1Pos.Y -= g.paddle1Rect.speed * float32(deltaTime)
	}
	if rl.IsKeyDown(rl.KeyS) && p1Pos.Y < 500 {
		p1Pos.Y += g.paddle1Rect.speed * float32(deltaTime)
	}
	g.paddle1Rect.SetPosition(p1Pos)

	// Update paddle 2 (Up/Down arrow keys)
	p2Pos := g.paddle2Rect.Position()
	if rl.IsKeyDown(rl.KeyUp) && p2Pos.Y > 0 {
		p2Pos.Y -= g.paddle2Rect.speed * float32(deltaTime)
	}
	if rl.IsKeyDown(rl.KeyDown) && p2Pos.Y < 500 {
		p2Pos.Y += g.paddle2Rect.speed * float32(deltaTime)
	}
	g.paddle2Rect.SetPosition(p2Pos)

	// Paddle collision
	g.checkPaddleCollision()

	g.Root().Update(deltaTime)
}

// checkPaddleCollision checks if the ball collides with either paddle.
func (g *GameScene) checkPaddleCollision() {
	ballPos := g.ballCircle.Position()

	// Paddle 1 collision
	p1Pos := g.paddle1Rect.Position()
	if ballPos.X-g.ballRadius < p1Pos.X+g.paddle1Rect.width &&
		ballPos.Y > p1Pos.Y && ballPos.Y < p1Pos.Y+g.paddle1Rect.height {
		g.ballVelX = -g.ballVelX
		ballPos.X = p1Pos.X + g.paddle1Rect.width + g.ballRadius
		g.ballCircle.SetPosition(ballPos)
	}

	// Paddle 2 collision
	p2Pos := g.paddle2Rect.Position()
	if ballPos.X+g.ballRadius > p2Pos.X &&
		ballPos.Y > p2Pos.Y && ballPos.Y < p2Pos.Y+g.paddle2Rect.height {
		g.ballVelX = -g.ballVelX
		ballPos.X = p2Pos.X - g.ballRadius
		g.ballCircle.SetPosition(ballPos)
	}
}

// resetBall resets the ball to the center.
func (g *GameScene) resetBall() {
	g.ballCircle.SetPosition(rl.NewVector2(400, 300))
	g.ballVelX = g.ballSpeed
	g.ballVelY = g.ballSpeed * 0.5
}

// updateScore updates the score label.
func (g *GameScene) updateScore() {
	g.scoreLabel.SetText(formatScore(g.score1, g.score2))
}

// formatScore formats the score as a string.
func formatScore(p1, p2 int) string {
	return string(rune('0'+p1)) + " : " + string(rune('0'+p2))
}

// Render renders the scene.
func (g *GameScene) Render() {
	g.ballCircle.Render()
	g.paddle1Rect.Render()
	g.paddle2Rect.Render()
	g.scoreLabel.Render()
}
