package issue

import (
	"strconv"

	"github.com/pterm/pterm"
	"gitlab.com/jstang/rasic/types"
	"gitlab.com/jstang/rasic/types/plugins"
)

// open a new issue in the given project
func Open(client types.HttpClient, api plugins.Api, project string, token string, issue types.Vulnerabilities, packagetarget string, packagetype string) error {
	projectId, _ := strconv.Atoi(project)
	newIssue := types.Issue{Title: issue.Title, Description: generateMarkdown(issue, packagetarget, packagetype), Id: projectId}

	// TODO: allow to configure Severity
	if issue.Severity == "HIGH" {
		api.CreateIssue(client, project, token, newIssue)
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
