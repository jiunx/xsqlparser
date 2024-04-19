#!/bin/zsh

    # antlr4 -Dlanguage=Go -no-visitor -package parsing *.g4
source ~/.zshrc

antlr4 -Dlanguage=Go -package parser -visitor  grammer/*.g4

mv grammer/*.go gen/

mv grammer/*.tokens gen/

mv grammer/*.interp gen/