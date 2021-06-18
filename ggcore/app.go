package ggcore

type App interface {
	PushLayer(Layer)
	PushOverlay(Layer)
	GetWindow() Window
}

var app App

func SetApp(application App) {
	if app != nil {
		CoreError("Two instances of application not allowed")
	}
	app = application
}

func GetApp() App {
	return app
}
