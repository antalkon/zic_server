package tools

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	pcpg "github.com/antalkon/zic_server/internal/db/db_pg/pc_pg"
)

const (
	timeout  = 4 * time.Second
	interval = 30 * time.Second // Интервал между проверками
)

func CheckStatus() {
	for {
		ips, err := pcpg.GetAllIp()
		if err != nil {
			fmt.Println("Error getting IPs:", err)
			return
		}

		var wg sync.WaitGroup

		for _, ip := range ips {
			wg.Add(1)
			go func(ip string) {
				defer wg.Done()
				checkIPStatus(ip)
			}(ip)
		}

		wg.Wait() // Ждем завершения всех горутин

		// Ожидаем 5 минут перед следующим запуском
		time.Sleep(interval)
	}
}

func checkIPStatus(ip string) {
	url := fmt.Sprintf("http://%s:3000/server/ping", ip)

	// Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Создаем HTTP-запрос с таймаутом
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		pcpg.UpdateComputerStatus(ip, false)
		return
	}

	// Выполняем запрос
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request to", ip, ":", err)
		pcpg.UpdateComputerStatus(ip, false)
		return
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode == http.StatusOK {
		pcpg.UpdateComputerStatus(ip, true)
	} else {
		pcpg.UpdateComputerStatus(ip, false)
	}
}
