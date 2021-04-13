// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountPlatformsColumns holds the columns for the "account_platforms" table.
	AccountPlatformsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uid", Type: field.TypeInt64},
		{Name: "platform_id", Type: field.TypeString, Size: 30, Default: ""},
		{Name: "platform_token", Type: field.TypeString, Size: 32, Default: ""},
		{Name: "nickname", Type: field.TypeString, Size: 60, Default: ""},
		{Name: "avatar", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "type", Type: field.TypeInt8},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// AccountPlatformsTable holds the schema information for the "account_platforms" table.
	AccountPlatformsTable = &schema.Table{
		Name:        "account_platforms",
		Columns:     AccountPlatformsColumns,
		PrimaryKey:  []*schema.Column{AccountPlatformsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "accountplatform_uid_platform_id",
				Unique:  false,
				Columns: []*schema.Column{AccountPlatformsColumns[1], AccountPlatformsColumns[2]},
			},
		},
	}
	// AccountUsersColumns holds the columns for the "account_users" table.
	AccountUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Size: 30, Default: ""},
		{Name: "phone", Type: field.TypeString, Size: 15, Default: ""},
		{Name: "password", Type: field.TypeString, Size: 32, Default: ""},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt8},
	}
	// AccountUsersTable holds the schema information for the "account_users" table.
	AccountUsersTable = &schema.Table{
		Name:        "account_users",
		Columns:     AccountUsersColumns,
		PrimaryKey:  []*schema.Column{AccountUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "accountuser_email_phone",
				Unique:  false,
				Columns: []*schema.Column{AccountUsersColumns[1], AccountUsersColumns[2]},
			},
		},
	}
	// AuthItemsColumns holds the columns for the "auth_items" table.
	AuthItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeInt8},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// AuthItemsTable holds the schema information for the "auth_items" table.
	AuthItemsTable = &schema.Table{
		Name:        "auth_items",
		Columns:     AuthItemsColumns,
		PrimaryKey:  []*schema.Column{AuthItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// AuthMsColumns holds the columns for the "auth_ms" table.
	AuthMsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeInt8},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// AuthMsTable holds the schema information for the "auth_ms" table.
	AuthMsTable = &schema.Table{
		Name:        "auth_ms",
		Columns:     AuthMsColumns,
		PrimaryKey:  []*schema.Column{AuthMsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// AuthMsMenusColumns holds the columns for the "auth_ms_menus" table.
	AuthMsMenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeInt8},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// AuthMsMenusTable holds the schema information for the "auth_ms_menus" table.
	AuthMsMenusTable = &schema.Table{
		Name:        "auth_ms_menus",
		Columns:     AuthMsMenusColumns,
		PrimaryKey:  []*schema.Column{AuthMsMenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// AuthRolesColumns holds the columns for the "auth_roles" table.
	AuthRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeInt8},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// AuthRolesTable holds the schema information for the "auth_roles" table.
	AuthRolesTable = &schema.Table{
		Name:        "auth_roles",
		Columns:     AuthRolesColumns,
		PrimaryKey:  []*schema.Column{AuthRolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// AuthRoleStaffsColumns holds the columns for the "auth_role_staffs" table.
	AuthRoleStaffsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeInt8},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// AuthRoleStaffsTable holds the schema information for the "auth_role_staffs" table.
	AuthRoleStaffsTable = &schema.Table{
		Name:        "auth_role_staffs",
		Columns:     AuthRoleStaffsColumns,
		PrimaryKey:  []*schema.Column{AuthRoleStaffsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProductAttrsColumns holds the columns for the "product_attrs" table.
	ProductAttrsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeInt8},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// ProductAttrsTable holds the schema information for the "product_attrs" table.
	ProductAttrsTable = &schema.Table{
		Name:        "product_attrs",
		Columns:     ProductAttrsColumns,
		PrimaryKey:  []*schema.Column{ProductAttrsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProductAttrValuesColumns holds the columns for the "product_attr_values" table.
	ProductAttrValuesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeInt8},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// ProductAttrValuesTable holds the schema information for the "product_attr_values" table.
	ProductAttrValuesTable = &schema.Table{
		Name:        "product_attr_values",
		Columns:     ProductAttrValuesColumns,
		PrimaryKey:  []*schema.Column{ProductAttrValuesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProductBrandsColumns holds the columns for the "product_brands" table.
	ProductBrandsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "desc", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "logo_url", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "status", Type: field.TypeInt8},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// ProductBrandsTable holds the schema information for the "product_brands" table.
	ProductBrandsTable = &schema.Table{
		Name:        "product_brands",
		Columns:     ProductBrandsColumns,
		PrimaryKey:  []*schema.Column{ProductBrandsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProductCategoriesColumns holds the columns for the "product_categories" table.
	ProductCategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "pid", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "desc", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "pic_url", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "path", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "status", Type: field.TypeInt8},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// ProductCategoriesTable holds the schema information for the "product_categories" table.
	ProductCategoriesTable = &schema.Table{
		Name:        "product_categories",
		Columns:     ProductCategoriesColumns,
		PrimaryKey:  []*schema.Column{ProductCategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProductSkusColumns holds the columns for the "product_skus" table.
	ProductSkusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "spu_id", Type: field.TypeInt64},
		{Name: "attrs", Type: field.TypeString, Size: 2147483647},
		{Name: "banner_url", Type: field.TypeString, Size: 2147483647},
		{Name: "main_url", Type: field.TypeString, Size: 2147483647},
		{Name: "price_fee", Type: field.TypeInt64},
		{Name: "price_scale", Type: field.TypeInt8},
		{Name: "market_price_fee", Type: field.TypeInt64},
		{Name: "market_price_scale", Type: field.TypeInt8},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt8},
	}
	// ProductSkusTable holds the schema information for the "product_skus" table.
	ProductSkusTable = &schema.Table{
		Name:        "product_skus",
		Columns:     ProductSkusColumns,
		PrimaryKey:  []*schema.Column{ProductSkusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProductSkuStocksColumns holds the columns for the "product_sku_stocks" table.
	ProductSkuStocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt8},
	}
	// ProductSkuStocksTable holds the schema information for the "product_sku_stocks" table.
	ProductSkuStocksTable = &schema.Table{
		Name:        "product_sku_stocks",
		Columns:     ProductSkuStocksColumns,
		PrimaryKey:  []*schema.Column{ProductSkuStocksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProductSpusColumns holds the columns for the "product_spus" table.
	ProductSpusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "brand_id", Type: field.TypeInt64},
		{Name: "category_id", Type: field.TypeInt64},
		{Name: "name", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "desc", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "selling_point", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "unit", Type: field.TypeString, Size: 25, Default: ""},
		{Name: "banner_url", Type: field.TypeString, Size: 2147483647},
		{Name: "main_url", Type: field.TypeString, Size: 2147483647},
		{Name: "price_fee", Type: field.TypeInt64},
		{Name: "price_scale", Type: field.TypeInt8},
		{Name: "market_price_fee", Type: field.TypeInt64},
		{Name: "market_price_scale", Type: field.TypeInt8},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt8},
	}
	// ProductSpusTable holds the schema information for the "product_spus" table.
	ProductSpusTable = &schema.Table{
		Name:        "product_spus",
		Columns:     ProductSpusColumns,
		PrimaryKey:  []*schema.Column{ProductSpusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProductSpuSkuAttrMapsColumns holds the columns for the "product_spu_sku_attr_maps" table.
	ProductSpuSkuAttrMapsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeInt8},
	}
	// ProductSpuSkuAttrMapsTable holds the schema information for the "product_spu_sku_attr_maps" table.
	ProductSpuSkuAttrMapsTable = &schema.Table{
		Name:        "product_spu_sku_attr_maps",
		Columns:     ProductSpuSkuAttrMapsColumns,
		PrimaryKey:  []*schema.Column{ProductSpuSkuAttrMapsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// SkrMembersColumns holds the columns for the "skr_members" table.
	SkrMembersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uid", Type: field.TypeInt64},
		{Name: "nickname", Type: field.TypeString, Size: 30, Default: ""},
		{Name: "avatar", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "gender", Type: field.TypeInt8, Default: 1},
		{Name: "role", Type: field.TypeInt8},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// SkrMembersTable holds the schema information for the "skr_members" table.
	SkrMembersTable = &schema.Table{
		Name:        "skr_members",
		Columns:     SkrMembersColumns,
		PrimaryKey:  []*schema.Column{SkrMembersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "skrmember_uid",
				Unique:  false,
				Columns: []*schema.Column{SkrMembersColumns[1]},
			},
		},
	}
	// StaffInfosColumns holds the columns for the "staff_infos" table.
	StaffInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uid", Type: field.TypeInt64},
		{Name: "email", Type: field.TypeString, Size: 30, Default: ""},
		{Name: "phone", Type: field.TypeString, Size: 15, Default: ""},
		{Name: "password", Type: field.TypeString, Size: 32, Default: ""},
		{Name: "name", Type: field.TypeString, Size: 30, Default: ""},
		{Name: "avatar", Type: field.TypeString, Size: 255, Default: ""},
		{Name: "gender", Type: field.TypeInt8, Default: 1},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// StaffInfosTable holds the schema information for the "staff_infos" table.
	StaffInfosTable = &schema.Table{
		Name:        "staff_infos",
		Columns:     StaffInfosColumns,
		PrimaryKey:  []*schema.Column{StaffInfosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "staffinfo_email_phone",
				Unique:  false,
				Columns: []*schema.Column{StaffInfosColumns[2], StaffInfosColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountPlatformsTable,
		AccountUsersTable,
		AuthItemsTable,
		AuthMsTable,
		AuthMsMenusTable,
		AuthRolesTable,
		AuthRoleStaffsTable,
		ProductAttrsTable,
		ProductAttrValuesTable,
		ProductBrandsTable,
		ProductCategoriesTable,
		ProductSkusTable,
		ProductSkuStocksTable,
		ProductSpusTable,
		ProductSpuSkuAttrMapsTable,
		SkrMembersTable,
		StaffInfosTable,
	}
)

func init() {
}
