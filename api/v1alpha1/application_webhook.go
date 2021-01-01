/*
Copyright 2020.

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

package v1alpha1

import (
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var applicationlog = logf.Log.WithName("application-resource")

func (r *Application) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-apps-martinheinz-dev-v1alpha1-application,mutating=true,failurePolicy=fail,sideEffects=None,groups=apps.martinheinz.dev,resources=applications,verbs=create;update,versions=v1alpha1,name=mapplication.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Defaulter = &Application{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Application) Default() {
	applicationlog.Info("default", "name", r.Name)

	// TODO Test this
	if r.Spec.Replicas == 0 {
		r.Spec.Replicas = 1
	}
}

// +kubebuilder:webhook:path=/validate-apps-martinheinz-dev-v1alpha1-application,mutating=false,failurePolicy=fail,sideEffects=None,groups=apps.martinheinz.dev,resources=applications,verbs=create;update,versions=v1alpha1,name=vapplication.kb.io,admissionReviewVersions={v1,v1beta1}

var _ webhook.Validator = &Application{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Application) ValidateCreate() error {
	applicationlog.Info("validate create", "name", r.Name)
	return r.validateApplication()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Application) ValidateUpdate(old runtime.Object) error {
	applicationlog.Info("validate update", "name", r.Name)
	return r.validateApplication()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Application) ValidateDelete() error {
	applicationlog.Info("validate delete", "name", r.Name)
	return nil
}

func (r *Application) validateApplication() error {
	var allErrs field.ErrorList
	if err := r.validateApplicationSpec(); err != nil {
		allErrs = append(allErrs, err)
	}
	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "apps.martinheinz.dev", Kind: "Application"},
		r.Name, allErrs)
}

func (r *Application) validateApplicationSpec() *field.Error {
	// Do some validation here... -> https://github.com/kubernetes-sigs/kubebuilder/blob/master/docs/book/src/cronjob-tutorial/testdata/project/api/v1/cronjob_webhook.go#L172
	return nil
}
