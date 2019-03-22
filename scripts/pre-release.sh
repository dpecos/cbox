export BRANCH=$(git rev-parse --abbrev-ref HEAD)
export COMMIT=$(git rev-parse HEAD)
export BUILD=$BRANCH-$COMMIT@$(date +%FT%T%z)
export TAG=$(git describe --tags 2> /dev/null)
export ENV=prod

# if [[ "$TAG" ]]
# then
#   export VERSION=$TAG
# else
#   export VERSION=0.0.0
# fi
