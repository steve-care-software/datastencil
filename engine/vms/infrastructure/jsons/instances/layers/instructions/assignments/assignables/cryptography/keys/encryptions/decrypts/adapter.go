package decrypts

import (
	"bytes"
	"encoding/json"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
)

// Adapter represents the decrypt adapter
type Adapter struct {
	builder decrypts.Builder
}

func createAdapter(
	builder decrypts.Builder,
) decrypts.Adapter {
	out := Adapter{
		builder: builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins decrypts.Decrypt) ([]byte, error) {
	ptr, err := app.DecryptToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to instance
func (app *Adapter) ToInstance(data []byte) (decrypts.Decrypt, error) {
	ins := new(Decrypt)
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(ins)
	if err != nil {
		return nil, err
	}

	return app.StructToDecrypt(*ins)
}

// DecryptToStruct converts a decrypt to struct
func (app *Adapter) DecryptToStruct(ins decrypts.Decrypt) (*Decrypt, error) {
	return &Decrypt{
		Cipher:     ins.Cipher(),
		PrivateKey: ins.PrivateKey(),
	}, nil
}

// StructToDecrypt converts a struct to decrypt
func (app *Adapter) StructToDecrypt(str Decrypt) (decrypts.Decrypt, error) {
	return app.builder.Create().
		WithCipher(str.Cipher).
		WithPrivateKey(str.PrivateKey).
		Now()
}
