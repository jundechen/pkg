// Code generated by codegen. DO NOT EDIT.
// Generated by sql/dmlgen. DO NOT EDIT.
package dmltestgeneratedMToM

import (
	"time"

	"github.com/corestoreio/pkg/storage/null"
)

// Athlete represents a single row for DB table athlete. Auto generated.
// Table comment: Athletes
type Athlete struct {
	AthleteID    uint32        `max_len:"10"`  // athlete_id int(10) unsigned NOT NULL PRI  auto_increment "Athlete ID"
	Firstname    null.String   `max_len:"340"` // firstname varchar(340) NULL  DEFAULT 'NULL'  "First Name"
	Lastname     null.String   `max_len:"340"` // lastname varchar(340) NULL  DEFAULT 'NULL'  "Last Name"
	AthleteTeams *AthleteTeams // Reversed M:N athlete.athlete_id via athlete_team_member.athlete_id => athlete_team.team_id
}

// Athletes represents a collection type for DB table athlete
// Not thread safe. Auto generated.
type Athletes struct {
	Data []*Athlete `json:"data,omitempty"`
}

// NewAthletes  creates a new initialized collection. Auto generated.
func NewAthletes() *Athletes {
	return &Athletes{
		Data: make([]*Athlete, 0, 5),
	}
}

// AthleteTeam represents a single row for DB table athlete_team. Auto generated.
// Table comment: Athlete Team
type AthleteTeam struct {
	TeamID   uint32    `max_len:"10"`  // team_id int(10) unsigned NOT NULL PRI  auto_increment ""
	Name     string    `max_len:"340"` // name varchar(340) NOT NULL    "Team name"
	Athletes *Athletes // Reversed M:N athlete_team.team_id via athlete_team_member.team_id => athlete.athlete_id
}

// AthleteTeams represents a collection type for DB table athlete_team
// Not thread safe. Auto generated.
type AthleteTeams struct {
	Data []*AthleteTeam `json:"data,omitempty"`
}

// NewAthleteTeams  creates a new initialized collection. Auto generated.
func NewAthleteTeams() *AthleteTeams {
	return &AthleteTeams{
		Data: make([]*AthleteTeam, 0, 5),
	}
}

// CustomerAddressEntity represents a single row for DB table
// customer_address_entity. Auto generated.
// Table comment: Customer Address Entity
type CustomerAddressEntity struct {
	EntityID          uint32      // entity_id int(10) unsigned NOT NULL PRI  auto_increment "Entity ID"
	IncrementID       null.String // increment_id varchar(50) NULL  DEFAULT 'NULL'  "Increment Id"
	ParentID          null.Uint32 // parent_id int(10) unsigned NULL MUL DEFAULT 'NULL'  "Parent ID"
	CreatedAt         time.Time   // created_at timestamp NOT NULL  DEFAULT 'current_timestamp()'  "Created At"
	UpdatedAt         time.Time   // updated_at timestamp NOT NULL  DEFAULT 'current_timestamp()' on update current_timestamp() "Updated At"
	IsActive          bool        // is_active smallint(5) unsigned NOT NULL  DEFAULT '1'  "Is Active"
	City              string      // city varchar(255) NOT NULL    "City"
	Company           null.String // company varchar(255) NULL  DEFAULT 'NULL'  "Company"
	CountryID         string      // country_id varchar(255) NOT NULL    "Country"
	Fax               null.String // fax varchar(255) NULL  DEFAULT 'NULL'  "Fax"
	Firstname         string      // firstname varchar(255) NOT NULL    "First Name"
	Lastname          string      // lastname varchar(255) NOT NULL    "Last Name"
	Middlename        null.String // middlename varchar(255) NULL  DEFAULT 'NULL'  "Middle Name"
	Postcode          null.String // postcode varchar(255) NULL  DEFAULT 'NULL'  "Zip/Postal Code"
	Prefix            null.String // prefix varchar(40) NULL  DEFAULT 'NULL'  "Name Prefix"
	Region            null.String // region varchar(255) NULL  DEFAULT 'NULL'  "State/Province"
	RegionID          null.Uint32 // region_id int(10) unsigned NULL  DEFAULT 'NULL'  "State/Province"
	Street            string      // street text NOT NULL    "Street Address"
	Suffix            null.String // suffix varchar(40) NULL  DEFAULT 'NULL'  "Name Suffix"
	Telephone         string      // telephone varchar(255) NOT NULL    "Phone Number"
	VatID             null.String // vat_id varchar(255) NULL  DEFAULT 'NULL'  "VAT number"
	VatIsValid        null.Bool   // vat_is_valid int(10) unsigned NULL  DEFAULT 'NULL'  "VAT number validity"
	VatRequestDate    null.String // vat_request_date varchar(255) NULL  DEFAULT 'NULL'  "VAT number validation request date"
	VatRequestID      null.String // vat_request_id varchar(255) NULL  DEFAULT 'NULL'  "VAT number validation request ID"
	VatRequestSuccess null.Uint32 // vat_request_success int(10) unsigned NULL  DEFAULT 'NULL'  "VAT number validation request success"
}

// CustomerAddressEntities represents a collection type for DB table
// customer_address_entity
// Not thread safe. Auto generated.
type CustomerAddressEntities struct {
	Data []*CustomerAddressEntity `json:"data,omitempty"`
}

// NewCustomerAddressEntities  creates a new initialized collection. Auto
// generated.
func NewCustomerAddressEntities() *CustomerAddressEntities {
	return &CustomerAddressEntities{
		Data: make([]*CustomerAddressEntity, 0, 5),
	}
}

// CustomerEntity represents a single row for DB table customer_entity. Auto
// generated.
// Table comment: Customer Entity
type CustomerEntity struct {
	EntityID                uint32                   `max_len:"10"`  // entity_id int(10) unsigned NOT NULL PRI  auto_increment "Entity ID"
	WebsiteID               null.Uint16              `max_len:"5"`   // website_id smallint(5) unsigned NULL MUL DEFAULT 'NULL'  "Website ID"
	Email                   null.String              `max_len:"255"` // email varchar(255) NULL MUL DEFAULT 'NULL'  "Email"
	GroupID                 uint16                   `max_len:"5"`   // group_id smallint(5) unsigned NOT NULL  DEFAULT '0'  "Group ID"
	IncrementID             null.String              `max_len:"50"`  // increment_id varchar(50) NULL  DEFAULT 'NULL'  "Increment Id"
	StoreID                 null.Uint16              `max_len:"5"`   // store_id smallint(5) unsigned NULL MUL DEFAULT '0'  "Store ID"
	CreatedAt               time.Time                // created_at timestamp NOT NULL  DEFAULT 'current_timestamp()'  "Created At"
	UpdatedAt               time.Time                // updated_at timestamp NOT NULL  DEFAULT 'current_timestamp()' on update current_timestamp() "Updated At"
	IsActive                bool                     `max_len:"5"`   // is_active smallint(5) unsigned NOT NULL  DEFAULT '1'  "Is Active"
	DisableAutoGroupChange  uint16                   `max_len:"5"`   // disable_auto_group_change smallint(5) unsigned NOT NULL  DEFAULT '0'  "Disable automatic group change based on VAT ID"
	CreatedIn               null.String              `max_len:"255"` // created_in varchar(255) NULL  DEFAULT 'NULL'  "Created From"
	Prefix                  null.String              `max_len:"40"`  // prefix varchar(40) NULL  DEFAULT 'NULL'  "Name Prefix"
	Firstname               null.String              `max_len:"255"` // firstname varchar(255) NULL MUL DEFAULT 'NULL'  "First Name"
	Middlename              null.String              `max_len:"255"` // middlename varchar(255) NULL  DEFAULT 'NULL'  "Middle Name/Initial"
	Lastname                null.String              `max_len:"255"` // lastname varchar(255) NULL MUL DEFAULT 'NULL'  "Last Name"
	Suffix                  null.String              `max_len:"40"`  // suffix varchar(40) NULL  DEFAULT 'NULL'  "Name Suffix"
	Dob                     null.Time                // dob date NULL  DEFAULT 'NULL'  "Date of Birth"
	PasswordHash            null.String              `max_len:"128"` // password_hash varchar(128) NULL  DEFAULT 'NULL'  "Password_hash"
	RpToken                 null.String              `max_len:"128"` // rp_token varchar(128) NULL  DEFAULT 'NULL'  "Reset password token"
	RpTokenCreatedAt        null.Time                // rp_token_created_at datetime NULL  DEFAULT 'NULL'  "Reset password token creation time"
	DefaultBilling          null.Uint32              `max_len:"10"` // default_billing int(10) unsigned NULL  DEFAULT 'NULL'  "Default Billing Address"
	DefaultShipping         null.Uint32              `max_len:"10"` // default_shipping int(10) unsigned NULL  DEFAULT 'NULL'  "Default Shipping Address"
	Taxvat                  null.String              `max_len:"50"` // taxvat varchar(50) NULL  DEFAULT 'NULL'  "Tax/VAT Number"
	Confirmation            null.String              `max_len:"64"` // confirmation varchar(64) NULL  DEFAULT 'NULL'  "Is Confirmed"
	Gender                  null.Uint16              `max_len:"5"`  // gender smallint(5) unsigned NULL  DEFAULT 'NULL'  "Gender"
	FailuresNum             null.Int16               `max_len:"5"`  // failures_num smallint(6) NULL  DEFAULT '0'  "Failure Number"
	FirstFailure            null.Time                // first_failure timestamp NULL  DEFAULT 'NULL'  "First Failure"
	LockExpires             null.Time                // lock_expires timestamp NULL  DEFAULT 'NULL'  "Lock Expiration Date"
	Customeraddressentities *CustomerAddressEntities // Reversed 1:M customer_entity.entity_id => customer_address_entity.parent_id
}

// CustomerEntities represents a collection type for DB table customer_entity
// Not thread safe. Auto generated.
type CustomerEntities struct {
	Data []*CustomerEntity `json:"data,omitempty"`
}

// NewCustomerEntities  creates a new initialized collection. Auto generated.
func NewCustomerEntities() *CustomerEntities {
	return &CustomerEntities{
		Data: make([]*CustomerEntity, 0, 5),
	}
}
