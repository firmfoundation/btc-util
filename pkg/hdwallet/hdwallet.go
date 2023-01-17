package hdwallet

import (
	"fmt"

	pb "github.com/LuisAcerv/btchdwallet/proto/btchdwallet"
	"github.com/blockcypher/gobcy"
	"github.com/brianium/mnemonic"
	crypt "github.com/firmfoundation/btcutil/pkg/crypt"
	"github.com/wemeetagain/go-hdwallet"
)

func CreateHDWallet() *pb.Response {
	// Create random 256 bit seed
	seed := crypt.CreateRandSeed()

	// create mnemonic
	mnemonic, _ := mnemonic.New([]byte(seed), mnemonic.English)

	fmt.Println(mnemonic.Sentence())
	// Create a master private key

	//mnemonic sentence for testing
	//var tmp string = "give balance snap idea globe option network shoot cram seek escape shrug"

	master_privatekey := hdwallet.MasterKey([]byte(mnemonic.Sentence()))

	// Generate public key from private key
	master_publickey := master_privatekey.Pub()

	//Get address
	address := master_publickey.Address()

	return &pb.Response{
		Address:  address,
		PubKey:   master_publickey.String(),
		PrivKey:  master_privatekey.Address(),
		Mnemonic: mnemonic.Sentence(),
	}
}

func DecodeHDWallet(mnemonic_sentence string) *pb.Response {
	// get private key from mnemonic sentence
	master_prv := hdwallet.MasterKey([]byte(mnemonic_sentence))

	//get public key from private key
	master_pub := master_prv.Pub()

	//get address
	address := master_pub.Address()

	return &pb.Response{
		Address: address,
		PubKey:  master_pub.String(),
		PrivKey: master_prv.String(),
	}
}

func GetBalance(address string) *pb.Response {
	btc := gobcy.API{
		Token: "",
		Coin:  "btc",
		Chain: "main",
	}
	addr, err := btc.GetAddrBal(address, nil)
	if err != nil {
		fmt.Println(err)
	}

	/*
	 returned values are in satoshi
	 1 satoshi = 0.00000001 bitcoin
	 so at the end mutiple the value with 0.00000001 to get the
	 bitcoin amount
	*/
	balance := addr.Balance
	totalReceived := addr.TotalReceived
	totalSent := addr.TotalSent
	unconfirmedBalance := addr.UnconfirmedBalance

	return &pb.Response{
		Address:            address,
		Balance:            balance.Int64(),
		TotalReceived:      totalReceived.Int64(),
		TotalSent:          totalSent.Int64(),
		UnconfirmedBalance: unconfirmedBalance.Int64(),
	}
}
