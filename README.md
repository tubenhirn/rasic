# cve2issue

a simple cli for sve scanning and issue creation on gitlab.com

## requirements

the cli uses trivy as cve-scanner.<br>
you need to have trivy installed on your system to use this cli.<br>

## build

```sh
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a
```

## TODO

- add jira issue creation
- add a dry-run mode
- ...
