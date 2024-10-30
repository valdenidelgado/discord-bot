package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	URL      = ""
	EMAIL    = ""
	PASSWORD = ""
)

type API struct {
	Token string
}

func New() *API {
	token := loginCubi()
	return &API{Token: token}
}

func (a *API) GetCompanyById(id string) string {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return "ID inválido"
	}
	token := a.Token
	url := fmt.Sprintf("%s/company/%d", URL, idInt)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var companyResp Response[Company]
	err = json.Unmarshal(body, &companyResp)
	if err != nil {

		log.Fatal(err)
	}

	company := fmt.Sprintf("Nome: %s\nCNPJ: %s\nEmail: %s\nTelefone: %s\nPremium: %d\nTradeName: %s\nNome do Contato: %s", companyResp.Data.Name, companyResp.Data.Cnpj, companyResp.Data.Contact.Email, companyResp.Data.Contact.Phone, companyResp.Data.Premium, companyResp.Data.Tradename, companyResp.Data.Contact.Name)
	return company
}

func (a *API) GetBranchById(id string) string {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return "ID inválido"
	}
	token := a.Token
	url := fmt.Sprintf("%s/branch/%d", URL, idInt)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var branchResp Response[Branch]
	err = json.Unmarshal(body, &branchResp)
	if err != nil {

		log.Fatal(err)
	}

	branch := fmt.Sprintf(
		"Nome: %s\nEmail: %s\nTelefone: %s\nCidade: %s\nEstado: %s\nLocationNumber: %s\nCNPJ: %s\nCompanyID: %d\nBillingCollectionState: %t\nBillingCollectionPermissionPerMonth: %t",
		branchResp.Data.Name,
		branchResp.Data.Email,
		branchResp.Data.Phone,
		branchResp.Data.City,
		branchResp.Data.State,
		branchResp.Data.LocationNumber,
		branchResp.Data.Cnpj,
		branchResp.Data.CompanyID,
		branchResp.Data.BillingCollectionState,
		branchResp.Data.BillingCollectionPermissionPerMonth,
	)

	return branch
}

func (a *API) GetBillingDetailsById(id string) string {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return "ID inválido"
	}
	token := a.Token
	url := fmt.Sprintf("%s/billingInfo/%d", URL, idInt)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var billingRes Response[BillingInfo]
	err = json.Unmarshal(body, &billingRes)
	if err != nil {

		log.Fatal(err)
	}

	billing := fmt.Sprintf(
		"BillingID: %d\nOriginalFileName: %s\nBranchID: %d\nType: %s\nisIgnored: %t\nStatus: %s\nBilledAt %s\n"+
			"FourvisionID: %s\nBillingFileId: %d\nMD5: %s\nMetaID: %d\nUserID: %d\nCompanyId: %d\nFourvisionStatus: %s",
		billingRes.Data.ID,
		billingRes.Data.Meta.BillingFile.OriginalFilename,
		billingRes.Data.Meta.BranchID,
		billingRes.Data.Meta.BillingFile.Pipeline,
		billingRes.Data.Meta.IsIgnored,
		billingRes.Data.Meta.Status,
		billingRes.Data.Meta.BilledAt,
		billingRes.Data.Meta.BillingFile.FourvisionID,
		billingRes.Data.Meta.BillingFile.ID,
		billingRes.Data.Meta.BillingFile.Md5,
		billingRes.Data.Meta.ID,
		billingRes.Data.Meta.BillingFile.UserID,
		billingRes.Data.Meta.BillingFile.CompanyID,
		billingRes.Data.Meta.BillingFile.Events[0].Status,
	)

	return billing
}

func loginCubi() string {
	url := fmt.Sprintf("%s/login", URL)
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("email=%s&password=%s",
			EMAIL, PASSWORD)))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)

	}
	var loginResp LoginResponse
	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		log.Fatal(err)
	}

	if !loginResp.Success {
		log.Fatal("Login failed")
	}

	token := loginResp.Data.Token
	return token
}
