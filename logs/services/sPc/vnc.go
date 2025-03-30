package spc

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	pcpg "github.com/antalkon/zic_server/internal/db/db_pg/pc_pg"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Разрешить подключение от любого источника
	},
}

func Vnc(c *gin.Context) {
	pcID := c.Param("id")
	log.Printf("Запрос на подключение к ПК с ID: %s", pcID)

	// Получаем IP-адрес и порт VNC-сервера из базы данных
	vncAddress, err := pcpg.GetIpByID(pcID)
	if err != nil {
		log.Printf("Ошибка получения IP для ID %s: %v", pcID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve IP from database"})
		return
	}
	vncAddress = fmt.Sprintf("%s:5900", vncAddress)

	// Проверка на пустой IP-адрес
	if vncAddress == "" {
		log.Printf("Адрес VNC для ID %s не найден в базе данных", pcID)
		c.JSON(http.StatusNotFound, gin.H{"error": "VNC address not found for provided ID"})
		return
	}
	log.Printf("Подключаемся к VNC-серверу на %s", vncAddress)

	// Устанавливаем WebSocket-соединение с клиентом
	clientConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Ошибка при обновлении до WebSocket с клиентом:", err)
		return
	}
	defer clientConn.Close()
	log.Println("Успешное подключение WebSocket с клиентом")

	// Устанавливаем TCP-соединение с VNC-сервером
	vncConn, err := net.Dial("tcp", vncAddress)
	if err != nil {
		log.Printf("Ошибка подключения к VNC-серверу на %s: %v", vncAddress, err)
		clientConn.WriteMessage(websocket.TextMessage, []byte("Failed to connect to VNC server"))
		return
	}
	defer vncConn.Close()
	log.Println("Успешное TCP-соединение с VNC-сервером установлено")

	// Канал для управления завершением горутин
	done := make(chan struct{})

	// Пересылка данных от клиента на VNC-сервер
	go func() {
		defer func() {
			done <- struct{}{}
		}()
		for {
			_, message, err := clientConn.ReadMessage()
			if err != nil {
				log.Println("Ошибка чтения сообщения от клиента:", err)
				return
			}
			_, err = vncConn.Write(message)
			if err != nil {
				log.Println("Ошибка записи в VNC-сервер:", err)
				return
			}
		}
	}()

	// Пересылка данных от VNC-сервера клиенту
	go func() {
		defer func() {
			done <- struct{}{}
		}()
		buffer := make([]byte, 1024)
		for {
			n, err := vncConn.Read(buffer)
			if err != nil {
				if err != io.EOF {
					log.Println("Ошибка чтения от VNC-сервера:", err)
				}
				return
			}
			log.Printf("Сообщение от VNC-сервера: %d байт", n)
			err = clientConn.WriteMessage(websocket.BinaryMessage, buffer[:n])
			if err != nil {
				log.Println("Ошибка записи клиенту:", err)
				return
			}
		}
	}()

	// Ожидание завершения работы горутин
	<-done
	log.Println("Передача данных между клиентом и VNC-сервером завершена")
}
