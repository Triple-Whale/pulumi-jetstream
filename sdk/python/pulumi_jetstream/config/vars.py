# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

import types

__config__ = pulumi.Config('jetstream')


class _ExportableConfig(types.ModuleType):
    @property
    def credential_data(self) -> Optional[str]:
        """
        The contents of the NATS 2.0 Credentials file to use
        """
        return __config__.get('credentialData')

    @property
    def credentials(self) -> Optional[str]:
        """
        Path to the NATS 2.0 credentials file to use for authentication
        """
        return __config__.get('credentials')

    @property
    def nkey(self) -> Optional[str]:
        """
        Connect using a NKEY seed stored in a file
        """
        return __config__.get('nkey')

    @property
    def password(self) -> Optional[str]:
        """
        Connect using a password
        """
        return __config__.get('password')

    @property
    def servers(self) -> Optional[str]:
        """
        Comma separated list of NATS servers to connect to
        """
        return __config__.get('servers')

    @property
    def tls(self) -> Optional[str]:
        return __config__.get('tls')

    @property
    def user(self) -> Optional[str]:
        """
        Connect using an username, used as token when no password is given
        """
        return __config__.get('user')

