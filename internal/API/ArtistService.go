package api

import (
	"encoding/json"
	"fmt"
	"groupie/internal/types"
	"net/http"
)

const (
	ArtistAPI   = "https://groupietrackers.herokuapp.com/api/artists"
	DatesAPI    = "https://groupietrackers.herokuapp.com/api/dates"
	RelationAPI = "https://groupietrackers.herokuapp.com/api/relation"
	LocationAPI = "https://groupietrackers.herokuapp.com/api/locations"
)

func FillArtists() ([]types.Artists, error) {
	var Artists []types.Artists

	respArtist, err := FetchJSON(ArtistAPI)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(respArtist, &Artists); err != nil {
		return nil, fmt.Errorf("проблема с декодированием артистов %w", err)
	}

	respRelation, err := FetchJSON(RelationAPI)
	if err != nil {
		return nil, err
	}

	var relation types.Relations
	if err := json.Unmarshal(respRelation, &relation); err != nil {
		return nil, fmt.Errorf("проблема с декодированием Relation %w", err)
	}

	var location types.Location
	respLocation, err := FetchJSON(LocationAPI)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(respLocation, &location); err != nil {
		return nil, fmt.Errorf("проблема с декодированием Location %w", err)
	}

	for i := 0; i < len(Artists); i++ { // Итерируемся и присваиваем значения незаполненным полям артистов
		Artists[i].Relations = relation.Index[i].Relation.DatesLocations
		Artists[i].Locations = location.Index[i].Location
	}
	return Artists, nil
}

func FillById(id string) (types.Artists, error) {
	var artist types.Artists //

	respArtist, err := FetchJSON(ArtistAPI + ("/" + id))
	if err != nil {
		return artist, fmt.Errorf("Не удалось запросить артиста по ID %s: %w", id, err)
	}

	if err := json.Unmarshal(respArtist, &artist); err != nil {
		return artist, fmt.Errorf("проблема с декодированием по ID артиста %w", err)
	}

	var artistRel types.Relation

	res, err := http.Get(RelationAPI + ("/" + id))
	if err != nil {
		return artist, fmt.Errorf("ошибка при запросе связей для артиста с ID %s: %w", id, err)
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&artistRel); err != nil {
		return artist, fmt.Errorf("ошибка при декодировании связей для артиста с ID %s: %w", id, err)
	}

	artist.Relations = artistRel.DatesLocations
	return artist, nil
}
