package issue

import (
	"github.com/pterm/pterm"
	"gitlab.com/jstang/rasic/api"
	"gitlab.com/jstang/rasic/types"
)

// open a new issue in the given project
// we use glab cli to make this more easy
// TODO: remove glab dependency and use a custom api-call
func Open(client types.HttpClient, project string, token string, issue types.Vulnerabilities, packagetarget string, packagetype string) error {
	newIssue := &types.Issue{Title: issue.Title, Description: generateMarkdown(issue, packagetarget, packagetype)}
	// TODO: allow to configure Severity
	if issue.Severity == "CRITICAL" {
		_, err := api.CreateIssue(client, project, token, newIssue)
		// app := "glab"
		// arg0 := "issue"
		// arg1 := "create"
		// TODO: check if other label color is possible
		// arg2 := "-l cve, " + issue.Severity
		// arg3 := "-t " + issue.VulnerabilityID
		// arg4 := "-d " + generateMarkdown(issue, packagetarget, packagetype)
		// arg5 := "-R " + project
		// cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4, arg5)
		// stdout, err := cmd.Output()
		if err != nil {
			pterm.Error.Println(err.Error())
			return err
		}
		pterm.Info.Println("issue created")
	}

	return nil
}

// generate markdown to populate the new issue
// TODO: check if this can be done with a template
func generateMarkdown(issue types.Vulnerabilities, packagetarget string, packagetype string) string {
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

// create a newline
func newline(input string) string {
	return input + "\n"
}

// create a break-tag
func dobreak(input string) string {
	return input + "<br>"
}

// write bold
func bold(input string) string {
	return "**" + input + "**"
}

// open a code-style bash block
func consoleStart() string {
	return "```bash\n"
}

// end a code-style bash block
func consoleEnd() string {
	return "\n```"
}
