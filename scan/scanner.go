package scan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"

	"tubenhirn.com/cve2issue/issue"
	"tubenhirn.com/cve2issue/types"
)

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
		fmt.Println(err.Error())
		return err
	}

	result, err := ioutil.ReadFile("result.json")
	if err != nil {
		fmt.Printf("Status: %s\n", "file read error")
	}

	var report types.CVEReport
	unmarshalerr := json.Unmarshal(result, &report)
	if unmarshalerr != nil {
		return unmarshalerr
	}

	for _, result := range report.Results {
		if len(result.Vulnerabilities) > 0 {
			fmt.Println(result.Target)

			for _, cve := range result.Vulnerabilities {
				// check for open issues
				exists := false
				for i := range issues {
					if issues[i].Title == cve.VulnerabilityID {
						exists = true
						break
					}
				}
				if !exists {
					// open issue if new
					fmt.Println(cve.VulnerabilityID + " issue created")
					issue.Open(project, &cve, result.Target, result.Type)

				}

			}
		}
	}
	return nil
}
