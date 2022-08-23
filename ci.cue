package ci

import (
	"dagger.io/dagger"
	"dagger.io/dagger/core"
	"universe.dagger.io/go"

	"universe.dagger.io/alpha/go/golangci"
	"tubenhirn.com/ci/releasing"
	"tubenhirn.com/ci/renovate"
)

dagger.#Plan & {
	client: filesystem: ".": read: contents:                       dagger.#FS
	client: filesystem: "./bin": write: contents:                  actions.build."rasic".output
	client: filesystem: "./bin/plugins/source": write: contents:   actions.build."source".output
	client: filesystem: "./bin/plugins/reporter": write: contents: actions.build."reporter".output
	client: env: {
		GITLAB_TOKEN: dagger.#Secret
		GITHUB_TOKEN: dagger.#Secret
	}

	actions: {
		_source:  client.filesystem["."].read.contents
		_version: core.#ReadFile & {
			input: _source
			path:  "version"
		}
		build: {
			"rasic": go.#Build & {
				source:  _source
				os:      client.platform.os
				arch:    client.platform.arch
				ldflags: "-X main.appVersion=\(_version.contents)"
				env: {
					CGO_ENABLED: "0"
				}
			}
			"source": go.#Build & {
				source:  _source
				package: "./plugins/source/gitlab.go"
				os:      client.platform.os
				arch:    client.platform.arch
				env: {
					CGO_ENABLED: "0"
				}

			}
			"reporter": go.#Build & {
				source:  _source
				package: "./plugins/reporter/gitlab.go"
				os:      client.platform.os
				arch:    client.platform.arch
				env: {
					CGO_ENABLED: "0"
				}

			}

		}
		lint: {
			go: golangci.#Lint & {
				source:  _source
				version: "1.48"
			}
		}
		release: releasing.#Release & {
			sourcecode: _source
			authToken:  client.env.GITLAB_TOKEN
			version:    "v2.4.3"
		}
		"renovate": renovate.#Run & {
			project:     "jstang/rasic"
			gitlabToken: client.env.GITLAB_TOKEN
			githubToken: client.env.GITHUB_TOKEN
		}
	}
}
