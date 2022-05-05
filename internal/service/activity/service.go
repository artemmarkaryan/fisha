package activity

import "context"

type Service struct{}

func (Service) GetNear(ctx context.Context, lon, lat float32, distanceMeters int) ([]Activity, error) {
	return new(repo).getNear(ctx, lon, lat, distanceMeters)
}
