#! /bin/bash

find .  -name "*.sh" | sed -e 's/.sh//g' |  cut -d '.' -f2 |  cut -d '/' -f2