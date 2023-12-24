// Code generated from ./grammar/fhirpath.g4 by ANTLR 4.13.1. DO NOT EDIT.

package fhirpath

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type fhirpathLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var FhirpathLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func fhirpathlexerLexerInit() {
	staticData := &FhirpathLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'.'", "'['", "']'", "'+'", "'-'", "'*'", "'/'", "'div'", "'mod'",
		"'&'", "'is'", "'as'", "'|'", "'<='", "'<'", "'>'", "'>='", "'='", "'~'",
		"'!='", "'!~'", "'in'", "'contains'", "'and'", "'or'", "'xor'", "'implies'",
		"'('", "')'", "'{'", "'}'", "'true'", "'false'", "'%'", "'$this'", "'$index'",
		"'$total'", "','", "'year'", "'month'", "'week'", "'day'", "'hour'",
		"'minute'", "'second'", "'millisecond'", "'years'", "'months'", "'weeks'",
		"'days'", "'hours'", "'minutes'", "'seconds'", "'milliseconds'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "DATETIME", "TIME", "IDENTIFIER", "DELIMITEDIDENTIFIER",
		"STRING", "NUMBER", "WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
		"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "T__48",
		"T__49", "T__50", "T__51", "T__52", "T__53", "DATETIME", "TIME", "TIMEFORMAT",
		"IDENTIFIER", "DELIMITEDIDENTIFIER", "STRING", "NUMBER", "WS", "COMMENT",
		"LINE_COMMENT", "ESC", "UNICODE", "HEX",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 63, 514, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43,
		1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1,
		44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45,
		1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48,
		1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1,
		50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51,
		1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1,
		53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53,
		1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1,
		54, 1, 54, 1, 54, 3, 54, 385, 8, 54, 3, 54, 387, 8, 54, 1, 54, 3, 54, 390,
		8, 54, 1, 54, 3, 54, 393, 8, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1,
		54, 1, 54, 3, 54, 402, 8, 54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56,
		1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 4, 56, 418, 8,
		56, 11, 56, 12, 56, 419, 3, 56, 422, 8, 56, 3, 56, 424, 8, 56, 3, 56, 426,
		8, 56, 1, 57, 3, 57, 429, 8, 57, 1, 57, 5, 57, 432, 8, 57, 10, 57, 12,
		57, 435, 9, 57, 1, 58, 1, 58, 1, 58, 5, 58, 440, 8, 58, 10, 58, 12, 58,
		443, 9, 58, 1, 58, 1, 58, 1, 59, 1, 59, 1, 59, 5, 59, 450, 8, 59, 10, 59,
		12, 59, 453, 9, 59, 1, 59, 1, 59, 1, 60, 4, 60, 458, 8, 60, 11, 60, 12,
		60, 459, 1, 60, 1, 60, 4, 60, 464, 8, 60, 11, 60, 12, 60, 465, 3, 60, 468,
		8, 60, 1, 61, 4, 61, 471, 8, 61, 11, 61, 12, 61, 472, 1, 61, 1, 61, 1,
		62, 1, 62, 1, 62, 1, 62, 5, 62, 481, 8, 62, 10, 62, 12, 62, 484, 9, 62,
		1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 63, 1, 63, 1, 63, 1, 63, 5, 63, 495,
		8, 63, 10, 63, 12, 63, 498, 9, 63, 1, 63, 1, 63, 1, 64, 1, 64, 1, 64, 3,
		64, 505, 8, 64, 1, 65, 1, 65, 1, 65, 1, 65, 1, 65, 1, 65, 1, 66, 1, 66,
		3, 441, 451, 482, 0, 67, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44,
		89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105,
		53, 107, 54, 109, 55, 111, 56, 113, 0, 115, 57, 117, 58, 119, 59, 121,
		60, 123, 61, 125, 62, 127, 63, 129, 0, 131, 0, 133, 0, 1, 0, 8, 1, 0, 48,
		57, 2, 0, 43, 43, 45, 45, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13,
		13, 8, 0, 39, 39, 47, 47, 92, 92, 96, 96, 102, 102, 110, 110, 114, 114,
		116, 116, 3, 0, 48, 57, 65, 70, 97, 102, 533, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0,
		27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0,
		0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0,
		0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0,
		0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1,
		0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65,
		1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0,
		73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0,
		0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0,
		0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0,
		0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103,
		1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0,
		0, 111, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1,
		0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0,
		127, 1, 0, 0, 0, 1, 135, 1, 0, 0, 0, 3, 137, 1, 0, 0, 0, 5, 139, 1, 0,
		0, 0, 7, 141, 1, 0, 0, 0, 9, 143, 1, 0, 0, 0, 11, 145, 1, 0, 0, 0, 13,
		147, 1, 0, 0, 0, 15, 149, 1, 0, 0, 0, 17, 153, 1, 0, 0, 0, 19, 157, 1,
		0, 0, 0, 21, 159, 1, 0, 0, 0, 23, 162, 1, 0, 0, 0, 25, 165, 1, 0, 0, 0,
		27, 167, 1, 0, 0, 0, 29, 170, 1, 0, 0, 0, 31, 172, 1, 0, 0, 0, 33, 174,
		1, 0, 0, 0, 35, 177, 1, 0, 0, 0, 37, 179, 1, 0, 0, 0, 39, 181, 1, 0, 0,
		0, 41, 184, 1, 0, 0, 0, 43, 187, 1, 0, 0, 0, 45, 190, 1, 0, 0, 0, 47, 199,
		1, 0, 0, 0, 49, 203, 1, 0, 0, 0, 51, 206, 1, 0, 0, 0, 53, 210, 1, 0, 0,
		0, 55, 218, 1, 0, 0, 0, 57, 220, 1, 0, 0, 0, 59, 222, 1, 0, 0, 0, 61, 224,
		1, 0, 0, 0, 63, 226, 1, 0, 0, 0, 65, 231, 1, 0, 0, 0, 67, 237, 1, 0, 0,
		0, 69, 239, 1, 0, 0, 0, 71, 245, 1, 0, 0, 0, 73, 252, 1, 0, 0, 0, 75, 259,
		1, 0, 0, 0, 77, 261, 1, 0, 0, 0, 79, 266, 1, 0, 0, 0, 81, 272, 1, 0, 0,
		0, 83, 277, 1, 0, 0, 0, 85, 281, 1, 0, 0, 0, 87, 286, 1, 0, 0, 0, 89, 293,
		1, 0, 0, 0, 91, 300, 1, 0, 0, 0, 93, 312, 1, 0, 0, 0, 95, 318, 1, 0, 0,
		0, 97, 325, 1, 0, 0, 0, 99, 331, 1, 0, 0, 0, 101, 336, 1, 0, 0, 0, 103,
		342, 1, 0, 0, 0, 105, 350, 1, 0, 0, 0, 107, 358, 1, 0, 0, 0, 109, 371,
		1, 0, 0, 0, 111, 403, 1, 0, 0, 0, 113, 407, 1, 0, 0, 0, 115, 428, 1, 0,
		0, 0, 117, 436, 1, 0, 0, 0, 119, 446, 1, 0, 0, 0, 121, 457, 1, 0, 0, 0,
		123, 470, 1, 0, 0, 0, 125, 476, 1, 0, 0, 0, 127, 490, 1, 0, 0, 0, 129,
		501, 1, 0, 0, 0, 131, 506, 1, 0, 0, 0, 133, 512, 1, 0, 0, 0, 135, 136,
		5, 46, 0, 0, 136, 2, 1, 0, 0, 0, 137, 138, 5, 91, 0, 0, 138, 4, 1, 0, 0,
		0, 139, 140, 5, 93, 0, 0, 140, 6, 1, 0, 0, 0, 141, 142, 5, 43, 0, 0, 142,
		8, 1, 0, 0, 0, 143, 144, 5, 45, 0, 0, 144, 10, 1, 0, 0, 0, 145, 146, 5,
		42, 0, 0, 146, 12, 1, 0, 0, 0, 147, 148, 5, 47, 0, 0, 148, 14, 1, 0, 0,
		0, 149, 150, 5, 100, 0, 0, 150, 151, 5, 105, 0, 0, 151, 152, 5, 118, 0,
		0, 152, 16, 1, 0, 0, 0, 153, 154, 5, 109, 0, 0, 154, 155, 5, 111, 0, 0,
		155, 156, 5, 100, 0, 0, 156, 18, 1, 0, 0, 0, 157, 158, 5, 38, 0, 0, 158,
		20, 1, 0, 0, 0, 159, 160, 5, 105, 0, 0, 160, 161, 5, 115, 0, 0, 161, 22,
		1, 0, 0, 0, 162, 163, 5, 97, 0, 0, 163, 164, 5, 115, 0, 0, 164, 24, 1,
		0, 0, 0, 165, 166, 5, 124, 0, 0, 166, 26, 1, 0, 0, 0, 167, 168, 5, 60,
		0, 0, 168, 169, 5, 61, 0, 0, 169, 28, 1, 0, 0, 0, 170, 171, 5, 60, 0, 0,
		171, 30, 1, 0, 0, 0, 172, 173, 5, 62, 0, 0, 173, 32, 1, 0, 0, 0, 174, 175,
		5, 62, 0, 0, 175, 176, 5, 61, 0, 0, 176, 34, 1, 0, 0, 0, 177, 178, 5, 61,
		0, 0, 178, 36, 1, 0, 0, 0, 179, 180, 5, 126, 0, 0, 180, 38, 1, 0, 0, 0,
		181, 182, 5, 33, 0, 0, 182, 183, 5, 61, 0, 0, 183, 40, 1, 0, 0, 0, 184,
		185, 5, 33, 0, 0, 185, 186, 5, 126, 0, 0, 186, 42, 1, 0, 0, 0, 187, 188,
		5, 105, 0, 0, 188, 189, 5, 110, 0, 0, 189, 44, 1, 0, 0, 0, 190, 191, 5,
		99, 0, 0, 191, 192, 5, 111, 0, 0, 192, 193, 5, 110, 0, 0, 193, 194, 5,
		116, 0, 0, 194, 195, 5, 97, 0, 0, 195, 196, 5, 105, 0, 0, 196, 197, 5,
		110, 0, 0, 197, 198, 5, 115, 0, 0, 198, 46, 1, 0, 0, 0, 199, 200, 5, 97,
		0, 0, 200, 201, 5, 110, 0, 0, 201, 202, 5, 100, 0, 0, 202, 48, 1, 0, 0,
		0, 203, 204, 5, 111, 0, 0, 204, 205, 5, 114, 0, 0, 205, 50, 1, 0, 0, 0,
		206, 207, 5, 120, 0, 0, 207, 208, 5, 111, 0, 0, 208, 209, 5, 114, 0, 0,
		209, 52, 1, 0, 0, 0, 210, 211, 5, 105, 0, 0, 211, 212, 5, 109, 0, 0, 212,
		213, 5, 112, 0, 0, 213, 214, 5, 108, 0, 0, 214, 215, 5, 105, 0, 0, 215,
		216, 5, 101, 0, 0, 216, 217, 5, 115, 0, 0, 217, 54, 1, 0, 0, 0, 218, 219,
		5, 40, 0, 0, 219, 56, 1, 0, 0, 0, 220, 221, 5, 41, 0, 0, 221, 58, 1, 0,
		0, 0, 222, 223, 5, 123, 0, 0, 223, 60, 1, 0, 0, 0, 224, 225, 5, 125, 0,
		0, 225, 62, 1, 0, 0, 0, 226, 227, 5, 116, 0, 0, 227, 228, 5, 114, 0, 0,
		228, 229, 5, 117, 0, 0, 229, 230, 5, 101, 0, 0, 230, 64, 1, 0, 0, 0, 231,
		232, 5, 102, 0, 0, 232, 233, 5, 97, 0, 0, 233, 234, 5, 108, 0, 0, 234,
		235, 5, 115, 0, 0, 235, 236, 5, 101, 0, 0, 236, 66, 1, 0, 0, 0, 237, 238,
		5, 37, 0, 0, 238, 68, 1, 0, 0, 0, 239, 240, 5, 36, 0, 0, 240, 241, 5, 116,
		0, 0, 241, 242, 5, 104, 0, 0, 242, 243, 5, 105, 0, 0, 243, 244, 5, 115,
		0, 0, 244, 70, 1, 0, 0, 0, 245, 246, 5, 36, 0, 0, 246, 247, 5, 105, 0,
		0, 247, 248, 5, 110, 0, 0, 248, 249, 5, 100, 0, 0, 249, 250, 5, 101, 0,
		0, 250, 251, 5, 120, 0, 0, 251, 72, 1, 0, 0, 0, 252, 253, 5, 36, 0, 0,
		253, 254, 5, 116, 0, 0, 254, 255, 5, 111, 0, 0, 255, 256, 5, 116, 0, 0,
		256, 257, 5, 97, 0, 0, 257, 258, 5, 108, 0, 0, 258, 74, 1, 0, 0, 0, 259,
		260, 5, 44, 0, 0, 260, 76, 1, 0, 0, 0, 261, 262, 5, 121, 0, 0, 262, 263,
		5, 101, 0, 0, 263, 264, 5, 97, 0, 0, 264, 265, 5, 114, 0, 0, 265, 78, 1,
		0, 0, 0, 266, 267, 5, 109, 0, 0, 267, 268, 5, 111, 0, 0, 268, 269, 5, 110,
		0, 0, 269, 270, 5, 116, 0, 0, 270, 271, 5, 104, 0, 0, 271, 80, 1, 0, 0,
		0, 272, 273, 5, 119, 0, 0, 273, 274, 5, 101, 0, 0, 274, 275, 5, 101, 0,
		0, 275, 276, 5, 107, 0, 0, 276, 82, 1, 0, 0, 0, 277, 278, 5, 100, 0, 0,
		278, 279, 5, 97, 0, 0, 279, 280, 5, 121, 0, 0, 280, 84, 1, 0, 0, 0, 281,
		282, 5, 104, 0, 0, 282, 283, 5, 111, 0, 0, 283, 284, 5, 117, 0, 0, 284,
		285, 5, 114, 0, 0, 285, 86, 1, 0, 0, 0, 286, 287, 5, 109, 0, 0, 287, 288,
		5, 105, 0, 0, 288, 289, 5, 110, 0, 0, 289, 290, 5, 117, 0, 0, 290, 291,
		5, 116, 0, 0, 291, 292, 5, 101, 0, 0, 292, 88, 1, 0, 0, 0, 293, 294, 5,
		115, 0, 0, 294, 295, 5, 101, 0, 0, 295, 296, 5, 99, 0, 0, 296, 297, 5,
		111, 0, 0, 297, 298, 5, 110, 0, 0, 298, 299, 5, 100, 0, 0, 299, 90, 1,
		0, 0, 0, 300, 301, 5, 109, 0, 0, 301, 302, 5, 105, 0, 0, 302, 303, 5, 108,
		0, 0, 303, 304, 5, 108, 0, 0, 304, 305, 5, 105, 0, 0, 305, 306, 5, 115,
		0, 0, 306, 307, 5, 101, 0, 0, 307, 308, 5, 99, 0, 0, 308, 309, 5, 111,
		0, 0, 309, 310, 5, 110, 0, 0, 310, 311, 5, 100, 0, 0, 311, 92, 1, 0, 0,
		0, 312, 313, 5, 121, 0, 0, 313, 314, 5, 101, 0, 0, 314, 315, 5, 97, 0,
		0, 315, 316, 5, 114, 0, 0, 316, 317, 5, 115, 0, 0, 317, 94, 1, 0, 0, 0,
		318, 319, 5, 109, 0, 0, 319, 320, 5, 111, 0, 0, 320, 321, 5, 110, 0, 0,
		321, 322, 5, 116, 0, 0, 322, 323, 5, 104, 0, 0, 323, 324, 5, 115, 0, 0,
		324, 96, 1, 0, 0, 0, 325, 326, 5, 119, 0, 0, 326, 327, 5, 101, 0, 0, 327,
		328, 5, 101, 0, 0, 328, 329, 5, 107, 0, 0, 329, 330, 5, 115, 0, 0, 330,
		98, 1, 0, 0, 0, 331, 332, 5, 100, 0, 0, 332, 333, 5, 97, 0, 0, 333, 334,
		5, 121, 0, 0, 334, 335, 5, 115, 0, 0, 335, 100, 1, 0, 0, 0, 336, 337, 5,
		104, 0, 0, 337, 338, 5, 111, 0, 0, 338, 339, 5, 117, 0, 0, 339, 340, 5,
		114, 0, 0, 340, 341, 5, 115, 0, 0, 341, 102, 1, 0, 0, 0, 342, 343, 5, 109,
		0, 0, 343, 344, 5, 105, 0, 0, 344, 345, 5, 110, 0, 0, 345, 346, 5, 117,
		0, 0, 346, 347, 5, 116, 0, 0, 347, 348, 5, 101, 0, 0, 348, 349, 5, 115,
		0, 0, 349, 104, 1, 0, 0, 0, 350, 351, 5, 115, 0, 0, 351, 352, 5, 101, 0,
		0, 352, 353, 5, 99, 0, 0, 353, 354, 5, 111, 0, 0, 354, 355, 5, 110, 0,
		0, 355, 356, 5, 100, 0, 0, 356, 357, 5, 115, 0, 0, 357, 106, 1, 0, 0, 0,
		358, 359, 5, 109, 0, 0, 359, 360, 5, 105, 0, 0, 360, 361, 5, 108, 0, 0,
		361, 362, 5, 108, 0, 0, 362, 363, 5, 105, 0, 0, 363, 364, 5, 115, 0, 0,
		364, 365, 5, 101, 0, 0, 365, 366, 5, 99, 0, 0, 366, 367, 5, 111, 0, 0,
		367, 368, 5, 110, 0, 0, 368, 369, 5, 100, 0, 0, 369, 370, 5, 115, 0, 0,
		370, 108, 1, 0, 0, 0, 371, 372, 5, 64, 0, 0, 372, 373, 7, 0, 0, 0, 373,
		374, 7, 0, 0, 0, 374, 375, 7, 0, 0, 0, 375, 392, 7, 0, 0, 0, 376, 377,
		5, 45, 0, 0, 377, 378, 7, 0, 0, 0, 378, 389, 7, 0, 0, 0, 379, 380, 5, 45,
		0, 0, 380, 381, 7, 0, 0, 0, 381, 386, 7, 0, 0, 0, 382, 384, 5, 84, 0, 0,
		383, 385, 3, 113, 56, 0, 384, 383, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385,
		387, 1, 0, 0, 0, 386, 382, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 390,
		1, 0, 0, 0, 388, 390, 5, 84, 0, 0, 389, 379, 1, 0, 0, 0, 389, 388, 1, 0,
		0, 0, 389, 390, 1, 0, 0, 0, 390, 393, 1, 0, 0, 0, 391, 393, 5, 84, 0, 0,
		392, 376, 1, 0, 0, 0, 392, 391, 1, 0, 0, 0, 392, 393, 1, 0, 0, 0, 393,
		401, 1, 0, 0, 0, 394, 402, 5, 90, 0, 0, 395, 396, 7, 1, 0, 0, 396, 397,
		7, 0, 0, 0, 397, 398, 7, 0, 0, 0, 398, 399, 5, 58, 0, 0, 399, 400, 7, 0,
		0, 0, 400, 402, 7, 0, 0, 0, 401, 394, 1, 0, 0, 0, 401, 395, 1, 0, 0, 0,
		401, 402, 1, 0, 0, 0, 402, 110, 1, 0, 0, 0, 403, 404, 5, 64, 0, 0, 404,
		405, 5, 84, 0, 0, 405, 406, 3, 113, 56, 0, 406, 112, 1, 0, 0, 0, 407, 408,
		7, 0, 0, 0, 408, 425, 7, 0, 0, 0, 409, 410, 5, 58, 0, 0, 410, 411, 7, 0,
		0, 0, 411, 423, 7, 0, 0, 0, 412, 413, 5, 58, 0, 0, 413, 414, 7, 0, 0, 0,
		414, 421, 7, 0, 0, 0, 415, 417, 5, 46, 0, 0, 416, 418, 7, 0, 0, 0, 417,
		416, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 417, 1, 0, 0, 0, 419, 420,
		1, 0, 0, 0, 420, 422, 1, 0, 0, 0, 421, 415, 1, 0, 0, 0, 421, 422, 1, 0,
		0, 0, 422, 424, 1, 0, 0, 0, 423, 412, 1, 0, 0, 0, 423, 424, 1, 0, 0, 0,
		424, 426, 1, 0, 0, 0, 425, 409, 1, 0, 0, 0, 425, 426, 1, 0, 0, 0, 426,
		114, 1, 0, 0, 0, 427, 429, 7, 2, 0, 0, 428, 427, 1, 0, 0, 0, 429, 433,
		1, 0, 0, 0, 430, 432, 7, 3, 0, 0, 431, 430, 1, 0, 0, 0, 432, 435, 1, 0,
		0, 0, 433, 431, 1, 0, 0, 0, 433, 434, 1, 0, 0, 0, 434, 116, 1, 0, 0, 0,
		435, 433, 1, 0, 0, 0, 436, 441, 5, 96, 0, 0, 437, 440, 3, 129, 64, 0, 438,
		440, 9, 0, 0, 0, 439, 437, 1, 0, 0, 0, 439, 438, 1, 0, 0, 0, 440, 443,
		1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 441, 439, 1, 0, 0, 0, 442, 444, 1, 0,
		0, 0, 443, 441, 1, 0, 0, 0, 444, 445, 5, 96, 0, 0, 445, 118, 1, 0, 0, 0,
		446, 451, 5, 39, 0, 0, 447, 450, 3, 129, 64, 0, 448, 450, 9, 0, 0, 0, 449,
		447, 1, 0, 0, 0, 449, 448, 1, 0, 0, 0, 450, 453, 1, 0, 0, 0, 451, 452,
		1, 0, 0, 0, 451, 449, 1, 0, 0, 0, 452, 454, 1, 0, 0, 0, 453, 451, 1, 0,
		0, 0, 454, 455, 5, 39, 0, 0, 455, 120, 1, 0, 0, 0, 456, 458, 7, 0, 0, 0,
		457, 456, 1, 0, 0, 0, 458, 459, 1, 0, 0, 0, 459, 457, 1, 0, 0, 0, 459,
		460, 1, 0, 0, 0, 460, 467, 1, 0, 0, 0, 461, 463, 5, 46, 0, 0, 462, 464,
		7, 0, 0, 0, 463, 462, 1, 0, 0, 0, 464, 465, 1, 0, 0, 0, 465, 463, 1, 0,
		0, 0, 465, 466, 1, 0, 0, 0, 466, 468, 1, 0, 0, 0, 467, 461, 1, 0, 0, 0,
		467, 468, 1, 0, 0, 0, 468, 122, 1, 0, 0, 0, 469, 471, 7, 4, 0, 0, 470,
		469, 1, 0, 0, 0, 471, 472, 1, 0, 0, 0, 472, 470, 1, 0, 0, 0, 472, 473,
		1, 0, 0, 0, 473, 474, 1, 0, 0, 0, 474, 475, 6, 61, 0, 0, 475, 124, 1, 0,
		0, 0, 476, 477, 5, 47, 0, 0, 477, 478, 5, 42, 0, 0, 478, 482, 1, 0, 0,
		0, 479, 481, 9, 0, 0, 0, 480, 479, 1, 0, 0, 0, 481, 484, 1, 0, 0, 0, 482,
		483, 1, 0, 0, 0, 482, 480, 1, 0, 0, 0, 483, 485, 1, 0, 0, 0, 484, 482,
		1, 0, 0, 0, 485, 486, 5, 42, 0, 0, 486, 487, 5, 47, 0, 0, 487, 488, 1,
		0, 0, 0, 488, 489, 6, 62, 0, 0, 489, 126, 1, 0, 0, 0, 490, 491, 5, 47,
		0, 0, 491, 492, 5, 47, 0, 0, 492, 496, 1, 0, 0, 0, 493, 495, 8, 5, 0, 0,
		494, 493, 1, 0, 0, 0, 495, 498, 1, 0, 0, 0, 496, 494, 1, 0, 0, 0, 496,
		497, 1, 0, 0, 0, 497, 499, 1, 0, 0, 0, 498, 496, 1, 0, 0, 0, 499, 500,
		6, 63, 0, 0, 500, 128, 1, 0, 0, 0, 501, 504, 5, 92, 0, 0, 502, 505, 7,
		6, 0, 0, 503, 505, 3, 131, 65, 0, 504, 502, 1, 0, 0, 0, 504, 503, 1, 0,
		0, 0, 505, 130, 1, 0, 0, 0, 506, 507, 5, 117, 0, 0, 507, 508, 3, 133, 66,
		0, 508, 509, 3, 133, 66, 0, 509, 510, 3, 133, 66, 0, 510, 511, 3, 133,
		66, 0, 511, 132, 1, 0, 0, 0, 512, 513, 7, 7, 0, 0, 513, 134, 1, 0, 0, 0,
		24, 0, 384, 386, 389, 392, 401, 419, 421, 423, 425, 428, 431, 433, 439,
		441, 449, 451, 459, 465, 467, 472, 482, 496, 504, 1, 0, 1, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// fhirpathLexerInit initializes any static state used to implement fhirpathLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewfhirpathLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func FhirpathLexerInit() {
	staticData := &FhirpathLexerLexerStaticData
	staticData.once.Do(fhirpathlexerLexerInit)
}

// NewfhirpathLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewfhirpathLexer(input antlr.CharStream) *fhirpathLexer {
	FhirpathLexerInit()
	l := new(fhirpathLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &FhirpathLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "fhirpath.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// fhirpathLexer tokens.
const (
	fhirpathLexerT__0                = 1
	fhirpathLexerT__1                = 2
	fhirpathLexerT__2                = 3
	fhirpathLexerT__3                = 4
	fhirpathLexerT__4                = 5
	fhirpathLexerT__5                = 6
	fhirpathLexerT__6                = 7
	fhirpathLexerT__7                = 8
	fhirpathLexerT__8                = 9
	fhirpathLexerT__9                = 10
	fhirpathLexerT__10               = 11
	fhirpathLexerT__11               = 12
	fhirpathLexerT__12               = 13
	fhirpathLexerT__13               = 14
	fhirpathLexerT__14               = 15
	fhirpathLexerT__15               = 16
	fhirpathLexerT__16               = 17
	fhirpathLexerT__17               = 18
	fhirpathLexerT__18               = 19
	fhirpathLexerT__19               = 20
	fhirpathLexerT__20               = 21
	fhirpathLexerT__21               = 22
	fhirpathLexerT__22               = 23
	fhirpathLexerT__23               = 24
	fhirpathLexerT__24               = 25
	fhirpathLexerT__25               = 26
	fhirpathLexerT__26               = 27
	fhirpathLexerT__27               = 28
	fhirpathLexerT__28               = 29
	fhirpathLexerT__29               = 30
	fhirpathLexerT__30               = 31
	fhirpathLexerT__31               = 32
	fhirpathLexerT__32               = 33
	fhirpathLexerT__33               = 34
	fhirpathLexerT__34               = 35
	fhirpathLexerT__35               = 36
	fhirpathLexerT__36               = 37
	fhirpathLexerT__37               = 38
	fhirpathLexerT__38               = 39
	fhirpathLexerT__39               = 40
	fhirpathLexerT__40               = 41
	fhirpathLexerT__41               = 42
	fhirpathLexerT__42               = 43
	fhirpathLexerT__43               = 44
	fhirpathLexerT__44               = 45
	fhirpathLexerT__45               = 46
	fhirpathLexerT__46               = 47
	fhirpathLexerT__47               = 48
	fhirpathLexerT__48               = 49
	fhirpathLexerT__49               = 50
	fhirpathLexerT__50               = 51
	fhirpathLexerT__51               = 52
	fhirpathLexerT__52               = 53
	fhirpathLexerT__53               = 54
	fhirpathLexerDATETIME            = 55
	fhirpathLexerTIME                = 56
	fhirpathLexerIDENTIFIER          = 57
	fhirpathLexerDELIMITEDIDENTIFIER = 58
	fhirpathLexerSTRING              = 59
	fhirpathLexerNUMBER              = 60
	fhirpathLexerWS                  = 61
	fhirpathLexerCOMMENT             = 62
	fhirpathLexerLINE_COMMENT        = 63
)
