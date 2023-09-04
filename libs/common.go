package libs

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

func GetTimestamp(format string) (timestamp string) {
	dt := time.Now().UTC()
	timestamp = dt.Format(format)
	return
}

func GenerateSignature(path string, method string, token string, timestamp string, body string, secret string) (sig string) {
	payload := "path=" + path + "&verb=" + method + "&token=" + token + "&timestamp=" + timestamp + "&body=" + body
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(payload))

	sig = base64.StdEncoding.EncodeToString(h.Sum(nil))
	return
}

func GenerateSha1Timestamp(salt string) string {
	key := fmt.Sprintf("%s-%d", salt, time.Now().UnixNano())

	h := sha1.New()
	h.Write([]byte(key))
	return fmt.Sprintf("%x", h.Sum(nil))
}
