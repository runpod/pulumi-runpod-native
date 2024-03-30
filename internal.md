
# Internal docs
This is an internal doc only. If you are an end user, please refer to the examples in the Readme.md or inside the examples folder. This is strictly internal.

## Steps in publishing:

### Step one
There are a bunch of steps you have to take when pushing to the repository. Number one, you have to always build before pushing. 

```
    make build
```

When you run the command above, it first builds the provider, then auto-generates code, then generates SDKs, etc. The same thing happens when you cut a release and in order to
not cut a dirty release, you must ensure that you have all the files that will be generated when the workflow is run. Since Pulumi's codegen system is still in its infancy, there
might be files that are not present when you actually release it. One of them is the setup.py inside the Python SDK. If you are building dirty because of that file, please 
consider removing these lines in that file:

```
    import os
    VERSION = os.getenv("PULUMI_PYTHON_VERSION", "0.0.0")
```

### Step two
Step two is publishing your changes. The current workflow has all the steps that you need. Only update it if it is strictly necessary. Besides changing versions
of actions, there isn't much you will need to do.

Commit your changes, push changes to origin and then tag the root release in the format: "v0.0.0". Always make sure that the tags are higher. It might seem obvious but many slip at this
and never see their changes published.

```
    git tag v0.9.0
```

Then tag the Go SDK:
```
    git tag sdk/v0.9.0
```

The push the sdk first:
```
    git push origin sdk/v0.9.0
```

Then:
```
    git push origin v0.9.0
```

If you are not seeing your tags be released, go inside provider/provider.go and then manually change the version that is in there. The workflow is automatically setup to deploy your changes.
So you can rest assured.

## Common problems

With Python, you might come up against a situation where your pip's path might not be what you want. Instead of directly downloading with pip as usual, try this:

```
    python3 -m pip install YOUR_PACKAGE
```

## Point of contact
Pierre worked on the most for this repo. So please reach out to him over Slack.