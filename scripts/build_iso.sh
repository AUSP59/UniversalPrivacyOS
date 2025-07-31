#!/bin/bash
echo "🌀 Building bootable ISO..."
mkdir -p iso_root
cp -r * iso_root/
xorriso -as mkisofs -o isos/UniversalPrivacyOS.iso -b isolinux.bin -c boot.cat -no-emul-boot iso_root
echo "✅ ISO created at isos/UniversalPrivacyOS.iso"
