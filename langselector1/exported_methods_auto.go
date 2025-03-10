// Code generated by "dbusutil-gen em -type LangSelector"; DO NOT EDIT.

package langselector

import (
	"github.com/linuxdeepin/go-lib/dbusutil"
)

func (v *LangSelector) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name:   "AddLocale",
			Fn:     v.AddLocale,
			InArgs: []string{"locale"},
		},
		{
			Name:   "DeleteLocale",
			Fn:     v.DeleteLocale,
			InArgs: []string{"locale"},
		},
		{
			Name:    "GetLanguageSupportPackages",
			Fn:      v.GetLanguageSupportPackages,
			InArgs:  []string{"locale"},
			OutArgs: []string{"packages"},
		},
		{
			Name:    "GetLocaleDescription",
			Fn:      v.GetLocaleDescription,
			InArgs:  []string{"locale"},
			OutArgs: []string{"description"},
		},
		{
			Name:    "GetLocaleList",
			Fn:      v.GetLocaleList,
			OutArgs: []string{"locales"},
		},
		{
			Name:    "GetLocaleRegion",
			Fn:      v.GetLocaleRegion,
			OutArgs: []string{"region"},
		},
		{
			Name: "Reset",
			Fn:   v.Reset,
		},
		{
			Name:   "SetLocale",
			Fn:     v.SetLocale,
			InArgs: []string{"locale"},
		},
		{
			Name:   "SetLocaleRegion",
			Fn:     v.SetLocaleRegion,
			InArgs: []string{"locale"},
		},
	}
}
