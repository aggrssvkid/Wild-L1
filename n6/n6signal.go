func sigTake() {
	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel,
		syscall.SIGINT,
		syscall.SIGQUIT)
	for {
		sig := <-signalChanel
		if sig == syscall.SIGINT {
			fmt.Println("Got ctrl+c.\nExit")
			os.Exit(0)
		}
	}
}

/*
Остановка горутины с помощью канала, который принимает сигналы.
Конкретная программа реализована в n4 папке.
*/