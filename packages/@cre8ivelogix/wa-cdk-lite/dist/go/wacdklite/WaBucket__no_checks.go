//go:build no_runtime_type_checking

// Well Architected CDK Constructs - Lite
package wacdklite

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WaBucket) validateAddCorsRuleParameters(rule *awss3.CorsRule) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateAddEventNotificationParameters(event awss3.EventType, dest awss3.IBucketNotificationDestination, filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateAddInventoryParameters(inventory *awss3.Inventory) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateAddLifecycleRuleParameters(rule *awss3.LifecycleRule) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateAddMetricParameters(metric *awss3.BucketMetrics) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateAddObjectCreatedNotificationParameters(dest awss3.IBucketNotificationDestination, filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateAddObjectRemovedNotificationParameters(dest awss3.IBucketNotificationDestination, filters *[]*awss3.NotificationKeyFilter) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateArnForObjectsParameters(keyPattern *string) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateGrantDeleteParameters(identity awsiam.IGrantable) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateGrantPutParameters(identity awsiam.IGrantable) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateGrantPutAclParameters(identity awsiam.IGrantable) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateGrantReadParameters(identity awsiam.IGrantable) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateGrantReadWriteParameters(identity awsiam.IGrantable) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateGrantWriteParameters(identity awsiam.IGrantable) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateOnCloudTrailEventParameters(id *string, options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateOnCloudTrailPutObjectParameters(id *string, options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateOnCloudTrailWriteObjectParameters(id *string, options *awss3.OnCloudTrailBucketEventOptions) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateTransferAccelerationUrlForObjectParameters(options *awss3.TransferAccelerationUrlOptions) error {
	return nil
}

func (w *jsiiProxy_WaBucket) validateVirtualHostedUrlForObjectParameters(options *awss3.VirtualHostedStyleUrlOptions) error {
	return nil
}

func validateWaBucket_FromBucketArnParameters(scope constructs.Construct, id *string, bucketArn *string) error {
	return nil
}

func validateWaBucket_FromBucketAttributesParameters(scope constructs.Construct, id *string, attrs *awss3.BucketAttributes) error {
	return nil
}

func validateWaBucket_FromBucketNameParameters(scope constructs.Construct, id *string, bucketName *string) error {
	return nil
}

func validateWaBucket_FromCfnBucketParameters(cfnBucket awss3.CfnBucket) error {
	return nil
}

func validateWaBucket_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWaBucket_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWaBucket_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWaBucket_ValidateBucketNameParameters(physicalName *string) error {
	return nil
}

func (j *jsiiProxy_WaBucket) validateSetAutoCreatePolicyParameters(val *bool) error {
	return nil
}

func validateNewWaBucketParameters(scope constructs.Construct, id *string, props *WaBucketProps) error {
	return nil
}

