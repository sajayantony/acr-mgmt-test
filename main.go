package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2017-10-01/containerregistry"
	"github.com/Azure/draft/pkg/azure/iam"
)

func main() {
	fmt.Println("Get Classic Registry")
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*60)
	defer cancel()

	client, err := GetRegistriesClient(ctx, "dfb63c8c-7c89-4ef8-af13-75c1d873c895")
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := client.Get(ctx, "sajaydev", "sajay")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("StorageAccount: %s\n", *r.StorageAccount.ID)
}

// GetRegistriesClient return Container Registry Management client
func GetRegistriesClient(ctx context.Context, subID string) (c containerregistry.RegistriesClient, err error) {
	registriesClient := containerregistry.NewRegistriesClient(subID)
	auth, err := iam.GetResourceManagementAuthorizer(iam.AuthGrantType())
	if err != nil {
		return c, fmt.Errorf("Failed to get client. Err: %v", err)
	}
	registriesClient.Authorizer = auth
	registriesClient.AddToUserAgent(containerregistry.UserAgent())
	return registriesClient, nil
}
