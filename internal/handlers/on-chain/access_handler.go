package handler

import (
	"context"
	"drone-backend/smart-contracts/PDP"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PDPHandler struct {
    Instance *PDP.PDP
    Auth     *bind.TransactOpts
    Client   *ethclient.Client
}

func NewPDPHandler(instance *PDP.PDP, auth *bind.TransactOpts, client *ethclient.Client) *PDPHandler {
    return &PDPHandler{
        Instance: instance,
        Auth:     auth,
        Client:   client,
    }
}

func (h *PDPHandler) Layer1EvaluateAccess(entityID uint, _model_type string, _zone int, _startTime string, _endTime string, _accessGranted bool) (bool, string, error) {
    if h.Client == nil {
        log.Printf("Ethereum client is nil")
        return false, "", fmt.Errorf("Ethereum client is not initialized")
    }

    if h.Instance == nil {
        log.Printf("PDP instance is nil")
        return false, "", fmt.Errorf("PDP instance is not initialized")
    }
    
    bigID := big.NewInt(int64(entityID))
    log.Printf("Starting EvaluateAccess for entityID: %d", entityID)
    bigZone := big.NewInt((int64(_zone)))

    tx, err := h.Instance.Layer1EvaluateAccess(h.Auth, bigID, _model_type, bigZone, _startTime, _endTime, _accessGranted )
    if err != nil {
        log.Printf("Failed to call EvaluateAccess: %v", err)
        return false, "", err
    }

    log.Printf("Transaction sent: %s", tx.Hash().Hex())

    log.Println("Waiting for transaction to be mined...")

    ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    defer cancel()
    
    receipt, err := bind.WaitMined(ctx, h.Client, tx)
    if err != nil {
        log.Printf("Failed to mine transaction: %v", err)
        return false, "", err
    }

    log.Printf("Transaction mined: %s", receipt.TxHash.Hex())

    eventSignature := []byte("AccessEvaluated(uint,bool)")
    eventSigHash := crypto.Keccak256Hash(eventSignature)
    accessGranted := false

    // var policy Policy.PolicyContractPolicy
    for _, vLog := range receipt.Logs {
        log.Printf("Processing log with topics: %v", vLog.Topics)
        if len(vLog.Topics) == 0 || vLog.Topics[0] != eventSigHash {
            log.Printf("No matching event signature found in log")
            continue
        }

        // Parse the event
        log.Println("Matching event signature found, attempting to parse event")
        var event struct {
            DroneId    *big.Int
            AccessGranted bool
        }
        err := h.parseAccessDecision(vLog, &event)
        if err != nil {
            log.Printf("Failed to parse AccessDecision event: %v", err)
            return false, "", fmt.Errorf("failed to parse AccessDecision event: %v", err)
        }

        accessGranted = event.AccessGranted
        // Check if the event matches the given entityId

    }

    log.Printf("EvaluateAccess completed: Granted=%t", accessGranted)
    return accessGranted, tx.Hash().Hex(), nil
}


func (h *PDPHandler) Layer2EvaluateAccess(entityID uint, _model_type string, _zone int, _startTime string, _endTime string) (bool, string, error) {
    if h.Client == nil {
        log.Printf("Ethereum client is nil")
        return false, "", fmt.Errorf("Ethereum client is not initialized")
    }

    if h.Instance == nil {
        log.Printf("PDP instance is nil")
        return false, "", fmt.Errorf("PDP instance is not initialized")
    }
    
    bigID := big.NewInt(int64(entityID))
    log.Printf("Starting EvaluateAccess for entityID: %d", entityID)
    bigZone := big.NewInt((int64(_zone)))

    tx, err := h.Instance.Layer2EvaluateAccess(h.Auth, bigID, _model_type, bigZone, _startTime, _endTime)
    if err != nil {
        log.Printf("Failed to call EvaluateAccess: %v", err)
        return false, "", err
    }

    log.Printf("Transaction sent: %s", tx.Hash().Hex())

    log.Println("Waiting for transaction to be mined...")

    ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    defer cancel()
    
    receipt, err := bind.WaitMined(ctx, h.Client, tx)
    if err != nil {
        log.Printf("Failed to mine transaction: %v", err)
        return false, "", err
    }

    log.Printf("Transaction mined: %s", receipt.TxHash.Hex())

    eventSignature := []byte("AccessEvaluated(uint,bool)")
    eventSigHash := crypto.Keccak256Hash(eventSignature)
    accessGranted := false

    // var policy Policy.PolicyContractPolicy
    for _, vLog := range receipt.Logs {
        log.Printf("Processing log with topics: %v", vLog.Topics)
        if len(vLog.Topics) == 0 || vLog.Topics[0] != eventSigHash {
            log.Printf("No matching event signature found in log")
            continue
        }

        // Parse the event
        log.Println("Matching event signature found, attempting to parse event")
        var event struct {
            DroneId    *big.Int
            AccessGranted bool
        }
        err := h.parseAccessDecision(vLog, &event)
        if err != nil {
            log.Printf("Failed to parse AccessDecision event: %v", err)
            return false, "", fmt.Errorf("failed to parse AccessDecision event: %v", err)
        }

        accessGranted = event.AccessGranted
        // Check if the event matches the given entityId

    }

    log.Printf("EvaluateAccess completed: Granted=%t", accessGranted)
    return accessGranted, tx.Hash().Hex(), nil
}

func (h *PDPHandler) Layer3EvaluateAccess(entityID uint, _model_type string, _zone int) (bool, string, error) {
    if h.Client == nil {
        log.Printf("Ethereum client is nil")
        return false, "", fmt.Errorf("Ethereum client is not initialized")
    }

    if h.Instance == nil {
        log.Printf("PDP instance is nil")
        return false, "", fmt.Errorf("PDP instance is not initialized")
    }
    
    bigID := big.NewInt(int64(entityID))
    log.Printf("Starting EvaluateAccess for entityID: %d", entityID)
    bigZone := big.NewInt(int64(_zone))
    tx, err := h.Instance.Layer3EvaluateAccess(h.Auth, bigID, _model_type, bigZone)
    if err != nil {
        log.Printf("Failed to call EvaluateAccess: %v", err)
        return false, "", err
    }

    log.Printf("Transaction sent: %s", tx.Hash().Hex())

    log.Println("Waiting for transaction to be mined...")

    ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    defer cancel()
    
    receipt, err := bind.WaitMined(ctx, h.Client, tx)
    if err != nil {
        log.Printf("Failed to mine transaction: %v", err)
        return false, "", err
    }

    log.Printf("Transaction mined: %s", receipt.TxHash.Hex())

    eventSignature := []byte("AccessEvaluated(uint,bool)")
    eventSigHash := crypto.Keccak256Hash(eventSignature)
    accessGranted := false

    // var policy Policy.PolicyContractPolicy
    for _, vLog := range receipt.Logs {
        log.Printf("Processing log with topics: %v", vLog.Topics)
        if len(vLog.Topics) == 0 || vLog.Topics[0] != eventSigHash {
            log.Printf("No matching event signature found in log")
            continue
        }

        // Parse the event
        log.Println("Matching event signature found, attempting to parse event")
        var event struct {
            DroneId    *big.Int
            AccessGranted bool
        }
        err := h.parseAccessDecision(vLog, &event)
        if err != nil {
            log.Printf("Failed to parse AccessDecision event: %v", err)
            return false, "", fmt.Errorf("failed to parse AccessDecision event: %v", err)
        }

        accessGranted = event.AccessGranted
        // Check if the event matches the given entityId

    }

    log.Printf("EvaluateAccess completed: Granted=%t", accessGranted)
    return accessGranted, tx.Hash().Hex(), nil
}

func (h *PDPHandler) Layer4EvaluateAccess(entityID uint) (bool, string, error) {
    if h.Client == nil {
        log.Printf("Ethereum client is nil")
        return false, "", fmt.Errorf("Ethereum client is not initialized")
    }

    if h.Instance == nil {
        log.Printf("PDP instance is nil")
        return false, "", fmt.Errorf("PDP instance is not initialized")
    }
    
    bigID := big.NewInt(int64(entityID))
    log.Printf("Starting EvaluateAccess for entityID: %d", entityID)
    
    tx, err := h.Instance.Layer4EvaluateAccess(h.Auth, bigID)
    if err != nil {
        log.Printf("Failed to call EvaluateAccess: %v", err)
        return false, "", err
    }

    log.Printf("Transaction sent: %s", tx.Hash().Hex())

    log.Println("Waiting for transaction to be mined...")

    ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    defer cancel()
    
    receipt, err := bind.WaitMined(ctx, h.Client, tx)
    if err != nil {
        log.Printf("Failed to mine transaction: %v", err)
        return false, "", err
    }

    log.Printf("Transaction mined: %s", receipt.TxHash.Hex())

    eventSignature := []byte("AccessEvaluated(uint,bool)")
    eventSigHash := crypto.Keccak256Hash(eventSignature)
    accessGranted := false

    // var policy Policy.PolicyContractPolicy
    for _, vLog := range receipt.Logs {
        log.Printf("Processing log with topics: %v", vLog.Topics)
        if len(vLog.Topics) == 0 || vLog.Topics[0] != eventSigHash {
            log.Printf("No matching event signature found in log")
            continue
        }

        // Parse the event
        log.Println("Matching event signature found, attempting to parse event")
        var event struct {
            DroneId    *big.Int
            AccessGranted bool
        }
        err := h.parseAccessDecision(vLog, &event)
        if err != nil {
            log.Printf("Failed to parse AccessDecision event: %v", err)
            return false, "", fmt.Errorf("failed to parse AccessDecision event: %v", err)
        }

        accessGranted = event.AccessGranted
        // Check if the event matches the given entityId

    }

    log.Printf("EvaluateAccess completed: Granted=%t", accessGranted)
    return accessGranted, tx.Hash().Hex(), nil
}

func (h *PDPHandler) parseAccessDecision(logEntry *types.Log, event *struct {
    DroneId *big.Int
    AccessGranted bool
}) error {
    abiDefinition := `[{
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "entityId",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "bool",
        "name": "accessGranted",
        "type": "bool"
      }
    ],
    "name": "AccessEvaluated",
    "type": "event"
    }]`
    parsedABI, err := abi.JSON(strings.NewReader(abiDefinition))
    if err != nil {
        return err
    }

    err = parsedABI.UnpackIntoInterface(event, "AccessEvaluated", logEntry.Data)
    if err != nil {
        return err
    }

    if len(logEntry.Topics) > 1 {
        event.DroneId = new(big.Int).SetBytes(logEntry.Topics[1].Bytes())
    } else {
        return fmt.Errorf("insufficient topics in log entry")
    }

    return nil
}
