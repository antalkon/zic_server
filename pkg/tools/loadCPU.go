package tools

import (
	"log"
	"time"

	loadpg "github.com/antalkon/zic_server/internal/db/db_pg/load_pg"
	"github.com/shirou/gopsutil/cpu"
)

func LoadCPU() {
	for {
		// Получаем процент загрузки CPU
		usage, err := cpu.Percent(0, false)
		if err != nil {
			log.Printf("Error getting CPU usage: %v", err)
			time.Sleep(10 * time.Minute) // Задержка перед следующим опросом при ошибке
			continue
		}

		// Сохраняем данные в таблицу
		if len(usage) > 0 {
			loadpg.NewCPURecord(usage)
		}

		// Ждём 10 минут до следующего сбора данных
		time.Sleep(10 * time.Minute)
	}
}
