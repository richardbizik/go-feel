#!/bin/bash

layer_id="f19a63f39269460fedfac3ac50b28f3b809cd5d9a0f01d9b0383fe21fc7eeec9"

docker image ls -qa \
  | while read -r imageid; do
      docker image inspect --format '{{json .GraphDriver.Data}}' "$imageid" \
        | grep -q "$layer_id" && echo "$imageid"
    done
