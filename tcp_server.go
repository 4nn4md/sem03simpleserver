package main

import (
	"io"
	"log"
	"net"
	"sync"
	"tmp/mycrypt"
	"tmp/conv"
	"strings"
	"strconv"
)

func main() {

	var wg sync.WaitGroup

	server, err := net.Listen("tcp", "172.17.0.4:1000")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("bundet til %s", server.Addr().String())
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			log.Println("før server.Accept() kallet")
			conn, err := server.Accept()
			if err != nil {
				return
			}
			go func(c net.Conn) {
				defer c.Close()
				for {
					buf := make([]byte, 1024)
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							log.Println(err)
						}
						return // fra for løkke
					}

alfLength := len(mycrypt.ALF_SEM03)
hentetMelding := string(buf[:n])
dekryptertMelding := mycrypt.Krypter([]rune(hentetMelding), mycrypt.ALF_SEM03, alfLength-4)
dekryptertString := string(dekryptertMelding)

					switch msg := dekryptertString; msg {
  				        case "ping":
						_, err = c.Write([]byte("pong"))


					case dekryptertString:
						a1 := strings.Split((msg), ";")
						a2, err := strconv.ParseFloat(a1[3], 64)
						a3 := conv.CelsiusToFahrenheit(a2)
						a4 := strconv.FormatFloat(a3, 'f', 2, 64)
						a5 := (a1[0] + ";" + a1[1] + ";" + a1[2] + ";" + a4)

							if err != nil {
								panic(err)
							}


					//kryptert tilbake
					kryptertBobRune := mycrypt.Krypter([]rune(a5), mycrypt.ALF_SEM03, 4)
					kryptertBobString := string(kryptertBobRune)

					_, err = c.Write([]byte(kryptertBobString))

					default:
						_, err = c.Write(buf[:n])
					}
					if err != nil {
						if err != io.EOF {
							log.Println(err)
						}
						return // fra for løkke
					}
				}
			}(conn)
		}
	}()
	wg.Wait()
}
