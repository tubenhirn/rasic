package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/urfave/cli/v2"
	"tubenhirn.com/cve2issue/types"
)

func GetProjectList(group string, token string) (types.Projects, error) {
	url := "https://gitlab.com/api/v4/groups/" + group + "/projects?per_page=100&include_subgroups=true&archived=false"
	client := http.Client{}
	req, reqerr := http.NewRequest("GET", url, nil)
	if reqerr != nil {
		log.Fatalln(reqerr)
		return nil, cli.NewExitError("request error", 1)
	}

	req.Header.Set("PRIVATE-TOKEN", token)

	res, reserr := client.Do(req)
	if reserr != nil {
		log.Fatalln(reserr)
		return nil, cli.NewExitError("request error", 1)
	}

	defer res.Body.Close()

	if res.Status == "200 OK" || res.Status == "200" {
		var projectlist types.Projects
		if err := json.NewDecoder(res.Body).Decode(&projectlist); err != nil {
			return nil, cli.NewExitError("decoder error", 2)
		}
		return projectlist, nil
	} else {
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return nil, cli.NewExitError("read error", 3)
		}
		return nil, cli.NewExitError(string(res.Status), 2)
	}

}

func GetProject(project string, token string) (*types.Project, error) {
	url := "https://gitlab.com/api/v4/projects/" + project
	client := http.Client{}
	req, reqerr := http.NewRequest("GET", url, nil)
	if reqerr != nil {
		log.Fatalln(reqerr)
		return nil, cli.NewExitError("request error", 1)
	}

	req.Header.Set("PRIVATE-TOKEN", token)

	res, reserr := client.Do(req)
	if reserr != nil {
		log.Fatalln(reserr)
		return nil, cli.NewExitError("request error", 1)
	}

	defer res.Body.Close()

	if res.Status == "200 OK" || res.Status == "200" {
		project := &types.Project{}
		if err := json.NewDecoder(res.Body).Decode(&project); err != nil {
			return nil, cli.NewExitError("decoder error", 2)
		}
		return project, nil
	} else {
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return nil, cli.NewExitError("read error", 3)
		}
		return nil, cli.NewExitError(string(res.Status), 2)
	}

}

func GetIssueList(project string, token string) (types.Issues, error) {
	url := "https://gitlab.com/api/v4/projects/" + project + "/issues?per_page=100"
	client := http.Client{}
	req, reqerr := http.NewRequest("GET", url, nil)
	if reqerr != nil {
		log.Fatalln(reqerr)
		return nil, cli.NewExitError("request error", 1)
	}

	req.Header.Set("PRIVATE-TOKEN", token)

	res, reserr := client.Do(req)
	if reserr != nil {
		log.Fatalln(reserr)
		return nil, cli.NewExitError("request error", 1)
	}

	defer res.Body.Close()

	if res.Status == "200 OK" || res.Status == "200" {
		var issuelist types.Issues
		if err := json.NewDecoder(res.Body).Decode(&issuelist); err != nil {
			return nil, cli.NewExitError("decoder error", 2)
		}
		return issuelist, nil
	} else {
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return nil, cli.NewExitError("read error", 3)
		}
		return nil, cli.NewExitError(string(res.Status), 2)
	}

}
