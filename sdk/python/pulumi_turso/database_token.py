# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DatabaseTokenArgs', 'DatabaseToken']

@pulumi.input_type
class DatabaseTokenArgs:
    def __init__(__self__, *,
                 database_name: pulumi.Input[str],
                 organization_name: pulumi.Input[str],
                 authorization: Optional[pulumi.Input[str]] = None,
                 expiration: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DatabaseToken resource.
        :param pulumi.Input[str] database_name: The name of the database.
        :param pulumi.Input[str] organization_name: The name of the organization or user.
        :param pulumi.Input[str] authorization: Authorization level for the token (full-access or read-only).
        :param pulumi.Input[str] expiration: Expiration time for the token (e.g., 2w1d30m).
        """
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "organization_name", organization_name)
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if expiration is not None:
            pulumi.set(__self__, "expiration", expiration)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[str]:
        """
        The name of the database.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="organizationName")
    def organization_name(self) -> pulumi.Input[str]:
        """
        The name of the organization or user.
        """
        return pulumi.get(self, "organization_name")

    @organization_name.setter
    def organization_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization_name", value)

    @property
    @pulumi.getter
    def authorization(self) -> Optional[pulumi.Input[str]]:
        """
        Authorization level for the token (full-access or read-only).
        """
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization", value)

    @property
    @pulumi.getter
    def expiration(self) -> Optional[pulumi.Input[str]]:
        """
        Expiration time for the token (e.g., 2w1d30m).
        """
        return pulumi.get(self, "expiration")

    @expiration.setter
    def expiration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration", value)


@pulumi.input_type
class _DatabaseTokenState:
    def __init__(__self__, *,
                 authorization: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 expiration: Optional[pulumi.Input[str]] = None,
                 jwt: Optional[pulumi.Input[str]] = None,
                 organization_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DatabaseToken resources.
        :param pulumi.Input[str] authorization: Authorization level for the token (full-access or read-only).
        :param pulumi.Input[str] database_name: The name of the database.
        :param pulumi.Input[str] expiration: Expiration time for the token (e.g., 2w1d30m).
        :param pulumi.Input[str] jwt: The generated authorization token (JWT).
        :param pulumi.Input[str] organization_name: The name of the organization or user.
        """
        if authorization is not None:
            pulumi.set(__self__, "authorization", authorization)
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if expiration is not None:
            pulumi.set(__self__, "expiration", expiration)
        if jwt is not None:
            pulumi.set(__self__, "jwt", jwt)
        if organization_name is not None:
            pulumi.set(__self__, "organization_name", organization_name)

    @property
    @pulumi.getter
    def authorization(self) -> Optional[pulumi.Input[str]]:
        """
        Authorization level for the token (full-access or read-only).
        """
        return pulumi.get(self, "authorization")

    @authorization.setter
    def authorization(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the database.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter
    def expiration(self) -> Optional[pulumi.Input[str]]:
        """
        Expiration time for the token (e.g., 2w1d30m).
        """
        return pulumi.get(self, "expiration")

    @expiration.setter
    def expiration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration", value)

    @property
    @pulumi.getter
    def jwt(self) -> Optional[pulumi.Input[str]]:
        """
        The generated authorization token (JWT).
        """
        return pulumi.get(self, "jwt")

    @jwt.setter
    def jwt(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "jwt", value)

    @property
    @pulumi.getter(name="organizationName")
    def organization_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the organization or user.
        """
        return pulumi.get(self, "organization_name")

    @organization_name.setter
    def organization_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_name", value)


class DatabaseToken(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 expiration: Optional[pulumi.Input[str]] = None,
                 organization_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Database Token resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorization: Authorization level for the token (full-access or read-only).
        :param pulumi.Input[str] database_name: The name of the database.
        :param pulumi.Input[str] expiration: Expiration time for the token (e.g., 2w1d30m).
        :param pulumi.Input[str] organization_name: The name of the organization or user.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabaseTokenArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Database Token resource

        :param str resource_name: The name of the resource.
        :param DatabaseTokenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabaseTokenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 expiration: Optional[pulumi.Input[str]] = None,
                 organization_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabaseTokenArgs.__new__(DatabaseTokenArgs)

            __props__.__dict__["authorization"] = authorization
            if database_name is None and not opts.urn:
                raise TypeError("Missing required property 'database_name'")
            __props__.__dict__["database_name"] = database_name
            __props__.__dict__["expiration"] = expiration
            if organization_name is None and not opts.urn:
                raise TypeError("Missing required property 'organization_name'")
            __props__.__dict__["organization_name"] = organization_name
            __props__.__dict__["jwt"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["jwt"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(DatabaseToken, __self__).__init__(
            'turso:index/databaseToken:DatabaseToken',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization: Optional[pulumi.Input[str]] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            expiration: Optional[pulumi.Input[str]] = None,
            jwt: Optional[pulumi.Input[str]] = None,
            organization_name: Optional[pulumi.Input[str]] = None) -> 'DatabaseToken':
        """
        Get an existing DatabaseToken resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorization: Authorization level for the token (full-access or read-only).
        :param pulumi.Input[str] database_name: The name of the database.
        :param pulumi.Input[str] expiration: Expiration time for the token (e.g., 2w1d30m).
        :param pulumi.Input[str] jwt: The generated authorization token (JWT).
        :param pulumi.Input[str] organization_name: The name of the organization or user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatabaseTokenState.__new__(_DatabaseTokenState)

        __props__.__dict__["authorization"] = authorization
        __props__.__dict__["database_name"] = database_name
        __props__.__dict__["expiration"] = expiration
        __props__.__dict__["jwt"] = jwt
        __props__.__dict__["organization_name"] = organization_name
        return DatabaseToken(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authorization(self) -> pulumi.Output[Optional[str]]:
        """
        Authorization level for the token (full-access or read-only).
        """
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[str]:
        """
        The name of the database.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter
    def expiration(self) -> pulumi.Output[Optional[str]]:
        """
        Expiration time for the token (e.g., 2w1d30m).
        """
        return pulumi.get(self, "expiration")

    @property
    @pulumi.getter
    def jwt(self) -> pulumi.Output[str]:
        """
        The generated authorization token (JWT).
        """
        return pulumi.get(self, "jwt")

    @property
    @pulumi.getter(name="organizationName")
    def organization_name(self) -> pulumi.Output[str]:
        """
        The name of the organization or user.
        """
        return pulumi.get(self, "organization_name")

