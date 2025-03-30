#!/bin/bash

# 如果没有提供版本号，使用默认版本
version=${VERSION:-"0.1.0"}
version_tag="v${version}"

# 切换到 gin-frontend 目录
cd gin-frontend || exit 1

# 确保有 git 标签
if [ -z "$(git tag -l)" ]; then
    echo "Creating initial tag ${version_tag}..."
    git tag ${version_tag}
fi

# 生成 changelog
echo "Generating changelog..."
git log --pretty=format:"- %s" $(git describe --tags --abbrev=0)..HEAD > changelog.txt

# 运行 goreleaser
echo "Running goreleaser..."
goreleaser release --clean

echo "Release process completed. Changelog saved to changelog.txt"
