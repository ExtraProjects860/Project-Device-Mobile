#!/bin/bash

OS=$(uname -s)

echo "Iniciando instalação do ngrok!"
echo "Antes de iniciar, tenha o curl ou netcat instalado na máquina"

if docker ps > /dev/null 2>&1; then
    echo "✅ Docker reconhecido na máquina, prosseguindo..."
else
    echo "❌ Docker não encontrado. Instale o Docker para continuar."
    exit 1
fi

echo -n "Insira seu token de autenticação: "
read AUTH_TOKEN

while true; do
    echo -n "Escolha a porta da aplicação: "
    read PORT_APP

    if command -v nc > /dev/null 2>&1; then
        if nc -z localhost $PORT_APP; then
            echo "✅ Porta $PORT_APP está respondendo"
            break
        fi
    elif command -v curl > /dev/null 2>&1; then
        if curl -s http://localhost:$PORT_APP > /dev/null 2>&1; then
            echo "✅ Porta $PORT_APP está respondendo"
            break
        fi
    fi

    echo "⚠️ Porta $PORT_APP não respondeu. Tente novamente..."
done

if docker pull ngrok/ngrok > /dev/null 2>&1; then
    echo "✅ Ngrok via Docker disponível, prosseguindo..."
else
    echo "❌ Erro ao instalar Ngrok via Docker"
    exit 1
fi

if docker inspect ngrok > /dev/null 2>&1; then
    echo "Ngrok já instalado no container, iniciando..."
    docker start -ai ngrok
else
    echo "Rodando ngrok pela primeira vez..."
    if [[ "$OS" == "Linux"* ]]; then
        docker run --net=host --name ngrok -it -e NGROK_AUTHTOKEN=$AUTH_TOKEN ngrok/ngrok:latest http $PORT_APP
    else
        docker run --name ngrok -it -e NGROK_AUTHTOKEN=$AUTH_TOKEN -p 4040:4040 ngrok/ngrok:latest http host.docker.internal:$PORT_APP
    fi
fi

echo "✅ Processo concluído. O domínio do ngrok deve aparecer acima."
