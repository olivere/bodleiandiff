#!/bin/sh
curl -H 'Content-Type: application/json' -XPUT 'localhost:9200/index01' -d @mapping.json
curl -H 'Content-Type: application/json' -XPUT 'localhost:9200/index02' -d @mapping.json
