package usecase

import (
	"fmt"
	"slices"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/constant"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RoleUseCase struct {
	RoleRepo IRepository.Role
}

func NewRoleUseCase(roleRepo IRepository.Role) IUseCase.Role {
	return &RoleUseCase{
		RoleRepo: roleRepo,
	}
}

func (c *RoleUseCase) CreateRole(ctx *gin.Context, name string) (IResponse.Role, error, int) {
	_, err := c.RoleRepo.GetRoleByName(ctx, name)

	if err == nil {
		global.Logger.Error("role is exist", zap.String("Status", "Error"))
		return IResponse.Role{}, fmt.Errorf("role is exist"), 409
	}

	role, err := global.DB.CreateRole(ctx, name)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Role{}, err, 401
	}

	return IResponse.Role{
		Id:         role.ID,
		Name:       role.Name,
		Permission: role.Permission,
	}, nil, 201
}

func (c *RoleUseCase) GetRoleUseCase(ctx *gin.Context, id int) (IResponse.Role, error, int) {
	Role, err := c.RoleRepo.GetRoleById(ctx, int64(id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Role{}, fmt.Errorf("get Role is not exist"), 401
	}

	return IResponse.Role{
		Id:         Role.ID,
		Name:       Role.Name,
		Permission: Role.Permission,
	}, nil, 200
}

func (c *RoleUseCase) GetAllRoleUseCase(ctx *gin.Context, req IRequest.GetAllRole) (IResponse.GetAllRole, error, int) {
	role, err := c.RoleRepo.GetAllRole(ctx, req)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.GetAllRole{}, fmt.Errorf("get Role is not exist"), 401
	}

	if len(role) == 0 {
		return IResponse.GetAllRole{
			Roles:      role,
			TotalCount: 0,
			TotalPage:  0,
		}, nil, 200
	}

	totalPage := utils.PageCount(int64(req.Limit), role[0].TotalCount)

	return IResponse.GetAllRole{
		Roles:      role,
		TotalCount: role[0].TotalCount,
		TotalPage:  totalPage,
	}, nil, 200
}

func (c *RoleUseCase) UpdateRoleUseCase(ctx *gin.Context, id int, name string, permission []string) (IResponse.Role, error, int) {
	idInt64 := int64(id)

	role, err := global.DB.GetRoleById(ctx, idInt64)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Role{}, fmt.Errorf("role is not exist"), 401
	}

	if slices.Contains(role.Permission, constant.CONFIG_PERMISSIONS["ADMIN"].(string)) || slices.Contains(role.Permission, constant.CONFIG_PERMISSIONS["BASIC"].(string)) || name == "Admin" || name == "Basic" {
		global.Logger.Error("role admin or basic don't remove!!", zap.String("Status", "Error"))
		return IResponse.Role{}, fmt.Errorf("ArrayID is empty"), 401
	}
	if len(permission) == 0 {
		updateRole, err := c.RoleRepo.UpdateRole(ctx, idInt64, name, role.Permission)

		if err != nil {
			global.Logger.Error(err.Error(), zap.String("Status", "Error"))
			return IResponse.Role{}, fmt.Errorf("update Role is fail"), 401
		}

		res := IResponse.Role{
			Id:         updateRole.ID,
			Name:       updateRole.Name,
			Permission: updateRole.Permission,
		}

		return res, nil, 200
	}

	if name == "" {
		updateRole, err := c.RoleRepo.UpdateRole(ctx, idInt64, role.Name, permission)

		if err != nil {
			global.Logger.Error(err.Error(), zap.String("Status", "Error"))
			return IResponse.Role{}, fmt.Errorf("update Role is fail"), 401
		}

		res := IResponse.Role{
			Id:         updateRole.ID,
			Name:       updateRole.Name,
			Permission: updateRole.Permission,
		}

		return res, nil, 200
	}

	updateRole, err := c.RoleRepo.UpdateRole(ctx, idInt64, name, permission)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Role{}, fmt.Errorf("update Role is fail"), 401
	}

	res := IResponse.Role{
		Id:         updateRole.ID,
		Name:       updateRole.Name,
		Permission: updateRole.Permission,
	}

	return res, nil, 200
}

func (c *RoleUseCase) DeleteRoleUseCase(ctx *gin.Context, id int) (error, int) {
	err := global.DB.DeleteRoleById(ctx, int64(id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("get Role is not exist"), 401
	}

	return nil, 200
}

func (c *RoleUseCase) DeleteManyRoleUseCase(ctx *gin.Context, arrayId []int) (error, int) {
	if len(arrayId) == 0 {
		global.Logger.Error("ArrayID is empty", zap.String("Status", "Error"))
		return fmt.Errorf("ArrayID is empty"), 401
	}

	// if slices.Contains(arrayId, constant.CONFIG_PERMISSIONS["ADMIN"].(string)) || slices.Contains(arrayId, constant.CONFIG_PERMISSIONS["BASIC"].(string)) {
	// 	global.Logger.Error("role admin or basic don't remove!!", zap.String("Status", "Error"))
	// 	return fmt.Errorf("ArrayID is empty"), 401
	// }

	arrayId64 := make([]int64, len(arrayId))

	for i, v := range arrayId {
		arrayId64[i] = int64(v)
	}

	err := global.DB.DeleteManyRolesByIds(ctx, arrayId64)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("get Role is not exist"), 401
	}

	return nil, 200
}
