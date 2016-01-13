#!/bin/bash
set -e
R=$(curl --write-out %{http_code} --silent --output /dev/null -F dockerfile=@Dockerfile.fail_unittest localhost:8080/validate  )

echo $R

if [[ $R -ne "200" ]]; then
    echo "Host is not in compliance with the dockerfile rules"
    exit 1
fi

echo "Dockerfile is valid"
exit 0
