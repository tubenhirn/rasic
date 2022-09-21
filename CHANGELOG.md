## [1.13.4](https://gitlab.com/jstang/rasic/compare/v1.13.3...v1.13.4) (2022-09-21)


### Bug Fixes

* fix job name for relase ([d3dc3e0](https://gitlab.com/jstang/rasic/commit/d3dc3e00c229343748e4ed3d1ef5e3bf14956dd9))
* fix jobs ([2ca4fcb](https://gitlab.com/jstang/rasic/commit/2ca4fcb32b3d69735345be76703054960d8324b1))
* fix release process and combine the two steps ([e1d4ac8](https://gitlab.com/jstang/rasic/commit/e1d4ac80c250fea80d6a073c06accd524c3a9e7e))

## [1.13.3](https://gitlab.com/jstang/rasic/compare/v1.13.2...v1.13.3) (2022-09-21)


### Bug Fixes

* fix releaseArtifacts job ([a17421c](https://gitlab.com/jstang/rasic/commit/a17421c4e84e4c3212b1079508e184e591df1bf6))

## [1.13.2](https://gitlab.com/jstang/rasic/compare/v1.13.1...v1.13.2) (2022-09-21)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 46d9e77 ([94dd40e](https://gitlab.com/jstang/rasic/commit/94dd40ee6bc393795dcf55ad3400f456cf514b6a))
* **deps:** update module github.com/hashicorp/go-hclog to v1.3.1 ([d07d113](https://gitlab.com/jstang/rasic/commit/d07d1132782a17101d74edb106ae09f0f4e42520))


### Code Refactoring

* **ci:** refactor build step and use goreleaser for snapshot builds ([5f0f0c9](https://gitlab.com/jstang/rasic/commit/5f0f0c94613a6853c118151e0ce7adb31d4d611f))


### Continuous Integration

* add packagerule for renovate ([316fbf2](https://gitlab.com/jstang/rasic/commit/316fbf226c321eaff7c010bd41cf032b163da551))
* update compile step ([e3d479e](https://gitlab.com/jstang/rasic/commit/e3d479eafd485d2c927bd93ff4e17e532988b6ea))
* use goreleaser with dagger for building ([07cb9e8](https://gitlab.com/jstang/rasic/commit/07cb9e8bd7ced6dc0dafac3a67e9194b6e50f335))

## [1.13.1](https://gitlab.com/jstang/rasic/compare/v1.13.0...v1.13.1) (2022-09-18)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 2d61f44 ([f146ebe](https://gitlab.com/jstang/rasic/commit/f146ebea2515a7047ddfb05c0885b42edc35b3c3))
* **deps:** update golang.org/x/exp digest to b168a2c ([805ff00](https://gitlab.com/jstang/rasic/commit/805ff00898913acca43a94c0c15cf83b9c6ea612))


### Miscellaneous Chores

* update dagger ([6f61c53](https://gitlab.com/jstang/rasic/commit/6f61c53aee9be348e870fecec5f417b326852f3a))

# [1.13.0](https://gitlab.com/jstang/rasic/compare/v1.12.0...v1.13.0) (2022-09-13)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 5c715a9 ([581f57a](https://gitlab.com/jstang/rasic/commit/581f57af7fa929c60506af092bb8f741784cb3ba))
* **deps:** update module github.com/urfave/cli/v2 to v2.16.3 ([f324da1](https://gitlab.com/jstang/rasic/commit/f324da1548c746a6e4bec9870c7f66c20a556bc0))


### Continuous Integration

* update release image ([e274f07](https://gitlab.com/jstang/rasic/commit/e274f07cac625fb069d815c39046202a13430a5a))


### Features

* **deps:** update module github.com/urfave/cli/v2 to v2.15.0 ([7218d08](https://gitlab.com/jstang/rasic/commit/7218d0852955e72df30d00299fc20c04228ce387))
* **deps:** update module github.com/urfave/cli/v2 to v2.16.0 ([6a7f264](https://gitlab.com/jstang/rasic/commit/6a7f2642422ddfa7a5f15ec3cd0f3b183136d1f8))


### Miscellaneous Chores

* add .envrc to ignore ([c0bdc58](https://gitlab.com/jstang/rasic/commit/c0bdc587dcfca6e1b27177b15c6c7179e1e17166))

# [1.12.0](https://gitlab.com/jstang/rasic/compare/v1.11.1...v1.12.0) (2022-09-07)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 145caa8 ([66327ca](https://gitlab.com/jstang/rasic/commit/66327caebb8bb13e8ca7a022c9cda4858eb4f327))


### Continuous Integration

* update dagger to v0.2.33 ([1b5bd8e](https://gitlab.com/jstang/rasic/commit/1b5bd8e99dd8451f5564e07ded3150d48c064611))


### Features

* **deps:** update module github.com/urfave/cli/v2 to v2.14.1 ([126f0a2](https://gitlab.com/jstang/rasic/commit/126f0a223a027561d3c34557dcac730f6e74089c))

## [1.11.1](https://gitlab.com/jstang/rasic/compare/v1.11.0...v1.11.1) (2022-09-05)


### Bug Fixes

* **deps:** update module github.com/pterm/pterm to v0.12.46 ([0f8c54c](https://gitlab.com/jstang/rasic/commit/0f8c54cdbe1a8c98b9849cfca7eb6534c97c4de4))


### Continuous Integration

* upload artifacts to gitlab.com package registry ([141afda](https://gitlab.com/jstang/rasic/commit/141afda6f9572cc8db8ebac3816bf01c6608cc55))

# [1.11.0](https://gitlab.com/jstang/rasic/compare/v1.10.3...v1.11.0) (2022-09-04)


### Features

* **deps:** update module github.com/hashicorp/go-hclog to v1.3.0 ([5636d96](https://gitlab.com/jstang/rasic/commit/5636d96d1402f65f01362958db4405044ae6aa8e))
* use goreleaser ([db32d32](https://gitlab.com/jstang/rasic/commit/db32d32c8ebc68ca91edd9f20fc7bd608e2189c1))

## [1.10.3](https://gitlab.com/jstang/rasic/compare/v1.10.2...v1.10.3) (2022-08-29)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 334a238 ([6edcfde](https://gitlab.com/jstang/rasic/commit/6edcfde5a765bfda7e4e7692bb98c4a776531a5c))
* **deps:** update golang.org/x/exp digest to bd9bcdd ([501bb53](https://gitlab.com/jstang/rasic/commit/501bb53599d52bcd7494e0491f602aff275505d9))


### Miscellaneous Chores

* update dagger ([53384bf](https://gitlab.com/jstang/rasic/commit/53384bfe4e790f53abeb1dfc736d7e5426c006a9))

## [1.10.2](https://gitlab.com/jstang/rasic/compare/v1.10.1...v1.10.2) (2022-08-23)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 807a232 ([aefc48d](https://gitlab.com/jstang/rasic/commit/aefc48dd827c30ad12769b951a19aae83b1c872e))
* **deps:** update module github.com/hashicorp/go-plugin to v1.4.5 ([e008adb](https://gitlab.com/jstang/rasic/commit/e008adbee6b565b1a8f55c06b870937ea78f4d9b))
* **deps:** update module github.com/urfave/cli/v2 to v2.11.2 ([8a14012](https://gitlab.com/jstang/rasic/commit/8a140124f852341a8091b75236c00398820ed5ad))


### Miscellaneous Chores

* update dagger ([97ca17f](https://gitlab.com/jstang/rasic/commit/97ca17fcf694abacf3f6e659e9e34983c2cee4df))
* update semantic-release version to v2.4.3 ([e180916](https://gitlab.com/jstang/rasic/commit/e18091697768e6d17815c5c516edfecd1c7ec689))

## [1.10.1](https://gitlab.com/jstang/rasic/compare/v1.10.0...v1.10.1) (2022-08-13)


### Bug Fixes

* cleanup golangci linting and do some linting ([5cef440](https://gitlab.com/jstang/rasic/commit/5cef440de134d3a16e69624443e9ea5b91688cf6))


### Continuous Integration

* update semantic-release to v2.4.2 ([353dc8c](https://gitlab.com/jstang/rasic/commit/353dc8c899b18fae1a66ccb45c29f09983e00739))


### Miscellaneous Chores

* more linting for go 1.19 ([7ebd00b](https://gitlab.com/jstang/rasic/commit/7ebd00bb23cece7f7eb5e3672c8bfa752de9b3a7))
* update dagger ([cd07c15](https://gitlab.com/jstang/rasic/commit/cd07c15b6aea6a66f3c3ae75963fc4d3ff9b5eef))
* update golang verion for golang-ci ([d49bb35](https://gitlab.com/jstang/rasic/commit/d49bb3543479706105528db3a728e519daeb790d))

# [1.10.0](https://gitlab.com/jstang/rasic/compare/v1.9.5...v1.10.0) (2022-08-03)


### Bug Fixes

* **deps:** update module github.com/hashicorp/go-hclog to v1.2.2 ([00a7506](https://gitlab.com/jstang/rasic/commit/00a7506878d71f6c23e560391946f5926c704c9c))
* **deps:** update module github.com/pterm/pterm to v0.12.45 ([4a985df](https://gitlab.com/jstang/rasic/commit/4a985df10834173b132a36890b3b27c4d717fd61))


### Features

* **deps:** update module go to 1.19 ([266d261](https://gitlab.com/jstang/rasic/commit/266d26101e72ba794bedb6db4278b8f33fba019b))


### Miscellaneous Chores

* update dagger ([181f802](https://gitlab.com/jstang/rasic/commit/181f8027d39711e8145a574edf8103da97073277))
* update dagger ([82a32b7](https://gitlab.com/jstang/rasic/commit/82a32b70e4d585225027c124f2423b2d9b5a07b3))

## [1.9.5](https://gitlab.com/jstang/rasic/compare/v1.9.4...v1.9.5) (2022-07-25)


### Bug Fixes

* fix linting errors and add a "positiv" scan output ([c22d9c0](https://gitlab.com/jstang/rasic/commit/c22d9c06557edd72400a35861036cadce053008c))

## [1.9.4](https://gitlab.com/jstang/rasic/compare/v1.9.3...v1.9.4) (2022-07-25)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to a9213ee ([1471d3a](https://gitlab.com/jstang/rasic/commit/1471d3a56f713c4086d134d009e484fec8c5378f))
* **deps:** update module github.com/pterm/pterm to v0.12.44 ([5467ec1](https://gitlab.com/jstang/rasic/commit/5467ec1363aec9ab4442aaf41611d0960b60418d))


### Code Refactoring

* change tmp directory creation ([ded5465](https://gitlab.com/jstang/rasic/commit/ded54650eeb8052890fd854b650b045fa4c7401e))

## [1.9.3](https://gitlab.com/jstang/rasic/compare/v1.9.2...v1.9.3) (2022-07-22)


### Bug Fixes

* **deps:** update module github.com/pterm/pterm to v0.12.43 ([076907c](https://gitlab.com/jstang/rasic/commit/076907c3c19f4678959c37b4c04de2e1e31a60d2))
* fix a bug where the tmp dir of a scanned project can not be deleted ([1ea3fce](https://gitlab.com/jstang/rasic/commit/1ea3fce06cb1ba50e60079d37a31aa959c4e6fa1))

## [1.9.2](https://gitlab.com/jstang/rasic/compare/v1.9.1...v1.9.2) (2022-07-21)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 79cabaa ([f42a0df](https://gitlab.com/jstang/rasic/commit/f42a0df2f569513be9e645bcc09d578ce11bfdfa))
* **deps:** update module github.com/urfave/cli/v2 to v2.11.1 ([8fb48da](https://gitlab.com/jstang/rasic/commit/8fb48da0ec870802aef699d789c8d83c49362185))


### Code Refactoring

* do some linting with golang-ci lint ([ae70c90](https://gitlab.com/jstang/rasic/commit/ae70c902ca66285a93be145600caf6c084e83a10))


### Continuous Integration

* update semantic-release version to v2.4.1 ([d7728d2](https://gitlab.com/jstang/rasic/commit/d7728d2b6a8abf17acbca7c761bc0354cf7d9139))


### Miscellaneous Chores

* update dagger ([4573889](https://gitlab.com/jstang/rasic/commit/457388969cc68fcab58cf4a2bfdefc25f0697e81))
* update dagger ([2bb8116](https://gitlab.com/jstang/rasic/commit/2bb8116c1e739f5d8fe8d609234f9af4d313b0ab))

## [1.9.1](https://gitlab.com/jstang/rasic/compare/v1.9.0...v1.9.1) (2022-07-14)


### Miscellaneous Chores

* update semantic-release version to v2.4.0 ([3280f43](https://gitlab.com/jstang/rasic/commit/3280f432c650b54f7628336f21f32ab7386ff67e))

# [1.9.0](https://gitlab.com/jstang/rasic/compare/v1.8.1...v1.9.0) (2022-07-14)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to b4a6d95 ([4ce9322](https://gitlab.com/jstang/rasic/commit/4ce9322e27aa61df3a6e7a25bc9d961df478d365))
* **deps:** update module github.com/urfave/cli/v2 to v2.10.3 ([80a10f7](https://gitlab.com/jstang/rasic/commit/80a10f7f90d33778cbd4dd4e0791308b3227c45d))


### Continuous Integration

* add version string config parameter to release job ([85958a0](https://gitlab.com/jstang/rasic/commit/85958a048319ed2670937bf1f011f39ca812aba9))


### Features

* **deps:** update module github.com/urfave/cli/v2 to v2.11.0 ([4d9f248](https://gitlab.com/jstang/rasic/commit/4d9f248c863a43a79417479debf7120ff9c086cb))


### Miscellaneous Chores

* update dagger ([4e54f26](https://gitlab.com/jstang/rasic/commit/4e54f2638922d4dcc694000905570f564fb784d2))
* update dagger ([b00218c](https://gitlab.com/jstang/rasic/commit/b00218cac15f9b749637e85bc068a5c0546014e5))

## [1.8.1](https://gitlab.com/jstang/rasic/compare/v1.8.0...v1.8.1) (2022-06-24)


### Bug Fixes

* unset GITHUB_TOKEN if GITLAB_TOKEN is set ([6286d5d](https://gitlab.com/jstang/rasic/commit/6286d5dc1be7b5aa95b579e4e217bebacf052f1b))


### Miscellaneous Chores

* update dagger ([9203f67](https://gitlab.com/jstang/rasic/commit/9203f6714c68847e6daa80ada031b7f64a23d623))

# [1.8.0](https://gitlab.com/jstang/rasic/compare/v1.7.2...v1.8.0) (2022-06-21)


### Bug Fixes

* **deps:** update module github.com/pterm/pterm to v0.12.42 ([da7a9e9](https://gitlab.com/jstang/rasic/commit/da7a9e9e1a6035bed60b869a800683ce4e80e83d))
* **deps:** update module github.com/urfave/cli/v2 to v2.10.2 ([e84bd50](https://gitlab.com/jstang/rasic/commit/e84bd5068040e672192edf1329ea0010742ea3fe))


### Features

* add issue close command ([2714b9b](https://gitlab.com/jstang/rasic/commit/2714b9bb0520aab64a56e12243af10c920094bbb))

## [1.7.2](https://gitlab.com/jstang/rasic/compare/v1.7.1...v1.7.2) (2022-06-20)


### Bug Fixes

* write result files to tmp directory ([a1c54cb](https://gitlab.com/jstang/rasic/commit/a1c54cb9235efa38b63f364d7432d51eb26e57e1))

## [1.7.1](https://gitlab.com/jstang/rasic/compare/v1.7.0...v1.7.1) (2022-06-19)


### Bug Fixes

* add label "cve" to new issues ([24b0669](https://gitlab.com/jstang/rasic/commit/24b06690755b352944c7f88de9950b79abf6dd7f))

# [1.7.0](https://gitlab.com/jstang/rasic/compare/v1.6.1...v1.7.0) (2022-06-19)


### Documentation

* update README.md with some usage examples ([fa37ce7](https://gitlab.com/jstang/rasic/commit/fa37ce713264c84b62ad7052d393873727e87a71))


### Features

* **deps:** update module github.com/urfave/cli/v2 to v2.10.0 ([c06a49b](https://gitlab.com/jstang/rasic/commit/c06a49bdff838202526cb5756237dcbb1b15b0a8))

## [1.6.1](https://gitlab.com/jstang/rasic/compare/v1.6.0...v1.6.1) (2022-06-15)


### Miscellaneous Chores

* update copyright year ([266cf01](https://gitlab.com/jstang/rasic/commit/266cf018c595e99b23f4f0e11e8e6ad21a1add16))

# [1.6.0](https://gitlab.com/jstang/rasic/compare/v1.5.7...v1.6.0) (2022-06-15)


### Features

* get issues with pagination ([f75faeb](https://gitlab.com/jstang/rasic/commit/f75faeb6d113262c43532fe2a6b7bc78e0b381ba))

## [1.5.7](https://gitlab.com/jstang/rasic/compare/v1.5.6...v1.5.7) (2022-06-14)


### Bug Fixes

* fix getIssues method ([5e98c9d](https://gitlab.com/jstang/rasic/commit/5e98c9de55db2ef069b114b96a7ac99acbc20bca))

## [1.5.6](https://gitlab.com/jstang/rasic/compare/v1.5.5...v1.5.6) (2022-06-14)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to b0d7811 ([16e7475](https://gitlab.com/jstang/rasic/commit/16e747517e9290c77517c829e69d1161ff6e2627))

## [1.5.5](https://gitlab.com/jstang/rasic/compare/v1.5.4...v1.5.5) (2022-06-10)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to a51bd04 ([0d5d581](https://gitlab.com/jstang/rasic/commit/0d5d581bd7b3c97d32be98e306236e20b2c84ad2))


### Miscellaneous Chores

* update dagger ([650a37d](https://gitlab.com/jstang/rasic/commit/650a37d3c7703df6be85c1f3c74440bc30e21687))
* update dagger ([6d5db81](https://gitlab.com/jstang/rasic/commit/6d5db8117c38982a56448184ff8d9844fd549efc))

## [1.5.4](https://gitlab.com/jstang/rasic/compare/v1.5.3...v1.5.4) (2022-06-07)


### Bug Fixes

* **deps:** update module github.com/hashicorp/go-hclog to v1.2.1 ([e428549](https://gitlab.com/jstang/rasic/commit/e42854916970b05321a96c10457baf762393a6c4))


### Miscellaneous Chores

* update dagger ([d0871e4](https://gitlab.com/jstang/rasic/commit/d0871e4042c32d1be5c1af5abc87deed5d483695))

## [1.5.3](https://gitlab.com/jstang/rasic/compare/v1.5.2...v1.5.3) (2022-06-06)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 4a0574d ([302249c](https://gitlab.com/jstang/rasic/commit/302249c35c99a781e8636cfe6b369e1356f32c1f))
* **deps:** update golang.org/x/net digest to c960675 ([4398abc](https://gitlab.com/jstang/rasic/commit/4398abcddd6943d6e04ebbbec45dd1aaef7d7752))
* **deps:** update google.golang.org/genproto digest to e326c6e ([480001d](https://gitlab.com/jstang/rasic/commit/480001db5b76684604eea8749b6c769257de1c9b))
* go mod tidy ([63eeef1](https://gitlab.com/jstang/rasic/commit/63eeef1fa3cf53362d14f51873a23b1e6ddbfb5f))
* move dep to indirect block ([d47df7b](https://gitlab.com/jstang/rasic/commit/d47df7b908289a82db742c993deb053c07a37073))

## [1.5.2](https://gitlab.com/jstang/rasic/compare/v1.5.1...v1.5.2) (2022-06-03)


### Miscellaneous Chores

* update dagger ([64eb003](https://gitlab.com/jstang/rasic/commit/64eb00385b48e1fa7b1bbb36cfcd56158fba2957))

## [1.5.1](https://gitlab.com/jstang/rasic/compare/v1.5.0...v1.5.1) (2022-05-30)


### Bug Fixes

* **deps:** update golang.org/x/net digest to 5463443 ([edacc8e](https://gitlab.com/jstang/rasic/commit/edacc8e8db592caa6aa64e3a948d0afea63a1109))
* **deps:** update golang.org/x/term digest to 065cf7b ([3886f72](https://gitlab.com/jstang/rasic/commit/3886f72dc66a7a6d54a7dae3c3780691970ee170))
* **deps:** update google.golang.org/genproto digest to 00d5c0f ([e265136](https://gitlab.com/jstang/rasic/commit/e265136a35ed5053a7a58759b8516e97e3bfa8dc))
* **deps:** update module github.com/urfave/cli/v2 to v2.8.1 ([5ceac29](https://gitlab.com/jstang/rasic/commit/5ceac297f5bc9bae4ece9431a844a6de04433672))


### Code Refactoring

* remove yaml dep ([f085793](https://gitlab.com/jstang/rasic/commit/f0857938817714e4e9c82594ab2578094d76dd1e))


### Miscellaneous Chores

* update dagger ([ba87c18](https://gitlab.com/jstang/rasic/commit/ba87c18fd6318ba5cecd20bee8b5664f18a61bc4))
* update dagger files ([9ec9199](https://gitlab.com/jstang/rasic/commit/9ec9199d999bce149611945e84fea11952613e7b))

# [1.5.0](https://gitlab.com/jstang/rasic/compare/v1.4.3...v1.5.0) (2022-05-23)


### Bug Fixes

* **deps:** update google.golang.org/genproto digest to 3a47de7 ([8ba58e1](https://gitlab.com/jstang/rasic/commit/8ba58e12008ea049ac2c70f79fc4f170746d05b2))


### Features

* **deps:** update module github.com/urfave/cli/v2 to v2.8.0 ([29baf87](https://gitlab.com/jstang/rasic/commit/29baf87c68bdebb6dbfd770f6c0242a778a0c8e6))

## [1.4.3](https://gitlab.com/jstang/rasic/compare/v1.4.2...v1.4.3) (2022-05-22)


### Bug Fixes

* **deps:** update golang.org/x/exp digest to 0b5c67f ([c6d499a](https://gitlab.com/jstang/rasic/commit/c6d499a41d41b09570a98c368bc53efc3a78544e))
* **deps:** update golang.org/x/exp digest to 24438e5 ([c36090e](https://gitlab.com/jstang/rasic/commit/c36090ee180bf05ee45a3975e0700b0fd066be64))
* **deps:** update golang.org/x/net digest to 2e3eb7b ([a04658f](https://gitlab.com/jstang/rasic/commit/a04658f36de671da7c972544e0ea5df3d5867727))
* **deps:** update golang.org/x/sys digest to 5e4e11f ([103e1bc](https://gitlab.com/jstang/rasic/commit/103e1bc3897a87aa28ee9a852dd25b5c33ef5f04))
* **deps:** update golang.org/x/sys digest to bc2c85a ([7341d31](https://gitlab.com/jstang/rasic/commit/7341d311996dbfe1e68dd80bffda6cf43c077924))

## [1.4.2](https://gitlab.com/jstang/rasic/compare/v1.4.1...v1.4.2) (2022-05-15)


### Bug Fixes

* **deps:** update golang.org/x/net digest to 9564170 ([ab8560d](https://gitlab.com/jstang/rasic/commit/ab8560d071c6fa7bda44eb025a259ebb435dbcf6))
* **deps:** update golang.org/x/sys digest to 45d2b45 ([6773308](https://gitlab.com/jstang/rasic/commit/67733086f291ccd2ba3661e001390b8428c40d12))
* **deps:** update module google.golang.org/grpc to v1.46.2 ([5f8c1a1](https://gitlab.com/jstang/rasic/commit/5f8c1a109fa3d1931d74ab5b19e4451b225754ab))


### Miscellaneous Chores

* update dagger ([4da834b](https://gitlab.com/jstang/rasic/commit/4da834b36ecc76edf225fd5a42aef06fb0ec1a79))

## [1.4.1](https://gitlab.com/jstang/rasic/compare/v1.4.0...v1.4.1) (2022-05-10)


### Miscellaneous Chores

* update semantic-release to v2.1.0 ([061badd](https://gitlab.com/jstang/rasic/commit/061badd0f28f53f872a48340b1425e73d69f5582))

# [1.4.0](https://gitlab.com/jstang/rasic/compare/v1.3.8...v1.4.0) (2022-05-10)


### Features

* **deps:** update dependency bats to v1.6.0 ([f29cfd4](https://gitlab.com/jstang/rasic/commit/f29cfd455d55abaffc3268c7a3d90ada77c114e8))

## [1.3.8](https://gitlab.com/jstang/rasic/compare/v1.3.7...v1.3.8) (2022-05-09)

## [1.3.7](https://gitlab.com/jstang/rasic/compare/v1.3.6...v1.3.7) (2022-05-09)

## [1.3.6](https://gitlab.com/jstang/rasic/compare/v1.3.5...v1.3.6) (2022-05-08)


### Code Refactoring

* apply linting rules ([a8f8c24](https://gitlab.com/jstang/rasic/commit/a8f8c2466ea8ab669a8b3b741ac018aa5366cfc2))


### Continuous Integration

* edit lint config ([fafc8ed](https://gitlab.com/jstang/rasic/commit/fafc8ed845c4411492a466e0eb898362847383f3))
* update golint config ([dd5993b](https://gitlab.com/jstang/rasic/commit/dd5993b89d9fc23a3573791e69881fb989b91725))

## [1.3.5](https://gitlab.com/jstang/rasic/compare/v1.3.4...v1.3.5) (2022-05-08)


### Bug Fixes

* dont pass username if empty ([37f59f0](https://gitlab.com/jstang/rasic/commit/37f59f08bff86d490a51fc7bdc8c8ead99c472ae))

## [1.3.4](https://gitlab.com/jstang/rasic/compare/v1.3.3...v1.3.4) (2022-05-07)


### Code Refactoring

* rename plugin api to source ([1e17e58](https://gitlab.com/jstang/rasic/commit/1e17e585b65e304cf0fbbcbfcbf2d1f2a0ab82fa))


### Documentation

* add ldflag comment for version ([318389d](https://gitlab.com/jstang/rasic/commit/318389d85c7ce93f7331c904bc44b689ae7a71e8))

## [1.3.3](https://gitlab.com/jstang/rasic/compare/v1.3.2...v1.3.3) (2022-05-07)


### Miscellaneous Chores

* upgrade deps ([c6f11fa](https://gitlab.com/jstang/rasic/commit/c6f11faeea1d97be3211419de1f88ad974a08d8c))

## [1.3.2](https://gitlab.com/jstang/rasic/compare/v1.3.1...v1.3.2) (2022-05-06)


### Miscellaneous Chores

* add a full release.config.js file instead of extending the base ([bacfcaf](https://gitlab.com/jstang/rasic/commit/bacfcaffc271c5068b855db18463aebae3bf958c))

# [1.3.0](https://gitlab.com/jstang/rasic/compare/v1.2.6...v1.3.0) (2022-05-05)


### Features

* add verion flag for go compiler ([6f89e9b](https://gitlab.com/jstang/rasic/commit/6f89e9b11dc65f2203827eb617ceceeb1df41060))

## [1.2.6](https://gitlab.com/jstang/rasic/compare/v1.2.5...v1.2.6) (2022-05-05)


### Bug Fixes

* add gitlab release creation ([3a08ab5](https://gitlab.com/jstang/rasic/commit/3a08ab523d604c1118e3801e1e18603bc5c2fc2a))

## [1.2.5](https://gitlab.com/jstang/rasic/compare/v1.2.4...v1.2.5) (2022-05-05)


### Bug Fixes

* add changelogs generation again ([8709bfe](https://gitlab.com/jstang/rasic/commit/8709bfe242d736586afea1c6bb252b8102455223))

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
