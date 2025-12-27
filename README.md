# MicroCenter Stock Alert (Go + Raspberry Pi)

A lightweight Go program that checks the stock status of a specific Micro Center product and sends an email notification when the item becomes available. Designed to run on a Raspberry Pi for low-power, always-on monitoring.

## Features

- Check the stock status of one Micro Center product
- Detects: IN STOCK, LIMITED AVAILABILITY, SOLD OUT, NO LONGER CARRIED
- Sends an email alert when the item becomes available
- Simple scheduling using cron
- Runs perfectly on Raspberry Pi or any Linux system
- No external Go dependencies (only net/http and net/smtp)

## How It Works

1. The program fetches a Micro Center product page.
2. It scans the HTML for stock status keywords.
3. If the product is IN STOCK or LIMITED AVAILABILITY, an email alert is sent.
4. A cron job runs the program automatically at a chosen interval.

## Project Structure

microcenteralert/
- main.go
- go.mod
- README.md

To build the binary, run:
go build -o microcenteralert

## Environment Variables

Set the required SMTP environment variables:

export SMTP_HOST="smtp.gmail.com"
export SMTP_PORT="587"
export SMTP_USER="you@gmail.com"
export SMTP_PASS="your_app_password"
export ALERT_EMAIL="destination@example.com"

For Gmail, generate an app password and use it for SMTP_PASS.

To store these permanently, add them to ~/.bashrc and reload:
source ~/.bashrc

## Running Manually

go run .

Or using the binary:

./microcenteralert

Example output:

Ryzen 5000 cpu – Local Store -> SOLD OUT
Not available, no email sent.

## Cron Job Setup (Raspberry Pi)

Edit the cron table:

crontab -e

Add this line to run every 30 minutes:

*/30 * * * * /home/john/microcenteralert/microcenteralert >> /home/john/microcenteralert.log 2>&1

Check the log:

cat /home/john/microcenteralert.log

## Customizing the Product URL

Update the product URL in main.go:

productURL := "https://www.microcenter.com/product/...?...&storeid=123"

Make sure your URL includes your Micro Center store ID.

## Requirements

- Go 1.18+
- Raspberry Pi OS or any Linux distro
- SMTP-capable email account (Gmail recommended)
- Micro Center product URL

## Why This Project Exists

Micro Center does not provide stock alerts for many products, especially rare or low-stock items. This program:

- Runs in the background
- Automatically notifies you on restocks
- Uses almost no system resources
- Serves as a practical, real-world Go + Raspberry Pi project

## Future Enhancements

- Config file (JSON or YAML)
- Monitor multiple items
- Add Discord alerts
- Better HTML parsing
- Log rotation or timestamped output
- Docker container support

## License

MIT License — see LICENSE for details.


