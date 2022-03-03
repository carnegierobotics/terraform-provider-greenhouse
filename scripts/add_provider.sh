#!/bin/bash
cd ../src
tfplugindocs generate
go build -o terraform-provider-greenhouse
export OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"
export VERSION="$(cat ../src/VERSION)"
export PLUGIN_DIR="$HOME/.terraform.d/plugins/carnegierobotics.com/hr/greenhouse/$VERSION/$OS_ARCH"
mkdir -p $PLUGIN_DIR
mv terraform-provider-greenhouse $PLUGIN_DIR
