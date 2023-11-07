package base

import (
	"context"
	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
	"github.com/suyuan32/simple-admin-message-center/types/mcms"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitMcmsDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitMcmsDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitMcmsDatabaseLogic {
	return &InitMcmsDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitMcmsDatabaseLogic) InitMcmsDatabase() (resp *types.BaseMsgResp, err error) {
	if !l.svcCtx.Config.ProjectConf.AllowInit {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.PermissionDeny)
	}

	if !l.svcCtx.Config.McmsRpc.Enabled {
		return nil, errorx.NewCodeUnavailableError(i18n.ServiceUnavailable)
	}
	check, err := l.svcCtx.CoreRpc.GetApiList(l.ctx, &core.ApiListReq{
		Page:     1,
		PageSize: 10,
		ApiGroup: pointy.GetPointer("email"),
	})
	if err != nil {
		return nil, err
	}

	if check.Total != 0 {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.AlreadyInit)
	}

	// insert api and menu data
	err = l.InsertApiData()
	if err != nil {
		return nil, err
	}

	err = l.InsertMenuData()
	if err != nil {
		return nil, err
	}

	result, err := l.svcCtx.McmsRpc.InitDatabase(l.ctx, &mcms.Empty{})
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.Casbin.LoadPolicy()
	if err != nil {
		logx.Errorw("failed to load Casbin policy", logx.Field("detail", err))
		return nil, errorx.NewCodeInternalError(i18n.DatabaseError)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
}

func (l *InitMcmsDatabaseLogic) InsertApiData() error {
	// Email Log
	_, err := l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/email_log/create"),
		Description: pointy.GetPointer("apiDesc.createEmailLog"),
		ApiGroup:    pointy.GetPointer("email_log"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/email_log/update"),
		Description: pointy.GetPointer("apiDesc.updateEmailLog"),
		ApiGroup:    pointy.GetPointer("email_log"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/email_log/delete"),
		Description: pointy.GetPointer("apiDesc.deleteEmailLog"),
		ApiGroup:    pointy.GetPointer("email_log"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/email_log/list"),
		Description: pointy.GetPointer("apiDesc.getEmailLogList"),
		ApiGroup:    pointy.GetPointer("email_log"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/email_log"),
		Description: pointy.GetPointer("apiDesc.getEmailLogById"),
		ApiGroup:    pointy.GetPointer("email_log"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	// Email Provider
	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/email_provider/create"),
		Description: pointy.GetPointer("apiDesc.createEmailProvider"),
		ApiGroup:    pointy.GetPointer("email_provider"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/email_provider/update"),
		Description: pointy.GetPointer("apiDesc.updateEmailProvider"),
		ApiGroup:    pointy.GetPointer("email_provider"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/email_provider/delete"),
		Description: pointy.GetPointer("apiDesc.deleteEmailProvider"),
		ApiGroup:    pointy.GetPointer("email_provider"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/email_provider/list"),
		Description: pointy.GetPointer("apiDesc.getEmailProviderList"),
		ApiGroup:    pointy.GetPointer("email_provider"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/email_provider"),
		Description: pointy.GetPointer("apiDesc.getEmailProviderById"),
		ApiGroup:    pointy.GetPointer("email_provider"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	// Sms Log
	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/sms_log/create"),
		Description: pointy.GetPointer("apiDesc.createSmsLog"),
		ApiGroup:    pointy.GetPointer("sms_log"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/sms_log/update"),
		Description: pointy.GetPointer("apiDesc.updateSmsLog"),
		ApiGroup:    pointy.GetPointer("sms_log"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/sms_log/delete"),
		Description: pointy.GetPointer("apiDesc.deleteSmsLog"),
		ApiGroup:    pointy.GetPointer("sms_log"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/sms_log/list"),
		Description: pointy.GetPointer("apiDesc.getSmsLogList"),
		ApiGroup:    pointy.GetPointer("sms_log"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/sms_log"),
		Description: pointy.GetPointer("apiDesc.getSmsLogById"),
		ApiGroup:    pointy.GetPointer("sms_log"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	// Sms Provider
	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/sms_provider/create"),
		Description: pointy.GetPointer("apiDesc.createSmsProvider"),
		ApiGroup:    pointy.GetPointer("sms_provider"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/sms_provider/update"),
		Description: pointy.GetPointer("apiDesc.updateSmsProvider"),
		ApiGroup:    pointy.GetPointer("sms_provider"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/sms_provider/delete"),
		Description: pointy.GetPointer("apiDesc.deleteSmsProvider"),
		ApiGroup:    pointy.GetPointer("sms_provider"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/sms_provider/list"),
		Description: pointy.GetPointer("apiDesc.getSmsProviderList"),
		ApiGroup:    pointy.GetPointer("sms_provider"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/sms_provider"),
		Description: pointy.GetPointer("apiDesc.getSmsProviderById"),
		ApiGroup:    pointy.GetPointer("sms_provider"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/sms/send"),
		Description: pointy.GetPointer("apiDesc.sendSms"),
		ApiGroup:    pointy.GetPointer("message_sender"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        pointy.GetPointer("/email/send"),
		Description: pointy.GetPointer("apiDesc.sendEmail"),
		ApiGroup:    pointy.GetPointer("message_sender"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	return nil
}

func (l *InitMcmsDatabaseLogic) InsertMenuData() error {
	menuData, err := l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Level:     pointy.GetPointer(uint32(1)),
		ParentId:  pointy.GetPointer(common.DefaultParentId),
		Path:      pointy.GetPointer("/mcms_dir"),
		Name:      pointy.GetPointer("MessageCenterManagement"),
		Component: pointy.GetPointer("LAYOUT"),
		Sort:      pointy.GetPointer(uint32(4)),
		Meta: &core.Meta{
			Title: pointy.GetPointer("route.messageCenterManagement"),
			Icon:  pointy.GetPointer("clarity:email-line"),
		},
		MenuType: pointy.GetPointer(uint32(1)),
	})
	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Level:     pointy.GetPointer(uint32(2)),
		ParentId:  pointy.GetPointer(menuData.Id),
		Path:      pointy.GetPointer("/mcms_email_provider"),
		Name:      pointy.GetPointer("EmailProviderManagement"),
		Component: pointy.GetPointer("/mcms/emailProvider/index"),
		Sort:      pointy.GetPointer(uint32(1)),
		Meta: &core.Meta{
			Title: pointy.GetPointer("route.emailProviderManagement"),
			Icon:  pointy.GetPointer("clarity:email-line"),
		},
		MenuType: pointy.GetPointer(uint32(2)),
	})
	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Level:     pointy.GetPointer(uint32(2)),
		ParentId:  pointy.GetPointer(menuData.Id),
		Path:      pointy.GetPointer("/mcms_sms_provider"),
		Name:      pointy.GetPointer("SmsProviderManagement"),
		Component: pointy.GetPointer("/mcms/smsProvider/index"),
		Sort:      pointy.GetPointer(uint32(2)),
		Meta: &core.Meta{
			Title: pointy.GetPointer("route.smsProviderManagement"),
			Icon:  pointy.GetPointer("clarity:mobile-line"),
		},
		MenuType: pointy.GetPointer(uint32(2)),
	})
	if err != nil {
		return err
	}

	return err
}
