#!/bin/bash


#update helm image version in origin if there is ome repo this particular job is not needed
sed -r -i '' 's/tag: ".*"/tag: "'$1'"/' /Users/prima/work/golang/src/github.com/prima101112/crudsample/helm/values.yaml
git commit helm/values.yaml -m "update helm version to $1"
#push to main origin to update so argodb could sync later in pipeline
git push origin main

printf "\n===== PRE : Create new branch from main =====\n\n"
git checkout -b $1
printf "\n===== ACTION : push to the branch version =====\n\n"
git push deployment $1
printf "\n===== POST : get back to main =====\n\n"
git checkout main
printf "\n===== POST : delete the branch from local =====\n\n"
git branch -d $1