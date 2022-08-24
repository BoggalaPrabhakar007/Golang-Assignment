package service

import (
	"context"
	"github.com/BoggalaPrabhakar007/golang-assignment/config"
	repomock "github.com/BoggalaPrabhakar007/golang-assignment/pkg/mocks"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/models"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func TestPortServ_GetPortsData(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []models.PortDetails
		wantErr bool
		err     error
	}{
		{
			name: "GetPortsData",
			args: args{
				ctx: context.Background(),
			},
			want:    []models.PortDetails{},
			wantErr: false,
			err:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rMock := repomock.PortRepoService{}
			rMock.On("GetPorts", mock.Anything).Return(tt.want, tt.err)
			p := PortServ{}
			config.LoadConfig()
			got, err := p.GetPortsData(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPortsData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPortsData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPortServ_GetPortDataByID(t *testing.T) {
	config.LoadConfig()
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    models.PortDetails
		wantErr bool
		err     error
	}{
		{
			name: "GetPortDataByID",
			args: args{
				ctx: context.Background(),
				id:  "id",
			},
			want:    models.PortDetails{},
			wantErr: false,
			err:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rMock := repomock.PortRepoService{}
			rMock.On("GetPortByID", mock.Anything, mock.Anything).Return(tt.want, tt.err)
			p := PortServ{}
			got, err := p.GetPortDataByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPortDataByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPortDataByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPortServ_DeletePortByID(t *testing.T) {
	config.LoadConfig()
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "DeletePortByID",
			args: args{
				ctx: context.Background(),
				id:  "id",
			},
			wantErr: false,
			err:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rMock := repomock.PortRepoService{}
			rMock.On("DeletePortByID", mock.Anything, mock.Anything).Return(tt.err)
			p := PortServ{}
			if err := p.DeletePortByID(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeletePortByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPortServ_UpdatePortByID(t *testing.T) {
	config.LoadConfig()
	type args struct {
		ctx  context.Context
		port models.PortDetails
		id   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "DeletePortByID",
			args: args{
				ctx:  context.Background(),
				port: models.PortDetails{},
				id:   "id",
			},
			wantErr: false,
			err:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rMock := repomock.PortRepoService{}
			rMock.On("UpdatePortByID", mock.Anything, mock.Anything, mock.Anything).Return(tt.err)
			p := PortServ{}
			if err := p.UpdatePortByID(tt.args.ctx, tt.args.port, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePortByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
