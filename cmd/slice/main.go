package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/EdgeNet-project/edgenet/pkg/bootstrap"
	"github.com/EdgeNet-project/edgenet/pkg/controller/core/v1alpha1/slice"
	informers "github.com/EdgeNet-project/edgenet/pkg/generated/informers/externalversions"
	"github.com/EdgeNet-project/edgenet/pkg/signals"

	"k8s.io/klog"
)

func main() {
	klog.InitFlags(nil)
	flag.String("kubeconfig-path", bootstrap.GetDefaultKubeconfigPath(), "Path to the kubeconfig file's directory")
	flag.Parse()

	stopCh := signals.SetupSignalHandler()
	var authentication string
	if authentication = strings.TrimSpace(os.Getenv("AUTHENTICATION_STRATEGY")); authentication != "kubeconfig" {
		authentication = "serviceaccount"
	}
	kubeclientset, err := bootstrap.CreateClientset(authentication)
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}
	edgenetclientset, err := bootstrap.CreateEdgeNetClientset(authentication)
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}

	// Start the controller to provide the functionalities of slice resource
	edgenetInformerFactory := informers.NewSharedInformerFactory(edgenetclientset, time.Second*30)

	controller := slice.NewController(kubeclientset,
		edgenetclientset,
		edgenetInformerFactory.Core().V1alpha1().SliceClaims(),
		edgenetInformerFactory.Core().V1alpha1().Slices())

	edgenetInformerFactory.Start(stopCh)

	if err = controller.Run(1, stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}
}
