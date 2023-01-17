# rasic

rapid security incident creation.

a simple cli for cve scanning with trivy and
issue creation on gitlab.com (and more to come).

## requirements

the cli uses trivy as cve-scanner.

you need to have **trivy installed** on your system to use this cli.

## build

### build with dagger

``` bash
dagger-cue do build
```

## install

### homebrew

``` sh
brew tap tubenhirn/homebrew-formulae
brew install rasic
```

## use

### scan a project

simple project or group scan.\
rasic will find all projects inside a group.

``` sh
rasic scan \
    --project <PROJECT_OR_GROUP_ID> \
    --token <A_OAUTH_TOKEN>
```

### scan a project or group including its container images

include container images.

``` sh
rasic scan \
    --container \
    --project <PROJECT_OR_GROUP_ID> \
    --token <A_OAUTH_TOKEN>
```

## release

``` sh
dagger-cue do release
```
