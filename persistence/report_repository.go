package persistence

import (
	"backend-assigment/domain"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

// IReportRepository defines the interface for report repository.
type IReportRepository interface {
	GetAllWithPagination(query string) []domain.ReportOutput
}

// ReportRepository implements the IReportRepository interface.
type ReportRepository struct {
	dbPool *pgxpool.Pool
}

// NewReportRepository creates a new instance of ReportRepository.
func NewReportRepository(dbPool *pgxpool.Pool) IReportRepository {
	return &ReportRepository{dbPool: dbPool}
}

// GetAllWithPagination retrieves reports with pagination from the database.
func (reportRepository *ReportRepository) GetAllWithPagination(query string) []domain.ReportOutput {
	ctx := context.Background()
	rows, err := reportRepository.dbPool.Query(ctx, query)
	log.Infof("Query: %s", query)
	if err != nil {
		log.Errorf("Error while getting reports: %v", err)
		return []domain.ReportOutput{}
	}
	return extractReports(rows)
}

// extractReports extracts report data from the database rows.
func extractReports(rows pgx.Rows) []domain.ReportOutput {
	var reports []domain.ReportOutput
	var row int64
	var mainUploadedVariation, mainExistingVariation, mainSymbol, details2Provean, linksMondo, linksPhenoPubmed string
	var mainAfVcf, mainDp float64
	var details2DannScore *float64

	for rows.Next() {
		err := rows.Scan(&row, &mainUploadedVariation, &mainExistingVariation, &mainSymbol, &mainAfVcf, &mainDp, &details2Provean, &details2DannScore, &linksMondo, &linksPhenoPubmed)
		if err != nil {
			log.Errorf("Error while scanning report: %v", err)
			return []domain.ReportOutput{}
		}
		reports = append(reports, domain.ReportOutput{
			Row:                   row,
			MainUploadedVariation: mainUploadedVariation,
			MainExistingVariation: mainExistingVariation,
			MainSymbol:            mainSymbol,
			MainAfVcf:             mainAfVcf,
			MainDp:                mainDp,
			Details2Provean:       details2Provean,
			Details2DannScore:     details2DannScore,
			LinksMondo:            linksMondo,
			LinksPhenoPubmed:      linksPhenoPubmed,
		})
	}
	return reports
}
