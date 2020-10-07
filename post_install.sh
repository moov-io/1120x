#!/bin/sh

OS_NAME=$TRAVIS_OS_NAME
if [ "$OS_NAME" = "" ]; then
  if [ $(uname -s) = "Darwin" ]; then
    export OS_NAME=osx
  else
	  export OS_NAME=linux
  fi
fi

if which xsltproc >/dev/null; then
    exit
fi

if [ "$OS_NAME" = "linux" ]; then
  sudo apt-get update -y
  sudo apt-get install -y libxml2
  sudo apt-get install -y xsltproc
fi

if [ "$OS_NAME" = "osx" ]; then
  brew install pkg-config
  brew install libxml2
  sudo ln -s /usr/local/opt/libxml2/include/libxml2/libxml /usr/local/include/libxml
  brew install libxslt
fi
