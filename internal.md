
# Internal docs
This is an internal doc only. If you are an end user, please refer to the examples in the Readme.md or inside the examples folder. This is strictly internal.

## Steps in publishing:

### Step one
There are a bunch of steps you have to take when pushing to the repository. Number one, you have to always build before pushing. 

```
    git checkout main
    VERSION=v1.9.9 make build-and-push
```

Replace v1.9.9 with a semver.

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

## Common problems

With Python, you might come up against a situation where your pip's path might not be what you want. Instead of directly downloading with pip as usual, try this:

```
    python3 -m pip install YOUR_PACKAGE
```

## Point of contact
Pierre worked on the most for this repo. So please reach out to him over Slack.