#!/bin/bash

ollama create editor-assistant:llama3 --file Modefiles/Modefile.llama3
ollama create editor-assistant:llama2 --file Modefiles/Modefile.llama2
ollama create editor-assistant:gemma2 --file Modefiles/Modefile.gemma2