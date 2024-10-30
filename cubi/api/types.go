package api

import "time"

type LoginResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Token string `json:"token"`
	} `json:"data"`
}

type Response[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

type Company struct {
	CreatedAt                   time.Time `json:"createdAt"`
	UpdatedAt                   time.Time `json:"updatedAt"`
	CreatedByID                 int       `json:"createdById"`
	DeletedAt                   any       `json:"deletedAt"`
	DeletedByID                 any       `json:"deletedById"`
	ID                          int       `json:"id"`
	UpdatedByID                 int       `json:"updatedById"`
	AddressID                   int       `json:"addressId"`
	Cnae                        string    `json:"cnae"`
	Cnpj                        string    `json:"cnpj"`
	ContactID                   int       `json:"contactId"`
	Department                  string    `json:"department"`
	Employees                   int       `json:"employees"`
	Name                        string    `json:"name"`
	SubDepartment               string    `json:"subDepartment"`
	Timezone                    string    `json:"timezone"`
	Tradename                   string    `json:"tradename"`
	Website                     string    `json:"website"`
	Premium                     int       `json:"premium"`
	PermissionCollectionBilling bool      `json:"permissionCollectionBilling"`
	Contact                     Contact   `json:"contact"`
}

type Contact struct {
	Email    string `json:"email"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Position string `json:"position"`
}

type Branch struct {
	ClassName                           string  `json:"className"`
	Email                               string  `json:"email"`
	ID                                  string  `json:"id"`
	Name                                string  `json:"name"`
	Phone                               string  `json:"phone"`
	Position                            string  `json:"position"`
	City                                string  `json:"city"`
	Complement                          string  `json:"complement"`
	District                            string  `json:"district"`
	Postcode                            string  `json:"postcode"`
	State                               string  `json:"state"`
	StreetName                          string  `json:"streetName"`
	StreetNumber                        string  `json:"streetNumber"`
	Cnpj                                string  `json:"cnpj"`
	Cnae                                string  `json:"cnae"`
	LocationNumber                      string  `json:"locationNumber"`
	Employees                           int     `json:"employees"`
	Department                          string  `json:"department"`
	SubDepartment                       string  `json:"subDepartment"`
	Address                             Address `json:"address"`
	Tradename                           string  `json:"tradename"`
	CompanyID                           int     `json:"companyId"`
	BillingCollectionState              bool    `json:"billingCollectionState"`
	BillingCollectionPermissionPerMonth bool    `json:"billingCollectionPermissionPerMonth"`
	NonDeletable                        bool    `json:"nonDeletable"`
}

type Address struct {
	ClassName    string `json:"className"`
	City         string `json:"city"`
	Complement   string `json:"complement"`
	District     string `json:"district"`
	ID           int    `json:"id"`
	Postcode     string `json:"postcode"`
	State        string `json:"state"`
	StreetName   string `json:"streetName"`
	StreetNumber string `json:"streetNumber"`
}

type Billing struct {
	CreatedAt        time.Time         `json:"createdAt"`
	UpdatedAt        time.Time         `json:"updatedAt"`
	ClassName        string            `json:"className"`
	CreatedByID      int               `json:"createdById"`
	DeletedAt        interface{}       `json:"deletedAt"`
	DeletedByID      interface{}       `json:"deletedById"`
	ID               int               `json:"id"`
	UpdatedByID      int               `json:"updatedById"`
	FourvisionID     string            `json:"fourvisionId"`
	S3Path           string            `json:"s3path"`
	OriginalFilename string            `json:"originalFilename"`
	DefaultFileName  interface{}       `json:"defaultFileName"`
	UploadDatetime   time.Time         `json:"uploadDatetime"`
	UserID           int               `json:"userId"`
	CompanyID        int               `json:"companyId"`
	Md5              string            `json:"md5"`
	PdfPassword      interface{}       `json:"pdfPassword"`
	Pipeline         string            `json:"pipeline"`
	Events           []FourvisionEvent `json:"__events__"`
}

type FourvisionEvent struct {
	ID           int       `json:"id"`
	FourvisionID string    `json:"fourvisionId"`
	Message      string    `json:"message"`
	Status       string    `json:"status"`
	ReceivedAt   time.Time `json:"receivedAt"`
	WhenChanged  time.Time `json:"whenChanged"`
}

type BillingMeta struct {
	ID            int         `json:"id"`
	IsIgnored     bool        `json:"isIgnored"`
	BranchID      int         `json:"branchId"`
	BillingFileID int         `json:"billingFileId"`
	UserID        int         `json:"userId"`
	Status        string      `json:"status"`
	PendingReason interface{} `json:"pendingReason"`
	PendingDetail interface{} `json:"pendingDetail"`
	IsManual      interface{} `json:"isManual"`
	IsHistory     bool        `json:"isHistory"`
	UUID          interface{} `json:"uuid"`
	UpdatedAt     time.Time   `json:"updatedAt"`
	BilledAt      time.Time   `json:"billedAt"`
	BillingFile   Billing     `json:"billingFile"`
}

type BillingInfo struct {
	ID         int         `json:"id"`
	BasicID    int         `json:"basicId"`
	MetaID     int         `json:"metaId"`
	RatesID    int         `json:"ratesId"`
	ValuesID   int         `json:"valuesId"`
	MeasuresID int         `json:"measuresId"`
	Meta       BillingMeta `json:"meta"`
}

// INFO: '/api/v2/billingInfo',
// INFO: '/api/v2/billingWater',
// INFO: '/api/v2/billingDanfe',
