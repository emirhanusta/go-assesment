package domain

type ReportOutput struct {
	Row                   int64    `json:"row"`
	MainUploadedVariation string   `json:"main_uploaded_variation"`
	MainExistingVariation string   `json:"main_existing_variation"`
	MainSymbol            string   `json:"main_symbol"`
	MainAfVcf             float64  `json:"main_af_vcf"`
	MainDp                float64  `json:"main_dp"`
	Details2Provean       string   `json:"details2_provean"`
	Details2DannScore     *float64 `json:"details2_dann_score,omitempty"`
	LinksMondo            string   `json:"links_mondo"`
	LinksPhenoPubmed      string   `json:"links_pheno_pubmed"`
}
