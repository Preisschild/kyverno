/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfigurations

import (
	v1 "github.com/kyverno/kyverno/api/kyverno/v1"
	v1alpha2 "github.com/kyverno/kyverno/api/kyverno/v1alpha2"
	v1beta1 "github.com/kyverno/kyverno/api/kyverno/v1beta1"
	v2alpha1 "github.com/kyverno/kyverno/api/kyverno/v2alpha1"
	v2beta1 "github.com/kyverno/kyverno/api/kyverno/v2beta1"
	policyreportv1alpha2 "github.com/kyverno/kyverno/api/policyreport/v1alpha2"
	kyvernov1 "github.com/kyverno/kyverno/pkg/client/applyconfigurations/kyverno/v1"
	kyvernov1alpha2 "github.com/kyverno/kyverno/pkg/client/applyconfigurations/kyverno/v1alpha2"
	kyvernov1beta1 "github.com/kyverno/kyverno/pkg/client/applyconfigurations/kyverno/v1beta1"
	kyvernov2alpha1 "github.com/kyverno/kyverno/pkg/client/applyconfigurations/kyverno/v2alpha1"
	kyvernov2beta1 "github.com/kyverno/kyverno/pkg/client/applyconfigurations/kyverno/v2beta1"
	applyconfigurationspolicyreportv1alpha2 "github.com/kyverno/kyverno/pkg/client/applyconfigurations/policyreport/v1alpha2"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=kyverno.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("AnyAllConditions"):
		return &kyvernov1.AnyAllConditionsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("APICall"):
		return &kyvernov1.APICallApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Attestation"):
		return &kyvernov1.AttestationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Attestor"):
		return &kyvernov1.AttestorApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AttestorSet"):
		return &kyvernov1.AttestorSetApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AutogenStatus"):
		return &kyvernov1.AutogenStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CEL"):
		return &kyvernov1.CELApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CertificateAttestor"):
		return &kyvernov1.CertificateAttestorApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CloneFrom"):
		return &kyvernov1.CloneFromApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CloneList"):
		return &kyvernov1.CloneListApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterPolicy"):
		return &kyvernov1.ClusterPolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Condition"):
		return &kyvernov1.ConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ConfigMapReference"):
		return &kyvernov1.ConfigMapReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ContextEntry"):
		return &kyvernov1.ContextEntryApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CTLog"):
		return &kyvernov1.CTLogApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Deny"):
		return &kyvernov1.DenyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DryRunOption"):
		return &kyvernov1.DryRunOptionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ForEachMutation"):
		return &kyvernov1.ForEachMutationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ForEachValidation"):
		return &kyvernov1.ForEachValidationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Generation"):
		return &kyvernov1.GenerationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ImageExtractorConfig"):
		return &kyvernov1.ImageExtractorConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ImageRegistry"):
		return &kyvernov1.ImageRegistryApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ImageRegistryCredentials"):
		return &kyvernov1.ImageRegistryCredentialsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ImageVerification"):
		return &kyvernov1.ImageVerificationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KeylessAttestor"):
		return &kyvernov1.KeylessAttestorApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Manifests"):
		return &kyvernov1.ManifestsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MatchResources"):
		return &kyvernov1.MatchResourcesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Mutation"):
		return &kyvernov1.MutationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ObjectFieldBinding"):
		return &kyvernov1.ObjectFieldBindingApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PodSecurity"):
		return &kyvernov1.PodSecurityApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PodSecurityStandard"):
		return &kyvernov1.PodSecurityStandardApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Policy"):
		return &kyvernov1.PolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("PolicyStatus"):
		return &kyvernov1.PolicyStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Rekor"):
		return &kyvernov1.RekorApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RequestData"):
		return &kyvernov1.RequestDataApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ResourceDescription"):
		return &kyvernov1.ResourceDescriptionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ResourceFilter"):
		return &kyvernov1.ResourceFilterApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ResourceSpec"):
		return &kyvernov1.ResourceSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Rule"):
		return &kyvernov1.RuleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RuleCountStatus"):
		return &kyvernov1.RuleCountStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SecretReference"):
		return &kyvernov1.SecretReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServiceCall"):
		return &kyvernov1.ServiceCallApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Spec"):
		return &kyvernov1.SpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("StaticKeyAttestor"):
		return &kyvernov1.StaticKeyAttestorApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("TargetResourceSpec"):
		return &kyvernov1.TargetResourceSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("UserInfo"):
		return &kyvernov1.UserInfoApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Validation"):
		return &kyvernov1.ValidationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ValidationFailureActionOverride"):
		return &kyvernov1.ValidationFailureActionOverrideApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("Variable"):
		return &kyvernov1.VariableApplyConfiguration{}

		// Group=kyverno.io, Version=v1alpha2
	case v1alpha2.SchemeGroupVersion.WithKind("AdmissionReport"):
		return &kyvernov1alpha2.AdmissionReportApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("AdmissionReportSpec"):
		return &kyvernov1alpha2.AdmissionReportSpecApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("BackgroundScanReport"):
		return &kyvernov1alpha2.BackgroundScanReportApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("BackgroundScanReportSpec"):
		return &kyvernov1alpha2.BackgroundScanReportSpecApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("ClusterAdmissionReport"):
		return &kyvernov1alpha2.ClusterAdmissionReportApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("ClusterBackgroundScanReport"):
		return &kyvernov1alpha2.ClusterBackgroundScanReportApplyConfiguration{}

		// Group=kyverno.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithKind("AdmissionRequestInfoObject"):
		return &kyvernov1beta1.AdmissionRequestInfoObjectApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("RequestInfo"):
		return &kyvernov1beta1.RequestInfoApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("UpdateRequest"):
		return &kyvernov1beta1.UpdateRequestApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("UpdateRequestSpec"):
		return &kyvernov1beta1.UpdateRequestSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("UpdateRequestSpecContext"):
		return &kyvernov1beta1.UpdateRequestSpecContextApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("UpdateRequestStatus"):
		return &kyvernov1beta1.UpdateRequestStatusApplyConfiguration{}

		// Group=kyverno.io, Version=v2alpha1
	case v2alpha1.SchemeGroupVersion.WithKind("CleanupPolicy"):
		return &kyvernov2alpha1.CleanupPolicyApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("CleanupPolicySpec"):
		return &kyvernov2alpha1.CleanupPolicySpecApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("CleanupPolicyStatus"):
		return &kyvernov2alpha1.CleanupPolicyStatusApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("ClusterCleanupPolicy"):
		return &kyvernov2alpha1.ClusterCleanupPolicyApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("Exception"):
		return &kyvernov2alpha1.ExceptionApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("PolicyException"):
		return &kyvernov2alpha1.PolicyExceptionApplyConfiguration{}
	case v2alpha1.SchemeGroupVersion.WithKind("PolicyExceptionSpec"):
		return &kyvernov2alpha1.PolicyExceptionSpecApplyConfiguration{}

		// Group=kyverno.io, Version=v2beta1
	case v2beta1.SchemeGroupVersion.WithKind("AnyAllConditions"):
		return &kyvernov2beta1.AnyAllConditionsApplyConfiguration{}
	case v2beta1.SchemeGroupVersion.WithKind("ClusterPolicy"):
		return &kyvernov2beta1.ClusterPolicyApplyConfiguration{}
	case v2beta1.SchemeGroupVersion.WithKind("Condition"):
		return &kyvernov2beta1.ConditionApplyConfiguration{}
	case v2beta1.SchemeGroupVersion.WithKind("Deny"):
		return &kyvernov2beta1.DenyApplyConfiguration{}
	case v2beta1.SchemeGroupVersion.WithKind("ImageVerification"):
		return &kyvernov2beta1.ImageVerificationApplyConfiguration{}
	case v2beta1.SchemeGroupVersion.WithKind("MatchResources"):
		return &kyvernov2beta1.MatchResourcesApplyConfiguration{}
	case v2beta1.SchemeGroupVersion.WithKind("Policy"):
		return &kyvernov2beta1.PolicyApplyConfiguration{}
	case v2beta1.SchemeGroupVersion.WithKind("Rule"):
		return &kyvernov2beta1.RuleApplyConfiguration{}
	case v2beta1.SchemeGroupVersion.WithKind("Spec"):
		return &kyvernov2beta1.SpecApplyConfiguration{}
	case v2beta1.SchemeGroupVersion.WithKind("Validation"):
		return &kyvernov2beta1.ValidationApplyConfiguration{}

		// Group=wgpolicyk8s.io, Version=v1alpha2
	case policyreportv1alpha2.SchemeGroupVersion.WithKind("ClusterPolicyReport"):
		return &applyconfigurationspolicyreportv1alpha2.ClusterPolicyReportApplyConfiguration{}
	case policyreportv1alpha2.SchemeGroupVersion.WithKind("PolicyReport"):
		return &applyconfigurationspolicyreportv1alpha2.PolicyReportApplyConfiguration{}
	case policyreportv1alpha2.SchemeGroupVersion.WithKind("PolicyReportResult"):
		return &applyconfigurationspolicyreportv1alpha2.PolicyReportResultApplyConfiguration{}
	case policyreportv1alpha2.SchemeGroupVersion.WithKind("PolicyReportSummary"):
		return &applyconfigurationspolicyreportv1alpha2.PolicyReportSummaryApplyConfiguration{}

	}
	return nil
}
