package engine

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/jstraney/nother/pkg/game"
	"github.com/jstraney/nother/pkg/scene"
)

var (
	ErrNoDefaultScene = errors.New("no default scene in game file")
	ErrUnknownSceneID = errors.New("unrecognized scened id")
)

type Engine struct {
	Game *game.Game
}

func New() *Engine {
	return &Engine{}
}

func (engine *Engine) LoadGameFromFile(gameFile fs.File) error {
	var b []byte
	var _, err = gameFile.Read(b)
	if err != nil {
		return fmt.Errorf("gameFile.Read err:%w", err)
	}
	var g *game.Game
	err = json.Unmarshal(b, &g)
	if err != nil {
		return fmt.Errorf("json.Unmarshal err:%w", err)
	}
	engine.LoadGame(g)
	return nil
}

func (engine *Engine) LoadGame(game *game.Game) {
	engine.Game = game
}

func (engine *Engine) PlayDefaultScene() error {
	if engine.Game.DefaultSceneID == "" {
		return ErrNoDefaultScene
	}
	return nil
}

func (engine *Engine) PlayGameScene(id string) error {
	var scene, err = engine.LoadGameSceneByID(id)
	if err != nil {
		return err
	}
	scene.Root()
	engine.Game.SceneManager().ChangeScene(id)
	return nil
}

func (engine *Engine) LoadGameSceneByID(id string) (scene.Scene, error) {
	var scene scene.Scene
	for _, sceneEntry := range engine.Game.Serial.SceneFiles {
		if sceneEntry.ID != id {
			continue
		}
		var b, err = os.ReadFile(sceneEntry.FilePath)
		if err != nil {
			return nil, fmt.Errorf("os.ReadFile err:%w", err)
		}
		err = json.Unmarshal(b, &scene)
		if err != nil {
			return nil, fmt.Errorf("json.Unmarshal err:%w", err)
		}

	}
	if scene == nil {
		return nil, ErrUnknownSceneID
	}
	return scene, nil
}
