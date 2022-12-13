package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
	"github.com/octane77/rova/identityService/common"
	"github.com/octane77/rova/identityService/common/errors"
	"github.com/octane77/rova/identityService/dtos"
	"github.com/octane77/rova/identityService/services"
	"github.com/octane77/rova/identityService/utils"
	"net/http"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetUserDetails(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
	apiService  services.APIService
	jwtService  services.JwtService
}

func (u userController) CreateUser(ctx *gin.Context) {
	var createUserDto dtos.CreateUserDto
	dtoErr := ctx.ShouldBind(&createUserDto)
	if dtoErr != nil {
		errors.ThrowUnprocessableEntityException(ctx, dtoErr.Error())
		return
	}
	createdUser, err := u.userService.CreateUser(createUserDto)
	if err != nil {
		errors.ThrowBadRequestException(ctx, err.Error())
		return
	}

	res := common.BuildResponse(true, utils.USER_CREATED_MESSAGE, createdUser)
	ctx.JSON(http.StatusCreated, res)
}
func (u userController) GetUserDetails(ctx *gin.Context) {
	var accountDetailsResDto dtos.AccountDetailsResponseDto
	idString, _ := ctx.Get("userId")
	id, err := utils.ConvertAnyToUint(idString)
	if err != nil {
		errors.ThrowBadRequestException(ctx, err.Error())
		return
	}
	user, err := u.userService.GetUser(*id)
	if err != nil {
		errors.ThrowBadRequestException(ctx, err.Error())
		return
	}
	err = smapping.FillStruct(&accountDetailsResDto, smapping.MapFields(&user))
	if err != nil {
		errors.ThrowBadRequestException(ctx, err.Error())
		return
	}
	accDetails, err := u.apiService.GetAccountDetails(user)

	if err != nil {
		errors.ThrowBadRequestException(ctx, err.Error())
		return
	}
	accountDetailsResDto.Accounts = (*accDetails).(map[string]interface{})["data"]
	res := common.BuildResponse(true, utils.ACCOUNT_GOTTEN, accountDetailsResDto)
	ctx.JSON(http.StatusOK, res)
}

func (u userController) Login(ctx *gin.Context) {
	var loginDto dtos.LoginDto
	dtoErr := ctx.ShouldBind(&loginDto)
	if dtoErr != nil {
		errors.ThrowUnprocessableEntityException(ctx, dtoErr.Error())
		return
	}
	user, err := u.userService.GetUserByName(loginDto.Name)
	if err != nil {
		errors.ThrowBadRequestException(ctx, utils.INVALID_CREDENTIALS)
		return
	}
	if isValidPassword := utils.ComparePasswordWithHash(loginDto.Password, user.Password); !isValidPassword {
		errors.ThrowBadRequestException(ctx, utils.INVALID_CREDENTIALS)
		return
	}
	token := u.jwtService.GenerateToken(user.ID)
	res := common.BuildResponse(true, utils.ACCOUNT_GOTTEN, dtos.LoginResponseDto{
		User:  *user,
		Token: token,
	})
	ctx.JSON(http.StatusOK, res)
}

func NewUserController(userService services.UserService, apiService services.APIService, jwtService services.JwtService) UserController {
	return userController{
		userService: userService,
		apiService:  apiService,
		jwtService:  jwtService,
	}
}
