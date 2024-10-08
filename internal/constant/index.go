package constant

type TMapPermission map[string]interface{}

var CONFIG_PERMISSIONS = TMapPermission{
	"ADMIN":     "ADMIN.GRANTED",
	"BASIC":     "BASIC.PUBLIC",
	"DASHBOARD": "DASHBOARD",
	"MANAGE_PRODUCT": TMapPermission{
		"PRODUCT": TMapPermission{
			"CREATE": "MANAGE_PRODUCT.PRODUCT.CREATE",
			"VIEW":   "MANAGE_PRODUCT.PRODUCT.VIEW",
			"UPDATE": "MANAGE_PRODUCT.PRODUCT.UPDATE",
			"DELETE": "MANAGE_PRODUCT.PRODUCT.DELETE",
		},
		"PRODUCT_TYPE": TMapPermission{
			"CREATE": "MANAGE_PRODUCT.PRODUCT_TYPE.CREATE",
			"UPDATE": "MANAGE_PRODUCT.PRODUCT_TYPE.UPDATE",
			"DELETE": "MANAGE_PRODUCT.PRODUCT_TYPE.DELETE",
		},
		"COMMENT": TMapPermission{
			"UPDATE": "MANAGE_PRODUCT.COMMENT.UPDATE",
			"DELETE": "MANAGE_PRODUCT.COMMENT.DELETE",
		},
	},
	"SYSTEM": TMapPermission{
		"USER": TMapPermission{
			"VIEW":   "SYSTEM.USER.VIEW",
			"CREATE": "SYSTEM.USER.CREATE",
			"UPDATE": "SYSTEM.USER.UPDATE",
			"DELETE": "SYSTEM.USER.DELETE",
		},
		"ROLE": TMapPermission{
			"VIEW":   "SYSTEM.ROLE.VIEW",
			"CREATE": "SYSTEM.ROLE.CREATE",
			"UPDATE": "SYSTEM.ROLE.UPDATE",
			"DELETE": "SYSTEM.ROLE.DELETE",
		},
	},
	"MANAGE_ORDER": TMapPermission{
		"REVIEW": TMapPermission{
			"UPDATE": "MANAGE_ORDER.REVIEW.UPDATE",
			"DELETE": "MANAGE_ORDER.REVIEW.DELETE",
		},
		"ORDER": TMapPermission{
			"VIEW":   "MANAGE_ORDER.ORDER.VIEW",
			"UPDATE": "MANAGE_ORDER.ORDER.UPDATE",
			"DELETE": "MANAGE_ORDER.ORDER.DELETE",
		},
	},
	"SETTING": TMapPermission{
		"PAYMENT_TYPE": TMapPermission{
			"CREATE": "SETTING.PAYMENT_TYPE.CREATE",
			"UPDATE": "SETTING.PAYMENT_TYPE.UPDATE",
			"DELETE": "SETTING.PAYMENT_TYPE.DELETE",
		},
		"DELIVERY_TYPE": TMapPermission{
			"CREATE": "SETTING.DELIVERY_TYPE.CREATE",
			"UPDATE": "SETTING.DELIVERY_TYPE.UPDATE",
			"DELETE": "SETTING.DELIVERY_TYPE.DELETE",
		},
		"CITY": TMapPermission{
			"CREATE": "CITY.CREATE",
			"UPDATE": "CITY.UPDATE",
			"DELETE": "CITY.DELETE",
		},
	},
}

var CONFIG_USER_TYPE = map[string]int{
	"FACEBOOK": 1,
	"GOOGLE":   2,
	"DEFAULT":  3,
}

const (
	AuthorizationHeader = "Authorization"
	AuthorizationType   = "Bearer"
	AuthorizationKey    = "authorization_payload"
)
