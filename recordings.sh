#!/bin/bash

cd recordings || true

for FILE in $(ls *.pcm); do
    ffmpeg -f s16le -ar 44.1k -ac 2 -i ${FILE} $(echo ${FILE} | sed s/\.pcm/\.wav/) > /dev/null
    rm ${FILE}
done

for FILE in $(ls *.wav); do
    zip $(echo ${FILE} | sed s/\.wav/\.zip/) ${FILE}
    rm ${FILE}
done