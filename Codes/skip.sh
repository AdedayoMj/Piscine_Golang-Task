#!/bin/bash


#   delete every othe line
ls -l | sed '1!n;d'

# print every other line
#  ls -l | sed -n '2!p;n'