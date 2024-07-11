// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Jetstream
{
    [JetstreamResourceType("jetstream:index/kvEntry:KvEntry")]
    public partial class KvEntry : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the bucket
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// The key of the entry
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The revision of the entry
        /// </summary>
        [Output("revision")]
        public Output<int> Revision { get; private set; } = null!;

        /// <summary>
        /// The value of the entry
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a KvEntry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KvEntry(string name, KvEntryArgs args, CustomResourceOptions? options = null)
            : base("jetstream:index/kvEntry:KvEntry", name, args ?? new KvEntryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KvEntry(string name, Input<string> id, KvEntryState? state = null, CustomResourceOptions? options = null)
            : base("jetstream:index/kvEntry:KvEntry", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KvEntry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KvEntry Get(string name, Input<string> id, KvEntryState? state = null, CustomResourceOptions? options = null)
        {
            return new KvEntry(name, id, state, options);
        }
    }

    public sealed class KvEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The key of the entry
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The value of the entry
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public KvEntryArgs()
        {
        }
        public static new KvEntryArgs Empty => new KvEntryArgs();
    }

    public sealed class KvEntryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// The key of the entry
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The revision of the entry
        /// </summary>
        [Input("revision")]
        public Input<int>? Revision { get; set; }

        /// <summary>
        /// The value of the entry
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public KvEntryState()
        {
        }
        public static new KvEntryState Empty => new KvEntryState();
    }
}
