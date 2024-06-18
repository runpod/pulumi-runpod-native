// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { NetworkStorageArgs } from "./networkStorage";
export type NetworkStorage = import("./networkStorage").NetworkStorage;
export const NetworkStorage: typeof import("./networkStorage").NetworkStorage = null as any;
utilities.lazyLoad(exports, ["NetworkStorage"], () => require("./networkStorage"));

export { PodArgs } from "./pod";
export type Pod = import("./pod").Pod;
export const Pod: typeof import("./pod").Pod = null as any;
utilities.lazyLoad(exports, ["Pod"], () => require("./pod"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { TemplateArgs } from "./template";
export type Template = import("./template").Template;
export const Template: typeof import("./template").Template = null as any;
utilities.lazyLoad(exports, ["Template"], () => require("./template"));


// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "runpod:index:NetworkStorage":
                return new NetworkStorage(name, <any>undefined, { urn })
            case "runpod:index:Pod":
                return new Pod(name, <any>undefined, { urn })
            case "runpod:index:Template":
                return new Template(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("runpod", "index", _module)
pulumi.runtime.registerResourcePackage("runpod", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:runpod") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
