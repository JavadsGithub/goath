package validators

import (
	"github.com/JavadsGithub/goath/auth"
)

type ClaimValidators map[string]auth.ClaimsValidator

func GetValidators() ClaimValidators {
	validators := make(ClaimValidators)

	validators["default"] = DefaultClaimsValidator

	return validators
}
