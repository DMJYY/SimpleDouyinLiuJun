package logic

import (
	"context"
	"doushen_by_liujun/internal/common"
	"doushen_by_liujun/internal/util"
	"doushen_by_liujun/service/chat/api/internal/svc"
	"doushen_by_liujun/service/chat/api/internal/types"
	"doushen_by_liujun/service/chat/rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type MessageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageListLogic {
	return &MessageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageListLogic) MessageList(req *types.MessageChatReq) (*types.MessageChatReqResp, error) {
	l.Logger.Info("MessageList方法请求参数：", req)
	var lastTime int64
	if req.PreMsgTime > 169268692200 {
		// 获取第三位数字
		thirdDigit := (req.PreMsgTime / 100) % 10
		// 进行四舍五入
		if thirdDigit >= 5 {
			lastTime = req.PreMsgTime/1000 + 1
		} else {
			lastTime = req.PreMsgTime / 1000
		}
	} else {
		lastTime = req.PreMsgTime
	}

	// parse token
	res, err := util.ParseToken(req.Token)
	if err != nil {
		l.Logger.Error(err)
		return &types.MessageChatReqResp{
			StatusCode:  common.TOKEN_EXPIRE_ERROR,
			StatusMsg:   common.MapErrMsg(common.TOKEN_EXPIRE_ERROR),
			MessageList: nil,
		}, nil
	}

	// get params
	userId := res.UserID
	toUserId := req.ToUserId

	request := pb.GetChatMessageByIdReq{
		UserId:     userId,
		ToUserId:   toUserId,
		PreMsgTime: lastTime,
	}
	// get chat messages
	message, err := l.svcCtx.ChatRpcClient.GetChatMessageById(l.ctx, &request)
	if err != nil {
		l.Logger.Error(err)
		return &types.MessageChatReqResp{
			StatusCode:  common.DB_ERROR,
			StatusMsg:   common.MapErrMsg(common.DB_ERROR),
			MessageList: nil,
		}, nil
	}

	var messages []types.Message
	for _, item := range message.MessageList {
		msg := types.Message{
			Id:         item.Id,
			ToUserId:   item.ToUserId,
			FromUserId: item.FromUserId,
			Content:    item.Content,
			CreateTime: *item.CreateTime,
		}
		messages = append(messages, msg)
	}

	return &types.MessageChatReqResp{
		StatusCode:  common.OK,
		StatusMsg:   common.MapErrMsg(common.OK),
		MessageList: messages,
	}, nil
}
