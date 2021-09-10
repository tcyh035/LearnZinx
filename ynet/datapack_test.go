package ynet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {
	// 模拟服务端
	listenner, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen failed", err)
		return
	}

	go func() {
		for {
			fmt.Println("Bingo")
			conn, err := listenner.Accept()
			if err != nil {
				fmt.Println("server accept error", err)
				return
			}

			go func(conn net.Conn) {
				dp := NewDataPack()
				// 第一次读， 从conn读，把包的head读出来
				headData := make([]byte, dp.GetHeadLen())
				for {
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head error", err)
						return
					}

					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server unpack err", err)
						return
					}

					// 第二次读，从conn读，把包head中的datalen读出来，再读取data内容
					if msgHead.GetDataLen() > 0 {
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetDataLen())
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("server unpack data, err", err)
							return
						}

						fmt.Println("----> Recv MsgID: ", msg.Id, "datalen =", msg.DataLen, "data =", string(msg.Data))
					}
				}
			}(conn)
		}
	}()

	// 模拟客户端
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err", err)
		return
	}

	dp := NewDataPack()

	// 模拟粘包，封装两个msg到一起
	msg1 := &Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}

	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 err", err)
	}

	msg2 := &Message{
		Id:      2,
		DataLen: 5,
		Data:    []byte{'w', 'o', 'r', 'l', 'd'},
	}

	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg2 err", err)
	}

	sendData1 = append(sendData1, sendData2...)

	conn.Write(sendData1)

	select {}
}
