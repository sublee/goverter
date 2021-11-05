// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package generated

import mismatched "github.com/jmattheis/goverter/example/mismatched"

type ConverterImpl struct{}

func (c *ConverterImpl) Convert(source mismatched.DBCustomers) mismatched.APICustomers {
	mismatchedAPICustomers := make(mismatched.APICustomers, len(source))
	for i := 0; i < len(source); i++ {
		mismatchedAPICustomers[i] = c.ToApiCustomer(source[i])
	}
	return mismatchedAPICustomers
}
func (c *ConverterImpl) ToAPIAddress(source mismatched.DBAddress) mismatched.APIAddress {
	var mismatchedAPIAddress mismatched.APIAddress
	mismatchedAPIAddress.Suburb = source.Suburb
	mismatchedAPIAddress.Postcode = source.Postcode
	return mismatchedAPIAddress
}
func (c *ConverterImpl) ToApiCustomer(source mismatched.DBCustomer) mismatched.APICustomer {
	var mismatchedAPICustomer mismatched.APICustomer
	mismatchedAPICustomer.APIPerson = c.mismatchedDBPersonToPMismatchedAPIPerson(source.DBPerson)
	mismatchedAPICustomer.APIAddress = mismatched.AddressOrDefault(c, source.DBAddress)
	return mismatchedAPICustomer
}
func (c *ConverterImpl) mismatchedDBPersonToMismatchedAPIPerson(source mismatched.DBPerson) mismatched.APIPerson {
	var mismatchedAPIPerson mismatched.APIPerson
	mismatchedAPIPerson.First = source.First
	mismatchedAPIPerson.Last = source.Last
	return mismatchedAPIPerson
}
func (c *ConverterImpl) mismatchedDBPersonToPMismatchedAPIPerson(source mismatched.DBPerson) *mismatched.APIPerson {
	mismatchedAPIPerson := c.mismatchedDBPersonToMismatchedAPIPerson(source)
	return &mismatchedAPIPerson
}