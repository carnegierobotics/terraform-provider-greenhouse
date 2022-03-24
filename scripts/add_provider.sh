#!/bin/bash
cd ../src
go build -o terraform-provider-greenhouse
export OS_ARCH="$(go env GOHOSTOS)_$(go env GOHOSTARCH)"
export VERSION="$(cat ../src/VERSION)"
export PLUGIN_DIR="$HOME/github/greenhouse_terraform_control_repo/terraform/terraform.d/plugins/carnegierobotics.com/hr/greenhouse/0.0.1/linux_amd64"
mkdir -p $PLUGIN_DIR
mv terraform-provider-greenhouse $PLUGIN_DIR
