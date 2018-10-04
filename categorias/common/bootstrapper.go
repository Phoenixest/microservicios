package common

func StartUp() {

	// Inicializa la variable AppConfig
	initConfig()
	// Arranca una sesion de MongoDB
	createDbSession()
}
