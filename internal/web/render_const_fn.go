package web

import (
	"github.com/eyepipe/eye/pkg/proto"
	"github.com/gofiber/fiber/v3"
)

func (w *Web) RenderConstJSONFn(obj any) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		return c.JSON(obj)
	}
}

func (w *Web) RenderContractJSONFn(contract proto.ContractV1) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		return c.JSON(contract.WithHost(c.BaseURL()))
	}
}
