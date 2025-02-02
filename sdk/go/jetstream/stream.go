// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package jetstream

import (
	"context"
	"reflect"

	"github.com/Triple-Whale/pulumi-jetstream/sdk/go/jetstream/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Stream struct {
	pulumi.CustomResourceState

	// If the Stream should support confirming receiving messages via acknowledgements
	Ack pulumi.BoolPtrOutput `pulumi:"ack"`
	// Allow higher performance, direct access to get individual messages via the $JS.DS.GET API
	AllowDirect pulumi.BoolPtrOutput `pulumi:"allowDirect"`
	// Allows the use of the Nats-Rollup header to replace all contents of a stream, or subject in a stream, with a single new message
	AllowRollupHdrs pulumi.BoolPtrOutput `pulumi:"allowRollupHdrs"`
	// Optional compression algorithm used for the Stream
	Compression pulumi.StringPtrOutput `pulumi:"compression"`
	// Restricts the ability to delete messages from a stream via the API. Cannot be changed once set to true
	DenyDelete pulumi.BoolPtrOutput `pulumi:"denyDelete"`
	// Restricts the ability to purge messages from a stream via the API. Cannot be change once set to true
	DenyPurge pulumi.BoolPtrOutput `pulumi:"denyPurge"`
	// Contains additional information about this stream
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// When a Stream reach it's limits either old messages are deleted or new ones are denied
	Discard pulumi.StringPtrOutput `pulumi:"discard"`
	// When discard policy is new and the stream is one with max messages per subject set, this will apply the new behavior to every subject. Essentially turning discard new from maximum number of subjects into maximum number of messages in a subject
	DiscardNewPerSubject pulumi.BoolPtrOutput `pulumi:"discardNewPerSubject"`
	// The size of the duplicate tracking windows, duration specified in seconds
	DuplicateWindow pulumi.IntPtrOutput `pulumi:"duplicateWindow"`
	// The maximum oldest message that can be kept in the stream, duration specified in seconds
	MaxAge pulumi.IntPtrOutput `pulumi:"maxAge"`
	// The maximum size of all messages that can be kept in the stream
	MaxBytes pulumi.IntPtrOutput `pulumi:"maxBytes"`
	// Number of consumers this stream allows
	MaxConsumers pulumi.IntPtrOutput `pulumi:"maxConsumers"`
	// The maximum individual message size that the stream will accept
	MaxMsgSize pulumi.IntPtrOutput `pulumi:"maxMsgSize"`
	// The maximum amount of messages that can be kept in the stream
	MaxMsgs pulumi.IntPtrOutput `pulumi:"maxMsgs"`
	// The maximum amount of messages that can be kept in the stream on a per-subject basis
	MaxMsgsPerSubject pulumi.IntPtrOutput `pulumi:"maxMsgsPerSubject"`
	// Free form metadata about the stream
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// Specifies a remote stream to mirror into this one
	Mirror StreamMirrorPtrOutput `pulumi:"mirror"`
	// The name of the source Stream
	Name pulumi.StringOutput `pulumi:"name"`
	// Place the stream in a specific cluster, influenced by placement_tags
	PlacementCluster pulumi.StringPtrOutput `pulumi:"placementCluster"`
	// Place the stream only on servers with these tags
	PlacementTags pulumi.StringArrayOutput `pulumi:"placementTags"`
	// How many replicas of the data to keep in a clustered environment
	Replicas pulumi.IntPtrOutput `pulumi:"replicas"`
	// The destination to publish messages to
	RepublishDestination pulumi.StringPtrOutput `pulumi:"republishDestination"`
	// Republish only message headers, no bodies
	RepublishHeadersOnly pulumi.BoolPtrOutput `pulumi:"republishHeadersOnly"`
	// Republish messages to republish_destination
	RepublishSource pulumi.StringPtrOutput `pulumi:"republishSource"`
	// The retention policy to apply over and above max*msgs, max*bytes and max_age
	Retention pulumi.StringPtrOutput `pulumi:"retention"`
	// The subject transform source
	Sources StreamSourceArrayOutput `pulumi:"sources"`
	// The storage engine to use to back the stream
	Storage pulumi.StringPtrOutput `pulumi:"storage"`
	// The subject filtering sources and associated destination transforms
	SubjectTransform StreamSubjectTransformPtrOutput `pulumi:"subjectTransform"`
	// The list of subjects that will be consumed by the Stream, may be empty when sources and mirrors are present
	Subjects pulumi.StringArrayOutput `pulumi:"subjects"`
}

// NewStream registers a new resource with the given unique name, arguments, and options.
func NewStream(ctx *pulumi.Context,
	name string, args *StreamArgs, opts ...pulumi.ResourceOption) (*Stream, error) {
	if args == nil {
		args = &StreamArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Stream
	err := ctx.RegisterResource("jetstream:index/stream:Stream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStream gets an existing Stream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamState, opts ...pulumi.ResourceOption) (*Stream, error) {
	var resource Stream
	err := ctx.ReadResource("jetstream:index/stream:Stream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stream resources.
type streamState struct {
	// If the Stream should support confirming receiving messages via acknowledgements
	Ack *bool `pulumi:"ack"`
	// Allow higher performance, direct access to get individual messages via the $JS.DS.GET API
	AllowDirect *bool `pulumi:"allowDirect"`
	// Allows the use of the Nats-Rollup header to replace all contents of a stream, or subject in a stream, with a single new message
	AllowRollupHdrs *bool `pulumi:"allowRollupHdrs"`
	// Optional compression algorithm used for the Stream
	Compression *string `pulumi:"compression"`
	// Restricts the ability to delete messages from a stream via the API. Cannot be changed once set to true
	DenyDelete *bool `pulumi:"denyDelete"`
	// Restricts the ability to purge messages from a stream via the API. Cannot be change once set to true
	DenyPurge *bool `pulumi:"denyPurge"`
	// Contains additional information about this stream
	Description *string `pulumi:"description"`
	// When a Stream reach it's limits either old messages are deleted or new ones are denied
	Discard *string `pulumi:"discard"`
	// When discard policy is new and the stream is one with max messages per subject set, this will apply the new behavior to every subject. Essentially turning discard new from maximum number of subjects into maximum number of messages in a subject
	DiscardNewPerSubject *bool `pulumi:"discardNewPerSubject"`
	// The size of the duplicate tracking windows, duration specified in seconds
	DuplicateWindow *int `pulumi:"duplicateWindow"`
	// The maximum oldest message that can be kept in the stream, duration specified in seconds
	MaxAge *int `pulumi:"maxAge"`
	// The maximum size of all messages that can be kept in the stream
	MaxBytes *int `pulumi:"maxBytes"`
	// Number of consumers this stream allows
	MaxConsumers *int `pulumi:"maxConsumers"`
	// The maximum individual message size that the stream will accept
	MaxMsgSize *int `pulumi:"maxMsgSize"`
	// The maximum amount of messages that can be kept in the stream
	MaxMsgs *int `pulumi:"maxMsgs"`
	// The maximum amount of messages that can be kept in the stream on a per-subject basis
	MaxMsgsPerSubject *int `pulumi:"maxMsgsPerSubject"`
	// Free form metadata about the stream
	Metadata map[string]string `pulumi:"metadata"`
	// Specifies a remote stream to mirror into this one
	Mirror *StreamMirror `pulumi:"mirror"`
	// The name of the source Stream
	Name *string `pulumi:"name"`
	// Place the stream in a specific cluster, influenced by placement_tags
	PlacementCluster *string `pulumi:"placementCluster"`
	// Place the stream only on servers with these tags
	PlacementTags []string `pulumi:"placementTags"`
	// How many replicas of the data to keep in a clustered environment
	Replicas *int `pulumi:"replicas"`
	// The destination to publish messages to
	RepublishDestination *string `pulumi:"republishDestination"`
	// Republish only message headers, no bodies
	RepublishHeadersOnly *bool `pulumi:"republishHeadersOnly"`
	// Republish messages to republish_destination
	RepublishSource *string `pulumi:"republishSource"`
	// The retention policy to apply over and above max*msgs, max*bytes and max_age
	Retention *string `pulumi:"retention"`
	// The subject transform source
	Sources []StreamSource `pulumi:"sources"`
	// The storage engine to use to back the stream
	Storage *string `pulumi:"storage"`
	// The subject filtering sources and associated destination transforms
	SubjectTransform *StreamSubjectTransform `pulumi:"subjectTransform"`
	// The list of subjects that will be consumed by the Stream, may be empty when sources and mirrors are present
	Subjects []string `pulumi:"subjects"`
}

type StreamState struct {
	// If the Stream should support confirming receiving messages via acknowledgements
	Ack pulumi.BoolPtrInput
	// Allow higher performance, direct access to get individual messages via the $JS.DS.GET API
	AllowDirect pulumi.BoolPtrInput
	// Allows the use of the Nats-Rollup header to replace all contents of a stream, or subject in a stream, with a single new message
	AllowRollupHdrs pulumi.BoolPtrInput
	// Optional compression algorithm used for the Stream
	Compression pulumi.StringPtrInput
	// Restricts the ability to delete messages from a stream via the API. Cannot be changed once set to true
	DenyDelete pulumi.BoolPtrInput
	// Restricts the ability to purge messages from a stream via the API. Cannot be change once set to true
	DenyPurge pulumi.BoolPtrInput
	// Contains additional information about this stream
	Description pulumi.StringPtrInput
	// When a Stream reach it's limits either old messages are deleted or new ones are denied
	Discard pulumi.StringPtrInput
	// When discard policy is new and the stream is one with max messages per subject set, this will apply the new behavior to every subject. Essentially turning discard new from maximum number of subjects into maximum number of messages in a subject
	DiscardNewPerSubject pulumi.BoolPtrInput
	// The size of the duplicate tracking windows, duration specified in seconds
	DuplicateWindow pulumi.IntPtrInput
	// The maximum oldest message that can be kept in the stream, duration specified in seconds
	MaxAge pulumi.IntPtrInput
	// The maximum size of all messages that can be kept in the stream
	MaxBytes pulumi.IntPtrInput
	// Number of consumers this stream allows
	MaxConsumers pulumi.IntPtrInput
	// The maximum individual message size that the stream will accept
	MaxMsgSize pulumi.IntPtrInput
	// The maximum amount of messages that can be kept in the stream
	MaxMsgs pulumi.IntPtrInput
	// The maximum amount of messages that can be kept in the stream on a per-subject basis
	MaxMsgsPerSubject pulumi.IntPtrInput
	// Free form metadata about the stream
	Metadata pulumi.StringMapInput
	// Specifies a remote stream to mirror into this one
	Mirror StreamMirrorPtrInput
	// The name of the source Stream
	Name pulumi.StringPtrInput
	// Place the stream in a specific cluster, influenced by placement_tags
	PlacementCluster pulumi.StringPtrInput
	// Place the stream only on servers with these tags
	PlacementTags pulumi.StringArrayInput
	// How many replicas of the data to keep in a clustered environment
	Replicas pulumi.IntPtrInput
	// The destination to publish messages to
	RepublishDestination pulumi.StringPtrInput
	// Republish only message headers, no bodies
	RepublishHeadersOnly pulumi.BoolPtrInput
	// Republish messages to republish_destination
	RepublishSource pulumi.StringPtrInput
	// The retention policy to apply over and above max*msgs, max*bytes and max_age
	Retention pulumi.StringPtrInput
	// The subject transform source
	Sources StreamSourceArrayInput
	// The storage engine to use to back the stream
	Storage pulumi.StringPtrInput
	// The subject filtering sources and associated destination transforms
	SubjectTransform StreamSubjectTransformPtrInput
	// The list of subjects that will be consumed by the Stream, may be empty when sources and mirrors are present
	Subjects pulumi.StringArrayInput
}

func (StreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamState)(nil)).Elem()
}

type streamArgs struct {
	// If the Stream should support confirming receiving messages via acknowledgements
	Ack *bool `pulumi:"ack"`
	// Allow higher performance, direct access to get individual messages via the $JS.DS.GET API
	AllowDirect *bool `pulumi:"allowDirect"`
	// Allows the use of the Nats-Rollup header to replace all contents of a stream, or subject in a stream, with a single new message
	AllowRollupHdrs *bool `pulumi:"allowRollupHdrs"`
	// Optional compression algorithm used for the Stream
	Compression *string `pulumi:"compression"`
	// Restricts the ability to delete messages from a stream via the API. Cannot be changed once set to true
	DenyDelete *bool `pulumi:"denyDelete"`
	// Restricts the ability to purge messages from a stream via the API. Cannot be change once set to true
	DenyPurge *bool `pulumi:"denyPurge"`
	// Contains additional information about this stream
	Description *string `pulumi:"description"`
	// When a Stream reach it's limits either old messages are deleted or new ones are denied
	Discard *string `pulumi:"discard"`
	// When discard policy is new and the stream is one with max messages per subject set, this will apply the new behavior to every subject. Essentially turning discard new from maximum number of subjects into maximum number of messages in a subject
	DiscardNewPerSubject *bool `pulumi:"discardNewPerSubject"`
	// The size of the duplicate tracking windows, duration specified in seconds
	DuplicateWindow *int `pulumi:"duplicateWindow"`
	// The maximum oldest message that can be kept in the stream, duration specified in seconds
	MaxAge *int `pulumi:"maxAge"`
	// The maximum size of all messages that can be kept in the stream
	MaxBytes *int `pulumi:"maxBytes"`
	// Number of consumers this stream allows
	MaxConsumers *int `pulumi:"maxConsumers"`
	// The maximum individual message size that the stream will accept
	MaxMsgSize *int `pulumi:"maxMsgSize"`
	// The maximum amount of messages that can be kept in the stream
	MaxMsgs *int `pulumi:"maxMsgs"`
	// The maximum amount of messages that can be kept in the stream on a per-subject basis
	MaxMsgsPerSubject *int `pulumi:"maxMsgsPerSubject"`
	// Free form metadata about the stream
	Metadata map[string]string `pulumi:"metadata"`
	// Specifies a remote stream to mirror into this one
	Mirror *StreamMirror `pulumi:"mirror"`
	// The name of the source Stream
	Name *string `pulumi:"name"`
	// Place the stream in a specific cluster, influenced by placement_tags
	PlacementCluster *string `pulumi:"placementCluster"`
	// Place the stream only on servers with these tags
	PlacementTags []string `pulumi:"placementTags"`
	// How many replicas of the data to keep in a clustered environment
	Replicas *int `pulumi:"replicas"`
	// The destination to publish messages to
	RepublishDestination *string `pulumi:"republishDestination"`
	// Republish only message headers, no bodies
	RepublishHeadersOnly *bool `pulumi:"republishHeadersOnly"`
	// Republish messages to republish_destination
	RepublishSource *string `pulumi:"republishSource"`
	// The retention policy to apply over and above max*msgs, max*bytes and max_age
	Retention *string `pulumi:"retention"`
	// The subject transform source
	Sources []StreamSource `pulumi:"sources"`
	// The storage engine to use to back the stream
	Storage *string `pulumi:"storage"`
	// The subject filtering sources and associated destination transforms
	SubjectTransform *StreamSubjectTransform `pulumi:"subjectTransform"`
	// The list of subjects that will be consumed by the Stream, may be empty when sources and mirrors are present
	Subjects []string `pulumi:"subjects"`
}

// The set of arguments for constructing a Stream resource.
type StreamArgs struct {
	// If the Stream should support confirming receiving messages via acknowledgements
	Ack pulumi.BoolPtrInput
	// Allow higher performance, direct access to get individual messages via the $JS.DS.GET API
	AllowDirect pulumi.BoolPtrInput
	// Allows the use of the Nats-Rollup header to replace all contents of a stream, or subject in a stream, with a single new message
	AllowRollupHdrs pulumi.BoolPtrInput
	// Optional compression algorithm used for the Stream
	Compression pulumi.StringPtrInput
	// Restricts the ability to delete messages from a stream via the API. Cannot be changed once set to true
	DenyDelete pulumi.BoolPtrInput
	// Restricts the ability to purge messages from a stream via the API. Cannot be change once set to true
	DenyPurge pulumi.BoolPtrInput
	// Contains additional information about this stream
	Description pulumi.StringPtrInput
	// When a Stream reach it's limits either old messages are deleted or new ones are denied
	Discard pulumi.StringPtrInput
	// When discard policy is new and the stream is one with max messages per subject set, this will apply the new behavior to every subject. Essentially turning discard new from maximum number of subjects into maximum number of messages in a subject
	DiscardNewPerSubject pulumi.BoolPtrInput
	// The size of the duplicate tracking windows, duration specified in seconds
	DuplicateWindow pulumi.IntPtrInput
	// The maximum oldest message that can be kept in the stream, duration specified in seconds
	MaxAge pulumi.IntPtrInput
	// The maximum size of all messages that can be kept in the stream
	MaxBytes pulumi.IntPtrInput
	// Number of consumers this stream allows
	MaxConsumers pulumi.IntPtrInput
	// The maximum individual message size that the stream will accept
	MaxMsgSize pulumi.IntPtrInput
	// The maximum amount of messages that can be kept in the stream
	MaxMsgs pulumi.IntPtrInput
	// The maximum amount of messages that can be kept in the stream on a per-subject basis
	MaxMsgsPerSubject pulumi.IntPtrInput
	// Free form metadata about the stream
	Metadata pulumi.StringMapInput
	// Specifies a remote stream to mirror into this one
	Mirror StreamMirrorPtrInput
	// The name of the source Stream
	Name pulumi.StringPtrInput
	// Place the stream in a specific cluster, influenced by placement_tags
	PlacementCluster pulumi.StringPtrInput
	// Place the stream only on servers with these tags
	PlacementTags pulumi.StringArrayInput
	// How many replicas of the data to keep in a clustered environment
	Replicas pulumi.IntPtrInput
	// The destination to publish messages to
	RepublishDestination pulumi.StringPtrInput
	// Republish only message headers, no bodies
	RepublishHeadersOnly pulumi.BoolPtrInput
	// Republish messages to republish_destination
	RepublishSource pulumi.StringPtrInput
	// The retention policy to apply over and above max*msgs, max*bytes and max_age
	Retention pulumi.StringPtrInput
	// The subject transform source
	Sources StreamSourceArrayInput
	// The storage engine to use to back the stream
	Storage pulumi.StringPtrInput
	// The subject filtering sources and associated destination transforms
	SubjectTransform StreamSubjectTransformPtrInput
	// The list of subjects that will be consumed by the Stream, may be empty when sources and mirrors are present
	Subjects pulumi.StringArrayInput
}

func (StreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamArgs)(nil)).Elem()
}

type StreamInput interface {
	pulumi.Input

	ToStreamOutput() StreamOutput
	ToStreamOutputWithContext(ctx context.Context) StreamOutput
}

func (*Stream) ElementType() reflect.Type {
	return reflect.TypeOf((**Stream)(nil)).Elem()
}

func (i *Stream) ToStreamOutput() StreamOutput {
	return i.ToStreamOutputWithContext(context.Background())
}

func (i *Stream) ToStreamOutputWithContext(ctx context.Context) StreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamOutput)
}

// StreamArrayInput is an input type that accepts StreamArray and StreamArrayOutput values.
// You can construct a concrete instance of `StreamArrayInput` via:
//
//	StreamArray{ StreamArgs{...} }
type StreamArrayInput interface {
	pulumi.Input

	ToStreamArrayOutput() StreamArrayOutput
	ToStreamArrayOutputWithContext(context.Context) StreamArrayOutput
}

type StreamArray []StreamInput

func (StreamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Stream)(nil)).Elem()
}

func (i StreamArray) ToStreamArrayOutput() StreamArrayOutput {
	return i.ToStreamArrayOutputWithContext(context.Background())
}

func (i StreamArray) ToStreamArrayOutputWithContext(ctx context.Context) StreamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamArrayOutput)
}

// StreamMapInput is an input type that accepts StreamMap and StreamMapOutput values.
// You can construct a concrete instance of `StreamMapInput` via:
//
//	StreamMap{ "key": StreamArgs{...} }
type StreamMapInput interface {
	pulumi.Input

	ToStreamMapOutput() StreamMapOutput
	ToStreamMapOutputWithContext(context.Context) StreamMapOutput
}

type StreamMap map[string]StreamInput

func (StreamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Stream)(nil)).Elem()
}

func (i StreamMap) ToStreamMapOutput() StreamMapOutput {
	return i.ToStreamMapOutputWithContext(context.Background())
}

func (i StreamMap) ToStreamMapOutputWithContext(ctx context.Context) StreamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamMapOutput)
}

type StreamOutput struct{ *pulumi.OutputState }

func (StreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Stream)(nil)).Elem()
}

func (o StreamOutput) ToStreamOutput() StreamOutput {
	return o
}

func (o StreamOutput) ToStreamOutputWithContext(ctx context.Context) StreamOutput {
	return o
}

// If the Stream should support confirming receiving messages via acknowledgements
func (o StreamOutput) Ack() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.BoolPtrOutput { return v.Ack }).(pulumi.BoolPtrOutput)
}

// Allow higher performance, direct access to get individual messages via the $JS.DS.GET API
func (o StreamOutput) AllowDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.BoolPtrOutput { return v.AllowDirect }).(pulumi.BoolPtrOutput)
}

// Allows the use of the Nats-Rollup header to replace all contents of a stream, or subject in a stream, with a single new message
func (o StreamOutput) AllowRollupHdrs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.BoolPtrOutput { return v.AllowRollupHdrs }).(pulumi.BoolPtrOutput)
}

// Optional compression algorithm used for the Stream
func (o StreamOutput) Compression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringPtrOutput { return v.Compression }).(pulumi.StringPtrOutput)
}

// Restricts the ability to delete messages from a stream via the API. Cannot be changed once set to true
func (o StreamOutput) DenyDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.BoolPtrOutput { return v.DenyDelete }).(pulumi.BoolPtrOutput)
}

// Restricts the ability to purge messages from a stream via the API. Cannot be change once set to true
func (o StreamOutput) DenyPurge() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.BoolPtrOutput { return v.DenyPurge }).(pulumi.BoolPtrOutput)
}

// Contains additional information about this stream
func (o StreamOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When a Stream reach it's limits either old messages are deleted or new ones are denied
func (o StreamOutput) Discard() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringPtrOutput { return v.Discard }).(pulumi.StringPtrOutput)
}

// When discard policy is new and the stream is one with max messages per subject set, this will apply the new behavior to every subject. Essentially turning discard new from maximum number of subjects into maximum number of messages in a subject
func (o StreamOutput) DiscardNewPerSubject() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.BoolPtrOutput { return v.DiscardNewPerSubject }).(pulumi.BoolPtrOutput)
}

// The size of the duplicate tracking windows, duration specified in seconds
func (o StreamOutput) DuplicateWindow() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.IntPtrOutput { return v.DuplicateWindow }).(pulumi.IntPtrOutput)
}

// The maximum oldest message that can be kept in the stream, duration specified in seconds
func (o StreamOutput) MaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.IntPtrOutput { return v.MaxAge }).(pulumi.IntPtrOutput)
}

// The maximum size of all messages that can be kept in the stream
func (o StreamOutput) MaxBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.IntPtrOutput { return v.MaxBytes }).(pulumi.IntPtrOutput)
}

// Number of consumers this stream allows
func (o StreamOutput) MaxConsumers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.IntPtrOutput { return v.MaxConsumers }).(pulumi.IntPtrOutput)
}

// The maximum individual message size that the stream will accept
func (o StreamOutput) MaxMsgSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.IntPtrOutput { return v.MaxMsgSize }).(pulumi.IntPtrOutput)
}

// The maximum amount of messages that can be kept in the stream
func (o StreamOutput) MaxMsgs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.IntPtrOutput { return v.MaxMsgs }).(pulumi.IntPtrOutput)
}

// The maximum amount of messages that can be kept in the stream on a per-subject basis
func (o StreamOutput) MaxMsgsPerSubject() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.IntPtrOutput { return v.MaxMsgsPerSubject }).(pulumi.IntPtrOutput)
}

// Free form metadata about the stream
func (o StreamOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// Specifies a remote stream to mirror into this one
func (o StreamOutput) Mirror() StreamMirrorPtrOutput {
	return o.ApplyT(func(v *Stream) StreamMirrorPtrOutput { return v.Mirror }).(StreamMirrorPtrOutput)
}

// The name of the source Stream
func (o StreamOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Place the stream in a specific cluster, influenced by placement_tags
func (o StreamOutput) PlacementCluster() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringPtrOutput { return v.PlacementCluster }).(pulumi.StringPtrOutput)
}

// Place the stream only on servers with these tags
func (o StreamOutput) PlacementTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringArrayOutput { return v.PlacementTags }).(pulumi.StringArrayOutput)
}

// How many replicas of the data to keep in a clustered environment
func (o StreamOutput) Replicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.IntPtrOutput { return v.Replicas }).(pulumi.IntPtrOutput)
}

// The destination to publish messages to
func (o StreamOutput) RepublishDestination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringPtrOutput { return v.RepublishDestination }).(pulumi.StringPtrOutput)
}

// Republish only message headers, no bodies
func (o StreamOutput) RepublishHeadersOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.BoolPtrOutput { return v.RepublishHeadersOnly }).(pulumi.BoolPtrOutput)
}

// Republish messages to republish_destination
func (o StreamOutput) RepublishSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringPtrOutput { return v.RepublishSource }).(pulumi.StringPtrOutput)
}

// The retention policy to apply over and above max*msgs, max*bytes and max_age
func (o StreamOutput) Retention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringPtrOutput { return v.Retention }).(pulumi.StringPtrOutput)
}

// The subject transform source
func (o StreamOutput) Sources() StreamSourceArrayOutput {
	return o.ApplyT(func(v *Stream) StreamSourceArrayOutput { return v.Sources }).(StreamSourceArrayOutput)
}

// The storage engine to use to back the stream
func (o StreamOutput) Storage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringPtrOutput { return v.Storage }).(pulumi.StringPtrOutput)
}

// The subject filtering sources and associated destination transforms
func (o StreamOutput) SubjectTransform() StreamSubjectTransformPtrOutput {
	return o.ApplyT(func(v *Stream) StreamSubjectTransformPtrOutput { return v.SubjectTransform }).(StreamSubjectTransformPtrOutput)
}

// The list of subjects that will be consumed by the Stream, may be empty when sources and mirrors are present
func (o StreamOutput) Subjects() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringArrayOutput { return v.Subjects }).(pulumi.StringArrayOutput)
}

type StreamArrayOutput struct{ *pulumi.OutputState }

func (StreamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Stream)(nil)).Elem()
}

func (o StreamArrayOutput) ToStreamArrayOutput() StreamArrayOutput {
	return o
}

func (o StreamArrayOutput) ToStreamArrayOutputWithContext(ctx context.Context) StreamArrayOutput {
	return o
}

func (o StreamArrayOutput) Index(i pulumi.IntInput) StreamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Stream {
		return vs[0].([]*Stream)[vs[1].(int)]
	}).(StreamOutput)
}

type StreamMapOutput struct{ *pulumi.OutputState }

func (StreamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Stream)(nil)).Elem()
}

func (o StreamMapOutput) ToStreamMapOutput() StreamMapOutput {
	return o
}

func (o StreamMapOutput) ToStreamMapOutputWithContext(ctx context.Context) StreamMapOutput {
	return o
}

func (o StreamMapOutput) MapIndex(k pulumi.StringInput) StreamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Stream {
		return vs[0].(map[string]*Stream)[vs[1].(string)]
	}).(StreamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StreamInput)(nil)).Elem(), &Stream{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamArrayInput)(nil)).Elem(), StreamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StreamMapInput)(nil)).Elem(), StreamMap{})
	pulumi.RegisterOutputType(StreamOutput{})
	pulumi.RegisterOutputType(StreamArrayOutput{})
	pulumi.RegisterOutputType(StreamMapOutput{})
}
