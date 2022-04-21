package issue

import (
	"strconv"

	"gitlab.com/jstang/rasic/types"
)

// open a new issue in the given project
func Template(project string, cve types.Vulnerabilities, packagetarget string, packagetype string) (types.RasicIssue, error) {
	projectId, _ := strconv.Atoi(project)
	var cveSeverity types.Severity
	cveSeverity = cve.Severity
	newIssue := types.RasicIssue{Title: cve.VulnerabilityID, Description: generateMarkdown(cve, packagetarget, packagetype), Id: projectId, Severity: cveSeverity}

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
