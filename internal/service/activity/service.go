package activity

import "context"

type Service struct{}

func (Service) GetNear(
	ctx context.Context,
	lon, lat float64,
	distanceMeters float64,
	limit uint64,
	exclude []int64,
) ([]Activity, error) {
	return new(repo).getNear(ctx, lon, lat, distanceMeters, limit, exclude)
}

func (Service) Get(ctx context.Context, id int64) (Activity, error) {
	return new(repo).get(ctx, id)
}
