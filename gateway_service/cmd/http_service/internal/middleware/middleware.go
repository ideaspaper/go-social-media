package middleware

import "golang.org/x/exp/slog"

type Middleware struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *Middleware {
	return &Middleware{
		logger: logger,
	}
}
