package provisioning

import (
	"time"

	"github.com/kyma-project/control-plane/components/kyma-environment-broker/internal"
	"github.com/kyma-project/control-plane/components/kyma-environment-broker/internal/process"
	"github.com/kyma-project/control-plane/components/kyma-environment-broker/internal/storage"
	"github.com/sirupsen/logrus"
)

const (
	KymaComponentNameKnativeProvisionerNatss = "knative-provisioner-natss"
	KebComponentNameKnativeProvisionerNatss  = "KnativeProvisionerNatss"
)

type KnativeProvisionerNatssStep struct {
	operationManager *process.ProvisionOperationManager
}

// ensure the interface is implemented
var _ Step = (*KnativeProvisionerNatssStep)(nil)

func NewKnativeProvisionerNatssStep(os storage.Operations) *KnativeProvisionerNatssStep {
	return &KnativeProvisionerNatssStep{
		operationManager: process.NewProvisionOperationManager(os),
	}
}

func (s *KnativeProvisionerNatssStep) Name() string {
	return "Provision Knative Nats Streaming"
}

func (s *KnativeProvisionerNatssStep) Run(operation internal.ProvisioningOperation, log logrus.FieldLogger) (internal.ProvisioningOperation, time.Duration, error) {
	parameters, err := operation.GetProvisioningParameters()
	if err != nil {
		log.Errorf("cannot fetch provisioning parameters from operation: %s", err)
		return s.operationManager.OperationFailed(operation, "invalid operation provisioning parameters")
	}
	log.Infof("Provisioning for PlanID: %s", parameters.PlanID)
	operation.InputCreator.EnableComponent(KebComponentNameKnativeProvisionerNatss)
	return operation, 0, nil
}
