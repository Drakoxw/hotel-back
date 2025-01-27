package services

import (
	"bytes"
	"encoding/base64"
	"html/template"
	"io"
	"mime/multipart"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
)

// SendEmailWithTemplate envía un correo con datos personalizados y plantillas.
func SendEmailWithTemplate(recipient, subject, templatePath string, emailData interface{}, attachments []string) error {
	// Configuración del servidor SMTP
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	senderEmail := os.Getenv("SENDER_EMAIL")
	senderPassword := os.Getenv("SENDER_PASSWORD")

	// Cargar la plantilla HTML
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	// Ejecutar la plantilla con los datos
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Encabezados del correo
	buf.WriteString("To: " + recipient + "\r\n")
	buf.WriteString("Subject: " + subject + "\r\n")
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString("Content-Type: multipart/mixed; boundary=" + writer.Boundary() + "\r\n")
	buf.WriteString("\r\n")

	// Parte del cuerpo del mensaje (HTML)
	part, err := writer.CreatePart(map[string][]string{
		"Content-Type": {"text/html; charset=utf-8"},
	})
	if err != nil {
		return err
	}

	err = tmpl.Execute(part, emailData)
	if err != nil {
		return err
	}

	// Adjuntar archivos (si los hay)
	for _, attachment := range attachments {
		if err := addAttachment(writer, attachment); err != nil {
			return err
		}
	}

	// Cerrar el escritor de multipartes
	writer.Close()

	// Autenticación para el servidor SMTP
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)

	// Envío del correo electrónico
	err = smtp.SendMail(smtpServer+":"+smtpPort, auth, senderEmail, []string{recipient}, buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}

// Función para añadir un archivo adjunto al correo
func addAttachment(writer *multipart.Writer, filename string) error {
	// Abrir el archivo
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Crear la parte del adjunto
	part, err := writer.CreatePart(map[string][]string{
		"Content-Type":              {detectContentType(filename)},
		"Content-Transfer-Encoding": {"base64"},
		"Content-Disposition":       {"attachment; filename=\"" + filepath.Base(filename) + "\""},
	})
	if err != nil {
		return err
	}

	// Codificar el archivo en base64 y escribirlo en la parte del adjunto
	encoder := base64.NewEncoder(base64.StdEncoding, part)
	_, err = io.Copy(encoder, file)
	if err != nil {
		return err
	}
	encoder.Close()

	return nil
}

// Función para detectar el tipo de contenido del archivo
func detectContentType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".txt":
		return "text/plain"
	case ".html":
		return "text/html"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".pdf":
		return "application/pdf"
	default:
		return "application/octet-stream"
	}
}
