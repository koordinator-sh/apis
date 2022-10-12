#!/bin/bash

# Define the relative path of API src and dst
API_PATHS_MAP=(
    "apis/config:config"
    "apis/runtime:runtime"
    "apis/slo:slo"
    #"apis/extension:extension"
    #"apis/scheduling:scheduling"
)

help() {
    echo "command format error, usage example"
    echo "$0 \${koordinator_version_tag}"
}

KOORDINATOR_REPO="https://github.com/koordinator-sh/koordinator.git"
VERSION_TAG="$1"
if [[ -z ${VERSION_TAG} ]]; then
    help
    exit 1
fi

TEMP_DIR=$(mktemp -d)
cleanup() {
    echo ">> Removing ${TEMP_DIR}"
    rm -rf ${TEMP_DIR}
}
trap "cleanup" EXIT SIGINT

echo ">> Temporary output directory ${TEMP_DIR}"

KOORDINATOR_REPO_DIR=${TEMP_DIR}/koordinator/
git clone -b ${VERSION_TAG} ${KOORDINATOR_REPO} ${KOORDINATOR_REPO_DIR}

# copy api files from ${src_path} to ${dst_path} excludes zz_generated*.go and *_test.go
# $1: soruce api full dir path: koordinator-sh/koordinator/api/slo
# $2: destination api full dir path, e.g. koordinator-sh/api/slo
function copy_api_files() {
    src_path=$1
    dst_path=$2
    echo ">> copy api files from" ${src_path} "to" ${dst_path} "excludes zz_generated*.go and *_test.go"
    api_files_src_path="$( find ${src_path} -type f \( -name "*.go" -and ! -name "zz_generated*.go" -and ! -name "*_test.go" \) )"
    echo ${api_files_src_path}
    for api_file_src_path in ${api_files_src_path};
    do
        file_relative_path=${api_file_src_path#${src_path}}
        dir_relative_path="$( dirname ${file_relative_path} )"
        dst_file_path="${dst_path}/${file_relative_path}"
        dst_dir_path=${dst_path}/${dir_relative_path}
        mkdir -p ${dst_dir_path}
        cp ${api_file_src_path} ${dst_file_path}
        echo ">> copy api file" ${api_file_src_path} "to" ${dst_file_path}
    done
}

API_REPO_DIR=$( cd $(dirname $0)/..; pwd )

echo ">> copy apis files from" ${KOORDINATOR_REPO_DIR}
for api_path_pair in "${API_PATHS_MAP[@]}";
do
    src_relative_path="${api_path_pair%%:*}"
    dst_relative_path="${api_path_pair##*:}"
    copy_api_files ${KOORDINATOR_REPO_DIR}/${src_relative_path} ${API_REPO_DIR}/${dst_relative_path}
done
