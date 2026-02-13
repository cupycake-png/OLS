package main

import (
	"bufio"
	"encoding/json"
	"log"
	"ols/lsp"
	"ols/rpc"
	"os"
)

func handleMessage(logger *log.Logger, msg []byte) {
	method, content, err := rpc.DecodeMessage(msg)

	logger.Println(msg)

	if err != nil {
		logger.Println("Error decoding message from client")
	}

	fileContents := ""

	switch method {

	case "initialize":
		_ = fileContents

		var request lsp.InitialiseRequest

		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("Error parsing %s", err)
		}

		logger.Printf("Connected to (%s, %s)", request.Params.ClientInfo.Name, *request.Params.ClientInfo.Version)

		// TODO: Check if token types and token modifiers are null before passing to NewInitialiseResponse

		msg := lsp.NewInitialiseResponse(request.ID, "utf-16", request.Params.Capabilities.TextDocument.SemanticTokens.TokenTypes, request.Params.Capabilities.TextDocument.SemanticTokens.TokenModifiers)
		reply := rpc.EncodeMessage(msg)

		_, err := os.Stdout.Write([]byte(reply))

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

		fileContents = notification.Params.TextDocument.Text

	case "textDocument/didChange":
		var notification lsp.DidChangeTextDocumentNotification

		if err := json.Unmarshal(content, &notification); err != nil {
			logger.Printf("Error parsing %s", err)
		}

		fileContents := notification.Params.ContentChanges[0].Text

		_ = fileContents

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

func getLogger(filename string) *log.Logger {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

	if err != nil {
		panic("error opening file")
	}

	return log.New(logFile, "[OLS] ", 0)
}

func main() {
	logger := getLogger("C:/Users/Charlie/Desktop/Programming/Go/OLS/logs.txt")

	logger.Println("Started language server")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()

		handleMessage(logger, msg)
	}
}
