package reaction

import "testing"

func Test_calcNewRanks(t *testing.T) {
	type args struct {
		userRank     float64
		activityRank float64
		k            float64
	}
	tests := []struct {
		name                string
		args                args
		wantNewUserRank     float64
		wantNewActivityRank float64
	}{
		{
			name: "1",
			args: args{
				userRank:     1,
				activityRank: -1,
				k:            0.1,
			},
			wantNewUserRank:     0.8,
			wantNewActivityRank: -0.8,
		}, {
			name: "1",
			args: args{
				userRank:     1,
				activityRank: -1,
				k:            0.5,
			},
			wantNewUserRank:     0,
			wantNewActivityRank: 0,
		}, {
			name: "1",
			args: args{
				userRank:     1,
				activityRank: 0,
				k:            0.1,
			},
			wantNewUserRank:     0.9,
			wantNewActivityRank: 0.1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewUserRank, gotNewActivityRank := calcNewRanks(tt.args.userRank, tt.args.activityRank, tt.args.k)
			if gotNewUserRank != tt.wantNewUserRank {
				t.Errorf("calcNewRanks() gotNewUserRank = %v, want %v", gotNewUserRank, tt.wantNewUserRank)
			}
			if gotNewActivityRank != tt.wantNewActivityRank {
				t.Errorf("calcNewRanks() gotNewActivityRank = %v, want %v", gotNewActivityRank, tt.wantNewActivityRank)
			}
		})
	}
}
