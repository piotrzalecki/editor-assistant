#!/bin/bash


INPUT=$1

echo $INPUT

echo "LLAMA3"
echo "<<$INPUT>>" | ollama run editor-assistant:llama3
echo
echo "LLAMA2"
echo "<<$INPUT>>"| ollama run editor-assistant:llama2
echo
echo "GEMMA"
echo "<<$INPUT>>" | ollama run editor-assistant:gemma2