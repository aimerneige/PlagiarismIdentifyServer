#!/bin/bash
# Copyright (c) 2021 AimerNeige
# aimer.neige@aimerneige.com
# All rights reserved.

./release_build.sh
scp ./release/PlagiarismIdentidyServer root@39.105.116.248:/root/plagiarism-identify-server/PlagiarismIdentidyServer
scp ./raml/api.html root@39.105.116.248:/root/plagiarism-identify-server/file/doc/api.html
scp ./raml/api.md root@39.105.116.248:/root/plagiarism-identify-server/file/doc/api.md
