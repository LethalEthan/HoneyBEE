package server

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/panjf2000/gnet"
)

// CFB stream with 8 bit segment size - thank you to https://stackoverflow.com/users/3477832/kostya for the implementation used here
// See http://csrc.nist.gov/publications/nistpubs/800-38a/sp800-38a.pdf
type CFB8 struct {
	block     cipher.Block
	blockSize int
	iv, out   []byte
	decrypt   bool
}

func (CFB8 *CFB8) XORKeyStream(dst, src []byte) {
	for i := range src {
		CFB8.block.Encrypt(CFB8.out, CFB8.iv)
		copy(CFB8.iv[:CFB8.blockSize-1], CFB8.iv[1:])
		if CFB8.decrypt {
			CFB8.iv[CFB8.blockSize-1] = src[i]
		}
		dst[i] = src[i] ^ CFB8.out[0]
		if !CFB8.decrypt {
			CFB8.iv[CFB8.blockSize-1] = dst[i]
		}
	}
}

func NewCFB8Encrypter(block cipher.Block, iv []byte) *CFB8 {
	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		panic("cipher.newCFB: IV length must equal block size")
	}
	CFB8cipherstream := &CFB8{
		block:     block,
		blockSize: blockSize,
		out:       make([]byte, blockSize),
		iv:        make([]byte, blockSize),
		decrypt:   false,
	}
	copy(CFB8cipherstream.iv, iv)
	return CFB8cipherstream
}

// NewCFB8Decrypter returns a Stream which decrypts with cipher feedback mode
// (segment size = 8), using the given Block. The iv must be the same length as
// the Block's block size.
func NewCFB8Decrypter(block cipher.Block, iv []byte) *CFB8 {
	blockSize := block.BlockSize()
	if len(iv) != blockSize {
		panic("cipher.newCFB: IV length must equal block size")
	}
	CFB8cipherstream := &CFB8{
		block:     block,
		blockSize: blockSize,
		out:       make([]byte, blockSize),
		iv:        make([]byte, blockSize),
		decrypt:   true,
	}
	copy(CFB8cipherstream.iv, iv)
	return CFB8cipherstream
}

func CreateStreamCipher(sharedsecret []byte) (*CFB8, *CFB8, error) {
	block, err := aes.NewCipher(sharedsecret) // the shared secret is used both for the key and initial vector
	if err != nil {
		return nil, nil, err
	}
	Log.Debug("Created CFB8 E/D")
	return NewCFB8Encrypter(block, sharedsecret), NewCFB8Decrypter(block, sharedsecret), nil
}

func (Client *Client) EncryptPacket(data []byte) []byte {
	Client.encryptstream.XORKeyStream(data, data)
	return data
}

func (Client *Client) DecryptPacket(data []byte) []byte {
	Client.decryptstream.XORKeyStream(data, data)
	return data
}

func (Client *Client) SendData(c gnet.Conn, data []byte) (err error) {
	if Client.onlinemode {
		Log.Debug("Sending onli")
		data = Client.EncryptPacket(data)
		err = c.AsyncWrite(data)
	} else {
		Log.Debug("Sending ofln")
		err = c.AsyncWrite(data)
	}
	return
}
