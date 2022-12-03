import {App, Stack} from "aws-cdk-lib";
import { expect as cdkExpect, haveResource, SynthUtils } from "@aws-cdk/assert";
import {WaBucket} from "../../lib";
import * as s3 from "aws-cdk-lib/aws-s3";

let app: App;
let stack: Stack;

describe("Default Well Architected S3 Bucket Configurations", () => {

    beforeEach(() => {
        app = new App();
        stack = new Stack(app, "TestStack");
    });

    const defaultEncryption = {
        ServerSideEncryptionConfiguration: [
            {
                ServerSideEncryptionByDefault: {
                    SSEAlgorithm: "AES256"
                }
            }
        ]
    };
    const customEncryption = {
        ServerSideEncryptionConfiguration: [
            {
                ServerSideEncryptionByDefault: {
                    KMSMasterKeyID: {
                        "Fn::GetAtt": ["MyBucketKeyC17130CF", "Arn"]
                    },
                    SSEAlgorithm: "aws:kms"
                }
            }
        ]
    };

    test.each`
      actual                  | expected
      ${undefined}            | ${defaultEncryption}
      ${null}                 | ${defaultEncryption}
      ${s3.BucketEncryption.KMS} | ${customEncryption}
   `(
        "Bucket encryption status should be $expected when $actual",
        ({ actual, expected }) => {
            new WaBucket(stack, "MyBucket", {
                encryption: actual
            });
            cdkExpect(stack).to(
                haveResource("AWS::S3::Bucket", {
                    BucketEncryption: expected
                })
            );
            expect(SynthUtils.toCloudFormation(stack)).toMatchSnapshot();
        }
    );

    const defaultAccess = {
        BlockPublicAcls: true,
        BlockPublicPolicy: true,
        IgnorePublicAcls: true,
        RestrictPublicBuckets: true
    };
    const customAccess = {
        BlockPublicAcls: true,
        IgnorePublicAcls: true
    };
    test.each`
      actual                             | expected
      ${undefined}                       | ${defaultAccess}
      ${null}                            | ${defaultAccess}
      ${s3.BlockPublicAccess.BLOCK_ACLS} | ${customAccess}
   `(
        "bucket public access should be $expected when $actual",
        ({ actual, expected }) => {
            new WaBucket(stack, "MyBucket", {
                blockPublicAccess: actual
            });
            cdkExpect(stack).to(
                haveResource("AWS::S3::Bucket", {
                    PublicAccessBlockConfiguration: expected
                })
            );
            expect(SynthUtils.toCloudFormation(stack)).toMatchSnapshot();
        }
    );

});
