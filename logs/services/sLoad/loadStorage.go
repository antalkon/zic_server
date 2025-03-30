package sload

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/disk"
)

// formatSize - функция для форматирования размера в удобные единицы (ГБ, ТБ и т.д.)
func formatSize(bytes uint64) (float64, string) {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
		TB = GB * 1024
	)

	switch {
	case bytes >= TB:
		return float64(bytes) / TB, "TB"
	case bytes >= GB:
		return float64(bytes) / GB, "GB"
	case bytes >= MB:
		return float64(bytes) / MB, "MB"
	default:
		return float64(bytes) / KB, "KB"
	}
}

// StorageLoad - функция для получения данных о свободном и занятом месте на дисках сервера
func StorageLoad(c *gin.Context) {
	// Получаем информацию о дисках
	usage, err := disk.Usage("/")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve disk usage"})
		return
	}

	// Форматируем значения для отображения
	total, unit := formatSize(usage.Total)
	used, _ := formatSize(usage.Used)
	free, _ := formatSize(usage.Free)
	usedPercent := usage.UsedPercent
	freePercent := 100 - usedPercent

	// Возвращаем данные в формате JSON
	c.JSON(http.StatusOK, gin.H{
		"total":        fmt.Sprintf("%.2f %s", total, unit),
		"used":         fmt.Sprintf("%.2f %s", used, unit),
		"free":         fmt.Sprintf("%.2f %s", free, unit),
		"unit":         unit,
		"used_percent": fmt.Sprintf("%.2f%%", usedPercent),
		"free_percent": fmt.Sprintf("%.2f%%", freePercent),
	})
}
