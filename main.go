package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"drone-backend/internal/database"
	"drone-backend/internal/handlers/api"
	handlers "drone-backend/internal/handlers/off-chain"
	handler "drone-backend/internal/handlers/on-chain"
	"drone-backend/smart-contracts/Attribute"
	"drone-backend/smart-contracts/Drone"
	"drone-backend/smart-contracts/PDP"
	"drone-backend/smart-contracts/Policy"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := database.Initialize()
	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}

	handlers.InitDB(database.DB)
	
	r := mux.NewRouter()
	// r.HandleFunc("/ping", handlers.PingHandler).Methods(http.MethodGet)

	r.HandleFunc("/off-chain/createDrone", handlers.CreateDroneHandler).Methods(http.MethodPost)
	r.HandleFunc("/off-chain/getDrones", handlers.GetDronesHandler).Methods(http.MethodGet)
	r.HandleFunc("/off-chain/getDrone/{id:[0-9]+}", handlers.GetDroneHandler).Methods(http.MethodGet)
	r.HandleFunc("/off-chain/updateDrone/{id:[0-9]+}", handlers.UpdateDroneHandler).Methods(http.MethodPut)
	r.HandleFunc("/off-chain/removeDrone/{id:[0-9]+}", handlers.DeleteDroneHandler).Methods(http.MethodDelete)
	r.HandleFunc("/off-chain/getDronesByZone/{zone}", handlers.GetDroneByZoneHandler).Methods(http.MethodGet)

	r.HandleFunc("/off-chain/createAttribute", handlers.CreateAttributeHandler).Methods(http.MethodPost)
	r.HandleFunc("/off-chain/getAttributes", handlers.GetAttributesHandler).Methods(http.MethodGet)
	r.HandleFunc("/off-chain/getAttribute/{id:[0-9]+}", handlers.GetAttributeHandler).Methods(http.MethodGet)
	r.HandleFunc("/off-chain/updateAttribute/{id:[0-9]+}", handlers.UpdateAttributeHandler).Methods(http.MethodPut)
	r.HandleFunc("/off-chain/removeAttribute/{id:[0-9]+}", handlers.DeleteAttributeHandler).Methods(http.MethodDelete)
	r.HandleFunc("/off-chain/getAttributeByName/{name}", handlers.GetAttributeByNameHandler).Methods(http.MethodGet)

	r.HandleFunc("/off-chain/createPolicy", handlers.CreatePolicyHandler).Methods(http.MethodPost)
	r.HandleFunc("/off-chain/getPolicies", handlers.GetPoliciesHandler).Methods(http.MethodGet)
	r.HandleFunc("/off-chain/getPolicy/{id:[0-9]+}", handlers.GetPolicyHandler).Methods(http.MethodGet)
	r.HandleFunc("/off-chain/updatePolicy/{id:[0-9]+}", handlers.UpdatePolicyHandler).Methods(http.MethodPut)
	r.HandleFunc("/off-chain/removePolicy/{id:[0-9]+}", handlers.DeletePolicyHandler).Methods(http.MethodDelete)

	
	env_err := godotenv.Load()
	if env_err != nil {
		log.Fatalf("Error loading .env file: %v", env_err)
	}
	
	clientURL := os.Getenv("clientURL")
	// password := os.Getenv("password")
	attributeContractAddress := os.Getenv("attributeContractAddress")
	policyContractAddress := os.Getenv("policyContractAddress")
	droneContractAddress := os.Getenv("droneContractAddress")
	pdpContractAddress := os.Getenv("pdpContractAddress")
	keystoreFile := "./wallet.json"
	
	// clientURL := "https://json-rpc.evm.testnet.iotaledger.net"
	password := "Wahaha123!@#";
	// attributeContractAddress := "0xBE090359fCa730387E674741eF912E3be05a2693"
	// policyContractAddress := "0xC83eED3676fd6422bE8B4E68dDecD1C3763C26A3"
	// droneContractAddress := "0x971c1979B399c7B6798A39066684f36C41724221"
	// pdpContractAddress := "0x5651987cd71bA9f5FF2d5d127a4C4116af53567d"
	
	client, err := ethclient.Dial(clientURL)
    if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }
	
	defer client.Close()
	
    keyJSON, err := ioutil.ReadFile(keystoreFile)
    if err != nil {
		log.Fatalf("Failed to read the keyfile: %v", err)
    }
	
    key, err := keystore.DecryptKey(keyJSON, password)
    if err != nil {
		log.Fatalf("Failed to decrypt the key: %v", err)
    }
	
    auth := bind.NewKeyedTransactor(key.PrivateKey)
	
    attributeInstance, err := Attribute.NewAttribute(common.HexToAddress(attributeContractAddress), client)
    if err != nil {
		log.Fatalf("Failed to instantiate AttributeContract: %v", err)
    }
	
    policyInstance, err := Policy.NewPolicy(common.HexToAddress(policyContractAddress), client)
    if err != nil {
		log.Fatalf("Failed to instantiate PolicyContract: %v", err)
    }
	
    droneInstance, err := Drone.NewDrone(common.HexToAddress(droneContractAddress), client)
    if err != nil {
		log.Fatalf("Failed to instantiate DroneContract: %v", err)
    }
	
	pdpInstance, err := PDP.NewPDP(common.HexToAddress(pdpContractAddress), client)
    if err != nil {
		log.Fatalf("Failed to instantiate PDPContract: %v", err)
    }
	
    attributeHandler := handler.NewAttributeHandler(attributeInstance, auth, client)
    policyHandler := handler.NewPolicyHandler(policyInstance, auth, client)
    droneHandler := handler.NewDroneHandler(droneInstance, auth, client)
	pdpHandler := handler.NewPDPHandler(pdpInstance, auth, client)

	attributeAPI := api.NewAttributeAPI(attributeHandler)
    policyAPI := api.NewPolicyAPI(policyHandler)
    droneAPI := api.NewDroneAPI(droneHandler)
	pdpAPI := api.NewAccessAPI(pdpHandler, droneAPI, policyAPI)
	
	r.HandleFunc("/on-chain/getAttribute/{id:[0-9]+}", attributeAPI.GetAttributeHandler)
	r.HandleFunc("/on-chain/createAttribute", attributeAPI.CreateAttributeHandler)
	r.HandleFunc("/on-chain/removeAttribute/{id:[0-9]+}", attributeAPI.RemoveAttributeHandler)
	r.HandleFunc("/on-chain/updateAttribute/{id:[0-9]+}", attributeAPI.UpdateAttributeHandler)
	r.HandleFunc("/on-chain/getAttributes", attributeAPI.GetAttributesHandler)
	r.HandleFunc("/on-chain/getAttributeByName/{name}", attributeAPI.GetAttributeByNameHandler)
	
	r.HandleFunc("/on-chain/getPolicies", policyAPI.GetPoliciesHandler)
	r.HandleFunc("/on-chain/createPolicy", policyAPI.CreatePolicyHandler)
	r.HandleFunc("/on-chain/updatePolicy/{id:[0-9]+}", policyAPI.UpdatePolicyHandler)
	r.HandleFunc("/on-chain/removePolicy/{id:[0-9]+}", policyAPI.RemovePolicyHandler)
	r.HandleFunc("/on-chain/getPolicyByZone/{zone}", policyAPI.GetPolicyByZoneHandler)
	
	r.HandleFunc("/on-chain/getDrones", droneAPI.GetDronesHandler)
	r.HandleFunc("/on-chain/getDrone/{id:[0-9]+}", droneAPI.GetDroneHandler)
	r.HandleFunc("/on-chain/createDrone", droneAPI.CreateDroneHandler)
	r.HandleFunc("/on-chain/updateDrone/{id:[0-9]+}", droneAPI.UpdateDroneHandler)
	r.HandleFunc("/on-chain/removeDrone/{id:[0-9]+}", droneAPI.RemoveDroneHandler)
	r.HandleFunc("/on-chain/getDronesByZone/{zone}", droneAPI.GetDronesByZoneHandler)
	
	r.HandleFunc("/layer-1/accessRequest", pdpAPI.Layer1AccessRequestHandler).Methods(http.MethodPost)
	r.HandleFunc("/layer-2/accessRequest", pdpAPI.Layer2AccessRequestHandler)
	r.HandleFunc("/layer-3/accessRequest", pdpAPI.Layer3AccessRequestHandler)
	r.HandleFunc("/layer-4/accessRequest", pdpAPI.Layer4AccessRequestHandler)
	
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
		AllowCredentials: true,
	})

	// srv := server.NewServer()
	handler := c.Handler((r));
	log.Fatal(http.ListenAndServe(":8080", handler))
}