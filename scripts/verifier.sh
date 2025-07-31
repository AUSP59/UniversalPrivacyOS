#!/bin/bash
echo "🔍 Verifying UniversalPrivacyOS..."
gpg --verify signatures/upos.sig upos.bin
sha256sum upos.bin > local_hash.txt
echo "✅ Verification complete."
