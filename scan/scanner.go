package scan

import (
	"encoding/json"
	"github.com/pterm/pterm"
	"io/ioutil"
	"os/exec"

	"tubenhirn.com/cve2issue/issue"
	"tubenhirn.com/cve2issue/types"
)

/**
  scan with trivy binary
  and save the output as result.json
 **/
func Scanner(project string, issues types.Issues) error {
	app := "trivy"
	arg0 := "-q"
	arg1 := "repo"
	arg2 := "--format=json"
	arg3 := "--output=result.json"
	arg4 := project
	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4)
	_, err := cmd.Output()
	if err != nil {
		pterm.Error.Printfln(err.Error())
		return err
	}

	result, err := ioutil.ReadFile("result.json")
	if err != nil {
		pterm.Error.Printf("Status: %s\n", "file read error")
	}

	var report types.CVEReport
	unmarshalerr := json.Unmarshal(result, &report)
	if unmarshalerr != nil {
		return unmarshalerr
	}

	for _, result := range report.Results {
		if len(result.Vulnerabilities) > 0 {
			for _, cve := range result.Vulnerabilities {
				// check for open issues
				exists := false
				for i := range issues {
					if issues[i].Title == cve.VulnerabilityID {
						// TODO: allow to control State
						// maybe check for label "wont-fix" on the closed issue - otherwise reopen it
						if issues[i].State == "opened" {
							pterm.Info.Println("open issue exists for " + cve.VulnerabilityID)
							exists = true
							break
						}
					}
				}
				if !exists {
					// open issue if no issuw present in thes current porject
					issue.Open(project, &cve, result.Target, result.Type)
				}

			}
		}
	}
	return nil
}
