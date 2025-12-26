# Creación de una Red Personalizada (Manual)

Crear una red sin el script `test-network` implica entender la configuración criptográfica y los artefactos de génesis.

## 1. Estructura de Archivos
Debes crear una carpeta con:

- `crypto-config.yaml`: Para generar certificados (si usas cryptogen).
- `configtx.yaml`: Para definir el Bloque Génesis y la configuración del Canal.
- `docker-compose.yaml`: Para definir los contenedores (Peers, Orderers, CLI).

## 2. Generación de Identidades (Criptografía)
Si no usas una CA Server, usa `cryptogen`:

```bash
cryptogen generate --config=./crypto-config.yaml --output="organizations"
```

## 3. Artefactos del Canal y Génesis
Debes generar el bloque génesis del sistema ordenadar:

```bash
configtxgen -profile TwoOrgsOrdererGenesis -channelID system-channel -outputBlock ./channel-artifacts/genesis.block
```

Y la transacción de creación del canal:

```bash
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID mycustomchannel
```

## 4. Despliegue
Ejecuta `docker-compose up -d`. Una vez arriba, debes entrar al contenedor CLI y ejecutar manualmente:

```bash
peer channel create ...
peer channel join ... #  para cada peer
peer channel update ... # para definir anchor peers
```