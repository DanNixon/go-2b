package powerbox

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/DanNixon/go-2b/pkg/types"
	"github.com/jacobsa/go-serial/serial"
)

type SerialPowerbox struct {
	Port               io.ReadWriteCloser
	Status             types.LockedStatus
	CommunicationMutex sync.Mutex
}

func NewSerialPowerbox(port string) (*SerialPowerbox, error) {
	c := serial.OpenOptions{
		PortName:              port,
		BaudRate:              9600,
		DataBits:              8,
		StopBits:              1,
		InterCharacterTimeout: 100,
	}

	var pb SerialPowerbox
	var err error

	pb.Port, err = serial.Open(c)

	if err == nil {
		go pb.backgroundUpdate()
	}

	return &pb, err
}

func (p *SerialPowerbox) Reset() (types.Status, error) {
	log.Printf("Resetting powerbox")
	p.CommunicationMutex.Lock()
	defer p.CommunicationMutex.Unlock()
	return p.sendCommandAndGetStatus(ResetCommand)
}

func (p *SerialPowerbox) Kill() (types.Status, error) {
	log.Printf("Killing powerbox")
	p.CommunicationMutex.Lock()
	defer p.CommunicationMutex.Unlock()
	return p.sendCommandAndGetStatus(KillCommand)
}

func (p *SerialPowerbox) Set(s types.Settings) (types.Status, error) {
	log.Printf("Setting powerbox output configuration")

	p.CommunicationMutex.Lock()
	defer p.CommunicationMutex.Unlock()

	p.Status.Mutex.Lock()
	st := p.Status.Status
	commands, err := GenerateDeltaCommands(p.Status.Status.Settings, s)
	p.Status.Mutex.Unlock()
	if err != nil {
		return st, err
	}

	for _, c := range commands {
		if st, err = p.sendCommandAndGetStatus(c); err != nil {
			return st, err
		}
	}

	if st.Settings != s {
		return st, errors.New("Powerbox reports different settings from requested after update")
	}

	return st, nil
}

func (p *SerialPowerbox) Get() (types.Status, error) {
	log.Printf("Serving powerbox status")
	return p.Status.Status, nil
}

func (p *SerialPowerbox) backgroundUpdate() {
	for {
		time.Sleep(1000 * time.Millisecond)
		log.Printf("Background status update...")
		p.CommunicationMutex.Lock()
		if _, err := p.sendCommandAndGetStatus(""); err != nil {
			log.Printf("Failed to fetch status: %v", err)
		}
		p.CommunicationMutex.Unlock()
	}
}

func (p *SerialPowerbox) sendCommandAndGetStatus(c string) (types.Status, error) {
	log.Printf("powerbox command = %s", c)

	var s types.Status

	commandData := []byte(c + "\n\r")
	log.Printf("powerbox command (hex): %s", hex.EncodeToString(commandData))

	txCount, err := p.Port.Write(commandData)
	if err != nil {
		return s, err
	}
	if txCount != len(commandData) {
		return s, errors.New(fmt.Sprintf("Did not send the expected amount of data (actual %d)", txCount))
	}

	rxString, err := p.serialRead()
	s, err = ParseStatusMessage(rxString)
	if err == nil {
		p.Status.Mutex.Lock()
		p.Status.Status = s
		p.Status.Mutex.Unlock()
		log.Printf("Powerbox status: %v", p.Status.Status)
	}

	return s, err
}

func ScanCRLF(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte{'\r', '\n'}); i >= 0 {
		// We have a full newline-terminated line.
		return i + 2, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

func (p *SerialPowerbox) serialRead() (string, error) {
	scanner := bufio.NewScanner(p.Port)
	scanner.Split(ScanCRLF)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	s := strings.Split(scanner.Text(), "\n")[0]
	log.Printf("powerbox reply = %s (length %d)", s, len(s))

	return s, nil
}
