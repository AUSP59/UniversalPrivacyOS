
package encryption

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "crypto/x509"
    "encoding/pem"
    "errors"
    "fmt"
    "io"
)

func RunSelfTest() {
    fmt.Println("🔐 Running AES-256 + RSA encryption self-test...")

    key := make([]byte, 32)
    rand.Read(key)

    plaintext := []byte("UniversalPrivacyOS secure message")

    ciphertext, err := EncryptAES(plaintext, key)
    if err != nil {
        fmt.Println("AES encryption failed:", err)
        return
    }

    decrypted, err := DecryptAES(ciphertext, key)
    if err != nil || !bytes.Equal(plaintext, decrypted) {
        fmt.Println("AES decryption failed")
        return
    }

    priv, pub := GenerateRSAKeyPair(2048)
    sig := SignMessageRSA(priv, plaintext)
    if !VerifySignatureRSA(pub, plaintext, sig) {
        fmt.Println("RSA signature verification failed")
        return
    }

    fmt.Println("✅ AES-256 and RSA crypto test passed.")
}

func EncryptAES(plainText []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    cipherText := make([]byte, aes.BlockSize+len(plainText))
    iv := cipherText[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

    return cipherText, nil
}

func DecryptAES(cipherText []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    if len(cipherText) < aes.BlockSize {
        return nil, errors.New("cipherText too short")
    }

    iv := cipherText[:aes.BlockSize]
    cipherText = cipherText[aes.BlockSize:]

    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(cipherText, cipherText)

    return cipherText, nil
}

func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
    privKey, _ := rsa.GenerateKey(rand.Reader, bits)
    return privKey, &privKey.PublicKey
}

func SignMessageRSA(priv *rsa.PrivateKey, msg []byte) []byte {
    hash := sha256.Sum256(msg)
    sig, _ := rsa.SignPKCS1v15(rand.Reader, priv, cryptoHash(), hash[:])
    return sig
}

func VerifySignatureRSA(pub *rsa.PublicKey, msg, sig []byte) bool {
    hash := sha256.Sum256(msg)
    return rsa.VerifyPKCS1v15(pub, cryptoHash(), hash[:], sig) == nil
}

func cryptoHash() crypto.Hash {
    return crypto.SHA256
}
