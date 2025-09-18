#!/bin/bash

# Script to introspect Runpod GraphQL schema
# Usage: RUNPOD_API_KEY=your_key ./introspect_schema.sh

if [ -z "$RUNPOD_API_KEY" ]; then
    echo "Error: RUNPOD_API_KEY environment variable is required"
    echo "Usage: RUNPOD_API_KEY=your_key ./introspect_schema.sh"
    exit 1
fi

echo "Introspecting Runpod GraphQL schema..."
/Users/zackmckenna/go/bin/gqlintrospect "https://api.runpod.io/graphql?api_key=$RUNPOD_API_KEY" > graphql/schema.graphqls

if [ $? -eq 0 ]; then
    echo "Schema saved to graphql/schema.graphqls"
    echo "File size: $(wc -l < graphql/schema.graphqls) lines"
else
    echo "Error: Failed to introspect schema"
    exit 1
fi