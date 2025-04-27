package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "s", 2)
	task(ctx)
}

func task(ctx context.Context) {
	ctx = context.WithValue(ctx, "d", 3)
	fmt.Println(ctx.Value("s"))
}
