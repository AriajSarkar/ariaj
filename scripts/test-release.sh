#!/bin/bash
# For Windows, save as test-release.bat and modify commands accordingly

# Clean dist directory
rm -rf dist/

# Run goreleaser in test mode
goreleaser build --clean --snapshot

# Test the binary
cd dist/ariaj_linux_amd64_v1/
./ariaj --help | grep "Usage: ariaj \[prompt\]"

if [ $? -eq 0 ]; then
    echo "Local test passed!"
else
    echo "Local test failed!"
    exit 1
fi
