package handler

import (
	"context"
	"drone-backend/smart-contracts/Attribute"
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

type AttributeHandler struct {
    Instance *Attribute.Attribute
    Auth     *bind.TransactOpts
	Client	*ethclient.Client
}

type AttributeUpdatedEvent struct {
    Id        *big.Int
    Name   string
    Value []*big.Int
}

func NewAttributeHandler(instance *Attribute.Attribute, auth *bind.TransactOpts, client *ethclient.Client) *AttributeHandler {
    return &AttributeHandler{
        Instance: instance,
        Auth:     auth,
		Client:	client,
    }
}

func (h *AttributeHandler) GetAttribute(id uint) (Attribute.AttributeContractAttribute, error) {
    bigID := big.NewInt(int64(id))
    return h.Instance.GetAttribute(&bind.CallOpts{}, bigID)
}

func (h *AttributeHandler) CreateAttribute(name string, values []uint) (string, error) {
	bigIntValues := make([]*big.Int, len(values))
	for i, val := range values {
		bigIntValues[i] = big.NewInt(int64(val))
	}

	tx, err := h.Instance.CreateAttribute(h.Auth, name, bigIntValues)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (h *AttributeHandler) RemoveAttribute(id uint) (string, error) {
	bigID := big.NewInt(int64(id))
	tx, err := h.Instance.DeleteAttribute(h.Auth, bigID)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (h *AttributeHandler) UpdateAttribute(id uint, name string, values []uint) (Attribute.AttributeContractAttribute, error) {
	bigID := big.NewInt(int64(id))
	bigIntValues := make([]*big.Int, len(values))
	for i, val := range values {
		bigIntValues[i] = big.NewInt(int64(val))
	}

	tx, err := h.Instance.UpdateAttribute(h.Auth, bigID, name, bigIntValues)
	if err != nil {
		return Attribute.AttributeContractAttribute{}, err
	}
	log.Printf("Transaction sent: %s", tx.Hash().Hex())

    log.Println("Waiting for transaction to be mined...")

    ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    defer cancel()
    
    receipt, err := bind.WaitMined(ctx, h.Client, tx)
    if err != nil {
        log.Printf("Failed to mine transaction: %v", err)
        return Attribute.AttributeContractAttribute{}, err
    }

    var attribute Attribute.AttributeContractAttribute
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
        var tempAttribute AttributeUpdatedEvent
        err := h.parseAttributeUpdated(vLog, &tempAttribute)
        if err != nil {
            log.Printf("Failed to parse AccessDecision event: %v", err)
            return Attribute.AttributeContractAttribute{}, fmt.Errorf("failed to parse AccessDecision event: %v", err)
        }

        attribute = Attribute.AttributeContractAttribute{
            Id: tempAttribute.Id,
			Name: tempAttribute.Name,
			Value: tempAttribute.Value,
        }
    }

    return attribute, nil
}

func (h *AttributeHandler) GetAttributes() ([]Attribute.AttributeContractAttribute, error) {
	var attrs []Attribute.AttributeContractAttribute

	attrs, err := h.Instance.GetAttributes(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return attrs, nil
}

func (h *AttributeHandler) GetAttributesByName(name string) ([]Attribute.AttributeContractAttribute, error) {
	var attrs []Attribute.AttributeContractAttribute
	
	attrs, err := h.Instance.GetAttributesByName(&bind.CallOpts{}, name)
	if err != nil {
		return nil, err
	}
	return attrs, nil
}

func (h *AttributeHandler) parseAttributeUpdated(logEntry *types.Log, attribute *AttributeUpdatedEvent) error {
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
        "name": "name",
        "type": "string"
      },
      {
        "indexed": false,
        "internalType": "uint256[]",
        "name": "value",
        "type": "uint256[]"
      }
    ],
    "name": "AttributeUpdated",
    "type": "event"
  }]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDefinition))
	if err != nil {
		log.Printf("Error parsing ABI: %v", err)
		return err
	}

	err = parsedABI.UnpackIntoInterface(attribute, "AttributeUpdated", logEntry.Data)
	if err != nil {
		log.Printf("Error unpacking into interface: %v", err)
		return err
	}

	if len(logEntry.Topics) > 1 {
		attribute.Id = new(big.Int).SetBytes(logEntry.Topics[1].Bytes())
	} else {
		return fmt.Errorf("insufficient topics in log entry")
	}

	return nil
}
