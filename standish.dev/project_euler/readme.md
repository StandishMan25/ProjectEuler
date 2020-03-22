# Getting Started
* Clone repo
* Add file to `problems` directory
* Reference problem in main.go

# Format of file

```go
package problems

import (
    "any/necessary/modules"
    "standish.dev/project_euler/utility"
    "etc"
)

func ProblemNumber() {  
    // Logic goes here

    fmt.Println(answer)
}
```

# Running
Once you're ready to execute your code, navigate to `standish.dev/project_euler`, run `go build`, then `go install`. Execute `project_euler` to see your results!

If `project_euler` isn't in the path, you can add it. Naviate again to `standish.dev/project_euler` and execute `export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))`