package grpc

import (
	"context"
	"github.com/chroma/chroma-coordinator/internal/common"
	"github.com/chroma/chroma-coordinator/internal/grpcutils"
	"github.com/chroma/chroma-coordinator/internal/proto/logservicepb"
	"github.com/chroma/chroma-coordinator/internal/types"
	"github.com/pingcap/log"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func (s *Server) PushLogs(ctx context.Context, req *logservicepb.PushLogsRequest) (*logservicepb.PushLogsResponse, error) {
	res := &logservicepb.PushLogsResponse{}
	collectionID, err := types.ToUniqueID(&req.CollectionId)
	if err != nil {
		log.Error("collection id format error", zap.String("collection.id", req.CollectionId))
		return nil, common.ErrCollectionIDFormat
	}
	var recordsContent [][]byte
	for _, record := range req.Records {
		data, err := proto.Marshal(record)
		if err != nil {
			log.Error("marshaling error", zap.Error(err))
			grpcError, err := grpcutils.BuildInvalidArgumentGrpcError("records", "marshaling error")
			if err != nil {
				return nil, grpcError
			}
			return nil, err
		}
		recordsContent = append(recordsContent, data)
	}
	recordCount, err := s.logService.PushLogs(ctx, collectionID, recordsContent)
	if err != nil {
		log.Error("error pushing logs", zap.Error(err))
		return nil, grpcutils.BuildInternalGrpcError("error pushing logs")
	}
	res.RecordCount = int32(recordCount)
	log.Info("PushLogs success", zap.String("collectionID", req.CollectionId), zap.Int("recordCount", recordCount))
	return res, nil
}