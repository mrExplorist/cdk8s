package k8s


// VolumeAttachmentSpec is the specification of a VolumeAttachment request.
type VolumeAttachmentSpecV1Alpha1 struct {
	// Attacher indicates the name of the volume driver that MUST handle this request.
	//
	// This is the name returned by GetPluginName().
	Attacher *string `field:"required" json:"attacher" yaml:"attacher"`
	// The node that the volume should be attached to.
	NodeName *string `field:"required" json:"nodeName" yaml:"nodeName"`
	// Source represents the volume that should be attached.
	Source *VolumeAttachmentSourceV1Alpha1 `field:"required" json:"source" yaml:"source"`
}

