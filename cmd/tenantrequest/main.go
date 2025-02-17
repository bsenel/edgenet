package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/EdgeNet-project/edgenet/pkg/bootstrap"
	"github.com/EdgeNet-project/edgenet/pkg/controller/registration/v1alpha1/tenantrequest"
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

	// Start the controller to provide the functionalities of tenantrequest resource
	edgenetInformerFactory := informers.NewSharedInformerFactory(edgenetclientset, 0)

	controller := tenantrequest.NewController(kubeclientset,
		edgenetclientset,
		edgenetInformerFactory.Registration().V1alpha1().TenantRequests())

	edgenetInformerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}
}
