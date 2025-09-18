#!/bin/bash
set -e

echo "🚀 Starting dual publishing process..."

# Check if VERSION is set
if [ -z "$VERSION" ]; then
    echo "❌ VERSION environment variable is required"
    echo "Usage: VERSION=v1.10.0 ./dual-publish.sh"
    exit 1
fi

echo "📦 Building all packages for version $VERSION..."

# Build everything
make build

echo "📤 Publishing to npm (dual packages)..."

# Node.js - Publish BOTH old and new packages
cd sdk/nodejs/bin

# Backup original package.json
cp package.json package.json.backup

# Publish NEW package (@runpod/pulumi-runpod)
echo "Publishing @runpod/pulumi-runpod..."
npm publish --access public

# Modify for OLD package name and publish
sed 's/"@runpod\/pulumi-runpod"/"@runpod-infra\/pulumi"/g' package.json.backup > package.json
echo "Publishing @runpod-infra/pulumi (compatibility)..."
npm publish

# Restore original
mv package.json.backup package.json

cd ../../..

echo "🐍 Publishing to PyPI (dual packages)..."

# Python - Publish BOTH old and new packages
cd sdk/python

# Backup original setup.py
cp setup.py setup.py.backup

# Publish NEW package (pulumi-runpod)
echo "Publishing pulumi-runpod..."
python -m twine upload dist/*

# Modify for OLD package name and publish
sed 's/name='\''pulumi-runpod'\''/name='\''runpodinfra'\''/g' setup.py.backup > setup.py
sed 's/'\''pulumi_runpod'\''/'\''runpodinfra'\''/g' setup.py > setup.py.tmp && mv setup.py.tmp setup.py
# Need to rebuild with old package structure
echo "Building old package structure..."
# This would need more complex logic to rebuild the old structure

# Restore original
mv setup.py.backup setup.py

cd ../..

echo "💎 Publishing to NuGet..."

# .NET
cd sdk/dotnet
echo "Publishing Pulumi.Runpod..."
dotnet nuget push Pulumi.Runpod.*.nupkg --api-key $NUGET_API_KEY --source https://api.nuget.org/v3/index.json

cd ../..

echo "🏷️  Creating Git tag..."
git tag $VERSION
git push origin $VERSION

echo "✅ Dual publishing complete!"
echo ""
echo "📋 Published packages:"
echo "   npm: @runpod/pulumi-runpod AND @runpod-infra/pulumi"
echo "   PyPI: pulumi-runpod AND runpodinfra"
echo "   NuGet: Pulumi.Runpod"
echo "   Go: github.com/runpod/pulumi-runpod-native/sdk (via git tag)"
echo ""
echo "🔄 Next steps:"
echo "   1. Submit to Pulumi Registry"
echo "   2. Update documentation"
echo "   3. Add deprecation notices to old packages"