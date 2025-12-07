package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/vlahanam/company-management/common"
	"github.com/vlahanam/company-management/internal/repositories"
	"github.com/vlahanam/company-management/internal/requests"
	"github.com/vlahanam/company-management/internal/services"
)

func CreateCompany(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rq requests.CreateCompanyRequest

		if err := c.BodyParser(&rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorBodyParser)
		}

		if err := rq.Validation(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.WrapDetail(err))
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewCompanyService(rp)

		company, err := svc.CreateCompany(c.UserContext(), &rq)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		company.Mask(1)
		return c.Status(fiber.StatusCreated).JSON(common.CreateSuccessResponse("company").WrapData(company))
	}
}

func GetListCompanies(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rq requests.ListCompanyRequest

		if err := c.QueryParser(&rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorQueryParser)
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewCompanyService(rp)

		companies, err := svc.GetListCompaniesWithPagination(c.UserContext(), rq)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		for _, company := range companies {
			company.Mask(1)
		}

		return c.Status(fiber.StatusOK).JSON(common.GetListSuccessResponse("companies").WrapData(companies))
	}
}

func GetCompany(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		uid, err := common.FromBase58(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.Clone().SetDetail("id", "invalid id format"))
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewCompanyService(rp)

		company, err := svc.FindByID(c.UserContext(), uint64(uid.GetLocalID()))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(common.ErrorNotFound.Clone().WrapMessage("company not found"))
		}

		company.Mask(1)
		return c.Status(fiber.StatusOK).JSON(common.GetSuccessResponse("company").WrapData(company))
	}
}

func UpdateCompany(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		uid, err := common.FromBase58(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.Clone().SetDetail("id", "invalid id format"))
		}

		var rq requests.UpdateCompanyRequest
		if err := c.BodyParser(&rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorBodyParser)
		}

		if err := rq.Validation(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.WrapDetail(err))
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewCompanyService(rp)

		if err := svc.UpdateCompany(c.UserContext(), uint64(uid.GetLocalID()), &rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.UpdateSuccessResponse("company"))
	}
}

func DeleteCompany(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		uid, err := common.FromBase58(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.Clone().SetDetail("id", "invalid id format"))
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewCompanyService(rp)

		if err := svc.DeleteCompany(c.UserContext(), uint64(uid.GetLocalID())); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.DeleteSuccessResponse("company"))
	}
}
