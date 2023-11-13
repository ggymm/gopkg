package gopkg

/*func main() {
	c := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go start()

EXIT:
	for {
		sig := <-sc

		// 信号处理
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			c = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	cleanup()

	time.Sleep(time.Second)
	os.Exit(c)
}

func start() {
	println("start")
}

func cleanup() {
	println("cleanup")
}*/
