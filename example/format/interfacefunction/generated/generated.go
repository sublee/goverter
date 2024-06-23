// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package generated

import (
	common "github.com/jmattheis/goverter/example/format/common"
	interfacefunction "github.com/jmattheis/goverter/example/format/interfacefunction"
)

func ConvertApartment(source common.DBApartment) common.APIApartment {
	var commonAPIApartment common.APIApartment
	commonAPIApartment.Position = source.Position
	commonAPIApartment.Owner = ConvertPerson(source.Owner)
	commonAPIApartment.OwnerName = source.Owner.Name
	if source.CoResident != nil {
		commonAPIApartment.CoResident = make([]common.APIPerson, len(source.CoResident))
		for i := 0; i < len(source.CoResident); i++ {
			commonAPIApartment.CoResident[i] = ConvertPerson(source.CoResident[i])
		}
	}
	return commonAPIApartment
}
func ConvertPerson(source common.DBPerson) common.APIPerson {
	var commonAPIPerson common.APIPerson
	commonAPIPerson.ID = source.ID
	commonAPIPerson.MiddleName = interfacefunction.SQLStringToPString(source.MiddleName)
	pString := source.Name
	commonAPIPerson.FirstName = &pString
	if source.Friends != nil {
		commonAPIPerson.Friends = make([]common.APIPerson, len(source.Friends))
		for i := 0; i < len(source.Friends); i++ {
			commonAPIPerson.Friends[i] = ConvertPerson(source.Friends[i])
		}
	}
	commonAPIPerson.Info = commonInfoToCommonInfo(source.Info)
	return commonAPIPerson
}
func commonInfoToCommonInfo(source common.Info) common.Info {
	var commonInfo common.Info
	commonInfo.Birthplace = source.Birthplace
	return commonInfo
}