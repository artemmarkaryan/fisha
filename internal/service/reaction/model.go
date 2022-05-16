package reaction

const (
	ReactionTypeUnknown = "UNKNOWN"
	ReactionTypeLike    = "LIKE"
	ReactionTypeDislike = "DISLIKE"
)

var ValidReactionTypes = []string{ReactionTypeLike, ReactionTypeDislike}

func ValidateReactionType(t string) bool {
	for _, rt := range ValidReactionTypes {
		if t == rt {
			return true
		}
	}

	return false
}

type Reaction struct {
	UserID     int64  `json:"user_id"`
	ActivityID int64  `json:"activity_id"`
	Reaction   string `json:"reaction"`
}
