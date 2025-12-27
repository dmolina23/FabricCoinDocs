# Creación de Red: Test Network Oficial

Hyperledger proporciona un script robusto llamado `test-network.sh` que levanta:

- 1 Org Ordenadora (Orderer)
- 2 Orgs Peer (Org1 y Org2)
- 1 Canal por defecto

## Pasos para iniciar

1. Navegar al directorio de `fabric-samples/test-network`.
2. Limpiar entornos previos:
    ```bash
    ./network.sh down
    ```
3. Levantar la red y crear el canal `mychannel`:
    ```bash
    ./network.sh up createChannel -c mychannel -ca
    ```
    *Nota: La bandera `-ca` levanta Autoridades de Certificación reales en lugar de usar cryptogen.*

4. Desplegar el chaincode:
    ```bash
    ./network.sh deployCC -ccn mycoin -ccp ../token-erc-20/chaincode-go/ -ccl go
    ```

## Validación
Verifica que los contenedores están corriendo:

```bash
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
```