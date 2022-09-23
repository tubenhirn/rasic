package ci

import (
	"dagger.io/dagger"
	"dagger.io/dagger/core"

	"universe.dagger.io/alpha/go/golangci"
	"universe.dagger.io/alpha/go/goreleaser"
	"github.com/tubenhirn/dagger-ci-modules/releasing"
)

dagger.#Plan & {
	client: filesystem: ".": read: contents:       dagger.#FS
	client: filesystem: "./dist": write: contents: actions.build.export.directories."/src/dist"

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

		release: {
			semanticRelease: releasing.#Release & {
				sourcecode: _source
				authToken:  client.env.GITLAB_TOKEN
				version:    "v2.5.0"
			}
		}

		releaseArtifacts: goreleaser.#Release & {
			source:     _source
			removeDist: true
			env: {
				"APP_VERSION":  _version.contents
				"GITLAB_TOKEN": client.env.GITLAB_TOKEN
			}
		}

	}
}
