package endpoint

import (
	"net/http"
	"testing"

	servicemock "github.com/BoggalaPrabhakar007/golang-assignment/pkg/mocks"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/models"

	"github.com/stretchr/testify/mock"
)

func TestInsertPortDataEndPoint(t *testing.T) {
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
			args: args{},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sMock := servicemock.PortService{}
			sMock.On("InsertPortData", mock.Anything, mock.Anything).Return(tt.err)
			InsertPortDataEndPoint(tt.args.w, tt.args.r)
		})
	}
}

func TestGetPortsDataEndPoint(t *testing.T) {
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
			args: args{},
			want: []models.PortDetails{},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sMock := servicemock.PortService{}
			sMock.On("GetPortsData", mock.Anything).Return(tt.want, tt.err)
			GetPortsDataEndPoint(tt.args.w, tt.args.r)
		})
	}
}

func TestGetPortDataByIDEndPoint(t *testing.T) {
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
			args: args{},
			want: models.PortDetails{},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sMock := servicemock.PortService{}
			sMock.On("GetPortDataByID", mock.Anything, mock.Anything).Return(tt.want, tt.err)
			GetPortDataByIDEndPoint(tt.args.w, tt.args.r)
		})
	}
}

func TestDeletePortByIDEndPoint(t *testing.T) {
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
			args: args{},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sMock := servicemock.PortService{}
			sMock.On("DeletePortByID", mock.Anything, mock.Anything).Return(tt.err)
			DeletePortByIDEndPoint(tt.args.w, tt.args.r)
		})
	}
}

func TestUpdatePortByIDEndPoint(t *testing.T) {
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
			args: args{},
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sMock := servicemock.PortService{}
			sMock.On("DeletePortByID", mock.Anything, mock.Anything).Return(tt.err)
			UpdatePortByIDEndPoint(tt.args.w, tt.args.r)
		})
	}
}
