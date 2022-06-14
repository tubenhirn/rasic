package core

import (
	"strconv"

	"github.com/pterm/pterm"
	"gitlab.com/jstang/rasic/types"
	"gitlab.com/jstang/rasic/types/plugins"
)

// open a new issue in the given project
func Template(project string, cve types.Vulnerabilities, packagetarget string, packagetype string) (types.RasicIssue, error) {
	projectID, _ := strconv.Atoi(project)

	var cveSeverity types.Severity
	cveSeverity = cve.Severity

	newIssue := types.RasicIssue{Title: cve.VulnerabilityID, Description: generateMarkdown(cve, packagetarget, packagetype), ID: projectID, Severity: cveSeverity, Labels: []string{cve.Severity.String()}}

	return newIssue, nil
}

// generate markdown to populate the new issue
// TODO: check if this can be done with a template
func generateMarkdown(issue types.Vulnerabilities, packagetarget string, packagetype string) string {
	markdown := newline("### " + issue.Title)
	markdown += newline(dobreak(issue.Description))
	markdown += newline(dobreak(issue.PrimaryURL))
	markdown += newline("### Severity")
	markdown += newline(dobreak(bold(issue.Severity.String())))
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

// open new issues using the current reporter
func OpenNewIssues(httpClient types.HTTPClient, reporterPlugin plugins.Reporter, project types.RasicProject, newIssues []types.RasicIssue, authToken string) {
	// get all issues for current project
	var projectIssues []types.RasicIssue
	projectIssues = reporterPlugin.GetIssues(httpClient, "projects", strconv.Itoa(project.ID), authToken)

	// check newIssues against projectIssues
	// if the issue does not exist in State="opened", create it with the current reporter
	for _, newIssue := range newIssues {
		issueExists := false
		for _, openIssue := range projectIssues {
			if newIssue.Title == openIssue.Title && openIssue.State == "opened" {
				issueExists = true
				break
			}
		}
		if !issueExists {
			reporterPlugin.CreateIssue(httpClient, strconv.Itoa(project.ID), authToken, newIssue)
			pterm.Info.Println("new issue opened for " + newIssue.Title + " - " + newIssue.Severity.String())
		}
	}
}
