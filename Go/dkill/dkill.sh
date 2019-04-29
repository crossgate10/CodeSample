#!/bin/bash

netstat -aon | grep "YourPort" | awk '{print $5}' | xargs kill -f $i