package k8s

import (
	"cowait/core"
	"flag"
	"path/filepath"

	"context"

	batchv1 "k8s.io/api/batch/v1"
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
}

func New() core.Cluster {
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
	}
}

func NewInCluster() (core.Cluster, error) {
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
	}, nil
}

func (k *kube) Spawn(ctx context.Context, def *core.TaskDef) (core.Task, error) {
	envdef, err := def.ToEnv()
	if err != nil {
		return nil, err
	}

	pod := apiv1.PodSpec{
		RestartPolicy: apiv1.RestartPolicyNever,
		Containers: []apiv1.Container{
			{
				Name:            "task",
				Image:           def.Image,
				ImagePullPolicy: apiv1.PullIfNotPresent,
				Env: []apiv1.EnvVar{
					{
						Name:  core.EnvTaskdef,
						Value: envdef,
					},
					{
						Name: EnvPodName,
						ValueFrom: &apiv1.EnvVarSource{
							FieldRef: &apiv1.ObjectFieldSelector{
								FieldPath: "metadata.name",
							},
						},
					},
				},
			},
		},
	}

	completions := int32(1)
	var deadline *int64

	job := batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      def.Name,
			Namespace: def.Namespace,
		},
		Spec: batchv1.JobSpec{
			Completions:           &completions,
			ActiveDeadlineSeconds: deadline,
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: def.Namespace,
				},
				Spec: pod,
			},
		},
	}

	// create kubernetes job
	taskjob, err := k.BatchV1().
		Jobs(def.Namespace).
		Create(ctx, &job, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return &task{
		job: taskjob.Name,
	}, nil
}
