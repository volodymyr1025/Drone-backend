package handler

import (
	"context"
	"drone-backend/smart-contracts/Drone" // Replace with your actual module name
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

type DroneHandler struct {
    Instance *Drone.Drone
    Auth     *bind.TransactOpts
    Client  *ethclient.Client
}

type DroneCreatedEvent struct {
    Id        *big.Int
    ModelType string
    Zone      *big.Int
}

func NewDroneHandler(instance *Drone.Drone, auth *bind.TransactOpts, client *ethclient.Client) *DroneHandler {
    return &DroneHandler{
        Instance: instance,
        Auth:     auth,
        Client: client,
    }
}

func (h *DroneHandler) GetDrones() ([]Drone.DroneContractDrone, error) {
    return h.Instance.GetDrones(&bind.CallOpts{})
}

func (h *DroneHandler) GetDronesByZone(zone int) ([]Drone.DroneContractDrone, error) {
	drones, err := h.Instance.GetDronesByZone(&bind.CallOpts{}, big.NewInt(int64(zone)))
    if err != nil {
        return []Drone.DroneContractDrone{}, err
    }
    return drones, nil
}
func (h *DroneHandler) CreateDrone(model string, zone int) (Drone.DroneContractDrone, error) {
    tx, err := h.Instance.CreateDrone(h.Auth, model, big.NewInt(int64(zone)))
    if err != nil {
        return Drone.DroneContractDrone{}, err
    }
    log.Printf("Transaction sent: %s", tx.Hash().Hex())

    log.Println("Waiting for transaction to be mined...")

    ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    defer cancel()
    
    receipt, err := bind.WaitMined(ctx, h.Client, tx)
    if err != nil {
        log.Printf("Failed to mine transaction: %v", err)
        return Drone.DroneContractDrone{}, err
    }

    var drone Drone.DroneContractDrone
    log.Printf("Transaction mined: %s", receipt.TxHash.Hex())

    eventSignature := []byte("DroneCreated(uint256,string,int256)")
    eventSigHash := crypto.Keccak256Hash(eventSignature)

    for _, vLog := range receipt.Logs {
        log.Printf("Processing log with topics: %v", vLog.Topics)
        if len(vLog.Topics) == 0 || vLog.Topics[0] != eventSigHash {
            log.Printf("No matching event signature found in log")
            continue
        }

        // Parse the event
        log.Println("Matching event signature found, attempting to parse event")
        var tempDrone DroneCreatedEvent
        err := h.parseDroneCreated(vLog, &tempDrone)
        if err != nil {
            log.Printf("Failed to parse AccessDecision event: %v", err)
            return Drone.DroneContractDrone{}, fmt.Errorf("failed to parse AccessDecision event: %v", err)
        }

        drone = Drone.DroneContractDrone{
            Id: tempDrone.Id,
            ModelType: tempDrone.ModelType,
            Zone: tempDrone.Zone,
        }
    }

    return drone, nil
}

func (h *DroneHandler) UpdateDrone(id uint, model string, zone int) (Drone.DroneContractDrone, error) {
    bigID := big.NewInt(int64(id))
    tx, err := h.Instance.UpdateDrone(h.Auth, bigID, model, big.NewInt(int64(zone)))
    if err != nil {
        return Drone.DroneContractDrone{}, err
    }
    log.Printf("Transaction sent: %s", tx.Hash().Hex())

    log.Println("Waiting for transaction to be mined...")

    ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    defer cancel()
    
    receipt, err := bind.WaitMined(ctx, h.Client, tx)
    if err != nil {
        log.Printf("Failed to mine transaction: %v", err)
        return Drone.DroneContractDrone{}, err
    }

    var drone Drone.DroneContractDrone
    log.Printf("Transaction mined: %s", receipt.TxHash.Hex())

    eventSignature := []byte("DroneUpdated(uint256,string,int256)")
    eventSigHash := crypto.Keccak256Hash(eventSignature)

    for _, vLog := range receipt.Logs {
        log.Printf("Processing log with topics: %v", vLog.Topics)
        if len(vLog.Topics) == 0 || vLog.Topics[0] != eventSigHash {
            log.Printf("No matching event signature found in log")
            continue
        }

        // Parse the event
        log.Println("Matching event signature found, attempting to parse event")
        var tempDrone DroneCreatedEvent
        err := h.parseDroneCreated(vLog, &tempDrone)
        if err != nil {
            log.Printf("Failed to parse AccessDecision event: %v", err)
            return Drone.DroneContractDrone{}, fmt.Errorf("failed to parse AccessDecision event: %v", err)
        }

        drone = Drone.DroneContractDrone{
            Id: tempDrone.Id,
            ModelType: tempDrone.ModelType,
            Zone: tempDrone.Zone,
        }
    }

    return drone, nil
}

func (h *DroneHandler) RemoveDrone(id uint) (string, error) {
    bigID := big.NewInt(int64(id))
    tx, err := h.Instance.DeleteDrone(h.Auth, bigID)
    if err != nil {
        return "", err
    }
    return tx.Hash().Hex(), nil
}

func (h *DroneHandler) GetDrone(id uint) (Drone.DroneContractDrone, error) {
    bigID := big.NewInt(int64(id))

    if h.Client == nil {
        log.Printf("Ethereum client is nil")
        return Drone.DroneContractDrone{}, fmt.Errorf("Ethereum client is not initialized")
    }

    drone, err := h.Instance.GetDrone(&bind.CallOpts{}, bigID)
    if err != nil {
        return Drone.DroneContractDrone{}, err
    }
    return drone, nil
}

func (h *DroneHandler) parseDroneCreated(logEntry *types.Log, drone *DroneCreatedEvent) error {
	abiDefinition := `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint256","name":"id","type":"uint256"},{"indexed":false,"internalType":"string","name":"model_type","type":"string"},{"indexed":false,"internalType":"int256","name":"zone","type":"int256"}],"name":"DroneCreated","type":"event"}]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDefinition))
	if err != nil {
		log.Printf("Error parsing ABI: %v", err)
		return err
	}

	err = parsedABI.UnpackIntoInterface(drone, "DroneCreated", logEntry.Data)
	if err != nil {
		log.Printf("Error unpacking into interface: %v", err)
		return err
	}

	if len(logEntry.Topics) > 1 {
		drone.Id = new(big.Int).SetBytes(logEntry.Topics[1].Bytes())
	} else {
		return fmt.Errorf("insufficient topics in log entry")
	}

	log.Printf("Unpacked Event: ID=%v, ModelType=%v, Zone=%v", drone.Id, drone.ModelType, drone.Zone)
	return nil
}

func (h *DroneHandler) parseDroneUpdated(logEntry *types.Log, drone *DroneCreatedEvent) error {
	abiDefinition := `[{
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "uint256",
        "name": "id",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "model_type",
        "type": "string"
      },
      {
        "indexed": false,
        "internalType": "int256",
        "name": "zone",
        "type": "int256"
      }
    ],
    "name": "DroneUpdated",
    "type": "event"
  }]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDefinition))
	if err != nil {
		log.Printf("Error parsing ABI: %v", err)
		return err
	}

	err = parsedABI.UnpackIntoInterface(drone, "DroneUpdated", logEntry.Data)
	if err != nil {
		log.Printf("Error unpacking into interface: %v", err)
		return err
	}

	if len(logEntry.Topics) > 1 {
		drone.Id = new(big.Int).SetBytes(logEntry.Topics[1].Bytes())
	} else {
		return fmt.Errorf("insufficient topics in log entry")
	}

	log.Printf("Unpacked Event: ID=%v, ModelType=%v, Zone=%v", drone.Id, drone.ModelType, drone.Zone)
	return nil
}

func (h *DroneHandler) parseDroneAccessed(logEntry *types.Log, event *struct {
    User    *big.Int
    Action string
    Timestamp  string
    Data string
}) error {
    abiDefinition := `[{
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "user",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "action",
        "type": "string"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "timestamp",
        "type": "string"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "data",
        "type": "string"
      }
    ],
    "name": "ActionLogged",
    "type": "event"
  }]`
    parsedABI, err := abi.JSON(strings.NewReader(abiDefinition))
    if err != nil {
        return err
    }

    err = parsedABI.UnpackIntoInterface(event, "ActionLogged", logEntry.Data)
    if err != nil {
        return err
    }

    if len(logEntry.Topics) > 1 {
        event.User = new(big.Int).SetBytes(logEntry.Topics[1].Bytes())
    } else {
        return fmt.Errorf("insufficient topics in log entry")
    }

    return nil
}
