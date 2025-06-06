# MLOps at Edge

We have a setup of golang microservices and python code that together supports the MLOps pipelines in Hedge.
Here is a brief description of the golang services. Refer to documentation under python-code for training and inferencing services written in python

| Service Name        | Description                                                                                                                                                                                                                                                                    |
|---------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| hedge-ml-management | The core ml management service that provides APIs and underlying implementations<br/>to support ML data configuration, training, deployment & custom Algorithm registration                                                                                                    |
| hedge-ml-agent      | ml-agent which is deployed on the node/edge orchestrate model deployment and upgrades. It downloads the trained model to the edge node and then <br/>instructs ml-broker to start the ml pipeline and also creates new containers for prediction based on the prediction image |
| hedge-ml-sandbox    | This is a docker in docker implementation where it spawns a new training docker container to train a model on receiving the command                                                                                                                                            |
| hedge-ml-broker     | This is dynamic ml data pipeline based on deployed ML model configuration. It calls prediction service as part of the pipeline, interprets the prediction to generate events                                                                                                   |


A flow depiction on how the ML model deployment works is as below:
![ML Deployment flow](../images/ml_deployment.png?raw=true)


The below diagrams illustrate how the ML prediction at edge is performed:
![ML Prediction at edge](../images/ml_inferencing_flow.png?raw=true)


