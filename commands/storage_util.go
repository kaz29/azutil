package commands

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/arm/examples/helpers"
	"github.com/Azure/azure-sdk-for-go/arm/resources/resources"
	"github.com/Azure/azure-sdk-for-go/arm/storage"
	storagem "github.com/Azure/azure-sdk-for-go/storage"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/spf13/viper"
)

type ServicePrinciple struct {
	Spt                 *azure.ServicePrincipalToken
	AzureClientID       string
	AzureClientSecret   string
	AzureSubscriptionId string
	AzureTenantId       string
}

var servicePrinicpleInstance *ServicePrinciple

func createServicePrinciple(servicePrincipleFileName string) (*ServicePrinciple, error) {
	if servicePrinicpleInstance == nil {
		azureConfig := servicePrincipleFileName
		if azureConfig == "" {
			azureConfig = "azureConfig"
		}
		viper.SetConfigName(servicePrincipleFileName)
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error: config file %s", err))
		}

		strageConfig := map[string]string{
			"AZURE_CLIENT_ID":       viper.Get("AZURE_CLIENT_ID").(string),
			"AZURE_CLIENT_SECRET":   viper.Get("AZURE_CLIENT_SECRET").(string),
			"AZURE_SUBSCRIPTION_ID": viper.Get("AZURE_SUBSCRIPTION_ID").(string),
			"AZURE_TENANT_ID":       viper.Get("AZURE_TENANT_ID").(string)}

		if err := checkEnvVar(&strageConfig); err != nil {
			log.Println("Error: %v", err)
			return nil, err
		}
		spt, err := helpers.NewServicePrincipalTokenFromCredentials(strageConfig, azure.PublicCloud.ResourceManagerEndpoint)
		if err != nil {
			log.Println("Error: %v", err)
			return nil, err
		}

		servicePrinicpleInstance = &ServicePrinciple{
			spt,
			strageConfig["AZURE_CLIENT_ID"],
			strageConfig["AZURE_CLIENT_SECRET"],
			strageConfig["AZURE_SUBSCRIPTION_ID"],
			strageConfig["AZURE_TENANT_ID"],
		}
	}
	return servicePrinicpleInstance, nil
}

func createOrUpdateResourceGroup(resourceGroup string, storageAccount string, location string, sp *ServicePrinciple) error {
	group := resources.Group{Location: to.StringPtr(location)}
	groupsClient := resources.NewGroupsClient(sp.AzureSubscriptionId)

	groupsClient.Authorizer = sp.Spt
	_, err := groupsClient.CreateOrUpdate(resourceGroup, group)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func checkEnvVar(envVars *map[string]string) error {
	var missingVars []string
	for varName, value := range *envVars {
		if value == "" {
			missingVars = append(missingVars, varName)
		}
	}
	if len(missingVars) > 0 {
		return fmt.Errorf("Missing environment variables %v", missingVars)
	}
	return nil
}

func getAccountClient(sp *ServicePrinciple) storage.AccountsClient {
	accoutClient := storage.NewAccountsClient(sp.AzureSubscriptionId)
	accoutClient.Authorizer = sp.Spt
	return accoutClient
}
func newFile(fn string) *os.File {
	fp, err := os.OpenFile(fn, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return fp
}

func createContainer(storageAccountName string, storageAccountKeyValue string, containerName string) error {
	client, err := storagem.NewBasicClient(storageAccountName, storageAccountKeyValue)
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	blobClient := client.GetBlobService()
	container := blobClient.GetContainerReference(containerName)
	_, err = container.CreateIfNotExists()
	if err != nil {
		log.Println("Error: %v", err)
		return err
	}
	log.Printf("Container %s has been created\n", containerName)
	return nil
}

func createStrageAccount(accoutClient storage.AccountsClient, resourceGroup string, strageName string) (string, error) {
	cna, err := accoutClient.CheckNameAvailability(
		storage.AccountCheckNameAvailabilityParameters{
			Name: to.StringPtr(strageName),
			Type: to.StringPtr("Microsoft.Storage/storageAccounts")})
	if err != nil {
		log.Printf("Error: %v", err)
		return "", err
	}

	if !to.Bool(cna.NameAvailable) {
		return "", errors.New("Strage name \"" + strageName + "\" is unavailable")
	}

	cp := storage.AccountCreateParameters{
		Sku: &storage.Sku{
			Name: storage.StandardLRS,
			Tier: storage.Standard},
		Location: to.StringPtr("japaneast")}
	cancel := make(chan struct{})
	if _, err = accoutClient.Create(resourceGroup, strageName, cp, cancel); err != nil {
		log.Println("Create '%s' storage account failed: %v\n", strageName, err)
	}

	keyResults, err := accoutClient.ListKeys(resourceGroup, strageName)
	if err != nil {
		log.Println("Error: %v", err)
		return "", err
	}
	accountKeyList := keyResults.Keys
	pl := *accountKeyList
	accountKey := pl[0]
	value := accountKey.Value
	log.Printf("AccountKey: %s\nValue: %s\n", strageName, *value)

	return *value, nil
}
