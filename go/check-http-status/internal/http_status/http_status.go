package http_status

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/supermarine1377/monitoring-scripts/go/check-http-status/timeutil"
)

type Monitorer struct {
	httpClient      *http.Client
	targetURL       string
	intervalSeconds int
	files           []io.Writer
}

func NewMonitorer(targetURL string, intervalSeconds int, files []io.Writer) *Monitorer {
	return &Monitorer{
		httpClient:      http.DefaultClient,
		targetURL:       targetURL,
		intervalSeconds: intervalSeconds,
		files:           files,
	}
}

func (m *Monitorer) Do(ctx context.Context) {
Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		default:
			r, err := m.result(ctx)
			if err != nil {
				m.logln(err.Error())
				continue
			}
			m.logln(r)
			time.Sleep(time.Second * time.Duration(m.intervalSeconds))
		}
	}
}

func (m *Monitorer) result(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		m.targetURL, nil,
	)
	if err != nil {
		return "", err
	}
	t := timeutil.NowStr()
	res, err := m.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	s := t + res.Status
	return s, nil
}

func (m *Monitorer) logln(s string) {
	b := []byte(s + "\n")
	for _, f := range m.files {
		f.Write(b)
	}
}
