package utils

import (
    "errors"
    "log"
    "time"

    "github.com/golang-jwt/jwt/v4"
)

type JWTUtil struct {
    secretKey       string
    tokenDuration   time.Duration
    refreshDuration time.Duration
    signingMethod   jwt.SigningMethod
}

type CustomClaims struct {
    UserID string `json:"user_id"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}

// NewJWTUtil membuat instance baru dari JWTUtil
func NewJWTUtil(secret string, tokenDuration, refreshDuration time.Duration, signingMethod jwt.SigningMethod) (*JWTUtil, error) {
    if secret == "" {
        return nil, errors.New("secret key tidak boleh kosong")
    }

    if signingMethod == nil {
        signingMethod = jwt.SigningMethodHS256
    }

    return &JWTUtil{
        secretKey:       secret,
        tokenDuration:   tokenDuration,
        refreshDuration: refreshDuration,
        signingMethod:   signingMethod,
    }, nil
}

// GenerateToken membuat token JWT baru
func (j *JWTUtil) GenerateToken(userID, role string) (string, error) {
    claims := &CustomClaims{
        UserID: userID,
        Role:   role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.tokenDuration)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(j.signingMethod, claims)
    return token.SignedString([]byte(j.secretKey))
}

// ValidateToken memvalidasi token JWT dan mengembalikan klaimnya
func (j *JWTUtil) ValidateToken(tokenString string) (*CustomClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(j.secretKey), nil
    })

    if err != nil {
        log.Printf("Error saat memvalidasi token: %v", err)
        return nil, err
    }

    claims, ok := token.Claims.(*CustomClaims)
    if !ok || !token.Valid {
        return nil, errors.New("token tidak valid")
    }

    return claims, nil
}

// GenerateRefreshToken membuat refresh token JWT baru
func (j *JWTUtil) GenerateRefreshToken(userID, role string) (string, error) {
    claims := &CustomClaims{
        UserID: userID,
        Role:   role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.refreshDuration)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(j.signingMethod, claims)
    return token.SignedString([]byte(j.secretKey))
}

// ValidateRefreshToken memvalidasi refresh token JWT dan mengembalikan klaimnya
func (j *JWTUtil) ValidateRefreshToken(tokenString string) (*CustomClaims, error) {
    return j.ValidateToken(tokenString)
}

type IJWTUtil interface {
    GenerateToken(userID, role string) (string, error)
    ValidateToken(token string) (*CustomClaims, error)
    GenerateRefreshToken(userID, role string) (string, error)
    ValidateRefreshToken(token string) (*CustomClaims, error)
}