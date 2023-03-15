#!/bin/sh

set -e

# If the first argument looks like a flag, assume we want to run the
# command with the default arguments.
if [ "${1#-}" != "$1" ]; then
    set -- /app/main "$@"
fi

# If the first argument is "docker-entrypoint.sh", we assume the user
# wants to run their own init process, for example a `bash` shell to
# explore the container.
if [ "$1" = 'entrypoint.sh' ]; then
    set -- "$@"
fi

exec "$@"
