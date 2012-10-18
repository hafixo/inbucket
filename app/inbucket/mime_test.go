package inbucket

import (
	"bufio"
	"fmt"
	"github.com/stretchrcom/testify/assert"
	"net/mail"
	"os"
	"path/filepath"
	"testing"
)

func TestIdentifyNonMime(t *testing.T) {
	msg := readMessage("non-mime.raw")
	assert.False(t, IsMIMEMessage(msg), "Failed to identify non-MIME message")
}

func TestIdentifyMime(t *testing.T) {
	msg := readMessage("html-mime-inline.raw")
	assert.True(t, IsMIMEMessage(msg), "Failed to identify MIME message")
}

func TestParseNonMime(t *testing.T) {
	msg := readMessage("non-mime.raw")

	_, err := ParseMIMEMessage(msg)
	assert.NotNil(t, err, "Expected error parsing a non-MIME message")
}

func TestParseInlineText(t *testing.T) {
	msg := readMessage("html-mime-inline.raw")

	mime, err := ParseMIMEMessage(msg)
	if err != nil {
		t.Fatalf("Failed to parse mime: %v", err)
	}

	assert.Equal(t, mime.Text, "Test of HTML section")
}

func TestParseInlineHtml(t *testing.T) {
	msg := readMessage("html-mime-inline.raw")

	mime, err := ParseMIMEMessage(msg)
	if err != nil {
		t.Fatalf("Failed to parse mime: %v", err)
	}

	assert.Contains(t, mime.Html, "<html>")
	assert.Contains(t, mime.Html, "Test of HTML section")
}

// readMessage is a test utility function to fetch a mail.Message object.
func readMessage(filename string) *mail.Message {
	// Open test email for parsing
	raw, err := os.Open(filepath.Join("..", "..", "test-data", filename))
	if err != nil {
		panic(fmt.Sprintf("Failed to open test data: %v", err))
	}

	// Parse email into a mail.Message object like we do
	reader := bufio.NewReader(raw)
	msg, err := mail.ReadMessage(reader)
	if err != nil {
		panic(fmt.Sprintf("Failed to read message: %v", err))
	}

	return msg
}
