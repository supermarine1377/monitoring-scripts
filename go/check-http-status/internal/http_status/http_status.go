package http_status

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Monitor struct {
	httpClient      *http.Client
	targetURL       string
	intervalSeconds int
}

func NewMonitor(targetURL string, intervalSeconds int) *Monitor {
	return &Monitor{
		httpClient:      http.DefaultClient,
		targetURL:       targetURL,
		intervalSeconds: intervalSeconds,
	}
}

// Writedown write down monitoring results on the given io.Writer
func (m *Monitor) Writedown(ctx context.Context, w io.Writer) {
Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		default:
			code, err := m.statusCode(ctx)
			if err != nil {
				errLog := []byte(err.Error() + "\n")
				w.Write(errLog)
				continue
			}
			log := strconv.Itoa(code) + "\n"
			w.Write([]byte(log))

			time.Sleep(time.Second * time.Duration(m.intervalSeconds))
		}
	}
}

func (m *Monitor) statusCode(ctx context.Context) (int, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		m.targetURL, nil,
	)
	if err != nil {
		return 0, err
	}
	res, err := m.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	return res.StatusCode, nil
}
