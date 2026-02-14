package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"ols/lsp"
	"ols/rpc"
	"os"
	"strings"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_python "github.com/tree-sitter/tree-sitter-python/bindings/go"
)

func sendMessage(connection net.Conn, msg any, logger *log.Logger) {
	_, err := connection.Write([]byte(rpc.EncodeMessage(msg)))

	if err != nil {
		logger.Printf("Error writing to connection: %s", err.Error())
	}
}

func handleMessage(logger *log.Logger, connection net.Conn, msg []byte) {
	method, content, err := rpc.DecodeMessage(msg)

	logger.Printf("Received message with method %s: %s\n", method, content)

	if err != nil {
		logger.Println("Error decoding message from client")
	}

	fileContents := ""
	uri := ""
	version := 0

	parser := tree_sitter.NewParser()

	defer parser.Close()

	switch method {

	case "initialize":
		_ = uri

		var request lsp.InitialiseRequest

		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("Error parsing %s", err)
		}

		logger.Printf("Connected to (%s, %s)", request.Params.ClientInfo.Name, *request.Params.ClientInfo.Version)

		msg := lsp.NewInitialiseResponse(*request.ID, "utf-16", request.Params.Capabilities.TextDocument.SemanticTokens.TokenTypes, request.Params.Capabilities.TextDocument.SemanticTokens.TokenModifiers)

		sendMessage(connection, msg, logger)

		if err != nil {
			logger.Printf("Error writing initialise reply %s", err.Error())
		} else {
			logger.Printf("Sent initialise reply")
		}

	case "textDocument/didOpen":
		var notification lsp.DidOpenTextDocumentNotification

		if err := json.Unmarshal(content, &notification); err != nil {
			logger.Printf("Error parsing %s", err)
		}

		uri = notification.Params.TextDocument.URI
		version = notification.Params.TextDocument.Version

		fileContents = notification.Params.TextDocument.Text

		logger.Printf("Received textDocument/didOpen notification with contents: %s", fileContents)

		// TODO: Change the parsing language depending on file type
		language := tree_sitter.NewLanguage(tree_sitter_python.Language())

		parser.SetLanguage(language)

		tree := parser.Parse([]byte(fileContents), nil)

		defer tree.Close()

		errorQuery, err := tree_sitter.NewQuery(language, "((ERROR) @error)")

		if err != nil {
			logger.Printf("Error created query: %s", err.Error())
		}

		defer errorQuery.Close()

		errorQueryCursor := tree_sitter.NewQueryCursor()

		defer errorQueryCursor.Close()

		matches := errorQueryCursor.Matches(errorQuery, tree.RootNode(), []byte(fileContents))

		var diagnostics []lsp.Diagnostic

		for match := matches.Next(); match != nil; match = matches.Next() {
			for _, capture := range match.Captures {
				logger.Printf("Error info: %s from position (%d, %d) to position (%d, %d)", capture.Node.Utf8Text([]byte(fileContents)), capture.Node.StartPosition().Row, capture.Node.StartPosition().Column, capture.Node.EndPosition().Row, capture.Node.EndPosition().Column)

				// 1 is the severity code for errors
				severity := 1

				diagnostic := lsp.Diagnostic{
					Range:    lsp.Range{Start: lsp.Position{Line: capture.Node.StartPosition().Row - 1, Character: capture.Node.StartPosition().Column - 1}, End: lsp.Position{Line: capture.Node.EndPosition().Row - 1, Character: capture.Node.EndPosition().Column - 1}},
					Severity: &severity,
					Message:  fmt.Sprintf("Error on line %d", capture.Node.StartPosition().Row),
				}

				diagnostics = append(diagnostics, diagnostic)
			}
		}

		msg := lsp.NewPublishDiagnosticsNotification(uri, version, diagnostics)

		sendMessage(connection, msg, logger)

	case "textDocument/didChange":
		var notification lsp.DidChangeTextDocumentNotification

		if err := json.Unmarshal(content, &notification); err != nil {
			logger.Printf("Error parsing %s", err)
		}

		version = notification.Params.TextDocument.Version

		splitContents := strings.Split(fileContents, "\n")
		splitContents[notification.Params.ContentChanges[0].Range.Start.Line] = notification.Params.ContentChanges[0].Text

		fileContents = strings.Join(splitContents, "\n")

		logger.Printf("Received textDocument/didChange notification with new content: %s", fileContents)

	case "textDocument/semanticTokens/full":
		var request lsp.SemanticTokensFullRequest

		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("Error parsing %s", err)
		}

		logger.Print(request.Params.TextDocument.URI)

	case "shutdown":
		logger.Printf("Shutdown request received")

		// TODO: handle shutdown request, not really sure what it does tbh
	}

}

func handleConnection(logger *log.Logger, connection net.Conn) {
	defer connection.Close()

	scanner := bufio.NewScanner(connection)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		handleMessage(logger, connection, msg)
	}

	if err := scanner.Err(); err != nil {
		logger.Printf("Connection error: %s", err)
	}
}

func getLogger(filename string) *log.Logger {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

	if err != nil {
		panic("error opening file")
	}

	return log.New(logFile, "[OLS] ", 0)
}

func main() {
	logger := getLogger("OLS_logs.txt")

	logger.Println("Started language server")

	listener, err := net.Listen("tcp", "127.0.0.1:2956")

	if err != nil {
		logger.Printf("ERROR: %s", err.Error())
		log.Fatal(err)
	}

	defer listener.Close()

	logger.Println("Listening on 127.0.0.1:2956")

	for {
		connection, err := listener.Accept()

		if err != nil {
			logger.Printf("Accept erroor: %s", err)
			continue
		}

		logger.Println("Client connected")

		go handleConnection(logger, connection)
	}
}
