// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Jetstream.Inputs
{

    public sealed class StreamSourceExternalGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The subject prefix for the remote API
        /// </summary>
        [Input("api")]
        public Input<string>? Api { get; set; }

        /// <summary>
        /// The subject prefix where messages will be delivered to
        /// </summary>
        [Input("deliver")]
        public Input<string>? Deliver { get; set; }

        public StreamSourceExternalGetArgs()
        {
        }
        public static new StreamSourceExternalGetArgs Empty => new StreamSourceExternalGetArgs();
    }
}
