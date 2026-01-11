# Engine Serialization Process

The game engine uses Go's built-in JSON marshaling for serializing nodes and scene data. Node types like `BaseNode` and `Node2D` have public fields that are automatically serialized to JSON by Go's `encoding/json` package. This approach requires minimal boilerplate and leverages Go's standard library conventions.

When custom serialization behavior is needed—such as handling private fields, computing derived values, or transforming data during marshaling—node types can implement the `json.Marshaler` and `json.Unmarshaler` interfaces. This allows fine-grained control over how specific node types are converted to and from JSON without introducing a separate serialization abstraction layer.

The editor communicates scene data with the Go backend via Wails, sending and receiving JSON-serialized nodes. Nodes maintain their hierarchy and state through this process, with the editor able to modify node properties and have those changes reflected when nodes are unserialized back into the runtime engine.

# Engine Entrypoints

The game engine accepts a game configuration file that defines the game's settings, assets, and default scene. This game file serves as the primary entry point for initializing the runtime engine. The game configuration is deserialized from JSON and used to populate the `Game` struct with window settings, target FPS, asset filesystem references, and other initialization parameters.

Optionally, a scene file can be passed to override the game's default scene. This is useful for testing specific scenes in isolation without needing to load the entire game's initial scene hierarchy. When a custom scene is provided, it is deserialized and loaded instead of the default scene specified in the game configuration.

By separating game configuration from scene data, the engine allows flexibility in how games are initialized and tested. Developers can maintain a consistent game setup while quickly switching between different scenes for development and debugging purposes.

# Scene Parsing and Initialization

When a scene file is loaded, the engine deserializes the JSON into a node hierarchy. A scene file contains a root node definition along with all of its descendant nodes in a tree structure. Each node in the JSON includes a `type` field that identifies which node type to instantiate. The engine maintains a registry of node types and their corresponding factory functions, allowing it to dynamically construct the correct node type during deserialization.

Scene initialization proceeds recursively from the root node, building the hierarchy by instantiating each node type and establishing parent-child relationships as specified in the serialized data. As each node is created, its public fields are populated from the JSON (position, name, enabled state, etc.), and any custom initialization logic is invoked through the `Load()` method on the scene itself.

This approach decouples the scene file format from specific node implementations. New node types can be registered with the engine without modifying core deserialization logic. The registry pattern makes the system extensible: developers can define custom node types, register them with factory functions, and immediately use them in scenes without engine changes.
