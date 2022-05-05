# [1.1.0](https://gitlab.com/jstang/rasic/compare/v1.0.0...v1.1.0) (2022-05-05)


### Features

* add ci pipeline with dagger ([c6af51a](https://gitlab.com/jstang/rasic/commit/c6af51ae16b241b70e27f156193337862eaffee0))

# 1.0.0 (2022-05-05)


### Bug Fixes

* add creation time to issue for gitlab.com ([6474d8b](https://gitlab.com/jstang/rasic/commit/6474d8bd52a3438293375b0ea280a86a3598be21))
* change os.exec to read .trivyignore file ([a9ad626](https://gitlab.com/jstang/rasic/commit/a9ad626b7a5dc94932927d8c3beb4250beaad174))
* change severity from HIGH to CRITICAL ([5df58ba](https://gitlab.com/jstang/rasic/commit/5df58ba732e5d3a72d9237bad31a478f446fd9a7))
* cleanup ([c2b30ad](https://gitlab.com/jstang/rasic/commit/c2b30adfc3dc8271f320db36650fdab1bcbab6b2))
* do not list archived projects ([d1e3016](https://gitlab.com/jstang/rasic/commit/d1e3016ccf13b752b03911530ab5df79be58431e))
* fix cve dublication check ([9cfa575](https://gitlab.com/jstang/rasic/commit/9cfa575399467771b6763f4e9a0d2f300f0c1565))
* fixed a bug where registryexcludes where not looked up properly ([abda001](https://gitlab.com/jstang/rasic/commit/abda0012dbbd0d9d1e3b1763bce7b91628173d1f))
* make plugin vars consts ([45c05a9](https://gitlab.com/jstang/rasic/commit/45c05a9225d442da538db05aabc1213bfbfaf7ca))
* make username not required ([151cd43](https://gitlab.com/jstang/rasic/commit/151cd43d295c1eb5f32c0f929ff27394a27ea6b1))
* pass auth to trivy ([f409b15](https://gitlab.com/jstang/rasic/commit/f409b150d29994e59d0e99790e7d821f5735be5f))
* refactor command functions with default values ([95386cf](https://gitlab.com/jstang/rasic/commit/95386cf19d0efd8a996e62b77088b0fda8b241e8))
* remove dublicate struct ([99d870f](https://gitlab.com/jstang/rasic/commit/99d870fc4ccc782df471351072228dc0da3504dc))
* remove logs ([ff474c6](https://gitlab.com/jstang/rasic/commit/ff474c6ac8835a27bdc189263ea235a98904c3c3))
* update run command for trivy scanner ([bb6f1de](https://gitlab.com/jstang/rasic/commit/bb6f1de84a98932ae2954661632d1f16734b30f9))
* use default values for main ([cbdb942](https://gitlab.com/jstang/rasic/commit/cbdb9420aca06303e205c0b1fb1939f6926c7691))


### Code Refactoring

* add plugin loader ([c36a6b3](https://gitlab.com/jstang/rasic/commit/c36a6b31498b76e124135a8f74e3343bca530565))
* move issue and scanner to core ([2dd2761](https://gitlab.com/jstang/rasic/commit/2dd2761a58b35545310865779d5bbd356348d8c2))
* move methods to matching packages ([41ddb58](https://gitlab.com/jstang/rasic/commit/41ddb58a89c48b1369b2cd6fdfe93bcb4bf114b1))
* move project list to subcommand section ([bc8be1d](https://gitlab.com/jstang/rasic/commit/bc8be1d5b97f5ffd54b0873c728c3907b4b71cf6))
* put baseUrl to var and format types🚆 ([f1eda28](https://gitlab.com/jstang/rasic/commit/f1eda2856c3da62e5afc09ea31eb8950f38885d6))
* remove unneeded methods from api and reporter plugins ([21c6bfa](https://gitlab.com/jstang/rasic/commit/21c6bfa92aea9777acb62b34c3768254ea509ae9))
* remove unused api methods ([e5d44bc](https://gitlab.com/jstang/rasic/commit/e5d44bc8d5b859e7b73d5cb23401b0794682d4b6))
* rename build command to compile ([bf0af6d](https://gitlab.com/jstang/rasic/commit/bf0af6d6764c800453c61f422069f40f5f2e2239))
* rename vars and add some todo's ([cb2ac32](https://gitlab.com/jstang/rasic/commit/cb2ac32353e84e4c9517048e883972acc1b19dad))
* rework cve creation ([d81965c](https://gitlab.com/jstang/rasic/commit/d81965c49934654dce84c65a88a2e7e802a58c66))
* sort imports🏹 ([5190585](https://gitlab.com/jstang/rasic/commit/51905856e422ef5cf52cb35b057ab73930ee3cc1))
* use gitlab types for gitlab reporter functions ([e1ad92a](https://gitlab.com/jstang/rasic/commit/e1ad92a6bf073cf4246ef7cb0feb60a3f8321da9))


### Continuous Integration

* add .trivyignore file ([94ca402](https://gitlab.com/jstang/rasic/commit/94ca402c707313320b471d48dc7702007c0b85e9))
* add cgo_enabled env flag ([27a84a9](https://gitlab.com/jstang/rasic/commit/27a84a9dc6ef56c5409553484187fb7cb68a4dc5))
* add ci helpers and fix linting ([3406689](https://gitlab.com/jstang/rasic/commit/340668921503b6e8792ae25989520462797fb250))
* add dagger ci build step ([c9602f1](https://gitlab.com/jstang/rasic/commit/c9602f11cf2f52bd863ef72c82f588667dd3a890))
* create a testfile ([3fede03](https://gitlab.com/jstang/rasic/commit/3fede034432dd5bf7a96eb458d63875efcbb40f2))
* update makefile install step to use ./bin dir ([a6ebd0e](https://gitlab.com/jstang/rasic/commit/a6ebd0e3572496d43eb39ae05f79eee63fbe8c0e))


### Documentation

* some doc strings ([cad17c6](https://gitlab.com/jstang/rasic/commit/cad17c627d63b8b70add5fc3481b2ea1ef72c4fd))
* update readme ([d874d32](https://gitlab.com/jstang/rasic/commit/d874d32e293b61d2bca441824206238c2b988f40))
* update README.md with build command ([00f2b0f](https://gitlab.com/jstang/rasic/commit/00f2b0f16c814f8141f9d312c561acc3d4383dc6))
* update README.md👓 ([6090672](https://gitlab.com/jstang/rasic/commit/60906723411bb21179416fd8aaa6b4f32b9a8df1))
* update README💣 ([4b1eca7](https://gitlab.com/jstang/rasic/commit/4b1eca7261e38c64dc55a293f08ab0f6f8b6af7c))


### Features

* add .trivyignore to be read when scanning a repo ([cf48f5a](https://gitlab.com/jstang/rasic/commit/cf48f5a93d6657a82b9c022e2175b53d4c1c68f0))
* add detection for open issues ([4858938](https://gitlab.com/jstang/rasic/commit/485893870f3c28105f05f43610d34b3b40424176))
* add first steps for issue command🌛 ([abfdde9](https://gitlab.com/jstang/rasic/commit/abfdde98ebd809f7ca1ca14bdb791bafafa125fb))
* add found cve count output to console ([3f7ae55](https://gitlab.com/jstang/rasic/commit/3f7ae55d6b524326c5f58b2f628888a9ff6795aa))
* add gitignore ([d76458f](https://gitlab.com/jstang/rasic/commit/d76458f409373dd47914e24465905d2761263944))
* add gitlab api ([549b14b](https://gitlab.com/jstang/rasic/commit/549b14b27fb1e539e103cae11a7b954aff1930b3))
* add issue labels to reporter plugins ([b29cce2](https://gitlab.com/jstang/rasic/commit/b29cce2e7c2ef69929732c2b89d4c0344787fd52))
* add issue labels to reporter plugins ([7a95c6c](https://gitlab.com/jstang/rasic/commit/7a95c6cfbd376d9b0f49c5c3ecc9b33be042a4f8))
* add plugin system ([4d62cf7](https://gitlab.com/jstang/rasic/commit/4d62cf7260d5330f7e91b040ff6f3de60539296b))
* add pluginhome flag ([cc67931](https://gitlab.com/jstang/rasic/commit/cc6793189f43300893dd16f9468256dce3b95607))
* add pluginhome var ([f2d0068](https://gitlab.com/jstang/rasic/commit/f2d00688532c35bbd99dd12416b081ee09daa24d))
* add registryExcludeFlag ([9d5d821](https://gitlab.com/jstang/rasic/commit/9d5d82104297ef6495c75a473761719a2f350ed0))
* add repository scanner ([e925ff0](https://gitlab.com/jstang/rasic/commit/e925ff0db428ac52943c2e74ebc75bb71cc7376d))
* add scan command ([6755452](https://gitlab.com/jstang/rasic/commit/6755452d4e2183d1c1b382df06a327b7b01a1b29))
* add severity flag ([a8ad1f3](https://gitlab.com/jstang/rasic/commit/a8ad1f34bb9d6cdfbd23d3331334c8a261b1f903))
* add single project support ([392e5ec](https://gitlab.com/jstang/rasic/commit/392e5eca898ed27b5527b37352a2cea39471ab10))
* create issues on gitlab ([f9f56f7](https://gitlab.com/jstang/rasic/commit/f9f56f753e14b09a324c5a863a13125361488555))
* declare core package ([1c9cb0e](https://gitlab.com/jstang/rasic/commit/1c9cb0ecd9dd08b9c08cff55f18c484289657530))
* make it possible to ignore a local .trivyignorefile ([fb8f7b3](https://gitlab.com/jstang/rasic/commit/fb8f7b3372341950d3a734070e43796e9c835492))
* make resultFilePath a variable ([a6d66bf](https://gitlab.com/jstang/rasic/commit/a6d66bf16f746176073e8b66309d9cd1038995ea))
* make severity compareabel ([4f88a6f](https://gitlab.com/jstang/rasic/commit/4f88a6f7139407b6f0a2d9e1fb885eb0e5868d93))
* refactor issue.Open method ([b636fed](https://gitlab.com/jstang/rasic/commit/b636fed74a4919a1601e5435303ffaedda212857))
* remove third for loop for cve checking ([72c8ac8](https://gitlab.com/jstang/rasic/commit/72c8ac821d550dc0ebead4b9a892e58a267e9d4c))
* rework logging with pterm ([056d946](https://gitlab.com/jstang/rasic/commit/056d946e621e5c28e4058d44462f419a9477d113))
* split plugins by type ([3ac411b](https://gitlab.com/jstang/rasic/commit/3ac411b6985b1708554e12c2caaba42c078ea509))
* support third party container registries (e.g. gcr) ([d647333](https://gitlab.com/jstang/rasic/commit/d6473335625993c0c58c47304249251eeedd8889))
* use rasic types instead of gitlab ones as default ([a242865](https://gitlab.com/jstang/rasic/commit/a242865fe37ebb8e860ac3a539899c607034f831))


### Miscellaneous Chores

* organize code ([ca1c5d5](https://gitlab.com/jstang/rasic/commit/ca1c5d5bb2235eef453378892c4bb964a0f8a35e))
* update dagger ([7e17de9](https://gitlab.com/jstang/rasic/commit/7e17de9173b1ffc0a251341fc5f5a0aff99d83b0))
* update deps ([5519391](https://gitlab.com/jstang/rasic/commit/5519391fa851730b8aa49de599465c7052e5e5ec))
* update deps ([0a37ee8](https://gitlab.com/jstang/rasic/commit/0a37ee8e4783e76429ea0aa54692b3018505dacf))
* update deps ([0d3aa25](https://gitlab.com/jstang/rasic/commit/0d3aa25677970e368c6dab7ba0c9ea150c74eb0e))
* update yamlv2 to v2.4.0 ([8dd0d2b](https://gitlab.com/jstang/rasic/commit/8dd0d2b5582b0486c2bec3e1c47de7e9185e3cda))