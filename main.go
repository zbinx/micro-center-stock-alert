package main

import (
    "fmt"
    "io"
    "net/http"
    "net/smtp"
    "os"
    "strings"
    "time"
)

func main() {
    productName := "SilverStone FLP02 – microcenter Mayfield Heights"
    productURL := "https://www.microcenter.com/product/702259/silverstone-flp02-retro-inspired-atx-mid-tower-computer-case-beige?storeid=051"

    status, err := checkStock(productURL)
    if err != nil {
        fmt.Printf("Error checking stock: %v\n", err)
        return
    }

    fmt.Printf("%s -> %s\n", productName, status)

    if isAvailable(status) {
        subject := "Micro Center stock alert"
        body := fmt.Sprintf("%s is %s!\n\nLink: %s\n", productName, status, productURL)

        if err := sendEmailAlert(subject, body); err != nil {
            fmt.Printf("Error sending email: %v\n", err)
        } else {
            fmt.Println("Email alert sent.")
        }
    } else {
        fmt.Println("Not available, no email sent.")
    }
}

func checkStock(url string) (string, error) {
    client := &http.Client{
        Timeout: 15 * time.Second,
    }

    resp, err := client.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return detectStock(string(body)), nil
}

func detectStock(html string) string {
    upper := strings.ToUpper(html)

    switch {
    case strings.Contains(upper, "NO LONGER CARRIED"):
        return "NO LONGER CARRIED"
    case strings.Contains(upper, "SOLD OUT"):
        return "SOLD OUT"
    case strings.Contains(upper, "IN STOCK"):
        return "IN STOCK"
    case strings.Contains(upper, "LIMITED AVAILABILITY"):
        return "LIMITED AVAILABILITY"
    default:
        return "UNKNOWN"
    }
}

func isAvailable(status string) bool {
    switch status {
    case "IN STOCK", "LIMITED AVAILABILITY":
        return true
    default:
        return false
    }
}

func sendEmailAlert(subject, body string) error {
    smtpHost := os.Getenv("SMTP_HOST")
    smtpPort := os.Getenv("SMTP_PORT")
    smtpUser := os.Getenv("SMTP_USER")
    smtpPass := os.Getenv("SMTP_PASS")
    alertEmail := os.Getenv("ALERT_EMAIL")

    if smtpHost == "" || smtpPort == "" || smtpUser == "" || smtpPass == "" || alertEmail == "" {
        return fmt.Errorf("SMTP env vars not set (SMTP_HOST, SMTP_PORT, SMTP_USER, SMTP_PASS, ALERT_EMAIL)")
    }

    from := smtpUser
    to := []string{alertEmail}

    // Simple plain-text email
    msg := []byte(
        "To: " + alertEmail + "\r\n" +
            "Subject: " + subject + "\r\n" +
            "From: " + from + "\r\n" +
            "\r\n" +
            body + "\r\n",
    )

    addr := smtpHost + ":" + smtpPort
    auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

    return smtp.SendMail(addr, auth, from, to, msg)
}
