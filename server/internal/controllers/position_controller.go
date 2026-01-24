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

func CreatePosition(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		companyIDStr := c.Params("company_id")
		companyID, err := strconv.ParseUint(companyIDStr, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.Clone().SetDetail("company_id", "invalid company id"))
		}

		var rq requests.CreatePositionRequest
		if err := c.BodyParser(&rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorBodyParser)
		}

		if err := rq.Validation(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.WrapDetail(err))
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewPositionService(rp)

		position, err := svc.CreatePosition(c.UserContext(), companyID, &rq)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		position.Mask(1)
		return c.Status(fiber.StatusCreated).JSON(common.CreateSuccessResponse("position").WrapData(position))
	}
}

func GetListPositions(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		companyIDStr := c.Params("company_id")
		companyID, err := strconv.ParseUint(companyIDStr, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.Clone().SetDetail("company_id", "invalid company id"))
		}

		var rq requests.ListPositionRequest
		if err := c.QueryParser(&rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorQueryParser)
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewPositionService(rp)

		positions, err := svc.GetPositionsByCompanyWithPagination(c.UserContext(), companyID, rq)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		for _, position := range positions {
			position.Mask(1)
		}

		return c.Status(fiber.StatusOK).JSON(common.GetListSuccessResponse("positions").WrapData(positions))
	}
}

func GetPosition(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		uid, err := common.FromBase58(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.Clone().SetDetail("id", "invalid id format"))
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewPositionService(rp)

		position, err := svc.FindByID(c.UserContext(), uint64(uid.GetLocalID()))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(common.ErrorNotFound.Clone().WrapMessage("position not found"))
		}

		position.Mask(1)
		return c.Status(fiber.StatusOK).JSON(common.GetSuccessResponse("position").WrapData(position))
	}
}

func UpdatePosition(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		uid, err := common.FromBase58(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.Clone().SetDetail("id", "invalid id format"))
		}

		var rq requests.UpdatePositionRequest
		if err := c.BodyParser(&rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorBodyParser)
		}

		if err := rq.Validation(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.WrapDetail(err))
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewPositionService(rp)

		if err := svc.UpdatePosition(c.UserContext(), uint64(uid.GetLocalID()), &rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.UpdateSuccessResponse("position"))
	}
}

func DeletePosition(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		uid, err := common.FromBase58(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.Clone().SetDetail("id", "invalid id format"))
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewPositionService(rp)

		if err := svc.DeletePosition(c.UserContext(), uint64(uid.GetLocalID())); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.DeleteSuccessResponse("position"))
	}
}
