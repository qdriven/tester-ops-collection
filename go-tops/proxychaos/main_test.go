package proxychaos

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/elazarl/goproxy"
)

func Test_chaosHandler_Handle(t *testing.T) {
	type fields struct {
		normalRate  float64
		blockRate   float64
		latencyRate float64
		latencySec  int64
	}
	type args struct {
		r *http.Request
		c *goproxy.ProxyCtx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantReq *http.Request
		wantRes *http.Response
	}{
		{
			name:    "should return request if normal rate is 1.0",
			fields:  fields{normalRate: 1.0, blockRate: 0.0, latencyRate: 0.0, latencySec: 0},
			args:    args{&http.Request{}, &goproxy.ProxyCtx{}},
			wantReq: &http.Request{},
			wantRes: nil,
		},
		{
			name:    "should return nil if block rate is 1.0",
			fields:  fields{normalRate: 0.0, blockRate: 1.0, latencyRate: 0.0, latencySec: 0},
			args:    args{&http.Request{}, &goproxy.ProxyCtx{}},
			wantReq: nil,
			wantRes: nil,
		},
		{
			name:    "should return request but delayed if latency rate is 1.0",
			fields:  fields{normalRate: 0.0, blockRate: 0.0, latencyRate: 1.0, latencySec: 1},
			args:    args{&http.Request{}, &goproxy.ProxyCtx{}},
			wantReq: &http.Request{},
			wantRes: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewChaosHandler(tt.fields.normalRate, tt.fields.blockRate, tt.fields.latencyRate, tt.fields.latencySec)
			req, res := h.Handle(tt.args.r, tt.args.c)
			if !reflect.DeepEqual(req, tt.wantReq) {
				t.Errorf("chaosHandler.Handle() request = %v, want %v", req, tt.wantReq)
			}
			if !reflect.DeepEqual(res, tt.wantRes) {
				t.Errorf("chaosHandler.Handle() response = %v, want %v", res, tt.wantRes)
			}
		})
	}
}
