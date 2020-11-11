// Copyright 2020 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	nginxv1alpha1 "github.com/tsuru/nginx-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"

	extensionsv1alpha1 "github.com/tsuru/rpaas-operator/api/v1alpha1"
)

// NewScheme creates a scheme with Rpaas, Nginx and the default Kubernetes
// types (Pod, Deployment, PersistentVolumeClaim, etc) added into.
//
// NOTE: It panics whether some clientset scheme cannot be added to scheme.
func NewScheme() *runtime.Scheme {
	scheme := runtime.NewScheme()
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(nginxv1alpha1.AddToScheme(scheme))
	utilruntime.Must(extensionsv1alpha1.AddToScheme(scheme))
	// +kubebuilder:scaffold:scheme
	return scheme
}
