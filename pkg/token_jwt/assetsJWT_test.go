package tokenjwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInitJWTKey(t *testing.T) {
	// Проверяем инициализацию ключа
	InitJWTKey("mysecretkey")
	assert.Equal(t, "mysecretkey", SecretKey, "JWT ключ должен быть инициализирован")
}

func TestGenerateJWT_ValidToken(t *testing.T) {
	InitJWTKey("mysecretkey")

	userID := uuid.New()
	token, err := GenerateJWT(userID)
	assert.NoError(t, err, "Не должно быть ошибки при генерации токена")
	assert.NotEmpty(t, token, "Токен не должен быть пустым")

	// Парсим токен
	claims := &Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	assert.NoError(t, err, "Токен должен успешно парситься")
	assert.True(t, parsedToken.Valid, "Токен должен быть валидным")
	assert.Equal(t, userID, claims.UserID, "UserID должен совпадать")
	assert.True(t, time.Now().Before(claims.ExpiresAt.Time), "Токен не должен быть просрочен")
}

func TestGenerateJWT_InvalidSecretKey(t *testing.T) {
	InitJWTKey("mysecretkey")

	userID := uuid.New()
	token, err := GenerateJWT(userID)
	assert.NoError(t, err, "Не должно быть ошибки при генерации токена")

	// Меняем ключ, чтобы токен стал невалидным
	SecretKey = "wrongkey"

	claims := &Claims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	assert.Error(t, err, "Ожидаем ошибку при разборе токена с неверным ключом")
}

func TestDecodeJWT_InvalidToken(t *testing.T) {
	InitJWTKey("mysecretkey")

	invalidToken := "invalid.token.here"
	_, err := DecodeJWT(invalidToken)

	assert.Error(t, err, "Ожидаем ошибку при разборе невалидного токена")
}

func TestDecodeJWT_ExpiredToken(t *testing.T) {
	InitJWTKey("mysecretkey")

	expiredTime := time.Now().Add(-24 * time.Hour) // Токен уже истёк
	claims := &Claims{
		UserID: uuid.New(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(SecretKey))
	assert.NoError(t, err, "Токен должен быть успешно подписан")

	_, err = DecodeJWT(signedToken)
	assert.Error(t, err, "Декодирование истёкшего токена должно вызвать ошибку")
}
