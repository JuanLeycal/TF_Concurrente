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

	"github.com/JuanLeycal/TF_Concurrente/hotPotato/kmeans"
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

func EnviarCargaSgteNodo(encodedStr string, IP_API string, nIteraciones int, nActual int) {
	// enviar carga
	indice := rand.Intn(len(bitacora)) // selecciono de manera aleatoria
	hostremoto := fmt.Sprintf("%s:%d", bitacora[indice], puerto_procesoHP)
	fmt.Printf("Enviando la carga %d al nodo %s\n", nActual, bitacora[indice])
	conn, _ := net.Dial("tcp", hostremoto)
	defer conn.Close()
	//fmt.Fprintf(conn, "%d,%d,%s\n", nIteraciones, nActual, IP_API)
	fmt.Fprintf(conn, "%s/%s/%d/%d", encodedStr, IP_API, nIteraciones, nActual)
}

func EnviarCargaFinal(encodedStr string, IP_API string, nIteraciones int, nActual int) {
	// enviar carga
	hostremoto := fmt.Sprintf("%s:%d", IP_API, puerto_procesoHP)
	fmt.Printf("Enviando la carga %d de vuelta al API %s\n", nActual, IP_API)
	conn, _ := net.Dial("tcp", hostremoto)
	defer conn.Close()
	fmt.Fprintf(conn, encodedStr)
}

func ManejadorServicioHP(conn net.Conn) {
	defer conn.Close()
	// leer la carga que llega al nodo
	bufferIn := bufio.NewReader(conn)
	load, _ := bufferIn.ReadString('\n')
	//fmt.Printf("Llegó la carga: %s\n", load)
	load = strings.TrimSpace(load)
	//fmt.Printf("Llegó la carga: %s\n", load)
	s := strings.Split(load, "/")
	encodedStr := s[0]
	IP_API := s[1]
	nIteraciones, _ := strconv.Atoi(s[2]) //total de iteraciones
	nActual, _ := strconv.Atoi(s[3])      //numero actual de iteraciones
	fmt.Printf("Total de Iteraciones: %d. Iteracion Actual: %d, IP API:%s\n", nIteraciones, nActual, IP_API)
	//
	nList, nCentroids, maxIter := kmeans.TrainingDecode(encodedStr)
	nList, nCentroids, maxIter = kmeans.KMeansTraining(nList, nCentroids, maxIter)
	encodedStr = kmeans.TrainingEncode(nList, nCentroids, maxIter)

	if nActual != 0 {
		EnviarCargaSgteNodo(encodedStr, IP_API, nIteraciones, nActual-1)
	} else {
		fmt.Println("Se terminó el entrenamiento del algoritmo")
		EnviarCargaFinal(encodedStr, IP_API, nIteraciones, nActual)
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
		if strings.HasPrefix(oiface.Name, "ens33") {
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
