package utils

import (
	"airbusexpert/model"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"log"
	"time"
)

func JwtBuilder(user model.User) string {

	// For testing create the RSA key pair in the code
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("generating random key: %v", err)
	}

	// Create Square.jose signing key
	key := jose.SigningKey{Algorithm: jose.RS256, Key: privKey}

	// Create a Square.jose RSA signer, used to sign the JWT
	var signerOpts = jose.SignerOptions{}
	signerOpts.WithType("JWT")

	rsaSigner, err := jose.NewSigner(key, &signerOpts)
	if err != nil {
		panic(err)
	}

	// Create an instance of Builder that uses the rsa signer
	builder := jwt.Signed(rsaSigner)

	// Public Claims
	pubClaims := jwt.Claims{
		Issuer:   "SE",
		Subject:  "SE",
		ID:       fmt.Sprintf("%v", user.ID),
		Audience: jwt.Audience{},
		IssuedAt: jwt.NewNumericDate(time.Now().UTC()),
		Expiry:   jwt.NewNumericDate(time.Now().UTC().AddDate(0, 0, 7)),
	}

	// Private claims, as payload is JSON use the generic json patterns
	privClaims := model.Payload{
		Name:     user.Name,
		Surname:  user.Surname,
		Role:     user.Role,
		Username: user.ID,
	}

	// Add the claims. Note Claims returns a Builder so can chain
	builder = builder.Claims(pubClaims).Claims(privClaims)

	// Validate all ok, sign with the RSA key, and return a compact JWT
	rawJWT, err := builder.CompactSerialize()
	if err != nil {
		panic(err)
	}

	log.Println(rawJWT)
	return rawJWT
}
