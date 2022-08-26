package repo

import (
	"context"
	"reflect"
	"testing"

	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/models"
	"github.com/BoggalaPrabhakar007/golang-assignment/repository-lib/pkg/mongodb"
	repolibmock "github.com/BoggalaPrabhakar007/golang-assignment/repository-lib/pkg/mongodb/mocks"

	"github.com/stretchr/testify/mock"
)

func TestPortRepoServ_InsertPorts(t *testing.T) {
	type fields struct {
		repoLib mongodb.RepoLib
	}
	type args struct {
		ctx          context.Context
		portsDetails []models.PortDetails
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		err     error
		eInter  interface{}
	}{
		{
			name: "Insert ports success",
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rMock := repolibmock.RepoLib{}
			rMock.On("InsertMultipleRecords", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.eInter, tt.err)
			p := NewPortRepoServ(&rMock)
			if err := p.InsertPorts(tt.args.ctx, tt.args.portsDetails); (err != nil) != tt.wantErr {
				t.Errorf("InsertPorts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPortRepoServ_GetPorts(t *testing.T) {
	type fields struct {
		repoLib mongodb.RepoLib
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.PortDetails
		wantErr bool
		err     error
	}{
		{
			name: "get ports success",
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rMock := repolibmock.RepoLib{}
			rMock.On("GetRecords", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.err)
			p := NewPortRepoServ(&rMock)
			got, err := p.GetPorts(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPorts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPorts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPortRepoServ_GetPortByID(t *testing.T) {
	type fields struct {
		repoLib mongodb.RepoLib
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.PortDetails
		wantErr bool
		err     error
	}{
		{
			name: "Get port by id success",
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rMock := repolibmock.RepoLib{}
			rMock.On("GetRecord", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.err)
			p := NewPortRepoServ(&rMock)
			got, err := p.GetPortByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPortByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPortByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPortRepoServ_DeletePortByID(t *testing.T) {
	type fields struct {
		repoLib mongodb.RepoLib
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Delete port by id success",
			err:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rMock := repolibmock.RepoLib{}
			rMock.On("DeleteRecordByID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.err)
			p := NewPortRepoServ(&rMock)
			if err := p.DeletePortByID(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeletePortByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPortRepoServ_UpdatePortByID(t *testing.T) {
	type fields struct {
		repoLib mongodb.RepoLib
	}
	type args struct {
		ctx  context.Context
		id   string
		port *models.PortDetails
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		err      error
		updated  int
		modified int
	}{
		{
			name: "Update port by id success",
			err:  nil,
			args: args{
				port: &models.PortDetails{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rMock := repolibmock.RepoLib{}
			rMock.On("UpdateRecord", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.updated, tt.modified, tt.err)
			p := NewPortRepoServ(&rMock)
			if err := p.UpdatePortByID(tt.args.ctx, tt.args.id, tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePortByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
