#!/bin/bash
set -e
R=$(curl --write-out %{http_code} --silent --output /dev/null -F dockerfile=@Dockerfile.fail_unittest localhost:8080/validate  )

if [[ $R -eq "200" ]]; then
    exit 1
fi

exit 0
