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

package controllers

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	appv1alpha1 "github.com/example/webpage-operator/api/v1alpha1"
)

// WebPageReconciler reconciles a WebPage object
type WebPageReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=app.my.domain,resources=webpages,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=app.my.domain,resources=webpages/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=app.my.domain,resources=webpages/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the WebPage object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *WebPageReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	log.Info("Reconciling")

	// Fetch the WebPage instance
	webpage := &appv1alpha1.WebPage{}
	err := r.Client.Get(context.TODO(), req.NamespacedName, webpage)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	dep := r.defineDeployment(webpage)
	result, err := r.ensureDeployment(dep)
	if result != nil {
		return *result, err
	}
	log.Info("Ensured Deployment")

	svc := r.defineService(webpage)
	result, err = r.ensureService(svc)
	if result != nil {
		return *result, err
	}
	log.Info("Ensured Service")

	result, err = r.handleChanges(webpage)
	if result != nil {
		return *result, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *WebPageReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appv1alpha1.WebPage{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Complete(r)
}

func (r *WebPageReconciler) defineDeployment(w *appv1alpha1.WebPage) *appsv1.Deployment {
	var env []corev1.EnvVar
	size := int32(1)
	if w.Spec.Title != "" {
		env = append(env, corev1.EnvVar{
			Name:  "REACT_APP_TITLE",
			Value: w.Spec.Title,
		})
	}
	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: w.Namespace,
			Name:      w.Name + "-webpage",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &size,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels(w),
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels(w),
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:  "visitors-webui",
						Image: "jdob/visitors-webui:1.0.0",
						Ports: []corev1.ContainerPort{{
							ContainerPort: 3000,
						}},
						Env: env,
					}},
				},
			},
		},
	}
	_ = controllerutil.SetControllerReference(w, dep, r.Scheme)
	return dep
}

func (r *WebPageReconciler) ensureDeployment(dep *appsv1.Deployment) (*ctrl.Result, error) {
	found := &appsv1.Deployment{}
	err := r.Client.Get(context.TODO(), types.NamespacedName{
		Name:      dep.Name,
		Namespace: dep.Namespace,
	}, found)
	if err != nil {
		if errors.IsNotFound(err) {
			err = r.Client.Create(context.TODO(), dep)
			if err != nil {
				return &ctrl.Result{}, err
			}
		}
		return &ctrl.Result{}, err
	}

	return nil, nil
}

func (r *WebPageReconciler) defineService(w *appv1alpha1.WebPage) *corev1.Service {
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: w.Namespace,
			Name:      w.Name + "-service",
		},
		Spec: corev1.ServiceSpec{
			Selector: labels(w),
			Ports: []corev1.ServicePort{{
				Protocol:   corev1.ProtocolTCP,
				Port:       3000,
				TargetPort: intstr.FromInt(3000),
				NodePort:   30686,
			}},
			Type: corev1.ServiceTypeNodePort,
		},
	}

	_ = controllerutil.SetControllerReference(w, svc, r.Scheme)

	return svc
}

func (r *WebPageReconciler) ensureService(svc *corev1.Service) (*ctrl.Result, error) {
	found := &corev1.Service{}
	err := r.Client.Get(context.TODO(), types.NamespacedName{
		Namespace: svc.Namespace,
		Name:      svc.Name,
	}, found)
	if err != nil {
		if errors.IsNotFound(err) {
			err = r.Client.Create(context.TODO(), svc)
			if err != nil {
				return &ctrl.Result{}, err
			}
		}
		return &ctrl.Result{}, err
	}

	return nil, nil

}

func labels(w *appv1alpha1.WebPage) map[string]string {
	return map[string]string{
		"webpage_cr": w.Name,
	}
}

func (r *WebPageReconciler) handleChanges(w *appv1alpha1.WebPage) (*ctrl.Result, error) {
	found := &appsv1.Deployment{}
	err := r.Client.Get(context.TODO(), types.NamespacedName{
		Namespace: w.Namespace,
		Name:      w.Name + "-webpage",
	}, found)
	if err != nil {
		return &ctrl.Result{RequeueAfter: 5 * time.Second}, err
	}

	title := w.Spec.Title
	existing := (*found).Spec.Template.Spec.Containers[0].Env[0].Value

	if title != existing {
		(*found).Spec.Template.Spec.Containers[0].Env[0].Value = title
		err = r.Client.Update(context.TODO(), found)
		if err != nil {
			return &ctrl.Result{}, err
		}

		return &ctrl.Result{Requeue: true}, nil
	}
	return nil, nil
}
