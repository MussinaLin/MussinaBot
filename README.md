# Mussina Bot
This project is named MussinaBot. Lending USD on Bitfinex automatically.

## How to Build  
GOOS=linux GOARCH=amd64 go build -o MussinaBot-EC2  

## Copy Execution File to EC2 Server  
scp -i ~/AWS/MussinaBotKeyPair.pem MussinaBot-EC2 ubuntu@18.181.241.55:
