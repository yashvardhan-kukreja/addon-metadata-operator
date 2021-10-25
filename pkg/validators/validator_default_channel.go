package validators

import (
	"fmt"

	"github.com/mt-sre/addon-metadata-operator/pkg/utils"
)

// ValidateDefaultChannel validates whether the 'defaultChannel' provided under an addon.yaml also exists under 'channels' field
func ValidateDefaultChannel(metabundle *utils.MetaBundle) (bool, error) {
	defaultChannel := metabundle.AddonMeta.DefaultChannel
	for _, channel := range metabundle.AddonMeta.Channels {
		if channel.Name == defaultChannel {
			return true, nil
		}
	}
	return false, fmt.Errorf("failed to validate")
}
