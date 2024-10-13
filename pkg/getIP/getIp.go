package getip

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

func GetIps() (string, string, error) {
	// Получение публичного IP
	publicIP, err := getPublicIP()
	if err != nil {
		return "", "", err
	}

	// Получение локального IP
	localIP, err := getLocalIP()
	if err != nil {
		return publicIP, "", err // Возвращаем публичный IP и ошибку по локальному
	}

	return publicIP, localIP, nil
}

// Вспомогательная функция для получения публичного IP
func getPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Вспомогательная функция для получения локального IP
func getLocalIP() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue // Интерфейс не активен или является loopback
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}

			ip = ip.To4()
			if ip == nil {
				continue // Не IPv4 адрес
			}

			return ip.String(), nil
		}
	}
	return "", fmt.Errorf("не удалось найти локальный IP")
}
