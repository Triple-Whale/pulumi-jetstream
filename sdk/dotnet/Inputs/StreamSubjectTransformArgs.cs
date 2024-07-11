// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Jetstream.Inputs
{

    public sealed class StreamSubjectTransformArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The subject transform destination
        /// </summary>
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

        /// <summary>
        /// The subject transform source
        /// </summary>
        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        public StreamSubjectTransformArgs()
        {
        }
        public static new StreamSubjectTransformArgs Empty => new StreamSubjectTransformArgs();
    }
}
