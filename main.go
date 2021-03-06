package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
)

type CLI struct {
	outStream, errStream io.Writer
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}

func (c *CLI) Run(args []string) int {
	if len(args) != 2 {
		fmt.Fprintf(c.errStream, "error: 第一引数にstatus codeを")
		return ExitCodeError
	}
	searchWord := args[1]
	description, err := findStatus(searchWord)
	if err != nil {
		fmt.Fprintf(c.errStream, "error: %s\n", err.Error())
		return ExitCodeError
	}
	fmt.Fprintf(c.outStream, "description: %s\n", description)
	return ExitCodeOK
}

func findStatus(searchWord string) (string, error) {
	i, err := strconv.Atoi(searchWord)
	if err != nil {
		return "", errors.New("引数は数値")
	}
	result := statuses()[i]
	if result == "" {
		return "", errors.New("存在しないステータスコード")
	}
	return result, nil
}

func statuses() map[int]string {
	m := map[int]string{
		100: "Continue",
		101: "Switching Protocols",
		102: "Processing", // RFC 2518 (WebDAV)
		200: "OK",
		201: "Created",
		202: "Accepted",
		203: "Non-Authoritative Information",
		204: "No Content",
		205: "Reset Content",
		206: "Partial Content",
		207: "Multi-Status",     // RFC 2518 (WebDAV)
		208: "Already Reported", // RFC 5842
		300: "Multiple Choices",
		301: "Moved Permanently",
		302: "Found",
		303: "See Other",
		304: "Not Modified",
		305: "Use Proxy",
		307: "Temporary Redirect",
		400: "Bad Request",
		401: "Unauthorized",
		402: "Payment Required",
		403: "Forbidden",
		404: "Not Found",
		405: "Method Not Allowed",
		406: "Not Acceptable",
		407: "Proxy Authentication Required",
		408: "Request Timeout",
		409: "Conflict",
		410: "Gone",
		411: "Length Required",
		412: "Precondition Failed",
		413: "Request Entity Too Large",
		414: "Request-URI Too Large",
		415: "Unsupported Media Type",
		416: "Request Range Not Satisfiable",
		417: "Expectation Failed",
		418: "I\"m a teapot",        // RFC 2324
		422: "Unprocessable Entity", // RFC 2518 (WebDAV)
		423: "Locked",               // RFC 2518 (WebDAV)
		424: "Failed Dependency",    // RFC 2518 (WebDAV)
		425: "No code",              // WebDAV Advanced Collections
		426: "Upgrade Required",     // RFC 2817
		428: "Precondition Required",
		429: "Too Many Requests",
		431: "Request Header Fields Too Large",
		449: "Retry with", // unofficial Microsoft
		500: "Internal Server Error",
		501: "Not Implemented",
		502: "Bad Gateway",
		503: "Service Unavailable",
		504: "Gateway Timeout",
		505: "HTTP Version Not Supported",
		506: "Variant Also Negotiates",  // RFC 2295
		507: "Insufficient Storage",     // RFC 2518 (WebDAV)
		509: "Bandwidth Limit Exceeded", // unofficial
		510: "Not Extended",             // RFC 2774
		511: "Network Authentication Required",
	}
	return m
}
