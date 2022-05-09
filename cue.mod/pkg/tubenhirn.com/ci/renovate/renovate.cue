package renovate

import (
	"dagger.io/dagger"
	"universe.dagger.io/docker"
)

// create a release using semenatic-release
#Run: {
	gitlabToken: dagger.#Secret
	githubToken: dagger.#Secret
	project:     string
	platform:    *"gitlab" | string

	_image: docker.#Pull & {
		source: "renovate/renovate"
	}

	docker.#Run & {
		input: _image.output
		env: {
			RENOVATE_TOKEN:                gitlabToken
			GITHUB_COM_TOKEN:              githubToken
			RENOVATE_PLATFORM:             platform
			RENOVATE_EXTENDS:              "github>whitesource/merge-confidence:beta"
			RENOVATE_REQUIRE_CONFIG:       "true"
			RENOVATE_GIT_AUTHOR:           "Renovate Bot <bot@renovateapp.com>"
			RENOVATE_PIN_DIGEST:           "true"
			RENOVATE_DEPENDENCY_DASHBOARD: "false"
			RENOVATE_LABELS:               "renovate"
			RENOVATE_AUTODISCOVER:         "true"
			RENOVATE_AUTODISCOVER_FILTER:  project
			LOG_LEVEL: "debug"
		}
	}
}
