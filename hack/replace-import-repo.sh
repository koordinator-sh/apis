#!/bin/bash

function custom_sed(){
    perl -i -pe $@
}

REPO_PATH="$( cd -- "$(dirname "$0")" >/dev/null 2>&1 ; cd ..; pwd -P )"
API_FILES="$( find $REPO_PATH -name "*.go" )"

echo ">> replace package cross imports"

for api_file in $API_FILES
do
    custom_sed 's#github.com/koordinator-sh/koordinator/apis#github.com/koordinator-sh/apis#g' ${api_file}
    echo ${api_file}
done


