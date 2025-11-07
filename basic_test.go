package find_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/find"
)

func ExampleFind_basic() {
	// find . -name "*.go" -maxdepth 2
	// Note: This would actually search the filesystem
	yup.MustRun(
		Find(".", Name("*.go"), MaxDepth(2)),
	)
	// Output:
	//
}

