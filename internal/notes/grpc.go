package notes

import (
	"context"
	"database/sql"
	"errors"

	"github.com/IEP/sqlite-wasm/gen/go/database/notes"
	pb "github.com/IEP/sqlite-wasm/gen/go/protos/notes/v1"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCService struct {
	db   *sql.DB
	repo *notes.Queries

	pb.UnimplementedNoteServiceServer
}

func NewNotesGRPCService(db *sql.DB) *GRPCService {
	return &GRPCService{
		db:   db,
		repo: notes.New(db),
	}
}

func (c *GRPCService) GetNote(ctx context.Context, req *pb.GetNoteRequest) (*pb.Note, error) {
	note, err := c.repo.GetNote(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	var content *string = nil
	if note.Content.Valid {
		content = &note.Content.String
	}

	return &pb.Note{
		Id:      note.ID,
		Name:    note.Name,
		Content: content,
	}, nil
}

func (c *GRPCService) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.Note, error) {
	if req.GetName() == "" {
		return nil, errors.New("field 'name' should not be null")
	}

	note, err := c.repo.CreateNote(ctx, notes.CreateNoteParams{
		Name: req.GetName(),
		Content: sql.NullString{
			String: req.GetContent(),
			Valid:  true,
		},
	})
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	var content *string = nil
	if note.Content.Valid {
		content = &note.Content.String
	}

	return &pb.Note{
		Id:      note.ID,
		Name:    note.Name,
		Content: content,
	}, nil
}

func (c *GRPCService) ListNotes(ctx context.Context, req *pb.ListNotesRequest) (*pb.ListNotesResponse, error) {
	offset := DecodePageToken(req.PageToken)
	if offset < 0 {
		offset = 0
	}
	pageSize := req.PageSize
	if pageSize < 10 {
		pageSize = 10
	}
	filter := ""
	if req.GetFilter() != "" {
		filter = "%" + req.GetFilter() + "%"
	}

	ns, err := c.repo.ListNotes(ctx, notes.ListNotesParams{
		Filter: filter,
		Limit:  int64(pageSize),
		Offset: int64(offset),
	})
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	var maxID int64 = 0

	pbNotes := lo.Map(ns, func(n notes.Note, _ int) *pb.Note {
		var content *string = nil
		if n.Content.Valid {
			content = &n.Content.String
		}

		maxID = max(maxID, n.ID)

		return &pb.Note{
			Id:      n.ID,
			Name:    n.Name,
			Content: content,
		}
	})

	return &pb.ListNotesResponse{
		Notes:         pbNotes,
		NextPageToken: EncodePageToken(int(maxID)),
	}, nil
}

func (c *GRPCService) UpdateNote(ctx context.Context, req *pb.UpdateNoteRequest) (*pb.Note, error) {
	note := req.GetNote()
	if note.GetName() == "" {
		return nil, errors.New("field 'name' should not be null")
	}

	n, err := c.repo.UpdateNote(ctx, notes.UpdateNoteParams{
		Name: note.GetName(),
		Content: sql.NullString{
			String: note.GetContent(),
			Valid:  true,
		},
		ID: note.GetId(),
	})
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	var content *string = nil
	if n.Content.Valid {
		content = &n.Content.String
	}

	return &pb.Note{
		Id:      n.ID,
		Name:    n.Name,
		Content: content,
	}, nil
}

func (c *GRPCService) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest) (*emptypb.Empty, error) {
	err := c.repo.DeleteNote(ctx, int64(req.GetId()))
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}
	return &emptypb.Empty{}, nil
}
