// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Database data source
 */
export function getDatabase(args: GetDatabaseArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("turso:index/getDatabase:getDatabase", {
        "group": args.group,
        "isSchema": args.isSchema,
        "name": args.name,
        "organizationName": args.organizationName,
        "schema": args.schema,
        "sizeLimit": args.sizeLimit,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabase.
 */
export interface GetDatabaseArgs {
    /**
     * The name of the group where the database should be created. The group must already exist.
     */
    group?: string;
    /**
     * Mark this database as the parent schema database that updates child databases with any schema changes.
     */
    isSchema?: boolean;
    /**
     * The name of the new database. Must contain only lowercase letters, numbers, dashes. No longer than 32 characters.
     */
    name: string;
    /**
     * Name of organization to create the database for
     */
    organizationName: string;
    /**
     * The name of the parent database to use as the schema.
     */
    schema?: string;
    /**
     * The maximum size of the database in bytes. Values with units are also accepted, e.g. 1mb, 256mb, 1gb.
     */
    sizeLimit?: string;
}

/**
 * A collection of values returned by getDatabase.
 */
export interface GetDatabaseResult {
    /**
     * The database universal unique identifier (UUID).
     */
    readonly dbId: string;
    /**
     * The name of the group where the database should be created. The group must already exist.
     */
    readonly group?: string;
    /**
     * The DNS hostname used for client libSQL and HTTP connections.
     */
    readonly hostname: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Mark this database as the parent schema database that updates child databases with any schema changes.
     */
    readonly isSchema?: boolean;
    /**
     * The name of the new database. Must contain only lowercase letters, numbers, dashes. No longer than 32 characters.
     */
    readonly name: string;
    /**
     * Name of organization to create the database for
     */
    readonly organizationName: string;
    /**
     * The name of the parent database to use as the schema.
     */
    readonly schema?: string;
    /**
     * The maximum size of the database in bytes. Values with units are also accepted, e.g. 1mb, 256mb, 1gb.
     */
    readonly sizeLimit?: string;
}
/**
 * Database data source
 */
export function getDatabaseOutput(args: GetDatabaseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabaseResult> {
    return pulumi.output(args).apply((a: any) => getDatabase(a, opts))
}

/**
 * A collection of arguments for invoking getDatabase.
 */
export interface GetDatabaseOutputArgs {
    /**
     * The name of the group where the database should be created. The group must already exist.
     */
    group?: pulumi.Input<string>;
    /**
     * Mark this database as the parent schema database that updates child databases with any schema changes.
     */
    isSchema?: pulumi.Input<boolean>;
    /**
     * The name of the new database. Must contain only lowercase letters, numbers, dashes. No longer than 32 characters.
     */
    name: pulumi.Input<string>;
    /**
     * Name of organization to create the database for
     */
    organizationName: pulumi.Input<string>;
    /**
     * The name of the parent database to use as the schema.
     */
    schema?: pulumi.Input<string>;
    /**
     * The maximum size of the database in bytes. Values with units are also accepted, e.g. 1mb, 256mb, 1gb.
     */
    sizeLimit?: pulumi.Input<string>;
}
