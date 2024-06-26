// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class Template extends pulumi.CustomResource {
    /**
     * Get an existing Template resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Template {
        return new Template(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'runpod:index:Template';

    /**
     * Returns true if the given object is an instance of Template.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Template {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Template.__pulumiType;
    }

    public readonly containerDiskInGb!: pulumi.Output<number>;
    public readonly containerRegistryAuthId!: pulumi.Output<string | undefined>;
    public readonly dockerArgs!: pulumi.Output<string>;
    public readonly env!: pulumi.Output<outputs.PodEnv[]>;
    public readonly imageName!: pulumi.Output<string>;
    public readonly isPublic!: pulumi.Output<boolean | undefined>;
    public readonly isServerless!: pulumi.Output<boolean | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly ports!: pulumi.Output<string | undefined>;
    public readonly readme!: pulumi.Output<string | undefined>;
    public readonly startJupyter!: pulumi.Output<boolean | undefined>;
    public readonly startSsh!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly template!: pulumi.Output<outputs.Template>;
    public readonly volumeInGb!: pulumi.Output<number>;
    public readonly volumeMountPath!: pulumi.Output<string | undefined>;

    /**
     * Create a Template resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.containerDiskInGb === undefined) && !opts.urn) {
                throw new Error("Missing required property 'containerDiskInGb'");
            }
            if ((!args || args.dockerArgs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dockerArgs'");
            }
            if ((!args || args.env === undefined) && !opts.urn) {
                throw new Error("Missing required property 'env'");
            }
            if ((!args || args.imageName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageName'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.volumeInGb === undefined) && !opts.urn) {
                throw new Error("Missing required property 'volumeInGb'");
            }
            resourceInputs["containerDiskInGb"] = args ? args.containerDiskInGb : undefined;
            resourceInputs["containerRegistryAuthId"] = args ? args.containerRegistryAuthId : undefined;
            resourceInputs["dockerArgs"] = args ? args.dockerArgs : undefined;
            resourceInputs["env"] = args ? args.env : undefined;
            resourceInputs["imageName"] = args ? args.imageName : undefined;
            resourceInputs["isPublic"] = args ? args.isPublic : undefined;
            resourceInputs["isServerless"] = args ? args.isServerless : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ports"] = args ? args.ports : undefined;
            resourceInputs["readme"] = args ? args.readme : undefined;
            resourceInputs["startJupyter"] = args ? args.startJupyter : undefined;
            resourceInputs["startSsh"] = args ? args.startSsh : undefined;
            resourceInputs["volumeInGb"] = args ? args.volumeInGb : undefined;
            resourceInputs["volumeMountPath"] = args ? args.volumeMountPath : undefined;
            resourceInputs["template"] = undefined /*out*/;
        } else {
            resourceInputs["containerDiskInGb"] = undefined /*out*/;
            resourceInputs["containerRegistryAuthId"] = undefined /*out*/;
            resourceInputs["dockerArgs"] = undefined /*out*/;
            resourceInputs["env"] = undefined /*out*/;
            resourceInputs["imageName"] = undefined /*out*/;
            resourceInputs["isPublic"] = undefined /*out*/;
            resourceInputs["isServerless"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["ports"] = undefined /*out*/;
            resourceInputs["readme"] = undefined /*out*/;
            resourceInputs["startJupyter"] = undefined /*out*/;
            resourceInputs["startSsh"] = undefined /*out*/;
            resourceInputs["template"] = undefined /*out*/;
            resourceInputs["volumeInGb"] = undefined /*out*/;
            resourceInputs["volumeMountPath"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Template.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Template resource.
 */
export interface TemplateArgs {
    containerDiskInGb: pulumi.Input<number>;
    containerRegistryAuthId?: pulumi.Input<string>;
    dockerArgs: pulumi.Input<string>;
    env: pulumi.Input<pulumi.Input<inputs.PodEnvArgs>[]>;
    imageName: pulumi.Input<string>;
    isPublic?: pulumi.Input<boolean>;
    isServerless?: pulumi.Input<boolean>;
    name: pulumi.Input<string>;
    ports?: pulumi.Input<string>;
    readme?: pulumi.Input<string>;
    startJupyter?: pulumi.Input<boolean>;
    startSsh?: pulumi.Input<boolean>;
    volumeInGb: pulumi.Input<number>;
    volumeMountPath?: pulumi.Input<string>;
}
