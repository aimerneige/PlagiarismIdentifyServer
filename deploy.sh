#!/bin/bash
# Copyright (c) 2021 AimerNeige
# aimer.neige@aimerneige.com
# All rights reserved.

./release_build.sh
cd ./raml
raml2html --theme 'raml2html-default-theme' api.raml > api.html
raml2html --theme 'raml2html-full-markdown-theme' api.raml > api.md
cd ..
scp ./release/PlagiarismIdentidyServer root@39.105.116.248:/root/plagiarism-identify-server/PlagiarismIdentidyServer
scp ./raml/api.html root@39.105.116.248:/root/plagiarism-identify-server/file/doc/api.html
scp ./raml/api.md root@39.105.116.248:/root/plagiarism-identify-server/file/doc/api.md
