// +build ignore

package salesrule

import (
	"github.com/corestoreio/csfw/config/model"
)

// PathPromoAutoGeneratedCouponCodesLength => Code Length.
// Excluding prefix, suffix and separators.
var PathPromoAutoGeneratedCouponCodesLength = model.NewStr(`promo/auto_generated_coupon_codes/length`, model.WithPkgCfg(PackageConfiguration))

// PathPromoAutoGeneratedCouponCodesFormat => Code Format.
// SourceModel: Otnegam\SalesRule\Model\System\Config\Source\Coupon\Format
var PathPromoAutoGeneratedCouponCodesFormat = model.NewStr(`promo/auto_generated_coupon_codes/format`, model.WithPkgCfg(PackageConfiguration))

// PathPromoAutoGeneratedCouponCodesPrefix => Code Prefix.
var PathPromoAutoGeneratedCouponCodesPrefix = model.NewStr(`promo/auto_generated_coupon_codes/prefix`, model.WithPkgCfg(PackageConfiguration))

// PathPromoAutoGeneratedCouponCodesSuffix => Code Suffix.
var PathPromoAutoGeneratedCouponCodesSuffix = model.NewStr(`promo/auto_generated_coupon_codes/suffix`, model.WithPkgCfg(PackageConfiguration))

// PathPromoAutoGeneratedCouponCodesDash => Dash Every X Characters.
// If empty no separation.
var PathPromoAutoGeneratedCouponCodesDash = model.NewStr(`promo/auto_generated_coupon_codes/dash`, model.WithPkgCfg(PackageConfiguration))

// PathRssCatalogDiscounts => Coupons/Discounts.
// SourceModel: Otnegam\Config\Model\Config\Source\Enabledisable
var PathRssCatalogDiscounts = model.NewBool(`rss/catalog/discounts`, model.WithPkgCfg(PackageConfiguration))
