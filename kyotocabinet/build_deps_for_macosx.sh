#!/bin/sh
cd ./deps/kyotocabinet-1.2.76
cp -f ../macosx/Makefile.in ./Makefile.in
./configure
make
sudo make install
