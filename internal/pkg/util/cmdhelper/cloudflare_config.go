package cmdhelper

import (
	"github.com/spf13/cobra"

	configPkg "github.com/nitschmann/cfdns/internal/pkg/config"
	"github.com/nitschmann/cfdns/internal/pkg/model"
)

// GetCloudflareConfigByFlags initializes an model.CloudflareConfig out of a command and its' flags
func GetCloudflareConfigByFlags(cmd *cobra.Command) (*model.CloudflareConfig, error) {
	var cloudflareConfig *model.CloudflareConfig

	apiKey, err := cmd.Flags().GetString("api-key")
	if err != nil {
		return cloudflareConfig, err
	}

	email, err := cmd.Flags().GetString("email")
	if err != nil {
		return cloudflareConfig, err
	}

	cloudflareConfig = &model.CloudflareConfig{
		APIKey: apiKey,
		Email:  email,
	}

	if cloudflareConfig.APIKey == "" || cloudflareConfig.Email == "" {
		configProfileName, err := cmd.Flags().GetString("profile")
		if err != nil {
			return cloudflareConfig, err
		}

		if configProfile, ok := configPkg.GetProfiles()[configProfileName]; ok {
			cloudflareConfig.APIKey = configProfile.APIKey
			cloudflareConfig.Email = configProfile.Email
		}
	}

	return cloudflareConfig, nil
}
