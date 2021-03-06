// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msp/msp_config.proto

package msp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MSPConfig collects all the configuration information for
// an MSP. The Config field should be unmarshalled in a way
// that depends on the Type
type MSPConfig struct {
	// Type holds the type of the MSP; the default one would
	// be of type FABRIC implementing an X.509 based provider
	Type int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	// Config is MSP dependent configuration info
	Config []byte `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (m *MSPConfig) Reset()                    { *m = MSPConfig{} }
func (m *MSPConfig) String() string            { return proto.CompactTextString(m) }
func (*MSPConfig) ProtoMessage()               {}
func (*MSPConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *MSPConfig) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MSPConfig) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

// FabricMSPConfig collects all the configuration information for
// a Fabric MSP.
// Here we assume a default certificate validation policy, where
// any certificate signed by any of the listed rootCA certs would
// be considered as valid under this MSP.
// This MSP may or may not come with a signing identity. If it does,
// it can also issue signing identities. If it does not, it can only
// be used to validate and verify certificates.
type FabricMSPConfig struct {
	// Name holds the identifier of the MSP; MSP identifier
	// is chosen by the application that governs this MSP.
	// For example, and assuming the default implementation of MSP,
	// that is X.509-based and considers a single Issuer,
	// this can refer to the Subject OU field or the Issuer OU field.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// List of root certificates trusted by this MSP
	// they are used upon certificate validation (see
	// comment for IntermediateCerts below)
	RootCerts [][]byte `protobuf:"bytes,2,rep,name=root_certs,json=rootCerts,proto3" json:"root_certs,omitempty"`
	// List of intermediate certificates trusted by this MSP;
	// they are used upon certificate validation as follows:
	// validation attempts to build a path from the certificate
	// to be validated (which is at one end of the path) and
	// one of the certs in the RootCerts field (which is at
	// the other end of the path). If the path is longer than
	// 2, certificates in the middle are searched within the
	// IntermediateCerts pool
	IntermediateCerts [][]byte `protobuf:"bytes,3,rep,name=intermediate_certs,json=intermediateCerts,proto3" json:"intermediate_certs,omitempty"`
	// Identity denoting the administrator of this MSP
	Admins [][]byte `protobuf:"bytes,4,rep,name=admins,proto3" json:"admins,omitempty"`
	// Identity revocation list
	RevocationList [][]byte `protobuf:"bytes,5,rep,name=revocation_list,json=revocationList,proto3" json:"revocation_list,omitempty"`
	// SigningIdentity holds information on the signing identity
	// this peer is to use, and which is to be imported by the
	// MSP defined before
	SigningIdentity *SigningIdentityInfo `protobuf:"bytes,6,opt,name=signing_identity,json=signingIdentity" json:"signing_identity,omitempty"`
	// OrganizationalUnitIdentifiers holds one or more
	// fabric organizational unit identifiers that belong to
	// this MSP configuration
	OrganizationalUnitIdentifiers []*FabricOUIdentifier `protobuf:"bytes,7,rep,name=organizational_unit_identifiers,json=organizationalUnitIdentifiers" json:"organizational_unit_identifiers,omitempty"`
	// FabricCryptoConfig contains the configuration parameters
	// for the cryptographic algorithms used by this MSP
	CryptoConfig *FabricCryptoConfig `protobuf:"bytes,8,opt,name=crypto_config,json=cryptoConfig" json:"crypto_config,omitempty"`
	// List of TLS root certificates trusted by this MSP.
	// They are returned by GetTLSRootCerts.
	TlsRootCerts [][]byte `protobuf:"bytes,9,rep,name=tls_root_certs,json=tlsRootCerts,proto3" json:"tls_root_certs,omitempty"`
	// List of TLS intermediate certificates trusted by this MSP;
	// They are returned by GetTLSIntermediateCerts.
	TlsIntermediateCerts [][]byte `protobuf:"bytes,10,rep,name=tls_intermediate_certs,json=tlsIntermediateCerts,proto3" json:"tls_intermediate_certs,omitempty"`
}

func (m *FabricMSPConfig) Reset()                    { *m = FabricMSPConfig{} }
func (m *FabricMSPConfig) String() string            { return proto.CompactTextString(m) }
func (*FabricMSPConfig) ProtoMessage()               {}
func (*FabricMSPConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *FabricMSPConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FabricMSPConfig) GetRootCerts() [][]byte {
	if m != nil {
		return m.RootCerts
	}
	return nil
}

func (m *FabricMSPConfig) GetIntermediateCerts() [][]byte {
	if m != nil {
		return m.IntermediateCerts
	}
	return nil
}

func (m *FabricMSPConfig) GetAdmins() [][]byte {
	if m != nil {
		return m.Admins
	}
	return nil
}

func (m *FabricMSPConfig) GetRevocationList() [][]byte {
	if m != nil {
		return m.RevocationList
	}
	return nil
}

func (m *FabricMSPConfig) GetSigningIdentity() *SigningIdentityInfo {
	if m != nil {
		return m.SigningIdentity
	}
	return nil
}

func (m *FabricMSPConfig) GetOrganizationalUnitIdentifiers() []*FabricOUIdentifier {
	if m != nil {
		return m.OrganizationalUnitIdentifiers
	}
	return nil
}

func (m *FabricMSPConfig) GetCryptoConfig() *FabricCryptoConfig {
	if m != nil {
		return m.CryptoConfig
	}
	return nil
}

func (m *FabricMSPConfig) GetTlsRootCerts() [][]byte {
	if m != nil {
		return m.TlsRootCerts
	}
	return nil
}

func (m *FabricMSPConfig) GetTlsIntermediateCerts() [][]byte {
	if m != nil {
		return m.TlsIntermediateCerts
	}
	return nil
}

// FabricCryptoConfig contains configuration parameters
// for the cryptographic algorithms used by the MSP
// this configuration refers to
type FabricCryptoConfig struct {
	// SignatureHashFamily is a string representing the hash family to be used
	// during sign and verify operations.
	// Allowed values are "SHA2" and "SHA3".
	SignatureHashFamily string `protobuf:"bytes,1,opt,name=signature_hash_family,json=signatureHashFamily" json:"signature_hash_family,omitempty"`
	// IdentityIdentifierHashFunction is a string representing the hash function
	// to be used during the computation of the identity identifier of an MSP identity.
	// Allowed values are "SHA256", "SHA384" and "SHA3_256", "SHA3_384".
	IdentityIdentifierHashFunction string `protobuf:"bytes,2,opt,name=identity_identifier_hash_function,json=identityIdentifierHashFunction" json:"identity_identifier_hash_function,omitempty"`
}

func (m *FabricCryptoConfig) Reset()                    { *m = FabricCryptoConfig{} }
func (m *FabricCryptoConfig) String() string            { return proto.CompactTextString(m) }
func (*FabricCryptoConfig) ProtoMessage()               {}
func (*FabricCryptoConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *FabricCryptoConfig) GetSignatureHashFamily() string {
	if m != nil {
		return m.SignatureHashFamily
	}
	return ""
}

func (m *FabricCryptoConfig) GetIdentityIdentifierHashFunction() string {
	if m != nil {
		return m.IdentityIdentifierHashFunction
	}
	return ""
}

// SigningIdentityInfo represents the configuration information
// related to the signing identity the peer is to use for generating
// endorsements
type SigningIdentityInfo struct {
	// PublicSigner carries the public information of the signing
	// identity. For an X.509 provider this would be represented by
	// an X.509 certificate
	PublicSigner []byte `protobuf:"bytes,1,opt,name=public_signer,json=publicSigner,proto3" json:"public_signer,omitempty"`
	// PrivateSigner denotes a reference to the private key of the
	// peer's signing identity
	PrivateSigner *KeyInfo `protobuf:"bytes,2,opt,name=private_signer,json=privateSigner" json:"private_signer,omitempty"`
}

func (m *SigningIdentityInfo) Reset()                    { *m = SigningIdentityInfo{} }
func (m *SigningIdentityInfo) String() string            { return proto.CompactTextString(m) }
func (*SigningIdentityInfo) ProtoMessage()               {}
func (*SigningIdentityInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *SigningIdentityInfo) GetPublicSigner() []byte {
	if m != nil {
		return m.PublicSigner
	}
	return nil
}

func (m *SigningIdentityInfo) GetPrivateSigner() *KeyInfo {
	if m != nil {
		return m.PrivateSigner
	}
	return nil
}

// KeyInfo represents a (secret) key that is either already stored
// in the bccsp/keystore or key material to be imported to the
// bccsp key-store. In later versions it may contain also a
// keystore identifier
type KeyInfo struct {
	// Identifier of the key inside the default keystore; this for
	// the case of Software BCCSP as well as the HSM BCCSP would be
	// the SKI of the key
	KeyIdentifier string `protobuf:"bytes,1,opt,name=key_identifier,json=keyIdentifier" json:"key_identifier,omitempty"`
	// KeyMaterial (optional) for the key to be imported; this is
	// properly encoded key bytes, prefixed by the type of the key
	KeyMaterial []byte `protobuf:"bytes,2,opt,name=key_material,json=keyMaterial,proto3" json:"key_material,omitempty"`
}

func (m *KeyInfo) Reset()                    { *m = KeyInfo{} }
func (m *KeyInfo) String() string            { return proto.CompactTextString(m) }
func (*KeyInfo) ProtoMessage()               {}
func (*KeyInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *KeyInfo) GetKeyIdentifier() string {
	if m != nil {
		return m.KeyIdentifier
	}
	return ""
}

func (m *KeyInfo) GetKeyMaterial() []byte {
	if m != nil {
		return m.KeyMaterial
	}
	return nil
}

// FabricOUIdentifier represents an organizational unit and
// its related chain of trust identifier.
type FabricOUIdentifier struct {
	// Certificate represents the second certificate in a certification chain.
	// (Notice that the first certificate in a certification chain is supposed
	// to be the certificate of an identity).
	// It must correspond to the certificate of root or intermediate CA
	// recognized by the MSP this message belongs to.
	// Starting from this certificate, a certification chain is computed
	// and boud to the OrganizationUnitIdentifier specified
	Certificate []byte `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	// OrganizationUnitIdentifier defines the organizational unit under the
	// MSP identified with MSPIdentifier
	OrganizationalUnitIdentifier string `protobuf:"bytes,2,opt,name=organizational_unit_identifier,json=organizationalUnitIdentifier" json:"organizational_unit_identifier,omitempty"`
}

func (m *FabricOUIdentifier) Reset()                    { *m = FabricOUIdentifier{} }
func (m *FabricOUIdentifier) String() string            { return proto.CompactTextString(m) }
func (*FabricOUIdentifier) ProtoMessage()               {}
func (*FabricOUIdentifier) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *FabricOUIdentifier) GetCertificate() []byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *FabricOUIdentifier) GetOrganizationalUnitIdentifier() string {
	if m != nil {
		return m.OrganizationalUnitIdentifier
	}
	return ""
}

func init() {
	proto.RegisterType((*MSPConfig)(nil), "msp.MSPConfig")
	proto.RegisterType((*FabricMSPConfig)(nil), "msp.FabricMSPConfig")
	proto.RegisterType((*FabricCryptoConfig)(nil), "msp.FabricCryptoConfig")
	proto.RegisterType((*SigningIdentityInfo)(nil), "msp.SigningIdentityInfo")
	proto.RegisterType((*KeyInfo)(nil), "msp.KeyInfo")
	proto.RegisterType((*FabricOUIdentifier)(nil), "msp.FabricOUIdentifier")
}

func init() { proto.RegisterFile("msp/msp_config.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0xd7, 0xad, 0xa3, 0xa7, 0x69, 0x37, 0xbc, 0x31, 0x72, 0xc1, 0x46, 0x57, 0xfe, 0x7a,
	0x43, 0x2b, 0x6d, 0x48, 0xdc, 0x70, 0xb5, 0xa2, 0x89, 0x0a, 0x26, 0xa6, 0x54, 0xbb, 0x41, 0x42,
	0x91, 0x9b, 0xba, 0xe9, 0x51, 0x13, 0xdb, 0xb2, 0xdd, 0xa1, 0x20, 0xde, 0x82, 0x57, 0xe5, 0x01,
	0x50, 0x6c, 0x6f, 0x4b, 0x61, 0xda, 0x9d, 0xfd, 0xfd, 0x9c, 0x1e, 0x7f, 0xe7, 0x34, 0xb0, 0x9f,
	0x6b, 0x39, 0xcc, 0xb5, 0x8c, 0x13, 0xc1, 0xe7, 0x98, 0x0e, 0xa4, 0x12, 0x46, 0x90, 0x7a, 0xae,
	0x65, 0xef, 0x3d, 0x34, 0x2f, 0x26, 0x97, 0x23, 0x8b, 0x13, 0x02, 0x9b, 0xa6, 0x90, 0x2c, 0xac,
	0x75, 0x6b, 0xfd, 0xad, 0xc8, 0x9e, 0xc9, 0x01, 0x34, 0x9c, 0x2b, 0xdc, 0xe8, 0xd6, 0xfa, 0x41,
	0xe4, 0x6f, 0xbd, 0x3f, 0x75, 0xd8, 0x39, 0xa7, 0x53, 0x85, 0xc9, 0x9a, 0x9f, 0xd3, 0xdc, 0xf9,
	0x9b, 0x91, 0x3d, 0x93, 0x43, 0x00, 0x25, 0x84, 0x89, 0x13, 0xa6, 0x8c, 0x0e, 0x37, 0xba, 0xf5,
	0x7e, 0x10, 0x35, 0x4b, 0x64, 0x54, 0x02, 0xe4, 0x2d, 0x10, 0xe4, 0x86, 0xa9, 0x9c, 0xcd, 0x90,
	0x1a, 0xe6, 0x65, 0x75, 0x2b, 0x7b, 0x5c, 0x65, 0x9c, 0xfc, 0x00, 0x1a, 0x74, 0x96, 0x23, 0xd7,
	0xe1, 0xa6, 0x95, 0xf8, 0x1b, 0x79, 0x03, 0x3b, 0x8a, 0x5d, 0x8b, 0x84, 0x1a, 0x14, 0x3c, 0xce,
	0x50, 0x9b, 0x70, 0xcb, 0x0a, 0x3a, 0x77, 0xf0, 0x17, 0xd4, 0x86, 0x8c, 0x60, 0x57, 0x63, 0xca,
	0x91, 0xa7, 0x31, 0xce, 0x18, 0x37, 0x68, 0x8a, 0xb0, 0xd1, 0xad, 0xf5, 0x5b, 0x27, 0xe1, 0x20,
	0xd7, 0x72, 0x30, 0x71, 0xe4, 0xd8, 0x73, 0x63, 0x3e, 0x17, 0xd1, 0x8e, 0x5e, 0x07, 0x49, 0x0c,
	0xcf, 0x85, 0x4a, 0x29, 0xc7, 0x9f, 0xb6, 0x30, 0xcd, 0xe2, 0x15, 0x47, 0xe3, 0x0b, 0xce, 0x91,
	0x29, 0x1d, 0x6e, 0x77, 0xeb, 0xfd, 0xd6, 0xc9, 0x53, 0x5b, 0xd3, 0xc5, 0xf4, 0xf5, 0x6a, 0x7c,
	0xcb, 0x47, 0x87, 0xeb, 0xfe, 0x2b, 0x8e, 0xe6, 0x8e, 0xd5, 0xe4, 0x03, 0xb4, 0x13, 0x55, 0x48,
	0x23, 0xfc, 0xc4, 0xc2, 0x47, 0xb6, 0xc5, 0x6a, 0xb9, 0x91, 0xe5, 0x5d, 0xf0, 0x51, 0x90, 0x54,
	0x6e, 0xe4, 0x25, 0x74, 0x4c, 0xa6, 0xe3, 0x4a, 0xec, 0x4d, 0x9b, 0x45, 0x60, 0x32, 0x1d, 0xdd,
	0x26, 0xff, 0x0e, 0x0e, 0x4a, 0xd5, 0x3d, 0xe9, 0x83, 0x55, 0xef, 0x9b, 0x4c, 0x8f, 0xff, 0x1d,
	0x40, 0xef, 0x77, 0x0d, 0xc8, 0xff, 0x0d, 0x90, 0x13, 0x78, 0x52, 0x86, 0x44, 0xcd, 0x4a, 0xb1,
	0x78, 0x41, 0xf5, 0x22, 0x9e, 0xd3, 0x1c, 0xb3, 0xc2, 0xaf, 0xc2, 0xde, 0x2d, 0xf9, 0x89, 0xea,
	0xc5, 0xb9, 0xa5, 0xc8, 0x18, 0x8e, 0x6f, 0x46, 0x50, 0x89, 0xce, 0xbb, 0x57, 0x3c, 0x29, 0xa3,
	0xb1, 0x4b, 0xd7, 0x8c, 0x8e, 0x6e, 0x84, 0x77, 0x21, 0xd9, 0x42, 0x5e, 0xd5, 0x13, 0xb0, 0x77,
	0xcf, 0xe0, 0xc8, 0x0b, 0x68, 0xcb, 0xd5, 0x34, 0xc3, 0x24, 0x2e, 0x7f, 0x9f, 0x29, 0xdb, 0x4d,
	0x10, 0x05, 0x0e, 0x9c, 0x58, 0x8c, 0x9c, 0x42, 0x47, 0x2a, 0xbc, 0x2e, 0x9f, 0xef, 0x55, 0x1b,
	0x36, 0xec, 0xc0, 0x86, 0xfd, 0x99, 0xb9, 0x1d, 0x68, 0x7b, 0x8d, 0x33, 0xf5, 0x26, 0xb0, 0xed,
	0x19, 0xf2, 0x0a, 0x3a, 0x4b, 0x56, 0x7d, 0x81, 0x7f, 0x73, 0x7b, 0xc9, 0x2a, 0xed, 0x92, 0x63,
	0x08, 0x4a, 0x59, 0x4e, 0x0d, 0x53, 0x48, 0x33, 0xff, 0x6f, 0x6a, 0x2d, 0x59, 0x71, 0xe1, 0xa1,
	0xde, 0xaf, 0x9b, 0x68, 0xab, 0xab, 0x42, 0xba, 0xd0, 0x2a, 0xc7, 0x82, 0x73, 0x4c, 0xa8, 0x61,
	0xfe, 0x09, 0x55, 0x88, 0x7c, 0x84, 0xa3, 0x87, 0xd7, 0xd1, 0xa7, 0xf8, 0xec, 0xa1, 0xa5, 0x3b,
	0xfb, 0x0e, 0xc7, 0x42, 0xa5, 0x83, 0x45, 0x21, 0x99, 0xca, 0xd8, 0x2c, 0x65, 0x6a, 0x30, 0xb7,
	0xdd, 0xb8, 0xcf, 0x85, 0x2e, 0xe3, 0x38, 0xdb, 0xbd, 0xd0, 0xd2, 0x8d, 0xfc, 0x92, 0x26, 0x4b,
	0x9a, 0xb2, 0x6f, 0xaf, 0x53, 0x34, 0x8b, 0xd5, 0x74, 0x90, 0x88, 0x7c, 0xf8, 0x63, 0x41, 0x0d,
	0x6a, 0x21, 0xe4, 0xd0, 0x39, 0x87, 0xce, 0x59, 0x7e, 0x7a, 0xa6, 0x0d, 0x7b, 0x3e, 0xfd, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x93, 0x19, 0x6d, 0x4c, 0x8c, 0x04, 0x00, 0x00,
}
