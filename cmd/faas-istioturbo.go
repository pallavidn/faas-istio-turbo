package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/pallavidn/faas-istio/pkg"
	"github.com/pallavidn/faas-istio/pkg/conf"
)

func main() {
	// The default is to log to both of stderr and file
	// These arguments can be overloaded from the command-line args
	flag.Set("logtostderr", "false")
	flag.Set("alsologtostderr", "true")
	//flag.Set("log_dir", "/var/log")
	defer glog.Flush()

	args := conf.NewFaasIstioTurboArgs(flag.CommandLine)
	flag.Parse()

	glog.Info("Starting Faas Istio Turbo...")
	s, err := pkg.NewFaasIstioTAPService(args)

	if err != nil {
		glog.Fatal("Failed creating FaasIstioTurbo.: %v", err)
	}

	s.Start()
}
