run:
	go run cmd/app/main.go

localIpMac: 
	ipconfig getifaddr en0 