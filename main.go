package main

import (
	"fmt"

	pb "github.com/LuisAcerv/btchdwallet/proto/btchdwallet"
	crypt "github.com/firmfoundation/btcutil/pkg/crypt"
	hdwallet "github.com/firmfoundation/btcutil/pkg/hdwallet"
)

func main() {
	fmt.Println("Lord Jesus thank you , you granted my request")

	crypt.CreateRandSeed()

	var hd *pb.Response
	hd = hdwallet.CreateHDWallet()
	fmt.Printf("address : %v\n", hd.Address)

	var info *pb.Response
	//get any address for testing from https://www.blockchain.com/explorer/
	var tmpAddress string = "35PCtVR7U9FF5atdCSfrztLXfQa3Mm17wo"
	info = hdwallet.GetBalance(tmpAddress)

	btc := float64(info.Balance) * 0.00000001
	fmt.Println("btcoin balance ", btc)

	// var dhd *pb.Response
	// dhd = hdwallet.DecodeHDWallet(hd.Mnemonic)
	// fmt.Printf("restored address : %v", dhd.Address)
}
