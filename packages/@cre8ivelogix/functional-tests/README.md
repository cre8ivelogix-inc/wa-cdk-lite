# Well Architected CDK Lite Functional Tests

This package contains functional tests for the constructs in package (@cre8ivelogix/wa-cdk-lite).

The `cdk.json` file tells the CDK Toolkit how to execute your app. In our case the entry point is bin/functional-tests.ts file

## Useful commands for functional testing.

Goto directory @cre8ivelogix/functional-tests and execute the following command.

* `npm run build`   compile typescript to js
* `npm run watch`   watch for changes and compile
* `npm run functional-tests` deploy functional-test stack to your default AWS account/region
* `npm run test`    perform the jest unit tests
* `cdk deploy`      deploy this stack to your default AWS account/region
* `cdk diff`        compare deployed stack with current state
* `cdk synth`       emits the synthesized CloudFormation template
