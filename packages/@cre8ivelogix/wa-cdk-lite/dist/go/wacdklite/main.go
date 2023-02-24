package wacdklite

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cre8ivelogix/wa-cdk-lite.WaBucket",
		reflect.TypeOf((*WaBucket)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addCorsRule", GoMethod: "AddCorsRule"},
			_jsii_.MemberMethod{JsiiMethod: "addEventNotification", GoMethod: "AddEventNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addInventory", GoMethod: "AddInventory"},
			_jsii_.MemberMethod{JsiiMethod: "addLifecycleRule", GoMethod: "AddLifecycleRule"},
			_jsii_.MemberMethod{JsiiMethod: "addMetric", GoMethod: "AddMetric"},
			_jsii_.MemberMethod{JsiiMethod: "addObjectCreatedNotification", GoMethod: "AddObjectCreatedNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addObjectRemovedNotification", GoMethod: "AddObjectRemovedNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "arnForObjects", GoMethod: "ArnForObjects"},
			_jsii_.MemberProperty{JsiiProperty: "autoCreatePolicy", GoGetter: "AutoCreatePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketDomainName", GoGetter: "BucketDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketDualStackDomainName", GoGetter: "BucketDualStackDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketRegionalDomainName", GoGetter: "BucketRegionalDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketWebsiteDomainName", GoGetter: "BucketWebsiteDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketWebsiteUrl", GoGetter: "BucketWebsiteUrl"},
			_jsii_.MemberProperty{JsiiProperty: "disallowPublicAccess", GoGetter: "DisallowPublicAccess"},
			_jsii_.MemberMethod{JsiiMethod: "enableEventBridgeNotification", GoMethod: "EnableEventBridgeNotification"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantDelete", GoMethod: "GrantDelete"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublicAccess", GoMethod: "GrantPublicAccess"},
			_jsii_.MemberMethod{JsiiMethod: "grantPut", GoMethod: "GrantPut"},
			_jsii_.MemberMethod{JsiiMethod: "grantPutAcl", GoMethod: "GrantPutAcl"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadWrite", GoMethod: "GrantReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "isWebsite", GoGetter: "IsWebsite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationsHandlerRole", GoGetter: "NotificationsHandlerRole"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailEvent", GoMethod: "OnCloudTrailEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailPutObject", GoMethod: "OnCloudTrailPutObject"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailWriteObject", GoMethod: "OnCloudTrailWriteObject"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberMethod{JsiiMethod: "s3UrlForObject", GoMethod: "S3UrlForObject"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "transferAccelerationUrlForObject", GoMethod: "TransferAccelerationUrlForObject"},
			_jsii_.MemberMethod{JsiiMethod: "urlForObject", GoMethod: "UrlForObject"},
			_jsii_.MemberMethod{JsiiMethod: "virtualHostedUrlForObject", GoMethod: "VirtualHostedUrlForObject"},
		},
		func() interface{} {
			j := jsiiProxy_WaBucket{}
			_jsii_.InitJsiiProxy(&j.Type__awss3Bucket)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@cre8ivelogix/wa-cdk-lite.WaBucketAlarms",
		reflect.TypeOf((*WaBucketAlarms)(nil)).Elem(),
		map[string]interface{}{
			"FOURXX_ERRORS_ALARM": WaBucketAlarms_FOURXX_ERRORS_ALARM,
			"FIVEXX_ERRORS_ALARM": WaBucketAlarms_FIVEXX_ERRORS_ALARM,
		},
	)
	_jsii_.RegisterStruct(
		"@cre8ivelogix/wa-cdk-lite.WaBucketProps",
		reflect.TypeOf((*WaBucketProps)(nil)).Elem(),
	)
}
