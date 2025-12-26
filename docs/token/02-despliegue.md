# Despliegue del Token ERC-20

Usaremos el código base oficial de `fabric-samples/token-erc-20/chaincode-go`

## 1. Empaquetado del Chaincode
Configura las variables de entorno y empaqueta:

```bash
peer lifecycle chaincode package token_erc20.tar.gz --path ../token-erc-20/chaincode-go/ --lang golang --label token_erc20_1.0
```


## 2. Instalación
Instala el paquete en los peers de ambas organizaciones (Org1 y Org2).

```bash
peer lifecycle chaincode install token_erc20.tar.gz
```

## 3. Aprobación y Commit
Ambas organizaciones deben aprobar la definición.

```bash
peer lifecycle chaincode approveformyorg -o localhost:7050 --channelID mychannel --name token_erc20 --version 1.0 --sequence 1 ... (flags TLS)
```

Una vez aprobado, haz commit:

```bash
peer lifecycle chaincode commit ... --name token_erc20 ...
```

## 4. Inicialización y Uso (Minting)
Inicializa el token creando el suministro inicial para un usuario.

```bash
peer chaincode invoke ... -c '{"function":"Mint","Args":["5000"]}'
```

### Transferir tokens a otro usuario:

```bash
peer chaincode invoke ... -c '{"function":"Transfer","Args":["x509::CN=User1...","100"]}'
```