package commands

import (
	"bufio"
	"log"

	storagem "github.com/Azure/azure-sdk-for-go/storage"
)

type UploadVhdCommand struct {
	ResourceGroup            string `short:"r" long:"resourceGroup" description:"Resource Group Name"`
	StorageAccount           string `short:"s" long:"storageAccount" description:"Storage Account Name"`
	Container                string `short:"c" long:"containerName" description:"Container Name"`
	Vhd                      string `short:"v" long:"vhd" description:"VHD file name"`
	ServicePrincipleFileName string `short:"f" long:"servicePrinciple" description:"Service PrincipleFileName"`
}

func (command *UploadVhdCommand) Execute(args []string) error {
	sp, err := createServicePrinciple(command.ServicePrincipleFileName)
	if err != nil {
		log.Println("Error: %v", err)
		return err
	}
	accoutClient := getAccountClient(sp)

	keyResults, err := accoutClient.ListKeys(command.ResourceGroup, command.StorageAccount)
	if err != nil {
		log.Printf("Can not access Storage Account %v¥n", err)
		return err
	}

	accountKeyList := keyResults.Keys
	pl := *accountKeyList
	accountKey := pl[0]
	log.Printf("AccountKey: %s\nValue: %s\n", command.StorageAccount, *accountKey.Value)

	client, err := storagem.NewBasicClient(command.StorageAccount, *accountKey.Value)
	blobClient := client.GetBlobService()

	fp := newFile(command.Vhd)
	defer fp.Close()
	reader := bufio.NewReader(fp)
	stat, err := fp.Stat()
	size := uint64(stat.Size())

	blobClient.CreateBlockBlobFromReader(command.Container, "storage.vhd", size, reader, nil)
	return nil
}
