package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/vlahanam/company-management/common"
	"github.com/vlahanam/company-management/internal/repositories"
	"github.com/vlahanam/company-management/internal/requests"
	"github.com/vlahanam/company-management/internal/services"
)

func CreateRole(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rq requests.CreateRoleRequest

		if err := c.BodyParser(&rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorBodyParser)
		}

		if err := rq.Validation(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.WrapDetail(err))
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewRoleService(rp)

		role, err := svc.CreateRole(c.UserContext(), &rq)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		return c.Status(fiber.StatusCreated).JSON(common.CreateSuccessResponse("role").WrapData(role))
	}
}

func GetListRoles(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rq requests.ListRoleRequest

		if err := c.QueryParser(&rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorQueryParser)
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewRoleService(rp)

		roles, err := svc.GetListRolesWithPagination(c.UserContext(), rq)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.GetListSuccessResponse("roles").WrapData(roles))
	}
}

func GetRole(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.Clone().SetDetail("id", "invalid id format"))
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewRoleService(rp)

		role, err := svc.FindByID(c.UserContext(), id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(common.ErrorNotFound.Clone().WrapMessage("role not found"))
		}

		return c.Status(fiber.StatusOK).JSON(common.GetSuccessResponse("role").WrapData(role))
	}
}

func UpdateRole(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.Clone().SetDetail("id", "invalid id format"))
		}

		var rq requests.UpdateRoleRequest
		if err := c.BodyParser(&rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorBodyParser)
		}

		if err := rq.Validation(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.WrapDetail(err))
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewRoleService(rp)

		if err := svc.UpdateRole(c.UserContext(), id, &rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.UpdateSuccessResponse("role"))
	}
}

func DeleteRole(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.Clone().SetDetail("id", "invalid id format"))
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewRoleService(rp)

		if err := svc.DeleteRole(c.UserContext(), id); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.DeleteSuccessResponse("role"))
	}
}
