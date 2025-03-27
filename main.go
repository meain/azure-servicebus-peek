package main

// This tool is useful for peeking messages from an Azure Service Bus queue,
// since the Azure CLI currently lacks a built-in method for this functionality.

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func main() {
	// Define command-line flags
	resourceGroup := flag.String("resource-group", "", "Azure Resource Group")
	namespace := flag.String("namespace", "", "Azure Service Bus Namespace")
	queueName := flag.String("queue", "", "Azure Service Bus Queue Name")
	flag.Parse()

	// Validate required flags
	if *resourceGroup == "" || *namespace == "" || *queueName == "" {
		flag.Usage()
		log.Fatal("All flags are required")
	}

	// Construct the fully qualified namespace
	fullyQualifiedNamespace := fmt.Sprintf("%s.servicebus.windows.net", *namespace)

	// Create a default Azure credential
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("Failed to obtain a credential: %v", err)
	}

	// Create a Service Bus client
	client, err := azservicebus.NewClient(fullyQualifiedNamespace, cred, nil)
	if err != nil {
		log.Fatalf("Failed to create Service Bus client: %v", err)
	}
	defer client.Close(context.Background())

	// Create a receiver for the specified queue
	receiver, err := client.NewReceiverForQueue(*queueName, nil)
	if err != nil {
		log.Fatalf("Failed to create receiver: %v", err)
	}
	defer receiver.Close(context.Background())

	// Peek messages from the queue
	messages, err := receiver.PeekMessages(context.Background(), 10, nil)
	if err != nil {
		log.Fatalf("Failed to peek messages: %v", err)
	}

	for _, msg := range messages {
		fmt.Printf("%s: %s\n", msg.MessageID, string(msg.Body))
	}
}
