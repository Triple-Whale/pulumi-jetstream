// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Jetstream.Outputs
{

    [OutputType]
    public sealed class StreamSourceSubjectTransform
    {
        /// <summary>
        /// The subject transform destination
        /// </summary>
        public readonly string Destination;
        /// <summary>
        /// The subject transform source
        /// </summary>
        public readonly string Source;

        [OutputConstructor]
        private StreamSourceSubjectTransform(
            string destination,

            string source)
        {
            Destination = destination;
            Source = source;
        }
    }
}
