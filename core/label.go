package core

import (
	"strconv"

	"gitlab.com/jstang/rasic/types"
	"gitlab.com/jstang/rasic/types/plugins"
	"golang.org/x/exp/slices"
)

// check project for labels
// if they dont exist, create them
// labels are severities here. we create them with a name and a color
// labels are added to issues (e.g. CRITICAL, HIGHT, MEDIUM....)
func CheckLabels(httpClient types.HttpClient, reporterPlugin plugins.Reporter, project types.RasicProject, authToken string) {

	// get all labels for current project
	var projectLabels []types.RasicLabel
	projectLabels = reporterPlugin.GetLabels(httpClient, strconv.Itoa(project.Id), authToken)

	var labelSlice []string
	// check if all required labels do exist
	// first put them in a list
	for _, label := range projectLabels {
		labelSlice = append(labelSlice, label.Name)
	}

	requiredLabels := []string{types.Critical.String(), types.High.String(), types.Medium.String(), types.Low.String(), types.Unknown.String()}

	for _, required := range requiredLabels {
		if !slices.Contains(labelSlice, required) {
			// create the missing label from the name string
			// cast it into our severity to add the color
			var severity types.Severity
			severity = severity.FromString(required)

			var newLabel types.RasicLabel
			newLabel.Name = required
			newLabel.Color = severity.Color()

			reporterPlugin.CreateLabel(httpClient, strconv.Itoa(project.Id), authToken, newLabel)
		}
	}
}
