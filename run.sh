#!/bin/bash

YEAR="$1"
QUESTION="$2"

if [[ -z $YEAR || -z $QUESTION ]];then
  "usage: ./run.sh YEAR QUESTION"
  exit 0
fi

if [[ "$YEAR" = "2015" ]]; then
  question="2015/$QUESTION/main.go"
  input="2015/$QUESTION/input"
  if [[ -f "2015/$QUESTION/input" ]];then
    go run "$question" < "$input"
  else
    go run "$question"
  fi
fi
