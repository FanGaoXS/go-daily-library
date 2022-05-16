package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

func main() {
	// 1、创建watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln("NewWatcher failed: ", err)
	}
	// 2、记得关闭watcher
	defer watcher.Close()

	// 3、利用协程来处理监听事件
	done := make(chan bool)
	go func() {
		defer close(done)

		for {
			select {
			case event, ok := <-watcher.Events: // 监听到文件修改的回调
				if !ok {
					return
				}
				log.Printf("%s %s\n", event.Name, event.Op)
			case err, ok := <-watcher.Errors: // 监听到错误的回调
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	// 4、监听的路径
	watcher.Add("./")
	if err != nil {
		log.Fatal("Add failed:", err)
	}
	<-done

	// 在 './' 路径下创建、删除或者重命名文件即会触发事件
}
