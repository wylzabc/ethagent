package main

import (
	"fmt"
	//"log"
	//"io/ioutil"
	"os"
	//"os/exec"
	"os/signal"
	"runtime"
	"syscall"

	//"github.com/golang/glog"
	"github.com/spf13/viper"

	"wylzabc/ethagent/common/log"
	httpserver "wylzabc/ethagent/httpserver"
)

var signalHandlerMap map[os.Signal]func()

func SigHupHandler() {
	fmt.Println("sig hup")
	log.Flush()
	os.Exit(0)
}

func SigIntHandler() {
	fmt.Println("sig int")
	log.Flush()
	os.Exit(0)
}

func SetSignalHandler(handler func(), sig os.Signal) {
	if signalHandlerMap == nil {
		signalHandlerMap = make(map[os.Signal]func())
	}
	signalHandlerMap[sig] = handler
}

func dealSignal() {
	SetSignalHandler(SigHupHandler, syscall.SIGHUP)
	SetSignalHandler(SigIntHandler, syscall.SIGINT)

	signalChan := make(chan os.Signal, 1)
	var signals []os.Signal
	for k := range signalHandlerMap {
		signals = append(signals, k)
	}
	signal.Notify(signalChan, signals...)
	go func() {
		for sig := range signalChan {
			if signalHandlerMap[sig] != nil {
				signalHandlerMap[sig]()
			}
		}
	}()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	log.SetupLogger(
		viper.GetString("log.level"),
		viper.GetString("log.filename"),
		viper.GetInt("log.maxsize"),
		viper.GetInt("log.maxrolls"),
		viper.GetString("log.errfilename"))

	go dealSignal()

	err = httpserver.Run()
	if err != nil {
		os.Exit(-1)
	}

	runtime.Goexit()
}
