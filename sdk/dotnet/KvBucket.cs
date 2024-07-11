// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Jetstream
{
    [JetstreamResourceType("jetstream:index/kvBucket:KvBucket")]
    public partial class KvBucket : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Contains additional information about this bucket
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// How many historical values to keep
        /// </summary>
        [Output("history")]
        public Output<int?> History { get; private set; } = null!;

        /// <summary>
        /// Maximum size of the entire bucket
        /// </summary>
        [Output("maxBucketSize")]
        public Output<int?> MaxBucketSize { get; private set; } = null!;

        /// <summary>
        /// Maximum size of any value
        /// </summary>
        [Output("maxValueSize")]
        public Output<int?> MaxValueSize { get; private set; } = null!;

        /// <summary>
        /// The name of the Bucket
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Place the bucket in a specific cluster, influenced by placement_tags
        /// </summary>
        [Output("placementCluster")]
        public Output<string?> PlacementCluster { get; private set; } = null!;

        /// <summary>
        /// Place the stream only on servers with these tags
        /// </summary>
        [Output("placementTags")]
        public Output<ImmutableArray<string>> PlacementTags { get; private set; } = null!;

        /// <summary>
        /// Number of cluster replicas to store
        /// </summary>
        [Output("replicas")]
        public Output<int?> Replicas { get; private set; } = null!;

        /// <summary>
        /// How many seconds a value will be kept in the bucket
        /// </summary>
        [Output("ttl")]
        public Output<int?> Ttl { get; private set; } = null!;


        /// <summary>
        /// Create a KvBucket resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KvBucket(string name, KvBucketArgs? args = null, CustomResourceOptions? options = null)
            : base("jetstream:index/kvBucket:KvBucket", name, args ?? new KvBucketArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KvBucket(string name, Input<string> id, KvBucketState? state = null, CustomResourceOptions? options = null)
            : base("jetstream:index/kvBucket:KvBucket", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing KvBucket resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KvBucket Get(string name, Input<string> id, KvBucketState? state = null, CustomResourceOptions? options = null)
        {
            return new KvBucket(name, id, state, options);
        }
    }

    public sealed class KvBucketArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Contains additional information about this bucket
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// How many historical values to keep
        /// </summary>
        [Input("history")]
        public Input<int>? History { get; set; }

        /// <summary>
        /// Maximum size of the entire bucket
        /// </summary>
        [Input("maxBucketSize")]
        public Input<int>? MaxBucketSize { get; set; }

        /// <summary>
        /// Maximum size of any value
        /// </summary>
        [Input("maxValueSize")]
        public Input<int>? MaxValueSize { get; set; }

        /// <summary>
        /// The name of the Bucket
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Place the bucket in a specific cluster, influenced by placement_tags
        /// </summary>
        [Input("placementCluster")]
        public Input<string>? PlacementCluster { get; set; }

        [Input("placementTags")]
        private InputList<string>? _placementTags;

        /// <summary>
        /// Place the stream only on servers with these tags
        /// </summary>
        public InputList<string> PlacementTags
        {
            get => _placementTags ?? (_placementTags = new InputList<string>());
            set => _placementTags = value;
        }

        /// <summary>
        /// Number of cluster replicas to store
        /// </summary>
        [Input("replicas")]
        public Input<int>? Replicas { get; set; }

        /// <summary>
        /// How many seconds a value will be kept in the bucket
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public KvBucketArgs()
        {
        }
        public static new KvBucketArgs Empty => new KvBucketArgs();
    }

    public sealed class KvBucketState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Contains additional information about this bucket
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// How many historical values to keep
        /// </summary>
        [Input("history")]
        public Input<int>? History { get; set; }

        /// <summary>
        /// Maximum size of the entire bucket
        /// </summary>
        [Input("maxBucketSize")]
        public Input<int>? MaxBucketSize { get; set; }

        /// <summary>
        /// Maximum size of any value
        /// </summary>
        [Input("maxValueSize")]
        public Input<int>? MaxValueSize { get; set; }

        /// <summary>
        /// The name of the Bucket
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Place the bucket in a specific cluster, influenced by placement_tags
        /// </summary>
        [Input("placementCluster")]
        public Input<string>? PlacementCluster { get; set; }

        [Input("placementTags")]
        private InputList<string>? _placementTags;

        /// <summary>
        /// Place the stream only on servers with these tags
        /// </summary>
        public InputList<string> PlacementTags
        {
            get => _placementTags ?? (_placementTags = new InputList<string>());
            set => _placementTags = value;
        }

        /// <summary>
        /// Number of cluster replicas to store
        /// </summary>
        [Input("replicas")]
        public Input<int>? Replicas { get; set; }

        /// <summary>
        /// How many seconds a value will be kept in the bucket
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public KvBucketState()
        {
        }
        public static new KvBucketState Empty => new KvBucketState();
    }
}
