// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("jetstream");

/**
 * The contents of the NATS 2.0 Credentials file to use
 */
export declare const credentialData: string | undefined;
Object.defineProperty(exports, "credentialData", {
    get() {
        return __config.get("credentialData");
    },
    enumerable: true,
});

/**
 * Path to the NATS 2.0 credentials file to use for authentication
 */
export declare const credentials: string | undefined;
Object.defineProperty(exports, "credentials", {
    get() {
        return __config.get("credentials");
    },
    enumerable: true,
});

/**
 * Connect using a NKEY seed stored in a file
 */
export declare const nkey: string | undefined;
Object.defineProperty(exports, "nkey", {
    get() {
        return __config.get("nkey");
    },
    enumerable: true,
});

/**
 * Connect using a password
 */
export declare const password: string | undefined;
Object.defineProperty(exports, "password", {
    get() {
        return __config.get("password");
    },
    enumerable: true,
});

/**
 * Comma separated list of NATS servers to connect to
 */
export declare const servers: string | undefined;
Object.defineProperty(exports, "servers", {
    get() {
        return __config.get("servers");
    },
    enumerable: true,
});

export declare const tls: outputs.config.Tls | undefined;
Object.defineProperty(exports, "tls", {
    get() {
        return __config.getObject<outputs.config.Tls>("tls");
    },
    enumerable: true,
});

/**
 * Connect using an username, used as token when no password is given
 */
export declare const user: string | undefined;
Object.defineProperty(exports, "user", {
    get() {
        return __config.get("user");
    },
    enumerable: true,
});

