package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"net/netip"
	"os"
	"strconv"
	"strings"
	"sync"
)

type command struct {
	name        string
	args        string
	description string
}

var COMMANDS = []command{
	{"menu", "", "Show this menu"},
	{"create-room", "name", "Creates a new room"},
	{"connect", "ip", "Connects to a room with certain IP"},
	{"show-ip", "4|6", "Shows your own public IP(v4|v6)"},
	{"exit", "", "Exits the program"},
}

var PUBLIC_IP_PROVIDERS = []ip_provider{
	ip_provider{
		ipv:    4,
		url:    "https://ifcfg.me/",
		method: "GET",
		opt:    map[string]string{},
		res:    IFCFG{},
	},
	ip_provider{
		ipv:    6,
		url:    "https://api64.ipify.org",
		method: "GET",
		opt:    map[string]string{},
		res:    IPIFY{},
	},
}

type peer struct {
	ip   string
	port int
}

type room struct {
	name  string
	peers []peer
}

func main() {
	__main__show_menu()

	var command []string
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("command: ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Invalid input!")
				continue
			} else {
				fmt.Println("\nExiting...")
				break
			}
		}

		command = strings.Split(scanner.Text(), " ")

		switch command[0] {
		case "menu":
			__main__show_menu()

		case "show-ip":
			var ip netip.Addr
			var ipv int
			var err error

			ipv, err = strconv.Atoi(command[1])

			if err != nil {
				fmt.Println("It seems the second argument is not an integer")
				break
			}

			ip, _ = get_self_public_ip(PUBLIC_IP_PROVIDERS, int8(ipv))
			fmt.Printf("IP: %s.\n", ip.String())

		case "exit":
			fmt.Println("Bye bye")
			return

		default:
			fmt.Println("Unknown command. Try again")
			continue
		}
	}
	return
}

func __main__show_menu() {
	var text string
	var largest_text string = COMMANDS[0].name + " " + COMMANDS[0].args

	for i := 0; i < len(COMMANDS); i++ {
		next_text := COMMANDS[i].name + " " + COMMANDS[i].args

		if len(next_text) > len(largest_text) {
			largest_text = next_text
		}
	}

	for i := 0; i < len(COMMANDS); i++ {
		cmd := COMMANDS[i].name

		if len(COMMANDS[i].args) > 0 {
			cmd = cmd + " " + COMMANDS[i].args
		}

		diff := len(largest_text) - len(cmd)
		right_pad := strings.Repeat(".", diff+3)
		text = text + cmd + right_pad + COMMANDS[i].description + "\n"
	}

	fmt.Println(text)
	return
}

// main - sub procedures
func __main__create_config_dir() {
	_, err := os.Stat(CONFIG_DIR_PATH)

	if errors.Is(err, fs.ErrNotExist) {
		os.Mkdir(CONFIG_DIR_PATH, os.ModeDir)
	}
}

func __main__get_ips() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		ipv6 := get_self_ipv6(PUBLIC_IP_PROVIDERS)
		fmt.Printf("IP: %s is IPv4? %t.\n", ipv6.String(), ipv6.Is4())
	}()

	go func() {
		ipv4 := get_self_ipv4(PUBLIC_IP_PROVIDERS)
		fmt.Printf("IP: %s is IPv4? %t.\n", ipv4.String(), ipv4.Is4())
	}()

	wg.Wait()
}
