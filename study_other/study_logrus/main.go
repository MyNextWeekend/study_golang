package main

func main() {
	InitLogrus()
	logger.Debug("this is debug")
	logger.Info("this is info")
	logger.Warning("this is warning")
	logger.Error("this is error")
}
