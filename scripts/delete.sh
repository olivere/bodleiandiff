#!/bin/sh
curl -H 'Content-Type: application/json' -XDELETE 'localhost:9200/index01'
curl -H 'Content-Type: application/json' -XDELETE 'localhost:9200/index02'
