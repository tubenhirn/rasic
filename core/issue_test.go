package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/jstang/rasic/types"
)

func TestNewline(t *testing.T) {
	test := newline("test")
	assert.Equal(t, "test\n", test, "newline() failed")
}

func TestDoBreak(t *testing.T) {
	test := dobreak("test")
	assert.Equal(t, "test<br>", test, "dobreak() failed")
}

func TestBold(t *testing.T) {
	test := bold("test")
	assert.Equal(t, "**test**", test, "bold() failed")
	assert.NotEqual(t, "test", test, "bold() failed")
}

func TestConsoleStart(t *testing.T) {
	test := consoleStart()
	assert.Equal(t, "```bash\n", test, "consoleStart() failed")
}

func TestGenerateMarkdown(t *testing.T) {
	test := generateMarkdown(testVulnerability, "testTarget", "testType")
	await := "### testTitle\ntestDescription<br>\ntestPrimaryURL<br>\n### Severity\n**LOW**<br>\n### Package-Information\n```bash\ntarget=testTarget\ntype=testType\npackagename=testPkgName\ninstalled_version=testInstalledVersion\nfixed_version=testFixedVersion\n\n```"
	assert.Equal(t, await, test, "generateMarkdown() failed")
}

func TestTemplate(t *testing.T) {
	test, _ := Template("1", testVulnerability, "testTarget", "testType")
	assert.Equal(t, "testTitle", test.Title, "Template() failed")
}

var testVulnerability = types.Vulnerabilities{
	Title:            "testTitle",
	VulnerabilityID:  "testTitle",
	Description:      "testDescription",
	PrimaryURL:       "testPrimaryURL",
	Severity:         types.Severity(1),
	PkgName:          "testPkgName",
	InstalledVersion: "testInstalledVersion",
	FixedVersion:     "testFixedVersion",
}
