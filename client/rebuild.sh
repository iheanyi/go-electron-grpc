#!/bin/bash

LDFLAGS=-L/usr/local/opt/openssl/lib CPPFLAGS=-I/usr/local/opt/openssl/include ./node_modules/.bin/electron-rebuild

