"use client";

import { useState } from "react";
import { Box, Camera, Gamepad2, Image, Zap, Skull } from "lucide-react";
import {
  ResizablePanelGroup,
  ResizablePanel,
  ResizableHandle,
} from "@/components/ui/resizable";
import { Separator } from "@/components/ui/separator";
import {
  ContextMenu,
  ContextMenuTrigger,
  ContextMenuContent,
  ContextMenuItem,
} from "@/components/ui/context-menu";
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
} from "@/components/ui/dropdown-menu";
import { Dialog, DialogContent } from "@/components/ui/dialog";
import { TreeList, TreeListItem } from "@/components/tree/TreeList";
import { useAppSelector } from "@/stores/hooks";
import { Welcome } from "@/features/welcome/pages/Welcome";

export function EditorLayout() {
  const [selectedNode, setSelectedNode] = useState<string | null>(null);
  const game = useAppSelector((state) => state.editor.game);

  return (
    <>
      <div className="flex flex-col h-screen bg-background">
        {/* Header with OS-style menus */}
        <header className="border-b bg-background">
          <div className="h-14 flex items-center px-4 gap-8">
            <h1 className="text-lg font-semibold">Nother Editor</h1>
            <nav className="flex gap-0 text-sm">
              <DropdownMenu>
                <DropdownMenuTrigger className="px-3 py-2 hover:bg-accent hover:text-accent-foreground rounded cursor-pointer">
                  File
                </DropdownMenuTrigger>
                <DropdownMenuContent align="start">
                  <DropdownMenuItem>Load Project</DropdownMenuItem>
                  <DropdownMenuItem>Save Project</DropdownMenuItem>
                  <DropdownMenuItem>Exit</DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>

              <DropdownMenu>
                <DropdownMenuTrigger className="px-3 py-2 hover:bg-accent hover:text-accent-foreground rounded cursor-pointer">
                  Edit
                </DropdownMenuTrigger>
                <DropdownMenuContent align="start">
                  <DropdownMenuItem>Undo</DropdownMenuItem>
                  <DropdownMenuItem>Redo</DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>

              <DropdownMenu>
                <DropdownMenuTrigger className="px-3 py-2 hover:bg-accent hover:text-accent-foreground rounded cursor-pointer">
                  View
                </DropdownMenuTrigger>
                <DropdownMenuContent align="start">
                  <DropdownMenuItem>Zoom In</DropdownMenuItem>
                  <DropdownMenuItem>Zoom Out</DropdownMenuItem>
                  <DropdownMenuItem>Reset Zoom</DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>

              <DropdownMenu>
                <DropdownMenuTrigger className="px-3 py-2 hover:bg-accent hover:text-accent-foreground rounded cursor-pointer">
                  Help
                </DropdownMenuTrigger>
                <DropdownMenuContent align="start">
                  <DropdownMenuItem>About</DropdownMenuItem>
                  <DropdownMenuItem>Documentation</DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            </nav>
          </div>
        </header>

        {/* Main content area with resizable panels */}
        <ResizablePanelGroup direction="horizontal" className="flex-1">
          {/* Left Panel: Node Tree + Asset Navigator */}
          <ResizablePanel defaultSize={20} minSize={15}>
            <ResizablePanelGroup direction="vertical" className="h-full">
              {/* Node Tree Section */}
              <ResizablePanel defaultSize={60} minSize={30}>
                <div className="h-full flex flex-col bg-background border-b">
                  <div className="h-10 border-b border-border flex items-center px-4 bg-muted/50">
                    <h3 className="text-sm font-semibold">Scene Nodes</h3>
                  </div>
                  <div className="flex-1 overflow-auto px-2 py-2">
                    <div className="space-y-1">
                      <NodeTreeItem
                        label="Scene"
                        icon={Box}
                        onSelect={setSelectedNode}
                        isSelected={selectedNode === "scene"}
                      />
                      <div className="pl-4 space-y-1">
                        <NodeTreeItem
                          label="Camera"
                          icon={Camera}
                          onSelect={setSelectedNode}
                          isSelected={selectedNode === "camera"}
                        />
                        <NodeTreeItem
                          label="Player"
                          icon={Gamepad2}
                          onSelect={setSelectedNode}
                          isSelected={selectedNode === "player"}
                        />
                        <div className="pl-4 space-y-1">
                          <NodeTreeItem
                            label="Sprite"
                            icon={Image}
                            onSelect={setSelectedNode}
                            isSelected={selectedNode === "sprite"}
                          />
                          <NodeTreeItem
                            label="Collider"
                            icon={Zap}
                            onSelect={setSelectedNode}
                            isSelected={selectedNode === "collider"}
                          />
                        </div>
                        <NodeTreeItem
                          label="Enemy"
                          icon={Skull}
                          onSelect={setSelectedNode}
                          isSelected={selectedNode === "enemy"}
                        />
                      </div>
                    </div>
                  </div>
                </div>
              </ResizablePanel>

              <ResizableHandle />

              {/* Asset Navigator Section */}
              <ResizablePanel defaultSize={40} minSize={20}>
                <div className="h-full flex flex-col bg-background">
                  <div className="h-10 border-b border-border flex items-center px-4 bg-muted/50">
                    <h3 className="text-sm font-semibold">Assets</h3>
                  </div>
                  <div className="flex-1 overflow-auto px-2 py-2">
                    <TreeList>
                      <TreeListItem>
                        <div className="px-2 py-1.5 rounded cursor-pointer text-sm text-foreground hover:bg-accent hover:text-accent-foreground transition-colors">
                          Sprites (12)
                        </div>
                        <TreeList>
                          <TreeListItem>
                            <div className="px-2 py-1.5 rounded cursor-pointer text-sm text-foreground hover:bg-accent hover:text-accent-foreground transition-colors">
                              player.png
                            </div>
                          </TreeListItem>
                          <TreeListItem>
                            <div className="px-2 py-1.5 rounded cursor-pointer text-sm text-foreground hover:bg-accent hover:text-accent-foreground transition-colors">
                              enemy.png
                            </div>
                          </TreeListItem>
                        </TreeList>
                      </TreeListItem>
                      <TreeListItem>
                        <div className="px-2 py-1.5 rounded cursor-pointer text-sm text-foreground hover:bg-accent hover:text-accent-foreground transition-colors">
                          Sounds (5)
                        </div>
                        <TreeList>
                          <TreeListItem>
                            <div className="px-2 py-1.5 rounded cursor-pointer text-sm text-foreground hover:bg-accent hover:text-accent-foreground transition-colors">
                              jump.wav
                            </div>
                          </TreeListItem>
                        </TreeList>
                      </TreeListItem>
                      <TreeListItem>
                        <div className="px-2 py-1.5 rounded cursor-pointer text-sm text-foreground hover:bg-accent hover:text-accent-foreground transition-colors">
                          Fonts (3)
                        </div>
                      </TreeListItem>
                      <TreeListItem>
                        <div className="px-2 py-1.5 rounded cursor-pointer text-sm text-foreground hover:bg-accent hover:text-accent-foreground transition-colors">
                          Scenes (2)
                        </div>
                      </TreeListItem>
                    </TreeList>
                  </div>
                </div>
              </ResizablePanel>
            </ResizablePanelGroup>
          </ResizablePanel>

          <ResizableHandle withHandle />

          {/* Center Panel: Scene Preview */}
          <ResizablePanel defaultSize={50} minSize={30}>
            <div className="h-full flex flex-col bg-slate-900 border-r">
              {/* Toolbar */}
              <div className="h-10 border-b border-border bg-background flex items-center px-3 gap-2">
                <button className="px-2 py-1 text-xs rounded bg-primary text-primary-foreground hover:bg-primary/90">
                  Play
                </button>
                <button className="px-2 py-1 text-xs rounded bg-secondary text-secondary-foreground hover:bg-secondary/90">
                  Pause
                </button>
                <Separator orientation="vertical" className="h-4" />
                <span className="text-xs text-muted-foreground">
                  Canvas View
                </span>
              </div>

              {/* Canvas Area */}
              <ContextMenu>
                <div className="flex-1 flex items-center justify-center overflow-hidden relative bg-linear-to-b from-slate-900 to-slate-950">
                  <ContextMenuTrigger className="absolute inset-0 flex items-center justify-center">
                    <div className="text-center text-muted-foreground">
                      <p className="mb-2">Game Preview</p>
                      <p className="text-xs">Right-click to add nodes</p>
                    </div>
                  </ContextMenuTrigger>

                  <ContextMenuContent>
                    <ContextMenuItem>Add Sprite</ContextMenuItem>
                    <ContextMenuItem>Add Empty</ContextMenuItem>
                    <ContextMenuItem>Add Camera</ContextMenuItem>
                  </ContextMenuContent>
                </div>
              </ContextMenu>
            </div>
          </ResizablePanel>

          <ResizableHandle withHandle />

          {/* Right Panel: Property Inspector */}
          <ResizablePanel defaultSize={25} minSize={20}>
            <div className="h-full flex flex-col bg-background border-l overflow-hidden">
              <div className="h-10 border-b border-border flex items-center px-4">
                <h3 className="text-sm font-semibold">
                  {selectedNode ? `Properties - ${selectedNode}` : "Properties"}
                </h3>
              </div>

              {selectedNode ? (
                <div className="flex-1 overflow-auto p-4 space-y-4">
                  {/* Transform Section */}
                  <div className="space-y-2">
                    <h4 className="text-xs font-semibold text-muted-foreground uppercase">
                      Transform
                    </h4>
                    <PropertyField label="Position" value="0, 0" />
                    <PropertyField label="Rotation" value="0°" />
                    <PropertyField label="Scale" value="1, 1" />
                  </div>

                  <Separator />

                  {/* Component Section */}
                  <div className="space-y-2">
                    <h4 className="text-xs font-semibold text-muted-foreground uppercase">
                      Components
                    </h4>
                    <div className="border rounded-md p-2 text-xs bg-muted/50">
                      <div className="flex justify-between items-center">
                        <span>Sprite Renderer</span>
                        <span className="text-xs text-muted-foreground">✕</span>
                      </div>
                    </div>
                    <div className="border rounded-md p-2 text-xs bg-muted/50">
                      <div className="flex justify-between items-center">
                        <span>Collider</span>
                        <span className="text-xs text-muted-foreground">✕</span>
                      </div>
                    </div>
                  </div>

                  <Separator />

                  {/* Add Component */}
                  <button className="w-full px-3 py-1.5 text-xs rounded bg-secondary text-secondary-foreground hover:bg-secondary/90">
                    + Add Component
                  </button>
                </div>
              ) : (
                <div className="flex-1 flex items-center justify-center text-muted-foreground text-sm">
                  Select a node to edit properties
                </div>
              )}
            </div>
          </ResizablePanel>
        </ResizablePanelGroup>
      </div>

      <Dialog open={game === null}>
        <DialogContent className="sm:max-w-4xl max-h-[90vh] overflow-y-auto p-0 border-none">
          <Welcome />
        </DialogContent>
      </Dialog>
    </>
  );
}

interface NodeTreeItemProps {
  label: string;
  icon: React.ComponentType<{ className?: string }>;
  onSelect: (id: string) => void;
  isSelected: boolean;
}

function NodeTreeItem({
  label,
  icon: Icon,
  onSelect,
  isSelected,
}: NodeTreeItemProps) {
  return (
    <div
      onClick={() => onSelect(label.toLowerCase())}
      className={`px-2 py-1.5 rounded cursor-pointer text-sm flex items-center gap-2 transition-colors ${
        isSelected
          ? "bg-primary text-primary-foreground"
          : "hover:bg-accent hover:text-accent-foreground text-foreground"
      }`}
    >
      <Icon className="w-4 h-4" />
      <span>{label}</span>
    </div>
  );
}

interface PropertyFieldProps {
  label: string;
  value: string;
}

function PropertyField({ label, value }: PropertyFieldProps) {
  return (
    <div className="flex justify-between items-center text-xs">
      <label className="text-muted-foreground">{label}</label>
      <input
        type="text"
        defaultValue={value}
        className="px-2 py-0.5 rounded border border-input bg-background text-foreground text-right w-20"
      />
    </div>
  );
}
