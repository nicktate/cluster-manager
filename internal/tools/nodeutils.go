package tools

import (
	corev1 "k8s.io/api/core/v1"

	provisioncsv3 "github.com/containership/cloud-agent/pkg/apis/provision.containership.io/v3"
)

// NodeIsTargetKubernetesVersion checks if the current node version matches the target version
// of the cluster upgrade that is being processed. This only checks that the
// kubelet is up to date, and does not check the static pods.
// NOTE: this should only be called with upgrades of type Kubernetes
func NodeIsTargetKubernetesVersion(cup *provisioncsv3.ClusterUpgrade, node *corev1.Node) bool {
	return node.Status.NodeInfo.KubeletVersion == cup.Spec.TargetVersion
}

// NodeIsReady returns true if the given node has a Ready status, else false.
// See https://kubernetes.io/docs/concepts/nodes/node/#condition for more info.
func NodeIsReady(node *corev1.Node) bool {
	for _, condition := range node.Status.Conditions {
		if condition.Type == corev1.NodeReady && condition.Status == corev1.ConditionTrue {
			return true
		}
	}
	return false
}