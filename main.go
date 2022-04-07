package main

import (
	"backend/config"
	"backend/controller/global/company"
	"backend/controller/global/departement"
	"backend/controller/global/division"
	"backend/controller/global/global"
	"backend/controller/global/log"
	"backend/controller/global/module"
	"backend/controller/global/partner"
	"backend/controller/global/permission"
	"backend/controller/global/user"
	_ "backend/docs"
	"backend/middleware"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var (
	//DB CONNECTION
	db *gorm.DB = config.SetupDBConnection()

	//USER ROUTES
	userRepository user.UsersRepository = user.NewUsersRepository(db)
	userService    user.UserService     = user.NewUserService(userRepository)
	userController user.UserController  = user.NewUserController(userService, jwtService)

	//AUTH ROUTES
	authService    user.AuthService    = user.NewAuthService(userRepository)
	jwtService     user.JWTService     = user.NewJWTService()
	authController user.AuthController = user.NewAuthController(authService, jwtService)

	//MODULE ROUTES
	modulesRepository module.ModulesRepository = module.NewModulesRepository(db)
	modulesService    module.ModulesService    = module.NewModulesService(modulesRepository)
	modulesController module.ModulesController = module.NewModulesController(modulesService, jwtService)

	//COMPANY ROUTES
	companyRepository company.CompanyRepository = company.NewCompanyRepository(db)
	companyService    company.CompanyService    = company.NewCompanyService(companyRepository)
	companyController company.CompanyController = company.NewCompanyController(companyService, jwtService)

	//DEPARTEMENT ROUTES
	departementRepository departement.DepartementRepository = departement.NewDepartementRepository(db)
	departementService    departement.DepartementService    = departement.NewDepartementService(departementRepository)
	departementController departement.DepartementController = departement.NewDepartementController(departementService, jwtService)

	//DIVISION ROUTES
	divisionRepository division.DivisionRepository = division.NewDivisionRepository(db)
	divisionService    division.DivisionService    = division.NewDivisionService(divisionRepository)
	divisionController division.DivisionController = division.NewDivisionController(divisionService, jwtService)

	//GLOBAL ROUTES
	globalRepository global.GlobalRepository = global.NewGlobalRepository(db)
	globalService    global.GlobalService    = global.NewGlobalService(globalRepository)
	globalController global.GlobalController = global.NewGlobalController(globalService, jwtService)

	//LOG ROUTES
	logRepository log.LogRepository = log.NewLogRepository(db)
	logService    log.LogService    = log.NewLogService(logRepository)
	logController log.LogController = log.NewLogController(logService, jwtService)

	//PARTNER ROUTES
	partnerRepository partner.PartnerRepository = partner.NewPartnerRepository(db)
	partnerService    partner.PartnerService    = partner.NewPartnerService(partnerRepository)
	partnerController partner.PartnerController = partner.NewPartnerController(partnerService, jwtService)

	//PERMISSION ROUTES
	permissionRepository permission.PermissionRepository = permission.NewPermissionRepository(db)
	permissionService    permission.PermissionService    = permission.NewPermissionService(permissionRepository)
	permissionController permission.PermissionController = permission.NewPermissionController(permissionService, jwtService)
)

// @title API DOCUMENTATION - API FOR SIPIL 2022
// @version 1.0
// @description This is a Api Documentation for SIPIL 2022 with Golang Backend, Gin Gonic Framework, GORM MySQL, and JWT Authentication
// @termsOfService http://swagger.io/terms/
// contact.name API Support
// contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// routes.GlobalRoutes(db)
	// routes.UserRoutes(db)
	r.Use(middleware.SetupCorsMiddleware())
	r.StaticFS("/web", http.Dir("web"))

	authRoute := r.Group("api/auth")
	{
		authRoute.POST("/login", authController.Login)
		authRoute.POST("/register", authController.Register)
	}

	usersRoute := r.Group("api/users")
	{
		usersRoute.GET("/", userController.Profile)
		usersRoute.PUT("/:id", userController.Update)
	}

	modulesRoute := r.Group("api/module")
	{
		modulesRoute.GET("/", modulesController.GetAllModules)
		modulesRoute.POST("/", modulesController.InsertModules)
		modulesRoute.GET("/:id", modulesController.FindModulesByID)
		modulesRoute.PUT("/:id", modulesController.UpdateModules)
		modulesRoute.DELETE("/:id", modulesController.DeleteModules)
	}

	companyRoute := r.Group("api/company")
	{
		companyRoute.GET("/", companyController.GetAllCompany)
		companyRoute.POST("/", companyController.InsertCompany)
		companyRoute.GET("/:id", companyController.FindCompanyByID)
		companyRoute.PUT("/:id", companyController.UpdateCompany)
		companyRoute.DELETE("/:id", companyController.DeleteCompany)
	}

	departementRoute := r.Group("api/departement")
	{
		departementRoute.GET("/", departementController.GetAllDepartement)
		departementRoute.POST("/", departementController.InsertDepartement)
		departementRoute.GET("/:id", departementController.FindDepartementByID)
		departementRoute.PUT("/:id", departementController.UpdateDepartement)
		departementRoute.DELETE("/:id", departementController.DeleteDepartement)
	}

	divisionRoute := r.Group("api/division")
	{
		divisionRoute.GET("/", divisionController.GetAllDivision)
		divisionRoute.POST("/", divisionController.InsertDivision)
		divisionRoute.GET("/:id", divisionController.FindDivisionByID)
		divisionRoute.PUT("/:id", divisionController.UpdateDivision)
		divisionRoute.DELETE("/:id", divisionController.DeleteDivision)
	}

	globalRoute := r.Group("api/global")
	{
		globalRoute.GET("/", globalController.GetAllGlobal)
		globalRoute.POST("/", globalController.InsertGlobal)
		globalRoute.GET("/:id", globalController.FindGlobalByID)
		globalRoute.PUT("/:id", globalController.UpdateGlobal)
		globalRoute.DELETE("/:id", globalController.DeleteGlobal)
	}

	logRoute := r.Group("api/log")
	{
		logRoute.GET("/", logController.GetAllLog)
		logRoute.POST("/", logController.InsertLog)
		logRoute.GET("/:id", logController.FindLogByID)
		logRoute.PUT("/:id", logController.UpdateLog)
		logRoute.DELETE("/:id", logController.DeleteLog)
	}

	partnerRoute := r.Group("api/partner")
	{
		partnerRoute.GET("/", partnerController.GetAllPartner)
		partnerRoute.POST("/", partnerController.InsertPartner)
		partnerRoute.GET("/:id", partnerController.FindPartnerByID)
		partnerRoute.PUT("/:id", partnerController.UpdatePartner)
		partnerRoute.DELETE("/:id", partnerController.DeletePartner)
	}

	permissionRoute := r.Group("api/permission")
	{
		permissionRoute.GET("/", permissionController.GetAllPermission)
		permissionRoute.POST("/", permissionController.InsertPermission)
		permissionRoute.GET("/:id", permissionController.FindPermissionByID)
		permissionRoute.PUT("/:id", permissionController.UpdatePermission)
		permissionRoute.DELETE("/:id", permissionController.DeletePermission)
	}

	r.Run(": " + os.Getenv("SERVER_PORT"))
}
