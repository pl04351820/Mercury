# Mercury
A step function implementation for connecting different cloud platform.

# Goal
The goal of this project is to build a scalable and fault-tolerant workflow system to connect serverless function on different cloud provider or your own platform.

# Philosophy
Instead of choosing to design workflow for specific platform, we build a universal solution based on same protocol. For now, different serverless provider usually share same API design Philosophy, Mercury only do works by using "invoke" interface.
State machine language is used to describe function workflow. Mercury chooses a standard that is compatible for most serverless providers.

# Protocol
Invoke interface should provide events and log interface.
State machine language in first stage will choose Amazon State Language as template, and will add more complex logic in future stage. You can find more details for this spec in
[Amazon State Language](https://states-language.net/spec.html)

# Usage
```
```

# Installation
```
```

# Import Hits
Events, all states share events during the whole state machine process.