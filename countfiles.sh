#!/bin/bash

count=$(find . -type f -o -type d | wc -l)

echo "$count"
