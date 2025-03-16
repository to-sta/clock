package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "clock",
		Width:  370,
		Height: 150,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Fullscreen: false,
		DisableResize: true,
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		AlwaysOnTop: true,
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent: true,
			BackdropType: windows.Acrylic,
			DisableWindowIcon: false,
			DisableFramelessWindowDecorations: true,
		},
		Debug: options.Debug{
            OpenInspectorOnStartup: true,
        },
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
