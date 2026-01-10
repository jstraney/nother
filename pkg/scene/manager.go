package scene

import (
	"fmt"
	"sync"
)

// SceneManager manages multiple scenes, allowing switching between them
// and ensuring only one scene is active at a time.
type SceneManager struct {
	scenes  map[string]Scene
	current string
	mu      sync.RWMutex
}

// New creates a new SceneManager.
func New() *SceneManager {
	return &SceneManager{
		scenes:  make(map[string]Scene),
		current: "",
	}
}

// AddScene registers a scene with the given name.
// If a scene with this name already exists, it returns an error.
func (sm *SceneManager) AddScene(name string, scene Scene) error {
	if scene == nil {
		return fmt.Errorf("cannot add nil scene")
	}

	sm.mu.Lock()
	defer sm.mu.Unlock()

	if _, exists := sm.scenes[name]; exists {
		return fmt.Errorf("scene %q already exists", name)
	}

	sm.scenes[name] = scene
	return nil
}

// RemoveScene unloads and removes a scene with the given name.
// If the removed scene is currently active, it cannot be removed.
func (sm *SceneManager) RemoveScene(name string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if sm.current == name {
		return fmt.Errorf("cannot remove the active scene %q", name)
	}

	scene, exists := sm.scenes[name]
	if !exists {
		return fmt.Errorf("scene %q not found", name)
	}

	if err := scene.Unload(); err != nil {
		return fmt.Errorf("failed to unload scene %q: %w", name, err)
	}

	delete(sm.scenes, name)
	return nil
}

// GetScene returns the scene with the given name, or nil if not found.
func (sm *SceneManager) GetScene(name string) Scene {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.scenes[name]
}

// Current returns the name of the currently active scene, or an empty string if none is active.
func (sm *SceneManager) Current() string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.current
}

// ChangeScene switches to the scene with the given name.
// It pauses the current scene and resumes the new scene.
// If the new scene has never been loaded, it calls Load().
// Returns an error if the scene doesn't exist or if loading fails.
func (sm *SceneManager) ChangeScene(name string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	scene, exists := sm.scenes[name]
	if !exists {
		return fmt.Errorf("scene %q not found", name)
	}

	// Pause the current scene
	if sm.current != "" {
		if current := sm.scenes[sm.current]; current != nil {
			current.Pause()
		}
	}

	// Load the new scene if needed
	if sm.current != name {
		if err := scene.Load(); err != nil {
			// Resume the previous scene on failure
			if sm.current != "" && sm.scenes[sm.current] != nil {
				sm.scenes[sm.current].Resume()
			}
			return fmt.Errorf("failed to load scene %q: %w", name, err)
		}
	}

	// Resume the new scene
	scene.Resume()
	sm.current = name
	return nil
}

// Update updates the currently active scene.
func (sm *SceneManager) Update(deltaTime float64) {
	sm.mu.RLock()
	current := sm.scenes[sm.current]
	sm.mu.RUnlock()

	if current != nil {
		current.Update(deltaTime)
	}
}

// PauseCurrent pauses the currently active scene.
func (sm *SceneManager) PauseCurrent() {
	sm.mu.RLock()
	current := sm.scenes[sm.current]
	sm.mu.RUnlock()

	if current != nil {
		current.Pause()
	}
}

// ResumeCurrent resumes the currently active scene.
func (sm *SceneManager) ResumeCurrent() {
	sm.mu.RLock()
	current := sm.scenes[sm.current]
	sm.mu.RUnlock()

	if current != nil {
		current.Resume()
	}
}

// IsPausedCurrent returns whether the currently active scene is paused.
func (sm *SceneManager) IsPausedCurrent() bool {
	sm.mu.RLock()
	current := sm.scenes[sm.current]
	sm.mu.RUnlock()

	if current == nil {
		return false
	}
	return current.Paused()
}

// Shutdown unloads all scenes and clears the manager.
func (sm *SceneManager) Shutdown() error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	var lastErr error
	for name, scene := range sm.scenes {
		if err := scene.Unload(); err != nil {
			lastErr = fmt.Errorf("failed to unload scene %q: %w", name, err)
		}
	}

	sm.scenes = make(map[string]Scene)
	sm.current = ""
	return lastErr
}
