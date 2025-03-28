// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package v1alpha1

const (
	// AnnotationPrefix is the prefix for all ACK annotations
	AnnotationPrefix = "services.k8s.aws/"
	// AnnotationAdopted is an annotation whose value is a boolean value,
	// If this annotation is set to true on a CR, that means the user is
	// indicating to the ACK service controller that it should expect a backend
	// AWS service API resource to already exist (and that ACK should "adopt"
	// the resource into its management). If this annotation is set to false on
	// a CR, that means the user expects the ACK service controller to create
	// the backend AWS service API resource.
	AnnotationAdopted = AnnotationPrefix + "adopted"
	// AnnotationOwnerAccountID is an annotation whose value is the identifier
	// for the AWS account to which the resource belongs.  If this annotation
	// is set on a CR, the Kubernetes user is indicating that the ACK service
	// controller should create/patch/delete the resource in the specified AWS
	// Account. In order for this cross-account resource management to succeed,
	// the AWS IAM Role that the ACK service controller runs as needs to have
	// the ability to call the AWS STS::AssumeRole API call and assume an IAM
	// Role in the target AWS Account.
	// TODO(jaypipes): Link to documentation on cross-account resource
	// management
	AnnotationOwnerAccountID = AnnotationPrefix + "owner-account-id"
	// AnnotationTeamID is an annotation whose value is the identifier
	// for the AWS team ID to manage the resources.  If this annotation
	// is set on a CR, the Kubernetes user is indicating that the ACK service
	// controller should create/patch/delete the resource in the specified AWS
	// role for this team ID. In order for this cross-account resource management
	// to succeed, the AWS IAM Role that the ACK service controller runs as needs
	// to have the ability to call the AWS STS::AssumeRole API call and assume an
	// IAM Role in the target AWS Account.
	AnnotationTeamID = AnnotationPrefix + "team-id"
	// AnnotationRegion is an annotation whose value is the identifier for the
	// the AWS region in which the resources should be created. If this annotation
	// is set on a CR metadata, that means the user is indicating to the ACK service
	// controller that the CR should be created on specific region. ACK service
	// controller will not override the resource region if this annotation is set.
	AnnotationRegion = AnnotationPrefix + "region"
	// AnnotationDefaultRegion is an annotation whose value is the identifier
	// for the default AWS region in which resources should be created. If this
	// annotation is set on a namespace, the Kubernetes user is indicating that
	// the ACK service controller should set the regions in which the resource
	// should be created, if a region annotation is not set on the CR metadata.
	// If this annotation - and AnnotationRegion - are not set, ACK service
	// controllers look for controller binary flags and environment variables
	// injected by POD IRSA, to decide in which region the resources should be
	// created.
	AnnotationDefaultRegion = AnnotationPrefix + "default-region"
	// AnnotationEndpointURL is an annotation whose value is the identifier
	// for the AWS endpoint in which the service controller will use to create
	// its resources. If this annotation is set on a namespace, the Kubernetes user
	// is indicating that the ACK service controller should create its resources using
	// that specific endpoint. If this annotation is not set, ACK service controller
	// will either use the default behavior	of aws-sdk-go to create endpoints or
	// aws-endpoint-url if it is set in controller binary flags and environment variables.
	AnnotationEndpointURL = AnnotationPrefix + "endpoint-url"
	// AnnotationDeletionPolicy is an annotation whose value is the identifier for the
	// the deletion policy for the current resource. If this annotation is set
	// to "delete" the resource manager will delete the AWS resource when the
	// K8s resource is deleted. If this annotation is set to "retain" the
	// resource manager will leave the AWS resource intact when the K8s resource
	// is deleted.
	AnnotationDeletionPolicy = AnnotationPrefix + "deletion-policy"
	// AnnotationReadOnly is an annotation whose value is a boolean indicating
	// whether the resource is read-only. If this annotation is set to true on a
	// CR, that means the user is indicating to the ACK service controller that
	// the resource is read-only and should not be created/patched/deleted by the
	// ACK service controller.
	AnnotationReadOnly = AnnotationPrefix + "read-only"
	// AnnotationAdoptionPolicy is an annotation whose value is the identifier for whether
	// we will attempt adoption only (value = adopt-only) or attempt a create if resource 
	// is not found (value adopt-or-create).
	//
	// NOTE (michaelhtm): Currently create-or-adopt is not supported
	AnnotationAdoptionPolicy = AnnotationPrefix + "adoption-policy"
	// AnnotationAdoptionFields is an annotation whose value contains a json-like
	// format of the requied fields to do a ReadOne when attempting to force-adopt
	// a Resource
	AnnotationAdoptionFields = AnnotationPrefix + "adoption-fields"
)
