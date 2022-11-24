# report-operator

This operator simply watched for `VulnerabilityReport` CRDs (already created by the `trivy` operator) and uploads all reports to an S3 bucket.

## Description

Because this relies on CRDs already provided by Aqua Security, no need to define/install them here. You need to first install the `trivy-operator`, which can be found [here](https://github.com/aquasecurity/trivy-operator). Once the `trivy-operator` is installed, you can install this operator. 

## Getting Started

You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster

1. Follow instructions for installing CRDs from `trivy-operator`.
2. Please take a look at the `build-and-deploy.sh` shell script and modify the variables as needed:
    * `IMAGE_NAME` - the name of the image to be built and deployed
    * `IMAGE_TAG` - the tag to use for the image
    * `NAMESPACE` - the namespace to deploy the operator to
    * `S3_BUCKET` - the S3 bucket to upload reports to
    * The below are optional, if not specified will default to environment variables
    * `AWS_ACCESS_KEY_ID` - the AWS access key ID to use
    * `AWS_SECRET_ACCESS_KEY` - the AWS secret access key to use
    * `AWS_REGION` - the AWS region to use
3. This script will automate all steps for you, including building the image, pushing it to a registry, and deploying the operator to the cluster. You can run it with the following command:

    ```bash
    ./build-and-deploy.sh
    ```

  * If using ECR, please uncomment `make ecr-login` line in the script.

### How it works

This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/) 
which provides a reconcile function responsible for synchronizing resources untile the desired state is reached on the cluster 


### Modifying the API definitions

If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

### Modify CRD

You must add the `processed` boolean to the `VulnerabilityReport` CRD. This is used to determine if the report has already been processed or not.

```yaml
additionalPrinterColumns:
  ...
  - description: Has report been processed
    name: processed
    jsonPath: .report.processed
    priority: 1
    type: boolean
  schema.properties.report:
    ...
    processed:
      description: Whether the report been processed and uploaded to S3
      type: boolean
```

For this either perform:

* `kubectl edit crd/vulnerabilityreports.aquasecurity.github.io`
* `kubectl patch`
* Kustomize:
  * Edit `config/crd/bases/aquasecurity.github.io_vulnerabilityreports.yaml`
  * Run `make manifests`
  * `kubectl apply -f config/crd/aquasecurity.github.io_vulnerabilityreports.yaml`.

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.