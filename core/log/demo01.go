package main

import "log"

func init() {
	// 日志前缀
	log.SetPrefix("TRACE: ")
	// log标志
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func demo0101() {

	log.Println("message")

	//// 需要使用 recover，否则打印调用栈后终止
	//log.Panicln("panic message")

	// 写入日志后，调用os.Exit(1)
	log.Fatalln("fatal message")

}
