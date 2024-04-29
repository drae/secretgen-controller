// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "carvel.dev/secretgen-controller/pkg/client2/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// SecretExports returns a SecretExportInformer.
	SecretExports() SecretExportInformer
	// SecretImports returns a SecretImportInformer.
	SecretImports() SecretImportInformer
	// SecretTemplates returns a SecretTemplateInformer.
	SecretTemplates() SecretTemplateInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// SecretExports returns a SecretExportInformer.
func (v *version) SecretExports() SecretExportInformer {
	return &secretExportInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SecretImports returns a SecretImportInformer.
func (v *version) SecretImports() SecretImportInformer {
	return &secretImportInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SecretTemplates returns a SecretTemplateInformer.
func (v *version) SecretTemplates() SecretTemplateInformer {
	return &secretTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
