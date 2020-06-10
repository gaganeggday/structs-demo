package main

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetStruct() v1.Pod {
	return v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "mem-example",
			Name:      "memory-demo",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				v1.Container{
					Name:  "memory-demo-ctr",
					Image: "polinux/stress",
					Command: []string{
						"stress",
					},
				},
			},
		},
	}
}

func TestStruct(t *testing.T) {
	want := GetStruct()
	got := createPod()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("The errors didnt match")
	}

}
