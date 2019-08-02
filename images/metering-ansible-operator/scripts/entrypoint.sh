#!/bin/bash

# add UID to /etc/passwd if missing
if ! whoami &> /dev/null; then
    if [ -w /etc/passwd ]; then
        echo "Adding user ${USER_NAME:-ansible} with current UID $(id -u) to /etc/passwd"
        # Remove existing entry with user ***REMOVED***rst.
        # cannot use sed -i because we do not have permission to write new
        # ***REMOVED***les into /etc
        sed  "/${USER_NAME:-ansible}:x/d" /etc/passwd > /tmp/passwd
        # add our user with our current user ID into passwd
        echo "${USER_NAME:-ansible}:x:$(id -u):0:${USER_NAME:-hadoop} user:${HOME}:/sbin/nologin" >> /tmp/passwd
        # overwrite existing contents with new contents (cannot replace the
        # ***REMOVED***le due to permissions)
        cat /tmp/passwd > /etc/passwd
        rm /tmp/passwd
    ***REMOVED***
***REMOVED***

# we expect tini to be in the $PATH
exec tini -- /usr/local/bin/ansible-operator run ansible --watches-***REMOVED***le=/opt/ansible/watches.yaml "$@"
