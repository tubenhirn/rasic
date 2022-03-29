package scan

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"

	"github.com/pterm/pterm"

	"gitlab.com/jstang/rasic/issue"
	"gitlab.com/jstang/rasic/types"
)

/**
  scan with trivy binary
  and save the output as result.json
 **/
func Scanner(client types.HttpClient, project string, token string, issues types.Issues, ignorefile string) error {
	// find path to trivy binary
	binary, lookErr := exec.LookPath("trivy")
	if lookErr != nil {
		panic(lookErr)
	}
	// build args
	args := []string{"-q", "repo", "--ignorefile=" + ignorefile, "--format=json", "--output=result.json", project}

	// get current environment
	env := os.Environ()

	// exec trivy with args and env
	// execErr := syscall.Exec(binary, args, env)
	cmd := exec.Command(binary, args...)
	cmd.Env = env
	_, execErr := cmd.Output()
	if execErr != nil {
		// pterm.Error.Printfln(execErr.Error())
		return execErr
	}

	result, err := ioutil.ReadFile("result.json")
	if err != nil {
		pterm.Error.Printf("Status: %s\n", "file read error")
	}

	var report types.CVEReport
	unmarshalerr := json.Unmarshal(result, &report)
	if unmarshalerr != nil {
		pterm.Error.Println(unmarshalerr)
		return unmarshalerr
	}

	for _, result := range report.Results {
		if len(result.Vulnerabilities) > 0 {

			pterm.Warning.Println(strconv.Itoa(len(result.Vulnerabilities)) + " vulnerabilities found")

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
					issue.Open(client, project, token, cve, result.Target, result.Type)
				}

			}
		}
	}
	return nil
}
