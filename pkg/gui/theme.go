package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Theme holds default styling and behavior for GUI components.
type Theme struct {
	// Default fonts
	DefaultFont     *rl.Font
	DefaultFontSize float32
	DefaultSpacing  float32

	// Default colors
	BackgroundColor rl.Color
	ForegroundColor rl.Color
	BorderColor     rl.Color
	AccentColor     rl.Color

	// Button colors
	ButtonNormalColor   rl.Color
	ButtonHoverColor    rl.Color
	ButtonPressedColor  rl.Color
	ButtonDisabledColor rl.Color

	// TextBox colors
	TextBoxBgColor      rl.Color
	TextBoxBorderColor  rl.Color
	TextBoxFocusedColor rl.Color

	// Default sizes
	DefaultPadding     float32
	DefaultBorderWidth float32
}

// NewDefaultTheme creates a theme with sensible defaults.
func NewDefaultTheme() *Theme {
	return &Theme{
		DefaultFont:     nil, // nil means use raylib default
		DefaultFontSize: 16.0,
		DefaultSpacing:  1.0,

		BackgroundColor: rl.White,
		ForegroundColor: rl.Black,
		BorderColor:     rl.DarkGray,
		AccentColor:     rl.SkyBlue,

		ButtonNormalColor:   rl.LightGray,
		ButtonHoverColor:    rl.Gray,
		ButtonPressedColor:  rl.DarkGray,
		ButtonDisabledColor: rl.NewColor(100, 100, 100, 255),

		TextBoxBgColor:      rl.White,
		TextBoxBorderColor:  rl.Gray,
		TextBoxFocusedColor: rl.Blue,

		DefaultPadding:     8.0,
		DefaultBorderWidth: 2.0,
	}
}

// NewDarkTheme creates a dark-themed GUI.
func NewDarkTheme() *Theme {
	return &Theme{
		DefaultFont:     nil,
		DefaultFontSize: 16.0,
		DefaultSpacing:  1.0,

		BackgroundColor: rl.NewColor(30, 30, 30, 255),
		ForegroundColor: rl.White,
		BorderColor:     rl.NewColor(60, 60, 60, 255),
		AccentColor:     rl.NewColor(70, 130, 180, 255),

		ButtonNormalColor:   rl.NewColor(50, 50, 50, 255),
		ButtonHoverColor:    rl.NewColor(70, 70, 70, 255),
		ButtonPressedColor:  rl.NewColor(40, 40, 40, 255),
		ButtonDisabledColor: rl.NewColor(60, 60, 60, 255),

		TextBoxBgColor:      rl.NewColor(40, 40, 40, 255),
		TextBoxBorderColor:  rl.NewColor(80, 80, 80, 255),
		TextBoxFocusedColor: rl.NewColor(70, 130, 180, 255),

		DefaultPadding:     8.0,
		DefaultBorderWidth: 2.0,
	}
}
