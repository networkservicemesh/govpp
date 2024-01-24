#!/bin/bash
set -ex

git config user.name "John Doe"
git config user.email "johndoe@example.com"

function git_cherry_pick ()
{
	refs=$1
	git fetch "https://gerrit.fd.io/r/vpp" ${refs}
	git cherry-pick FETCH_HEAD
	git commit --amend -m "gerrit:${refs#refs/changes/*/} $(git log -1 --pretty=%B)"
}

# NSM cherry picks
git_cherry_pick refs/changes/03/39503/1 # 39503: vppinfra: fix setns typo | https://gerrit.fd.io/r/c/vpp/+/39503
git_cherry_pick refs/changes/28/39528/9 # 39528: ping: Simple binary API for running ping based on events | https://gerrit.fd.io/r/c/vpp/+/39528
git_cherry_pick refs/changes/24/39824/1 # 39824: af_packet: remove UNIX_FILE_EVENT_EDGE_TRIGGERED flag | https://gerrit.fd.io/r/c/vpp/+/39824
git_cherry_pick refs/changes/19/40119/1 # 40119: af_packet: set next0 for AF_PACKET_IF_MODE_ETHERNET mode | https://gerrit.fd.io/r/c/vpp/+/40119
git_cherry_pick refs/changes/46/40246/6 # 40246: ping: Check only PING_RESPONSE_IP4 and PING_RESPONSE_IP6 events | https://gerrit.fd.io/r/c/vpp/+/40246
# Calico cherry picks
git_cherry_pick refs/changes/26/34726/3 # 34726: interface: add buffer stats api | https://gerrit.fd.io/r/c/vpp/+/34726

# Copy Calico local patches
git clone -b v3.26.0 https://github.com/projectcalico/vpp-dataplane.git /vpp-dataplane/
cp /vpp-dataplane/vpplink/generated/patches/* patch/

if [ "$(ls ./patch/*.patch 2> /dev/null)" ]; then
  git apply patch/*.patch
  git add --all
  git commit -m "misc patches"
fi
