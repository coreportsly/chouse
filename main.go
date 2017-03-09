package chouse

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	defaultURL  = "https://api.companieshouse.gov.uk"
	contentType = "application/json"
)

type Company struct {
	Etag                  string         `json:"etag"`
	Accounts              AnnualAccounts `json:"accounts"`
	AnnualReturn          AnnualReturn   `json:"annual_return"`
	Branch                Branch         `json:"branch_company_details"`
	CanFile               bool           `json:"can_file"`
	CompanyName           string         `json:"company_name"`
	CompanyNumber         string         `json:"company_number"`
	CompanyStatus         string         `json:"company_status"`
	CompanyStatusDetail   string         `json:"company_status_detail"`
	ConfirmationStatement AnnualReturn   `json:"confirmation_statement"`
	DateOfCessation       string         `json:"date_of_cessation"`
	DateOfCreation        string         `json:"date_of_creation"`
	ForeignCompany        ForeignCompany `json:"foreign_company_details"`
	Liquidated            bool           `json:"has_been_liquidated"`
	Charges               bool           `json:"has_charges"`
	InsolvencyHistory     bool           `json:"has_insolvency_history"`
	Cic                   bool           `json:"is_community_interest_company"`
	Jurisdiction          string         `json:"jurisdiction"`
	LastMembersList       string         `json:"last_full_members_list_date"`
	Links                 Links          `json:"links"`
	PartialData           string         `json:"partial_data_available"`
	PreviousNames         []PreviousName `json:"previous_company_names"`
	RegisteredOffice      Address        `json:"registered_office_address"`
	RoDispute             bool           `json:"registered_office_is_in_dispute"`
	SicCodes              []string       `json:"sic_codes"`
	CompanyType           string         `json:"type"`
	RoUndeliverable       bool           `json:"undeliverable_registered_office_address"`
}

type AnnualAccounts struct {
	RefDate      RefDate `json:"accounting_reference_date"`
	LastAccounts struct {
		MadeUpTo string `json:"made_up_to"`
		Type     string `json:"type"`
	} `json:"last_accounts"`
	NextDue      string `json:"next_due"`
	NextMadeUpTo string `json:"next_made_up_to"`
	Overdue      bool   `json:"overdue"`
}

type RefDate struct {
	Day   string `json:"day"`
	Month string `json:"month"`
}

type AnnualReturn struct {
	LastMadeUpTo string `json:"last_made_up_to"`
	NextDue      string `json:"next_due"`
	NextMadeUpTo string `json:"next_made_up_to"`
	Overdue      bool   `json:"overdue"`
}

type Branch struct {
	Activity            string `json:"business_activity"`
	ParentCompanyCame   string `json:"parent_company_name"`
	ParentCompanyNumber string `json:"parent_company_number"`
}

type ForeignCompany struct {
	AccountingRequirement struct {
		AccountType string `json:"foreign_account_type"`
		Terms       string `json:"terms_of_account_publication"`
	} `json:"accounting_requirement"`
	Accounts struct {
		From   RefDate `json:"account_period_from"`
		To     RefDate `json:"account_period_to"`
		Within struct {
			Months int `json:"months"`
		} `json:"must_file_within"`
	} `json:"accounts`
	BusinessActivity    string `json:"business_activity"`
	CompanyType         string `json:"company_type"`
	GovernedBy          string `json:"governed_by"`
	FinanceInstitution  bool   `json:"is_a_credit_finance_institution"`
	OriginatingRegistry struct {
		Country string `json:"country"`
		Name    string `json:"name"`
	} `json:"originating_registry"`
	RegistrationNumber string `json:"registration_number"`
}

type Links struct {
	Charges       string `json:"charges"`
	FilingHistory string `json:"filing_history"`
	Insolvency    string `json:"insolvency"`
	Officers      string `json:"officers"`
	Psc           string `json:"persons_with_significant_control"`
	PscStatements string `json:"persons_with_significant_control_statements`
	Registers     string `json:"registers"`
	Self          string `json:"self"`
}

type PreviousName struct {
	CeasedOn      string `json:"ceased_on"`
	EffectiveFrom string `json:"effective_from"`
	Name          string `json:"name"`
}

type Address struct {
	Address1 string `json:"address_line_1"`
	Address2 string `json:"address_line_2"`
	CareOf   string `json:"care_of"`
	Country  string `json:"country"`
	Locality string `json:"locality"`
	PoBox    string `json:"po_box"`
	Postcode string `json:"postal_code"`
	Premises string `json:"premises"`
	Region   string `json:"region"`
}

type FilingHistoryList struct {
	Etag         string              `json:"etag"`
	Status       string              `json:"filing_history_status"`
	Items        []FilingHistoryItem `json:"items"`
	ItemsPerPage int                 `json:"items_per_page"`
	Kind         string              `json:"kind"`
	Start        int                 `json:"start_index"`
	TotalCount   int                 `json:"total_count"`
}

type FilingHistoryItem struct {
	Annotations []struct {
		Annotation        string `json:"annotation"`
		Category          string `json:"category"`
		Date              string `json:"date"`
		Description       string `json:"description"`
		DescriptionValues struct {
			Description string `json:"description"`
		} `json:"description_values"`
		Type string `json:"type"`
	} `json:"annotations"`
	Associated []struct {
		ActionDate        int    `json:"action_date"`
		Category          string `json:"category"`
		Date              string `json:"date"`
		Description       string `json:"description"`
		DescriptionValues struct {
			Capital []struct {
				Currency string `json:"currency"`
				Figure   string `json:"figure"`
			} `json:"capital"`
			Date string `json:"date"`
		} `json:"description_values"`
		OriginalDescription string `json:"original_description"`
		Type                string `json:"type"`
	} `json:"associated_filings"`
	DescriptionValues struct {
		MadeUpDate      string `json:"made_up_date"`
		NewDate         string `json:"new_date"`
		Description     string `json:"description"`
		OfficerName     string `json:"officer_name"`
		AppointmentDate string `json:"appointment_date"`
		NewAddress      string `json:"new_address"`
		ChangeDate      string `json:"change_date"`
		OldAddress      string `json:"old_address"`
		Date            string `json:"date"`
		Capital         []struct {
			Figure   string `json:"figure"`
			Currency string `json:"currency"`
		} `json:"capital"`
	} `json:"description_values"`
	Barcode     string `json:"barcode"`
	Category    string `json:"category"`
	Date        string `json:"date"`
	ActionDate  string `json:"action_date"`
	Description string `json:"description"`
	Links       Links  `json:"links"`
	Pages       int    `json:"pages"`
	PaperFiled  bool   `json:"paper_filed"`
	Resolutions []struct {
		Category          string `json:"category"`
		DeltaAt           string `json:"delta_at"`
		Description       string `json:"description"`
		DescriptionValues struct {
			Description string `json:"description"`
			ResType     string `json:"res_type"`
		} `json:"description_values"`
		DocumentID  string `json:"document_id"`
		ReceiveDate string `json:"receive_date"`
		// It's either Array or String...
		// Subcategory struct `json:"subcategory"`
		Type string `json:"type"`
	} `json:"resolutions"`
	Subcategory   string `json:"subcategory"`
	TransactionID string `json:"transaction_id"`
	Type          string `json:"type"`
}

type OfficerList struct {
	Etag                   string        `json:"etag"`
	Kind                   string        `json:"kind"`
	Start                  int           `json:"start_index"`
	ItemsPerPage           int           `json:"items_per_page"`
	TotalResults           int           `json:"total_results"`
	ActiveAppointments     int           `json:"active_count"`
	ResignatedAppointments int           `json:"resigned_count"`
	Officers               []OfficerItem `json:"items"`
	Links                  struct {
		self string `json:"self"`
	} `json:"Links"`
}

type OfficerItem struct {
	Address            Address `json:"address"`
	AppointedOn        string  `json:"appointed_on"`
	CountryOfResidence string  `json:"country_of_residence"`
	Dob                struct {
		Day   int `json:"day"`
		Month int `json:"month"`
		Year  int `json:"year"`
	} `json:"date_of_birth"`
	FomerNames []struct {
		Forenames string `json:"forenames"`
		Surname   string `json:"surname"`
	} `json:"former_names"`
	Identification struct {
		IDType             string `json:"identification_type"`
		Authority          string `json:"legal_authority"`
		LegalForm          string `json:"legal_form"`
		PlaceRegistered    string `json:"place_registered"`
		RegistrationNumber string `json:"registration_number"`
	} `json:"identification"`
	Links struct {
		Officer struct {
			Appointments string `json:"appointments"`
		} `json:"officer"`
	} `json:"links"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Occupation  string `json:"occupation"`
	Role        string `json:"officer_role"`
	ResignedOn  string `json:"resigned_on"`
}

type CoHouseAPI struct {
	apiKey        string
	companyNumber string
}

func Explore(n string) *CoHouseAPI {
	ak := os.Getenv("CHOUSE_APIKEY")
	if ak == "" {
		fmt.Println("ERR: Env variable 'CHOUSE_APIKEY' is empty.")
		os.Exit(0)
	}
	return &CoHouseAPI{apiKey: ak, companyNumber: n}
}

func (ch *CoHouseAPI) callApi(path string) ([]byte, error) {
	url := defaultURL + path

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(ch.apiKey, "")
	req.Header.Set("Accept", contentType)

	client := &http.Client{}
	resp, errc := client.Do(req)
	if errc != nil {
		return nil, errc
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

func (ch *CoHouseAPI) Company() (*Company, error) {
	c := &Company{}

	body, err := ch.callApi("/company/" + ch.companyNumber)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (ch *CoHouseAPI) AccountFilings() (*FilingHistoryList, error) {
	return ch.Filings("accounts")
}

func (ch *CoHouseAPI) AddressFilings() (*FilingHistoryList, error) {
	return ch.Filings("address")
}

func (ch *CoHouseAPI) AnnualReturnsFilings() (*FilingHistoryList, error) {
	return ch.Filings("annual-return")
}

func (ch *CoHouseAPI) Filings(category string) (*FilingHistoryList, error) {
	fhl := &FilingHistoryList{}

	body, err := ch.callApi("/company/" + ch.companyNumber + "/filing-history?category=" + category + "&items_per_page=50")
	if err != nil {
		return fhl, err
	}

	err = json.Unmarshal(body, &fhl)
	if err != nil {
		return fhl, err
	}

	return fhl, nil
}

func (ch *CoHouseAPI) Officers() (*OfficerList, error) {
	ofr := &OfficerList{}

	body, err := ch.callApi("/company/" + ch.companyNumber + "/officers")
	if err != nil {
		return ofr, err
	}

	err = json.Unmarshal(body, &ofr)
	if err != nil {
		return ofr, err
	}

	return ofr, nil
}
