//
// Copyright 2021 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package resources

import (
	"context"
	"fmt"
	"strings"

	routesv1 "github.com/openshift/api/route/v1"

	operatorsv1alpha1 "github.com/IBM/ibm-commonui-operator/api/v1alpha1"
	version "github.com/IBM/ibm-commonui-operator/version"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

const affirm = "true"

func createConfigMap(ctx context.Context, client client.Client, cm *corev1.ConfigMap, instance *operatorsv1alpha1.CommonWebUI, needToRequeue *bool) error {
	reqLogger := log.WithValues("func", "createConfigMap", "instance.Name", instance.Name, "configmap.Name", cm.Name)

	err := controllerutil.SetControllerReference(instance, cm, client.Scheme())
	if err != nil {
		reqLogger.Error(err, fmt.Sprintf("Failed to set owner for configmap: %s", cm.Name))
		return err
	}

	reqLogger.Info(fmt.Sprintf("Creating configmap: %s", cm.Name))
	err = client.Create(ctx, cm)
	if err != nil {
		reqLogger.Error(err, fmt.Sprintf("Failed to create configmap: %s", cm.Name))
		return err
	}

	// Created log4js configmap successfully, requiring a requeue
	*needToRequeue = true

	return nil
}

func ReconcileLog4jsConfigMap(ctx context.Context, client client.Client, instance *operatorsv1alpha1.CommonWebUI, needToRequeue *bool) error {
	reqLogger := log.WithValues("func", "reconcileLog4jsConfigMap", "instance.Name", instance.Name, "instance.Namespace", instance.Namespace)
	reqLogger.Info("Reconciling log4js configmap")

	cm := &corev1.ConfigMap{}

	// Check if the log4js configmap already exists, if not create a new one
	err := client.Get(ctx, types.NamespacedName{Name: Log4jsConfigMapName, Namespace: instance.Namespace}, cm)
	if err != nil {
		if errors.IsNotFound(err) {
			metaLabels := LabelsForMetadata(Log4jsConfigMapName)
			cm = &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      Log4jsConfigMapName,
					Namespace: instance.Namespace,
					Labels:    metaLabels,
				},
				Data: Log4jsConfigMapData,
			}

			err = createConfigMap(ctx, client, cm, instance, needToRequeue)
			if err != nil {
				return err
			}
		} else {
			reqLogger.Error(err, "Failed to get log4js configmap")
			return err
		}
	}

	return nil
}

func ZenLeftNavExtensionsConfigMap(namespace string, data map[string]string) *corev1.ConfigMap {
	reqLogger := log.WithValues("func", "ZenLeftNavExtensionsConfigMap")
	reqLogger.Info("CS??? Entry")
	metaLabels := LabelsForMetadata(ZenLeftNavExtensionsConfigMapName)
	metaLabels["icpdata_addon"] = affirm
	configmap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ZenLeftNavExtensionsConfigMapName,
			Namespace: namespace,
			Labels:    metaLabels,
		},
		Data: data,
	}
	return configmap
}

func ZenCardExtensionsConfigMap(name string, namespace string, version string, data map[string]string) *corev1.ConfigMap {
	reqLogger := log.WithValues("func", "ZenCardExtensionsConfigMap")
	reqLogger.Info("CS??? Entry")
	metaLabels := LabelsForMetadata(name)
	metaLabels["icpdata_addon"] = affirm
	metaLabels["icpdata_addon_version"] = "v" + version
	configmap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    metaLabels,
		},
		Data: data,
	}
	return configmap
}

func CommonWebUIConfigMap(namespace string) *corev1.ConfigMap {
	reqLogger := log.WithValues("func", "CommonWebUIConfigMap")
	reqLogger.Info("CS??? Entry")
	metaLabels := LabelsForMetadata(CommonConfigMapName)
	configmap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      CommonConfigMapName,
			Namespace: namespace,
			Labels:    metaLabels,
		},
	}

	return configmap
}

func ReconcileZenProductConfigMap(ctx context.Context, client client.Client, request ctrl.Request, needToRequeue *bool) error {
	reqLogger := log.WithValues("func", "ReconcileZenProductConfigMap", "request", request)

	productConfigMap := &corev1.ConfigMap{}
	err := client.Get(ctx, types.NamespacedName{Name: ZenProductConfigMapName, Namespace: request.Namespace}, productConfigMap)
	if err != nil {
		if errors.IsNotFound(err) {
			//The zen product-configmap doesn't exist, if this is a zen adminhub instance then there is nothing to do, however
			//in this case the product-configmap will already exist so in theory we should never get here in that condition
			reqLogger.Info("Zen product configmap " + ZenProductConfigMapName + " was not found - assumed to be deleted")
			return nil
		}

		reqLogger.Error(err, "Failed to get zen product configmap "+ZenProductConfigMapName)
		return err
	}

	//If the requested fields are missing or not equal, then update the configmap
	update := false

	if productConfigMap.Data == nil {
		productConfigMap.Data = map[string]string{}
	}

	for key, element := range ZenPcmMap {
		zelement, ok := productConfigMap.Data[key]
		if !ok || element != zelement {
			productConfigMap.Data[key] = element
			update = true
			log.Info("updated zen product configmap "+ZenProductConfigMapName, "key", key, "value", element)
		}
	}

	if update {
		updateErr := client.Update(ctx, productConfigMap)
		if updateErr != nil {
			reqLogger.Info("Could not update zen product configmap "+ZenProductConfigMapName, updateErr)
			return updateErr
		}
		reqLogger.Info("Zen product configmap " + ZenProductConfigMapName + " updated with meta values")
		*needToRequeue = true
	} else {
		reqLogger.Info("No updates required for commonwebui in zen product configmap " + ZenProductConfigMapName)
	}

	return nil

}

func ReconcileConfigMapsZen(ctx context.Context, client client.Client, namespace, nameOfCM string) error {
	reqLogger := log.WithValues("func", "ReconcileConfigMapsZen")

	reqLogger.Info("Checking if config map: " + nameOfCM + " exists")
	// Check if the config map already exists, if not create a new one
	currentConfigMap := &corev1.ConfigMap{}
	err := client.Get(ctx, types.NamespacedName{Name: nameOfCM, Namespace: namespace}, currentConfigMap)
	if err != nil && errors.IsNotFound(err) {
		// Define a new ConfigMap
		newConfigMap := &corev1.ConfigMap{}
		if nameOfCM == ZenCardExtensionsConfigMapName {
			reqLogger.Info("Creating zen card extensions config map")
			var ExtensionsData = map[string]string{
				"nginx.conf": ZenNginxConfig,
				"extensions": ZenCardExtensions,
			}
			newConfigMap = ZenCardExtensionsConfigMap(ZenCardExtensionsConfigMapName, namespace, version.Version, ExtensionsData)
		} else if nameOfCM == ZenCardExtensionsConfigMapNameCncf {
			reqLogger.Info("Creating zen card extensions config map for CNCF")
			var ExtensionsData = map[string]string{
				"nginx.conf": ZenNginxConfig,
				"extensions": ZenCardExtensionsCncf,
			}
			newConfigMap = ZenCardExtensionsConfigMap(ZenCardExtensionsConfigMapNameCncf, namespace, version.Version, ExtensionsData)
		} else if nameOfCM == ZenQuickNavExtensionsConfigMapName {
			reqLogger.Info("Creating zen quick nav extensions config map")
			var ExtensionsData = map[string]string{
				"extensions": ZenQuickNavExtensions,
			}
			newConfigMap = ZenCardExtensionsConfigMap(ZenQuickNavExtensionsConfigMapName, namespace, version.Version, ExtensionsData)
		} else if nameOfCM == ZenWalkmeExtensionsConfigMapName {
			reqLogger.Info("Creating zen walkme extensions config map")
			var ExtensionsData = map[string]string{
				"extensions": ZenWalkmeExtensions,
			}
			newConfigMap = ZenCardExtensionsConfigMap(ZenWalkmeExtensionsConfigMapName, namespace, version.Version, ExtensionsData)
		} else if nameOfCM == CommonConfigMapName {
			reqLogger.Info("Creating common-web-ui-config config map")
			newConfigMap = CommonWebUIConfigMap(namespace)
		} else if nameOfCM == ZenLeftNavExtensionsConfigMapName {
			currentRoute := &routesv1.Route{}
			//Get the cp-console route and add it to the configmap below
			err2 := client.Get(ctx, types.NamespacedName{Name: "cp-console", Namespace: namespace}, currentRoute)
			if err2 != nil {
				reqLogger.Error(err2, "Failed to get route for cp-console, try again later")
				return err2
			}
			reqLogger.Info("Current route is: " + currentRoute.Spec.Host)

			var ExtensionsData = map[string]string{
				"extensions": strings.Replace(ZenLeftNavExtensionsConfigMapData, "/common-nav/dashboard", "https://"+currentRoute.Spec.Host+"/common-nav/dashboard", 1),
			}

			newConfigMap = ZenLeftNavExtensionsConfigMap(namespace, ExtensionsData)

		}

		reqLogger.Info("Creating a config map", "Namespace", newConfigMap.Namespace, "Name", newConfigMap.Name)
		err = client.Create(ctx, newConfigMap)
		if err != nil {
			reqLogger.Error(err, "Failed to create a config map", "Namespace", newConfigMap.Namespace, "Name", newConfigMap.Name)
			return err
		}
	} else if err != nil {
		reqLogger.Error(err, "Failed to get Zen Config map")
		return err
	}

	reqLogger.Info("Created config map", "Name", nameOfCM)

	return nil

}
