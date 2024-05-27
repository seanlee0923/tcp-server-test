package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 서버가 바인딩할 주소와 포트 지정
	address := "localhost:8080"

	// 주소를 TCP 주소로 변환
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// TCP 리스너 생성
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server is listening on", address)

	for {
		// 클라이언트 연결 수락
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		fmt.Println("Client connected:", conn.RemoteAddr().String())

		// 클라이언트와 통신을 처리하기 위해 고루틴 실행
		go handleClient(conn)
	}
}

// 클라이언트와의 통신을 처리하는 함수
func handleClient(conn *net.TCPConn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		// 클라이언트로부터 데이터 읽기
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		// 수신한 데이터 출력
		fmt.Println("Received data:", string(buf[:n]))

		// 에코 서버처럼 수신한 데이터를 다시 클라이언트로 전송
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
	}
}
