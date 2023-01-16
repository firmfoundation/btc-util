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

	var dhd *pb.Response
	dhd = hdwallet.DecodeHDWallet(hd.Mnemonic)
	fmt.Printf("restored address : %v", dhd.Address)
}
