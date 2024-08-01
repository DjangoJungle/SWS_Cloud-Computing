package main

import (
    "context"
    "fmt"
    "time"

    "k8s.io/apimachinery/pkg/fields"
    "k8s.io/apimachinery/pkg/util/runtime"
    "k8s.io/apimachinery/pkg/util/wait"
    "k8s.io/client-go/informers"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/cache"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/util/workqueue"
    "k8s.io/klog/v2"
)

func main() {
    // Set up Kubernetes client
    kubeconfig := clientcmd.RecommendedHomeFile
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        klog.Fatalf("Error building kubeconfig: %s", err.Error())
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        klog.Fatalf("Error building Kubernetes clientset: %s", err.Error())
    }

    // Set up Informer for ConfigMaps
    factory := informers.NewSharedInformerFactory(clientset, time.Minute)
    informer := factory.Core().V1().ConfigMaps().Informer()

    // Set up WorkQueue
    queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

    // Event Handlers
    informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
        AddFunc: func(obj interface{}) {
            key, err := cache.MetaNamespaceKeyFunc(obj)
            if err == nil {
                queue.Add(key)
            }
        },
        UpdateFunc: func(oldObj, newObj interface{}) {
            key, err := cache.MetaNamespaceKeyFunc(newObj)
            if err == nil {
                queue.Add(key)
            }
        },
        DeleteFunc: func(obj interface{}) {
            key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
            if err == nil {
                queue.Add(key)
            }
        },
    })

    // Start Informer
    stopCh := make(chan struct{})
    defer close(stopCh)
    go factory.Start(stopCh)

    // Wait for cache to sync
    if !cache.WaitForCacheSync(stopCh, informer.HasSynced) {
        runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
        return
    }

    // Process items from WorkQueue
    wait.Until(func() {
        for processNextItem(queue) {
        }
    }, time.Second, stopCh)
}

func processNextItem(queue workqueue.RateLimitingInterface) bool {
    key, quit := queue.Get()
    if quit {
        return false
    }
    defer queue.Done(key)

    // Process the item
    fmt.Printf("Processing key: %s\n", key)
    queue.Forget(key)
    return true
}
