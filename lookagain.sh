#! /bin/bash

find .  -name "*.sh" |  cut -d '.' -f2 |  cut -d '/' -f2 | sed 's/.sh//g' | sed 's/tes//g' 