package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/vlahanam/company-management/common"
	"github.com/vlahanam/company-management/internal/repositories"
	"github.com/vlahanam/company-management/internal/requests"
	"github.com/vlahanam/company-management/internal/services"
)

func CreateContract(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rq requests.CreateContractRequest

		if err := c.BodyParser(&rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorBodyParser)
		}

		if err := rq.Validation(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.WrapDetail(err))
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewContractService(rp)

		contract, err := svc.CreateContract(c.UserContext(), &rq)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		contract.Mask(1)
		return c.Status(fiber.StatusCreated).JSON(common.CreateSuccessResponse("contract").WrapData(contract))
	}
}

func GetListContracts(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rq requests.ListContractRequest

		if err := c.QueryParser(&rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorQueryParser)
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewContractService(rp)

		contracts, err := svc.GetListContractsWithPagination(c.UserContext(), rq)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		for _, contract := range contracts {
			contract.Mask(1)
		}

		return c.Status(fiber.StatusOK).JSON(common.GetListSuccessResponse("contracts").WrapData(contracts))
	}
}

func GetContract(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		uid, err := common.FromBase58(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.Clone().SetDetail("id", "invalid id format"))
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewContractService(rp)

		contract, err := svc.FindByID(c.UserContext(), uint64(uid.GetLocalID()))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(common.ErrorNotFound.Clone().WrapMessage("contract not found"))
		}

		contract.Mask(1)
		return c.Status(fiber.StatusOK).JSON(common.GetSuccessResponse("contract").WrapData(contract))
	}
}

func UpdateContract(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		uid, err := common.FromBase58(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.Clone().SetDetail("id", "invalid id format"))
		}

		var rq requests.UpdateContractRequest
		if err := c.BodyParser(&rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorBodyParser)
		}

		if err := rq.Validation(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.WrapDetail(err))
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewContractService(rp)

		if err := svc.UpdateContract(c.UserContext(), uint64(uid.GetLocalID()), &rq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.UpdateSuccessResponse("contract"))
	}
}

func DeleteContract(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		uid, err := common.FromBase58(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(common.ErrorValidation.Clone().SetDetail("id", "invalid id format"))
		}

		rp := repositories.NewMySQLStorage(db)
		svc := services.NewContractService(rp)

		if err := svc.DeleteContract(c.UserContext(), uint64(uid.GetLocalID())); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		return c.Status(fiber.StatusOK).JSON(common.DeleteSuccessResponse("contract"))
	}
}
