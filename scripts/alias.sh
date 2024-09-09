#!/bin/bash


if ds >> /dev/null; then
    echo "Alias 'ds' exists."
    exit 0
else
    echo "Alias 'ds' not found. Adding it..."
    echo 'alias ds="dcshort"' >> ~/.zshrc
    echo "Alias added successfully."
    zsh
fi
