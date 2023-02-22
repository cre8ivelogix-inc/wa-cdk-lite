import * as s3 from "aws-cdk-lib/aws-s3";
import {BucketEncryption} from "aws-cdk-lib/aws-s3";
import {Construct} from "constructs";
import {Duration} from "aws-cdk-lib";
import {Alarm, ComparisonOperator, Metric} from "aws-cdk-lib/aws-cloudwatch";

/**
 * Well Architected Bucket alarms
 */
export enum WaBucketAlarms {
    /**
     * This enum is used to represent an alarm name for 4XX errors e.g. 404 Not Found
     */
    FOURXX_ERRORS_ALARM = "4xxErrorsAlarm",
    /**
     * This enum is used to represent an alarm name for 5XX errors e.g. 500 Internal Server Error
     */
    FIVEXX_ERRORS_ALARM = "5xxErrorsAlarm",
}

/**
 * Well Architected Bucket Properties
 */
export interface WaBucketProps extends s3.BucketProps {
    /**
     * This flag is used to skip the default alarms.
     */
    readonly waDoNotAddDefaultAlarms?: boolean;
}

/**
 * Well Architected S3 Bucket that uses S3_MANAGED encryption, enforces ssl, denies public access
 * is versioned, and uses CloudFront distribution with bucket as origin.
 *
 * ### Default Alarms
 * By default it is configured with FourxxErrorsAlarm and FivexxErrorsAlarm alarms.
 *
 * Note that in freemium version the default alarm does not have any action. However, in premium version
 * it uses the Well Architected Alarm construct, which sets up an alarm action to notify the SNS
 * Topic *AlarmEventsTopic* by default.
 *
 * ##Example
 * Default Usage
 * ```ts
 * new WaBucket(this, "LogicalId", {});
 * ```
 *
 * ##Example
 * Custom Configuration
 * ```ts
 * new WaBucket(this, "LogicalId", {
 *    enforceSSL: false
 * });
 * ```
 *
 * ### Compliance
 * It addresses the following compliance requirements
 * * Blocks public access
 * > PCI, HIPAA, GDPR, APRA, MAS, NIST4
 * * Bucket versioning enabled
 * > PCI, APRA, MAS, NIST4
 * * Only allow secure transport protocols
 * > PCI, APRA, MAS, NIST4
 * * Server side encryption
 * > PCI, HIPAA, GDPR, APRA, MAS, NIST4
 * */
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

        new Alarm(this, WaBucketAlarms.FOURXX_ERRORS_ALARM, {
            alarmName: `${this.bucketName + WaBucketAlarms.FOURXX_ERRORS_ALARM}`,
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

        new Alarm(this, WaBucketAlarms.FIVEXX_ERRORS_ALARM, {
            alarmName: `${this.bucketName + WaBucketAlarms.FIVEXX_ERRORS_ALARM}`,
            comparisonOperator: ComparisonOperator.GREATER_THAN_THRESHOLD,
            metric: error5xxMetric,
            alarmDescription:
                `S3 5xxError Rate Alarm for Bucket ${this.bucketName}`,
            threshold: 1,
            evaluationPeriods: 5
        });
    }
}
