// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/aquasecurity/trivy-operator/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// CISKubeBenchReports returns a CISKubeBenchReportInformer.
	CISKubeBenchReports() CISKubeBenchReportInformer
	// ClusterComplianceDetailReports returns a ClusterComplianceDetailReportInformer.
	ClusterComplianceDetailReports() ClusterComplianceDetailReportInformer
	// ClusterComplianceReports returns a ClusterComplianceReportInformer.
	ClusterComplianceReports() ClusterComplianceReportInformer
	// ClusterConfigAuditReports returns a ClusterConfigAuditReportInformer.
	ClusterConfigAuditReports() ClusterConfigAuditReportInformer
	// ClusterVulnerabilityReports returns a ClusterVulnerabilityReportInformer.
	ClusterVulnerabilityReports() ClusterVulnerabilityReportInformer
	// ConfigAuditReports returns a ConfigAuditReportInformer.
	ConfigAuditReports() ConfigAuditReportInformer
	// VulnerabilityReports returns a VulnerabilityReportInformer.
	VulnerabilityReports() VulnerabilityReportInformer
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

// CISKubeBenchReports returns a CISKubeBenchReportInformer.
func (v *version) CISKubeBenchReports() CISKubeBenchReportInformer {
	return &cISKubeBenchReportInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterComplianceDetailReports returns a ClusterComplianceDetailReportInformer.
func (v *version) ClusterComplianceDetailReports() ClusterComplianceDetailReportInformer {
	return &clusterComplianceDetailReportInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterComplianceReports returns a ClusterComplianceReportInformer.
func (v *version) ClusterComplianceReports() ClusterComplianceReportInformer {
	return &clusterComplianceReportInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterConfigAuditReports returns a ClusterConfigAuditReportInformer.
func (v *version) ClusterConfigAuditReports() ClusterConfigAuditReportInformer {
	return &clusterConfigAuditReportInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterVulnerabilityReports returns a ClusterVulnerabilityReportInformer.
func (v *version) ClusterVulnerabilityReports() ClusterVulnerabilityReportInformer {
	return &clusterVulnerabilityReportInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ConfigAuditReports returns a ConfigAuditReportInformer.
func (v *version) ConfigAuditReports() ConfigAuditReportInformer {
	return &configAuditReportInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VulnerabilityReports returns a VulnerabilityReportInformer.
func (v *version) VulnerabilityReports() VulnerabilityReportInformer {
	return &vulnerabilityReportInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
