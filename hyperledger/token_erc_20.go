package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract define la estructura del contrato
type SmartContract struct {
	contractapi.Contract
}

// TokenConfig guarda los datos de la moneda
type TokenConfig struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals string `json:"decimals"`
}

// Initialize: Configura el nombre, símbolo y decimales del token
func (s *SmartContract) Initialize(ctx contractapi.TransactionContextInterface, name string, symbol string, decimals string) error {
	// Guardamos la configuración en la ledger bajo la clave "config"
	config := TokenConfig{Name: name, Symbol: symbol, Decimals: decimals}
	configBytes, _ := json.Marshal(config)

	err := ctx.GetStub().PutState("config", configBytes)
	if err != nil {
		return fmt.Errorf("fallo al guardar config: %v", err)
	}
	return nil
}

// Mint: Crea nuevos tokens y los asigna al que llama a la función (msg.sender)
func (s *SmartContract) Mint(ctx contractapi.TransactionContextInterface, amount int) error {
	// 1. Obtener ID del cliente (quien ejecuta el comando)
	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return fmt.Errorf("no se pudo obtener ID del cliente: %v", err)
	}

	// 2. Obtener saldo actual
	currentBalance, _ := s.getBalance(ctx, clientID)

	// 3. Sumar y guardar
	newBalance := currentBalance + amount
	return s.putBalance(ctx, clientID, newBalance)
}

// Transfer: Envía tokens del usuario actual a otro
func (s *SmartContract) Transfer(ctx contractapi.TransactionContextInterface, recipientID string, amount int) error {
	// 1. Obtener ID del remitente
	senderID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return fmt.Errorf("error obteniendo ID: %v", err)
	}

	// 2. Verificar saldo del remitente
	senderBalance, _ := s.getBalance(ctx, senderID)
	if senderBalance < amount {
		return fmt.Errorf("saldo insuficiente. Tienes %d, intentas enviar %d", senderBalance, amount)
	}

	// 3. Restar al remitente
	err = s.putBalance(ctx, senderID, senderBalance-amount)
	if err != nil {
		return err
	}

	// 4. Sumar al destinatario
	recipientBalance, _ := s.getBalance(ctx, recipientID)
	err = s.putBalance(ctx, recipientID, recipientBalance+amount)

	return err
}

// ClientAccountBalance: Consulta el saldo del propio usuario
func (s *SmartContract) ClientAccountBalance(ctx contractapi.TransactionContextInterface) (int, error) {
	clientID, _ := ctx.GetClientIdentity().GetID()
	return s.getBalance(ctx, clientID)
}

// --- Funciones Auxiliares (Privadas) ---

func (s *SmartContract) getBalance(ctx contractapi.TransactionContextInterface, id string) (int, error) {
	bytes, _ := ctx.GetStub().GetState(id)
	if bytes == nil {
		return 0, nil
	}
	var balance int
	json.Unmarshal(bytes, &balance)
	return balance, nil
}

func (s *SmartContract) putBalance(ctx contractapi.TransactionContextInterface, id string, balance int) error {
	bytes, _ := json.Marshal(balance)
	return ctx.GetStub().PutState(id, bytes)
}

func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		log.Panicf("Error creando chaincode: %v", err)
	}
	if err := chaincode.Start(); err != nil {
		log.Panicf("Error iniciando chaincode: %v", err)
	}
}
