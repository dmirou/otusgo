package telnet

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Client struct {
	network string
	address string
	timeout time.Duration

	dialer *net.Dialer
}

func NewClient(network, address string, timeout time.Duration) *Client {
	return &Client{
		network,
		address,
		timeout,
		&net.Dialer{},
	}
}

func (c Client) Run() error {
	ctx, _ := context.WithTimeout(context.Background(), c.timeout) // nolint: lostcancel

	conn, err := c.dialer.DialContext(
		ctx,
		"tcp",
		fmt.Sprintf("%s:%s", c.network, c.address),
	)
	if err != nil {
		return fmt.Errorf("cannot connect to the specified server: %v", err)
	}

	quit := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}

	wg.Add(1)

	go c.read(ctx, conn, wg, quit)

	wg.Add(1)

	go c.write(ctx, conn, wg, quit)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		log.Printf("main quit received")
	case <-sigChan:
		log.Printf("main terminate signal received ")
	}

	log.Printf("main before cancel called")
	cancel()
	log.Printf("main after cancel called")

	wg.Wait()

	conn.Close()
	log.Printf("main connection closed")

	return nil
}

func (c Client) read(ctx context.Context, conn io.Reader, wg *sync.WaitGroup, quit chan<- struct{}) {
	s := bufio.NewScanner(conn)

	defer wg.Done()
	defer log.Println("read: finished")

	for {
		select {
		case <-ctx.Done():
			log.Println("read: context done received")

			return
		default:
		}

		if !s.Scan() {
			log.Println("read: cannot scan")
			close(quit)

			return
		}

		str := s.Text()

		log.Printf("read: from server: %s", str)
	}
}

func (c Client) write(ctx context.Context, conn io.Writer, wg *sync.WaitGroup, quit chan<- struct{}) {
	r := bufio.NewReader(os.Stdin)

	defer wg.Done()
	defer log.Println("write: finished")

	for {
		select {
		case <-ctx.Done():
			log.Println("write: context done received")

			return
		default:
		}

		text, _, err := r.ReadLine()

		if err == io.EOF {
			log.Println("write: eof received from stdin")
			close(quit)

			continue
		}

		if err != nil {
			log.Printf("write: cannot read from stdin: %v\n", err)
			close(quit)

			return
		}

		_, err = conn.Write([]byte(fmt.Sprintf("%s\n", text)))
		if err != nil {
			log.Printf("write: error: %v", err)
			close(quit)

			return
		}

		log.Printf("write: to server: %s", text)
	}
}
