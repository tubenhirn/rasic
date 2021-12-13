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
	url := "https://gitlab.com/api/v4/groups/" + group + "/projects?per_page=100&include_subgroups=true"
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
