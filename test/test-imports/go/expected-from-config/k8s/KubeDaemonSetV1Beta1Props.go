package k8s


// DEPRECATED - This group version of DaemonSet is deprecated by apps/v1beta2/DaemonSet.
//
// See the release notes for more information. DaemonSet represents the configuration of a daemon set.
type KubeDaemonSetV1Beta1Props struct {
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// The desired behavior of this daemon set.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *DaemonSetSpecV1Beta1 `field:"optional" json:"spec" yaml:"spec"`
}

