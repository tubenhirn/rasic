# rasic

rapid security incident creation

a simple cli for cve scanning with trivy and issue creation on gitlab.com (and more to come).

## requirements

the cli uses trivy as cve-scanner.\
you need to have trivy installed on your system to use this cli.

## build

```sh
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a
```

## TODO

- add jira issue creation
- add a dry-run mode
