package tools

import (
	"log"
	"time"

	loadpg "github.com/antalkon/zic_server/internal/db/db_pg/load_pg"
	"github.com/shirou/gopsutil/net"
)

func LoadNetwork() {
	var previousRecv uint64
	var previousSent uint64

	for {
		// Получаем статистику сетевого трафика
		ioStats, err := net.IOCounters(false)
		if err != nil {
			log.Printf("Error getting network load: %v", err)
			time.Sleep(10 * time.Minute)
			continue
		}

		// Вычисляем разницу с предыдущими значениями, чтобы получить новые мегабайты
		if len(ioStats) > 0 {
			mbRecv := (ioStats[0].BytesRecv - previousRecv) / (1024 * 1024) // переводим в мегабайты
			mbSent := (ioStats[0].BytesSent - previousSent) / (1024 * 1024) // переводим в мегабайты

			// Обновляем предыдущие значения
			previousRecv = ioStats[0].BytesRecv
			previousSent = ioStats[0].BytesSent

			// Сохраняем данные в таблицу (уже в мегабайтах)
			err = loadpg.NewNetworkRecord(mbRecv, mbSent)
			if err != nil {
				log.Printf("Error saving network load to database: %v", err)
			}
		}

		// Ждём 10 минут до следующего сбора данных
		time.Sleep(1 * time.Minute)
	}
}
