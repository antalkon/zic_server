package sessions

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Секретный ключ для шифрования куков
var secretKey = "secret.key"

// InitSessionStore инициализирует хранилище сессий с максимальным временем жизни 12 часов
func InitSessionStore() gin.HandlerFunc {
	// Хранилище для сессий с использованием куков
	store := cookie.NewStore([]byte(secretKey))
	store.Options(sessions.Options{
		Path:     "/",          // Устанавливаем путь куки на корневой (вся область сайта)
		MaxAge:   12 * 60 * 60, // Время жизни сессии 12 часов
		HttpOnly: true,         // Куки недоступны через JS
		Secure:   false,        // Устанавливаем true для HTTPS
	})
	return sessions.Sessions("session-name", store)
}

// SetUserSession сохраняет сессию для пользователя
func SetUserSession(c *gin.Context, username string) error {
	session := sessions.Default(c)
	session.Set("user", username) // Сохраняем имя пользователя
	err := session.Save()         // Сохраняем сессию
	if err != nil {
		return errors.New("failed to save session")
	}
	return nil
}

// GetUserFromSession проверяет сессию и возвращает имя пользователя
func GetUserFromSession(c *gin.Context) (string, error) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		return "", errors.New("session not found or expired")
	}
	return user.(string), nil
}

// ClearSession очищает текущую сессию (например, для выхода)
func ClearSession(c *gin.Context) error {
	session := sessions.Default(c)
	session.Clear() // Очищаем сессию
	return session.Save()
}
