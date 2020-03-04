#!/bin/bash

set -e;

DIR="mychroot"
if [ -d "$DIR" ]; then
    echo "folder exists, exiting."
    exit
fi

mkdir mychroot

wget -qO- https://mirror.yandex.ru/gentoo-distfiles/releases/amd64/autobuilds/20200226T214502Z/stage3-amd64-20200226T214502Z.tar.xz | tar xvJf - -C ./mychroot