package main

import (
	"fmt"

	crypt "github.com/firmfoundation/btcutil/pkg/crypt"
	hdwallet "github.com/firmfoundation/btcutil/pkg/hdwallet"
)

func main() {
	fmt.Println("Lord Jesus thank you , you granted my request")

	crypt.CreateRandSeed()

	hdwallet.CreateHDWallet()
}
