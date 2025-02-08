#!/bin/bash

echo "Script execution time:" $(date '+%d-%m-%Y %H:%M:%S')
echo "...................................................."
echo "  BASH SCRIPT FOR COMPILE GO PROGRAM                "
echo "...................................................."
echo

## Variables
RANDOM_NAME="executable_file"

read -p "Are you sure to compile the GO program? (say 'yes' or 'no' in lowercase): " USER_ANS
echo "User entered the input as: ${USER_ANS}"

if [[ $USER_ANS == "yes" || $USER_ANS == "y" ]]; then
    echo "Good to compile the code..."
    cd ../
    ls -l
    read -p "Enter the script name please (e.g., program.go): " SCRIPT_NAME
    echo "User entered the script name as: $SCRIPT_NAME"
    echo "GO compiling begin in a few milliseconds..."
    
    # Create a dynamic executable name using script name and timestamp
    EXEC_NAME="$RANDOM_NAME_$(basename $SCRIPT_NAME .go)_$(date '+%d-%m-%Y_%H-%M-%S')"
    
    go build -o "$EXEC_NAME" "$SCRIPT_NAME"
    
    echo "Cross-checking if the code is compiled or not..."
    
    # Check if the compiled executable exists
    if [ -f "$EXEC_NAME" ]; then
        echo "Compilation successful! Executable: $EXEC_NAME"
    else
        echo "Compilation failed."
    fi
elif [[ $USER_ANS == "no" || $USER_ANS == "n" ]]; then
    echo "No compile, since user has entered 'no'..."
else
    echo "Invalid input! Please type 'yes' or 'no'..."
fi
