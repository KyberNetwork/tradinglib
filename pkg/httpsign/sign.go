package httpsign

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/KyberNetwork/tradinglib/pkg/sb"
)

const (
	requestTarget = "(request-target)"
	nonce         = "nonce"
	digest        = "digest"
)

var headers = []string{requestTarget, nonce, digest} // nolint: gochecknoglobals

// Sign a http request with selected header and return a signed request
// Use with http client only.
func Sign(r *http.Request, keyID string, secret []byte) (*http.Request, error) {
	// Set digest to body
	digestBody, err := calculateDigest(r)
	if err != nil {
		return nil, err
	}
	r.Header.Set(digest, digestBody)
	// Set nonce
	currentNonce := time.Now().UnixNano() / int64(time.Millisecond)
	r.Header.Set(nonce, strconv.FormatInt(currentNonce, 10)) // nolint: gomnd
	// Create sign string
	var signBuffer bytes.Buffer
	for i, h := range headers {
		var value string
		switch h {
		case requestTarget:
			value = sb.Concat(strings.ToLower(r.Method), " ", r.URL.RequestURI())
		default:
			value = r.Header.Get(h)
		}
		signBuffer.WriteString(sb.Concat(h, ": ", value))
		if i < len(headers)-1 {
			signBuffer.WriteString("\n")
		}
	}
	// Create signature header
	signature, err := sign(signBuffer.Bytes(), secret)
	if err != nil {
		return nil, err
	}
	signatureHeader := constructHeader(headers, keyID, signature)
	r.Header.Set("Signature", signatureHeader)
	return r, nil
}

func calculateDigest(r *http.Request) (string, error) {
	if r.Body == nil || r.ContentLength == 0 {
		return "", nil
	}
	data, err := io.ReadAll(r.Body)
	_ = r.Body.Close()
	if err != nil {
		return "", err
	}
	r.Body = io.NopCloser(bytes.NewBuffer(data))
	h := sha256.Sum256(data)
	sig := sb.Concat("SHA-256=", base64.StdEncoding.EncodeToString(h[:]))
	return sig, nil
}

func sign(msg, secret []byte) (string, error) {
	mac := hmac.New(sha512.New, secret)
	if _, err := mac.Write(msg); err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return signature, nil
}

func constructHeader(headers []string, keyID, signature string) string {
	var signBuffer bytes.Buffer
	signBuffer.WriteString(fmt.Sprintf(`keyId="%s",`, keyID))
	signBuffer.WriteString(`algorithm="hmac-sha512",`)
	signBuffer.WriteString(fmt.Sprintf(`headers="%s",`, strings.Join(headers, " ")))
	signBuffer.WriteString(fmt.Sprintf(`signature="%s"`, signature))
	return signBuffer.String()
}
