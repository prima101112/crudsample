printf "\n===== PRE : Create new branch from main =====/n/n"
git checkout -b $1
printf "/n===== ACTION : push to the branch version =====/n/n"
git push deployment $1
printf "/n===== POST : get back to main =====/n/n"
git checkout main
printf "/n===== POST : delete the branch from local =====/n/n"
git branch -d $1