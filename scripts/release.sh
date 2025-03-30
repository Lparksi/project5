#!/bin/bash

# Prompt for version
read -p "Enter the version number: " version

# Generate changelog
echo "Generating changelog..."
git log --pretty=format:"- %s" $(git describe --tags --abbrev=0)..HEAD > changelog.txt

# Install goreleaser
echo "Installing goreleaser..."
go install github.com/goreleaser/goreleaser@latest

# Run goreleaser
echo "Running goreleaser..."
goreleaser release --snapshot --skip-publish --rm-dist

# List compiled binaries with checksums
echo "Listing compiled binaries with checksums..."
find dist -type f -name "*.tar.gz" -o -name "*.zip" | while read -r file; do
    echo -e "### $file\n\`\`\`\n$(sha256sum "$file" | awk '{print $1}')\n\`\`\`\n" >> binaries.md
done

echo "Release script completed. Changelog saved to changelog.txt and binaries list with checksums saved to binaries.md."
