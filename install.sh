#!/bin/bash

REPO="josafamarengo/k-cli"
TOOL="k"

INSTALL_DIR="/usr/local/bin"

command -v curl >/dev/null 2>&1 || { echo >&2 "curl is not installed. Please install curl."; exit 1; }
command -v tar >/dev/null 2>&1 || { echo >&2 "tar is not installed. Please install tar."; exit 1; }

# LATEST_VERSION=$(curl -s https://api.github.com/repos/$REPO/releases/latest | grep 'tag_name' | cut -d\" -f4)
LATEST_VERSION="v0.1.0"

if [ -z "$LATEST_VERSION" ]; then
  echo "Error getting latest version. Please check if the repository is correct."
  exit 1
fi

URL="https://github.com/$REPO/releases/download/$LATEST_VERSION/$TOOL-linux-amd64.tar.gz"

TEMP_FILE="/tmp/$TOOL-$LATEST_VERSION.tar.gz"

echo "Downloading $TOOL version $LATEST_VERSION..."
curl -L $URL -o $TEMP_FILE

if [ $? -ne 0 ]; then
  echo "Error downloading the file. Check the URL and connection."
  exit 1
fi

if [ ! -f $TEMP_FILE ]; then
  echo "The download failed or the file was not found."
  exit 1
fi

echo "Unpacking $TOOL..."
tar -xzvf $TEMP_FILE -C /tmp

if [ $? -ne 0 ]; then
  echo "Error extracting the file. Check the format."
  rm -f $TEMP_FILE
  exit 1
fi

echo "Moving $TOOL to $INSTALL_DIR..."
sudo mv /tmp/$TOOL $INSTALL_DIR

if [ $? -ne 0 ]; then
  echo "Error moving the binary to $INSTALL_DIR. Check the permissions."
  exit 1
fi

chmod +x $INSTALL_DIR/$TOOL

if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
  echo "Adding $INSTALL_DIR to PATH..."
  echo "export PATH=\$PATH:$INSTALL_DIR" >> ~/.bashrc
  source ~/.bashrc
fi

rm -f $TEMP_FILE

echo "$TOOL version $LATEST_VERSION installed successfully!"


