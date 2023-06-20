package signaturevalidator

import (
	"encoding/json"

	"bruno.works/urlshortener/types"
)

func IsSignatureValid(publicKey string, configFile []byte) (bool, error) {
	var config types.ConfigurationFile
	parseErr := json.Unmarshal(configFile, &config)

	if parseErr != nil {
		return false, parseErr
	}

	configWithoutSignature := json.

	return true, nil
}
