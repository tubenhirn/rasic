package releasing

import (
	"dagger.io/dagger"

	"universe.dagger.io/docker"
)

// create a release using semenatic-release
#Release: {
	authToken: dagger.#Secret
	sourcecode: dagger.#FS

	_image: docker.#Pull & {
		source: "tubenhirn/semantic-release-gitlab:v1.0.0@sha256:44f6a3efd1cb0bd32ef2b8397cb2c67633855c52796d168a3fd01d18058b22e1"
	}

	docker.#Run & {
		input:    _image.output
		mounts: source: {
			dest: "/src"
			contents: sourcecode
		}
		workdir: "/src"
		env: {
			GITLAB_TOKEN: authToken
		}
	}
}
