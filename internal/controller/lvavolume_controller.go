/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	demooperatorv1 "demooperator/api/v1"
)

// LvaVolumeReconciler reconciles a LvaVolume object
type LvaVolumeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=demooperator.operator.caas-0002.beta.austrianopencloudcommunity.org,resources=lvavolumes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demooperator.operator.caas-0002.beta.austrianopencloudcommunity.org,resources=lvavolumes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=demooperator.operator.caas-0002.beta.austrianopencloudcommunity.org,resources=lvavolumes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the LvaVolume object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *LvaVolumeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	l.Info("You are executing the reconcile loop right now", "req", req)
	volume := &demooperatorv1.LvaVolume{}

	r.Get(ctx, req.NamespacedName, volume)
	l.Info("Volume=", "magnitude", volume.Spec.Magnitude)
	l.Info("VolumeStatus=", "magnitude", volume.Status.Magnitude)

	if volume.Spec.Magnitude != volume.Status.Magnitude {
		l.Info("im trying to overwrite the Status now")
		volume.Status.Magnitude = volume.Spec.Magnitude
		r.Status().Update(ctx, volume)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *LvaVolumeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demooperatorv1.LvaVolume{}).
		Complete(r)
}
