# Token ERC-20 en Hyperledger Fabric

Aunque ERC-20 es un estándar de Ethereum, en Hyperledger Fabric podemos implementar la misma lógica de negocio (interfaz) para crear un token fungible.

## Diferencias Clave
| Característica | Ethereum ERC-20 | Fabric Token Chaincode |
| :--- | :--- | :--- |
| **Costo (Gas)** | Se paga gas por tx | No hay gas (costo de infraestructura) |
| **Identidad** | Wallet pública (0x...) | Certificado X.509 (MSP) |
| **Privacidad** | Pública (generalmente) | Canales privados o Colecciones |

## Métodos a Implementar
El Chaincode debe soportar:

- `Mint()`: Crear monedas.
- `Burn()`: Destruir monedas.
- `Transfer(to, amount)`: Mover saldo.
- `BalanceOf(account)`: Consultar saldo.
- `TotalSupply()`: Ver total en circulación.

&nbsp;

> En los archivos de _fabric-samples_ podemos encontrar un ejemplo de token ERC20 semi-implementado