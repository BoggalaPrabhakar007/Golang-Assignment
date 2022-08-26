package endpoint

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	servicemock "github.com/BoggalaPrabhakar007/golang-assignment/pkg/mocks"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/models"

	"github.com/stretchr/testify/mock"
)

func TestInsertPortDataEndPoint(t *testing.T) {
	w := httptest.NewRecorder()
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "InsertPortDataEndPoint",
			args: args{
				w: w,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sMock := servicemock.PortService{}
			ePortServ := NewEndpoint(&sMock)
			sMock.On("InsertPortData", mock.Anything, mock.Anything).Return(tt.err)
			ePortServ.InsertPortDataEndPoint(tt.args.w, tt.args.r)
		})
	}
}

func TestGetPortsDataEndPoint(t *testing.T) {
	w := httptest.NewRecorder()
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want []models.PortDetails
		err  error
	}{
		{
			name: "GetPortsDataEndPoint",
			args: args{w: w},
			want: []models.PortDetails{},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sMock := servicemock.PortService{}
			ePortServ := NewEndpoint(&sMock)
			sMock.On("GetPortsData", mock.Anything).Return(tt.want, tt.err)
			ePortServ.GetPortsDataEndPoint(tt.args.w, tt.args.r)
		})
	}
}

func TestGetPortDataByIDEndPoint(t *testing.T) {
	httpVals := url.Values{}
	httpVals.Add("id", "id")
	req, _ := http.NewRequest("POST", " url", strings.NewReader(httpVals.Encode()))
	w := httptest.NewRecorder()
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want models.PortDetails
		err  error
	}{
		{
			name: "GetPortDataByIDEndPoint",
			args: args{w: w, r: req},
			want: models.PortDetails{},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sMock := servicemock.PortService{}
			ePortServ := NewEndpoint(&sMock)
			sMock.On("GetPortDataByID", mock.Anything, mock.Anything).Return(tt.want, tt.err)
			ePortServ.GetPortDataByIDEndPoint(tt.args.w, tt.args.r)
		})
	}
}

func TestDeletePortByIDEndPoint(t *testing.T) {
	httpVals := url.Values{}
	httpVals.Add("id", "id")
	req, _ := http.NewRequest("POST", " url", strings.NewReader(httpVals.Encode()))
	w := httptest.NewRecorder()
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "DeletePortByIDEndPoint",
			args: args{w: w, r: req},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sMock := servicemock.PortService{}
			ePortServ := NewEndpoint(&sMock)
			sMock.On("DeletePortByID", mock.Anything, mock.Anything).Return(tt.err)
			ePortServ.DeletePortByIDEndPoint(tt.args.w, tt.args.r)
		})
	}
}

func TestUpdatePortByIDEndPoint(t *testing.T) {
	httpVals := url.Values{}
	httpVals.Add("id", "id")
	req, _ := http.NewRequest("POST", " url", strings.NewReader(httpVals.Encode()))
	updateportBody, _ := json.Marshal(models.PortDetails{})
	req.Body = ioutil.NopCloser(strings.NewReader(string(updateportBody)))
	w := httptest.NewRecorder()
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "UpdatePortByIDEndPoint",
			args: args{w: w, r: req},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sMock := servicemock.PortService{}
			ePortServ := NewEndpoint(&sMock)
			sMock.On("UpdatePortByID", mock.Anything, mock.Anything, mock.Anything).Return(tt.err)
			ePortServ.UpdatePortByIDEndPoint(tt.args.w, tt.args.r)
		})
	}
}
