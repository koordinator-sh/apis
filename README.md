# koordinator-api

Schema of the API types that are served by Koordinator.

## Purpose

This library is the canonical location of the Koordinator API definition.

We recommend using the go types in this repo. You may serialize them directly to JSON.

## Where does it come from?

`koordinator-sh/api` is synced from [https://github.com/koordinator-sh/koordinator/apis](https://github.com/koordinator-sh/koordinator/tree/master/apis).
All modification Code changes are made in origin location(`koordinator-sh/koordinator/apis`), then merged into `koordinator-sh/apis` later.

## For Koordinator developers
`koordinator-sh/apis` is a read-only repo. Please contribute to [https://github.com/koordinator-sh/koordinator/tree/master/apis](https://github.comkoordinator-sh/koordinator/tree/master/apis)
if you have new features to develop. API Changes will be synced to `koordinator-sh/apis` when there is new version released of Koordinator .

## For Koordinator maintainers
To update Koordinator API repo.
1. (optional) add more api directories to clone, append source and destination to API_PATHS_MAP in clone-api-files.sh
```shell script
# vim hack/api-files.sh
API_PATHS_MAP=(
    "apis/slo:slo"
)
```

2. move API files from `koordinator-sh/koordinator/apis` with specified version tag
```shell script
make clone KOORD_VERSION=v0.7.0
```
3. generate deepcopy and clients for API files

3.1 (optional) update binary tool versions and clean local binary tools if necessary in Makefile

3.2 generate files
```shell script
make generate
```
4. (optional)update go.mod
```shell script
make mod
```
5. create a pull request to merge new files
6. create tag for new commit and push
```shell script
git tag ${version-tag}
git push --tags
```
