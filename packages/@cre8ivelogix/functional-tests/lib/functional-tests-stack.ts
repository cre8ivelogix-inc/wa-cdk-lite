import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { WaBucket } from "@cre8ivelogix/wa-cdk-lite";

export class FunctionalTestsStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);
    
    new WaBucket(this, "MyTestBucket", {});
  }
}
