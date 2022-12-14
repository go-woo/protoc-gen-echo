// Auth use data type.
// versions:
// - protoc-gen-echo v0.1.1
// - protoc  v3.12.4
// source: example/v1/greeter.proto

package v1

import "github.com/golang-jwt/jwt"

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}
