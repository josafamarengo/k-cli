#!/bin/bash

REPO="josafamarengo/k-cli"
TOOL="k"

INSTALL_DIR="/usr/local/bin"

# Verificar dependências
command -v curl >/dev/null 2>&1 || { echo >&2 "curl n  o est   instalado. Por favor, instale curl."; exit 1; }
command -v tar >/dev/null 2>&1 || { echo >&2 "tar n  o est   instalado. Por favor, instale tar."; exit 1; }
command -v sudo >/dev/null 2>&1 || { echo >&2 "sudo n  o est   instalado. Por favor, instale sudo."; exit 1; }

# Obter a última versão
# LATEST_VERSION=$(curl -s https://api.github.com/repos/$REPO/releases/latest | grep 'tag_name' | cut -d\" -f4)
LATEST_VERSION="v0.1.0"

if [ -z "$LATEST_VERSION" ]; then
  echo "Erro ao obter a   ltima vers  o. Verifique se o reposit  rio est   correto."
  exit 1
fi

URL="https://github.com/$REPO/releases/download/$LATEST_VERSION/$TOOL-linux-amd64.tar.gz"

TEMP_FILE="/tmp/$TOOL-$LATEST_VERSION.tar.gz"

echo "Baixando $TOOL vers  o $LATEST_VERSION..."
curl -L $URL -o $TEMP_FILE

if [ $? -ne 0 ]; then
  echo "Erro ao baixar o arquivo. Verifique a URL e a conex  o."
  exit 1
fi

if [ ! -f $TEMP_FILE ]; then
  echo "O download falhou ou o arquivo não foi encontrado."
  exit 1
fi

echo "Descompactando $TOOL..."
tar -xzvf $TEMP_FILE -C /tmp

if [ $? -ne 0 ]; then
  echo "Erro ao descompactar o arquivo. Verifique o formato."
  rm -f $TEMP_FILE
  exit 1
fi

echo "Movendo $TOOL para $INSTALL_DIR..."
sudo mv /tmp/$TOOL $INSTALL_DIR

if [ $? -ne 0 ]; then
  echo "Erro ao mover o bin  rio para $INSTALL_DIR. Verifique as permiss  es."
  exit 1
fi

chmod +x $INSTALL_DIR/$TOOL

if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
  echo "Adicionando $INSTALL_DIR ao PATH..."
  echo "export PATH=\$PATH:$INSTALL_DIR" >> ~/.bashrc
  source ~/.bashrc
fi

rm -f $TEMP_FILE

echo "$TOOL versão $LATEST_VERSION instalado com sucesso!"


