# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: in_toto_attestation/predicates/vsa/v1/vsa.proto
"""Generated protocol buffer code."""

from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder

# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2

DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(
    b"\n/in_toto_attestation/predicates/vsa/v1/vsa.proto\x12%in_toto_attestation.predicates.vsa.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"\xcc\x08\n\x13VerificationSummary\x12U\n\x08verifier\x18\x01 \x01(\x0b\x32\x43.in_toto_attestation.predicates.vsa.v1.VerificationSummary.Verifier\x12?\n\rtime_verified\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x0ctimeVerified\x12!\n\x0cresource_uri\x18\x03 \x01(\tR\x0bresourceUri\x12Q\n\x06policy\x18\x04 \x01(\x0b\x32\x41.in_toto_attestation.predicates.vsa.v1.VerificationSummary.Policy\x12z\n\x12input_attestations\x18\x05 \x03(\x0b\x32K.in_toto_attestation.predicates.vsa.v1.VerificationSummary.InputAttestationR\x11inputAttestations\x12/\n\x13verification_result\x18\x06 \x01(\tR\x12verificationResult\x12'\n\x0fverified_levels\x18\x07 \x01(\tR\x0everifiedLevels\x12}\n\x11\x64\x65pendency_levels\x18\x08 \x03(\x0b\x32P.in_toto_attestation.predicates.vsa.v1.VerificationSummary.DependencyLevelsEntryR\x10\x64\x65pendencyLevels\x12!\n\x0cslsa_version\x18\t \x01(\tR\x0bslsaVersion\x1a\x16\n\x08Verifier\x12\n\n\x02id\x18\x01 \x01(\t\x1a\xa3\x01\n\x06Policy\x12\x0b\n\x03uri\x18\x01 \x01(\t\x12]\n\x06\x64igest\x18\x02 \x03(\x0b\x32M.in_toto_attestation.predicates.vsa.v1.VerificationSummary.Policy.DigestEntry\x1a-\n\x0b\x44igestEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a\xb7\x01\n\x10InputAttestation\x12\x0b\n\x03uri\x18\x01 \x01(\t\x12g\n\x06\x64igest\x18\x02 \x03(\x0b\x32W.in_toto_attestation.predicates.vsa.v1.VerificationSummary.InputAttestation.DigestEntry\x1a-\n\x0b\x44igestEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a\x37\n\x15\x44\x65pendencyLevelsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\x04:\x02\x38\x01\x42\x65\n.io.github.intoto.attestation.predicates.vsa.v1Z3github.com/in-toto/attestation/go/predicates/vsa/v1b\x06proto3"
)

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(
    DESCRIPTOR, "in_toto_attestation.predicates.vsa.v1.vsa_pb2", _globals
)
if _descriptor._USE_C_DESCRIPTORS == False:
    _globals["DESCRIPTOR"]._options = None
    _globals[
        "DESCRIPTOR"
    ]._serialized_options = b"\n.io.github.intoto.attestation.predicates.vsa.v1Z3github.com/in-toto/attestation/go/predicates/vsa/v1"
    _globals["_VERIFICATIONSUMMARY_POLICY_DIGESTENTRY"]._options = None
    _globals["_VERIFICATIONSUMMARY_POLICY_DIGESTENTRY"]._serialized_options = b"8\001"
    _globals["_VERIFICATIONSUMMARY_INPUTATTESTATION_DIGESTENTRY"]._options = None
    _globals[
        "_VERIFICATIONSUMMARY_INPUTATTESTATION_DIGESTENTRY"
    ]._serialized_options = b"8\001"
    _globals["_VERIFICATIONSUMMARY_DEPENDENCYLEVELSENTRY"]._options = None
    _globals[
        "_VERIFICATIONSUMMARY_DEPENDENCYLEVELSENTRY"
    ]._serialized_options = b"8\001"
    _globals["_VERIFICATIONSUMMARY"]._serialized_start = 124
    _globals["_VERIFICATIONSUMMARY"]._serialized_end = 1224
    _globals["_VERIFICATIONSUMMARY_VERIFIER"]._serialized_start = 793
    _globals["_VERIFICATIONSUMMARY_VERIFIER"]._serialized_end = 815
    _globals["_VERIFICATIONSUMMARY_POLICY"]._serialized_start = 818
    _globals["_VERIFICATIONSUMMARY_POLICY"]._serialized_end = 981
    _globals["_VERIFICATIONSUMMARY_POLICY_DIGESTENTRY"]._serialized_start = 936
    _globals["_VERIFICATIONSUMMARY_POLICY_DIGESTENTRY"]._serialized_end = 981
    _globals["_VERIFICATIONSUMMARY_INPUTATTESTATION"]._serialized_start = 984
    _globals["_VERIFICATIONSUMMARY_INPUTATTESTATION"]._serialized_end = 1167
    _globals[
        "_VERIFICATIONSUMMARY_INPUTATTESTATION_DIGESTENTRY"
    ]._serialized_start = 936
    _globals["_VERIFICATIONSUMMARY_INPUTATTESTATION_DIGESTENTRY"]._serialized_end = 981
    _globals["_VERIFICATIONSUMMARY_DEPENDENCYLEVELSENTRY"]._serialized_start = 1169
    _globals["_VERIFICATIONSUMMARY_DEPENDENCYLEVELSENTRY"]._serialized_end = 1224
# @@protoc_insertion_point(module_scope)
