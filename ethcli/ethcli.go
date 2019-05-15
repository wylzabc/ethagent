package ethcli

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"wylzabc/ethagent/contract"
)

var ServerAddr = "http://192.168.19.128:8101"
var ContractAddr = "0xe2a61E8118E9858E12561AE45971066Cf69064a8"
var PrivateKeyHex = "a80295784d2917e82ecf09ea6f2fce0305376b2607247307eae47ce59364f236"

var rpcClient *rpc.Client
var privateKey *ecdsa.PrivateKey
var auth *bind.TransactOpts
var contractInHospitalData *contract.InHospitalData

func subscribeEvent() error {
	client, err := ethclient.Dial("ws://192.168.19.128:30002/ws")
	if err != nil {
		return err
	}
	contractAddress := common.HexToAddress(ContractAddr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return err
	}

	abifile, err := os.Open("./contract/inhostpitaldata.abi")
	if err != nil {
		return err
	}
	defer abifile.Close()

	contractABI, err := abi.JSON(abifile)
	go func() {
		for {
			select {
			case err := <-sub.Err():
				fmt.Println(err)
			case vlog := <-logs:
				event := struct {
					Personid   string
					Timestamp  string
					Hospitalid string
					Devid      string
				}{}
				err := contractABI.Unpack(&event, "dataPushed", vlog.Data)
				if err != nil {
					fmt.Println("contractABI.Unpack error")
				} else {
					fmt.Println(event.Personid, ",", event.Timestamp, ",", event.Hospitalid, ",", event.Devid)
				}
			}
		}
	}()
	return nil
}

func Init() error {
	var err error

	fmt.Println("ethclient init ok")
	rpcClient, err = rpc.Dial(ServerAddr)
	if err != nil {
		fmt.Println("rpc.Dial err:", err)
		return err
	}

	privateKey, err = crypto.HexToECDSA(PrivateKeyHex)
	if err != nil {
		fmt.Println("crypto.HexToECDSA err:", err)
		return err
	}
	auth = bind.NewKeyedTransactor(privateKey)

	contractInHospitalData, err = contract.NewInHospitalData(common.HexToAddress(ContractAddr), ethclient.NewClient(rpcClient))
	if err != nil {
		fmt.Println("contract.NewInHospitalData err:", err)
		return err
	}

	err = subscribeEvent()
	if err != nil {
		fmt.Println("subscribe event error:", err)
	}

	return nil
}

func PutData(personid, timestamp, hospitalid, devid, personveindata, personpicdata, additionalinfo string) error {
    fmt.Println(personid, timestamp, hospitalid, devid, personveindata, personpicdata, additionalinfo)
	_, err := contractInHospitalData.InHospitalDataTransactor.PutData(auth, personid, timestamp, hospitalid, devid, personid, personveindata, personpicdata, additionalinfo)
	if err != nil {
		fmt.Println("contract.NewInHospitalData err:", err)
		return err
	}
	return nil
}

func GetDataByPersonid(personid string) (string, error) {
	data, err := contractInHospitalData.InHospitalDataCaller.GetDataByPersonId(nil, personid)
	return data, err
}

func GetDataByPersonidAndDate(personid, date string) (string, error) {
	data, err := contractInHospitalData.InHospitalDataCaller.GetDataByPersonidAndDate(nil, personid, date)
	return data, err
}
