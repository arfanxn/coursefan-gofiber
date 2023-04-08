package enums

// requests.QueryFilter.Operator
const (
	QueryFilterOperatorEquals      string = "eq"
	QueryFilterOperatorOrEquals    string = "oeq"
	QueryFilterOperatorNotEquals   string = "neq"
	QueryFilterOperatorOrNotEquals string = "oneq"

	QueryFilterOperatorContains      string = "ct"
	QueryFilterOperatorOrContains    string = "oct"
	QueryFilterOperatorNotContains   string = "nct"
	QueryFilterOperatorOrNotContains string = "onct"

	QueryFilterOperatorStartsWith      string = "sw"
	QueryFilterOperatorOrStartsWith    string = "osw"
	QueryFilterOperatorNotStartsWith   string = "nsw"
	QueryFilterOperatorOrNotStartsWith string = "onsw"

	QueryFilterOperatorEndsWith      string = "ew"
	QueryFilterOperatorOrEndsWith    string = "oew"
	QueryFilterOperatorNotEndsWith   string = "new"
	QueryFilterOperatorOrNotEndsWith string = "onew"

	QueryFilterOperatorGreaterThan   string = "gt"
	QueryFilterOperatorOrGreaterThan string = "ogt"

	QueryFilterOperatorGreaterThanEqual   string = "gte"
	QueryFilterOperatorOrGreaterThanEqual string = "ogte"

	QueryFilterOperatorLesserThan   string = "lt"
	QueryFilterOperatorOrLesserThan string = "olt"

	QueryFilterOperatorLesserThanEqual   string = "lte"
	QueryFilterOperatorOrLesserThanEqual string = "olte"

	QueryFilterOperatorIn      string = "in"
	QueryFilterOperatorOrIn    string = "oin"
	QueryFilterOperatorNotIn   string = "nin"
	QueryFilterOperatorOrNotIn string = "onin"

	QueryFilterOperatorBetween      string = "bt"
	QueryFilterOperatorOrBetween    string = "obt"
	QueryFilterOperatorNotBetween   string = "nbt"
	QueryFilterOperatorOrNotBetween string = "onbt"
)

// QueryFilterOperators returns slice of requests.QueryFilter.Operator enums
func QueryFilterOperators() []string {
	return []string{
		QueryFilterOperatorEquals,
		QueryFilterOperatorOrEquals,
		QueryFilterOperatorNotEquals,
		QueryFilterOperatorOrNotEquals,

		QueryFilterOperatorContains,
		QueryFilterOperatorOrContains,
		QueryFilterOperatorNotContains,
		QueryFilterOperatorOrNotContains,

		QueryFilterOperatorStartsWith,
		QueryFilterOperatorOrStartsWith,
		QueryFilterOperatorNotStartsWith,
		QueryFilterOperatorOrNotStartsWith,

		QueryFilterOperatorEndsWith,
		QueryFilterOperatorOrEndsWith,
		QueryFilterOperatorNotEndsWith,
		QueryFilterOperatorOrNotEndsWith,

		QueryFilterOperatorGreaterThan,
		QueryFilterOperatorOrGreaterThan,

		QueryFilterOperatorGreaterThanEqual,
		QueryFilterOperatorOrGreaterThanEqual,

		QueryFilterOperatorLesserThan,
		QueryFilterOperatorOrLesserThan,

		QueryFilterOperatorLesserThanEqual,
		QueryFilterOperatorOrLesserThanEqual,

		QueryFilterOperatorIn,
		QueryFilterOperatorOrIn,
		QueryFilterOperatorNotIn,
		QueryFilterOperatorOrNotIn,

		QueryFilterOperatorBetween,
		QueryFilterOperatorOrBetween,
		QueryFilterOperatorNotBetween,
		QueryFilterOperatorOrNotBetween,
	}
}
