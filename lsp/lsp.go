package lsp

type Message struct {
	JSONRPC string `json:"jsonrpc"`
}

type RequestMessage struct {
	Message
	ID     int    `json:"id"`
	Method string `json:"method"`
}

type ResponseError struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    *string `json:"data"`
}

type ResponseMessage struct {
	Message
	ID *int `json:"id"`
}

type NotificationMessage struct {
	Message
	Method string `json:"method"`
}

type ClientInfo struct {
	Name    string  `json:"name"`
	Version *string `json:"version"`
}

type DocumentURI = string

type ResourceOperationKind = string

type FailureHandlingKind = string

type ChangeAnnotationSupport struct {
	GroupsOnLabel *bool `json:"groupsOnLabel"`
}

type WorkspaceEditClientCapabilities struct {
	DocumentChanges         *bool                    `json:"documentChanges"`
	ResourceOperations      *[]ResourceOperationKind `json:"resourceOperations"`
	FailureHandling         *FailureHandlingKind     `json:"failureHandling"`
	NormaliseLineEndings    *bool                    `json:"normalizesLineEndings"`
	ChangeAnnotationSupport *ChangeAnnotationSupport `json:"changeAnnotationSupport"`
}

type DidChangeConfigurationClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

type DidChangeWatchedFilesClientCapabilities struct {
	DynamicRegistration    *bool `json:"dynamicRegistration"`
	RelativePatternSupport *bool `json:"relativePatternSupport"`
}

type SymbolKind struct {
	ValueSet *[]int `json:"valueSet"`
}

type TagSupport struct {
	ValueSet []int `json:"valueSet"`
}

type ResolveSupport struct {
	Properties []string `json:"properties"`
}

type WorkspaceSymbolClientCapabilities struct {
	DynamicRegistration *bool           `json:"dynamicRegistration"`
	SymbolKind          *SymbolKind     `json:"symbolKind"`
	TagSupport          *TagSupport     `json:"tagSupport"`
	ResolveSupport      *ResolveSupport `json:"resolveSupport"`
}

type ExecuteCommandClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

type SemanticTokensWorkspaceClientCapabilities struct {
	RefreshSupport *bool `json:"refreshSupport"`
}

type CodeLensWorkspaceClientCapabilities struct {
	RefreshSupport *bool `json:"refreshSupport"`
}

type FileOperations struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
	DidCreate           *bool `json:"didCreate"`
	WillCreate          *bool `json:"willCreate"`
	DidRename           *bool `json:"didRename"`
	WillRename          *bool `json:"willRename"`
	DidDelete           *bool `json:"didDelete"`
	WillDelete          *bool `json:"willDelete"`
}

type InlineValueWorkspaceClientCapabilities struct {
	RefreshSupport *bool `json:"refreshSupport"`
}

type InlayHintWorkspaceClientCapabilities struct {
	RefreshSupport *bool `json:"refreshSupport"`
}

type DiagnosticWorkspaceClientCapabilities struct {
	RefreshSupport *bool `json:"refreshSupport"`
}

type Workspace struct {
	ApplyEdit              *bool                                      `json:"applyEdit"`
	WorkspaceEdit          *WorkspaceEditClientCapabilities           `json:"workspaceEdit"`
	DidChangeConfiguration *DidChangeConfigurationClientCapabilities  `json:"didChangeConfiguration"`
	DidChangeWatchedFiles  *DidChangeWatchedFilesClientCapabilities   `json:"didChangeWatchedFiles"`
	Symbol                 *WorkspaceSymbolClientCapabilities         `json:"symbol"`
	ExecuteCommand         *ExecuteCommandClientCapabilities          `json:"executeCommand"`
	WorkspaceFolders       *bool                                      `json:"workspaceFolders"`
	Configuration          *bool                                      `json:"configuration"`
	SemanticTokens         *SemanticTokensWorkspaceClientCapabilities `json:"semanticTokens"`
	CodeLens               *CodeLensWorkspaceClientCapabilities       `json:"codeLens"`
	FileOperations         *FileOperations                            `json:"fileOperations"`
	InlineValue            *InlineValueWorkspaceClientCapabilities    `json:"inlineValue"`
	InlayHint              *InlayHintWorkspaceClientCapabilities      `json:"inlayHint"`
	Diagnostics            *DiagnosticWorkspaceClientCapabilities     `json:"diagnostics"`
}

type TextDocumentSyncClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
	WillSave            *bool `json:"willSave"`
	WillSaveWaitUntil   *bool `json:"willSaveWaitUntil"`
	DidSave             *bool `json:"didSave"`
}

type MarkupKind = string

type InsertTextMode = int

type InsertTextModeSupport struct {
	ValueSet []InsertTextMode `json:"valueSet"`
}

type CompletionItem struct {
	SnippetSupport          *bool                  `json:"snippetSupport"`
	CommitCharactersSupport *bool                  `json:"commitCharactersSupport"`
	DocumentationFormat     *[]MarkupKind          `json:"documentationFormat"`
	DeprecatedSupport       *bool                  `json:"deprecatedSupport"`
	PreselectSupport        *bool                  `json:"preselectSupport"`
	TagSupport              *TagSupport            `json:"tagSupport"`
	InsertReplaceSupport    *bool                  `json:"insertReplaceSupport"`
	ResolveSupport          *ResolveSupport        `json:"resolveSupport"`
	InsertTextModeSupport   *InsertTextModeSupport `json:"insertTextModeSupport"`
	LabelDetailsSupport     *bool                  `json:"labelDetailsSupport"`
}

type CompletionItemKind struct {
	ValueSet []int `json:"valueSet"`
}

type CompletionList struct {
	ItemDefaults *[]string `json:"itemDefaults"`
}

type CompletionClientCapabilities struct {
	DynamicRegistration *bool               `json:"dynamicRegistration"`
	CompletionItem      *CompletionItem     `json:"completionItem"`
	CompletionItemKind  *CompletionItemKind `json:"completionItemKind"`
	ContextSupport      *bool               `json:"contextSupport"`
	InsertTextMode      *InsertTextMode     `json:"insertTextMode"`
	CompletionList      *CompletionList     `json:"completionList"`
}

type HoverClientCapabilities struct {
	DynamicRegistration *bool         `json:"dynamicRegistration"`
	ContentFormat       *[]MarkupKind `json:"contentFormat"`
}

type ParameterInformation struct {
	LabelOffsetSupport *bool `json:"labelOffsetSupport"`
}

type SignatureInformation struct {
	DocumentationFormat    *[]MarkupKind         `json:"documentationFormat"`
	ParameterInformation   *ParameterInformation `json:"parameterInformation"`
	ActiveParameterSupport *bool                 `json:"activeParameterSupport"`
}

type SignatureHelpClientCapabilities struct {
	DynamicRegistration  *bool                 `json:"dynamicRegistration"`
	SignatureInformation *SignatureInformation `json:"signatureInformation"`
	ContextSupport       *bool                 `json:"contextSupport"`
}

type DeclarationClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
	LinkSupport         *bool `json:"linkSupport"`
}

type DefinitionClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
	LinkSupport         *bool `json:"linkSupport"`
}

type TypeDefinitionClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
	LinkSupport         *bool `json:"linkSupport"`
}

type ImplementationClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
	LinkSupport         *bool `json:"linkSupport"`
}

type ReferenceClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

type DocumentHighlightClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

type DocumentSymbolClientCapabilities struct {
	DynamicRegistration               *bool       `json:"dynamicRegistration"`
	SymbolKind                        *SymbolKind `json:"symbolKind"`
	HierarchicalDocumentSymbolSupport *bool       `json:"hierarchicalDocumentSymbolSupport"`
	TagSupport                        *TagSupport `json:"tagSupport"`
	LabelSupport                      *bool       `json:"labelSupport"`
}

type CodeActionKind struct {
	ValueSet []string `json:"valueSet"`
}

type CodeActionLiteralSupport struct {
	CodeActionKind CodeActionKind `json:"codeActionKind"`
}

type CodeActionClientCapabilities struct {
	DynamicRegistration      *bool                     `json:"dynamicRegistration"`
	CodeActionLiteralSupport *CodeActionLiteralSupport `json:"codeActionLiteralSupport"`
	IsPreferredSupport       *bool                     `json:"isPreferredSupport"`
	DisabledSupport          *bool                     `json:"disabledSupport"`
	DataSupport              *bool                     `json:"dataSupport"`
	ResolveSupport           *ResolveSupport           `json:"resolveSupport"`
	HonoursChangeAnnotations *bool                     `json:"honorsChangeAnnotation"`
}

type CodeLensClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

type DocumentLinkClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
	TooltipSupport      *bool `json:"tooltipSupport"`
}

type DocumentColourClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

type DocumentFormattingClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

type DocumentRangeFormattingClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

type DocumentOnTypeFormattingClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

type PrepareSupportDefaultBehavior = int

type RenameClientCapabilities struct {
	DynamicRegistration           *bool                          `json:"dynamicRegistration"`
	PrepareSupport                *bool                          `json:"prepareSupport"`
	PrepareSupportDefaultBehavior *PrepareSupportDefaultBehavior `json:"prepareSupportDefaultBehavior"`
	HonoursChangeAnnotations      *bool                          `json:"honorsChangeAnnotations"`
}

type PublishDiagnosticsClientCapabilities struct {
	RelatedInformation     *bool       `json:"relatedInformation"`
	TagSupport             *TagSupport `json:"tagSupport"`
	VersionSupport         *bool       `json:"versionSupport"`
	CodeDescriptionSupport *bool       `json:"codeDescriptionSupport"`
	DataSupport            *bool       `json:"dataSupport"`
}

type FoldingRangeKind struct {
	ValueSet *[]string `json:"valueSet"`
}

type FoldingRange struct {
	CollapsedText *bool `json:"collapsedText"`
}

type FoldingRangeClientCapabilities struct {
	DynamicRegistration *bool             `json:"dynamicRegistration"`
	RangeLimit          *uint             `json:"rangeLimit"`
	LineFoldingOnly     *bool             `json:"lineFoldingOnly"`
	FoldingRangeKind    *FoldingRangeKind `json:"foldingRangeKind"`
	FoldingRange        *FoldingRange     `json:"foldingRange"`
}

type SelectionRangeClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

type LinkedEditingRangeClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

type CallHierarchyClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

type TokenFormat = string

type SemanticTokensClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`

	// TODO: Implement 'requests' field

	TokenTypes              []string      `json:"tokenTypes"`
	TokenModifiers          []string      `json:"tokenModifiers"`
	Formats                 []TokenFormat `json:"formats"`
	OverlappingTokenSupport *bool         `json:"overlappingTokenSupport"`
	MultilineTokenSupport   *bool         `json:"multilineTokenSupport"`
	ServerCancelSupport     *bool         `json:"serverCancelSupport"`
	AugmentsSyntaxTokens    *bool         `json:"augmentsSyntaxTokens"`
}

type MonikerClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

type TypeHierarchyClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

type InlineValueClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration"`
}

type InlayHintClientCapabilities struct {
	DynamicRegistration *bool           `json:"dynamicRegistration"`
	ResolveSupport      *ResolveSupport `json:"resolveSupport"`
}

type DiagnosticClientCapabilities struct {
	DynamicRegistration    *bool `json:"dynamicRegistration"`
	RelatedDocumentSupport *bool `json:"relatedDocumentSupport"`
	RelatedInformation     *bool `json:"relatedInformation"`

	// TODO: Implement 'tagSupport' field

	CodeDescriptionSupport *bool `json:"codeDescriptionSupport"`
	DataSupport            *bool `json:"dataSupport"`
}

type TextDocumentClientCapabilities struct {
	Synchronisation    *TextDocumentSyncClientCapabilities         `json:"synchronization"`
	Completion         *CompletionClientCapabilities               `json:"completion"`
	Hover              *HoverClientCapabilities                    `json:"hover"`
	SignatureHelp      *SignatureHelpClientCapabilities            `json:"signatureHelp"`
	Declaration        *DeclarationClientCapabilities              `json:"declaration"`
	Definition         *DefinitionClientCapabilities               `json:"definition"`
	TypeDefinition     *TypeDefinitionClientCapabilities           `json:"typeDefinition"`
	Implementation     *ImplementationClientCapabilities           `json:"implementation"`
	References         *ReferenceClientCapabilities                `json:"references"`
	DocumentHighlight  *DocumentHighlightClientCapabilities        `json:"documentHighlight"`
	DocumentSymbol     *DocumentSymbolClientCapabilities           `json:"documentSymbol"`
	CodeAction         *CodeActionClientCapabilities               `json:"codeAction"`
	CodeLens           *CodeLensClientCapabilities                 `json:"codeLens"`
	DocumentLink       *DocumentLinkClientCapabilities             `json:"documentLink"`
	ColourProvider     *DocumentColourClientCapabilities           `json:"colorProvider"`
	Formatting         *DocumentFormattingClientCapabilities       `json:"formatting"`
	RangeFormatting    *DocumentRangeFormattingClientCapabilities  `json:"rangeFormatting"`
	OnTypeFormatting   *DocumentOnTypeFormattingClientCapabilities `json:"onTypeFormatting"`
	Rename             *RenameClientCapabilities                   `json:"rename"`
	PublishDiagnostics *PublishDiagnosticsClientCapabilities       `json:"publishDiagnostics"`
	FoldingRange       *FoldingRangeClientCapabilities             `json:"foldingRange"`
	SelectionRange     *SelectionRangeClientCapabilities           `json:"selectionRange"`
	LinkedEditingRange *LinkedEditingRangeClientCapabilities       `json:"linkedEditingRange"`
	CallHierarchy      *CallHierarchyClientCapabilities            `json:"callHierarchy"`
	SemanticTokens     *SemanticTokensClientCapabilities           `json:"semanticTokens"`
	Moniker            *MonikerClientCapabilities                  `json:"moniker"`
	TypeHierarchy      *TypeHierarchyClientCapabilities            `json:"typeHierarchy"`
	InlineValue        *InlineValueClientCapabilities              `json:"inlineValue"`
	InlayHint          *InlayHintClientCapabilities                `json:"inlayHint"`
	Diagnostic         *DiagnosticClientCapabilities               `json:"diagnostic"`
}

type NotebookDocumentSyncClientCapabilities struct {
	DynamicRegistration     *bool `json:"dynamicRegistration"`
	ExecutionSummarySupport *bool `json:"executionSummarySupport"`
}

type NotebookDocumentClientCapabilities struct {
	Synchronisation NotebookDocumentSyncClientCapabilities `json:"synchronisation"`
}

type MessageActionItem struct {
	AdditionalPropertiesSupport *bool `json:"additionalPropertiesSupport"`
}

type ShowMessageRequestClientCapabilities struct {
	MessageActionItem *MessageActionItem `json:"messageActionItem"`
}

type ShowDocumentClientCapabilities struct {
	Support bool `json:"support"`
}

type Window struct {
	WorkDoneProgress *bool                                 `json:"workDoneProgress"`
	ShowMessage      *ShowMessageRequestClientCapabilities `json:"showMessage"`
	ShowDocument     *ShowDocumentClientCapabilities       `json:"showDocument"`
}

type StaleRequestSupport struct {
	Cancel                 bool     `json:"cancel"`
	RetryOnContentModified []string `json:"retryOnContentModified"`
}

type RegularExpressionsClientCapabilities struct {
	Engine  string  `json:"engine"`
	Version *string `json:"version"`
}

type MarkdownClientCapabilities struct {
	Parser      string    `json:"parser"`
	Version     *string   `json:"version"`
	AllowedTags *[]string `json:"allowedTags"`
}

type PositionEncodingKind = string

type General struct {
	StaleRequestSupport *StaleRequestSupport                  `json:"staleRequestSupport"`
	RegularExpressions  *RegularExpressionsClientCapabilities `json:"regularExpressions"`
	Markdown            *MarkdownClientCapabilities           `json:"markdown"`
	PositionEncodings   *[]PositionEncodingKind               `json:"positionEncodings"`
}

type ClientCapabilities struct {
	Workspace        *Workspace                          `json:"workspace"`
	TextDocument     *TextDocumentClientCapabilities     `json:"textDocument"`
	NotebookDocument *NotebookDocumentClientCapabilities `json:"NotebookDocumentClientCapabilities"`
	Window           *Window                             `json:"window"`
	General          *General                            `json:"general"`
	Experimental     *string                             `json:"experimental"`
}

type TraceValue = string

type URI = string

type WorkspaceFolder struct {
	URI  URI    `json:"uri"`
	Name string `json:"name"`
}

type InitialiseParams struct {
	ProcessID             *int               `json:"processID"`
	ClientInfo            *ClientInfo        `json:"clientInfo"`
	Locale                *string            `json:"locale"`
	RootPath              *string            `json:"rootPath"`
	RootURI               *DocumentURI       `json:"rootUri"`
	InitialisationOptions *any               `json:"initializationOptions"`
	Capabilities          ClientCapabilities `json:"clientCapabilities"`
	Trace                 *TraceValue        `json:"trace"`
	WorkspaceFolders      *[]WorkspaceFolder `json:"workspaceFolders"`
}

type InitialiseRequest struct {
	RequestMessage
	Params InitialiseParams `json:"params"`
}

type TextDocumentSyncKind = int

type TextDocumentSyncOptions struct {
	OpenClose *bool                 `json:"openClose,omitempty"`
	Change    *TextDocumentSyncKind `json:"change,omitempty"`
}

type SemanticTokensLegend struct {
	TokenTypes     []string `json:"tokenTypes"`
	TokenModifiers []string `json:"tokenModifiers"`
}

type SemanticTokensOptions struct {
	Legend SemanticTokensLegend `json:"legend"`
	Range  *bool                `json:"range"`
	Full   *bool                `json:"full"`
}

type ServerCapabilities struct {
	PositionEncoding       *PositionEncodingKind    `json:"positionEncoding,omitempty"`
	TextDocumentSync       *TextDocumentSyncOptions `json:"textDocumentSync,omitempty"`
	SemanticTokensProvider *SemanticTokensOptions   `json:"semanticTokenOptions,omitempty"`

	// TODO: Implement rest of ServerCapabilities
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitialiseResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type InitialiseResponse struct {
	ResponseMessage
	Result InitialiseResult `json:"result,omitempty"`
	Error  *ResponseError   `json:"error,omitempty"`
}

type TextDocumentItem struct {
	URI        DocumentURI `json:"uri"`
	LanguageID string      `json:"languageId"`
	Version    int         `json:"version"`
	Text       string      `json:"text"`
}

type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

type DidOpenTextDocumentNotification struct {
	NotificationMessage
	Params DidOpenTextDocumentParams `json:"params"`
}

type TextDocumentIdentifier struct {
	URI DocumentURI `json:"uri"`
}

type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}

type Position struct {
	Line      uint `json:"line"`
	Character uint `json:"character"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

type TextDocumentContentChangeEvent struct {
	Range Range  `json:"range"`
	Text  string `json:"text"`
}

type DidChangeTextDocumentParams struct {
	TextDocument   VersionedTextDocumentIdentifier  `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type DidChangeTextDocumentNotification struct {
	NotificationMessage
	Params DidChangeTextDocumentParams `json:"params"`
}

type SemanticTokenParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

type SemanticTokensFullRequest struct {
	RequestMessage
	Params SemanticTokenParams `json:"params"`
}

func NewInitialiseResponse(id int, positionEncoding string, tokenTypes []string, tokenModifiers []string) InitialiseResponse {
	openClose := true
	change := 1

	return InitialiseResponse{
		ResponseMessage: ResponseMessage{
			Message: Message{
				JSONRPC: "2.0",
			},
			ID: &id,
		},

		Result: InitialiseResult{
			Capabilities: ServerCapabilities{
				PositionEncoding: &positionEncoding,
				TextDocumentSync: &TextDocumentSyncOptions{
					OpenClose: &openClose,
					Change:    &change,
				},
				SemanticTokensProvider: &SemanticTokensOptions{
					Legend: SemanticTokensLegend{
						TokenTypes:     tokenTypes,
						TokenModifiers: tokenModifiers,
					},
				},
			},
			ServerInfo: ServerInfo{
				Name:    "OLS",
				Version: "0.0",
			},
		},
	}
}
