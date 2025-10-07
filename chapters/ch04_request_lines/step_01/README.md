# Task

- Create two files `main.go` and `test.go`.
- Write the code below in `test.go`.

```go
package main

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestRequestLineParse(t *testing.T) {
    assert.Equal(t, "Hey, Chat!", "Hey, Chat!")
}
```

- Ensure that the `main.go` file contains a `main` function that calls `TestRequestLineParse`.

```go
package main

import "testing"

func main() {
    TestRequestLineParse(&testing.T{})
}
```

- The `test.go` file should only contain the necessary imports and the `package` declaration.
- Ensure that both files are syntactically correct and can be compiled without errors.

**This task is just to ensure that `testify` is working**
