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
		source: "tubenhirn/semantic-release-gitlab"
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
