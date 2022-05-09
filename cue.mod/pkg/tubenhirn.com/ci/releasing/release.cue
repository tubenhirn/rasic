package releasing

import (
	"dagger.io/dagger"
	"universe.dagger.io/docker"
)

// create a release using semenatic-release
#Release: {
	authToken:  dagger.#Secret
	sourcecode: dagger.#FS

	_image: docker.#Pull & {
		source: "tubenhirn/semantic-release-gitlab:v2.0.0@sha256:11ecd980b986f3e78f9a6cbe72dd9466bb6a5bf14eca471c15472f96553d5c98"
	}

	docker.#Run & {
		input: _image.output
		mounts: code: {
			dest:     "/src"
			contents: sourcecode
		}
		workdir: "/src"
		env: {
			GITLAB_TOKEN: authToken
		}
	}
}
