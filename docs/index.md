# Documentación de Implementación de una Criptomoneda en Hyperledger Fabric

Esta documentación sirve como guía personal basada en la experiencia adquirida con el repositorio [chaincode_seminv_uja](https://github.com/dmolina23/chaincode_seminv_uja) y la documentación oficial.

## Objetivos
1. Implementar un contrato inteligente de activo fungible (Token ERC-20).
2. Desplegar el contrato dentro de una red simulada en Hyperledger Fabric.
3. Analizar la arquitectura de consenso (Ordering Service).

## Prerrequisitos
Antes de comenzar, asegúrate de tener instalado:

*   **Docker & Docker Compose**
*   **Go (Golang)**: Versión 1.20+
*   **Binarios de Hyperledger Fabric**: Descargados e incluidos en el PATH.

## Antes de empezar
> Debes tener en cuenta que esta documentación está orientada a sistemas con Ubuntu, Ubuntu Server o Linux. La implementación puede variar en otros Sistemas Operativos.