package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
)

var bitacora []string //Ips de los nodos de la red
const (
	puerto_registro  = 8000
	puerto_notifica  = 8001
	puerto_procesoHP = 8002
)

var direccionIP_Nodo string

func ManejadorNotificacion(conn net.Conn) {
	defer conn.Close()
	bufferIn := bufio.NewReader(conn)
	IpNuevoNodo, _ := bufferIn.ReadString('\n')
	IpNuevoNodo = strings.TrimSpace(IpNuevoNodo)
	bitacora = append(bitacora, IpNuevoNodo)
	fmt.Println(bitacora)
}
func AtenderNotificaciones() {
	hostlocal := fmt.Sprintf("%s:%d", direccionIP_Nodo, puerto_notifica)
	ln, _ := net.Listen("tcp", hostlocal)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go ManejadorNotificacion(conn)
	}
}

func RegistrarSolicitud(ipConectar string) {
	hostremoto := fmt.Sprintf("%s:%d", ipConectar, puerto_registro)
	conn, _ := net.Dial("tcp", hostremoto)
	defer conn.Close()
	fmt.Fprintf(conn, "%s\n", direccionIP_Nodo)
	bufferIn := bufio.NewReader(conn)
	msgBitacora, _ := bufferIn.ReadString('\n')
	var arrAuxiliar []string
	json.Unmarshal([]byte(msgBitacora), &arrAuxiliar)
	bitacora = append(arrAuxiliar, ipConectar)
	fmt.Println(bitacora)
}

func Notificar(ipremoto, ipNuevoNodo string) {
	hostremoto := fmt.Sprintf("%s:%d", ipremoto, puerto_notifica)
	conn, _ := net.Dial("tcp", hostremoto)
	defer conn.Close()
	fmt.Fprintf(conn, "%s\n", ipNuevoNodo)
}

func NotificarTodos(ipNuevoNodo string) {
	for _, dirIp := range bitacora {
		Notificar(dirIp, ipNuevoNodo)
	}
}

func ManejadorSolicitudes(conn net.Conn) {
	defer conn.Close()
	bufferIn := bufio.NewReader(conn)
	ip, _ := bufferIn.ReadString('\n')
	ip = strings.TrimSpace(ip)
	bytesBitacora, _ := json.Marshal(bitacora)
	fmt.Fprintf(conn, "%s\n", string(bytesBitacora))
	NotificarTodos(ip)
	bitacora = append(bitacora, ip)
	fmt.Println(bitacora)
}

func AtenderSolicitudRegistro() {
	hostlocal := fmt.Sprintf("%s:%d", direccionIP_Nodo, puerto_registro)
	ln, _ := net.Listen("tcp", hostlocal)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go ManejadorSolicitudes(conn)
	}
}

func EnviarCargaSgteNodo(nIteraciones int, nActual int) {
	// enviar carga
	indice := rand.Intn(len(bitacora)) // selecciono de manera aleatoria
	hostremoto := fmt.Sprintf("%s:%d", bitacora[indice], puerto_procesoHP)
	fmt.Printf("Enviando la carga %d al nodo %s\n", nActual, bitacora[indice])
	conn, _ := net.Dial("tcp", hostremoto)
	defer conn.Close()
	fmt.Fprintf(conn, "%d,%d\n", nIteraciones, nActual)

}
func ManejadorServicioHP(conn net.Conn) {
	defer conn.Close()
	// leer la carga que llega al nodo
	bufferIn := bufio.NewReader(conn)
	load, _ := bufferIn.ReadString('\n')
	load = strings.TrimSpace(load)
	fmt.Printf("Llegó la carga: %s\n", load)
	s := strings.Split(load, ",")
	nIteracionesString := s[0]
	nActualString := s[1]
	nIteraciones, _ := strconv.Atoi(nIteracionesString)
	nActual, _ := strconv.Atoi(nActualString)
	fmt.Printf("Total de Iteraciones: %d. Iteracion Actual: %d\n", nIteraciones, nActual)
	if nActual == nIteraciones {
		fmt.Println("Inicializamos el algoritmo")
		EnviarCargaSgteNodo(nIteraciones, nActual-1)
	} else if nActual != nIteraciones && nActual != 0 {
		EnviarCargaSgteNodo(nIteraciones, nActual-1)
	} else {
		fmt.Println("LLegó a su fin, proceso terminado!!!!")
	}

}
func AtenderServicioHP() {
	hostlocal := fmt.Sprintf("%s:%d", direccionIP_Nodo, puerto_procesoHP)
	ln, _ := net.Listen("tcp", hostlocal)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go ManejadorServicioHP(conn)
	}
}

func main() {
	direccionIP_Nodo = localAddress()
	fmt.Println("IP: ", direccionIP_Nodo)
	go AtenderSolicitudRegistro()
	go AtenderServicioHP()
	bufferIn := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese la ip remota: ")
	ipConectar, _ := bufferIn.ReadString('\n')
	ipConectar = strings.TrimSpace(ipConectar)
	if ipConectar != "" {
		RegistrarSolicitud(ipConectar)
	}

	AtenderNotificaciones()
}

func localAddress() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("localAddress: %v\n", err.Error()))
		return "127.0.0.1"
	}
	for _, oiface := range ifaces {
		if strings.HasPrefix(oiface.Name, "Wi-Fi") {
			addrs, err := oiface.Addrs()
			if err != nil {
				log.Print(fmt.Errorf("localAddress: %v\n", err.Error()))
				continue
			}
			for _, dir := range addrs {
				switch d := dir.(type) {
				case *net.IPNet:
					if strings.HasPrefix(d.IP.String(), "192") {
						return d.IP.String()
					}

				}
			}
		}
	}
	return "127.0.0.1"
}
