// Package valuechain holds the study material as files so it stays reviewable and
// diffable. The server embeds it; nothing about the content lives in the database.
package valuechain

import "embed"

// Files is the content tree: lecture notes, question banks, cases and the glossary.
//
//go:embed all:content
var Files embed.FS
