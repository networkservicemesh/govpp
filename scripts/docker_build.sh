#!/bin/bash -x
# shellcheck disable=SC2086

ip=$1
sshkey=$2
SSH_OPTS="-o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -o IdentitiesOnly=yes -i ${sshkey}"

# Create github_runner folder
ssh ${SSH_OPTS} root@${ip} mkdir actions-runner
curl -o actions-runner-linux-x64-2.305.0.tar.gz -L https://github.com/actions/runner/releases/download/v2.305.0/actions-runner-linux-x64-2.305.0.tar.gz

