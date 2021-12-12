package issue

import (
	"fmt"
	"os/exec"

	"tubenhirn.com/cve2issue/types"
)

func Open(issue *types.Vulnerabilities, packagetarget string, packagetype string) error {
	app := "glab"
	arg0 := "issue"
	arg1 := "create"
	arg2 := "-l cve"
	arg3 := "-t " + issue.VulnerabilityID
	arg4 := "-d " + generateMarkdown(issue, packagetarget, packagetype)
	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(string(stdout))

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
