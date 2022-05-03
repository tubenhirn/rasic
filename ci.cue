package ci

import (
	"dagger.io/dagger"
	"universe.dagger.io/go"

	"rasic.io/ci/golangci"
)

dagger.#Plan & {
	client: filesystem: ".": read: contents:                       dagger.#FS
	client: filesystem: "./bin": write: contents:                  actions.build."rasic".output
	client: filesystem: "./bin/plugins/api": write: contents:      actions.build."api".output
	client: filesystem: "./bin/plugins/reporter": write: contents: actions.build."reporter".output

	actions: {
		_source: client.filesystem["."].read.contents
		build: {
			"rasic": go.#Build & {
				source: _source
				os:     client.platform.os
				arch:   client.platform.arch
				env: {
					CGO_ENABLED: "0"
				}
			}
			"api": go.#Build & {
				source:  _source
				package: "./plugins/api/gitlab.go"
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
				version: "1.45"
			}

		}
	}
}
