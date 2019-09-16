package util

import (
	"crypto/aes"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

const (
	ecbKey  = "rFgB&h#%2?^eDg:Q"
	eapiKey = "e82ckenh8dichen8"
)

func EAPIAesEncryptECB(data, url string) (p string) {
	message := `nobody` + url + `use` + data + `md5forencrypt`
	hash := md5.New()
	_, err := hash.Write([]byte(message))
	if err != nil {
		panic(err)
	}
	digest := hex.EncodeToString(hash.Sum(nil))
	text := url + `-36cd479b6b5-` + data + `-36cd479b6b5-` + digest
	return AesEncryptECB(text, "eapi")

}

func AesEncryptECB(data, api string) (p string) {
	var key string
	if api == "lapi" {
		key = ecbKey
	} else if api == "eapi" {
		key = eapiKey
	}
	origData := []byte(data)
	ciphers, _ := aes.NewCipher(generateKey([]byte(key)))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, ciphers.BlockSize(); bs <= len(origData); bs, be = bs+ciphers.BlockSize(), be+ciphers.BlockSize() {
		ciphers.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	en := hex.EncodeToString(encrypted)
	return strings.ToUpper(en)
}

func AesDecryptECB(encrypted []byte) (decrypted []byte) {
	ciphers, _ := aes.NewCipher(generateKey([]byte(ecbKey)))
	decrypted = make([]byte, len(encrypted))
	for bs, be := 0, ciphers.BlockSize(); bs < len(encrypted); bs, be = bs+ciphers.BlockSize(), be+ciphers.BlockSize() {
		ciphers.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}
	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	return decrypted[:trim]
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
