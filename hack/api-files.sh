#!/bin/bash

# Define the relative path of API src and dst
API_PATHS_MAP=(
    "apis/analysis:analysis"
    "apis/config:config"
    "apis/configuration:configuration"
    "apis/extension:extension"
    "apis/quota:quota"
    "apis/runtime:runtime"
    "apis/scheduling:scheduling"
    "apis/scheme:scheme"
    "apis/slo:slo"
    "apis/thirdparty:thirdparty"
)

# Define files in apis should be excluded
EXCLUDED_PATHS=(
    #"apis/extension/xxx/"
)