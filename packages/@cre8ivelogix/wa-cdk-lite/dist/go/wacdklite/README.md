# Well Architected CDK Lite

This package contains the Well Architected S3 Bucket - WABucket construct.
WABucket constructs extend the Bucket construct provided by AWS CDK and
configures them to satisfy the Well Architected pillars, like
security, reliability, performance, cost-effectiveness and operational-excellence.
If you would like to use premium version which offers 50+ Well Architected constructs
for most commonly used usecases please email info@cre8ivelogix.com

## Create a CDK project

To try out Well Architected CDK Lite by itself you can create a new CDK project by running the following commands

```shell
cdk init app --language typescript
```

This will create a bare bone CDK project.

## Adding WA-CDK Lite dependency

Open CDK project in your favorite IDE, goto package.json file and add the WA-CDK-Lite dependency

```json
{
   "dependencies": {
      "@cre8ivelogix/wa-cdk-lite": "CURRENT_VERSION"
   }
}
```

## Install Dependencies

Once package.json file is updated, you can install the dependencies by running the following commands

```shell
npm install
```

## Using Well Architected CDK constructs

Now that you have all the constructs available from Well Architected CDK Lite,
use them just like any other class in typescript.

For example in the following code snippet we are using WaBucket from AWS-CDK-Lite package

```go
import {WaBucket} from "./wa-bucket";

new WaBucket(this, "MyBucket", {})
```

The Bucket created using WaBucket by default satisfies various security and compliance requirements,
checkout WaBucket class documentation for details.

### Building Java/.Net/Python/Javascript Distribution

```shell
docker run -it --rm --entrypoint sh -v $(pwd)/wa-cdk-lite:/wa-cdk-lite jsii/superchain:1-buster-slim-node16
```

or you can run the script docker-package.sh, when docker prompt is shown,

```shell
cd wa-cdk-lite
npm run package
```

This will cross compile the typescript construct to supported languages.

### Well Architected CDK Version Updates

The version of all CDK packages are fixed to the same version. If an
upgrade is desired, please make sure that all CDK packages are upgraded
to the same version at the same time. Different versions of CDK mixed
in the project may cause type incompatibilities.

### Useful commands

* `npm run build` compile typescript to js
* `npm run watch` watch for changes and compile
* `npm run test` perform the jest unit tests
* `npm run build-test` builds and runs unit tests
* `npm run doc-gen` generate documentation
* `npm run pretty` runs prettier on source code
