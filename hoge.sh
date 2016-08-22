#!/bin/sh

HOGE="hoge"
case $HOGE in
  *huga*)
    echo 'ok';;
  *)
    break;;
esac
