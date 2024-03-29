package main

import (
	"fmt"
	"os"

	arg "github.com/alexflint/go-arg"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"

	akmv1alpha1 "gitlab.com/GeorgeRaven/authentik-manager/operator/api/v1alpha1"
	"gitlab.com/GeorgeRaven/authentik-manager/operator/controllers"
	"gitlab.com/GeorgeRaven/authentik-manager/operator/utils"
	//+kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(akmv1alpha1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func main() {
	o := utils.Opts{}
	arg.MustParse(&o)

	if o.Debug {
		p, err := utils.PrettyPrint(o)
		if err != nil {
			setupLog.Error(err, "Unable to print options")
			os.Exit(1)
		}
		fmt.Println(p)
	}
	fmt.Println(fmt.Sprintf("Starting Authentik-Manager (%v) for Authentik (%v)", o.SrcVersion, o.AppVersion))

	opts := zap.Options{
		Development: o.Debug,
	}

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{
		Scheme: scheme,
		//MetricsBindAddress:     o.MetricsAddr,
		Metrics: metricsserver.Options{
			BindAddress: o.MetricsAddr,
		},
		//Port:                   o.Port,
		HealthProbeBindAddress: o.ProbeAddr,
		LeaderElection:         o.EnableLeaderElection,
		LeaderElectionID:       o.LeaderElectionID,
		// Specified the namespace the leader "lease" resource belongs
		// this will also affect clusterwide searches by operator which we dont want
		// so we specify so that these roles do not need to be granted
		LeaderElectionNamespace: o.OperatorNamespace,
		//Namespace:               o.WatchedNamespace,
		//Cache: cache.Options{
		//	DefaultNamespaces: map[string]cache.Config{
		//		o.WatchedNamespace
		//	},
		//},
	})

	//mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
	//	Scheme: scheme,
	//	//MetricsBindAddress:     o.MetricsAddr,
	//	Metrics:                ctrlSrv.Options{},
	//	Port:                   o.Port,
	//	HealthProbeBindAddress: o.ProbeAddr,
	//	LeaderElection:         o.EnableLeaderElection,
	//	LeaderElectionID:       o.LeaderElectionID,
	//	// Specified the namespace the leader "lease" resource belongs
	//	// this will also affect clusterwide searches by operator which we dont want
	//	// so we specify so that these roles do not need to be granted
	//	LeaderElectionNamespace: o.OperatorNamespace,
	//	Namespace:               o.WatchedNamespace,
	//})
	if err != nil {
		setupLog.Error(err, "unable to create manager")
		os.Exit(1)
	}

	if err = (&controllers.AkReconciler{
		ControlBase: utils.ControlBase{
			Client: mgr.GetClient(),
			Scheme: mgr.GetScheme(),
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Ak")
		os.Exit(1)
	}
	if err = (&controllers.AkBlueprintReconciler{
		ControlBase: utils.ControlBase{
			Client: mgr.GetClient(),
			Scheme: mgr.GetScheme(),
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "AkBlueprint")
		os.Exit(1)
	}
	if err = (&controllers.OIDCReconciler{
		ControlBase: utils.ControlBase{
			Client: mgr.GetClient(),
			Scheme: mgr.GetScheme(),
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "OIDC")
		os.Exit(1)
	}
	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
