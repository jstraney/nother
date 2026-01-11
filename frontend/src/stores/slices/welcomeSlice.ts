import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface RecentProject {
  name: string;
  path: string;
  lastOpened: string;
}

interface WelcomeState {
  recentProjects: RecentProject[];
}

const initialState: WelcomeState = {
  recentProjects: [
    {
      name: "Space Shooter",
      path: "/projects/space-shooter",
      lastOpened: "2 days ago",
    },
    {
      name: "Platformer Demo",
      path: "/projects/platformer-demo",
      lastOpened: "1 week ago",
    },
    {
      name: "Puzzle Game",
      path: "/projects/puzzle-game",
      lastOpened: "2 weeks ago",
    },
  ],
};

const welcomeSlice = createSlice({
  name: "welcome",
  initialState,
  reducers: {
    addRecentProject: (state, action: PayloadAction<RecentProject>) => {
      state.recentProjects.unshift(action.payload);
      // Keep only the 10 most recent
      state.recentProjects = state.recentProjects.slice(0, 10);
    },
    removeRecentProject: (state, action: PayloadAction<string>) => {
      state.recentProjects = state.recentProjects.filter(
        (p) => p.path !== action.payload,
      );
    },
    setRecentProjects: (state, action: PayloadAction<RecentProject[]>) => {
      state.recentProjects = action.payload;
    },
  },
});

export const { addRecentProject, removeRecentProject, setRecentProjects } =
  welcomeSlice.actions;
export default welcomeSlice.reducer;
