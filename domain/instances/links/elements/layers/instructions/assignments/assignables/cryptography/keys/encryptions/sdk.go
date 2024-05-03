package encryptions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
)

// Builder represents an encryption builder
type Builder interface {
	Create() Builder
	IsGeneratePrivateKey() Builder
	WithFetchPublicKey(fetchPublicKey string) Builder
	WithEncrypt(encrypt encrypts.Encrypt) Builder
	WithDecrypt(decrypt decrypts.Decrypt) Builder
	Now() (Encryption, error)
}

// Encryption represents encryption
type Encryption interface {
	Hash() hash.Hash
	IsGeneratePrivateKey() bool
	IsFetchPublicKey() bool
	FetchPublicKey() string
	IsEncrypt() bool
	Encrypt() encrypts.Encrypt
	IsDecrypt() bool
	Decrypt() decrypts.Decrypt
}