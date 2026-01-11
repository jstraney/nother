import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface SceneFileEntry {
  id: string;
  filePath: string;
}

export interface GameState {
  windowWidth: number;
  windowHeight: number;
  windowTitle: string;
  targetFPS: number;
  assetDirectory: string;
  defaultSceneID: string;
  sceneFiles: SceneFileEntry[];
}

export interface Node {
  id: string;
  name: string;
  enabled: boolean;
  children?: Node[];
}

export interface CurrentScene {
  id: string;
  paused: boolean;
  rootNode: Node | null;
}

interface EditorState {
  game: GameState | null;
  currentScene: CurrentScene | null;
  isLoading: boolean;
  error: string | null;
}

const initialState: EditorState = {
  game: null,
  currentScene: null,
  isLoading: false,
  error: null,
};

const editorSlice = createSlice({
  name: "editor",
  initialState,
  reducers: {
    setGame: (state, action: PayloadAction<GameState>) => {
      state.game = action.payload;
      state.error = null;
    },
    setCurrentScene: (state, action: PayloadAction<CurrentScene>) => {
      state.currentScene = action.payload;
      state.error = null;
    },
    updateScenePausedState: (state, action: PayloadAction<boolean>) => {
      if (state.currentScene) {
        state.currentScene.paused = action.payload;
      }
    },
    updateRootNode: (state, action: PayloadAction<Node | null>) => {
      if (state.currentScene) {
        state.currentScene.rootNode = action.payload;
      }
    },
    setLoading: (state, action: PayloadAction<boolean>) => {
      state.isLoading = action.payload;
    },
    setError: (state, action: PayloadAction<string | null>) => {
      state.error = action.payload;
      state.isLoading = false;
    },
    clearEditor: (state) => {
      state.game = null;
      state.currentScene = null;
      state.error = null;
      state.isLoading = false;
    },
  },
});

export const {
  setGame,
  setCurrentScene,
  updateScenePausedState,
  updateRootNode,
  setLoading,
  setError,
  clearEditor,
} = editorSlice.actions;

export default editorSlice.reducer;
