# hEdge
## hEdge IoT Platform

Welcome to hEdge IoT Platform! If you have been using edgex foundry or heard of it, this is the perfect place to build your complete solution.
If you are not aware of edgex foundry, no worries, all the complexities have been taken under the hoods so you can focus on solving your IoT problem.
As a contributor to this platform, BMC Helix leveraged its expertise in IT domain & AIOps and made use of IoT infrastructure services from edgex foundry to bring to you a simple, low foot-print, yet very powerful platform.
So let us get started.

hEdge is a complete IoT Platform built on top of edgex foundry leveraging its SDK that adds to its functionality so it is a self-sufficient IoT platform.
One of the key functions that it adds is the machine learning capability that provides the ability for a non data-scientist to define ML model data configuration, train and deploy trained models to the edge. 
The deployed model then drives real-time prediction as well as event/alert generation at the edge.
Customers can also register their code (python ML docker image) that enables data-science persona to Bring your own Algorithm (BYOA) while taking advantage of in-built ML data pipelines and predictions.
There was also a realization that to get better outcomes from ML predictions, it is important that we have the complete data. This is where hEdge added the concept of context data
by extending device metadata. Device data is then enriched in real time with this context data to provide a data pipeline that drives differentiated outcomes when it comes to machine learning or rule/workflows.

As an IoT platform, hEdge also addresses the challenge in managing multiple edge deployments. 
For this, many of the concepts from the IT industry have been leveraged and applied to the edge.
The edge deployment management is still retaining the edge requirements of real-time issue identification and remediation at edge.

On how hEdge addresses the management of edge deployments, refer to the Deployment architecture diagram.
The management layer (to the right), also called as core, serves the long-term storage requirements. 
It also provides the user a management UI and the dashboard (grafana) to view device data in the form of metrics, predictions and events in near-real time.

The edge nodes on the left are the edge services in an air-gapped OT environment. 
The data-collection, ML inferencing and rule/workflow execution all take place in this layer in near real time.
If you are concerned with latency and data volume being sent from edge to core, this is addressed by providing the user ability to
configure down-sampling intervals and/or aggregations (avg, min. max).

![hEdge Deployment Architecture](./images/hEdgeDeploymentArch.png?raw=true)

The below diagram is a depiction of core hEdge Services that were added on top of edgex foundry to make it a complete platform.

![hEdge Architecture](./images/hEdgeArchitectureLayered.png?raw=true)

The next few years will be very exciting as we embark on journey where-in we will be adding the concept of digital twin and simulations, so we need contributions from both data-scientists and golang developers.
Here is a quick guide to get started with your development environment and the build environment. If you need to just start using the product, refer to simple steps at hedge-deployments/docker.

## How to set up local Development Environment
**Prerequisites**
- Go (Min version 1.24.2)
- IDE: Goland (Preferred IDE)
- Redis
- Set the environment variable EDGEX_SECURITY_SECRET_STORE=false so it is available in IDE (For Mac add this to .zshrc). This bypasses the security enforcement in the development environment
- We leverage the simplified configuration concept from edgex foundry. All common configurations are saved in a global configuration.yaml file and we can override the same using service specific configuration.yaml or environment variable. 
So set up the environment variable as below. This step is important to get all common configurations.
  EDGEX_COMMON_CONFIG={edgex-go-dir}/cmd/core-common-config-bootstrapper/res/configuration.yaml
- To ensure TLS cipher compatibility in Go v1.24.2, set the environment variable GODEBUG=tlsrsakex=1 so it is available in your IDE. (For Mac users, you can add this line to your .zshrc file to apply it globally)
- Install pkg-config (for Mac, $brew install pkg-config)
- edgex foundry services (core-metadata & support notifications for local development)
- mosquitto (MQTT broker) for some services
- For some export services, backend database where it is best to refer to one of the deployment environments


**Set up Redis server locally**
* For Windows installation set up [wsl]([https://learn.microsoft.com/en-us/windows/wsl/install) and then follow [Redis instructions](https://redis.com/blog/install-redis-windows-11/)

**Set up edgex foundry**

Following are the steps to set up edgex environment. Make sure you check out edgexfoundry/edgex-go with version v3.1.0
1. Clone this GitHub repo on your machine. As a best practice, 
   create a workdirectory called hedge under which carry out the below steps 
    > git clone https://github.com/edgexfoundry/edgex-go
2. goto edgex-go directory
3. check out the supported version of edgex-go for edge-iot
    > git checkout v3.1.0
4. Make sure that redis-server is running 
5. We only need to run (core-metadata & support-notifications) on development environment, 
   so from the IDE. For each of the services in here, right click on cmd/main.go 
   and select create Debug or Run configuration. Thereafter, edit it to rename the configuration 
   and correct the work-directory to a path where ./res folder or main.go is located. 
   An example is attached in here

![Sample Golang Debug/Run Configuration](images/golangDebugConfigEg.png?raw=true "Debug Configuration")

## 3. Set up hEdge project 
1. Clone this Github repo on your machine 
    > git clone https://github.com/bmchelix/hedge.git 
   let this be outside of edgex foundry setup, one of the ways could be to local edgex foundry project from earlier step into _local folder that is excluded from git )
2. goto hedge directory
3. Ensure you download goland dependencies by running
   > go mod download all
4. In most cases, you might need to have a device service running so it registers the device and 
   generates the data
5. First check res/configiration.yaml so it refers to localhost, not logical names 
   that is required in docker containers 
6. Start with device-services/hedge-virtual-device, right-click cmd/main.go and click Run, 
   edit the working directory so it points to where main.go is
7. If the above is running, your dev environment is all setup
8. Some of the services depend on MQTT, you can refer to mqtt broker ( mosquitto) of your 
   shared dev environment. Same for database like victoria and elastic.
9. If you are working on Hedge UI, goto ui/admin and look for the corresponding documents in there

## hEdge Code Organization

The following are the main directories at the root with a brief description

| Directory                   | Description                                                                                                                                                                                                           |
|-----------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| common                      | Contains shared library functions, DTOs & constants that are shared across services                                                                                                                                   |
| device-services             | device services support different protocols, Included here is device-virtual service that <br/>helps developers with data generation                                                                                  |
| app-services                | All hEdge core services other than ML services are in here. These all services are based on<br/>app sdk of edgex foundry. Some examples include data-enrichment, export-event, device-extensions, user-app-management |
| edge-ml-service             | hEdge ML services responsible for ML data configuration, training, deployment, automatic pipeline(ml-broker) are in here<br/>                                                                                         |
| edge-ml-service/python-code | Python code that implements various ML algorithsm for training and prediction are in here                                                                                                                             |
| hedge-deployment            | deployment related Makefile docker-compose files are in here                                                                                                                                                          |
| hedge-init                  | initialization scripts that are used during deployments.<br/>Example include opensearch index creation, ml-algo registration                                                                                          |
| hedge-swagger-ui            | swagger doc related UI code and where swagger.json is the generated specs using a make command                                                                                                                        |
| external                    | hEdge infrastructure services outside of hedge, but where we need to update configs or upgrade libs for security reasons<br/>It includes edgex-submodule as well                                                      |
| contents                    | Sample workflows and rules to get started in kept here and is part of installation                                                                                                                                    |
| mocks                       | generated mocks for interfaces that is used in unit tests                                                                                                                                                             |


## hEdge Development Resources
Below are some of the relevant resources to get started on Hedge
1. If you are new to Hedge, here is where you can refer to get a good overview of edgex foundry 
   which is the underlying software, hedge uses(without modification):
> https://www.edgexfoundry.org/
> https://docs.edgexfoundry.org/3.1/getting-started/Ch-GettingStartedGoDevelopers/

2. To get to specific service API documentation, navigate to the specific service first and then 
   look for their Swagger documentation
> https://docs.edgexfoundry.org/3.1/api/core/Ch-APICoreMetadata/

2. If you are developing a north-bound service, here is the SDK reference to get started
> https://docs.edgexfoundry.org/3.1/microservices/application/ApplicationServices/

3. If you are developing a device-service to pull data from propriety protocol, here is the device SDK to get you started
> https://docs.edgexfoundry.org/3.1/microservices/device/DeviceService/

4. If you are working to secure the microservices using a secure gateway, there is where you can get started.
   edgex foundry is bundled with nginx gateway & with built-in protection using jwt token
> https://docs.edgexfoundry.org/3.1/security/Ch-APIGateway/

        
## Setting up Build Environment
**Prerequisites**

The below setup is required when first time setting up a build machine 
- git
- Go
- make
- Docker

**Build Steps**
1. Clone this GitHub repo on your machine 
    > git clone https://github.com/bmchelix/hedge.git
2. goto hedge directory
3. When it is a brand new machine etc, run the following from hEdge directory where go.mod is located
    > go mode download all
4. Run a clean build of all the micro services if you want to build executables only
    > $make clean build

    App binaries will be available inside respective microservice folders with same name as that of the microservice

    *NOTE: There may be special config instructions for individual services, make sure to check out the individual Readme.

5. **For Mac M1 Silicon only**: set another ENV value instead of default 'jenkins' in hedge/Makefile - for creating suitable images suitable for this specific platform (will pull images from the default source - Docker Hub):
    > ENV ?= jenkins

6. To build docker images add _docker_ as a target in above make command, or run a new command
    > _'make docker'_   
      or, _'make clean build docker'_   

7. To build a specific docker image, you can refer to Makefile for specific microservice name and run the make command
> eg make docker SERVICE=hedge-ml-management --> microservice name

## Push the images to docker registry
If you want to push specific image that is built to the registry, execute the below command. Note that the image name is not a fully-qualified image name
> make docker SERVICE=<service-name> IMAGE_PUSH=true

> eg, $make docker SERVICE=hedge-ml-management IMAGE_PUSH=true

You might get an error if you are not logged in to the registry in which case, login using the below
> docker login docker.io

To push all hedge images, run the below command
> make push SERVICE=hedge-services

To push all edgex-foundry images
> make docker SERVICE=edgex-services IMAGE_PUSH=true
> make push SERVICE=edgex-services

To push hedge infrastructure (opensearch, victoria, mosquitto, node-red, redis), execute
> make docker SERVICE=external-services IMAGE_PUSH=true
> make push SERVICE=external-services


To create the deployment package,
> make package

You will find the deployment package files under package directory (hedge-node-docker.tgz & hedge-core-docker.tgz)

## Deployment
For deployment Makefile and instructions, refer to Readme.md under hedge-deployment/docker

# Jenkins Builds
In case you want to setup Jenkins build locally, here are a few guidelines that you can get started wiuth:
Configure Jenkins Daily Builds on one of the machines. Let us call the Jenkins job as 'Daily_Checkin_Builds'.  This Jenkins job will be triggered for every code check-in that happens in GIT.  The url to access Jenkins by default is (http://<jenkins-machine>:8080). \
\
**-Note:** Jenkins is typically configured using the Master – Slave concept.  Here '<jenkins-machine>' can be the master and <jenkins-agent-machine> can be the slave machine.\
\
At the heart of the Jenkins job is the Jenkinsfile. This is a Jenkins pipeline job which contains the following stages:\
**-Start:** This stage sends an email to the configured recipients informing them that the Jenkins build has started. It stops the containers (if any), deletes the existing hedge-images on the build server before it starts to build the new images.\
\
**-Git Checkout:** This stage clones the latest code checked into the master branch in GIT, into the Jenkins workspace. \
\
**-Build:** This stage builds the binaries for each of the microservices. It triggers the command 'make build'. This make command invokes the 'build' target specified in the Makefile.\
\
**-Test:** This stage is in-place for unit testing. Currently there are no unit testing scripts available.\
\
**-Create Docker Images:** This stage creates the docker images of each of the hedge microservices. It triggers the command 'make docker'. This command invokes the 'docker' target specified in the Makefile.\
\
<br />
The 'post' section is configured for the below sections:\
\
**-Always:** This section will be executed always irrespective of whether the Jenkins job is successful or not. Here we clean up all the generated binaries for each of the hedge microservices that were created in the 'Build' stage.  It triggers the command 'make clean'. This command invokes the “clean” target specified in the Makefile.\
\
**-Successful:** This section will be executed if the Jenkins build was successful. This section is configured to push the new docker images to Harbor. It triggers the command 'make push'. This command invokes the 'push' target specified in the Makefile. It also sends out an email to all the configured recipients informing them about the build status.\
\
**-Failure:** This section will be executed if the Jenkins build fails. This section is configured to send out an email to all the configured recipients informing them about the build status. 

**Note:**
    You might typically setup 2 Jenkins jobs configured for hEdge
        a. Daily_Checkin_Builds - This job can run every time a checkin is done in GIT.
        b. Hedge_Nightly_Build - This job can be scheduled to run at a schedule time everyday.
        
#### **-To configure all the required credentials in Jenkins:**
    1.	Go to Dashboard->Manage Jenkins->Manage Credentials. Specify all the credentials that Jenkins would required to connect to the respective systems like GIT, docker.io or your custom docker registry, hEdge servers etc, and select the appropriate credentials where ever required.
    2.	Under Stores scoped to Jenkins-> Click Jenkins link-> Click Global Credentials(unrestricted) link >Click Add Credentials link.
    3.	Specify the username, password, ID and description.
    4.	Click OK. 
    
#### **-To configure Jenkins to connect to GIT:**
    1.	For Jenkins to pull the latest source code from Git into Jenkins workspace, Jenkins would need the GIT credentials to connect to GIT.
    2.	Select a Jenkins job, click Configure. Under Build Triggers, Pipeline (for Pipeline jobs), select the below options:
            Definition: Pipeline script from SCM
            SCM: GIT
            Repositories: Specify the repository url
            Credentials: Select the required git credentials which is required for Jenkins to connect to GIT.
            Branches to build: Specify the branch which needs to be checkout.
            Additional Behaviours: Select the option “Wipe out repository and force clone”.
            Script: Specify the path to the Jenkinsfile in GIT.
    3.	Click Save.


## Python Unit Tests

Python unit tests are running using prebuilt docker image to improve performance. When `edge-ml-service/python-code/tests/requirements.txt` is updated with new dependencies,
the image needs to be rebuilt. Follow the steps:

1. Update  `edge-ml-service/python-code/tests/requirements.txt` with new dependencies.
NOTE: It is important to use canonical names for python libraries. Pip is case-insensitive and also doesn't discern between underscore and minus sign.
2. Rebuild and push the image to registry by running `make push-python-test-coverage-base` for `edge-iot` project root folder.
NOTE: This might take about 20 mins or even more.
3. Push your changes to GitHub, create a PR and merge.
MOTE: If you create a PR before step 2 is complete python unit tests might fail because the image hasn't been pushed to registry yet.

## Swagger Documentation

Hedge Public API can be used with Swagger. Read this [page](IT-Operations-Management.Operations-Management.EDGEMASTER.Developing.Accessing-API-using-Swagger-UI.WebHome) to find out more.   
Current implementation use _swag_ generator by _swaggo_. Learn about _swaggo/swag_ [here](https://github.com/swaggo/swag?tab=readme-ov-file). 
We have several guidelines and hints for composing swagger documentation(comments).
1. Put comments over the method that only contain exactly one route definition
2. Such method should be named as `addRouteRest<...>`
3. Pay attention to `@Router` url and HTTP method (HTTP method should be in square brackets)
4. It's highly recommended to run Swagger specification file generation every time documentation editing is completed

For the reference check out `.../app-services/hedge-device-extensions/internal/router/router.go` file: 
```go
// @Summary		Update Device Extensions
// @Tags	Hedge Device Extensions - Device Extensions
// @Produce	json
// @Accept	json
// @Param			deviceName	path		string				true	"device name"
// @Param 			q 			body 		[]dto.DeviceExt 	true 	"List of the device extensions objects"
// @Success			200			{array}		[]dto.DeviceExtension 	""
// @Success			200			{array}		[]string 	"empty array"
// @Failure			400			{object}	error	"{"message":"Error message"}"
// @Failure			404			{object}	error	"{"message":"Error message"}"
// @Failure			409			{object}	error	"{"message":"Error message"}"
// @Failure			500			{object}	error	"{"message":"Error message"}"
// @Router			/api/v3/metadata/device/extn/{deviceName} [put]

func (r *Router) addRouteRestUpdateDeviceExtn() {
	...
}
```  
<br/>

#### Running Swagger specification file generation 
There are several ways to generate Swagger specification file. After completing adding/editing/removing _swag_ comments over REST API method you can either:
- Run `make generate-swagger-spec` from the project root folder
- Run in the terminal from project root folder `swag init --parseInternal=true --generalInfo=doc.go --pd=true --ot=json --output=./swagger-ui/res/swagger/`
- If in IDE, go to `doc.go` and run `go:generate` command 

You can start `swagger-ui` service in order to check generated file. Before commiting the changes to the Git, be sure to 
regenerate swagger spec file (since swagger-ui restructures swagger.json file in the runtime).

#### Swagger General Info
General info tags for Swagger (general description, license info, base path, etc.) is located in `doc.go`


