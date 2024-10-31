package sessions

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Секретный ключ для шифрования куков
var secretKey = "secret.key"

// InitSessionStore инициализирует хранилище сессий с максимальным временем жизни 12 часов
func InitSessionStore() gin.HandlerFunc {
	store := cookie.NewStore([]byte(secretKey))

	return func(c *gin.Context) {
		// Получаем текущий хост (IP или домен) из запроса
		host := c.Request.Host

		// Настраиваем параметры куки, используя полученный хост как домен
		store.Options(sessions.Options{
			Path:     "/",
			Domain:   host,         // Динамически подставляем текущий домен или IP
			MaxAge:   12 * 60 * 60, // Время жизни сессии 12 часов
			HttpOnly: true,
			Secure:   false, // Установи true для продакшн-сервера на HTTPS

		})

		// Применяем middleware сессий с динамическим доменом
		sessions.Sessions("session-name", store)(c)
	}
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

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		// Если пользователя нет в сессии, возвращаем 403
		if user == nil {
			c.Redirect(http.StatusFound, "/403")
			c.Abort() // Прекращаем дальнейшую обработку
			return
		}

		// Продолжаем обработку, если пользователь найден
		c.Next()
	}
}
