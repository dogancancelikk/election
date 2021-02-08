package services

import (
	"github.com/dogancancelikk/election/domain/parties"
	"github.com/dogancancelikk/election/utils/errors"
)

func CreateParty(party parties.Party) (*parties.Party, *errors.RestError) {

	if err := party.Save(); err != nil {
		return nil, err
	}

	return &party, nil
}

func GetParties() ([]parties.Party, *errors.RestError) {
	var party parties.Party
	result, err := party.GetParties()
	if err != nil {
		return nil, err
	}
	return result, nil
}
