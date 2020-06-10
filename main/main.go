package main

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       "Pod",
		APIVersion: "v1",
	}
}

func createObjectType() metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Namespace: "mem-example",
		Name:      "memory-demo",
	}
}

func createPodSpec() v1.PodSpec {
	return v1.PodSpec{
		Containers: []v1.Container{
			v1.Container{
				Name:  "memory-demo-ctr",
				Image: "polinux/stress",
				Command: []string{
					"stress",
				},
			},
		},
	}
}

func createPod() v1.Pod {
	return v1.Pod{
		TypeMeta:   createTypeMeta(),
		ObjectMeta: createObjectType(),
		Spec:       createPodSpec(),
	}
}
