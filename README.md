# Go CLI Examples
This repository contains a set of sample Command Line Interface implementaions using Go. 

# Projects
At the moment, this repo has one sample implementation: **deployer**. It is a scaffoling for an imaginary deployment automation tool for a 3-tier application, which has Web, API and Database artifacts to deploy. This CLI will act as an entry point for the automation actions such as deploy, undeploy and status check. 


# Building the project

- Go to **deployer** sub directory
- On Linux, go to a terminal and execute the following:
   - `$ go build -o build/deployer`
     - This will create an executable called deployer in build sub directory.

- On Windows, it can be done by going to command prompt and executing the following:
  - `$ go build -o build/deployer.exe`
    - This will generate an executable called deployer.exe in build sub directory.

- In you would like to generate binaries for multiple platforms & architectures, you could do as follows (On Linux, Windows or Mac). 

  - For Linux 64 bit target system:
    - `$ env GOOS=linux GOARCH=amd64 go build -o build/deployer`
  - For Windows 64 bit target system:
    - `$ env GOOS=windows GOARCH=amd64 go build -o build/deployer.exe`

# Execution
Go to **deployer/build** sub directory.

On Linx
- `$ ./deployer help`
- `$ ./deployer deploy help`
- `$ ./deployer deploy web help`
- `$ ./deployer deploy web --config web-deploy-config.yaml`

# References
If you would like to explore more, please refer to the official repositories of Cobra and Viper. 
 - Cobra: https://github.com/spf13/cobra
 - Viper: https://github.com/spf13/viper

To know more on Go, please refer to the official Go website. 
- Go: https://golang.org/