package hdwallet

import (
	"fmt"

	pb "github.com/LuisAcerv/btchdwallet/proto/btchdwallet"
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
