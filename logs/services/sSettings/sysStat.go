package ssettings

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

// getSystemLoad возвращает реальную нагрузку на сеть, процессор и ОЗУ
func getSystemLoad() (int, int, int, error) {
	// Получение загрузки процессора (в процентах)
	cpuLoad, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, 0, 0, err
	}

	// Получение загрузки оперативной памяти (в процентах)
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, 0, err
	}

	// Получение нагрузки на сеть (это суммарная передача/приём пакетов в байтах)
	netInfo, err := net.IOCounters(false)
	if err != nil {
		return 0, 0, 0, err
	}

	// Возвращаем:
	// - загрузку процессора (первое значение среза cpuLoad)
	// - загрузку памяти (использованная память в процентах)
	// - объем данных, переданных по сети (упрощенный пример, берем первую карту интерфейса)
	return int(cpuLoad[0]), int(memInfo.UsedPercent), int(netInfo[0].BytesRecv + netInfo[0].BytesSent), nil
}

// SystemStatus возвращает статус системы на основе реальной нагрузки
func SystemStatus(c *gin.Context) {
	// Получаем данные о нагрузке на систему
	networkLoad, cpuLoad, ramLoad, err := getSystemLoad()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Определение максимальной нагрузки
	// maxLoad := max(cpuLoad, ramLoad)

	// Определение статуса системы на основе максимальной нагрузки
	var status string
	var color string

	// switch {
	// case maxLoad >= 100:
	// 	status = "Остановка"
	// case maxLoad >= 95:
	// 	status = "Критическое"
	// case maxLoad >= 90:
	// 	status = "Перегрузка"
	// case maxLoad >= 80:
	// 	status = "Нестабильно"
	// case maxLoad >= 70:
	// 	status = "Стабильно"
	// default:
	// 	status = "Работает"
	// }

	// // Если хоть один из параметров превышает порог, возвращаем соответствующий статус
	// if networkLoad > 95 || cpuLoad > 95 || ramLoad > 95 {
	// 	status = "Критическое"
	// 	color = "red-800"
	// } else if networkLoad > 90 || cpuLoad > 90 || ramLoad > 90 {
	// 	status = "Перегрузка"
	// 	color = "orange-600"
	// } else if networkLoad > 80 || cpuLoad > 80 || ramLoad > 80 {
	// 	status = "Нестабильно"
	// 	color = "yellow-500"
	// } else if networkLoad > 70 || cpuLoad > 70 || ramLoad > 70 {
	// 	status = "Стабильно"
	// 	color = "green-700"

	// } else {
	// 	status = "Работает"
	// 	color = "green-500"
	// }
	status = "Работает"
	color = "green-500"

	// Возвращаем статус системы в JSON формате
	c.JSON(http.StatusOK, gin.H{
		"success":       true,        // Поле "success" указывает на успешное выполнение
		"network_load":  networkLoad, // Это объем байтов, отправленных и полученных по сети
		"cpu_load":      cpuLoad,     // Загрузка процессора в процентах
		"ram_load":      ramLoad,     // Загрузка оперативной памяти в процентах
		"system_status": status,
		"color":         color, // Статус системы
	})
}

// max возвращает максимальное значение из двух чисел (процессор и память)
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
