"use client";

import { useState } from "react";
import { FileText, FolderOpen, Plus } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { OpenProjectDialog } from "@wailsjs/go/editor/Editor";

export function Welcome() {
  const [isLoading, setIsLoading] = useState(false);

  const handleOpenProject = async () => {
    setIsLoading(true);
    try {
      const path = await OpenProjectDialog();
      if (path) {
        console.log("Selected project path:", path);
        // TODO: Load the project from the selected path
        // This will dispatch actions to load the game config and update Redux state
      }
    } catch (error) {
      console.error("Failed to open project dialog:", error);
    } finally {
      setIsLoading(false);
    }
  };

  const recentProjects = [
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
  ];

  return (
    <div className="min-h-screen bg-linear-to-br from-background via-background to-muted/30 flex flex-col">
      {/* Header */}
      <header className="border-b bg-background/80 backdrop-blur-sm">
        <div className="h-16 flex items-center px-8 justify-between">
          <div className="flex items-center gap-3">
            <div className="w-10 h-10 rounded-lg bg-linear-to-br from-primary to-primary/60 flex items-center justify-center">
              <span className="text-white font-bold text-lg">N</span>
            </div>
            <div>
              <h1 className="text-xl font-bold">Nother</h1>
              <p className="text-xs text-muted-foreground">
                Game Engine Editor
              </p>
            </div>
          </div>
          <div className="text-right">
            <p className="text-sm text-muted-foreground">Version 0.1.0</p>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="flex-1 px-8 py-12">
        <div className="max-w-6xl mx-auto">
          {/* Hero Section */}
          <div className="mb-16">
            <h2 className="text-4xl font-bold mb-2">Welcome to Nother</h2>
            <p className="text-lg text-muted-foreground mb-8">
              Create 2D games with ease. Start a new project or continue where
              you left off.
            </p>

            {/* Action Buttons */}
            <div className="flex gap-4">
              <Button size="lg" className="gap-2">
                <Plus className="w-5 h-5" />
                New Project
              </Button>
              <Button
                size="lg"
                variant="outline"
                className="gap-2"
                onClick={handleOpenProject}
                disabled={isLoading}
              >
                <FolderOpen className="w-5 h-5" />
                {isLoading ? "Opening..." : "Open Project"}
              </Button>
            </div>
          </div>

          {/* Recent Projects Section */}
          {recentProjects.length > 0 && (
            <div className="mb-16">
              <div className="flex items-center gap-2 mb-6">
                <FileText className="w-5 h-5 text-muted-foreground" />
                <h3 className="text-2xl font-semibold">Recent Projects</h3>
              </div>

              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                {recentProjects.map((project) => (
                  <Card
                    key={project.path}
                    className="p-4 hover:bg-accent transition-colors cursor-pointer group"
                  >
                    <div className="flex items-start justify-between mb-3">
                      <div className="w-10 h-10 rounded-lg bg-muted flex items-center justify-center group-hover:bg-primary/20 transition-colors">
                        <FolderOpen className="w-5 h-5 text-muted-foreground group-hover:text-primary transition-colors" />
                      </div>
                      <span className="text-xs px-2 py-1 rounded-full bg-muted text-muted-foreground">
                        {project.lastOpened}
                      </span>
                    </div>
                    <h4 className="font-semibold text-foreground mb-1 group-hover:text-primary transition-colors">
                      {project.name}
                    </h4>
                    <p className="text-xs text-muted-foreground truncate">
                      {project.path}
                    </p>
                  </Card>
                ))}
              </div>
            </div>
          )}

          {/* Quick Start Section */}
          <div className="bg-muted/30 border rounded-lg p-8">
            <h3 className="text-xl font-semibold mb-6">Quick Start</h3>

            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {/* Getting Started */}
              <div className="space-y-3">
                <h4 className="font-semibold text-foreground">
                  Getting Started
                </h4>
                <ul className="text-sm text-muted-foreground space-y-2">
                  <li className="hover:text-foreground transition-colors cursor-pointer">
                    → View Documentation
                  </li>
                  <li className="hover:text-foreground transition-colors cursor-pointer">
                    → Create Your First Game
                  </li>
                  <li className="hover:text-foreground transition-colors cursor-pointer">
                    → Learn the Basics
                  </li>
                </ul>
              </div>

              {/* Examples */}
              <div className="space-y-3">
                <h4 className="font-semibold text-foreground">Examples</h4>
                <ul className="text-sm text-muted-foreground space-y-2">
                  <li className="hover:text-foreground transition-colors cursor-pointer">
                    → Pong Game
                  </li>
                  <li className="hover:text-foreground transition-colors cursor-pointer">
                    → Simple Platformer
                  </li>
                  <li className="hover:text-foreground transition-colors cursor-pointer">
                    → Particle Effects
                  </li>
                </ul>
              </div>

              {/* Resources */}
              <div className="space-y-3">
                <h4 className="font-semibold text-foreground">Resources</h4>
                <ul className="text-sm text-muted-foreground space-y-2">
                  <li className="hover:text-foreground transition-colors cursor-pointer">
                    → API Reference
                  </li>
                  <li className="hover:text-foreground transition-colors cursor-pointer">
                    → Community Forum
                  </li>
                  <li className="hover:text-foreground transition-colors cursor-pointer">
                    → Report Issues
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </main>

      {/* Footer */}
      <footer className="border-t bg-background/50">
        <div className="px-8 py-6 max-w-6xl mx-auto flex items-center justify-between text-sm text-muted-foreground">
          <p>© 2024 Nother Game Engine. All rights reserved.</p>
          <div className="flex gap-6">
            <a href="#" className="hover:text-foreground transition-colors">
              Documentation
            </a>
            <a href="#" className="hover:text-foreground transition-colors">
              GitHub
            </a>
            <a href="#" className="hover:text-foreground transition-colors">
              Contact
            </a>
          </div>
        </div>
      </footer>
    </div>
  );
}
