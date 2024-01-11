#!/bin/bash

cd /app

export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:`pwd`/libs

exec ./applet-server