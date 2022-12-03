import * as s3 from "aws-cdk-lib/aws-s3";
import {BucketEncryption} from "aws-cdk-lib/aws-s3";
import {Construct} from "constructs";
import {Duration} from "aws-cdk-lib";
import {Alarm, ComparisonOperator, Metric} from "aws-cdk-lib/aws-cloudwatch";

export enum WaBucketAlarms {
    /**
     * This enum is used to represent and alarm name for 4XX errors e.g. 404 Not Found
     */
    FourxxErrorsAlarm = "4xxErrorsAlarm",
    /**
     * This enum is used to represent and alarm name for 5XX errors e.g. 500 Internal Server Error
     */
    FivexxErrorsAlarm = "5xxErrorsAlarm",
}

export interface WaBucketProps extends s3.BucketProps {
    /**
     * This flag is used to skip the default alarms.
     */
    readonly waDoNotAddDefaultAlarms?: boolean;
}

export class WaBucket extends s3.Bucket {

    constructor(scope: Construct, id: string, props: WaBucketProps) {
        super(scope, id, {
            ...props,
            encryption: props.encryption ?? BucketEncryption.S3_MANAGED,
            enforceSSL: props.enforceSSL ?? true,
            publicReadAccess: props.publicReadAccess ?? false,
            blockPublicAccess:
                props.blockPublicAccess ?? s3.BlockPublicAccess.BLOCK_ALL
        });

        this.add4xxErrorsAlarm();
        this.add5xxErrorsAlarm();
    }

    private add4xxErrorsAlarm(): void {
        const error4xxMetric = new Metric({
            namespace: "AWS/S3",
            metricName: "4xxErrors",
            statistic: "Sum",
            period: Duration.minutes(5)
        });

        new Alarm(this, WaBucketAlarms.FourxxErrorsAlarm, {
            alarmName: `${this.bucketName + WaBucketAlarms.FourxxErrorsAlarm}`,
            comparisonOperator: ComparisonOperator.GREATER_THAN_THRESHOLD,
            metric: error4xxMetric,
            alarmDescription: `S3 4xxError Rate Alarm for Bucket ${this.bucketName}`,
            threshold: 1,
            evaluationPeriods: 5
        });
    }

    private add5xxErrorsAlarm(): void {
        const error5xxMetric = new Metric({
            namespace: "AWS/S3",
            metricName: "5xxErrors",
            statistic: "Sum",
            period: Duration.minutes(5)
        });

        new Alarm(this, WaBucketAlarms.FivexxErrorsAlarm, {
            alarmName: `${this.bucketName + WaBucketAlarms.FivexxErrorsAlarm}`,
            comparisonOperator: ComparisonOperator.GREATER_THAN_THRESHOLD,
            metric: error5xxMetric,
            alarmDescription:
                `S3 5xxError Rate Alarm for Bucket ${this.bucketName}`,
            threshold: 1,
            evaluationPeriods: 5
        });
    }
}
