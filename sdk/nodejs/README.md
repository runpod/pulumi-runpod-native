# @runpod/pulumi-runpod

The Runpod Pulumi provider provides resources to interact with Runpod's native APIs.

## Installation

```bash
npm install @runpod/pulumi-runpod
```

or

```bash
yarn add @runpod/pulumi-runpod
```

## Usage

```typescript
import * as runpod from "@runpod/pulumi-runpod";

// Create a new pod template
const myTemplate = new runpod.Template("myTemplate", {
    containerDiskInGb: 5,
    dockerArgs: "python handler.py",
    imageName: "runpod/base:0.0.0",
    name: "my-template",
});
```

## Documentation

For detailed documentation, please visit [Runpod Documentation](https://docs.runpod.io/overview).
