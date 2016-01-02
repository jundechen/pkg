// +build ignore

package weee

import (
	"github.com/corestoreio/csfw/config/model"
)

// PathTaxWeeeEnable => Enable FPT.
// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
var PathTaxWeeeEnable = model.NewBool(`tax/weee/enable`, model.WithPkgCfg(PackageConfiguration))

// PathTaxWeeeDisplayList => Display Prices In Product Lists.
// SourceModel: Otnegam\Weee\Model\Config\Source\Display
var PathTaxWeeeDisplayList = model.NewStr(`tax/weee/display_list`, model.WithPkgCfg(PackageConfiguration))

// PathTaxWeeeDisplay => Display Prices On Product View Page.
// SourceModel: Otnegam\Weee\Model\Config\Source\Display
var PathTaxWeeeDisplay = model.NewStr(`tax/weee/display`, model.WithPkgCfg(PackageConfiguration))

// PathTaxWeeeDisplaySales => Display Prices In Sales Modules.
// SourceModel: Otnegam\Weee\Model\Config\Source\Display
var PathTaxWeeeDisplaySales = model.NewStr(`tax/weee/display_sales`, model.WithPkgCfg(PackageConfiguration))

// PathTaxWeeeDisplayEmail => Display Prices In Emails.
// SourceModel: Otnegam\Weee\Model\Config\Source\Display
var PathTaxWeeeDisplayEmail = model.NewStr(`tax/weee/display_email`, model.WithPkgCfg(PackageConfiguration))

// PathTaxWeeeApplyVat => Apply Tax To FPT.
// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
var PathTaxWeeeApplyVat = model.NewBool(`tax/weee/apply_vat`, model.WithPkgCfg(PackageConfiguration))

// PathTaxWeeeIncludeInSubtotal => Include FPT In Subtotal.
// SourceModel: Otnegam\Config\Model\Config\Source\Yesno
var PathTaxWeeeIncludeInSubtotal = model.NewBool(`tax/weee/include_in_subtotal`, model.WithPkgCfg(PackageConfiguration))

// PathSalesTotalsSortWeee => Fixed Product Tax.
var PathSalesTotalsSortWeee = model.NewStr(`sales/totals_sort/weee`, model.WithPkgCfg(PackageConfiguration))
