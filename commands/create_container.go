package commands

import (
	"log"
)

type CreateContainerCommand struct {
	ResourceGroup            string `short:"r" long:"resourceGroup" description:"Resource Group Name"`
	Location                 string `short:"l" long:"location" description:"Location. e.g. japaneast"`
	StorageAccount           string `short:"s" long:"storageAccount" description:"Storage Account Name"`
	Container                string `short:"c" long:"containerName" description:"Container Name"`
	ServicePrincipleFileName string `short:"f" long:"servicePrinciple" description:"Service PrincipleFileName"`
}

func (command *CreateContainerCommand) Execute(args []string) error {

	sp, err := createServicePrinciple(command.ServicePrincipleFileName)
	if err != nil {
		log.Println("Error: %v", err)
		return err
	}

	err = createOrUpdateResourceGroup(command.ResourceGroup, command.StorageAccount, command.Location, sp)
	if err != nil {
		log.Println("Error: %v", err)
		return err
	}

	accoutClient := getAccountClient(sp)

	accountKey, err := createStrageAccount(accoutClient, command.ResourceGroup, command.StorageAccount)
	if err != nil {
		log.Printf("Storage Account Creation Error : %v", err)
		return err
	}

	err = createContainer(command.StorageAccount, accountKey, command.Container)
	if err != nil {
		log.Printf("Container Creation Error : %v¥n", err)
		return err
	}
	return nil

}
