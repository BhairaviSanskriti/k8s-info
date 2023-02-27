# Table of Contents
- [k8s-info](#k8s-info)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Retrieve service data](#retrieve-service-data)
    - [Retrieve deployment data](#retrieve-deployment-data)
  - [Output formatting](#output-formatting)
  - [Advanced Usage](#advanced-usage)
  - [Conclusion](#conclusion)

# k8s-info
A command-line tool that retrieves deployment or service data from a Kubernetes cluster. This tool uses the Kubernetes Go client library to communicate with a Kubernetes API server to retrieve information about deployments and services.

This CLI tool was developed using the Cobra library for Go.

## Installation
Before running the `k8s-info` command, ensure that the following prerequisites are met:
- Go version 1.13 or higher is installed.
- A Kubernetes cluster is accessible and a *kubeconfig* file is available.

To install *k8s-info*, follow the steps below:
1. Clone the *k8s-info* repository from GitHub:
```
git clone https://github.com/BhairaviSanskriti/k8s-info.git
```
2. Change to the *k8s-info* directory:
```
cd k8s-info
```
3. Build the `k8s-info` executable by running the following command:
```
go build -o k8s-info
```
4. Make it easier to run the `k8s-info` command from anywhere in the terminal by executing the below command:
```
sudo mv ./k8s-info /usr/local/bin/
```

## Usage
`k8s-info` has the following usage pattern:
```
k8s-info [flags]
```
- `-n`, `--namespace`: The namespace to retrieve deployments from. The default is "default".
- `-k`, `--kubeconfig`: The absolute path to the kubeconfig file. If this flag is not set, the function will check for the kubeconfig file in the following locations in order:
    - The path specified in the `KUBECONFIG` environment variable.
    - `$HOME/.kube/config`.
    - `$PWD/kubeconfig`.
    
If none of the above locations contains a kubeconfig file, the function will assume that it is running inside a Kubernetes cluster and try to obtain the configuration from the in-cluster environment. Set `KUBERNETES_SERVICE_HOST` and `KUBERNETES_SERVICE_PORT` in case the function uses in-cluster info.


The application has two main sub-commands: `get svc` and `get deploy`.

### Retrieve service data
To retrieve service data, run the following command:
```
k8s-info get svc --kubeconfig </path/to/kubeconfig>
```
![image](https://user-images.githubusercontent.com/106534693/221403374-2ace7dee-8926-430c-b80f-23f80ae7ccc0.png)

This command retrieves the services in the *default* namespace and outputs their names, types, and exposed ports in JSON format.

### Retrieve deployment data
To retrieve deployment data, run the following command:
```
k8s-info get deploy -n kube-system
```
![image](https://user-images.githubusercontent.com/106534693/221397451-5adc5265-9ea8-4311-863e-4dbb6ae8d935.png)

This command retrieves the deployments in the *kube-system* namespace and outputs their names, healthy replicas, and unhealthy replicas in JSON format.

## Output formatting
By default, the output of the `k8s-info` command is in plain text format. However, you can also use the jq command-line tool to format the output as JSON or manipulate it in other ways.

To install [jq](https://stedolan.github.io/jq/), follow the installation instructions for your operating system or package manager. Once installed, you can pipe the output of `k8s-info` to `jq` like this:
```
k8s-info get deploy -n kube-system | jq 
```
![image](https://user-images.githubusercontent.com/106534693/221397510-59ca1fad-0147-4f31-8740-341535210070.png)

This will pretty-print the output in JSON format, which can make it easier to read and parse. You can also use `jq` to filter the output based on specific criteria, extract specific fields, and perform other operations.

Note that using `jq` is completely optional and depends on your needs. However, it can be a powerful tool for working with JSON output from various command-line tools, including `k8s-info`.

## Conclusion
`k8s-info` is a simple command-line tool that retrieves deployment and service data from a Kubernetes cluster. It is easy to install and use, and can be extended to support more functionality as needed. If you encounter any issues or have any feedback, please feel free to create an issue on the *k8s-info* GitHub repository.
