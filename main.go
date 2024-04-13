package main

import (
	"fmt"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"log"
	"os"
	"strconv"
)

func main() {
	var accs = ""
	for i := 6; i < 50; i++ {
		entropy, err := bip39.NewEntropy(128)
		if err != nil {
			log.Fatal(err)
		}

		mnemonic, _ := bip39.NewMnemonic(entropy)
		//var mnemonic = "pepper hair process town say voyage exhibit over carry property follow define"
		fmt.Println(mnemonic)
		seed := bip39.NewSeed(mnemonic, "") //这里可以选择传入指定密码或者空字符串，不同密码生成的助记词不同

		wallet, err := hdwallet.NewFromSeed(seed)
		if err != nil {
			log.Fatal(err)
		}

		path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0") //最后一位是同一个助记词的地址id，从0开始，相同助记词可以生产无限个地址
		account, err := wallet.Derive(path, false)
		if err != nil {
			log.Fatal(err)
		}

		address := account.Address.Hex()
		privateKey, _ := wallet.PrivateKeyHex(account)
		publicKey, _ := wallet.PublicKeyHex(account)

		fmt.Println(address)    // id为0的钱包地址
		fmt.Println(privateKey) // 私钥
		fmt.Println(publicKey)  // 公钥
		sprintf := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", strconv.Itoa(i), mnemonic, address, privateKey, publicKey)
		accs = accs + sprintf
	}
	wFile(accs)
	//path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1") //生成id为1的钱包地址
	//account, err = wallet.Derive(path, false)
	//if err != nil {
	//	log.Fatal(err)
	//}
	////0x46758Df513C5AdF347dFcEB26666F2Ad91e03631
	////7660389a44225b09fec999d9ed8ab18321eb975e2545890dd95ca23385bac0fc
	////7660389a44225b09fec999d9ed8ab18321eb975e2545890dd95ca23385bac0fc
	//fmt.Println("address1:", account.Address.Hex())

	// 创建一个文件，如果文件已存在则截断文件
}

func wFile(acc string) {
	// 创建一个文件，如果文件已存在则截断文件
	file, err := os.Create("fileacc.txt")
	if err != nil {
		fmt.Println("创建文件时出错：", err)
		return
	}
	defer file.Close()

	// 写入内容到文件
	content := []byte(acc)
	_, err = file.Write(content)
	if err != nil {
		fmt.Println("写入文件时出错：", err)
		return
	}

	fmt.Println("文件写入成功")
}
