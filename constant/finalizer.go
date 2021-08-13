package constant

const (
	// RsyncSourceProtectionFinalizer used to maintain resource created for rsync source.
	// This finalizer is used to cleanup the resource created for a rsync source cr.
	RsyncSourceProtectionFinalizer = "demo.io/rsync-source-protection"
)
