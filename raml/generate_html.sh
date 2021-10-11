#!/bin/bash
# Copyright (c) 2021 AimerNeige
# aimer.neige@aimerneige.com
# All rights reserved.

PS3='Please enter your choice: '
options=("raml2html-default-theme" "raml2html-plain-theme" "raml2html-slate-theme" "raml2html-kaa-theme" "raml2html-full-markdown-theme" "Quit")
select opt in "${options[@]}"
do
    case $opt in
        "raml2html-default-theme")
            echo "you chose raml2html-default-theme"
            raml2html --theme 'raml2html-default-theme' api.raml > api.html
            break
            ;;
        "raml2html-plain-theme")
            echo "you chose raml2html-plain-theme"
            raml2html --theme 'raml2html-plain-theme' api.raml > api.html
            break
            ;;
        "raml2html-slate-theme")
            echo "you chose raml2html-slate-theme"
            raml2html --theme 'raml2html-slate-theme' api.raml > api.html
            break
            ;;
        "raml2html-kaa-theme")
            echo "you chose raml2html-kaa-theme"
            raml2html --theme 'raml2html-kaa-theme' api.raml > api.html
            break
            ;;
        "raml2html-full-markdown-theme")
            echo "you chose raml2html-kaa-theme"
            raml2html --theme 'raml2html-full-markdown-theme' api.raml > api.md
            break
            ;;
        "Quit")
            break
            ;;
        *) echo "invalid option $REPLY";;
    esac
done
