package broker

import (
	"strings"

	"github.com/pkg/errors"
)

const (
	KymaServiceID   = "47c9dcbf-ff30-448e-ab36-d3bad66ba281"
	KymaServiceName = "kymaruntime"
)

var planIDsMapping = map[string]string{
	AzurePlanName:      AzurePlanID,
	AzureLitePlanName:  AzureLitePlanID,
	AzureTrialPlanName: AzureTrialPlanID,
	GCPPlanName:        GCPPlanID,
	GcpTrialPlanName:   GcpTrialPlanID,
}

type KymaEnvironmentBroker struct {
	*ServicesEndpoint
	*ProvisionEndpoint
	*DeprovisionEndpoint
	*UpdateEndpoint
	*GetInstanceEndpoint
	*LastOperationEndpoint
	*BindEndpoint
	*UnbindEndpoint
	*GetBindingEndpoint
	*LastBindingOperationEndpoint
}

// Config represents configuration for broker
type Config struct {
	EnablePlans EnablePlans `envconfig:"default=azure"`
}

// EnablePlans defines the plans that should be available for provisioning
type EnablePlans []string

// Unmarshal provides custom parsing of enabled plans.
// Implements envconfig.Unmarshal interface.
func (m *EnablePlans) Unmarshal(in string) error {
	plans := strings.Split(in, ",")
	for _, name := range plans {
		if _, exists := planIDsMapping[name]; !exists {
			return errors.Errorf("unrecognized %v plan name ", name)
		}
	}

	*m = plans
	return nil
}
