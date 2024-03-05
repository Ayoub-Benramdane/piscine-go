#!/bin/bash

curl https://learn.zone01oujda.ma/assets/superhero/all.json | jq ".[] | select(.id == $HERO_ID) | .connections.relatives" | tr -d "\""
 
