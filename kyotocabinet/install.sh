#!/bin/sh
. ./deps.sh
CGO_CFLAGS="-I$DEPS_PATH" CGO_LDFLAGS="-L$DEPS_PATH" LD_PRELOAD=$DEPS_PATH/libkyotocabinet.so go install
