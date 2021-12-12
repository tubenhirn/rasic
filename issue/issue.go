package issue

import (
	"fmt"

	"tubenhirn.com/cve2issue/types"
)

func Open(issue *types.Vulnerabilities) {
	fmt.Printf(string(issue.VulnerabilityID))
}
