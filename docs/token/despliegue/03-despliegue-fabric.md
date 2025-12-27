Este despliegue del chaincode es mucho más sencillo que el manual ya que no tenemos que realizar ninguna tarea auxiliar.

## Ejecución del comando de despliegue
```bash
./network.sh deployCC -ccn mycoin -ccp ../token-erc-20/chaincode-go/ -ccl go
```