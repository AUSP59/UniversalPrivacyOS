#!/bin/bash
echo "🔧 Installing UniversalPrivacyOS..."
go version || { echo 'Go not installed'; exit 1; }
mkdir -p ~/UniversalPrivacyOS/logs
cp -r * ~/UniversalPrivacyOS/
echo "✅ Installed at ~/UniversalPrivacyOS/"
