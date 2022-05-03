package ci

import (
	"dagger.io/dagger"
	"universe.dagger.io/go"

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
			}
			"api": go.#Build & {
				source:  _source
				package: "./plugins/api/gitlab.go"
			}
			"reporter": go.#Build & {
				source:  _source
				package: "./plugins/reporter/gitlab.go"
			}

		}
	}
}
