#!/bin/sh -eux

sed -i Makefile -e "s/^VERSION=${1}/VERSION=${2}/"
git add Makefile
git commit -m "Update version to ${2}"
