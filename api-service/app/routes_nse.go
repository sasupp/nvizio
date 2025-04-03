package app

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"xtrinio.com/db"
	"xtrinio.com/model"
	"xtrinio.com/repository"
	"xtrinio.com/service"
)

func sendServerError(c *fiber.Ctx, err error) error {
	return c.Status(500).JSON(&fiber.Map{"error": err.Error()})
}

func sendSearchResults(c *fiber.Ctx, searchService service.INseCompanySearchService, word string) error {
	searchResults, err := searchService.GetNseCompanyByWord(word)
	if err != nil {
		return sendServerError(c, err)
	}
	if len(searchResults) > 0 {
		return c.Status(300).JSON(searchResults)
	}
	apiError := model.ApiError{
		Code:    404,
		Message: "No companies found.",
	}
	return c.Status(404).JSON(apiError)
}

func SetupNseFeedRoutes(srv *fiber.App, db *db.Db) {

	nseFeedService := service.NewNseFeedService(repository.NewNseFeedRepository(db))
	nseShareholdingService := service.NewNseShareholdingService(repository.NewNseShareholdingRepository(db))
	nseInsiderService := service.NewNseInsiderService(repository.NewNseInsiderRepository(db))
	filingsService := service.NewFilingsService(repository.NewFilingsRepository(db))
	symbolsService := service.NewSymbolsService(repository.NewSymbolsRepository(db))
	nseCompanySearchService := service.NewNseCompanySearchService(repository.NewNseCompanySearchRepository(db))

	srv.Get("/nse/feed", func(c *fiber.Ctx) error {
		nse_feed, err := nseFeedService.GetNseFeedAll()
		if err != nil {
			return sendServerError(c, err)
		}
		return c.JSON(nse_feed)
	})

	srv.Get("/nse/company/:ticker/feed", func(c *fiber.Ctx) error {
		word := strings.ToLower(c.Params("ticker"))
		ticker := strings.ToUpper(word)

		if exists, err := nseCompanySearchService.TickerExists(ticker); err != nil {
			return sendServerError(c, err)
		} else if !exists {
			return sendSearchResults(c, nseCompanySearchService, word)
		}

		nse_feed, err := nseFeedService.GetNseFeedByCompany(ticker)
		if err != nil {
			return sendServerError(c, err)
		}
		return c.JSON(nse_feed)
	})

	srv.Get("/nse/company/:ticker/shp", func(c *fiber.Ctx) error {
		word := strings.ToLower(c.Params("ticker"))
		ticker := strings.ToUpper(word)

		if exists, err := nseCompanySearchService.TickerExists(ticker); err != nil {
			return sendServerError(c, err)
		} else if !exists {
			return sendSearchResults(c, nseCompanySearchService, word)
		}

		nse_shp, err := nseShareholdingService.GetNseShareholdingByCompany(ticker)
		if err != nil {
			return sendServerError(c, err)
		}
		return c.JSON(nse_shp)
	})

	srv.Get("/nse/company/:ticker/insider", func(c *fiber.Ctx) error {
		word := strings.ToLower(c.Params("ticker"))
		ticker := strings.ToUpper(word)

		if exists, err := nseCompanySearchService.TickerExists(ticker); err != nil {
			return sendServerError(c, err)
		} else if !exists {
			return sendSearchResults(c, nseCompanySearchService, word)
		}

		nse_insider, err := nseInsiderService.GetNseInsiderByCompany(ticker)
		if err != nil {
			return sendServerError(c, err)
		}
		return c.JSON(nse_insider)
	})

	srv.Get("/nse/company/:ticker/financials", func(c *fiber.Ctx) error {
		word := strings.ToLower(c.Params("ticker"))
		ticker := strings.ToUpper(word)

		if exists, err := nseCompanySearchService.TickerExists(ticker); err != nil {
			return sendServerError(c, err)
		} else if !exists {
			return sendSearchResults(c, nseCompanySearchService, word)
		}

		companyId, err := symbolsService.GetCompanyId(ticker, "NSE")
		if err != nil {
			return c.Status(404).JSON(&fiber.Map{
				"error": fmt.Sprintf("Company not found: %s", ticker),
			})
		}
		filings, err := filingsService.GetFilings(companyId)
		if err != nil {
			return sendServerError(c, err)
		}
		return c.JSON(filings)
	})

	srv.Get("/nse/company/:ticker/financials/:filing_id/docs", func(c *fiber.Ctx) error {
		filing_id := c.Params("filing_id")
		var docs []model.FilingDocumentsDto
		docs, err := filingsService.GetFilingDocuments(filing_id)
		if err != nil {
			return sendServerError(c, err)
		}
		return c.JSON(docs)
	})

	srv.Get("/nse/company/:ticker/financials/:filing_id/docs/:doc_id", func(c *fiber.Ctx) error {
		filing_id := c.Params("filing_id")
		doc_id, err := c.ParamsInt("doc_id")
		if err != nil {
			return c.Status(400).JSON(&fiber.Map{"error": "Invalid doc_id"})
		}
		doc, err := filingsService.GetFilingDocument(filing_id, doc_id)
		if err != nil {
			return sendServerError(c, err)
		}
		return c.JSON(doc)
	})
}
