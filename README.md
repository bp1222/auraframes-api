# Aura Frames OpenAPI [unofficial]
This project attempts to map the [Aura Frames](https://auraframes.com) REST API in use by the iPhone application

This API Currently supports simplistic data access.  Accessing your frames, assets and the recursive objects below them.

## Installation
### Golang
    go get github.com/bp1222/auraframes-api/go

## Language Documentation
### [Golang](go/README.md)

## Development
To build the API you need to install `swagger-cli` and `openapi-generator`

### MacOS
    npm install -g swagger-cli
    brew install openapi-generatorur


## Notice
The Aura Frames REST API does not allow for direct uploading of images.  That seems to happen over a non-http protocol.

However, each [frame](go/docs/Frame.md) has a `email_address` field which is an email you can send images to upload.  Below is basic code to send an image (in the form of a URL) to the frame

```golang
import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"os"
	"path"
	"path/filepath"
	"strings"

    auraframes "github.com/bp1222/auraframes-api/go"
)

var (
	host       = os.Getenv("EMAIL_HOST")
	username   = os.Getenv("EMAiL_USERNAME")
	password   = os.Getenv("EMAIL_PASSWORD")
	portNumber = os.Getenv("EMAIL_PORT")
)

func SendAuraFrameEmail(frame auraframes.Frame, imageUrl string) {
	sender := New()
	m := NewMessage("Image Upload", imageUrl)
	m.To = []string{frame.GetEmailAddress()}
	m.AttachUrlImage(image)
	fmt.Println(sender.Send(m))
}

type Sender struct {
	auth smtp.Auth
}

type Message struct {
	To          []string
	Subject     string
	Body        string
	Attachments map[string][]byte
}

func New() *Sender {
	auth := smtp.PlainAuth("", username, password, host)
	return &Sender{auth}
}

func (s *Sender) Send(m *Message) error {
	return smtp.SendMail(fmt.Sprintf("%s:%s", host, portNumber), s.auth, username, m.To, m.ToBytes())
}

func NewMessage(s, b string) *Message {
	return &Message{
		Subject:     s,
		Body:        b,
		Attachments: make(map[string][]byte),
	}
}

func (m *Message) AttachUrlImage(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	imageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	filename := path.Base(resp.Request.URL.Path)

	m.Attachments[filename] = imageBytes

	return nil
}

func (m *Message) AttachFile(src string) error {
	b, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	_, fileName := filepath.Split(src)
	m.Attachments[fileName] = b
	return nil
}

func (m *Message) ToBytes() []byte {
	buf := bytes.NewBuffer(nil)
	withAttachments := len(m.Attachments) > 0
	buf.WriteString(fmt.Sprintf("Subject: %s\n", m.Subject))
	buf.WriteString(fmt.Sprintf("To: %s\n", strings.Join(m.To, ",")))

	buf.WriteString("MIME-Version: 1.0\n")

	mixedWriter := multipart.NewWriter(buf)
	altWriter := multipart.NewWriter(buf)
	mixedBoundary := mixedWriter.Boundary()
	altBoundary := altWriter.Boundary()

	if withAttachments {
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\n", mixedBoundary))
		buf.WriteString(fmt.Sprintf("--%s\n", mixedBoundary))
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/alternative; boundary=\"%s\"\n", altBoundary))
		buf.WriteString(fmt.Sprintf("--%s\n", altBoundary))
	}

	buf.WriteString("Content-Type: text/plain; charset=utf-8\n")
	buf.WriteString(m.Body)
	buf.WriteString(fmt.Sprintf("\n\n--%s--\n", altBoundary))
	if withAttachments {
		for k, v := range m.Attachments {
			buf.WriteString(fmt.Sprintf("\n\n--%s\n", mixedBoundary))
			buf.WriteString(fmt.Sprintf("Content-Type: %s; name=\"%s\"\n", http.DetectContentType(v), k))
			buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\n", k))
			buf.WriteString("Content-Transfer-Encoding: base64\n\n")

			b := make([]byte, base64.StdEncoding.EncodedLen(len(v)))
			base64.StdEncoding.Encode(b, v)
			buf.Write(b)
			buf.WriteString(fmt.Sprintf("\n--%s", mixedBoundary))
		}

		buf.WriteString("--")
	}

	return buf.Bytes()
}
```

