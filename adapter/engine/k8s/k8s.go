package k8s

import (
	"cowait/core/cluster"
	"cowait/core/task"

	"flag"
	"fmt"
	"path/filepath"

	"context"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

const EnvPodName = "COWAIT_K8S_POD"

type kube struct {
	*kubernetes.Clientset
	name      string
	namespace string
}

func New() cluster.Driver {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return &kube{
		Clientset: clientset,
		name:      "kubernetes",
		namespace: "default",
	}
}

func NewInCluster() (cluster.Driver, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	clients, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &kube{
		Clientset: clients,
		name:      "kubernetes",
		namespace: "default",
	}, nil
}

func (k *kube) Name() string {
	return k.name
}

func (k *kube) Spawn(ctx context.Context, id task.ID, spec *task.Spec) error {
	envdef, err := spec.ToEnv()
	if err != nil {
		return err
	}

	meta := metav1.ObjectMeta{
		Name:      string(id),
		Namespace: k.namespace,
	}
	pod := apiv1.PodSpec{
		RestartPolicy: apiv1.RestartPolicyNever,
		Containers: []apiv1.Container{
			{
				Name:            "task",
				Image:           spec.Image,
				ImagePullPolicy: apiv1.PullAlways,
				Env: []apiv1.EnvVar{
					{
						Name:  task.EnvTaskdef,
						Value: envdef,
					},
					{
						Name:  task.EnvTaskID,
						Value: string(id),
					},
				},
			},
		},
	}

	_, err = k.CoreV1().Pods(k.namespace).Create(ctx, &apiv1.Pod{
		ObjectMeta: meta,
		Spec:       pod,
	}, metav1.CreateOptions{})
	return err
}

func (k *kube) Kill(ctx context.Context, id task.ID) error {
	return k.CoreV1().Pods(k.namespace).Delete(ctx, string(id), metav1.DeleteOptions{})
}

func (k *kube) Poke(ctx context.Context, id task.ID) error {
	pod, err := k.CoreV1().Pods(k.namespace).Get(ctx, string(id), metav1.GetOptions{})
	if err != nil {
		return err
	}

	switch pod.Status.Phase {
	case apiv1.PodPending:
		// check why its pending
		for _, cond := range pod.Status.ContainerStatuses {
			switch cond.State.Waiting.Reason {
			case "ErrImagePull":
				fallthrough
			case "ImagePullBackOff":
				return fmt.Errorf("failed to pull container image")
			}
		}

	case apiv1.PodFailed:
		return fmt.Errorf("pod is in failed phase")
	}

	return nil
}
