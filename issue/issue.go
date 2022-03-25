package issue

import (
	"os/exec"

	"github.com/pterm/pterm"
	"tubenhirn.com/risc/types"
)

func Open(project string, issue *types.Vulnerabilities, packagetarget string, packagetype string) error {
	// TODO: allow to configure Severity
	if issue.Severity == "CRITICAL" {
		app := "glab"
		arg0 := "issue"
		arg1 := "create"
		// TODO: check if other label color is possible
		arg2 := "-l cve, " + issue.Severity
		arg3 := "-t " + issue.VulnerabilityID
		arg4 := "-d " + generateMarkdown(issue, packagetarget, packagetype)
		arg5 := "-R " + project
		cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4, arg5)
		stdout, err := cmd.Output()
		if err != nil {
			pterm.Error.Println(err.Error())
			return err
		}
		pterm.Info.Println(string(stdout))
	}

	return nil
}

func generateMarkdown(issue *types.Vulnerabilities, packagetarget string, packagetype string) string {
	markdown := newline("### " + issue.Title)
	markdown += newline(dobreak(issue.Description))
	markdown += newline(dobreak(issue.PrimaryURL))
	markdown += newline("### Severity")
	markdown += newline(dobreak(bold(issue.Severity)))
	markdown += newline("### Package-Information")
	markdown += consoleStart()
	markdown += newline("target=" + packagetarget)
	markdown += newline("type=" + packagetype)
	markdown += newline("packagename=" + issue.PkgName)
	markdown += newline("installed_version=" + issue.InstalledVersion)
	markdown += newline("fixed_version=" + issue.FixedVersion)
	markdown += consoleEnd()
	return markdown
}

func newline(input string) string {
	return input + "\n"
}

func dobreak(input string) string {
	return input + "<br>"
}

func bold(input string) string {
	return "**" + input + "**"
}

func consoleStart() string {
	return "```bash\n"
}
func consoleEnd() string {
	return "\n```"
}
