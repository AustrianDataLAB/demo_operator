/*
Copyright 2024.

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

	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	demov1 "github.com/AustrianDataLab/demooperator/api/v1"
)

// IcecreamReconciler reconciles a Icecream object
type IcecreamReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=demo.lva.io,resources=icecreams,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=demo.lva.io,resources=icecreams/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=demo.lva.io,resources=icecreams/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Icecream object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *IcecreamReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// Fetch the Icecream instance
	icecream := &demov1.Icecream{}
	err := r.Get(ctx, req.NamespacedName, icecream)
	if err != nil {
		if apierrors.IsNotFound(err) {
			// Object not found, return. Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return ctrl.Result{}, err
	}

	// List all Icecream instances
	icecreamList := &demov1.IcecreamList{}
	if err := r.List(ctx, icecreamList); err != nil {
		return ctrl.Result{}, err
	}

	// Count the number of Icecream instances with the same flavour and keep track of the newest one
	count := 0
	var newestIcecream *demov1.Icecream
	for _, i := range icecreamList.Items {
		for _, flavour := range i.Spec.Flavour {
			if contains(icecream.Spec.Flavour, flavour) {
				count++
				if newestIcecream == nil || i.CreationTimestamp.After(newestIcecream.CreationTimestamp.Time) {
					newestIcecream = &i
				}
				break
			}
		}
	}

	// If count > 3, set SoldOut to true and delete the newest Icecream of the same flavour
	if count > 3 {
		icecream.Status.SoldOut = true
		if err := r.Delete(ctx, newestIcecream); err != nil {
			return ctrl.Result{}, err
		}
	} else {
		icecream.Status.SoldOut = false
	}

	// Update the ToppingsCount in the status
	icecream.Status.ToppingsCount = len(icecream.Spec.Topping)

	// Update the Icecream status
	err = r.Status().Update(ctx, icecream)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *IcecreamReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&demov1.Icecream{}).
		Complete(r)
}

// contains checks if a slice contains a string
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
