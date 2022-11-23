package zalo

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

type ZaloHelper interface {
	GetVerifierAndChallengeCode(codeVerifier string) map[string]interface{}
	Base64urlEncode(text string) string
}

func GetVerifierAndChallengeCode(codeVerifier string) map[string]interface{} {
	base64CodeVerifier := Base64urlEncode(codeVerifier)
	h := sha256.New()
	h.Write([]byte(base64CodeVerifier))
	codeVerifierEncodeSHA256 := h.Sum(nil)
	codeChallenge := Base64urlEncode(string(codeVerifierEncodeSHA256))
	result := make(map[string]interface{})
	result["codeChallenge"] = codeChallenge
	result["codeVerifier"] = base64CodeVerifier
	return result
}

//challenge

func Base64urlEncode(text string) string {
	base64Text := base64.StdEncoding.EncodeToString([]byte(text))
	base64Text = strings.Trim(base64Text, "=")
	base64url := strings.Replace(base64Text, "+/", "-_", -1)
	return base64url
}
