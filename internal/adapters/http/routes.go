package http

import (
	db "capston-lms/internal/adapters/db/mysql"
	handler "capston-lms/internal/adapters/http/handler"
	middlewares "capston-lms/internal/adapters/http/middleware"
	repository "capston-lms/internal/adapters/repository"
	usecase "capston-lms/internal/application/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	// user management
	userRepo    repository.UserRepository
	userHandler handler.UserHandler
	userUsecase usecase.UserUseCase
	// auth
	AuthHandler handler.AuthHandler
	// class
	classRepo    repository.ClassRepository
	classHandler handler.ClassHandler
	classUsecase usecase.ClassUseCase
	// Category
	categoryRepo    repository.CategoryRepository
	categoryHandler handler.CategoryHandler
	categoryUsecase usecase.CategoryUseCase
	// Major
	majorRepo    repository.MajorRepository
	majorHandler handler.MajorHandler
	majorUsecase usecase.MajorUseCase
	// Course
	courseEnrollmentRepo    repository.CourseEnrollmentRepository
	courseEnrollmentHandler handler.CourseEnrollmentHandler
	courseEnrollmentUseCase usecase.CourseEnrollmentUseCase

	// course enrollment
	courseRepo    repository.CourseRepository
	courseHandler handler.CourseHandler
	courseUsecase usecase.CourseUseCase

	// folder
	folderRepo    repository.FolderRepository
	folderHandler handler.FolderHandler
	folderUsecase usecase.FolderUseCase

	// attachment
	attachmentRepo    repository.AttachmentRepository
	attachmentHandler handler.AttachmentHandler
	attachmentUsecase usecase.AttachmentUseCase
)

func declare() {
	// user
	userRepo = repository.UserRepository{DB: db.DbMysql}
	userUsecase = usecase.UserUseCase{Repo: userRepo}
	userHandler = handler.UserHandler{UserUsecase: userUsecase}
	// auth
	AuthHandler = handler.AuthHandler{Usecase: userUsecase}
	// class
	classRepo = repository.ClassRepository{DB: db.DbMysql}
	classUsecase = usecase.ClassUseCase{Repo: classRepo}
	classHandler = handler.ClassHandler{ClassUsecase: classUsecase}
	// category
	categoryRepo = repository.CategoryRepository{DB: db.DbMysql}
	categoryUsecase = usecase.CategoryUseCase{Repo: categoryRepo}
	categoryHandler = handler.CategoryHandler{CategoryUsecase: categoryUsecase}
	// Major
	majorRepo = repository.MajorRepository{DB: db.DbMysql}
	majorUsecase = usecase.MajorUseCase{Repo: majorRepo}
	majorHandler = handler.MajorHandler{MajorUsecase: majorUsecase}
	// Major
	courseRepo = repository.CourseRepository{DB: db.DbMysql}
	courseUsecase = usecase.CourseUseCase{Repo: courseRepo}
	courseHandler = handler.CourseHandler{CourseUsecase: courseUsecase}
	// course enrrolment
	courseEnrollmentRepo = repository.CourseEnrollmentRepository{DB: db.DbMysql}
	courseEnrollmentUseCase = usecase.CourseEnrollmentUseCase{CourseEnrollmentRepo: courseEnrollmentRepo}
	courseEnrollmentHandler = handler.CourseEnrollmentHandler{CourseEnrollmentUseCase: courseEnrollmentUseCase}
	// folder
	folderRepo = repository.FolderRepository{DB: db.DbMysql}
	folderUsecase = usecase.FolderUseCase{Repo: folderRepo}
	folderHandler = handler.FolderHandler{FolderUsecase: folderUsecase}

	// folder
	attachmentRepo = repository.AttachmentRepository{DB: db.DbMysql}
	attachmentUsecase = usecase.AttachmentUseCase{Repo: attachmentRepo}
	attachmentHandler = handler.AttachmentHandler{AttachmentUsecase: attachmentUsecase}
}

func InitRoutes() *echo.Echo {
	db.Init()
	declare()

	e := echo.New()
	e.POST("/login", AuthHandler.Login())
	e.POST("/registrasi", AuthHandler.Register())
	e.POST("/verify-otp", AuthHandler.VerifyOTP())

	// montor group
	mentors := e.Group("/mentors")
	mentors.Use(middleware.Logger())
	mentors.Use(middlewares.AuthMiddleware())
	mentors.Use(middlewares.RequireRole("mentors"))

	mentors.GET("/users", userHandler.GetAllUsers())
	mentors.GET("/users/:id", userHandler.GetUser())
	mentors.POST("/users", userHandler.CreateUser())
	mentors.DELETE("/users/:id", userHandler.DeleteUser())

	mentors.GET("/chat/students/:id", courseEnrollmentHandler.GetAllStudents())
	mentors.GET("/chat/courses", courseEnrollmentHandler.GetAllCourse())
	// route folders
	mentors.GET("/folders", folderHandler.GetAllFolders())
	mentors.GET("/folders/:id", folderHandler.GetFolder())
	mentors.POST("/folders", folderHandler.CreateFolder())
	mentors.DELETE("/folders/:id", folderHandler.DeleteFolder())

	// route attachment
	mentors.GET("/attachment/:id", attachmentHandler.GetAllAttachments())
	mentors.GET("/attachment/find/:id", attachmentHandler.GetAttachment())
	mentors.POST("/attachment", attachmentHandler.CreateAttachment())
	mentors.DELETE("/attachment/:id", attachmentHandler.DeleteAttachment())

	mentors.GET("/classes", classHandler.GetAllClasses())
	mentors.GET("/classes/:id", classHandler.GetClass())
	mentors.PUT("/classes/:id", classHandler.UpdateClass())
	mentors.POST("/classes", classHandler.CreateClass())
	mentors.DELETE("/classes/:id", classHandler.DeleteClass())

	mentors.GET("/categories", categoryHandler.GetAllCategories())
	mentors.GET("/categories/:id", categoryHandler.GetCategory())
	mentors.PUT("/cateories/:id", categoryHandler.UpdateCategory())
	mentors.POST("/categories", categoryHandler.CreateCategory())
	mentors.DELETE("/categories/:id", categoryHandler.DeleteCategory())

	mentors.GET("/majors", majorHandler.GetAllMajors())
	mentors.GET("/majors/:id", majorHandler.CreateMajor())
	mentors.PUT("/majors/:id", majorHandler.UpdateMajor())
	mentors.POST("/majors", majorHandler.CreateMajor())
	mentors.DELETE("/majors/:id", majorHandler.DeleteMajor())

	mentors.GET("/courses", courseHandler.GetAllCourses())
	mentors.GET("/courses/:id", courseHandler.CreateCourse())
	// mentors.GET("/courses/:id", courseHandler.CreateCourse())
	mentors.PUT("/courses/:id", courseHandler.UpdateCourse())
	mentors.POST("/courses", courseHandler.CreateCourse())
	mentors.DELETE("/courses/:id", courseHandler.DeleteCourse())

	// students group
	students := e.Group("/students")
	students.Use(middleware.Logger())
	students.Use(middlewares.AuthMiddleware())
	students.Use(middlewares.RequireRole("students"))

	return e
}
