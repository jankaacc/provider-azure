/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this SentinelAlertRuleFusion
func (mg *SentinelAlertRuleFusion) GetTerraformResourceType() string {
	return "azurerm_sentinel_alert_rule_fusion"
}

// GetConnectionDetailsMapping for this SentinelAlertRuleFusion
func (tr *SentinelAlertRuleFusion) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelAlertRuleFusion
func (tr *SentinelAlertRuleFusion) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelAlertRuleFusion
func (tr *SentinelAlertRuleFusion) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelAlertRuleFusion
func (tr *SentinelAlertRuleFusion) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelAlertRuleFusion
func (tr *SentinelAlertRuleFusion) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelAlertRuleFusion
func (tr *SentinelAlertRuleFusion) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this SentinelAlertRuleFusion
func (tr *SentinelAlertRuleFusion) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this SentinelAlertRuleFusion using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelAlertRuleFusion) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelAlertRuleFusionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelAlertRuleFusion) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SentinelAlertRuleMachineLearningBehaviorAnalytics
func (mg *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetTerraformResourceType() string {
	return "azurerm_sentinel_alert_rule_machine_learning_behavior_analytics"
}

// GetConnectionDetailsMapping for this SentinelAlertRuleMachineLearningBehaviorAnalytics
func (tr *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelAlertRuleMachineLearningBehaviorAnalytics
func (tr *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelAlertRuleMachineLearningBehaviorAnalytics
func (tr *SentinelAlertRuleMachineLearningBehaviorAnalytics) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelAlertRuleMachineLearningBehaviorAnalytics
func (tr *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelAlertRuleMachineLearningBehaviorAnalytics
func (tr *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelAlertRuleMachineLearningBehaviorAnalytics
func (tr *SentinelAlertRuleMachineLearningBehaviorAnalytics) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this SentinelAlertRuleMachineLearningBehaviorAnalytics
func (tr *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this SentinelAlertRuleMachineLearningBehaviorAnalytics using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelAlertRuleMachineLearningBehaviorAnalytics) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelAlertRuleMachineLearningBehaviorAnalyticsParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelAlertRuleMachineLearningBehaviorAnalytics) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SentinelAlertRuleMSSecurityIncident
func (mg *SentinelAlertRuleMSSecurityIncident) GetTerraformResourceType() string {
	return "azurerm_sentinel_alert_rule_ms_security_incident"
}

// GetConnectionDetailsMapping for this SentinelAlertRuleMSSecurityIncident
func (tr *SentinelAlertRuleMSSecurityIncident) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelAlertRuleMSSecurityIncident
func (tr *SentinelAlertRuleMSSecurityIncident) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelAlertRuleMSSecurityIncident
func (tr *SentinelAlertRuleMSSecurityIncident) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelAlertRuleMSSecurityIncident
func (tr *SentinelAlertRuleMSSecurityIncident) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelAlertRuleMSSecurityIncident
func (tr *SentinelAlertRuleMSSecurityIncident) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelAlertRuleMSSecurityIncident
func (tr *SentinelAlertRuleMSSecurityIncident) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this SentinelAlertRuleMSSecurityIncident
func (tr *SentinelAlertRuleMSSecurityIncident) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this SentinelAlertRuleMSSecurityIncident using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelAlertRuleMSSecurityIncident) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelAlertRuleMSSecurityIncidentParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelAlertRuleMSSecurityIncident) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SentinelAutomationRule
func (mg *SentinelAutomationRule) GetTerraformResourceType() string {
	return "azurerm_sentinel_automation_rule"
}

// GetConnectionDetailsMapping for this SentinelAutomationRule
func (tr *SentinelAutomationRule) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelAutomationRule
func (tr *SentinelAutomationRule) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelAutomationRule
func (tr *SentinelAutomationRule) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelAutomationRule
func (tr *SentinelAutomationRule) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelAutomationRule
func (tr *SentinelAutomationRule) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelAutomationRule
func (tr *SentinelAutomationRule) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this SentinelAutomationRule
func (tr *SentinelAutomationRule) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this SentinelAutomationRule using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelAutomationRule) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelAutomationRuleParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelAutomationRule) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this SentinelDataConnectorIOT
func (mg *SentinelDataConnectorIOT) GetTerraformResourceType() string {
	return "azurerm_sentinel_data_connector_iot"
}

// GetConnectionDetailsMapping for this SentinelDataConnectorIOT
func (tr *SentinelDataConnectorIOT) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelDataConnectorIOT
func (tr *SentinelDataConnectorIOT) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelDataConnectorIOT
func (tr *SentinelDataConnectorIOT) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelDataConnectorIOT
func (tr *SentinelDataConnectorIOT) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelDataConnectorIOT
func (tr *SentinelDataConnectorIOT) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelDataConnectorIOT
func (tr *SentinelDataConnectorIOT) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this SentinelDataConnectorIOT
func (tr *SentinelDataConnectorIOT) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this SentinelDataConnectorIOT using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelDataConnectorIOT) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelDataConnectorIOTParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelDataConnectorIOT) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SentinelLogAnalyticsWorkspaceOnboarding
func (mg *SentinelLogAnalyticsWorkspaceOnboarding) GetTerraformResourceType() string {
	return "azurerm_sentinel_log_analytics_workspace_onboarding"
}

// GetConnectionDetailsMapping for this SentinelLogAnalyticsWorkspaceOnboarding
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelLogAnalyticsWorkspaceOnboarding
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelLogAnalyticsWorkspaceOnboarding
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelLogAnalyticsWorkspaceOnboarding
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelLogAnalyticsWorkspaceOnboarding
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelLogAnalyticsWorkspaceOnboarding
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this SentinelLogAnalyticsWorkspaceOnboarding
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this SentinelLogAnalyticsWorkspaceOnboarding using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelLogAnalyticsWorkspaceOnboardingParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelLogAnalyticsWorkspaceOnboarding) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SentinelWatchlist
func (mg *SentinelWatchlist) GetTerraformResourceType() string {
	return "azurerm_sentinel_watchlist"
}

// GetConnectionDetailsMapping for this SentinelWatchlist
func (tr *SentinelWatchlist) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SentinelWatchlist
func (tr *SentinelWatchlist) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SentinelWatchlist
func (tr *SentinelWatchlist) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SentinelWatchlist
func (tr *SentinelWatchlist) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SentinelWatchlist
func (tr *SentinelWatchlist) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SentinelWatchlist
func (tr *SentinelWatchlist) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this SentinelWatchlist
func (tr *SentinelWatchlist) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this SentinelWatchlist using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SentinelWatchlist) LateInitialize(attrs []byte) (bool, error) {
	params := &SentinelWatchlistParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SentinelWatchlist) GetTerraformSchemaVersion() int {
	return 0
}
