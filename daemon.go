package log

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// StartDaemon starts the log daemon for setting log level
func StartDaemon() {
	go std.daemon.start()
}

type daemon struct {
	Port int
}

func (d *daemon) start() {
	c, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	d.Port = c.Addr().(*net.TCPAddr).Port
	d.logPort()
	for {
		sock, err := c.Accept()
		if err != nil {
			Error(err)
			continue
		}
		go chatter{sock}.chat()
	}
}

func (d *daemon) logPort() {
	Fields{
		"service": "Log Daemon",
		"port":    d.Port,
	}.Info("Log Daemon is up and listening")
}

type chatter struct {
	c net.Conn
}

func (d chatter) chat() {
	d.help()
	for {
		r := bufio.NewReader(d.c)
		b, _, err := r.ReadLine()
		if err != nil {
			Error(err)
			return
		}

		line := strings.TrimSpace(string(b))
		parts := strings.Split(line, " ")
		cmd, para := "", ""
		switch len(parts) {
		case 1:
			cmd = parts[0]
		case 2:
			cmd, para = parts[0], parts[1]
		default:
			d.unsupported(line)
		}
		switch cmd {
		case "help":
			d.help()
		case "quit":
			d.c.Close()
			return
		case "show":
			switch para {
			case "level":
				d.echo(GetLevel().String())
			default:
				d.unsupported(line)
			}
		case "level":
			switch para {
			case "panic":
				std.SetLevel(PanicLevel)
			case "fatal":
				std.SetLevel(FatalLevel)
			case "error":
				std.SetLevel(ErrorLevel)
			case "warn":
				std.SetLevel(WarnLevel)
			case "info":
				std.SetLevel(InfoLevel)
			case "debug":
				std.SetLevel(DebugLevel)
			default:
				d.unsupported(line)
			}
		default:
			d.unsupported(line)
		}
	}
}

func (d chatter) unsupported(line string) {
	Fields{
		"service": "Log Daemon",
		"req":     line,
	}.Error("unsupported command")
	d.echo("unsupported command")
}

func (d chatter) echo(s string) {
	fmt.Fprintln(d.c, s)
}

func (d chatter) help() {
	d.echo(`
Log Daemon Usage:

quit
    quit the session.
show level
    print the current level.
level (panic|fatal|error|warn|info|debug)
    change the current level.
help
    print the help message.
`)
}
