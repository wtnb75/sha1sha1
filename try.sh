#! /bin/bash

seed=${1-da39a3ee5e6b4b0d3255bfef95601890afd80709}
seedlen=$(echo $(echo -n ${seed} | wc -c))

[ -x sha1sha1 ] || go build
for i in $(seq 8); do
  ns=$(echo ${seed} | cut -c1-$(expr $seedlen - $i))
  ./sha1sha1 -prefix $ns -len $i > ${seed}.${i}.txt
done
