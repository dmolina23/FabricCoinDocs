> Ten en cuenta que existen varias formas de instalar los archivos de Hyperledger Fabric. En esta guía utilizamos esta forma de instalación porque es la que minimiza los posibles errores derivados de la propia instalación y de dependencias externas a Hyperledger.

## 1. Crear directorio de trabajo
```bash
mkdir -p ~/hyperledger/fabric
cd ~/hyperledger/fabric
```

## 2. Descargar el script de bootstrap y ejecutar
```bash
curl -sSL https://bit.ly/2ysbOFE -o bootstrap.sh
chmod +x bootstrap.sh
./bootstrap.sh 2.2.5 1.5.2
```

## 3. Añadir binarios de Fabric al PATH
```bash
echo 'export PATH=$PATH:$HOME/hyperledger/fabric/fabric-samples/bin' >> ~/.profile
source ~/.profile
```