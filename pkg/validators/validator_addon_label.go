package validators

import (
	"fmt"

	"github.com/mt-sre/addon-metadata-operator/pkg/utils"
)

// ValidateAddonLabel validates whether the 'label' field under an addon.yaml follows the format 'api.openshift.com/addon-<id>'
func ValidateAddonLabel(metabundle *utils.MetaBundle) (bool, error) {
	operatorName, label := metabundle.AddonMeta.OperatorName, metabundle.AddonMeta.Label
	if label != "api.openshift.com/addon-"+operatorName {
		return false, nil
	}

	return true, nil
}
