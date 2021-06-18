package gogame

func CreateApplication(clientApp ClientApplication) {
	app := newCoreApplication(clientApp)
	app.run()
}
