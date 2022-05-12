package activity

import "context"

type Service struct{}

func (Service) GetNear(ctx context.Context, lon, lat float64, distanceMeters int, limit uint64) ([]Activity, error) {
	return new(repo).getNear(ctx, lon, lat, distanceMeters, limit)
}
