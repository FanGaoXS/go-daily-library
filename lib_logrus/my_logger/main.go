package main

func main() {
	logger := New("info", true, "json")

	logger.Info("info msg")
}
