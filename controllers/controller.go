package controllers

import "context"

type Controllers struct {
	ctx context.Context
}

func NewControllers(ctx context.Context) Controllers {
	return Controllers{ctx}
}
