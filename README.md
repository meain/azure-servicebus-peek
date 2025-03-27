# Azure Service Bus Message Peek Tool

A command-line utility to peek messages from an Azure Service Bus queue without removing them.

## Prerequisites

- Go 1.16 or later
- Azure CLI authenticated with appropriate permissions
- Access to an Azure Service Bus namespace and queue

## Installation

```bash
go install
```

## Usage

```bash
azservicebus-peek -resource-group <resource-group> -namespace <namespace> -queue <queue-name>
```

### Parameters

- `-resource-group`: Azure Resource Group name
- `-namespace`: Azure Service Bus Namespace
- `-queue`: Azure Service Bus Queue Name

All parameters are required.

### Example

```bash
azservicebus-peek -resource-group myResourceGroup -namespace myNamespace -queue myQueue
```

## Output

The tool will display the MessageID and body content of up to 10 messages from the specified queue.

## Authentication

The tool uses Azure Default Credentials. Make sure you are logged in with the Azure CLI:

```bash
az login
```
