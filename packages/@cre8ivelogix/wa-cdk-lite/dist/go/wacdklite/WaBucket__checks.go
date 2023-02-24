//go:build !no_runtime_type_checking

// Well Architected CDK Constructs - Lite
package wacdklite

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

func (w *jsiiProxy_WaBucket) validateAddCorsRuleParameters(rule *awss3.CorsRule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(rule, func() string { return "parameter rule" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateAddEventNotificationParameters(event awss3.EventType, dest awss3.IBucketNotificationDestination, filters *[]*awss3.NotificationKeyFilter) error {
	if event == "" {
		return fmt.Errorf("parameter event is required, but nil was provided")
	}

	if dest == nil {
		return fmt.Errorf("parameter dest is required, but nil was provided")
	}

	for idx_47e227, v := range *filters {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter filters[%#v]", idx_47e227) }); err != nil {
			return err
		}
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateAddInventoryParameters(inventory *awss3.Inventory) error {
	if inventory == nil {
		return fmt.Errorf("parameter inventory is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(inventory, func() string { return "parameter inventory" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateAddLifecycleRuleParameters(rule *awss3.LifecycleRule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(rule, func() string { return "parameter rule" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateAddMetricParameters(metric *awss3.BucketMetrics) error {
	if metric == nil {
		return fmt.Errorf("parameter metric is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(metric, func() string { return "parameter metric" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateAddObjectCreatedNotificationParameters(dest awss3.IBucketNotificationDestination, filters *[]*awss3.NotificationKeyFilter) error {
	if dest == nil {
		return fmt.Errorf("parameter dest is required, but nil was provided")
	}

	for idx_47e227, v := range *filters {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter filters[%#v]", idx_47e227) }); err != nil {
			return err
		}
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateAddObjectRemovedNotificationParameters(dest awss3.IBucketNotificationDestination, filters *[]*awss3.NotificationKeyFilter) error {
	if dest == nil {
		return fmt.Errorf("parameter dest is required, but nil was provided")
	}

	for idx_47e227, v := range *filters {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter filters[%#v]", idx_47e227) }); err != nil {
			return err
		}
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	if permission == nil {
		return fmt.Errorf("parameter permission is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateArnForObjectsParameters(keyPattern *string) error {
	if keyPattern == nil {
		return fmt.Errorf("parameter keyPattern is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	if arnAttr == nil {
		return fmt.Errorf("parameter arnAttr is required, but nil was provided")
	}

	if arnComponents == nil {
		return fmt.Errorf("parameter arnComponents is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(arnComponents, func() string { return "parameter arnComponents" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateGrantDeleteParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateGrantPutParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateGrantPutAclParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateGrantReadParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateGrantReadWriteParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateGrantWriteParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateOnCloudTrailEventParameters(id *string, options *awss3.OnCloudTrailBucketEventOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateOnCloudTrailPutObjectParameters(id *string, options *awss3.OnCloudTrailBucketEventOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateOnCloudTrailWriteObjectParameters(id *string, options *awss3.OnCloudTrailBucketEventOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateTransferAccelerationUrlForObjectParameters(options *awss3.TransferAccelerationUrlOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WaBucket) validateVirtualHostedUrlForObjectParameters(options *awss3.VirtualHostedStyleUrlOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateWaBucket_FromBucketArnParameters(scope constructs.Construct, id *string, bucketArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if bucketArn == nil {
		return fmt.Errorf("parameter bucketArn is required, but nil was provided")
	}

	return nil
}

func validateWaBucket_FromBucketAttributesParameters(scope constructs.Construct, id *string, attrs *awss3.BucketAttributes) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attrs == nil {
		return fmt.Errorf("parameter attrs is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attrs, func() string { return "parameter attrs" }); err != nil {
		return err
	}

	return nil
}

func validateWaBucket_FromBucketNameParameters(scope constructs.Construct, id *string, bucketName *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if bucketName == nil {
		return fmt.Errorf("parameter bucketName is required, but nil was provided")
	}

	return nil
}

func validateWaBucket_FromCfnBucketParameters(cfnBucket awss3.CfnBucket) error {
	if cfnBucket == nil {
		return fmt.Errorf("parameter cfnBucket is required, but nil was provided")
	}

	return nil
}

func validateWaBucket_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateWaBucket_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateWaBucket_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateWaBucket_ValidateBucketNameParameters(physicalName *string) error {
	if physicalName == nil {
		return fmt.Errorf("parameter physicalName is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_WaBucket) validateSetAutoCreatePolicyParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewWaBucketParameters(scope constructs.Construct, id *string, props *WaBucketProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

