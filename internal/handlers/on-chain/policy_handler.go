package handler

import (
	"context"
	"drone-backend/smart-contracts/Policy"
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

type PolicyHandler struct {
	Instance *Policy.Policy
	Auth     *bind.TransactOpts
    Client *ethclient.Client
}

type PolicyCreatedEvent struct {
    Id        *big.Int
    Zone      *big.Int
    StartTime   string
    EndTime string
}

func NewPolicyHandler(instance *Policy.Policy, auth *bind.TransactOpts, client *ethclient.Client) *PolicyHandler {
	return &PolicyHandler{
		Instance: instance,
		Auth: auth,
        Client: client,
	}
}

func (h *PolicyHandler) GetPolicies() ([]Policy.PolicyContractPolicy, error) {
	policies, err := h.Instance.GetPolicies(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return policies, nil
}

func (h *PolicyHandler) CreatePolicy(zone int, startTime, endTime string) (Policy.PolicyContractPolicy, error) {
    tx, err := h.Instance.CreatePolicy(h.Auth, big.NewInt(int64(zone)), startTime, endTime)
    if err != nil {
        return Policy.PolicyContractPolicy{}, err
    }
    log.Printf("Transaction sent: %s", tx.Hash().Hex())

    log.Println("Waiting for transaction to be mined...")

    ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    defer cancel()
    
    receipt, err := bind.WaitMined(ctx, h.Client, tx)
    if err != nil {
        log.Printf("Failed to mine transaction: %v", err)
        return Policy.PolicyContractPolicy{}, err
    }

    var policy Policy.PolicyContractPolicy
    log.Printf("Transaction mined: %s", receipt.TxHash.Hex())

    eventSignature := []byte("PolicyCreated(uint256,int256,string,string)")
    eventSigHash := crypto.Keccak256Hash(eventSignature)

    for _, vLog := range receipt.Logs {
        log.Printf("Processing log with topics: %v", vLog.Topics)
        if len(vLog.Topics) == 0 || vLog.Topics[0] != eventSigHash {
            log.Printf("No matching event signature found in log")
            continue
        }

        // Parse the event
        log.Println("Matching event signature found, attempting to parse event")
        var tempPolicy PolicyCreatedEvent
        err := h.parsePolicyCreated(vLog, &tempPolicy)
        if err != nil {
            log.Printf("Failed to parse AccessDecision event: %v", err)
            return Policy.PolicyContractPolicy{}, fmt.Errorf("failed to parse AccessDecision event: %v", err)
        }

        policy = Policy.PolicyContractPolicy{
            Id: tempPolicy.Id,
            Zone: tempPolicy.Zone,
            StartTime: tempPolicy.StartTime,
            EndTime: tempPolicy.EndTime,
        }
    }

    return policy, nil

}

func (h *PolicyHandler) UpdatePolicy(id uint, zone int, startTime, endTime string) (Policy.PolicyContractPolicy, error) {
    bigID := big.NewInt(int64(id))
    tx, err := h.Instance.UpdatePolicy(h.Auth, bigID, big.NewInt(int64(zone)), startTime, endTime)
    if err != nil {
        return Policy.PolicyContractPolicy{}, err
    }
    log.Printf("Transaction sent: %s", tx.Hash().Hex())

    log.Println("Waiting for transaction to be mined...")

    ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    defer cancel()
    
    receipt, err := bind.WaitMined(ctx, h.Client, tx)
    if err != nil {
        log.Printf("Failed to mine transaction: %v", err)
        return Policy.PolicyContractPolicy{}, err
    }

    var policy Policy.PolicyContractPolicy
    log.Printf("Transaction mined: %s", receipt.TxHash.Hex())

    eventSignature := []byte("PolicyUpdated(uint256,int256,string,string)")
    eventSigHash := crypto.Keccak256Hash(eventSignature)

    for _, vLog := range receipt.Logs {
        log.Printf("Processing log with topics: %v", vLog.Topics)
        if len(vLog.Topics) == 0 || vLog.Topics[0] != eventSigHash {
            log.Printf("No matching event signature found in log")
            continue
        }

        // Parse the event
        log.Println("Matching event signature found, attempting to parse event")
        var tempPolicy PolicyCreatedEvent
        err := h.parsePolicyUpdated(vLog, &tempPolicy)
        if err != nil {
            log.Printf("Failed to parse AccessDecision event: %v", err)
            return Policy.PolicyContractPolicy{}, fmt.Errorf("failed to parse AccessDecision event: %v", err)
        }

        policy = Policy.PolicyContractPolicy{
            Id: tempPolicy.Id,
            Zone: tempPolicy.Zone,
            StartTime: tempPolicy.StartTime,
            EndTime: tempPolicy.EndTime,
        }
    }

    return policy, nil
}

func (h *PolicyHandler) RemovePolicy(id uint) (string, error) {
    bigID := big.NewInt(int64(id))
    tx, err := h.Instance.DeletePolicy(h.Auth, bigID)
    if err != nil {
        return "", err
    }
    return tx.Hash().Hex(), nil
}

func (h *PolicyHandler) GetPolicyByZone(zone int) (Policy.PolicyContractPolicy, error) {
    bigZone := big.NewInt(int64(zone))
    
    if h.Client == nil {
        log.Printf("Ethereum client is nil")
        return Policy.PolicyContractPolicy{}, fmt.Errorf("Ethereum client is not initialized")
    }

    policy, err := h.Instance.GetPolicyByZone(&bind.CallOpts{}, bigZone)
    if err != nil {
        return Policy.PolicyContractPolicy{}, err
    }

    return policy, nil
}

func (h *PolicyHandler) parseLog(logEntry *types.Log, event *struct {
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

func (h *PolicyHandler) parsePolicyCreated(logEntry *types.Log, policy *PolicyCreatedEvent) error {
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
        "internalType": "int256",
        "name": "zone",
        "type": "int256"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "startTime",
        "type": "string"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "endTime",
        "type": "string"
      }
    ],
    "name": "PolicyCreated",
    "type": "event"
  }]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDefinition))
	if err != nil {
		log.Printf("Error parsing ABI: %v", err)
		return err
	}

	err = parsedABI.UnpackIntoInterface(policy, "PolicyCreated", logEntry.Data)
	if err != nil {
		log.Printf("Error unpacking into interface: %v", err)
		return err
	}

	if len(logEntry.Topics) > 1 {
		policy.Id = new(big.Int).SetBytes(logEntry.Topics[1].Bytes())
	} else {
		return fmt.Errorf("insufficient topics in log entry")
	}

	return nil
}

func (h *PolicyHandler) parsePolicyUpdated(logEntry *types.Log, policy *PolicyCreatedEvent) error {
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
        "internalType": "int256",
        "name": "zone",
        "type": "int256"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "startTime",
        "type": "string"
      },
      {
        "indexed": false,
        "internalType": "string",
        "name": "endTime",
        "type": "string"
      }
    ],
    "name": "PolicyUpdated",
    "type": "event"
  }]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDefinition))
	if err != nil {
		log.Printf("Error parsing ABI: %v", err)
		return err
	}

	err = parsedABI.UnpackIntoInterface(policy, "PolicyUpdated", logEntry.Data)
	if err != nil {
		log.Printf("Error unpacking into interface: %v", err)
		return err
	}

	if len(logEntry.Topics) > 1 {
		policy.Id = new(big.Int).SetBytes(logEntry.Topics[1].Bytes())
	} else {
		return fmt.Errorf("insufficient topics in log entry")
	}

	return nil
}