# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: in_toto_attestation/predicates/provenance/v1/provenance.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder

# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2

from in_toto_attestation.v1 import (
    resource_descriptor_pb2 as in__toto__attestation_dot_v1_dot_resource__descriptor__pb2,
)

DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(
    b'\n=in_toto_attestation/predicates/provenance/v1/provenance.proto\x12,in_toto_attestation.predicates.provenance.v1\x1a\x30in_toto_attestation/v1/resource_descriptor.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1fgoogle/protobuf/timestamp.proto"\xb4\x01\n\nProvenance\x12W\n\x10\x62uild_definition\x18\x01 \x01(\x0b\x32=.in_toto_attestation.predicates.provenance.v1.BuildDefinition\x12M\n\x0brun_details\x18\x02 \x01(\x0b\x32\x38.in_toto_attestation.predicates.provenance.v1.RunDetails"\xdc\x01\n\x0f\x42uildDefinition\x12\x12\n\nbuild_type\x18\x01 \x01(\t\x12\x34\n\x13\x65xternal_parameters\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x34\n\x13internal_parameters\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12I\n\x15resolved_dependencies\x18\x04 \x03(\x0b\x32*.in_toto_attestation.v1.ResourceDescriptor"\xe3\x01\n\nRunDetails\x12\x46\n\x07\x62uilder\x18\x01 \x01(\x0b\x32\x35.in_toto_attestation.predicates.provenance.v1.Builder\x12M\n\x08metadata\x18\x02 \x01(\x0b\x32;.in_toto_attestation.predicates.provenance.v1.BuildMetadata\x12>\n\nbyproducts\x18\x03 \x03(\x0b\x32*.in_toto_attestation.v1.ResourceDescriptor"\xe4\x01\n\x07\x42uilder\x12\n\n\x02id\x18\x01 \x01(\t\x12S\n\x07version\x18\x02 \x03(\x0b\x32\x42.in_toto_attestation.predicates.provenance.v1.Builder.VersionEntry\x12H\n\x14\x62uilder_dependencies\x18\x03 \x03(\x0b\x32*.in_toto_attestation.v1.ResourceDescriptor\x1a.\n\x0cVersionEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01"\x87\x01\n\rBuildMetadata\x12\x15\n\rinvocation_id\x18\x01 \x01(\t\x12.\n\nstarted_on\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12/\n\x0b\x66inished_on\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampBs\n5io.github.intoto.attestation.predicates.provenance.v1Z:github.com/in-toto/attestation/go/predicates/provenance/v1b\x06proto3'
)

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(
    DESCRIPTOR, "in_toto_attestation.predicates.provenance.v1.provenance_pb2", _globals
)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals["DESCRIPTOR"]._options = None
    _globals[
        "DESCRIPTOR"
    ]._serialized_options = b"\n5io.github.intoto.attestation.predicates.provenance.v1Z:github.com/in-toto/attestation/go/predicates/provenance/v1"
    _globals["_BUILDER_VERSIONENTRY"]._options = None
    _globals["_BUILDER_VERSIONENTRY"]._serialized_options = b"8\001"
    _globals["_PROVENANCE"]._serialized_start = 225
    _globals["_PROVENANCE"]._serialized_end = 405
    _globals["_BUILDDEFINITION"]._serialized_start = 408
    _globals["_BUILDDEFINITION"]._serialized_end = 628
    _globals["_RUNDETAILS"]._serialized_start = 631
    _globals["_RUNDETAILS"]._serialized_end = 858
    _globals["_BUILDER"]._serialized_start = 861
    _globals["_BUILDER"]._serialized_end = 1089
    _globals["_BUILDER_VERSIONENTRY"]._serialized_start = 1043
    _globals["_BUILDER_VERSIONENTRY"]._serialized_end = 1089
    _globals["_BUILDMETADATA"]._serialized_start = 1092
    _globals["_BUILDMETADATA"]._serialized_end = 1227
# @@protoc_insertion_point(module_scope)
