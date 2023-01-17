package ci

import (
	"dagger.io/dagger"
	"dagger.io/dagger/core"

	"universe.dagger.io/alpha/go/golangci"
	"github.com/tubenhirn/dagger-ci-modules/goreleaser"
)

dagger.#Plan & {
	client: filesystem: ".": read: contents:       dagger.#FS
	client: filesystem: "./dist": write: contents: actions.build.export.directories."/src/dist"

	client: env: {
		GITHUB_TOKEN: dagger.#Secret
	}

	actions: {
		_source:  client.filesystem["."].read.contents
		_version: core.#ReadFile & {
			input: _source
			path:  "version"
		}
		build: goreleaser.#Release & {
			source:     _source
			snapshot:   true
			removeDist: true
			env: {
				"APP_VERSION": _version.contents
			}
		}

		lint: {
			go: golangci.#Lint & {
				source:  _source
				version: "1.48"
			}
		}

		semanticRelease: {
			semanticRelease: releasing.#Release & {
				sourcecode: _source
				platform:   "git"
				authToken:  client.env.GITHUB_TOKEN
				version:    "v2.9.0"
			}
		}

		release: goreleaser.#Release & {
			source:     _source
			removeDist: true
			env: {
				"APP_VERSION":  _version.contents
				"GITHUB_TOKEN": client.env.GITHUB_TOKEN
			}
		}

	}
}
