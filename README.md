# rasic

rapid security incident creation

a simple cli for cve scanning with trivy and issue creation on gitlab.com (and more to come).

## requirements

the cli uses trivy as cve-scanner.

you need to have **trivy installed** on your system to use this cli.

## test

```sh
make test
```

## build

```sh
make build
```

## install

```sh
make install
```

## TODO

- add jira issue creation
- add a dry-run mode
