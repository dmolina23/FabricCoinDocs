# Investigación: Algoritmos de Consenso en Fabric

Hyperledger Fabric separa la ejecución de transacciones (Peers) del ordenamiento de las mismas (Orderers). El **Consenso** ocurre en el servicio de ordenamiento.

## Algoritmos Soportados Oficialmente

1.  **Raft (EtcdRaft):**
    *   Es el estándar actual (Crash Fault Tolerant - CFT).
    *   Sigue un modelo "Líder y Seguidores".
    *   Si el líder cae, los nodos votan uno nuevo.

2.  **SmartBFT (Byzantine Fault Tolerant):**
    *   Introducido en Fabric v3.0 (preview en 2.5).
    *   Tolera nodos maliciosos, no solo caídos.

## ¿Cómo implementar un Consenso Personal?

Fabric es modular, permitiendo un "Pluggable Consensus", pero **no es una configuración trivial**. Requiere desarrollo en Go.

### Pasos para desarrollar un Consenso Propio:

#### 1.  **Interfaz `Consenter`**
Debes implementar la interfaz definida en el código fuente de Fabric (`orderer/consensus/consensus.go`).
    
```go
type Consenter interface {
    HandleChain(support ConsenterSupport, metadata *cb.Metadata) (Chain, error)
}
```

#### 2.  **Interfaz `Chain`**
Define cómo se ordenan y empaquetan los bloques.
```go
type Chain interface {
    Order(env *cb.Envelope, configSeq uint64) error
    Configure(config *cb.Envelope, configSeq uint64) error
    WaitReady() error
    Start()
    Halt()
}
```

#### 3.  **Compilación Personalizada**
Una vez escrito tu plugin de consenso en Go, debes recompilar el binario del `orderer` para incluir tu módulo, ya que Go es estático. No se puede cargar como una DLL dinámica fácilmente en versiones antiguas.

### Conclusión
Para propósitos académicos o de producción estándar, se recomienda usar **Raft**. Crear un consenso propio solo se justifica si se requiere una lógica de ordenamiento matemática específica no cubierta por CFT o BFT estándar.