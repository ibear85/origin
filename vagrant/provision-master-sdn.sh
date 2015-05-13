#!/bin/bash
set -ex
source $(dirname $0)/provision-config.sh

pushd $HOME
# build openshift-sdn
if [ -d openshift-sdn ]; then
    cd openshift-sdn
    git fetch origin
    git reset --hard origin/master
else
    git clone https://github.com/openshift/openshift-sdn
    cd openshift-sdn
fi

make clean
make
make install
popd

# no need to start openshift-sdn, as it is integrated with openshift binary
