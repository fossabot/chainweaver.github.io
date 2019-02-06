#!/bin/bash

if [ $# -ne 1 ]; then
  echo "Requires at least 1 argument" 1>&2
  exit 1
fi

herokuBaseAppId=$1

if [ -z "$CIRCLE_TAG" ]; then
    if [ -z "$CIRCLE_BRANCH" ]; then
        appSuffix=mainline
    else
        appSuffix=$CIRCLE_BRANCH
    fi
else
    appSuffix=${CIRCLE_TAG//./-}
fi
herokuAppId="$herokuBaseAppId-$appSuffix"
echo $herokuAppId