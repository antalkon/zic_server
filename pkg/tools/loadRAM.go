package tools

import (
	"log"
	"time"

	loadpg "github.com/antalkon/zic_server/internal/db/db_pg/load_pg"
	"github.com/shirou/gopsutil/mem"
)

func LoadRAM() {
	for {
		vmStat, err := mem.VirtualMemory()
		if err != nil {
			log.Printf("Error getting RAM usage: %v", err)
			time.Sleep(10 * time.Minute)
			continue
		}

		// Сохраняем данные в таблицу
		err = loadpg.NewRAMRecord(vmStat.UsedPercent)
		if err != nil {
			log.Printf("Error saving RAM usage to database: %v", err)
		}

		// Ждём 10 минут до следующего сбора данных
		time.Sleep(10 * time.Minute)
	}
}
