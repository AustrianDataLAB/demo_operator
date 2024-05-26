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

package v1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var icecreamlog = logf.Log.WithName("icecream-resource")

// SetupWebhookWithManager will setup the manager to manage the webhooks
func (r *Icecream) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-demo-lva-io-v1-icecream,mutating=false,failurePolicy=fail,sideEffects=None,groups=demo.lva.io,resources=icecreams,verbs=create;update,versions=v1,name=vicecream.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Icecream{}

// Define a list of valid flavours
var validFlavours = map[string]bool{
	"vanilla":    true,
	"chocolate":  true,
	"strawberry": true,
	// add more flavours here
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Icecream) ValidateCreate() (admission.Warnings, error) {
	icecreamlog.Info("validate create", "name", r.Name)

	// Check if the new Icecream object has a valid flavour
	if !validFlavours[r.Spec.Flavour] {
		return nil, fmt.Errorf("invalid flavour: %s", r.Spec.Flavour)
	}

	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Icecream) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	icecreamlog.Info("validate update", "name", r.Name)

	// Check if the new Icecream object has a valid flavour
	if !validFlavours[r.Spec.Flavour] {
		return nil, fmt.Errorf("invalid flavour: %s", r.Spec.Flavour)
	}

	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Icecream) ValidateDelete() (admission.Warnings, error) {
	icecreamlog.Info("validate delete", "name", r.Name)

	// Log the deletion of the Icecream object
	icecreamlog.Info("Icecream object has been deleted", "name", r.Name)

	return nil, nil
}
