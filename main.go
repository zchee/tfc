package main

import (
	_ "github.com/google/go-github/v45/github"
	_ "github.com/google/go-querystring/query"
	_ "github.com/hashicorp/go-retryablehttp"
	_ "github.com/hashicorp/go-tfe"
	_ "github.com/hashicorp/hcl/v2"
	_ "github.com/spf13/cobra"
	_ "github.com/spf13/pflag"
	_ "golang.org/x/time/rate"
)

func main() {}
