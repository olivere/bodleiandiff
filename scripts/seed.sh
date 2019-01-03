#!/bin/sh
# example seed file for elasticsearch index comparsion
curl -H 'Content-Type: application/json' -XPUT 'localhost:9200/index01/fim/2IpLWmcBa7Etpo2IE72A' -d '{"file.path":"/a_folder/about.txt","file.size":"262","hash.sha512":"b1f76143422792f03889fe21d97f4579dd4faa44e56ecfe748d3ff79e694425b632ce9bc45d5747c459dc2127c51e8a2d73042ded721002d0c0b234b6de52b86"}'
curl -H 'Content-Type: application/json' -XPUT 'localhost:9200/index01/fim/HopLWmcBa7Etpo2IFcAz' -d '{"file.path":"/b_folder/basic.xml","file.size":"995","hash.sha512":"bd799a593f26e5ea02c44ddeb76bc6ac1bd92dcc02427a997652b149fb9d6b6bac05aac664e2d59b0cb29d9fa37818952e5971418f84e563a6c0e355a8a56d57"}'
curl -H 'Content-Type: application/json' -XPUT 'localhost:9200/index01/fim/A4pLWmcBa7Etpo2IFL9D' -d '{"file.path":"/c_folder/cropped.jp2","file.size":"8194","hash.sha512":"06b3025808c60f7362312e334c4b179ec9f003d85c7bfeaa2c4c56e7a7d173008ea95ad283cd9e324814c9a2492055a7d019f7b430aa3e72e38c7cc174c40da3"}'

curl -H 'Content-Type: application/json' -XPUT 'localhost:9200/index02/fim/jYpLWmcBa7Etpo2IFL5D1' -d '{"file.path":"/a_folder/about.txt","file.size":"262","hash.sha512":"b1f76143422792f03889fe21d97f4579dd4faa44e56ecfe748d3ff79e694425b632ce9bc45d5747c459dc2127c51e8a2d73042ded721002d0c0b234b6de52b86"}'
curl -H 'Content-Type: application/json' -XPUT 'localhost:9200/index02/fim/PIpLWmcBa7Etpo2IFL-P' -d '{"file.path":"/b_folder/basic.xml","file.size":"995","hash.sha512":"bd799a593f26e5ea02c44ddeb76bc6ac1bd92dcc02427a997652b149fb9d6b6bac05aac664e2d59b0cb29d9fa37818952e5971418f84e563a6c0e355a8a56d57"}'
curl -H 'Content-Type: application/json' -XPUT 'localhost:9200/index02/fim/vIpLWmcBa7Etpo2IFL_j' -d '{"file.path":"/c_folder/cropped.jp2","file.size":"8194","hash.sha512":"fb19a3d584f2fff59ffbfdc706b964fbca4c74d6c3af0b51e2945810d9fc986b4806f91080da599835f73b6e5bfe1fd5a3fe48376aed8565f964ca0e8260a667"}'