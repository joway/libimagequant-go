# /bin/bash

git submodule foreach git reset --hard
git submodule update --remote --init
rm ./pngquant/*.h
rm ./pngquant/*.c
cp ./lib/*.h ./pngquant/
cp ./lib/*.c ./pngquant/
rm ./pngquant/example.c
